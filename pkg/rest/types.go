package rest

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	Create(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	UpdateMany(ctx *gin.Context) (interface{}, error)
	Delete(ctx *gin.Context) (interface{}, error)
	DeleteMany(ctx *gin.Context) (interface{}, error)

	GetField(ctx *gin.Context) (interface{}, error)
	GetFields(ctx *gin.Context) (interface{}, error)

	Relations(ctx *gin.Context) (interface{}, error)
}

type Objer interface {
	GetObj() interface{}
	GetObjs() interface{}
}

type BaseRest struct {
	DB *gorm.DB
	Objer
}

type GetOneById struct {
	Id int64 `json:"id" uri:"id" form:"id"`
}

func (g *GetOneById) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", g.Id)
	})
	return scopes, nil
}

type GetManyReference struct {
	GetList
	GetOneById
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

type filter map[string]interface{}

type GetManyByIds struct {
	Ids []int64 `json:"ids" form:"ids"`
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
	b.DB = db
}

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

	var result []map[string]interface{}
	err = b.DB.Model(b.GetObj()).Count(&count).Scopes(scopes...).Find(&result).Error
	return map[string]interface{}{"list": result, "count": count}, errors.WithMessage(err, "DB find")

}

func (b *BaseRest) GetOne(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = b.DB.Model(b.GetObj()).Scopes(scopes...).First(&result).Error
	return result, err
}

func (b *BaseRest) Create(ctx *gin.Context) (interface{}, error) {
	obj := b.GetObj()
	err := ctx.BindJSON(obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = b.DB.Create(obj).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return obj, nil
}

func (b *BaseRest) Update(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var m map[string]interface{}
	err = ctx.BindJSON(&m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetObj()
	err = b.DB.Model(obj).Clauses(clause.Returning{}).Scopes(scopes...).Updates(m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return obj, nil
}

func (b *BaseRest) UpdateMany(ctx *gin.Context) (interface{}, error) {
	var getManyByIds GetManyByIds
	err := ctx.BindQuery(&getManyByIds)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := getManyByIds.Scopes()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var m []map[string]interface{}
	err = ctx.BindJSON(&m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	objs := b.GetObjs()
	err = b.DB.Model(objs).Clauses(clause.Returning{}).Scopes(scopes...).Updates(m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return objs, nil
}

func (b *BaseRest) Delete(ctx *gin.Context) (interface{}, error) {
	var getOneById GetOneById
	err := ctx.BindUri(&getOneById)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := getOneById.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetObj()
	err = b.DB.Clauses(clause.Returning{}).Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return obj, nil
}

func (b *BaseRest) DeleteMany(ctx *gin.Context) (interface{}, error) {
	var getManyByIds GetManyByIds
	err := ctx.BindQuery(&getManyByIds)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scopes, err := getManyByIds.Scopes()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := b.GetObj()
	err = b.DB.Where("1 = 1").Scopes(scopes...).Delete(obj).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return getManyByIds.Ids, nil
}

type QueryField struct {
	KeyWord *string `json:"query" form:"_keyWord"`
	Name    string  `uri:"name"`
}

func (q *QueryField) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Select("id", q.Name)
		if q.KeyWord != nil {
			db = db.Where(fmt.Sprintf("%s LIKE ?", q.Name), "%"+*q.KeyWord+"%")
		}
		return db
	})
	return scopes, nil

}

func (b *BaseRest) GetField(ctx *gin.Context) (interface{}, error) {
	var queryField QueryField
	err := ctx.ShouldBind(&queryField)
	if err != nil {
		return nil, err
	}
	spew.Dump(queryField)

	scopes, err := queryField.Scopes()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = b.DB.Model(b.GetObj()).Scopes(scopes...).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

type QueryFields struct {
	KeyWord *string  `json:"query" form:"_keyWord"`
	Fields  []string `json:"name" form:"_fields"`
}

func (q *QueryFields) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Select("id").Select(q.Fields)
		if q.KeyWord != nil {
			for _, v := range q.Fields {
				db = db.Where(fmt.Sprintf("%s LIKE ?", v), "%"+*q.KeyWord+"%")
			}
		}
		return db
	})
	return scopes, nil

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

	var result []map[string]interface{}
	err = b.DB.Scopes(scopes...).Find(&result).Error
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
	err := ctx.ShouldBind(&getRelation)
	if err != nil {
		return nil, err
	}

	scopes, err := getRelation.Scopes()
	if err != nil {
		return nil, err
	}

	obj := b.GetObj()

	err = b.DB.Find(obj, getRelation.Id).Error
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	var count int64
	err = b.DB.Model(obj).Count(&count).Scopes(scopes...).Association(getRelation.RelationName).Find(&result)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"list": result, "count": count}, nil
}
