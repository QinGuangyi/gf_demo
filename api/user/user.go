package user

import "github.com/gogf/gf/v2/frame/g"

type AddReq struct {
	g.Meta `path:"/adduser" method:"post"`
}
type AddRes struct {
	Name string
	Age  int
}
