package service

import (
	"errors"
	"gf-app/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

type PostService struct {
	Base *BaseService
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		Base: new(BaseService),
	}
}
func (ser *CategoryService) CountBy(where string ,params []interface{}) (int, error){
	count ,err := model.ModelBingoCategory.Where(where, params).Count()
	if err != nil{
		return 0, err
	}
	return count, nil
}
func (service *CategoryService) ListBy(r *ghttp.Request,where string, params[]interface{}, page int, pageSize int, orderBy string)(map[string]interface{}, int, error){
	tableName := model.TableBingoCategory
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

func (service *CategoryService) FindBy(where string, params []interface{}) (*model.BingoPost, error) {
	var post *model.BingoCategory
	input := map[string]interface{}{
		"where":where,
		"params":params,
	}
	ar := service.Base.FindBy(model.TableBingoCategory,input)
	err := ar.Struct(&post)
	return post,err
}
func (service *CategoryService) FindById(id int) (*model.BingoCategory, error){
	where := "id=?"
	var params []interface{}
	params = append(params, id)
	post, err := service.FindBy(where,params)
	return post, err
}
func (ps *CategoryService) Create(r *ghttp.Request) (*model.BingoCategory,error) {
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
func (ps *CategoryService) Update(r *ghttp.Request) (*model.BingoCategory, error) {
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
	_, err:=db.Table(model.BingoCategory).Data(post).Where("id=?", params).Update()
	if err != nil{
		return post, errors.New(err.Error())
	}
	return post,nil
}