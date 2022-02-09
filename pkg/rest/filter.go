package rest

import (
	"fmt"
	"github.com/fitan/magic/dao/dal/query"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strings"
)

type GetConstraint struct {
	Filters []Filter
	Sorts   Sort
}

type Filters []Filter

type Filter struct {
	Field string `json:"field" validate:""`
	// =
	// !=
	// <
	// >
	// >=
	// <=
	// BETWEEN NOT
	// BETWEEN
	// IN
	// NOT IN
	Expr string   `json:"expr"`
	Val  []string `json:"val"`
}

type ValidateField struct {
	Expr string `json:"expr"`
	Val  string `json:"val"`
}

func (f Filters) ValidateFields() map[string]ValidateField {
	return map[string]ValidateField{"name": ValidateField{
		Expr: "required",
		Val:  "",
	}}
}

func (f Filters) GenScopes(db *gorm.DB) (*gorm.DB, error) {
	save := db
	for _, v := range f {
		if v.Expr == "BETWEEN" {
			save = db.Where(fmt.Sprintf("%s %s ? AND ?", v.Field, v.Expr), v.Val[0], v.Val[1])
			continue
		}

		if v.Expr == "IN" {
			save = db.Where(fmt.Sprintf("%s %s ?", v.Field, v.Expr))
		}
		save = db.Where(fmt.Sprintf("%s %s ?", v.Field, v.Expr), v.Val[0])
	}
	return save, nil
}

type GenObjInterface interface {
	GetFieldByName(fieldName string) (field.OrderExpr, bool)
}

// 排序
type Sort struct {
	// -name,+age
	Sorts []string `json:"sorts"`
}

func (s Sort) GetFieldByNameFn(query *query.Query) GenObjInterface {
	return &query.User
}

func (s Sort) GormScopes(i GenObjInterface, db *gorm.DB) (*gorm.DB, error) {
	save := db
	for _, s := range s.Sorts {
		trimName := strings.TrimPrefix(s, "-")
		if strings.HasPrefix(s, "-") {
			_, has := i.GetFieldByName(trimName)
			if !has {
				return nil, fmt.Errorf("sort not find field %v", trimName)
			}
			save = save.Order(trimName + " desc")
			continue
		}

		if strings.HasPrefix(s, "+") {
			_, has := i.GetFieldByName(trimName)
			if !has {
				return nil, fmt.Errorf("sort not find field %v", trimName)
			}
			save = save.Order(trimName)
			continue
		}
		return nil, fmt.Errorf("sort field: -%v or +%v", s, s)
	}
	return save, nil
}

func (s Sort) GormGen(i GenObjInterface, tx gen.Dao) (gen.Dao, error) {
	save := tx
	for _, s := range s.Sorts {
		trimName := strings.TrimPrefix(s, "-")
		if strings.HasPrefix(s, "-") {
			expr, has := i.GetFieldByName(trimName)
			if !has {
				return nil, fmt.Errorf("sort not find field %v", trimName)
			}
			save.Order(expr.Desc())
			continue
		}

		if strings.HasPrefix(s, "+") {
			expr, has := i.GetFieldByName(trimName)
			if !has {
				return nil, fmt.Errorf("sort not find field %v", trimName)
			}
			save.Order(expr)
			continue
		}
		return nil, fmt.Errorf("sort field: -%v or +%v", s, s)
	}
	return save, nil
}

// 全文查询
type Q struct {
	Q string `json:"q"`
}

// 分页
type Pagination struct {
	Page  int
	Limit int
}

func (p Pagination) GormScope(db *gorm.DB) (*gorm.DB, error) {
	return db.Limit(p.Limit).Offset((p.Page - 1) * p.Limit), nil
}
