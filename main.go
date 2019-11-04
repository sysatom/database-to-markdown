package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

type Column struct {
	TableCatalog           string         `db:"TABLE_CATALOG"`
	TableSchema            string         `db:"TABLE_SCHEMA"`
	TableName              string         `db:"TABLE_NAME"`
	ColumnName             string         `db:"COLUMN_NAME"`
	OrdinalPosition        int            `db:"ORDINAL_POSITION"`
	ColumnDefault          sql.NullString `db:"COLUMN_DEFAULT"`
	IsNullable             string         `db:"IS_NULLABLE"`
	DataType               string         `db:"DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       sql.NullInt64  `db:"NUMERIC_PRECISION"`
	NumericScale           sql.NullInt64  `db:"NUMERIC_SCALE"`
	DatetimePrecision      sql.NullInt64  `db:"DATETIME_PRECISION"`
	CharacterSetName       sql.NullString `db:"CHARACTER_SET_NAME"`
	CollationName          sql.NullString `db:"COLLATION_NAME"`
	ColumnType             string         `db:"COLUMN_TYPE"`
	ColumnKey              string         `db:"COLUMN_KEY"`
	Extra                  string         `db:"EXTRA"`
	Privileges             string         `db:"PRIVILEGES"`
	ColumnComment          string         `db:"COLUMN_COMMENT"`
	GenerationExpression   string         `db:"GENERATION_EXPRESSION"`
}

type Table struct {
	TableCatalog   string         `db:"TABLE_CATALOG"`
	TableSchema    string         `db:"TABLE_SCHEMA"`
	TableName      string         `db:"TABLE_NAME"`
	TableType      string         `db:"TABLE_TYPE"`
	Engine         string         `db:"ENGINE"`
	Version        int            `db:"VERSION"`
	RowFormat      string         `db:"ROW_FORMAT"`
	TableRows      int            `db:"TABLE_ROWS"`
	AvgRowLength   int            `db:"AVG_ROW_LENGTH"`
	DataLength     int            `db:"DATA_LENGTH"`
	MaxDataLength  int            `db:"MAX_DATA_LENGTH"`
	IndexLength    int            `db:"INDEX_LENGTH"`
	DataFree       int            `db:"DATA_FREE"`
	AutoIncrement  sql.NullInt64  `db:"AUTO_INCREMENT"`
	CreateTime     string         `db:"CREATE_TIME"`
	UpdateTime     sql.NullString `db:"UPDATE_TIME"`
	CheckTime      sql.NullString `db:"CHECK_TIME"`
	TableCollation string         `db:"TABLE_COLLATION"`
	Checksum       sql.NullInt64  `db:"CHECKSUM"`
	CreateOptions  string         `db:"CREATE_OPTIONS"`
	TableComment   string         `db:"TABLE_COMMENT"`
}

func NullString2String(s sql.NullString) string {
	if s.Valid {
		return s.String
	} else {
		return "NULL"
	}
}

func main() {

	user := os.Args[1]
	password := os.Args[2]
	host := os.Args[3]
	database := os.Args[4]

	mysqlDSN := user + ":" + password + "@tcp(" + host + ")/" + database + "?parseTime=true&loc=Asia%2FShanghai"
	db, err := sqlx.Open("mysql", mysqlDSN)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	var tables []Table
	err = db.Select(&tables, "SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '"+database+"'")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	markdown := ""
	for _, table := range tables {

		var columns []Column
		err = db.Select(&columns, "SELECT * FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = '"+database+"' AND TABLE_NAME = '"+table.TableName+"'")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		comment := ""
		if utf8.RuneCountInString(table.TableComment) > 0 {
			comment += " (" + table.TableComment + ") "
		}
		markdown += "## " + table.TableName + comment + "\n\n"
		markdown += "| COLUMN_NAME |    COLUMN_TYPE   | COLUMN_DEFAULT | IS_NULLABLE | COLUMN_KEY |     EXTRA      | COLUMN_COMMENT |\n"
		markdown += "|-------------|------------------|----------------|-------------|------------|----------------|----------------|\n"

		for _, column := range columns {
			markdown += "| " + column.ColumnName + " | " + column.ColumnType + " | " + NullString2String(column.ColumnDefault) + " | " + column.IsNullable + " | " + column.ColumnKey + " | " + column.Extra + " | " + column.ColumnComment + " |\n"
		}

		markdown += "\n\n"
	}

	err = ioutil.WriteFile("./schema.md", []byte(markdown), 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("See schema.md")
}