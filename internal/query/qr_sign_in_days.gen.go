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

func newQrSignInDay(db *gorm.DB, opts ...gen.DOOption) qrSignInDay {
	_qrSignInDay := qrSignInDay{}

	_qrSignInDay.qrSignInDayDo.UseDB(db, opts...)
	_qrSignInDay.qrSignInDayDo.UseModel(&model.QrSignInDay{})

	tableName := _qrSignInDay.qrSignInDayDo.TableName()
	_qrSignInDay.ALL = field.NewAsterisk(tableName)
	_qrSignInDay.ID = field.NewInt32(tableName, "id")
	_qrSignInDay.Qq = field.NewString(tableName, "qq")
	_qrSignInDay.MonthDays = field.NewInt32(tableName, "month_days")
	_qrSignInDay.TotalDays = field.NewInt32(tableName, "total_days")
	_qrSignInDay.UpdateDate = field.NewField(tableName, "update_date")
	_qrSignInDay.CreateDate = field.NewField(tableName, "create_date")
	_qrSignInDay.CreateBy = field.NewString(tableName, "create_by")
	_qrSignInDay.CreateTime = field.NewField(tableName, "create_time")
	_qrSignInDay.UpdateBy = field.NewString(tableName, "update_by")
	_qrSignInDay.UpdateTime = field.NewField(tableName, "update_time")
	_qrSignInDay.DelFlag = field.NewField(tableName, "del_flag")

	_qrSignInDay.fillFieldMap()

	return _qrSignInDay
}

// qrSignInDay 签到天数关联表
type qrSignInDay struct {
	qrSignInDayDo

	ALL        field.Asterisk
	ID         field.Int32  // 主键
	Qq         field.String // 用户QQ
	MonthDays  field.Int32  // 月连续签到天数
	TotalDays  field.Int32  // 总天数
	UpdateDate field.Field  // 更新日期
	CreateDate field.Field  // 创建日期
	CreateBy   field.String // 创建人
	CreateTime field.Field  // 创建时间
	UpdateBy   field.String // 更新人
	UpdateTime field.Field  // 更新时间
	DelFlag    field.Field  // 删除标识;0否1是

	fieldMap map[string]field.Expr
}

func (q qrSignInDay) Table(newTableName string) *qrSignInDay {
	q.qrSignInDayDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qrSignInDay) As(alias string) *qrSignInDay {
	q.qrSignInDayDo.DO = *(q.qrSignInDayDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qrSignInDay) updateTableName(table string) *qrSignInDay {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt32(table, "id")
	q.Qq = field.NewString(table, "qq")
	q.MonthDays = field.NewInt32(table, "month_days")
	q.TotalDays = field.NewInt32(table, "total_days")
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

func (q *qrSignInDay) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qrSignInDay) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 11)
	q.fieldMap["id"] = q.ID
	q.fieldMap["qq"] = q.Qq
	q.fieldMap["month_days"] = q.MonthDays
	q.fieldMap["total_days"] = q.TotalDays
	q.fieldMap["update_date"] = q.UpdateDate
	q.fieldMap["create_date"] = q.CreateDate
	q.fieldMap["create_by"] = q.CreateBy
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_by"] = q.UpdateBy
	q.fieldMap["update_time"] = q.UpdateTime
	q.fieldMap["del_flag"] = q.DelFlag
}

func (q qrSignInDay) clone(db *gorm.DB) qrSignInDay {
	q.qrSignInDayDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qrSignInDay) replaceDB(db *gorm.DB) qrSignInDay {
	q.qrSignInDayDo.ReplaceDB(db)
	return q
}

type qrSignInDayDo struct{ gen.DO }

type IQrSignInDayDo interface {
	gen.SubQuery
	Debug() IQrSignInDayDo
	WithContext(ctx context.Context) IQrSignInDayDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQrSignInDayDo
	WriteDB() IQrSignInDayDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQrSignInDayDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQrSignInDayDo
	Not(conds ...gen.Condition) IQrSignInDayDo
	Or(conds ...gen.Condition) IQrSignInDayDo
	Select(conds ...field.Expr) IQrSignInDayDo
	Where(conds ...gen.Condition) IQrSignInDayDo
	Order(conds ...field.Expr) IQrSignInDayDo
	Distinct(cols ...field.Expr) IQrSignInDayDo
	Omit(cols ...field.Expr) IQrSignInDayDo
	Join(table schema.Tabler, on ...field.Expr) IQrSignInDayDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQrSignInDayDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQrSignInDayDo
	Group(cols ...field.Expr) IQrSignInDayDo
	Having(conds ...gen.Condition) IQrSignInDayDo
	Limit(limit int) IQrSignInDayDo
	Offset(offset int) IQrSignInDayDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSignInDayDo
	Unscoped() IQrSignInDayDo
	Create(values ...*model.QrSignInDay) error
	CreateInBatches(values []*model.QrSignInDay, batchSize int) error
	Save(values ...*model.QrSignInDay) error
	First() (*model.QrSignInDay, error)
	Take() (*model.QrSignInDay, error)
	Last() (*model.QrSignInDay, error)
	Find() ([]*model.QrSignInDay, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSignInDay, err error)
	FindInBatches(result *[]*model.QrSignInDay, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QrSignInDay) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQrSignInDayDo
	Assign(attrs ...field.AssignExpr) IQrSignInDayDo
	Joins(fields ...field.RelationField) IQrSignInDayDo
	Preload(fields ...field.RelationField) IQrSignInDayDo
	FirstOrInit() (*model.QrSignInDay, error)
	FirstOrCreate() (*model.QrSignInDay, error)
	FindByPage(offset int, limit int) (result []*model.QrSignInDay, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQrSignInDayDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q qrSignInDayDo) Debug() IQrSignInDayDo {
	return q.withDO(q.DO.Debug())
}

func (q qrSignInDayDo) WithContext(ctx context.Context) IQrSignInDayDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qrSignInDayDo) ReadDB() IQrSignInDayDo {
	return q.Clauses(dbresolver.Read)
}

func (q qrSignInDayDo) WriteDB() IQrSignInDayDo {
	return q.Clauses(dbresolver.Write)
}

func (q qrSignInDayDo) Session(config *gorm.Session) IQrSignInDayDo {
	return q.withDO(q.DO.Session(config))
}

func (q qrSignInDayDo) Clauses(conds ...clause.Expression) IQrSignInDayDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qrSignInDayDo) Returning(value interface{}, columns ...string) IQrSignInDayDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qrSignInDayDo) Not(conds ...gen.Condition) IQrSignInDayDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qrSignInDayDo) Or(conds ...gen.Condition) IQrSignInDayDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qrSignInDayDo) Select(conds ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qrSignInDayDo) Where(conds ...gen.Condition) IQrSignInDayDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qrSignInDayDo) Order(conds ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qrSignInDayDo) Distinct(cols ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qrSignInDayDo) Omit(cols ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qrSignInDayDo) Join(table schema.Tabler, on ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qrSignInDayDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qrSignInDayDo) RightJoin(table schema.Tabler, on ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qrSignInDayDo) Group(cols ...field.Expr) IQrSignInDayDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qrSignInDayDo) Having(conds ...gen.Condition) IQrSignInDayDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qrSignInDayDo) Limit(limit int) IQrSignInDayDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qrSignInDayDo) Offset(offset int) IQrSignInDayDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qrSignInDayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQrSignInDayDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qrSignInDayDo) Unscoped() IQrSignInDayDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qrSignInDayDo) Create(values ...*model.QrSignInDay) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qrSignInDayDo) CreateInBatches(values []*model.QrSignInDay, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qrSignInDayDo) Save(values ...*model.QrSignInDay) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qrSignInDayDo) First() (*model.QrSignInDay, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDay), nil
	}
}

func (q qrSignInDayDo) Take() (*model.QrSignInDay, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDay), nil
	}
}

func (q qrSignInDayDo) Last() (*model.QrSignInDay, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDay), nil
	}
}

func (q qrSignInDayDo) Find() ([]*model.QrSignInDay, error) {
	result, err := q.DO.Find()
	return result.([]*model.QrSignInDay), err
}

func (q qrSignInDayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QrSignInDay, err error) {
	buf := make([]*model.QrSignInDay, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qrSignInDayDo) FindInBatches(result *[]*model.QrSignInDay, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qrSignInDayDo) Attrs(attrs ...field.AssignExpr) IQrSignInDayDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qrSignInDayDo) Assign(attrs ...field.AssignExpr) IQrSignInDayDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qrSignInDayDo) Joins(fields ...field.RelationField) IQrSignInDayDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qrSignInDayDo) Preload(fields ...field.RelationField) IQrSignInDayDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qrSignInDayDo) FirstOrInit() (*model.QrSignInDay, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDay), nil
	}
}

func (q qrSignInDayDo) FirstOrCreate() (*model.QrSignInDay, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QrSignInDay), nil
	}
}

func (q qrSignInDayDo) FindByPage(offset int, limit int) (result []*model.QrSignInDay, count int64, err error) {
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

func (q qrSignInDayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qrSignInDayDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qrSignInDayDo) Delete(models ...*model.QrSignInDay) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qrSignInDayDo) withDO(do gen.Dao) *qrSignInDayDo {
	q.DO = *do.(*gen.DO)
	return q
}
