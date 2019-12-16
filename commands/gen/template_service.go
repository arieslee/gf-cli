/*
@Time : 2019/12/16 8:57 上午
@Author : sunmoon
@File : template_service
@Software: GoLand
*/
package gen

const templateBaseServiceContent = `
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
`

const templateConstServiceContent = `
package service

import (
	"errors"
	"{moduleName}/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type PostService struct {
	Base *BaseService
}

func New{UpperTableName}Service() *{UpperTableName}Service {
	return &{UpperTableName}Service{
		Base: new(BaseService),
	}
}
func (ser *{UpperTableName}Service) CountBy(where string ,params []interface{}) (int, error){
	count ,err := model.Model{fullTableName}.Where(where, params).Count()
	if err != nil{
		return 0, err
	}
	return count, nil
}
func (service *{UpperTableName}Service) ListBy(r *ghttp.Request,where string, params[]interface{}, page int, pageSize int, orderBy string)(map[string]interface{}, int, error){
	tableName := model.Table{fullTableName}
	listInput := map[string]interface{}{
		"where":where,
		"params":params,
		"orderBy":orderBy,
		"select":"p.id,p.created_at,p.title",
	}
	result,err := service.Base.List(r,tableName+" AS p", listInput)
	if  err!=nil {
		return result, gconv.Int(result["totalCount"]), errors.New(err.Error())
	}
	return result, gconv.Int(result["totalCount"]), nil
}

func (service *{UpperTableName}Service) FindBy(where string, params []interface{}) (*model.BingoPost, error) {
	var post *model.{fullTableName}
	input := map[string]interface{}{
		"where":where,
		"params":params,
	}
	ar := service.Base.FindBy(model.Table{fullTableName},input)
	err := ar.Struct(&post)
	return post,err
}
func (service *{UpperTableName}Service) FindById(id int) (*model.{fullTableName}, error){
	where := "id=?"
	var params []interface{}
	params = append(params, id)
	post, err := service.FindBy(where,params)
	return post, err
}
func (ps *{UpperTableName}Service) Create(r *ghttp.Request) (*model.{fullTableName},error) {
	request := r.GetPostMap()
	//rules := map[string]string {
	//	"title"  : "required|length:1,200",
	//	"content"  : "required",
	//	"cid" : "required",
	//	"intro" : "length:0,500",
	//}
	//msgs  := map[string]interface{} {
	//	"title" : "标题不能为空|标题的长度应当在:1到:200之间",
	//	"content" : "内容不能为空",
	//	"cid":"请选择分类",
	//	"intro":"摘要的长度只多只能为500个字符",
	//}
	//post := &model.BingoPost{}
	//if bindErr := r.GetRequestToStruct(post);bindErr!=nil{
	//	return post,bindErr
	//}
	//if e := gvalid.CheckMap(request, rules, msgs); e != nil {
	//	return post, errors.New(e.FirstString())
	//}
	_, err := post.Insert()
	if err != nil{
		return post, errors.New(err.Error())
	}
	return post, nil
}
func (ps *{UpperTableName}Service) Update(r *ghttp.Request) (*model.{fullTableName}, error) {
	request := r.GetPostMap()
	//rules := map[string]string {
	//	"title"  : "required|length:1,200",
	//	"content"  : "required",
	//	"cid" : "required",
	//	"intro" : "length:0,500",
	//}
	//msgs  := map[string]interface{} {
	//	"title" : "标题不能为空|标题的长度应当在:1到:200之间",
	//	"content" : "内容不能为空",
	//	"cid":"请选择分类",
	//	"intro":"摘要的长度只多只能为500个字符",
	//}
	//post := &model.{tableStruct}
	//if bindErr := r.GetRequestToStruct(post);bindErr!=nil{
	//	return post,errors.New(bindErr.Error())
	//}
	//if e := gvalid.CheckMap(request,rules,msgs);e != nil{
	//	return post, errors.New(e.FirstString())
	//}
	db := g.DB()
	db.SetDebug(true)
	params := []interface{}{post.Id}
	_, err:=db.Table(model.{fullTableName}).Data(post).Where("id=?", params).Update()
	if err != nil{
		return post, errors.New(err.Error())
	}
	return post,nil
}
`