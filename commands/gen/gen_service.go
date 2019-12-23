package gen

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	_ "github.com/lib/pq"
	"strings"
	//_ "github.com/mattn/go-oci8"
	//_ "github.com/mattn/go-sqlite3"
)

const (
	DEFAULT_GEN_SERVICE_PATH = "./app/service"
	DEFULAT_GEN_BASE_PATH = "./app"
)

func genBaseService()  {
	file := DEFAULT_GEN_SERVICE_PATH + gfile.Separator + "base_service.go"
	isExists := gfile.Exists(file)
	if !isExists{
		if err := gfile.PutContents(file, templateBaseServiceContent); err != nil {
			mlog.Fatalf("write base service file to '%s' failed: %v", file, err)
		}
	}
}
// doGenModel implements the "gen model" command.
func doGenService(parser *gcmd.Parser) {
	var err error
	basePath := parser.GetOpt("path")
	if basePath == ""{
		basePath = DEFULAT_GEN_BASE_PATH
	}
	genPath := basePath + "/service"
	modelGenPath := basePath + "/model"
	//genPath := parser.GetArg(3, DEFAULT_GEN_SERVICE_PATH)
	if !gfile.IsEmpty(genPath) {
		s := gcmd.Scanf("path '%s' is not empty, files might be overwrote, continue? [y/n]: ", genPath)
		if strings.EqualFold(s, "n") {
			return
		}
	}
	//生成base service
	genBaseService()
	tableOpt := parser.GetOpt("table")
	linkInfo := parser.GetOpt("link")
	configFile := parser.GetOpt("config")
	configGroup := parser.GetOpt("group", gdb.DEFAULT_GROUP_NAME)
	//要忽略的列前缀
	prefixArray := gstr.SplitAndTrim(parser.GetOpt("prefix"), ",")
	moduleName := parser.GetOpt("module")

	if moduleName == ""{
		mlog.Fatalf("Please input module name, ex.-m=gf-app")
		return
	}
	if linkInfo != "" {
		path := gfile.TempDir() + gfile.Separator + "config.toml"
		if err := gfile.PutContents(path, fmt.Sprintf("[database]\n\tlink=\"%s\"", linkInfo)); err != nil {
			mlog.Fatalf("write configuration file to '%s' failed: %v", path, err)
		}
		defer gfile.Remove(path)
		if err := g.Cfg().SetPath(gfile.TempDir()); err != nil {
			mlog.Fatalf("set configuration path '%s' failed: %v", gfile.TempDir(), err)
		}
	}

	if configFile != "" {
		path, err := gfile.Search(configFile)
		if err != nil {
			mlog.Fatalf("search configuration file '%s' failed: %v", configFile, err)
		}
		if err := g.Cfg().SetPath(path); err != nil {
			mlog.Fatalf("set configuration path '%s' failed: %v", path, err)
		}
		if err := g.Cfg().SetFileName(gfile.Basename(path)); err != nil {
			mlog.Fatalf("set configuration file name '%s' failed: %v", gfile.Basename(path), err)
		}
	}

	db := g.DB(configGroup)
	if db == nil {
		mlog.Fatal("database initialization failed")
	}

	if err := gfile.Mkdir(genPath); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", genPath, err)
	}

	tables := ([]string)(nil)
	if tableOpt != "" {
		tables = gstr.SplitAndTrim(tableOpt, ",")
	} else {
		tables, err = db.Tables()
		if err != nil {
			mlog.Fatalf("fetching tables failed: \n %v", err)
		}
	}
	for _, table := range tables {
		variable := table
		if len(variable) == 0{
			continue
		}
		//去掉表前缀
		for _, v := range prefixArray {
			cutLen := len(v)
			variable = variable[cutLen:len(variable)]
		}
		//生成model
		generateModelContentFile(db, table, variable, modelGenPath, configGroup)
		//生成service
		generateServiceContentFile(table, variable, moduleName)
	}
	mlog.Print("done!")
}
func generateServiceContentFile(table,variable,moduleName string)  {
	fullTableName := gstr.CamelCase(table)
	UpperTableName := gstr.CamelCase(variable)
	folderPath := DEFAULT_GEN_SERVICE_PATH

	serviceContent := gstr.ReplaceByMap(templateConstServiceContent, g.MapStrStr{
		"{fullTableName}": fullTableName,
		"{UpperTableName}":UpperTableName,
		"{moduleName}":moduleName,
	})
	fileName := gstr.SnakeCase(variable)
	path := gfile.Join(folderPath, fileName+"_service.go")
	if err := gfile.PutContents(path, strings.TrimSpace(serviceContent)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		mlog.Print("generated:", path)
	}
}
