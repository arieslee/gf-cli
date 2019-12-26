/*
@Time : 2019/12/16 8:57 上午
@Author : sunmoon
@File : template_service
@Software: GoLand
*/
package gen


const templateBaseApiContent = `
// ==========================================================================
// @Author : sunmoon
// @generator: gf-cli
// ==========================================================================
package api

type BaseApi struct {

}
`

const templateConstApiContent = `
// ==========================================================================
// @Time : {nowTime}
// @Author : sunmoon
// @File : {tableName}.go
// @generator: gf-cli
// ==========================================================================
package {tableName}
import (
    "{moduleName}/app/api"
    "github.com/gogf/gf/net/ghttp"
)
type {UpperTableName}Api struct {
	Base *api.BaseApi
}
func New{UpperTableName}Api() *{UpperTableName}Api {
	return &{UpperTableName}Api{
		Base: new(api.BaseApi),
	}
}
// 创建
func CreateHandler(r *ghttp.Request) {
    r.Response.Writeln("this is create @ {tableName}")
}
// 更新
func UpdateHandler(r *ghttp.Request) {
    r.Response.Writeln("this is update @ {tableName}")
}
// 删除
func DeleteHandler(r *ghttp.Request) {
    r.Response.Writeln("this is delete @ {tableName}")
}
// list
func ListHandler(r *ghttp.Request) {
    r.Response.Writeln("this is list @ {tableName}")
}
// view
func DetailHandler(r *ghttp.Request) {
    r.Response.Writeln("this is detail @ {tableName}")
}
`