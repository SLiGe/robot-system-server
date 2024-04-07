// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"robot-system-server/internal/model"
)

func newQrFortune(db *gorm.DB, opts ...gen.DOOption) qrFortune {
	_qrFortune := qrFortune{}

	_qrFortune.qrFortuneDo.UseDB(db, opts...)
	_qrFortune.qrFortuneDo.UseModel(&model.QrFortune{})

	tableName := _qrFortune.qrFortuneDo.TableName()
	_qrFortune.ALL = field.NewAsterisk(tableName)
	_qrFortune.ID = field.NewInt64(tableName, "id")
	_qrFortune.FortuneSummary = field.NewString(tableName, "fortune_summary")
	_qrFortune.LuckyStar = field.NewString(tableName, "lucky_star")
	_qrFortune.SignText = field.NewString(tableName, "sign_text")
	_qrFortune.UnSignText = field.NewString(tableName, "un_sign_text")

	_qrFortune.fillFieldMap()

	return _qrFortune
}

// qrFortune 运势表
type qrFortune struct {
	qrFortuneDo

	ALL            field.Asterisk
	ID             field.Int64
	FortuneSummary field.String // 运情总结
	LuckyStar      field.String // 幸运星
	SignText       field.String // 签文
	UnSignText     field.String // 解签

	fieldMap map[string]field.Expr
}

func (q qrFortune) Table(newTableName string) *qrFortune {
	q.qrFortuneDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qrFortune) As(alias string) *qrFortune {
	q.qrFortuneDo.DO = *(q.qrFortuneDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qrFortune) updateTableName(table string) *qrFortune {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.FortuneSummary = field.NewString(table, "fortune_summary")
	q.LuckyStar = field.NewString(table, "lucky_star")
	q.SignText = field.NewString(table, "sign_text")
	q.UnSignText = field.NewString(table, "un_sign_text")

	q.fillFieldMap()

	return q
}

func (q *qrFortune) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qrFortune) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 5)
	q.fieldMap["id"] = q.ID
	q.fieldMap["fortune_summary"] = q.FortuneSummary
	q.fieldMap["lucky_star"] = q.LuckyStar
	q.fieldMap["sign_text"] = q.SignText
	q.fieldMap["un_sign_text"] = q.UnSignText
}

func (q qrFortune) clone(db *gorm.DB) qrFortune {
	q.qrFortuneDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qrFortune) replaceDB(db *gorm.DB) qrFortune {
	q.qrFortuneDo.ReplaceDB(db)
	return q
}

type qrFortuneDo struct{ gen.DO }

type IQrFortuneDo interface {
	gen.SubQuery
	Debug() IQrFortuneDo
	WithContext(ctx context.Context) IQrFortuneDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQrFortuneDo
	WriteDB() IQrFortuneDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQrFortuneDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQrFortuneDo
	Not(conds ...gen.Condition) IQrFortuneDo
	Or(conds ...gen.Condition) IQrFortuneDo
	Select(conds ...field.Expr) IQrFortuneDo
	Where(conds ...gen.Condition) IQrFortuneDo
	Order(conds ...field.Expr) IQrFortuneDo
	Distinct(cols ...field.Expr) IQrFortuneDo
	Omit(cols ...field.Expr) IQrFortuneDo
	Join(table schema.Tabler, on ...field.Expr) IQrFortuneDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQrFortuneDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQrFortuneDo
	Group(cols ...field.Expr) IQrFortuneDo
	Having(conds ...gen.Condition) IQrFortuneDo
	Limit(limit int) IQrFortuneDo
	Offset(offset int) IQrFortuneDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQrFortuneDo
	Unscoped() IQrFortuneDo
	Create(values ...*model.QrFortune) error
	CreateInBatches(values []*model.QrFortune, batchSize int) error
	Save(values ...*model.QrFortune) error
	First() (*model.QrFortune, error)
	Take() (*model.QrFortune, error)
	Last() (*model.QrFortune, error)
	Find() ([]*model.QrFortune, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrFortune, err error)
	FindInBatches(result *[]*model.QrFortune, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QrFortune) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQrFortuneDo
	Assign(attrs ...field.AssignExpr) IQrFortuneDo
	Joins(fields ...field.RelationField) IQrFortuneDo
	Preload(fields ...field.RelationField) IQrFortuneDo
	FirstOrInit() (*model.QrFortune, error)
	FirstOrCreate() (*model.QrFortune, error)
	FindByPage(offset int, limit int) (result []*model.QrFortune, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQrFortuneDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qrFortuneDo) Debug() IQrFortuneDo {
	return q.withDO(q.DO.Debug())
}

func (q qrFortuneDo) WithContext(ctx context.Context) IQrFortuneDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qrFortuneDo) ReadDB() IQrFortuneDo {
	return q.Clauses(dbresolver.Read)
}

func (q qrFortuneDo) WriteDB() IQrFortuneDo {
	return q.Clauses(dbresolver.Write)
}

func (q qrFortuneDo) Session(config *gorm.Session) IQrFortuneDo {
	return q.withDO(q.DO.Session(config))
}

func (q qrFortuneDo) Clauses(conds ...clause.Expression) IQrFortuneDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qrFortuneDo) Returning(value interface{}, columns ...string) IQrFortuneDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qrFortuneDo) Not(conds ...gen.Condition) IQrFortuneDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qrFortuneDo) Or(conds ...gen.Condition) IQrFortuneDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qrFortuneDo) Select(conds ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qrFortuneDo) Where(conds ...gen.Condition) IQrFortuneDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qrFortuneDo) Order(conds ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qrFortuneDo) Distinct(cols ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qrFortuneDo) Omit(cols ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qrFortuneDo) Join(table schema.Tabler, on ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qrFortuneDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qrFortuneDo) RightJoin(table schema.Tabler, on ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qrFortuneDo) Group(cols ...field.Expr) IQrFortuneDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qrFortuneDo) Having(conds ...gen.Condition) IQrFortuneDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qrFortuneDo) Limit(limit int) IQrFortuneDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qrFortuneDo) Offset(offset int) IQrFortuneDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qrFortuneDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQrFortuneDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qrFortuneDo) Unscoped() IQrFortuneDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qrFortuneDo) Create(values ...*model.QrFortune) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qrFortuneDo) CreateInBatches(values []*model.QrFortune, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qrFortuneDo) Save(values ...*model.QrFortune) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qrFortuneDo) First() (*model.QrFortune, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrFortune), nil
	}
}

func (q qrFortuneDo) Take() (*model.QrFortune, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrFortune), nil
	}
}

func (q qrFortuneDo) Last() (*model.QrFortune, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrFortune), nil
	}
}

func (q qrFortuneDo) Find() ([]*model.QrFortune, error) {
	result, err := q.DO.Find()
	return result.([]*model.QrFortune), err
}

func (q qrFortuneDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrFortune, err error) {
	buf := make([]*model.QrFortune, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qrFortuneDo) FindInBatches(result *[]*model.QrFortune, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qrFortuneDo) Attrs(attrs ...field.AssignExpr) IQrFortuneDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qrFortuneDo) Assign(attrs ...field.AssignExpr) IQrFortuneDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qrFortuneDo) Joins(fields ...field.RelationField) IQrFortuneDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qrFortuneDo) Preload(fields ...field.RelationField) IQrFortuneDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qrFortuneDo) FirstOrInit() (*model.QrFortune, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrFortune), nil
	}
}

func (q qrFortuneDo) FirstOrCreate() (*model.QrFortune, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrFortune), nil
	}
}

func (q qrFortuneDo) FindByPage(offset int, limit int) (result []*model.QrFortune, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q qrFortuneDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qrFortuneDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qrFortuneDo) Delete(models ...*model.QrFortune) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qrFortuneDo) withDO(do gen.Dao) *qrFortuneDo {
	q.DO = *do.(*gen.DO)
	return q
}
