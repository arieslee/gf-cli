/*
@Time : 2019/12/16 8:57 上午
@Author : sunmoon
@File : template_service
@Software: GoLand
*/
package gen

const templateConstApiContent = `
// ==========================================================================
// @Time : {nowTime}
// @Author : sunmoon
// @File : {tableName}
// @generator: gf-cli
// ==========================================================================
package {tableName}
import (
    "github.com/gogf/gf/net/ghttp"
)

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