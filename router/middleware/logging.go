package middleware

import (
	"apiServerDemo/handler"
	"apiServerDemo/pkg/errno"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"io/ioutil"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().UTC()
		path := ctx.Request.URL.Path

		// 该中间件只记录业务请求，比如 /v1/user 和 /login 路径。
		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		// 忽略健康检查请求
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// 读取body内容
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}

		// 将io.ReadCloser恢复到原始状态
		//
		//  HTTP 的请求 Body，在读取过后会被置空，所以这里读取完后会重新赋值：
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// 基础信息/请求模式|请求者ip
		method := ctx.Request.Method
		ip := ctx.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = blw

		ctx.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// 得到code和message
		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "回复消息体(body)不能被转换成对象，body的内容为: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = errno.OK.Code
			message = errno.OK.Message
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)

	}
}
