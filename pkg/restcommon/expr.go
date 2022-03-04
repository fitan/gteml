package restcommon

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

const (
	Eq         = "EQ"
	Neq        = "NEQ"
	Gt         = "GT"
	Gte        = "GTE"
	Lt         = "LT"
	Lte        = "LTE"
	Between    = "BETWEEN"
	NotBetween = "NOTBETWEEN"
	In         = "IN"
	NotIn      = "NOTIN"
	Like       = "LIKE"
	NotLike    = "NOTLIKE"
	Regexp     = "REGEXP"
	NotRegexp  = "NOTREGEXP"
)

type ExprT string

func (e ExprT) Build(Field string, Value interface{}) (res func(db *gorm.DB) *gorm.DB, err error) {
	switch strings.ToUpper(string(e)) {
	case Eq:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s = ?", Field), Value)
		}, nil
	case Neq:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s != ?", Field), Value)
		}, nil
	case Gt:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s > ?", Field), Value)
		}, nil
	case Gte:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s >= ?", Field), Value)
		}, nil
	case Lt:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s < ?", Field), Value)
		}, nil
	case Lte:
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s <= ?", Field), Value)
		}, nil
	case Between:
		valueS, ok := Value.(string)
		if !ok {
			return nil, fmt.Errorf("%v is not string", Value)
		}
		valueSS := strings.Split(valueS, "&")
		if len(valueSS) != 2 {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", Field), valueSS[0], valueSS[1])
		}, nil
	case NotBetween:
		valueS, ok := Value.(string)
		if !ok {
			return nil, fmt.Errorf("%v is not string", Value)
		}
		valueSS := strings.Split(valueS, "&")
		if len(valueSS) != 2 {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Not(fmt.Sprintf("%s BETWEEN ? AND ?", Field), valueSS[0], valueSS[1])
		}, nil
	case In:
		valueS, ok := Value.([]interface{})
		if !ok {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s IN ?", Field), valueS)
		}, nil
	case NotIn:
		valueS, ok := Value.([]interface{})
		if !ok {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Not(map[string]interface{}{Field: valueS})
		}, nil
	case Like:
		valueS, ok := Value.(string)
		if !ok {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s LIKE ?", Field), "%"+valueS+"%")
		}, nil
	case NotLike:
		valueS, ok := Value.(string)
		if !ok {
			return nil, fmt.Errorf("%v format is wrong", Value)
		}
		return func(db *gorm.DB) *gorm.DB {
			return db.Not(fmt.Sprintf("%s LIKE ?", Field), "%"+valueS+"%")
		}, nil
		//case Regexp:
		//
		//case NotRegexp:

	}
	return nil, errors.New("unknown expression")
}

type Expr struct {
	Field string      `json:"field"`
	Expr  ExprT       `json:"expr"`
	Value interface{} `json:"value"`
	Or    *Expr       `json:"or"`
}

type Exprs []Expr

func (e *Exprs) Scopes() (scopes []func(db *gorm.DB) *gorm.DB, err error) {
	for _, v := range *e {
		s, err := v.Expr.Build(v.Field, v.Value)
		if err != nil {
			return nil, err
		}
	}
}

func depthScope(e *Expr, scopes *[]func(db *gorm.DB) *gorm.DB) error {
	if e != nil {
		s, err := e.Expr.Build(e.Field, e.Value)
		if err != nil {
			return err
		}
		*scopes = append(*scopes, s)
		depthScope(e.or)
	}
}
