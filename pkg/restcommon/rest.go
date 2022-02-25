package restcommon

import (
	"fmt"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/pkg/utils/slices"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type Hook interface {
	Before(ctx types.ServiceCore, req interface{}, body interface{}) error
	After(ctx types.ServiceCore, data interface{}, err error) (interface{}, error)
}

type Restful interface {
	Objer
	FieldConfer
	Wrap(ctx *gin.Context, data interface{}, err error)
	GetListScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetList(ctx types.ServiceCore, objs interface{}, count *int64) (interface{}, error)
	GetOneScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetOne(ctx types.ServiceCore, obj interface{}) (interface{}, error)
	CreateScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Create(ctx types.ServiceCore) (interface{}, error)
	UpdateScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Update(ctx types.ServiceCore) (interface{}, error)
	DeleteScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	Delete(ctx types.ServiceCore) (interface{}, error)
	DeleteManyScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	DeleteMany(ctx types.ServiceCore) (interface{}, error)
	GetFieldScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetField(ctx types.ServiceCore) (interface{}, error)
	GetFieldsScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error)
	GetFields(ctx types.ServiceCore) (interface{}, error)
	RelationGet(ctx types.ServiceCore, relationName string, relationObjs interface{}, count *int64) (interface{}, error)
	RelationCreate(ctx types.ServiceCore, relationName string) (interface{}, error)
	RelationUpdate(ctx types.ServiceCore, relationName string) (interface{}, error)
	RelationsCreate(ctx types.ServiceCore) (interface{}, error)
	RelationsUpdate(ctx types.ServiceCore) (interface{}, error)
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

func (b *BaseRest) GetListScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getList GetList
	err = ctx.GetGinX().GinCtx().ShouldBindQuery(&getList)
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

func (b *BaseRest) GetList(ctx types.ServiceCore, objs interface{}, count *int64) (interface{}, error) {
	scopes, err := b.GetListScopes(ctx)
	if err != nil {
		return nil, err
	}

	if objs == nil {
		objs = b.GetFindObj()
	}
	db := ctx.GetDao().DB().Model(b.GetModelObj())
	if count != nil {
		db = db.Count(count)
	}
	err = db.Scopes(scopes...).Find(objs).Error
	return objs, err
}

func (b *BaseRest) GetOneScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GetGinX().GinCtx().ShouldBindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err = getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	return
}

func (b *BaseRest) GetOne(ctx types.ServiceCore, obj interface{}) (interface{}, error) {
	scopes, err := b.GetOneScopes(ctx)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		obj = b.GetFirstObj()
	}
	err = ctx.GetDao().DB().Model(b.GetModelObj()).Scopes(scopes...).First(obj).Error
	return obj, err
}

func (b *BaseRest) CreateScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	s, o := b.Objer.CreateField()
	// 禁止关联创建，防止误创建。如果要创建使用relattioncreate
	o = append(o, clause.Associations)
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Select(s).Omit(o...)
	})
	return
}

func (b *BaseRest) Create(ctx types.ServiceCore) (interface{}, error) {
	obj := b.GetModelObj()
	err := ctx.GetGinX().GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	scopes, err := b.CreateScopes(ctx)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().Scopes(scopes...).Create(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) UpdateScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GetGinX().GinCtx().ShouldBindUri(&getOneById)
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

func (b *BaseRest) Update(ctx types.ServiceCore) (interface{}, error) {

	data := b.GetModelObj()
	err := ctx.GetGinX().GinCtx().ShouldBindJSON(data)
	if err != nil {
		return nil, err
	}

	scopes, err := b.UpdateScopes(ctx)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().Model(b.GetModelObj()).Scopes(scopes...).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

//func (b *BaseRest) UpdateMany(ctx types.ServiceCore) (interface{}, error) {
//	var getManyByIds GetManyByIds
//	err := ctx.GetGinX().GinCtx().ShouldBindQuery(&getManyByIds)
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
//	err = ctx.GetGinX().GinCtx().ShouldBindJSON(data)
//	if err != nil {
//		return nil, err
//	}
//
//	err = ctx.GetDao().DB().Model(b.GetModelObj()).Scopes(scopes...).Save(data).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return "ok", nil
//}

func (b *BaseRest) DeleteScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getOneById GetOneById
	err = ctx.GetGinX().GinCtx().ShouldBindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err = getOneById.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) Delete(ctx types.ServiceCore) (interface{}, error) {
	scopes, err := b.DeleteScopes(ctx)
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()
	err = ctx.GetDao().DB().Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) DeleteManyScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var getManyByIds GetManyByIds
	err = ctx.GetGinX().GinCtx().ShouldBindQuery(&getManyByIds)
	if err != nil {
		return nil, err
	}

	scopes, err = getManyByIds.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) DeleteMany(ctx types.ServiceCore) (interface{}, error) {
	scopes, err := b.DeleteManyScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetModelObjs()
	err = ctx.GetDao().DB().Scopes(scopes...).Delete(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *BaseRest) GetFieldScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var queryField QueryField
	err = ctx.GetGinX().GinCtx().ShouldBindUri(&queryField)
	if err != nil {
		return nil, err
	}
	err = ctx.GetGinX().GinCtx().ShouldBindQuery(&queryField)
	if err != nil {
		return nil, err
	}

	scopes, err = queryField.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) GetField(ctx types.ServiceCore) (interface{}, error) {
	scopes, err := b.GetFieldScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = ctx.GetDao().DB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *BaseRest) GetFieldsScopes(ctx types.ServiceCore) (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	var queryFields QueryFields
	err = ctx.GetGinX().GinCtx().ShouldBind(&queryFields)
	if err != nil {
		return nil, err
	}

	scopes, err = queryFields.Scopes()
	if err != nil {
		return nil, err
	}
	return
}

func (b *BaseRest) GetFields(ctx types.ServiceCore) (interface{}, error) {
	scopes, err := b.GetFieldsScopes(ctx)
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = ctx.GetDao().DB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
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
func (b *BaseRest) RelationGet(ctx types.ServiceCore, relationName string, relationObjs interface{}, count *int64) (interface{}, error) {

	var getRelation RelationGet
	err := ctx.GetGinX().GinCtx().ShouldBindUri(&getRelation)
	if err != nil {
		return nil, err
	}
	err = ctx.GetGinX().GinCtx().ShouldBind(&getRelation)
	if err != nil {
		return nil, err
	}

	scopes, err := getRelation.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()

	err = ctx.GetDao().DB().Find(obj, getRelation.Id).Error
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
		*count = ctx.GetDao().DB().Model(obj).Scopes(scopes...).Association(tableName).Count()
	}
	if relationObjs == nil {
		relationObjs = rc.GetFindObj()
	}

	err = ctx.GetDao().DB().Model(obj).Scopes(scopes...).Association(tableName).Find(relationObjs)
	if err != nil {
		return nil, err
	}

	return relationObjs, nil
}

func (b *BaseRest) RelationCreate(ctx types.ServiceCore, relationName string) (interface{}, error) {
	var relationCommon RelationsCommon
	err := ctx.GetGinX().GinCtx().ShouldBindUri(&relationCommon)
	if err != nil {
		return nil, err
	}

	obj := b.GetFirstObj()
	err = ctx.GetGinX().GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().First(obj, relationCommon.Id).Error
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
	err = ctx.GetDao().DB().Model(obj).Select(rf.GetTableName()).Create(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) RelationUpdate(ctx types.ServiceCore, relationName string) (interface{}, error) {
	var relationCommon RelationsCommon
	err := ctx.GetGinX().GinCtx().ShouldBindUri(&relationCommon)
	if err != nil {
		return nil, err
	}

	obj := b.GetFirstObj()
	err = ctx.GetGinX().GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().First(obj, relationCommon.Id).Error
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

	err = ctx.GetDao().DB().Session(&gorm.Session{FullSaveAssociations: true}).Select(s).Omit(o...).Updates(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil

}

func (b *BaseRest) RelationsCreate(ctx types.ServiceCore) (interface{}, error) {
	var relationCreate RelationCreate
	err := ctx.GetGinX().GinCtx().ShouldBindUri(&relationCreate)
	if err != nil {
		return nil, err
	}

	err = ctx.GetGinX().GinCtx().ShouldBindQuery(&relationCreate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.GetGinX().GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().First(obj, relationCreate.Id).Error
	if err != nil {
		return nil, err
	}

	db := ctx.GetDao().DB()

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

func (b *BaseRest) RelationsUpdate(ctx types.ServiceCore) (interface{}, error) {
	var relationUpdate RelationCreate
	err := ctx.GetGinX().GinCtx().ShouldBindUri(&relationUpdate)
	if err != nil {
		return nil, err
	}

	err = ctx.GetGinX().GinCtx().ShouldBindQuery(&relationUpdate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.GetGinX().GinCtx().ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = ctx.GetDao().DB().First(obj, relationUpdate.Id).Error
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

	err = ctx.GetDao().DB().Session(&gorm.Session{FullSaveAssociations: true}).Select(s).Omit(o...).Updates(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
