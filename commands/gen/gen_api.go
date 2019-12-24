package gen

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	_ "github.com/lib/pq"
	"strings"
	"time"

	//_ "github.com/mattn/go-oci8"
	//_ "github.com/mattn/go-sqlite3"
)

const (
	DEFAULT_GEN_API_PATH = "api"
)
// doGenModel implements the "gen model" command.

func generateApiContentFile(table,variable,moduleName string)  {
	UpperTableName := gstr.CamelCase(variable)
	folderPath := DEFULAT_GEN_BASE_PATH + gfile.Separator + DEFAULT_GEN_API_PATH
	nowTime := time.Now().Format("2006-01-02 15:16:17")
	serviceContent := gstr.ReplaceByMap(templateConstApiContent, g.MapStrStr{
		"{fullTableName}": variable,
		"{nowTime}":nowTime,
		"{tableName}":variable,
		"{UpperTableName}":UpperTableName,
		"{moduleName}":moduleName,
	})
	fileName := gstr.SnakeCase(variable)
	path := gfile.Join(folderPath, fileName, fileName+".go")
	if err := gfile.PutContents(path, strings.TrimSpace(serviceContent)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		mlog.Print("generated:", path)
	}
}
