package db

import (
	"bytes"
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"strings"
	"vitess.io/vitess/go/vt/sqlparser"
)

// GetTableInfo query table details from information schema
func GetTableInfo(sqlInfo SQLInfo) (tables []Table, err error) {
	sqlArr := strings.Split(sqlInfo.Sql, ";")
	for i := range sqlArr {
		sql := sqlArr[i]
		stat, err := sqlparser.ParseStrictDDL(sql)
		if err != nil {
			elog.Error("sqlparser.Parse error", elog.FieldErr(err))
			continue
		}
		switch stat.(type) {
		case *sqlparser.CreateTable:
			table := parse(stat)
			tables = append(tables, table)
		default:
			elog.Error("sql not creat table")
		}
	}
	return
}

func parse(stat sqlparser.Statement) (table Table) {
	createTable, ok := stat.(*sqlparser.CreateTable)
	if !ok {
		elog.Errorf("stat.(*sqlparser.DDL) convert fail")
		return
	}
	table = Table{
		Name:    createTable.Table.Name.String(),
		Columns: make([]Column, 0, 10),
	}

	spec := createTable.GetTableSpec()

	for _, col := range spec.Columns {
		column := Column{
			Name:       col.Name.String(),
			Comment:    string(col.Type.Comment.Val),
			Type:       col.Type.Type,
			NotNull:    col.Type.NotNull,
			AutoInc:    col.Type.Autoincrement,
			Unique:     col.Type.KeyOpt == 4,
			PrimaryKey: col.Type.KeyOpt == 1,
		}
		table.Columns = append(table.Columns, column)
	}
	for _, option := range spec.Options {
		if strings.ToUpper(option.Name) == "COMMENT" {
			table.Comment = string(option.Value.Val)
			break
		}
	}
	return
}

func isAutoInc(extra string) bool {
	return strings.Contains(extra, "auto_increment")
}

func isUnique(columnKey string) bool {
	return isPrimaryKey(columnKey) || strings.Contains(columnKey, "UNI")
}

func isPrimaryKey(columnKey string) bool {
	return strings.Contains(columnKey, "PRI")
}

// generateMysqlTypes go struct entries for a map[string]interface{} structure
func generateModelFields(table *Table, depth int, option *GenerateOption) string {

	structure := ""

	for _, column := range table.Columns {

		// Get the corresponding go value type for this mysql type
		valueType := mysqlTypeToGoType(column.Type, !column.NotNull, option.WithGureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(column.Name))
		var annotations []string
		if option.WithGormAnnotation {
			annotations = append(annotations, generateGormAnnotation(column))
		}
		if option.WithJsonAnnotation {
			annotations = append(annotations, generateJSONAnnotation(column))
		}
		if option.WithDBAnnotation {
			annotations = append(annotations, generateDBAnnotation(column))
		}
		if option.WithXmlAnnotation {
			annotations = append(annotations, generateXMLAnnotation(column))
		}
		if option.WithXormAnnotation {
			annotations = append(annotations, generateXormAnnotation(column))
		}
		if option.WithFakerAnnotation {
			annotations = append(annotations, generateFakerAnnotation(column))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			structure += fmt.Sprintf("\n%s %s",
				fieldName,
				valueType)
		}
	}
	return structure
}

func generateGormAnnotation(col Column) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("column:%s", col.Name))
	if col.PrimaryKey {
		buf.WriteString(";primaryKey")
	}
	if col.Unique {
		buf.WriteString(";unique")
	}
	if col.NotNull {
		buf.WriteString(";not null")
	}
	if col.AutoInc {
		buf.WriteString(";autoIncrement")
	}
	return fmt.Sprintf(`gorm:"%s"`, buf.String())
}
func generateJSONAnnotation(col Column) string {
	return fmt.Sprintf(`json:"%s"`, col.Name)
}
func generateDBAnnotation(col Column) string {
	return fmt.Sprintf(`db:"%s"`, col.Name)
}
func generateXMLAnnotation(col Column) string {
	return fmt.Sprintf(`xml:"%s"`, col.Name)
}
func generateXormAnnotation(col Column) string {
	return fmt.Sprintf(`xorm:"%s"`, col.Name)
}
func generateFakerAnnotation(col Column) string {
	return fmt.Sprintf(`faker:"%s"`, col.Name)
}

// mysqlTypeToGoType converts the mysql types to go compatible sql.NullAble (https://golang.org/pkg/database/sql/) types
func mysqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext", "json":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
