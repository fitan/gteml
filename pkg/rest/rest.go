package rest

import (
	"fmt"
	"github.com/fitan/magic/pkg/utils/slices"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type Hook interface {
	Before(ctx Context, req interface{}, body interface{}) error
	After(ctx Context, data interface{}, err error) (interface{}, error)
}

type GetListHook Hook
type GetOneHook Hook
type CreateHook Hook
type UpdateHook Hook
type DeleteHook Hook
type DeleteManyCycle Hook
type RelationGetHook Hook
type RelationCreateHook Hook
type RelationUpdateHook Hook

type Restful interface {
	Objer
	FieldConfer
	Wrap(ctx *gin.Context, data interface{}, err error)
	GetListScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetList(ctx Context, objs interface{}, count *int64) (interface{}, error)
	GetOneScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetOne(ctx Context, obj interface{}) (interface{}, error)
	CreateScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Create(ctx Context) (interface{}, error)
	UpdateScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Update(ctx Context) (interface{}, error)
	DeleteScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Delete(ctx Context) (interface{}, error)
	DeleteManyScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	DeleteMany(ctx Context) (interface{}, error)
	GetFieldScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetField(ctx Context) (interface{}, error)
	GetFieldsScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetFields(ctx Context) (interface{}, error)
	RelationGet(ctx Context, relationName string, relationObjs interface{}, count *int64) (interface{}, error)
	RelationCreate(ctx Context, relationName string) (interface{}, error)
	RelationUpdate(ctx Context, relationName string) (interface{}, error)
	RelationsCreate(ctx Context) (interface{}, error)
	RelationsUpdate(ctx Context) (interface{}, error)
}

type GetListRes struct {
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

type BaseRest struct {
	DB *gorm.DB
	Objer
}

func NewBaseRest(DB *gorm.DB, objer Objer) *BaseRest {
	return &BaseRest{DB: DB, Objer: objer}
}

func (b *BaseRest) Wrap(ctx *gin.Context, data interface{}, err error) {
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
	return
}

func (b *BaseRest) GetListScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getList GetList
	err = ctx.GinCtx().BindQuery(&getList)
	if err != nil {
		return
	}
	scopes, err = getList.Scopes()
	if err != nil {
		return
	}

	if slices.StringContains(getList.Includes, "_all") {
		for _, r := range b.RelationsField() {
			v := r.GetTableName()
			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
				return db.Preload(v)
			})
		}
	} else {
		for _, i := range getList.Includes {
			relations, ok := b.RelationsField()[i]
			if !ok {
				return nil, fmt.Errorf("includes err: %v not exist", i)
			}
			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
				return db.Preload(relations.GetTableName())
			})
		}
	}
	return
}

func (b *BaseRest) GetList(ctx Context, objs interface{}, count *int64) (interface{}, error) {
	scopes, err := b.GetListScopes(ctx)
	if err != nil {
		return nil, err
	}

	if objs == nil {
		objs = b.GetFindObj()
	}
	db := ctx.DB().Model(b.GetModelObj())
	if count != nil {
		db = db.Count(count)
	}
	err = db.Scopes(scopes...).Find(objs).Error
	return objs, err
}

func (b *BaseRest) GetOneScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GinCtx().BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err = getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	return
}

func (b *BaseRest) GetOne(ctx Context, obj interface{}) (interface{}, error) {
	scopes, err := b.GetOneScopes(ctx)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		obj = b.GetFirstObj()
	}
	err = ctx.DB().Model(b.GetModelObj()).Scopes(scopes...).First(obj).Error
	return obj, err
}

func (b *BaseRest) CreateScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	s, o := b.Objer.CreateField()
	// 禁止关联创建，防止误创建。如果要创建使用relattioncreate
	o = append(o, clause.Associations)
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Select(s).Omit(o...)
	})
	return
}

func (b *BaseRest) Create(ctx Context) (interface{}, error) {
	obj := b.GetModelObj()
	err := ctx.GinCtx().BindJSON(obj)
	if err != nil {
		return nil, err
	}

	scopes, err := b.CreateScopes(ctx)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().Scopes(scopes...).Create(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) UpdateScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GinCtx().BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err = getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	s, o := b.Objer.UpdateField()
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Select(s).Omit(o...)
	})

	return
}

func (b *BaseRest) Update(ctx Context) (interface{}, error) {

	data := b.GetModelObj()
	err := ctx.GinCtx().BindJSON(data)
	if err != nil {
		return nil, err
	}

	scopes, err := b.UpdateScopes(ctx)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().Model(b.GetModelObj()).Scopes(scopes...).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

//func (b *BaseRest) UpdateMany(ctx Context) (interface{}, error) {
//	var getManyByIds GetManyByIds
//	err := ctx.GinCtx().BindQuery(&getManyByIds)
//	if err != nil {
//		return nil, err
//	}
//
//	scopes, err := getManyByIds.Scopes()
//	if err != nil {
//		return nil, err
//	}
//
//	data := b.GetModelObj()
//	err = ctx.GinCtx().BindJSON(data)
//	if err != nil {
//		return nil, err
//	}
//
//	err = ctx.DB().Model(b.GetModelObj()).Scopes(scopes...).Save(data).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return "ok", nil
//}

func (b *BaseRest) DeleteScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GinCtx().BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err = getOneById.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) Delete(ctx Context) (interface{}, error) {
	scopes, err := b.DeleteScopes(ctx)
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()
	err = ctx.DB().Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) DeleteManyScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getManyByIds GetManyByIds
	err = ctx.GinCtx().BindQuery(&getManyByIds)
	if err != nil {
		return nil, err
	}

	scopes, err = getManyByIds.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) DeleteMany(ctx Context) (interface{}, error) {
	scopes, err := b.DeleteManyScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetModelObjs()
	err = ctx.DB().Scopes(scopes...).Delete(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *BaseRest) GetFieldScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var queryField QueryField
	err = ctx.GinCtx().ShouldBindUri(&queryField)
	if err != nil {
		return nil, err
	}
	err = ctx.GinCtx().ShouldBindQuery(&queryField)
	if err != nil {
		return nil, err
	}

	scopes, err = queryField.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) GetField(ctx Context) (interface{}, error) {
	scopes, err := b.GetFieldScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = ctx.DB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *BaseRest) GetFieldsScopes(ctx Context) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var queryFields QueryFields
	err = ctx.GinCtx().ShouldBind(&queryFields)
	if err != nil {
		return nil, err
	}

	scopes, err = queryFields.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) GetFields(ctx Context) (interface{}, error) {
	scopes, err := b.GetFieldsScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = ctx.DB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

type RelationCreate struct {
	Id     int64    `json:"id" uri:"id"`
	Fields []string `json:"fields" form:"_fields"`
}

type RelationsCommon struct {
	Id           int64  `uri:"id"`
	RelationName string `uri:"relationName"`
}

// 一对多查询
func (b *BaseRest) RelationGet(ctx Context, relationName string, relationObjs interface{}, count *int64) (interface{}, error) {

	var getRelation RelationGet
	err := ctx.GinCtx().ShouldBindUri(&getRelation)
	if err != nil {
		return nil, err
	}
	err = ctx.GinCtx().ShouldBind(&getRelation)
	if err != nil {
		return nil, err
	}

	scopes, err := getRelation.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()

	err = ctx.DB().Find(obj, getRelation.Id).Error
	if err != nil {
		return nil, err
	}

	if relationName == "" {
		relationName = getRelation.RelationName
	}

	rc, ok := b.RelationsField()[relationName]
	if !ok {
		return nil, fmt.Errorf("%v relation field permission denied", getRelation.RelationName)
	}

	tableName := rc.GetTableName()

	if count != nil {
		*count = ctx.DB().Model(obj).Scopes(scopes...).Association(tableName).Count()
	}
	if relationObjs == nil {
		relationObjs = rc.GetFindObj()
	}

	err = ctx.DB().Model(obj).Scopes(scopes...).Association(tableName).Find(relationObjs)
	if err != nil {
		return nil, err
	}

	return relationObjs, nil
}

func (b *BaseRest) RelationCreate(ctx Context, relationName string) (interface{}, error) {
	var relationCommon RelationsCommon
	err := ctx.GinCtx().ShouldBindUri(&relationCommon)
	if err != nil {
		return nil, err
	}

	obj := b.GetFirstObj()
	err = ctx.GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().First(obj, relationCommon.Id).Error
	if err != nil {
		return nil, err
	}

	if relationName == "" {
		relationName = relationCommon.RelationName
	}

	rf, ok := b.RelationsField()[relationName]
	if !ok {
		return nil, fmt.Errorf("relationName does not exist: %v", relationName)
	}
	err = ctx.DB().Model(obj).Select(rf.GetTableName()).Create(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) RelationUpdate(ctx Context, relationName string) (interface{}, error) {
	var relationCommon RelationsCommon
	err := ctx.GinCtx().ShouldBindUri(&relationCommon)
	if err != nil {
		return nil, err
	}

	obj := b.GetFirstObj()
	err = ctx.GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().First(obj, relationCommon.Id).Error
	if err != nil {
		return nil, err
	}

	if relationName == "" {
		relationName = relationCommon.RelationName
	}

	rf, ok := b.RelationsField()[relationName]
	if !ok {
		return nil, fmt.Errorf("relationName does not exist: %v", relationName)
	}

	tableName := rf.GetTableName()
	s, o := b.Objer.UpdateField()

	s = append(s, tableName)
	rfS, rfO := rf.UpdateField()
	for _, v := range rfS {
		s = append(s, tableName+"."+v)
	}
	for _, v := range rfO {
		o = append(o, tableName+"."+v)
	}

	err = ctx.DB().Session(&gorm.Session{FullSaveAssociations: true}).Select(s).Omit(o...).Updates(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil

}

func (b *BaseRest) RelationsCreate(ctx Context) (interface{}, error) {
	var relationCreate RelationCreate
	err := ctx.GinCtx().ShouldBindUri(&relationCreate)
	if err != nil {
		return nil, err
	}

	err = ctx.GinCtx().ShouldBindQuery(&relationCreate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().First(obj, relationCreate.Id).Error
	if err != nil {
		return nil, err
	}

	db := ctx.DB()

	if len(relationCreate.Fields) == 0 {
		return nil, errors.New("Field does not exist: _fields")
	} else {
		for _, f := range relationCreate.Fields {
			if rf, ok := b.RelationsField()[f]; ok {
				db = db.Select(rf.GetTableName())
				//s = append(s,rf.GetTableName())
			}
		}
	}

	err = db.Create(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (b *BaseRest) RelationsUpdate(ctx Context) (interface{}, error) {
	var relationUpdate RelationCreate
	err := ctx.GinCtx().ShouldBindUri(&relationUpdate)
	if err != nil {
		return nil, err
	}

	err = ctx.GinCtx().ShouldBindQuery(&relationUpdate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.DB().First(obj, relationUpdate.Id).Error
	if err != nil {
		return nil, err
	}

	s, o := b.Objer.UpdateField()

	if len(relationUpdate.Fields) == 0 {
		return nil, errors.New("Field does not exist: _fields")
	} else {
		for _, f := range relationUpdate.Fields {
			if rf, ok := b.RelationsField()[f]; ok {
				tableName := rf.GetTableName()
				s = append(s, tableName)

				rfS, rfO := rf.UpdateField()
				for _, v := range rfS {
					s = append(s, tableName+"."+v)
				}

				for _, v := range rfO {
					o = append(o, tableName+"."+v)
				}

			}
		}
	}

	err = ctx.DB().Session(&gorm.Session{FullSaveAssociations: true}).Select(s).Omit(o...).Updates(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
