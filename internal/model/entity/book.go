// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure for table book.
type Book struct {
	Id          uint        `json:"id"          orm:"id"           ` // ID
	Name        string      `json:"name"        orm:"name"         ` // 书名
	Author      string      `json:"author"      orm:"author"       ` // 作者
	Price       float64     `json:"price"       orm:"price"        ` // 价格
	PublishTime *gtime.Time `json:"publishTime" orm:"publish_time" ` // 出版时间
}
