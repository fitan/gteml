// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/fitan/magic/dao/dal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newService(db *gorm.DB) service {
	_service := service{}

	_service.serviceDo.UseDB(db)
	_service.serviceDo.UseModel(&model.Service{})

	tableName := _service.serviceDo.TableName()
	_service.ALL = field.NewField(tableName, "*")
	_service.ID = field.NewUint(tableName, "id")
	_service.CreatedAt = field.NewTime(tableName, "created_at")
	_service.UpdatedAt = field.NewTime(tableName, "updated_at")
	_service.DeletedAt = field.NewField(tableName, "deleted_at")
	_service.Name = field.NewString(tableName, "name")
	_service.Alias = field.NewString(tableName, "alias")
	_service.Description = field.NewString(tableName, "description")
	_service.ParentId = field.NewUint(tableName, "parent_id")
	_service.Services = serviceServices{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Services", "model.Service"),
		Services: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Services.Services", "model.Service"),
		},
	}

	_service.fillFieldMap()

	return _service
}

type service struct {
	serviceDo

	ALL         field.Field
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Name        field.String
	Alias       field.String
	Description field.String
	ParentId    field.Uint
	Services    serviceServices

	fieldMap map[string]field.Expr
}

func (s service) As(alias string) *service {
	s.serviceDo.DO = *(s.serviceDo.As(alias).(*gen.DO))

	s.ALL = field.NewField(alias, "*")
	s.ID = field.NewUint(alias, "id")
	s.CreatedAt = field.NewTime(alias, "created_at")
	s.UpdatedAt = field.NewTime(alias, "updated_at")
	s.DeletedAt = field.NewField(alias, "deleted_at")
	s.Name = field.NewString(alias, "name")
	s.Alias = field.NewString(alias, "alias")
	s.Description = field.NewString(alias, "description")
	s.ParentId = field.NewUint(alias, "parent_id")

	s.fillFieldMap()

	return &s
}

func (s *service) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *service) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["alias"] = s.Alias
	s.fieldMap["description"] = s.Description
	s.fieldMap["parent_id"] = s.ParentId

}

func (s service) clone(db *gorm.DB) service {
	s.serviceDo.ReplaceDB(db)
	return s
}

type serviceServices struct {
	db *gorm.DB

	field.RelationField

	Services struct {
		field.RelationField
	}
}

func (a serviceServices) Where(conds ...field.Expr) *serviceServices {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a serviceServices) WithContext(ctx context.Context) *serviceServices {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a serviceServices) Model(m *model.Service) *serviceServicesTx {
	return &serviceServicesTx{a.db.Model(m).Association(a.Name())}
}

type serviceServicesTx struct{ tx *gorm.Association }

func (a serviceServicesTx) Find() (result *model.Service, err error) {
	return result, a.tx.Find(&result)
}

func (a serviceServicesTx) Append(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a serviceServicesTx) Replace(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a serviceServicesTx) Delete(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a serviceServicesTx) Clear() error {
	return a.tx.Clear()
}

func (a serviceServicesTx) Count() int64 {
	return a.tx.Count()
}

type serviceDo struct{ gen.DO }

func (s serviceDo) Debug() *serviceDo {
	return s.withDO(s.DO.Debug())
}

func (s serviceDo) WithContext(ctx context.Context) *serviceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serviceDo) Clauses(conds ...clause.Expression) *serviceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serviceDo) Not(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serviceDo) Or(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serviceDo) Select(conds ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serviceDo) Where(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serviceDo) Order(conds ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serviceDo) Distinct(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serviceDo) Omit(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serviceDo) Join(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *serviceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serviceDo) Group(cols ...field.Expr) *serviceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serviceDo) Having(conds ...gen.Condition) *serviceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serviceDo) Limit(limit int) *serviceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serviceDo) Offset(offset int) *serviceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *serviceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serviceDo) Unscoped() *serviceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serviceDo) Create(values ...*model.Service) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serviceDo) CreateInBatches(values []*model.Service, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serviceDo) Save(values ...*model.Service) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serviceDo) First() (*model.Service, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Take() (*model.Service, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Last() (*model.Service, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) Find() ([]*model.Service, error) {
	result, err := s.DO.Find()
	return result.([]*model.Service), err
}

func (s serviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Service, err error) {
	buf := make([]*model.Service, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serviceDo) FindInBatches(result *[]*model.Service, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serviceDo) Attrs(attrs ...field.AssignExpr) *serviceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serviceDo) Assign(attrs ...field.AssignExpr) *serviceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serviceDo) Joins(field field.RelationField) *serviceDo {
	return s.withDO(s.DO.Joins(field))
}

func (s serviceDo) Preload(field field.RelationField) *serviceDo {
	return s.withDO(s.DO.Preload(field))
}

func (s serviceDo) FirstOrInit() (*model.Service, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) FirstOrCreate() (*model.Service, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Service), nil
	}
}

func (s serviceDo) FindByPage(offset int, limit int) (result []*model.Service, count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = s.Offset(offset).Limit(limit).Find()
	return
}

func (s serviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *serviceDo) withDO(do gen.Dao) *serviceDo {
	s.DO = *do.(*gen.DO)
	return s
}
