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

func newQrSignInDatum(db *gorm.DB, opts ...gen.DOOption) qrSignInDatum {
	_qrSignInDatum := qrSignInDatum{}

	_qrSignInDatum.qrSignInDatumDo.UseDB(db, opts...)
	_qrSignInDatum.qrSignInDatumDo.UseModel(&model.QrSignInDatum{})

	tableName := _qrSignInDatum.qrSignInDatumDo.TableName()
	_qrSignInDatum.ALL = field.NewAsterisk(tableName)
	_qrSignInDatum.ID = field.NewInt32(tableName, "id")
	_qrSignInDatum.Qq = field.NewString(tableName, "qq")
	_qrSignInDatum.QqNumber = field.NewInt64(tableName, "qq_number")
	_qrSignInDatum.Points = field.NewInt32(tableName, "points")
	_qrSignInDatum.DayID = field.NewInt32(tableName, "day_id")
	_qrSignInDatum.SignInDate = field.NewField(tableName, "sign_in_date")
	_qrSignInDatum.UpdateDate = field.NewField(tableName, "update_date")
	_qrSignInDatum.CreateDate = field.NewField(tableName, "create_date")
	_qrSignInDatum.CreateBy = field.NewString(tableName, "create_by")
	_qrSignInDatum.CreateTime = field.NewField(tableName, "create_time")
	_qrSignInDatum.UpdateBy = field.NewString(tableName, "update_by")
	_qrSignInDatum.UpdateTime = field.NewField(tableName, "update_time")
	_qrSignInDatum.DelFlag = field.NewField(tableName, "del_flag")

	_qrSignInDatum.fillFieldMap()

	return _qrSignInDatum
}

// qrSignInDatum QQ签到表
type qrSignInDatum struct {
	qrSignInDatumDo

	ALL        field.Asterisk
	ID         field.Int32  // 主键
	Qq         field.String // 用户QQ
	QqNumber   field.Int64  // QQ号码
	Points     field.Int32  // 积分
	DayID      field.Int32  // 连续签到天数
	SignInDate field.Field  // 签到日期
	UpdateDate field.Field  // 更新日期
	CreateDate field.Field  // 创建日期
	CreateBy   field.String // 创建人
	CreateTime field.Field  // 创建时间
	UpdateBy   field.String // 更新人
	UpdateTime field.Field  // 更新时间
	DelFlag    field.Field  // 删除标识;0否1是

	fieldMap map[string]field.Expr
}

func (q qrSignInDatum) Table(newTableName string) *qrSignInDatum {
	q.qrSignInDatumDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qrSignInDatum) As(alias string) *qrSignInDatum {
	q.qrSignInDatumDo.DO = *(q.qrSignInDatumDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qrSignInDatum) updateTableName(table string) *qrSignInDatum {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt32(table, "id")
	q.Qq = field.NewString(table, "qq")
	q.QqNumber = field.NewInt64(table, "qq_number")
	q.Points = field.NewInt32(table, "points")
	q.DayID = field.NewInt32(table, "day_id")
	q.SignInDate = field.NewField(table, "sign_in_date")
	q.UpdateDate = field.NewField(table, "update_date")
	q.CreateDate = field.NewField(table, "create_date")
	q.CreateBy = field.NewString(table, "create_by")
	q.CreateTime = field.NewField(table, "create_time")
	q.UpdateBy = field.NewString(table, "update_by")
	q.UpdateTime = field.NewField(table, "update_time")
	q.DelFlag = field.NewField(table, "del_flag")

	q.fillFieldMap()

	return q
}

func (q *qrSignInDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qrSignInDatum) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 13)
	q.fieldMap["id"] = q.ID
	q.fieldMap["qq"] = q.Qq
	q.fieldMap["qq_number"] = q.QqNumber
	q.fieldMap["points"] = q.Points
	q.fieldMap["day_id"] = q.DayID
	q.fieldMap["sign_in_date"] = q.SignInDate
	q.fieldMap["update_date"] = q.UpdateDate
	q.fieldMap["create_date"] = q.CreateDate
	q.fieldMap["create_by"] = q.CreateBy
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_by"] = q.UpdateBy
	q.fieldMap["update_time"] = q.UpdateTime
	q.fieldMap["del_flag"] = q.DelFlag
}

func (q qrSignInDatum) clone(db *gorm.DB) qrSignInDatum {
	q.qrSignInDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qrSignInDatum) replaceDB(db *gorm.DB) qrSignInDatum {
	q.qrSignInDatumDo.ReplaceDB(db)
	return q
}

type qrSignInDatumDo struct{ gen.DO }

type IQrSignInDatumDo interface {
	gen.SubQuery
	Debug() IQrSignInDatumDo
	WithContext(ctx context.Context) IQrSignInDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQrSignInDatumDo
	WriteDB() IQrSignInDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQrSignInDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQrSignInDatumDo
	Not(conds ...gen.Condition) IQrSignInDatumDo
	Or(conds ...gen.Condition) IQrSignInDatumDo
	Select(conds ...field.Expr) IQrSignInDatumDo
	Where(conds ...gen.Condition) IQrSignInDatumDo
	Order(conds ...field.Expr) IQrSignInDatumDo
	Distinct(cols ...field.Expr) IQrSignInDatumDo
	Omit(cols ...field.Expr) IQrSignInDatumDo
	Join(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo
	Group(cols ...field.Expr) IQrSignInDatumDo
	Having(conds ...gen.Condition) IQrSignInDatumDo
	Limit(limit int) IQrSignInDatumDo
	Offset(offset int) IQrSignInDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSignInDatumDo
	Unscoped() IQrSignInDatumDo
	Create(values ...*model.QrSignInDatum) error
	CreateInBatches(values []*model.QrSignInDatum, batchSize int) error
	Save(values ...*model.QrSignInDatum) error
	First() (*model.QrSignInDatum, error)
	Take() (*model.QrSignInDatum, error)
	Last() (*model.QrSignInDatum, error)
	Find() ([]*model.QrSignInDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSignInDatum, err error)
	FindInBatches(result *[]*model.QrSignInDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QrSignInDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQrSignInDatumDo
	Assign(attrs ...field.AssignExpr) IQrSignInDatumDo
	Joins(fields ...field.RelationField) IQrSignInDatumDo
	Preload(fields ...field.RelationField) IQrSignInDatumDo
	FirstOrInit() (*model.QrSignInDatum, error)
	FirstOrCreate() (*model.QrSignInDatum, error)
	FindByPage(offset int, limit int) (result []*model.QrSignInDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQrSignInDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qrSignInDatumDo) Debug() IQrSignInDatumDo {
	return q.withDO(q.DO.Debug())
}

func (q qrSignInDatumDo) WithContext(ctx context.Context) IQrSignInDatumDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qrSignInDatumDo) ReadDB() IQrSignInDatumDo {
	return q.Clauses(dbresolver.Read)
}

func (q qrSignInDatumDo) WriteDB() IQrSignInDatumDo {
	return q.Clauses(dbresolver.Write)
}

func (q qrSignInDatumDo) Session(config *gorm.Session) IQrSignInDatumDo {
	return q.withDO(q.DO.Session(config))
}

func (q qrSignInDatumDo) Clauses(conds ...clause.Expression) IQrSignInDatumDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qrSignInDatumDo) Returning(value interface{}, columns ...string) IQrSignInDatumDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qrSignInDatumDo) Not(conds ...gen.Condition) IQrSignInDatumDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qrSignInDatumDo) Or(conds ...gen.Condition) IQrSignInDatumDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qrSignInDatumDo) Select(conds ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qrSignInDatumDo) Where(conds ...gen.Condition) IQrSignInDatumDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qrSignInDatumDo) Order(conds ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qrSignInDatumDo) Distinct(cols ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qrSignInDatumDo) Omit(cols ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qrSignInDatumDo) Join(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qrSignInDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qrSignInDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qrSignInDatumDo) Group(cols ...field.Expr) IQrSignInDatumDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qrSignInDatumDo) Having(conds ...gen.Condition) IQrSignInDatumDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qrSignInDatumDo) Limit(limit int) IQrSignInDatumDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qrSignInDatumDo) Offset(offset int) IQrSignInDatumDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qrSignInDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSignInDatumDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qrSignInDatumDo) Unscoped() IQrSignInDatumDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qrSignInDatumDo) Create(values ...*model.QrSignInDatum) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qrSignInDatumDo) CreateInBatches(values []*model.QrSignInDatum, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qrSignInDatumDo) Save(values ...*model.QrSignInDatum) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qrSignInDatumDo) First() (*model.QrSignInDatum, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDatum), nil
	}
}

func (q qrSignInDatumDo) Take() (*model.QrSignInDatum, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDatum), nil
	}
}

func (q qrSignInDatumDo) Last() (*model.QrSignInDatum, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDatum), nil
	}
}

func (q qrSignInDatumDo) Find() ([]*model.QrSignInDatum, error) {
	result, err := q.DO.Find()
	return result.([]*model.QrSignInDatum), err
}

func (q qrSignInDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSignInDatum, err error) {
	buf := make([]*model.QrSignInDatum, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qrSignInDatumDo) FindInBatches(result *[]*model.QrSignInDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qrSignInDatumDo) Attrs(attrs ...field.AssignExpr) IQrSignInDatumDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qrSignInDatumDo) Assign(attrs ...field.AssignExpr) IQrSignInDatumDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qrSignInDatumDo) Joins(fields ...field.RelationField) IQrSignInDatumDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qrSignInDatumDo) Preload(fields ...field.RelationField) IQrSignInDatumDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qrSignInDatumDo) FirstOrInit() (*model.QrSignInDatum, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDatum), nil
	}
}

func (q qrSignInDatumDo) FirstOrCreate() (*model.QrSignInDatum, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDatum), nil
	}
}

func (q qrSignInDatumDo) FindByPage(offset int, limit int) (result []*model.QrSignInDatum, count int64, err error) {
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

func (q qrSignInDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qrSignInDatumDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qrSignInDatumDo) Delete(models ...*model.QrSignInDatum) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qrSignInDatumDo) withDO(do gen.Dao) *qrSignInDatumDo {
	q.DO = *do.(*gen.DO)
	return q
}
