package conf

import (
	// "fmt"

	"github.com/beego/beego/v2/core/logs"
)

func ErrorLog(message string) {
	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/www/html/PKM_Sync/errorLog/PKMapprove.log", "perm":"777"}`) //prod
	// logs.SetLogger(logs.AdapterFile, `{"filename":"successLog/PKMmobile.log", "perm":"777"}`) //development
	logs.Info(message)
}
