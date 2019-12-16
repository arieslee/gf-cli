
package service

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"math"
)

type BaseService struct {

}
const (
	LeftJoin string = "left"
	RightJoin string = "right"
	InnerJoin string = "inner"
)
/**
input示例
listInput := map[string]interface{}{
	"where":"p.status=?",
	"params":[]interface{}{1},
	"orderBy":"p.id DESC",
	"select":"p.id,p.created_at,p.title",
	"join":map[string]interface{}{
		"left":map[string]interface{}{
			"table":model.TableBingoCategory+" AS c",
			"on":"c.id=p.cid",
		},
	},
}
 */
func (bs *BaseService) List(r *ghttp.Request, tableName string, input map[string]interface{})(map[string]interface{},error) {
	page, ok := input["page"]
	if !ok {
		page = r.GetInt("page")
		if page == 0{
			page = 1
		}
	}
	pageSize := r.GetInt("pageSize")
	defaultPageSize := g.Config().GetInt("app.pageSize")
	if pageSize == 0{
		pageSize = defaultPageSize
	}
	db := g.DB()
	model := db.Table(tableName)

	if _,ok:=input["join"];ok{
		joins := gconv.Map(input["join"])
		for key,value := range joins{
			joinCondition := gconv.Map(value)
			if key == LeftJoin{
				model.LeftJoin(joinCondition["table"].(string), joinCondition["on"].(string))
			}else if key == RightJoin{
				model.RightJoin(joinCondition["table"].(string), joinCondition["on"].(string))
			}else if key == InnerJoin{
				model.InnerJoin(joinCondition["table"].(string), joinCondition["on"].(string))
			}
		}
	}
	where := ""
	if _,ok := input["where"];ok{
		where += gconv.String(input["where"])
	}
	var bindParams []interface{}
	if _,ok :=input["params"];ok{
		bindParams = gconv.Interfaces(input["params"])
	}
	model = model.Where(where, bindParams)
	countModel := model.Clone()

	if _, ok := input["select"];ok{
		model = model.Fields(input["select"].(string))
	}
	if _,ok:=input["orderBy"];ok{
		model = model.OrderBy(gconv.String(input["orderBy"]))
	}
	result := map[string]interface{}{
		"currentPage":page,
		"pageSize":pageSize,
		"totalPage":0,
		"totalCount":0,
		"list": nil,
	}

	totalCount,countErr := countModel.Count()
	if countErr!=nil{
		return result, countErr
	}
	result["totalCount"] = totalCount
	divPage := float64(totalCount/pageSize)
	totalPage := math.Ceil(divPage)
	result["totalPage"] = gconv.Int(totalPage)
	data,err:=model.ForPage(gconv.Int(page),pageSize).Select()
	if err != nil{
		if err == sql.ErrNoRows{
			result["totalCount"] = 0
			result["totalPage"] = 0
		}
		return result, err
	}
	result["list"] = data
	return result, nil
}
/**
使用方法
ar := bs.FindBy(tableName, params)
ar.Struct(&post)
 */
func (bs *BaseService) FindBy(table string, params map[string]interface{}) *gdb.Model {
	var (
		where string
		bindParams []interface{}
	)
	if _,ok:=params["where"];ok{
		where += gconv.String(params["where"])
	}
	if _,ok:=params["params"];ok{
		bindParams = gconv.Interfaces(params["params"])
	}
	db:=g.DB()
	return db.Table(table).Where(where,bindParams).Limit(1)
}
