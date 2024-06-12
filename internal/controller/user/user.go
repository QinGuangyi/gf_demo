package user

import (
	"context"
	"gf_demo/api/user"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// func Handler(ctx context.Context, req *Request) (res *Response, err error)

func (c *Controller) Add(ctx context.Context, req *user.AddReq) (res *user.AddRes, err error) {
	res = &user.AddRes{
		Name: "martha",
		Age:  25,
	}
	// g.RequestFromCtx(ctx).Response.Writeln("hello") //如果采用该方法，将以自定义方式返回内容“hello”，并不会通过将内容转成json返回
	return

}

// func (c *Controller) Add(r *ghttp.Request) {
// 	r.Response.Writeln("添加用户")
// }

func (c *Controller) Update(r *ghttp.Request) {
	r.Response.Writeln("更新用户")
}

func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Writeln("删除用户")
}

func (c *Controller) Get(r *ghttp.Request) {
	r.Response.Writeln("查询用户")
}
