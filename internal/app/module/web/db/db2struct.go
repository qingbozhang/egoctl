package db

import (
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/egoctl/internal/app/module/web/core"
)

type SQLInfo struct {
	Sql string `json:"sql"`
}

func DbToGoStruct(ctx *core.Context) {
	var sqlInfo SQLInfo
	err := ctx.ShouldBindJSON(&sqlInfo)
	if err != nil {
		ctx.JSONE(1, "获取参数失败: err"+err.Error(), err)
		return
	}

	tables, err := GetTableInfo(sqlInfo)
	structs := make([]string, 0, len(tables))

	for _, table := range tables {
		structBytes, err := Generate(&table, "", "", &GenerateOption{
			WithJsonAnnotation:  true,
			WithDBAnnotation:    true,
			WithGormAnnotation:  true,
			WithXmlAnnotation:   false,
			WithXormAnnotation:  false,
			WithFakerAnnotation: false,
			WithGureguTypes:     false,
			StructSorted:        false,
		})

		if err != nil {
			elog.Error("to struct error", elog.FieldErr(err))
		}
		structs = append(structs, string(structBytes))
	}
}
