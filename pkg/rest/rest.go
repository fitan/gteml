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

type RestfulV1 interface {
	Objer
	FieldConfer
	GetDB() *gorm.DB
	Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error))

	GetListBind(ctx *gin.Context) (GetList, error)
	GetListBefore(ctx *gin.Context, req GetList) error
	GetList(ctx *gin.Context, req GetList) (interface{}, error)
	GetListAfter(ctx *gin.Context, data interface{}, err error) (interface{}, error)

	GetOneBind(ctx *gin.Context) (GetOneById, error)
	GetOneBefore(ctx *gin.Context, req GetOneById) (interface{}, error)
	GetOne(ctx *gin.Context) (interface{}, error)
	GetOneAfter(ctx *gin.Context, data interface{}, err error) (interface{}, error)

	Create(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	UpdateMany(ctx *gin.Context) (interface{}, error)
	Delete(ctx *gin.Context) (interface{}, error)
	DeleteMany(ctx *gin.Context) (interface{}, error)

	GetField(ctx *gin.Context) (interface{}, error)
	GetFields(ctx *gin.Context) (interface{}, error)

	Relations(ctx *gin.Context) (interface{}, error)
	RelationCreate(ctx *gin.Context) (interface{}, error)
	RelationUpdate(ctx *gin.Context) (interface{}, error)
}

type Restful interface {
	Objer
	FieldConfer
	GetDB() *gorm.DB
	Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error))

	GetList(ctx *gin.Context) (interface{}, error)

	GetOne(ctx *gin.Context) (interface{}, error)
	Create(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	UpdateMany(ctx *gin.Context) (interface{}, error)
	Delete(ctx *gin.Context) (interface{}, error)
	DeleteMany(ctx *gin.Context) (interface{}, error)

	GetField(ctx *gin.Context) (interface{}, error)
	GetFields(ctx *gin.Context) (interface{}, error)

	Relations(ctx *gin.Context) (interface{}, error)
	RelationCreate(ctx *gin.Context) (interface{}, error)
	RelationUpdate(ctx *gin.Context) (interface{}, error)
}

type BaseRest struct {
	DB *gorm.DB
	Objer
}

func NewBaseRest(DB *gorm.DB, objer Objer) *BaseRest {
	return &BaseRest{DB: DB, Objer: objer}
}

func (b *BaseRest) Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error)) {
	data, err := fn(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
	return
}

func (b *BaseRest) GetDB() *gorm.DB {
	return b.DB
}

func (b *BaseRest) GetList(ctx *gin.Context) (interface{}, error) {
	var getList GetList
	err := ctx.BindQuery(&getList)
	if err != nil {
		return nil, errors.WithMessage(err, "bindQuery")
	}
	scopes, err := getList.Scopes()
	if err != nil {
		return nil, errors.WithMessage(err, "getList.Scopes")
	}

	var count int64

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

	objs := b.GetFindObj()
	err = b.GetDB().Model(b.GetModelObj()).Count(&count).Scopes(scopes...).Find(objs).Error
	return map[string]interface{}{"list": objs, "count": count}, errors.WithMessage(err, "DB find")

}

func (b *BaseRest) GetOne(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetFirstObj()
	err = b.GetDB().Model(b.GetModelObj()).Scopes(scopes...).First(obj).Error
	return obj, err
}

func (b *BaseRest) Create(ctx *gin.Context) (interface{}, error) {
	obj := b.GetModelObj()
	err := ctx.BindJSON(obj)
	if err != nil {
		return nil, err
	}

	s, o := b.Objer.CreateField()
	// 禁止关联创建，防止误创建。如果要创建使用relattioncreate
	o = append(o, clause.Associations)
	err = b.GetDB().Select(s).Omit(o...).Create(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) UpdateField() (s []string, o []string) {
	return []string{"*"}, []string{"id", "updated_at", "deleted_at", "created_at"}
}

func (b *BaseRest) Update(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	data := b.GetModelObj()
	err = ctx.BindJSON(data)
	if err != nil {
		return nil, err
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()
	s, o := b.UpdateField()
	err = b.GetDB().Model(obj).Select(s).Omit(o...).Scopes(scopes...).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return "ok", nil
}

func (b *BaseRest) UpdateMany(ctx *gin.Context) (interface{}, error) {
	var getManyByIds GetManyByIds
	err := ctx.BindQuery(&getManyByIds)
	if err != nil {
		return nil, err
	}

	scopes, err := getManyByIds.Scopes()
	if err != nil {
		return nil, err
	}

	data := b.GetModelObj()
	err = ctx.BindJSON(data)
	if err != nil {
		return nil, err
	}

	err = b.GetDB().Model(b.GetModelObj()).Scopes(scopes...).Save(data).Error
	if err != nil {
		return nil, err
	}

	return "ok", nil
}

func (b *BaseRest) Delete(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, err
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()
	err = b.GetDB().Clauses(clause.Returning{}).Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (b *BaseRest) DeleteMany(ctx *gin.Context) (interface{}, error) {
	var getManyByIds GetManyByIds
	err := ctx.BindQuery(&getManyByIds)
	if err != nil {
		return nil, err
	}

	scopes, err := getManyByIds.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()
	err = b.GetDB().Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, err
	}
	return getManyByIds.Ids, nil
}

func (b *BaseRest) GetField(ctx *gin.Context) (interface{}, error) {
	var queryField QueryField
	err := ctx.ShouldBindUri(&queryField)
	if err != nil {
		return nil, err
	}
	err = ctx.ShouldBindQuery(&queryField)
	if err != nil {
		return nil, err
	}

	scopes, err := queryField.Scopes()
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = b.GetDB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *BaseRest) GetFields(ctx *gin.Context) (interface{}, error) {
	var queryFields QueryFields
	err := ctx.ShouldBind(&queryFields)
	if err != nil {
		return nil, err
	}

	scopes, err := queryFields.Scopes()
	if err != nil {
		return nil, err
	}

	objs := b.GetFindObj()
	err = b.GetDB().Model(b.GetModelObj()).Scopes(scopes...).Find(objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

type RelationCreate struct {
	Id     int64    `json:"id" uri:"id"`
	Fields []string `json:"fields" form:"_fields"`
}

// 一对多查询
func (b *BaseRest) Relations(ctx *gin.Context) (interface{}, error) {

	var getRelation RelationGet
	err := ctx.ShouldBindUri(&getRelation)
	if err != nil {
		return nil, err
	}
	err = ctx.ShouldBind(&getRelation)
	if err != nil {
		return nil, err
	}

	scopes, err := getRelation.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetModelObj()

	err = b.GetDB().Find(obj, getRelation.Id).Error
	if err != nil {
		return nil, err
	}

	rc, ok := b.RelationsField()[getRelation.RelationName]
	if !ok {
		return nil, fmt.Errorf("%v relation field permission denied", getRelation.RelationName)
	}

	tableName := rc.GetTableName()
	result := rc.GetFindObj()

	count := b.GetDB().Model(obj).Scopes(scopes...).Association(tableName).Count()

	err = b.GetDB().Model(obj).Scopes(scopes...).Association(tableName).Find(result)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"list": result, "count": count}, nil
}

func (b *BaseRest) RelationCreate(ctx *gin.Context) (interface{}, error) {
	var relationCreate RelationCreate
	err := ctx.ShouldBindUri(&relationCreate)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindQuery(&relationCreate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = b.GetDB().First(obj, relationCreate.Id).Error
	if err != nil {
		return nil, err
	}

	//s,o := b.UpdateField()

	db := b.GetDB()

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

func (b *BaseRest) RelationUpdate(ctx *gin.Context) (interface{}, error) {
	var relationUpdate RelationCreate
	err := ctx.ShouldBindUri(&relationUpdate)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindQuery(&relationUpdate)
	if err != nil {
		return nil, err
	}
	obj := b.GetFirstObj()
	err = ctx.ShouldBindJSON(obj)
	if err != nil {
		return nil, err
	}

	err = b.GetDB().First(obj, relationUpdate.Id).Error
	if err != nil {
		return nil, err
	}

	s, o := b.UpdateField()

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

	err = b.GetDB().Session(&gorm.Session{FullSaveAssociations: true}).Select(s).Omit(o...).Updates(obj).Error
	if err != nil {
		return nil, err
	}

	return obj, nil
}
