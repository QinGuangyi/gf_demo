// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct{
	// g.Meta `path:"/params/{name}/{id}.html" method:"all"` //动态路由
	g.Meta `method:"all"`

	Username string `p:"name" d:"覃光怡"`
	Password string 
	Age int
	
}

type ParamsRes struct{

}


