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
// 自动生成可以修改或删除
func (a *{UpperTableName}Api) CreateHandler(r *ghttp.Request) {
    r.Response.Writeln("this is create @ {UpperTableName}Api")
}
// 更新
// 自动生成可以修改或删除
func (a *{UpperTableName}Api) UpdateHandler(r *ghttp.Request) {
    r.Response.Writeln("this is update @ {UpperTableName}Api")
}
// 删除
// 自动生成可以修改或删除
func (a *{UpperTableName}Api) DeleteHandler(r *ghttp.Request) {
    r.Response.Writeln("this is delete @ {UpperTableName}Api")
}
// 列表
// 自动生成可以修改或删除
func (a *{UpperTableName}Api) ListHandler(r *ghttp.Request) {
    r.Response.Writeln("this is list @ {UpperTableName}Api")
}
// 详情
// 自动生成可以修改或删除
func (a *{UpperTableName}Api) DetailHandler(r *ghttp.Request) {
    r.Response.Writeln("this is detail @ {UpperTableName}Api")
}
`