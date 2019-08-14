package bootstrap

import (
	"github.com/kataras/iris"
	"iris-blog/common"
	"strings"
	"time"
)

// 配置函数类型
type Configurator func(b *BootStrapper)
type Eorrorhandler func(context iris.Context)

type BootStrapper struct {
	*iris.Application

	// 实例编码
	code string

	// 启动时间
	runtime time.Time
}

/*
code
cfgs  分组  配置函数的分组
*/
func NewBootStrapper(code string, cfgs ...Configurator) *BootStrapper {

	code = strings.ReplaceAll(code, " ", "")
	if code == "" {
		code = common.GenUUID()
	}

	b := &BootStrapper{
		Application: iris.New(),
		code:        code,
		runtime:     time.Now(),
	}

	//for _, c := range cfgs {
	//	c(b)
	//}
	b.Configure(cfgs...)

	return b

}

// 配置
func (self *BootStrapper) Configure(cfgs ...Configurator) {
	for _, c := range cfgs {
		c(self)
	}
}

//
func (self *BootStrapper) Setup() {
	self.SetupServerErrorHandler(400, 405)

}

//-------------http 错误码处理  start-----------------------

//	 100 <=code <200 或 400 <= code <=511 都会在handler 处理
func (self *BootStrapper) SetupHttpAnyErrorHandler() {
	handler := func(context iris.Context) {
		err := iris.Map{
			"app_code":    self.code,
			"status_code": context.GetStatusCode(),
		}
		// json
		context.JSON(err)
		return

	}

	self.OnAnyErrorCode(handler)
}

// 指定status code由 handler处理
func (self *BootStrapper) SetupServerErrorHandler(start, end int) {

	serverEorrorHandler := func(c iris.Context) {
		err := iris.Map{
			"app_code":    self.code,
			"status_code": c.GetStatusCode(),
			"time":        "now",
		}
		c.JSON(err)
		return
	}

	for status_code := start; status_code <= end; status_code++ {
		// 指定status_code
		self.OnErrorCode(status_code, serverEorrorHandler)
	}

}

//-------------http 错误码处理--  end ---------------------

func (self *BootStrapper) Listen(addr string, cfgs ...iris.Configurator) {

	self.Run(iris.Addr(addr), cfgs...)
}
