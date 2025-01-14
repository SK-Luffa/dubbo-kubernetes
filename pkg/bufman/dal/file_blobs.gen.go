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

func newFileBlob(db *gorm.DB, opts ...gen.DOOption) fileBlob {
	_fileBlob := fileBlob{}

	_fileBlob.fileBlobDo.UseDB(db, opts...)
	_fileBlob.fileBlobDo.UseModel(&model.FileBlob{})

	tableName := _fileBlob.fileBlobDo.TableName()
	_fileBlob.ALL = field.NewAsterisk(tableName)
	_fileBlob.ID = field.NewInt64(tableName, "id")
	_fileBlob.Digest = field.NewString(tableName, "digest")
	_fileBlob.Content = field.NewBytes(tableName, "content")

	_fileBlob.fillFieldMap()

	return _fileBlob
}

type fileBlob struct {
	fileBlobDo

	ALL     field.Asterisk
	ID      field.Int64
	Digest  field.String
	Content field.Bytes

	fieldMap map[string]field.Expr
}

func (f fileBlob) Table(newTableName string) *fileBlob {
	f.fileBlobDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fileBlob) As(alias string) *fileBlob {
	f.fileBlobDo.DO = *(f.fileBlobDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fileBlob) updateTableName(table string) *fileBlob {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.Digest = field.NewString(table, "digest")
	f.Content = field.NewBytes(table, "content")

	f.fillFieldMap()

	return f
}

func (f *fileBlob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fileBlob) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["id"] = f.ID
	f.fieldMap["digest"] = f.Digest
	f.fieldMap["content"] = f.Content
}

func (f fileBlob) clone(db *gorm.DB) fileBlob {
	f.fileBlobDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fileBlob) replaceDB(db *gorm.DB) fileBlob {
	f.fileBlobDo.ReplaceDB(db)
	return f
}

type fileBlobDo struct{ gen.DO }

type IFileBlobDo interface {
	gen.SubQuery
	Debug() IFileBlobDo
	WithContext(ctx context.Context) IFileBlobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFileBlobDo
	WriteDB() IFileBlobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFileBlobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFileBlobDo
	Not(conds ...gen.Condition) IFileBlobDo
	Or(conds ...gen.Condition) IFileBlobDo
	Select(conds ...field.Expr) IFileBlobDo
	Where(conds ...gen.Condition) IFileBlobDo
	Order(conds ...field.Expr) IFileBlobDo
	Distinct(cols ...field.Expr) IFileBlobDo
	Omit(cols ...field.Expr) IFileBlobDo
	Join(table schema.Tabler, on ...field.Expr) IFileBlobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFileBlobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFileBlobDo
	Group(cols ...field.Expr) IFileBlobDo
	Having(conds ...gen.Condition) IFileBlobDo
	Limit(limit int) IFileBlobDo
	Offset(offset int) IFileBlobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFileBlobDo
	Unscoped() IFileBlobDo
	Create(values ...*model.FileBlob) error
	CreateInBatches(values []*model.FileBlob, batchSize int) error
	Save(values ...*model.FileBlob) error
	First() (*model.FileBlob, error)
	Take() (*model.FileBlob, error)
	Last() (*model.FileBlob, error)
	Find() ([]*model.FileBlob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileBlob, err error)
	FindInBatches(result *[]*model.FileBlob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FileBlob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFileBlobDo
	Assign(attrs ...field.AssignExpr) IFileBlobDo
	Joins(fields ...field.RelationField) IFileBlobDo
	Preload(fields ...field.RelationField) IFileBlobDo
	FirstOrInit() (*model.FileBlob, error)
	FirstOrCreate() (*model.FileBlob, error)
	FindByPage(offset int, limit int) (result []*model.FileBlob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFileBlobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fileBlobDo) Debug() IFileBlobDo {
	return f.withDO(f.DO.Debug())
}

func (f fileBlobDo) WithContext(ctx context.Context) IFileBlobDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileBlobDo) ReadDB() IFileBlobDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileBlobDo) WriteDB() IFileBlobDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileBlobDo) Session(config *gorm.Session) IFileBlobDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileBlobDo) Clauses(conds ...clause.Expression) IFileBlobDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileBlobDo) Returning(value interface{}, columns ...string) IFileBlobDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileBlobDo) Not(conds ...gen.Condition) IFileBlobDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileBlobDo) Or(conds ...gen.Condition) IFileBlobDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileBlobDo) Select(conds ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileBlobDo) Where(conds ...gen.Condition) IFileBlobDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileBlobDo) Order(conds ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileBlobDo) Distinct(cols ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileBlobDo) Omit(cols ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileBlobDo) Join(table schema.Tabler, on ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileBlobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileBlobDo) RightJoin(table schema.Tabler, on ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileBlobDo) Group(cols ...field.Expr) IFileBlobDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileBlobDo) Having(conds ...gen.Condition) IFileBlobDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileBlobDo) Limit(limit int) IFileBlobDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileBlobDo) Offset(offset int) IFileBlobDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileBlobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFileBlobDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileBlobDo) Unscoped() IFileBlobDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileBlobDo) Create(values ...*model.FileBlob) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileBlobDo) CreateInBatches(values []*model.FileBlob, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileBlobDo) Save(values ...*model.FileBlob) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileBlobDo) First() (*model.FileBlob, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileBlob), nil
	}
}

func (f fileBlobDo) Take() (*model.FileBlob, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileBlob), nil
	}
}

func (f fileBlobDo) Last() (*model.FileBlob, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileBlob), nil
	}
}

func (f fileBlobDo) Find() ([]*model.FileBlob, error) {
	result, err := f.DO.Find()
	return result.([]*model.FileBlob), err
}

func (f fileBlobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileBlob, err error) {
	buf := make([]*model.FileBlob, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileBlobDo) FindInBatches(result *[]*model.FileBlob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileBlobDo) Attrs(attrs ...field.AssignExpr) IFileBlobDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileBlobDo) Assign(attrs ...field.AssignExpr) IFileBlobDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileBlobDo) Joins(fields ...field.RelationField) IFileBlobDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileBlobDo) Preload(fields ...field.RelationField) IFileBlobDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileBlobDo) FirstOrInit() (*model.FileBlob, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileBlob), nil
	}
}

func (f fileBlobDo) FirstOrCreate() (*model.FileBlob, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileBlob), nil
	}
}

func (f fileBlobDo) FindByPage(offset int, limit int) (result []*model.FileBlob, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileBlobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileBlobDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileBlobDo) Delete(models ...*model.FileBlob) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileBlobDo) withDO(do gen.Dao) *fileBlobDo {
	f.DO = *do.(*gen.DO)
	return f
}
