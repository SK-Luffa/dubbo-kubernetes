// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/apache/dubbo-kubernetes/pkg/bufman/model"
)

func newRepository(db *gorm.DB, opts ...gen.DOOption) repository {
	_repository := repository{}

	_repository.repositoryDo.UseDB(db, opts...)
	_repository.repositoryDo.UseModel(&model.Repository{})

	tableName := _repository.repositoryDo.TableName()
	_repository.ALL = field.NewAsterisk(tableName)
	_repository.ID = field.NewInt64(tableName, "id")
	_repository.UserID = field.NewString(tableName, "user_id")
	_repository.UserName = field.NewString(tableName, "user_name")
	_repository.RepositoryID = field.NewString(tableName, "repository_id")
	_repository.RepositoryName = field.NewString(tableName, "repository_name")
	_repository.CreatedTime = field.NewTime(tableName, "created_time")
	_repository.UpdateTime = field.NewTime(tableName, "update_time")
	_repository.Visibility = field.NewUint8(tableName, "visibility")
	_repository.Deprecated = field.NewBool(tableName, "deprecated")
	_repository.DeprecationMsg = field.NewString(tableName, "deprecation_msg")
	_repository.Url = field.NewString(tableName, "url")
	_repository.Description = field.NewString(tableName, "description")

	_repository.fillFieldMap()

	return _repository
}

type repository struct {
	repositoryDo

	ALL            field.Asterisk
	ID             field.Int64
	UserID         field.String
	UserName       field.String
	RepositoryID   field.String
	RepositoryName field.String
	CreatedTime    field.Time
	UpdateTime     field.Time
	Visibility     field.Uint8
	Deprecated     field.Bool
	DeprecationMsg field.String
	Url            field.String
	Description    field.String

	fieldMap map[string]field.Expr
}

func (r repository) Table(newTableName string) *repository {
	r.repositoryDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r repository) As(alias string) *repository {
	r.repositoryDo.DO = *(r.repositoryDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *repository) updateTableName(table string) *repository {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.UserID = field.NewString(table, "user_id")
	r.UserName = field.NewString(table, "user_name")
	r.RepositoryID = field.NewString(table, "repository_id")
	r.RepositoryName = field.NewString(table, "repository_name")
	r.CreatedTime = field.NewTime(table, "created_time")
	r.UpdateTime = field.NewTime(table, "update_time")
	r.Visibility = field.NewUint8(table, "visibility")
	r.Deprecated = field.NewBool(table, "deprecated")
	r.DeprecationMsg = field.NewString(table, "deprecation_msg")
	r.Url = field.NewString(table, "url")
	r.Description = field.NewString(table, "description")

	r.fillFieldMap()

	return r
}

func (r *repository) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *repository) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 12)
	r.fieldMap["id"] = r.ID
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["user_name"] = r.UserName
	r.fieldMap["repository_id"] = r.RepositoryID
	r.fieldMap["repository_name"] = r.RepositoryName
	r.fieldMap["created_time"] = r.CreatedTime
	r.fieldMap["update_time"] = r.UpdateTime
	r.fieldMap["visibility"] = r.Visibility
	r.fieldMap["deprecated"] = r.Deprecated
	r.fieldMap["deprecation_msg"] = r.DeprecationMsg
	r.fieldMap["url"] = r.Url
	r.fieldMap["description"] = r.Description
}

func (r repository) clone(db *gorm.DB) repository {
	r.repositoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r repository) replaceDB(db *gorm.DB) repository {
	r.repositoryDo.ReplaceDB(db)
	return r
}

type repositoryDo struct{ gen.DO }

type IRepositoryDo interface {
	gen.SubQuery
	Debug() IRepositoryDo
	WithContext(ctx context.Context) IRepositoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRepositoryDo
	WriteDB() IRepositoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRepositoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRepositoryDo
	Not(conds ...gen.Condition) IRepositoryDo
	Or(conds ...gen.Condition) IRepositoryDo
	Select(conds ...field.Expr) IRepositoryDo
	Where(conds ...gen.Condition) IRepositoryDo
	Order(conds ...field.Expr) IRepositoryDo
	Distinct(cols ...field.Expr) IRepositoryDo
	Omit(cols ...field.Expr) IRepositoryDo
	Join(table schema.Tabler, on ...field.Expr) IRepositoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRepositoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRepositoryDo
	Group(cols ...field.Expr) IRepositoryDo
	Having(conds ...gen.Condition) IRepositoryDo
	Limit(limit int) IRepositoryDo
	Offset(offset int) IRepositoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRepositoryDo
	Unscoped() IRepositoryDo
	Create(values ...*model.Repository) error
	CreateInBatches(values []*model.Repository, batchSize int) error
	Save(values ...*model.Repository) error
	First() (*model.Repository, error)
	Take() (*model.Repository, error)
	Last() (*model.Repository, error)
	Find() ([]*model.Repository, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Repository, err error)
	FindInBatches(result *[]*model.Repository, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Repository) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRepositoryDo
	Assign(attrs ...field.AssignExpr) IRepositoryDo
	Joins(fields ...field.RelationField) IRepositoryDo
	Preload(fields ...field.RelationField) IRepositoryDo
	FirstOrInit() (*model.Repository, error)
	FirstOrCreate() (*model.Repository, error)
	FindByPage(offset int, limit int) (result []*model.Repository, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRepositoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r repositoryDo) Debug() IRepositoryDo {
	return r.withDO(r.DO.Debug())
}

func (r repositoryDo) WithContext(ctx context.Context) IRepositoryDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r repositoryDo) ReadDB() IRepositoryDo {
	return r.Clauses(dbresolver.Read)
}

func (r repositoryDo) WriteDB() IRepositoryDo {
	return r.Clauses(dbresolver.Write)
}

func (r repositoryDo) Session(config *gorm.Session) IRepositoryDo {
	return r.withDO(r.DO.Session(config))
}

func (r repositoryDo) Clauses(conds ...clause.Expression) IRepositoryDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r repositoryDo) Returning(value interface{}, columns ...string) IRepositoryDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r repositoryDo) Not(conds ...gen.Condition) IRepositoryDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r repositoryDo) Or(conds ...gen.Condition) IRepositoryDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r repositoryDo) Select(conds ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r repositoryDo) Where(conds ...gen.Condition) IRepositoryDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r repositoryDo) Order(conds ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r repositoryDo) Distinct(cols ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r repositoryDo) Omit(cols ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r repositoryDo) Join(table schema.Tabler, on ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r repositoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r repositoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r repositoryDo) Group(cols ...field.Expr) IRepositoryDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r repositoryDo) Having(conds ...gen.Condition) IRepositoryDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r repositoryDo) Limit(limit int) IRepositoryDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r repositoryDo) Offset(offset int) IRepositoryDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r repositoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRepositoryDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r repositoryDo) Unscoped() IRepositoryDo {
	return r.withDO(r.DO.Unscoped())
}

func (r repositoryDo) Create(values ...*model.Repository) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r repositoryDo) CreateInBatches(values []*model.Repository, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r repositoryDo) Save(values ...*model.Repository) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r repositoryDo) First() (*model.Repository, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repository), nil
	}
}

func (r repositoryDo) Take() (*model.Repository, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repository), nil
	}
}

func (r repositoryDo) Last() (*model.Repository, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repository), nil
	}
}

func (r repositoryDo) Find() ([]*model.Repository, error) {
	result, err := r.DO.Find()
	return result.([]*model.Repository), err
}

func (r repositoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Repository, err error) {
	buf := make([]*model.Repository, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r repositoryDo) FindInBatches(result *[]*model.Repository, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r repositoryDo) Attrs(attrs ...field.AssignExpr) IRepositoryDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r repositoryDo) Assign(attrs ...field.AssignExpr) IRepositoryDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r repositoryDo) Joins(fields ...field.RelationField) IRepositoryDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r repositoryDo) Preload(fields ...field.RelationField) IRepositoryDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r repositoryDo) FirstOrInit() (*model.Repository, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repository), nil
	}
}

func (r repositoryDo) FirstOrCreate() (*model.Repository, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Repository), nil
	}
}

func (r repositoryDo) FindByPage(offset int, limit int) (result []*model.Repository, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r repositoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r repositoryDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r repositoryDo) Delete(models ...*model.Repository) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *repositoryDo) withDO(do gen.Dao) *repositoryDo {
	r.DO = *do.(*gen.DO)
	return r
}
