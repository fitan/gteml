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

func newRole(db *gorm.DB) role {
	_role := role{}

	_role.roleDo.UseDB(db)
	_role.roleDo.UseModel(&model.Role{})

	tableName := _role.roleDo.TableName()
	_role.ALL = field.NewField(tableName, "*")
	_role.ID = field.NewUint(tableName, "id")
	_role.CreatedAt = field.NewTime(tableName, "created_at")
	_role.UpdatedAt = field.NewTime(tableName, "updated_at")
	_role.DeletedAt = field.NewField(tableName, "deleted_at")
	_role.Name = field.NewString(tableName, "name")
	_role.OnlyKey = field.NewString(tableName, "only_key")
	_role.Enabled = field.NewBool(tableName, "enabled")
	_role.Description = field.NewString(tableName, "description")
	_role.Level = field.NewInt(tableName, "level")
	_role.Permissions = rolePermissions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Permissions", "model.Permission"),
	}

	_role.fillFieldMap()

	return _role
}

type role struct {
	roleDo

	ALL         field.Field
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Name        field.String
	OnlyKey     field.String
	Enabled     field.Bool
	Description field.String
	Level       field.Int
	Permissions rolePermissions

	fieldMap map[string]field.Expr
}

func (r role) As(alias string) *role {
	r.roleDo.DO = *(r.roleDo.As(alias).(*gen.DO))

	r.ALL = field.NewField(alias, "*")
	r.ID = field.NewUint(alias, "id")
	r.CreatedAt = field.NewTime(alias, "created_at")
	r.UpdatedAt = field.NewTime(alias, "updated_at")
	r.DeletedAt = field.NewField(alias, "deleted_at")
	r.Name = field.NewString(alias, "name")
	r.OnlyKey = field.NewString(alias, "only_key")
	r.Enabled = field.NewBool(alias, "enabled")
	r.Description = field.NewString(alias, "description")
	r.Level = field.NewInt(alias, "level")

	r.fillFieldMap()

	return &r
}

func (r *role) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *role) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 10)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["name"] = r.Name
	r.fieldMap["only_key"] = r.OnlyKey
	r.fieldMap["enabled"] = r.Enabled
	r.fieldMap["description"] = r.Description
	r.fieldMap["level"] = r.Level

}

func (r role) clone(db *gorm.DB) role {
	r.roleDo.ReplaceDB(db)
	return r
}

type rolePermissions struct {
	db *gorm.DB

	field.RelationField
}

func (a rolePermissions) Where(conds ...field.Expr) *rolePermissions {
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

func (a rolePermissions) WithContext(ctx context.Context) *rolePermissions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a rolePermissions) Model(m *model.Role) *rolePermissionsTx {
	return &rolePermissionsTx{a.db.Model(m).Association(a.Name())}
}

type rolePermissionsTx struct{ tx *gorm.Association }

func (a rolePermissionsTx) Find() (result *model.Permission, err error) {
	return result, a.tx.Find(&result)
}

func (a rolePermissionsTx) Append(values ...*model.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a rolePermissionsTx) Replace(values ...*model.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a rolePermissionsTx) Delete(values ...*model.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a rolePermissionsTx) Clear() error {
	return a.tx.Clear()
}

func (a rolePermissionsTx) Count() int64 {
	return a.tx.Count()
}

type roleDo struct{ gen.DO }

func (r roleDo) Debug() *roleDo {
	return r.withDO(r.DO.Debug())
}

func (r roleDo) WithContext(ctx context.Context) *roleDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleDo) Clauses(conds ...clause.Expression) *roleDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleDo) Not(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleDo) Or(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleDo) Select(conds ...field.Expr) *roleDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleDo) Where(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleDo) Order(conds ...field.Expr) *roleDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleDo) Distinct(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleDo) Omit(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleDo) Join(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleDo) RightJoin(table schema.Tabler, on ...field.Expr) *roleDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleDo) Group(cols ...field.Expr) *roleDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleDo) Having(conds ...gen.Condition) *roleDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleDo) Limit(limit int) *roleDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleDo) Offset(offset int) *roleDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *roleDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleDo) Unscoped() *roleDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleDo) Create(values ...*model.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleDo) CreateInBatches(values []*model.Role, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleDo) Save(values ...*model.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleDo) First() (*model.Role, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Take() (*model.Role, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Last() (*model.Role, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) Find() ([]*model.Role, error) {
	result, err := r.DO.Find()
	return result.([]*model.Role), err
}

func (r roleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Role, err error) {
	buf := make([]*model.Role, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleDo) FindInBatches(result *[]*model.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleDo) Attrs(attrs ...field.AssignExpr) *roleDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleDo) Assign(attrs ...field.AssignExpr) *roleDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleDo) Joins(field field.RelationField) *roleDo {
	return r.withDO(r.DO.Joins(field))
}

func (r roleDo) Preload(field field.RelationField) *roleDo {
	return r.withDO(r.DO.Preload(field))
}

func (r roleDo) FirstOrInit() (*model.Role, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) FirstOrCreate() (*model.Role, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Role), nil
	}
}

func (r roleDo) FindByPage(offset int, limit int) (result []*model.Role, count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = r.Offset(offset).Limit(limit).Find()
	return
}

func (r roleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *roleDo) withDO(do gen.Dao) *roleDo {
	r.DO = *do.(*gen.DO)
	return r
}
