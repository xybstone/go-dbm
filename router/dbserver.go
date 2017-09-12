package dbm_router

import "github.com/valyala/fasthttp"

//DBServer database server api router
type DBServer struct {
	BaseServer
}

//GetServers 获取数据库服务
func (dbs *DBServer) GetServers(ctx *fasthttp.RequestCtx) {
	res := APIBaseResponse{}
	res.HTTPCode = 200
	res.Result = "待实现"
	dbs.ServerJSON(ctx, &res)
	return
}
