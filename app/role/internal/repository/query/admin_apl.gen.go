// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"role/internal/repository"
)

func newAdminAPL(db *gorm.DB, opts ...gen.DOOption) adminAPL {
	_adminAPL := adminAPL{}

	_adminAPL.adminAPLDo.UseDB(db, opts...)
	_adminAPL.adminAPLDo.UseModel(&repository.AdminAPL{})

	tableName := _adminAPL.adminAPLDo.TableName()
	_adminAPL.ALL = field.NewAsterisk(tableName)
	_adminAPL.ID = field.NewUint(tableName, "id")
	_adminAPL.CreatedAt = field.NewTime(tableName, "created_at")
	_adminAPL.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminAPL.DeletedAt = field.NewField(tableName, "deleted_at")
	_adminAPL.UserId = field.NewUint(tableName, "user_id")
	_adminAPL.Status = field.NewUint(tableName, "status")
	_adminAPL.APLComment = field.NewString(tableName, "apl_comment")
	_adminAPL.REVComment = field.NewString(tableName, "rev_comment")

	_adminAPL.fillFieldMap()

	return _adminAPL
}

type adminAPL struct {
	adminAPLDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	UserId     field.Uint
	Status     field.Uint
	APLComment field.String
	REVComment field.String

	fieldMap map[string]field.Expr
}

func (a adminAPL) Table(newTableName string) *adminAPL {
	a.adminAPLDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminAPL) As(alias string) *adminAPL {
	a.adminAPLDo.DO = *(a.adminAPLDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminAPL) updateTableName(table string) *adminAPL {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.UserId = field.NewUint(table, "user_id")
	a.Status = field.NewUint(table, "status")
	a.APLComment = field.NewString(table, "apl_comment")
	a.REVComment = field.NewString(table, "rev_comment")

	a.fillFieldMap()

	return a
}

func (a *adminAPL) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminAPL) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["user_id"] = a.UserId
	a.fieldMap["status"] = a.Status
	a.fieldMap["apl_comment"] = a.APLComment
	a.fieldMap["rev_comment"] = a.REVComment
}

func (a adminAPL) clone(db *gorm.DB) adminAPL {
	a.adminAPLDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminAPL) replaceDB(db *gorm.DB) adminAPL {
	a.adminAPLDo.ReplaceDB(db)
	return a
}

type adminAPLDo struct{ gen.DO }

type IAdminAPLDo interface {
	gen.SubQuery
	Debug() IAdminAPLDo
	WithContext(ctx context.Context) IAdminAPLDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminAPLDo
	WriteDB() IAdminAPLDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminAPLDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminAPLDo
	Not(conds ...gen.Condition) IAdminAPLDo
	Or(conds ...gen.Condition) IAdminAPLDo
	Select(conds ...field.Expr) IAdminAPLDo
	Where(conds ...gen.Condition) IAdminAPLDo
	Order(conds ...field.Expr) IAdminAPLDo
	Distinct(cols ...field.Expr) IAdminAPLDo
	Omit(cols ...field.Expr) IAdminAPLDo
	Join(table schema.Tabler, on ...field.Expr) IAdminAPLDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminAPLDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminAPLDo
	Group(cols ...field.Expr) IAdminAPLDo
	Having(conds ...gen.Condition) IAdminAPLDo
	Limit(limit int) IAdminAPLDo
	Offset(offset int) IAdminAPLDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminAPLDo
	Unscoped() IAdminAPLDo
	Create(values ...*repository.AdminAPL) error
	CreateInBatches(values []*repository.AdminAPL, batchSize int) error
	Save(values ...*repository.AdminAPL) error
	First() (*repository.AdminAPL, error)
	Take() (*repository.AdminAPL, error)
	Last() (*repository.AdminAPL, error)
	Find() ([]*repository.AdminAPL, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*repository.AdminAPL, err error)
	FindInBatches(result *[]*repository.AdminAPL, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*repository.AdminAPL) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminAPLDo
	Assign(attrs ...field.AssignExpr) IAdminAPLDo
	Joins(fields ...field.RelationField) IAdminAPLDo
	Preload(fields ...field.RelationField) IAdminAPLDo
	FirstOrInit() (*repository.AdminAPL, error)
	FirstOrCreate() (*repository.AdminAPL, error)
	FindByPage(offset int, limit int) (result []*repository.AdminAPL, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminAPLDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminAPLDo) Debug() IAdminAPLDo {
	return a.withDO(a.DO.Debug())
}

func (a adminAPLDo) WithContext(ctx context.Context) IAdminAPLDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminAPLDo) ReadDB() IAdminAPLDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminAPLDo) WriteDB() IAdminAPLDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminAPLDo) Session(config *gorm.Session) IAdminAPLDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminAPLDo) Clauses(conds ...clause.Expression) IAdminAPLDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminAPLDo) Returning(value interface{}, columns ...string) IAdminAPLDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminAPLDo) Not(conds ...gen.Condition) IAdminAPLDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminAPLDo) Or(conds ...gen.Condition) IAdminAPLDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminAPLDo) Select(conds ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminAPLDo) Where(conds ...gen.Condition) IAdminAPLDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminAPLDo) Order(conds ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminAPLDo) Distinct(cols ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminAPLDo) Omit(cols ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminAPLDo) Join(table schema.Tabler, on ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminAPLDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminAPLDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminAPLDo) Group(cols ...field.Expr) IAdminAPLDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminAPLDo) Having(conds ...gen.Condition) IAdminAPLDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminAPLDo) Limit(limit int) IAdminAPLDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminAPLDo) Offset(offset int) IAdminAPLDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminAPLDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminAPLDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminAPLDo) Unscoped() IAdminAPLDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminAPLDo) Create(values ...*repository.AdminAPL) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminAPLDo) CreateInBatches(values []*repository.AdminAPL, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminAPLDo) Save(values ...*repository.AdminAPL) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminAPLDo) First() (*repository.AdminAPL, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*repository.AdminAPL), nil
	}
}

func (a adminAPLDo) Take() (*repository.AdminAPL, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*repository.AdminAPL), nil
	}
}

func (a adminAPLDo) Last() (*repository.AdminAPL, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*repository.AdminAPL), nil
	}
}

func (a adminAPLDo) Find() ([]*repository.AdminAPL, error) {
	result, err := a.DO.Find()
	return result.([]*repository.AdminAPL), err
}

func (a adminAPLDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*repository.AdminAPL, err error) {
	buf := make([]*repository.AdminAPL, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminAPLDo) FindInBatches(result *[]*repository.AdminAPL, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminAPLDo) Attrs(attrs ...field.AssignExpr) IAdminAPLDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminAPLDo) Assign(attrs ...field.AssignExpr) IAdminAPLDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminAPLDo) Joins(fields ...field.RelationField) IAdminAPLDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminAPLDo) Preload(fields ...field.RelationField) IAdminAPLDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminAPLDo) FirstOrInit() (*repository.AdminAPL, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*repository.AdminAPL), nil
	}
}

func (a adminAPLDo) FirstOrCreate() (*repository.AdminAPL, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*repository.AdminAPL), nil
	}
}

func (a adminAPLDo) FindByPage(offset int, limit int) (result []*repository.AdminAPL, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminAPLDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminAPLDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminAPLDo) Delete(models ...*repository.AdminAPL) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminAPLDo) withDO(do gen.Dao) *adminAPLDo {
	a.DO = *do.(*gen.DO)
	return a
}
