package logger

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Info struct {
	Package string
	Action  string
	Message string
}

func setLogCtx(c *gin.Context, p string, action string) *log.Entry {

	ctx := log.WithFields(log.Fields{
		"Package": p,
		"Action":  action,
	})
	return ctx
}

func setLog(p string, action string) *log.Entry {
	ctx := log.WithFields(log.Fields{
		"Package": p,
		"Action":  action,
	})
	return ctx
}

/*
- Fatal
- Panic
- Error
- Warning
- Info
- Debug
*/

func FatalCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Fatal(info.Message)
}
func FatalLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Fatal(info.Message)
}

func PanicCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Panic(info.Message)
}
func PanicLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Panic(info.Message)
}
func ErrorCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Error(info.Message)
}

func ErrorLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Error(info.Message)
}
func WarningCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Warn(info.Message)
}

func WarningLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Warn(info.Message)
}
func InfoCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Info(info.Message)
}

func InfoLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Info(info.Message)
}
func DebugCtx(context *gin.Context, info Info) {
	ctx := setLogCtx(context, info.Package, info.Action)
	ctx.Debug(info.Message)
}

func DebugLog(info Info) {
	ctx := setLog(info.Package, info.Action)
	ctx.Debug(info.Message)
}
