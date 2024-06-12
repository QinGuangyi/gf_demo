package cmd

import (
	"context"
	"gf_demo/internal/controller/hello"
	"gf_demo/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

// func handler(req *ghttp.Request) {
// 	req.Response.Writeln("Hello GF")
// }

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// s.BindHandler("/he", handler)
			// hello := hello.NewHello()
			// s.BindHandler("/say", hello.SayHello)
			// s.BindObject("/user", user.New(), "Add, Get")

			//第三种写法： /api/hello  /api/user
			// s.Group("api", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)

			// 	group.Group("/hello", func(group1 *ghttp.RouterGroup) {
			// 		group1.Bind(
			// 			hello.NewHello(),
			// 		)
			// 	})
			// 	group.Group("/user", func(group2 *ghttp.RouterGroup) {
			// 		group2.Bind(
			// 			user.New(),
			// 		)
			// 	})

			// })

			//第一种写法： api/delete  api/hello
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewHello(),
					user.New(),
				)
			})

			//第二种写法： /user/ /hello
			// s.Group("/hello", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)
			// 	group.Bind(
			// 		hello.NewHello(),
			// 	)
			// })
			// s.Group("/user", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)
			// 	group.Bind(
			// 		user.New(),
			// 	)
			// })
			// s.SetServerRoot("resource/public")
			s.Run()
			return nil
		},
	}
)
