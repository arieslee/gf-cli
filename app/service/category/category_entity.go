// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package category

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table bingo_category.
type Entity struct {
    Id        int    `orm:"id,primary"       json:"id"`         //           
    CateName  string `orm:"cate_name,unique" json:"cate_name"`  // 名称      
    Slug      string `orm:"slug,unique"      json:"slug"`       // 缩略名    
    Counts    int    `orm:"counts"           json:"counts"`     // 文章数量  
    ParentId  int    `orm:"parent_id"        json:"parent_id"`  // 上级id    
    Intro     string `orm:"intro"            json:"intro"`      // 介绍      
    ListOrder int    `orm:"list_order"       json:"list_order"` // 排序      
    CreatedAt int    `orm:"created_at"       json:"created_at"` // 创建时间  
    UpdatedAt int    `orm:"updated_at"       json:"updated_at"` // 更新时间  
    Cover     string `orm:"cover"            json:"cover"`      // 封面      
    Template  string `orm:"template"         json:"template"`   // 模板      
    Status    int    `orm:"status"           json:"status"`     // 状态      
}

// Category is alias of Entity, which some developers say they just want.
type Category = Entity

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}