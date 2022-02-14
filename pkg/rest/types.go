package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type CURD struct {
}

type Field struct {
	Name string
	Type string

	// Create
	EnableCreate bool
	// 自己输入
	Input bool
	// 远程获取
	RemoteAuto bool
	RemoteFn   func(ctx *gin.Context)

	// Table
	EnableTable bool
	// 别名
	Lable string
	// 可修改
	EnableModify bool

	// query
	EnableQuery      bool
	EnableEq         bool
	EnableGt         bool
	EnableLt         bool
	EnableLte        bool
	EnableIn         bool
	EnableNotIn      bool
	EnableBetween    bool
	EnableNotBetween bool
	EnableLike       bool
	EnableNotLike    bool
	EnableSort       bool
}

type Restful interface {
	SetDB(db *gorm.DB)
	GetObj() interface{}
	GetObjs() interface{}
	Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error))
	GetList(ctx *gin.Context) (interface{}, error)
	GetOne(ctx *gin.Context) (interface{}, error)
	//GetMany(ctx *gin.Context) (interface{},error)
	//GetManyReference(ctx *gin.Context) (interface{},error)
	Create(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	UpdateMany(ctx *gin.Context) (interface{}, error)
	Delete(ctx *gin.Context) (interface{}, error)
	DeleteMany(ctx *gin.Context) (interface{}, error)

	GetField(ctx *gin.Context) (interface{}, error)
	GetFields(ctx *gin.Context) (interface{}, error)

	// 关联查询
	Relations(ctx *gin.Context) (interface{}, error)
}

type Objer interface {
	GetObj() interface{}
	GetObjs() interface{}
}

type BaseRest struct {
	db *gorm.DB
	Objer
}

type GetOne struct {
	Id int64 `json:"id" uri:"id" form:"id"`
}

func (g *GetOne) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", g.Id)
	})
	return scopes, nil
}

type GetManyReference struct {
	GetList
	GetOne
	Target string `json:"target" form:"target"`
}

func (g *GetManyReference) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	return g.GetList.Scopes()
}

type GetList struct {
	Page   *int    `json:"page" form:"_page"`
	Limit  *int    `json:"limit" form:"_limit"`
	Sort   *string `json:"sort" form:"_sort"`
	Order  *string `json:"order" form:"_order"`
	Filter *string `json:"filter" form:"_filter"`
	// GetMany must
	Ids []int `json:"ids" form:"ids"`
	// GetManyReference must
	Target *string `json:"target" form:"target"`
	Id     *int64  `json:"id" form:"id"`
}

func (g *GetList) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	if g.Sort != nil && g.Order != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%v %v", *g.Sort, *g.Order))
		})
	}

	if g.Ids != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id IN ?", g.Ids)
		})
		// 如果查询ids 直接返回
		return
	}

	if g.Page != nil && g.Limit != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Limit(*g.Limit).Offset((*g.Page - 1) * *g.Limit)
		})
	}

	if g.Filter != nil {
		var filter filter
		err = json.Unmarshal([]byte(*g.Filter), &filter)
		if err != nil {
			return
		}

		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			save := db
			for k, v := range filter {
				save = save.Where(fmt.Sprintf("%s = ?", k), v)
			}
			return save
		})
	}
	return
}

type filter map[string]string

type GetManyByIds struct {
	Ids []int `json:"filter" form:"ids"`
}

type Ids struct {
	Id []int64 `json:"id"`
}

func (g *GetManyByIds) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	if g.Ids != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id IN ?", g.Ids)
		})
	}
	return
}

func (b *BaseRest) SetDB(db *gorm.DB) {
	b.db = db
}

//func (b *BaseRest) GetObj() interface{} {
//	return nil
//}
//
//func (b *BaseRest) GetObjs() interface{} {
//	return nil
//}

func (b *BaseRest) Wrap(ctx *gin.Context, fn func(ctx *gin.Context) (interface{}, error)) {
	data, err := fn(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
	return

	//code := 2000
	//msg := ""
	//data,err := fn(ctx)
	//if err != nil {
	//	code = 5003
	//	msg = err.Error()
	//}
	//ctx.JSON(http.StatusOK, map[string]interface{}{
	//	"code": code,
	//	"data": data,
	//	"msg": msg,
	//})
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
	defer func() {
		ctx.Header("X-Total-Count", cast.ToString(count))
	}()
	// GetManyReference
	if getList.Target != nil && getList.Id != nil {
		obj := b.GetObj()
		err = b.db.First(obj, *getList.Id).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var result []map[string]interface{}
		err = b.db.Model(obj).Scopes(scopes...).Count(&count).Association(*getList.Target).Find(&result)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return result, nil

	}

	var result []map[string]interface{}
	err = b.db.Model(b.GetObj()).Count(&count).Scopes(scopes...).Find(&result).Error
	fmt.Printf("count %v", count)
	return result, errors.WithMessage(err, "db find")

}

func (b *BaseRest) GetOne(ctx *gin.Context) (interface{}, error) {
	var id GetOne
	err := ctx.BindUri(&id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result map[string]interface{}
	err = b.db.Model(b.GetObj()).First(&result, "id = ?", id.Id).Error
	return result, err
}

func (b *BaseRest) Create(ctx *gin.Context) (interface{}, error) {
	obj := b.GetObj()
	err := ctx.BindJSON(obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = b.db.Create(obj).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return obj, nil
}

func (b *BaseRest) Update(ctx *gin.Context) (interface{}, error) {
	var update GetOne
	err := ctx.BindUri(&update)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var m map[string]interface{}
	err = ctx.BindJSON(&m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := b.GetObj()
	err = b.db.Model(obj).Clauses(clause.Returning{}).Where("id = ?", update.Id).Updates(m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return obj, nil
}

func (b *BaseRest) UpdateMany(ctx *gin.Context) (interface{}, error) {
	var update GetManyByIds
	err := ctx.BindQuery(&update)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := update.Scopes()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var m []map[string]interface{}
	err = ctx.BindJSON(&m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	objs := b.GetObjs()
	err = b.db.Model(objs).Clauses(clause.Returning{}).Scopes(scopes...).Updates(m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return objs, nil
}

func (b *BaseRest) Delete(ctx *gin.Context) (interface{}, error) {
	var delete GetOne
	err := ctx.BindUri(&delete)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := b.GetObj()
	err = b.db.Clauses(clause.Returning{}).Delete(obj, delete.Id).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return obj, nil
}

func (b *BaseRest) DeleteMany(ctx *gin.Context) (interface{}, error) {
	var delete GetManyByIds
	err := ctx.BindQuery(&delete)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := delete.Scopes()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := b.GetObj()
	err = b.db.Where("1 = 1").Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return delete.Ids, nil
}

type QueryField struct {
	Key       *string `json:"query" form:"_key"`
	FieldName string  `json:"name" uri:"fieldName"`
}

func (q *QueryField) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Select("id", q.FieldName)
		if q.Key != nil {
			db = db.Where(fmt.Sprintf("%s LIKE ?", q.FieldName), "%"+*q.Key+"%")
		}
		return db
	})
	return scopes, nil

}

func (b *BaseRest) GetField(ctx *gin.Context) (interface{}, error) {
	var queryField QueryField
	err := ctx.Bind(&queryField)
	if err != nil {
		return nil, err
	}

	scopes, err := queryField.Scopes()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = b.db.Model(b.GetObj()).Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

type QueryFields struct {
	Key    *string  `json:"query" form:"_key"`
	Fields []string `json:"name" form:"_fields"`
}

func (q *QueryFields) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Select("id").Select(q.Fields)
		if q.Key != nil {
			for _, v := range q.Fields {
				db = db.Where(fmt.Sprintf("%s LIKE ?", v), "%"+*q.Key+"%")
			}
		}
		return db
	})
	return scopes, nil

}

func (b *BaseRest) GetFields(ctx *gin.Context) (interface{}, error) {
	var queryFields QueryFields
	err := ctx.Bind(&queryFields)
	if err != nil {
		return nil, err
	}

	scopes, err := queryFields.Scopes()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = b.db.Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

type GetRelation struct {
	Id           int64   `json:"id" uri:"id"`
	RelationName string  `json:"relationName" uri:"relationName"`
	Page         *int    `json:"page" form:"_page"`
	Limit        *int    `json:"limit" form:"_limit"`
	Sort         *string `json:"sort" form:"_sort"`
	Order        *string `json:"order" form:"_order"`
	Filter       *string `json:"filter" form:"_filter"`
}

func (g *GetRelation) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	if g.Sort != nil && g.Order != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%v %v", *g.Sort, *g.Order))
		})
	}

	if g.Page != nil && g.Limit != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Limit(*g.Limit).Offset((*g.Page - 1) * *g.Limit)
		})
	}

	if g.Filter != nil {
		var filter filter
		err = json.Unmarshal([]byte(*g.Filter), &filter)
		if err != nil {
			return
		}

		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			save := db
			for k, v := range filter {
				save = save.Where(fmt.Sprintf("%s = ?", k), v)
			}
			return save
		})
	}
	return
}

// 一对多查询
func (b *BaseRest) Relations(ctx *gin.Context) (interface{}, error) {

	var getRelation GetRelation
	err := ctx.Bind(&getRelation)
	if err != nil {
		return nil, err
	}

	scopes, err := getRelation.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetObj()

	err = b.db.Find(obj, getRelation.Id).Error
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = b.db.Model(obj).Scopes(scopes...).Association(getRelation.RelationName).Find(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
