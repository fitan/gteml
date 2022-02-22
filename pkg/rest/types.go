package rest

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type GetOneById struct {
	Id int64 `uri:"id"`
}

func (g *GetOneById) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", g.Id)
	})
	return scopes, nil
}

type GetList struct {
	Page     *int     `json:"page" form:"_page"`
	Limit    *int     `json:"limit" form:"_limit"`
	Sort     *string  `json:"sort" form:"_sort"`
	Order    *string  `json:"order" form:"_order"`
	Filter   *string  `json:"filter" form:"_filter"`
	Includes []string `json:"includes" form:"_includes"`
	Ids      []int    `json:"ids" form:"_ids"`
}

type filter map[string]interface{}

func (f *filter) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		save := db
		for k, v := range *f {
			save = save.Where(fmt.Sprintf("%s = ?", k), v)
		}
		return save
	})
	return
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

		var filterScopes []func(db *gorm.DB) *gorm.DB
		filterScopes, err = filter.Scopes()
		if err != nil {
			return
		}

		scopes = append(scopes, filterScopes...)
	}
	return
}

type GetManyByIds struct {
	Ids []int64 `json:"ids" form:"_ids"`
}

func (g *GetManyByIds) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	if g.Ids != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id IN ?", g.Ids)
		})
	}
	return
}

type QueryField struct {
	KeyWord *string `form:"_keyWord"`
	Name    string  `uri:"name"`
}

func (q *QueryField) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		db = db.Distinct(q.Name).Select("id", q.Name)
		if q.KeyWord != nil {
			db = db.Where(fmt.Sprintf("%s LIKE ?", q.Name), "%"+*q.KeyWord+"%")
		}
		return db
	})
	return scopes, nil
}

type QueryFields struct {
	KeyWord *string  `json:"query" form:"_keyWord"`
	Fields  []string `json:"name" form:"_fields"`
}

func (q *QueryFields) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
		var cy []interface{}
		cy = append(cy, q.Fields)
		db = db.Distinct(cy...).Select(append([]string{"id"}, q.Fields...))
		if q.KeyWord != nil {
			for _, v := range q.Fields {
				db = db.Where(fmt.Sprintf("%s LIKE ?", v), "%"+*q.KeyWord+"%")
			}
		}
		return db
	})
	return scopes, nil
}

type RelationGet struct {
	Id           int64   `json:"id" uri:"id"`
	RelationName string  `json:"relationName" uri:"relationName"`
	Page         *int    `json:"page" form:"_page"`
	Limit        *int    `json:"limit" form:"_limit"`
	Sort         *string `json:"sort" form:"_sort"`
	Order        *string `json:"order" form:"_order"`
	Filter       *string `json:"filter" form:"_filter"`
}

func (g *RelationGet) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
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

		var filterScopes []func(db *gorm.DB) *gorm.DB
		filterScopes, err = filter.Scopes()
		if err != nil {
			return
		}

		scopes = append(scopes, filterScopes...)
	}
	return
}
