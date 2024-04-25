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

func newQrSpiritSignUDatum(db *gorm.DB, opts ...gen.DOOption) qrSpiritSignUDatum {
	_qrSpiritSignUDatum := qrSpiritSignUDatum{}

	_qrSpiritSignUDatum.qrSpiritSignUDatumDo.UseDB(db, opts...)
	_qrSpiritSignUDatum.qrSpiritSignUDatumDo.UseModel(&model.QrSpiritSignUDatum{})

	tableName := _qrSpiritSignUDatum.qrSpiritSignUDatumDo.TableName()
	_qrSpiritSignUDatum.ALL = field.NewAsterisk(tableName)
	_qrSpiritSignUDatum.ID = field.NewInt64(tableName, "id")
	_qrSpiritSignUDatum.UserID = field.NewInt64(tableName, "user_id")
	_qrSpiritSignUDatum.UserQq = field.NewString(tableName, "user_qq")
	_qrSpiritSignUDatum.SignType = field.NewString(tableName, "sign_type")
	_qrSpiritSignUDatum.SignID = field.NewInt64(tableName, "sign_id")
	_qrSpiritSignUDatum.SignDate = field.NewField(tableName, "sign_date")
	_qrSpiritSignUDatum.CreateBy = field.NewString(tableName, "create_by")
	_qrSpiritSignUDatum.CreateTime = field.NewField(tableName, "create_time")
	_qrSpiritSignUDatum.UpdateBy = field.NewString(tableName, "update_by")
	_qrSpiritSignUDatum.UpdateTime = field.NewField(tableName, "update_time")

	_qrSpiritSignUDatum.fillFieldMap()

	return _qrSpiritSignUDatum
}

// qrSpiritSignUDatum 用户灵签数据
type qrSpiritSignUDatum struct {
	qrSpiritSignUDatumDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	UserID     field.Int64  // 用户ID
	UserQq     field.String // 用户QQ
	SignType   field.String // 灵签类型
	SignID     field.Int64  // 灵签表ID
	SignDate   field.Field  // 灵签时间
	CreateBy   field.String // 创建人
	CreateTime field.Field  // 创建时间
	UpdateBy   field.String // 更新人
	UpdateTime field.Field  // 更新时间

	fieldMap map[string]field.Expr
}

func (q qrSpiritSignUDatum) Table(newTableName string) *qrSpiritSignUDatum {
	q.qrSpiritSignUDatumDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qrSpiritSignUDatum) As(alias string) *qrSpiritSignUDatum {
	q.qrSpiritSignUDatumDo.DO = *(q.qrSpiritSignUDatumDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qrSpiritSignUDatum) updateTableName(table string) *qrSpiritSignUDatum {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.UserID = field.NewInt64(table, "user_id")
	q.UserQq = field.NewString(table, "user_qq")
	q.SignType = field.NewString(table, "sign_type")
	q.SignID = field.NewInt64(table, "sign_id")
	q.SignDate = field.NewField(table, "sign_date")
	q.CreateBy = field.NewString(table, "create_by")
	q.CreateTime = field.NewField(table, "create_time")
	q.UpdateBy = field.NewString(table, "update_by")
	q.UpdateTime = field.NewField(table, "update_time")

	q.fillFieldMap()

	return q
}

func (q *qrSpiritSignUDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qrSpiritSignUDatum) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 10)
	q.fieldMap["id"] = q.ID
	q.fieldMap["user_id"] = q.UserID
	q.fieldMap["user_qq"] = q.UserQq
	q.fieldMap["sign_type"] = q.SignType
	q.fieldMap["sign_id"] = q.SignID
	q.fieldMap["sign_date"] = q.SignDate
	q.fieldMap["create_by"] = q.CreateBy
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_by"] = q.UpdateBy
	q.fieldMap["update_time"] = q.UpdateTime
}

func (q qrSpiritSignUDatum) clone(db *gorm.DB) qrSpiritSignUDatum {
	q.qrSpiritSignUDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qrSpiritSignUDatum) replaceDB(db *gorm.DB) qrSpiritSignUDatum {
	q.qrSpiritSignUDatumDo.ReplaceDB(db)
	return q
}

type qrSpiritSignUDatumDo struct{ gen.DO }

type IQrSpiritSignUDatumDo interface {
	gen.SubQuery
	Debug() IQrSpiritSignUDatumDo
	WithContext(ctx context.Context) IQrSpiritSignUDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQrSpiritSignUDatumDo
	WriteDB() IQrSpiritSignUDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQrSpiritSignUDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQrSpiritSignUDatumDo
	Not(conds ...gen.Condition) IQrSpiritSignUDatumDo
	Or(conds ...gen.Condition) IQrSpiritSignUDatumDo
	Select(conds ...field.Expr) IQrSpiritSignUDatumDo
	Where(conds ...gen.Condition) IQrSpiritSignUDatumDo
	Order(conds ...field.Expr) IQrSpiritSignUDatumDo
	Distinct(cols ...field.Expr) IQrSpiritSignUDatumDo
	Omit(cols ...field.Expr) IQrSpiritSignUDatumDo
	Join(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo
	Group(cols ...field.Expr) IQrSpiritSignUDatumDo
	Having(conds ...gen.Condition) IQrSpiritSignUDatumDo
	Limit(limit int) IQrSpiritSignUDatumDo
	Offset(offset int) IQrSpiritSignUDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSpiritSignUDatumDo
	Unscoped() IQrSpiritSignUDatumDo
	Create(values ...*model.QrSpiritSignUDatum) error
	CreateInBatches(values []*model.QrSpiritSignUDatum, batchSize int) error
	Save(values ...*model.QrSpiritSignUDatum) error
	First() (*model.QrSpiritSignUDatum, error)
	Take() (*model.QrSpiritSignUDatum, error)
	Last() (*model.QrSpiritSignUDatum, error)
	Find() ([]*model.QrSpiritSignUDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSpiritSignUDatum, err error)
	FindInBatches(result *[]*model.QrSpiritSignUDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QrSpiritSignUDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQrSpiritSignUDatumDo
	Assign(attrs ...field.AssignExpr) IQrSpiritSignUDatumDo
	Joins(fields ...field.RelationField) IQrSpiritSignUDatumDo
	Preload(fields ...field.RelationField) IQrSpiritSignUDatumDo
	FirstOrInit() (*model.QrSpiritSignUDatum, error)
	FirstOrCreate() (*model.QrSpiritSignUDatum, error)
	FindByPage(offset int, limit int) (result []*model.QrSpiritSignUDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQrSpiritSignUDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qrSpiritSignUDatumDo) Debug() IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Debug())
}

func (q qrSpiritSignUDatumDo) WithContext(ctx context.Context) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qrSpiritSignUDatumDo) ReadDB() IQrSpiritSignUDatumDo {
	return q.Clauses(dbresolver.Read)
}

func (q qrSpiritSignUDatumDo) WriteDB() IQrSpiritSignUDatumDo {
	return q.Clauses(dbresolver.Write)
}

func (q qrSpiritSignUDatumDo) Session(config *gorm.Session) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Session(config))
}

func (q qrSpiritSignUDatumDo) Clauses(conds ...clause.Expression) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qrSpiritSignUDatumDo) Returning(value interface{}, columns ...string) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qrSpiritSignUDatumDo) Not(conds ...gen.Condition) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qrSpiritSignUDatumDo) Or(conds ...gen.Condition) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qrSpiritSignUDatumDo) Select(conds ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qrSpiritSignUDatumDo) Where(conds ...gen.Condition) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qrSpiritSignUDatumDo) Order(conds ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qrSpiritSignUDatumDo) Distinct(cols ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qrSpiritSignUDatumDo) Omit(cols ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qrSpiritSignUDatumDo) Join(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qrSpiritSignUDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qrSpiritSignUDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qrSpiritSignUDatumDo) Group(cols ...field.Expr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qrSpiritSignUDatumDo) Having(conds ...gen.Condition) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qrSpiritSignUDatumDo) Limit(limit int) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qrSpiritSignUDatumDo) Offset(offset int) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qrSpiritSignUDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qrSpiritSignUDatumDo) Unscoped() IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qrSpiritSignUDatumDo) Create(values ...*model.QrSpiritSignUDatum) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qrSpiritSignUDatumDo) CreateInBatches(values []*model.QrSpiritSignUDatum, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qrSpiritSignUDatumDo) Save(values ...*model.QrSpiritSignUDatum) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qrSpiritSignUDatumDo) First() (*model.QrSpiritSignUDatum, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSpiritSignUDatum), nil
	}
}

func (q qrSpiritSignUDatumDo) Take() (*model.QrSpiritSignUDatum, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSpiritSignUDatum), nil
	}
}

func (q qrSpiritSignUDatumDo) Last() (*model.QrSpiritSignUDatum, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSpiritSignUDatum), nil
	}
}

func (q qrSpiritSignUDatumDo) Find() ([]*model.QrSpiritSignUDatum, error) {
	result, err := q.DO.Find()
	return result.([]*model.QrSpiritSignUDatum), err
}

func (q qrSpiritSignUDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSpiritSignUDatum, err error) {
	buf := make([]*model.QrSpiritSignUDatum, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qrSpiritSignUDatumDo) FindInBatches(result *[]*model.QrSpiritSignUDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qrSpiritSignUDatumDo) Attrs(attrs ...field.AssignExpr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qrSpiritSignUDatumDo) Assign(attrs ...field.AssignExpr) IQrSpiritSignUDatumDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qrSpiritSignUDatumDo) Joins(fields ...field.RelationField) IQrSpiritSignUDatumDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qrSpiritSignUDatumDo) Preload(fields ...field.RelationField) IQrSpiritSignUDatumDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qrSpiritSignUDatumDo) FirstOrInit() (*model.QrSpiritSignUDatum, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSpiritSignUDatum), nil
	}
}

func (q qrSpiritSignUDatumDo) FirstOrCreate() (*model.QrSpiritSignUDatum, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSpiritSignUDatum), nil
	}
}

func (q qrSpiritSignUDatumDo) FindByPage(offset int, limit int) (result []*model.QrSpiritSignUDatum, count int64, err error) {
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

func (q qrSpiritSignUDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qrSpiritSignUDatumDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qrSpiritSignUDatumDo) Delete(models ...*model.QrSpiritSignUDatum) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qrSpiritSignUDatumDo) withDO(do gen.Dao) *qrSpiritSignUDatumDo {
	q.DO = *do.(*gen.DO)
	return q
}