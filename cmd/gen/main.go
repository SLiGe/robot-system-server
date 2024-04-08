package main

import (
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/query",
		FieldNullable: true,
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/niurensec?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	FieldJSONTagWithNS := func(columnName string) string {
		return strutil.CamelCase(strings.ToLower(columnName))
	}
	g.WithOpts(gen.FieldModify(func(field gen.Field) gen.Field {
		field.Name = strings.ToLower(field.Name)
		field.ColumnName = strings.ToLower(field.ColumnName)
		if field.Name == "create_date" {
			field.Type = "*db.LocalDateTime"
			field.GORMTag.Append("autoCreateTime")
		}
		if field.Name == "update_date" {
			field.Type = "*db.LocalDateTime"
			field.GORMTag.Append("autoUpdateTime")
		}
		if field.Name == "del_flag" {
			field.Type = "soft_delete.DeletedAt"
			field.GORMTag.Append("softDelete:flag")
		}

		return field
	}))
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	// json tag
	g.WithJSONTagNameStrategy(FieldJSONTagWithNS)
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		g.GenerateModel("qr_chinese_bqb"),
		g.GenerateModel("qr_fortune"),
		g.GenerateModel("qr_fortune_data"),
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
