package db

import (
	"strings"
	"testing"
	"vitess.io/vitess/go/vt/sqlparser"
)

func TestParse(t *testing.T) {
	sql := "CREATE TABLE `table_1` (\n  `id` int NOT NULL AUTO_INCREMENT,\n  `address` varchar(191) NOT NULL DEFAULT '',\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='table 111'"
	stat, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		t.Errorf("sqlparser.Parse error : %v \n", err)
		return
	}
	t.Logf("sqlparser.Parse success \n")
	t.Logf("%#v \n", stat)

	createTable, ok := stat.(*sqlparser.CreateTable)
	if !ok {
		t.Errorf("stat.(*sqlparser.DDL) convert fail")
		return
	}
	t.Logf("%#v \n", createTable)
	spec := createTable.GetTableSpec()
	for _, column := range spec.Columns {
		t.Logf("column %+v", column)
	}
	for _, option := range spec.Options {
		if strings.ToUpper(option.Name) == "COMMENT" {
			t.Logf("COMMENT %s", option.Value.Val)
		} else {
			t.Logf("option %+v", option)
		}
	}
	t.Logf("table %+v", createTable.GetToTables())
}
