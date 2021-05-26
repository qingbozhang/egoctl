package db

import (
	"io/ioutil"
	"strings"
	"testing"
	"vitess.io/vitess/go/vt/sqlparser"
)

func TestParse(t *testing.T) {
	b, err := ioutil.ReadFile("demo.sql")
	if err != nil {
		t.Errorf("read file error : %v \n", err)
		return
	}
	sqlStr := string(b)
	t.Logf("%s", sqlStr)
	sqlArr := strings.Split(sqlStr, ";")
	for i := range sqlArr {
		sql := sqlArr[i]
		stat, err := sqlparser.ParseStrictDDL(sql)
		if err != nil {
			t.Logf("sqlparser.Parse error ")
			continue
		}
		switch stat.(type) {
		case *sqlparser.CreateTable:
			parse(stat)
		default:
			t.Logf("ffff")
		}

	}

}
