// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/fitan/magic/dao/dal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"
)

func newUser(db *gorm.DB) user {
	_user := user{}

	_user.userDo.UseDB(db)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewField(tableName, "*")
	_user.ID = field.NewUint(tableName, "id")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Name = field.NewString(tableName, "name")
	_user.Email = field.NewString(tableName, "email")
	_user.PassWord = field.NewString(tableName, "pass_word")
	_user.Token = field.NewString(tableName, "token")
	_user.Enable = field.NewBool(tableName, "enable")
	_user.Roles = userRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "model.Role"),
		Permissions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Roles.Permissions", "model.Permission"),
		},
	}

	_user.Services = userServices{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Services", "model.Service"),
		Services: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Services.Services", "model.Service"),
		},
	}

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL       field.Field
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Email     field.String
	PassWord  field.String
	Token     field.String
	Enable    field.Bool
	Roles     userRoles

	Services userServices

	fieldMap map[string]field.Expr
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))

	u.ALL = field.NewField(alias, "*")
	u.ID = field.NewUint(alias, "id")
	u.CreatedAt = field.NewTime(alias, "created_at")
	u.UpdatedAt = field.NewTime(alias, "updated_at")
	u.DeletedAt = field.NewField(alias, "deleted_at")
	u.Name = field.NewString(alias, "name")
	u.Email = field.NewString(alias, "email")
	u.PassWord = field.NewString(alias, "pass_word")
	u.Token = field.NewString(alias, "token")
	u.Enable = field.NewBool(alias, "enable")

	u.fillFieldMap()

	return &u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 11)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["name"] = u.Name
	u.fieldMap["email"] = u.Email
	u.fieldMap["pass_word"] = u.PassWord
	u.fieldMap["token"] = u.Token
	u.fieldMap["enable"] = u.Enable

}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userRoles struct {
	db *gorm.DB

	field.RelationField

	Permissions struct {
		field.RelationField
	}
}

func (a userRoles) Where(conds ...field.Expr) *userRoles {
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

func (a userRoles) WithContext(ctx context.Context) *userRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userRoles) Model(m *model.User) *userRolesTx {
	return &userRolesTx{a.db.Model(m).Association(a.Name())}
}

type userRolesTx struct{ tx *gorm.Association }

func (a userRolesTx) Find() (result *model.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a userRolesTx) Append(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userRolesTx) Replace(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userRolesTx) Delete(values ...*model.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a userRolesTx) Count() int64 {
	return a.tx.Count()
}

type userServices struct {
	db *gorm.DB

	field.RelationField

	Services struct {
		field.RelationField
	}
}

func (a userServices) Where(conds ...field.Expr) *userServices {
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

func (a userServices) WithContext(ctx context.Context) *userServices {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userServices) Model(m *model.User) *userServicesTx {
	return &userServicesTx{a.db.Model(m).Association(a.Name())}
}

type userServicesTx struct{ tx *gorm.Association }

func (a userServicesTx) Find() (result *model.Service, err error) {
	return result, a.tx.Find(&result)
}

func (a userServicesTx) Append(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userServicesTx) Replace(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userServicesTx) Delete(values ...*model.Service) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userServicesTx) Clear() error {
	return a.tx.Clear()
}

func (a userServicesTx) Count() int64 {
	return a.tx.Count()
}

type userDo struct{ gen.DO }

//where("id=@id")
func (u userDo) GetByID(id uint) (result *model.User, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = u.UnderlyingDB().Where(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = u.UnderlyingDB().Where(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

//where("email=@email and pass_word=@password")
func (u userDo) CheckAccount(email string, password string) (result *model.User, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["email"] = email
	params["password"] = password
	generateSQL.WriteString("email=@email and pass_word=@password ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = u.UnderlyingDB().Where(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = u.UnderlyingDB().Where(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

//update @@table {{set}} pass_word=@password {{end}} {{where}} id=@id {{end}}
func (u userDo) ModifyPassword(id int, password string) (err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	generateSQL.WriteString("update users ")
	var setSQL0 strings.Builder
	params["password"] = password
	setSQL0.WriteString("pass_word=@password ")
	helper.JoinSetBuilder(&generateSQL, setSQL0)
	var whereSQL0 strings.Builder
	params["id"] = id
	whereSQL0.WriteString("id=@id ")
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params)
	} else {
		executeSQL = u.UnderlyingDB().Exec(generateSQL.String())
	}
	err = executeSQL.Error
	return
}

//select * from @@table
func (u userDo) FindApi() (result []model.ApiUser, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select * from users ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Find(&result)
	err = executeSQL.Error
	return
}

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(field field.RelationField) *userDo {
	return u.withDO(u.DO.Joins(field))
}

func (u userDo) Preload(field field.RelationField) *userDo {
	return u.withDO(u.DO.Preload(field))
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}