// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/fitan/magic/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newAudit(db *gorm.DB) audit {
	_audit := audit{}

	_audit.auditDo.UseDB(db)
	_audit.auditDo.UseModel(&model.Audit{})

	tableName := _audit.auditDo.TableName()
	_audit.ALL = field.NewField(tableName, "*")
	_audit.ID = field.NewUint(tableName, "id")
	_audit.CreatedAt = field.NewTime(tableName, "created_at")
	_audit.UpdatedAt = field.NewTime(tableName, "updated_at")
	_audit.DeletedAt = field.NewField(tableName, "deleted_at")
	_audit.Url = field.NewString(tableName, "url")
	_audit.Query = field.NewString(tableName, "query")
	_audit.Method = field.NewString(tableName, "method")
	_audit.Request = field.NewString(tableName, "request")
	_audit.Response = field.NewString(tableName, "response")
	_audit.Header = field.NewString(tableName, "header")
	_audit.StatusCode = field.NewInt(tableName, "status_code")
	_audit.RemoteIP = field.NewString(tableName, "remote_ip")
	_audit.ClientIP = field.NewString(tableName, "client_ip")
	_audit.CostTime = field.NewString(tableName, "cost_time")

	_audit.fillFieldMap()

	return _audit
}

type audit struct {
	auditDo

	ALL        field.Field
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Url        field.String
	Query      field.String
	Method     field.String
	Request    field.String
	Response   field.String
	Header     field.String
	StatusCode field.Int
	RemoteIP   field.String
	ClientIP   field.String
	CostTime   field.String

	fieldMap map[string]field.Expr
}

func (a audit) As(alias string) *audit {
	a.auditDo.DO = *(a.auditDo.As(alias).(*gen.DO))

	a.ALL = field.NewField(alias, "*")
	a.ID = field.NewUint(alias, "id")
	a.CreatedAt = field.NewTime(alias, "created_at")
	a.UpdatedAt = field.NewTime(alias, "updated_at")
	a.DeletedAt = field.NewField(alias, "deleted_at")
	a.Url = field.NewString(alias, "url")
	a.Query = field.NewString(alias, "query")
	a.Method = field.NewString(alias, "method")
	a.Request = field.NewString(alias, "request")
	a.Response = field.NewString(alias, "response")
	a.Header = field.NewString(alias, "header")
	a.StatusCode = field.NewInt(alias, "status_code")
	a.RemoteIP = field.NewString(alias, "remote_ip")
	a.ClientIP = field.NewString(alias, "client_ip")
	a.CostTime = field.NewString(alias, "cost_time")

	a.fillFieldMap()

	return &a
}

func (a *audit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *audit) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["url"] = a.Url
	a.fieldMap["query"] = a.Query
	a.fieldMap["method"] = a.Method
	a.fieldMap["request"] = a.Request
	a.fieldMap["response"] = a.Response
	a.fieldMap["header"] = a.Header
	a.fieldMap["status_code"] = a.StatusCode
	a.fieldMap["remote_ip"] = a.RemoteIP
	a.fieldMap["client_ip"] = a.ClientIP
	a.fieldMap["cost_time"] = a.CostTime
}

func (a audit) clone(db *gorm.DB) audit {
	a.auditDo.ReplaceDB(db)
	return a
}

type auditDo struct{ gen.DO }

func (a auditDo) Debug() *auditDo {
	return a.withDO(a.DO.Debug())
}

func (a auditDo) WithContext(ctx context.Context) *auditDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a auditDo) Clauses(conds ...clause.Expression) *auditDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a auditDo) Not(conds ...gen.Condition) *auditDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a auditDo) Or(conds ...gen.Condition) *auditDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a auditDo) Select(conds ...field.Expr) *auditDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a auditDo) Where(conds ...gen.Condition) *auditDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a auditDo) Order(conds ...field.Expr) *auditDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a auditDo) Distinct(cols ...field.Expr) *auditDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a auditDo) Omit(cols ...field.Expr) *auditDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a auditDo) Join(table schema.Tabler, on ...field.Expr) *auditDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a auditDo) LeftJoin(table schema.Tabler, on ...field.Expr) *auditDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a auditDo) RightJoin(table schema.Tabler, on ...field.Expr) *auditDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a auditDo) Group(cols ...field.Expr) *auditDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a auditDo) Having(conds ...gen.Condition) *auditDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a auditDo) Limit(limit int) *auditDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a auditDo) Offset(offset int) *auditDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a auditDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *auditDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a auditDo) Unscoped() *auditDo {
	return a.withDO(a.DO.Unscoped())
}

func (a auditDo) Create(values ...*model.Audit) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a auditDo) CreateInBatches(values []*model.Audit, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a auditDo) Save(values ...*model.Audit) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a auditDo) First() (*model.Audit, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Audit), nil
	}
}

func (a auditDo) Take() (*model.Audit, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Audit), nil
	}
}

func (a auditDo) Last() (*model.Audit, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Audit), nil
	}
}

func (a auditDo) Find() ([]*model.Audit, error) {
	result, err := a.DO.Find()
	return result.([]*model.Audit), err
}

func (a auditDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Audit, err error) {
	buf := make([]*model.Audit, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a auditDo) FindInBatches(result *[]*model.Audit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a auditDo) Attrs(attrs ...field.AssignExpr) *auditDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a auditDo) Assign(attrs ...field.AssignExpr) *auditDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a auditDo) Joins(field field.RelationField) *auditDo {
	return a.withDO(a.DO.Joins(field))
}

func (a auditDo) Preload(field field.RelationField) *auditDo {
	return a.withDO(a.DO.Preload(field))
}

func (a auditDo) FirstOrInit() (*model.Audit, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Audit), nil
	}
}

func (a auditDo) FirstOrCreate() (*model.Audit, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Audit), nil
	}
}

func (a auditDo) FindByPage(offset int, limit int) (result []*model.Audit, count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = a.Offset(offset).Limit(limit).Find()
	return
}

func (a auditDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a *auditDo) withDO(do gen.Dao) *auditDo {
	a.DO = *do.(*gen.DO)
	return a
}
