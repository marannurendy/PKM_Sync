package conf

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB
var DbBRNET *sql.DB
var DbPKM *sql.DB

func init() {
	Db, _ = sql.Open("mssql", "sqlserver://"+beego.AppConfig.String("mssqluser")+":"+beego.AppConfig.String("mssqlpass")+"@"+beego.AppConfig.String("mssqlurls")+"?database="+beego.AppConfig.String("mssqldb")+"&connection+timeout=0")
	DbBRNET, _ = sql.Open("mssql", "sqlserver://"+beego.AppConfig.String("BRNETmssqluser")+":"+beego.AppConfig.String("BRNETmssqlpass")+"@"+beego.AppConfig.String("BRNETmssqlurls")+"?database="+beego.AppConfig.String("BRNETmssqldb")+"&connection+timeout=0")
	DbPKM, _ = sql.Open("mssql", "sqlserver://"+beego.AppConfig.String("PKMmssqluser")+":"+beego.AppConfig.String("PKMmssqlpass")+"@"+beego.AppConfig.String("PKMmssqlurls")+"?database="+beego.AppConfig.String("PKMmssqldb")+"&connection+timeout=0")
}
