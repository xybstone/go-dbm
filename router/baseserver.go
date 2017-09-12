package dbm_router

import (
	"encoding/json"
	"os"
	"strings"
	"ucenter/config/error_code"

	"gitlab.gaodun.com/golib/filetool"

	"github.com/valyala/fasthttp"
)

//BaseServer 基础服务
type BaseServer struct {
	Ctx *fasthttp.RequestCtx
}

//ServerJSON 服务器返回
func (ba *BaseServer) ServerJSON(ctx *fasthttp.RequestCtx, v interface{}) {
	if b, err := json.Marshal(v); err == nil {
		ctx.Write(b)
	}
}

//HandleRoot 测试页
func (ba *BaseServer) HandleRoot(ctx *fasthttp.RequestCtx) {
	SetAPICustomHeader(ctx)
	ctx.Write([]byte("{\"version\": \"1.0.0.2\"}"))
	return
}

//GetStatus 状态页
func (ba *BaseServer) GetStatus(ctx *fasthttp.RequestCtx) {
	dir, _ := os.Getwd()
	fileOperate := filetool.FileOperate{Filename: dir + string(os.PathSeparator) + "DEPLOY"}
	if fileOperate.CheckFileExist(fileOperate.Filename) {
		text, _ := fileOperate.ReadFile(fileOperate.Filename)
		lineList := strings.Split(string(text), "\n")
		begin := "{"
		tmpStr := ""
		for _, line := range lineList {
			if len(line) == 0 {
				break
			}
			n := strings.Split(line, "|")
			jsonKey := "\"" + n[0] + "\""
			jsonValue := ":\"" + n[1] + "\""
			tmpStr += jsonKey + jsonValue + ","
		}
		mStr := strings.TrimRight(tmpStr, ",")
		end := "}"
		info := begin + mStr + end
		ctx.Write([]byte("{\"status\": \"1\",\"data\":" + info + "}"))
		return
	}

	ctx.Write([]byte("{\"status\": \"1\"}"))
	return
}

//SetAPICustomHeader 设置头
func SetAPICustomHeader(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.SetServerBytes([]byte("ucenter-api v1.0.0"))
	ctx.SetContentType("application/json; charset=utf8")
}

//APIBaseResponse  返回参数
type APIBaseResponse struct {
	HTTPCode int                      `json:"http_code"`
	Status   error_code.Code_type_int `json:"status"`
	Info     string                   `json:"info"`
	Result   interface{}              `json:"result"`
}

//NewBaseServer 新服务
func NewBaseServer() *BaseServer {
	return new(BaseServer)
}
