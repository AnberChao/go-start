package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context

	//	是否超时标记位
	hasTimeout bool
	//	写保护机制
	writerMux *sync.Mutex

	// 当前请求的handler链条
	handlers []ControllerHandler
	index    int //当前请求调用到调用链到哪个节点

	params map[string]string //url路由匹配的参数
}

//NewContext 初始化一个Context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

//	封装基本函数方法

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// 核心函数，调用Context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

//	实现Context接口
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

////	封装http.Request，实现查询url
//func (ctx *Context) QueryAll() map[string][]string {
//	if ctx.request != nil {
//		return map[string][]string(ctx.request.URL.Query())
//	}
//	return map[string][]string{}
//}
//
//func (ctx *Context) QueryInt(key string, def int) int {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			intval, err := strconv.Atoi(vals[len-1])
//			if err != nil {
//				return def
//			}
//			return intval
//		}
//	}
//	return def
//}

//func (ctx *Context) QueryString(key string, def string) string {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			return vals[len-1]
//		}
//	}
//	return def
//}

//func (ctx *Context) QueryArray(key string, def []string) []string {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		return vals
//	}
//	return def
//}
//
//// #region form post
//func (ctx *Context) FormInt(key string, def int) int {
//	params := ctx.FormAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			intval, err := strconv.Atoi(vals[len-1])
//			if err != nil {
//				return def
//			}
//			return intval
//		}
//	}
//	return def
//}
//
//func (ctx *Context) FormString(key string, def string) string {
//	params := ctx.FormAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			return vals[len-1]
//		}
//	}
//	return def
//}
//
//func (ctx *Context) FormArray(key string, def []string) []string {
//	params := ctx.FormAll()
//	if vals, ok := params[key]; ok {
//		return vals
//	}
//	return def
//}

//func (ctx *Context) FormAll() map[string][]string {
//	if ctx.request != nil {
//		return map[string][]string(ctx.request.PostForm)
//	}
//	return map[string][]string{}
//}

////	绑定请求
//func (ctx *Context) BindJson(obj interface{}) error {
//	if ctx.request != nil {
//		body, err := ioutil.ReadAll(ctx.request.Body)
//		if err != nil {
//			return err
//		}
//		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
//
//		err = json.Unmarshal(body, obj)
//		if err != nil {
//			return err
//		}
//	} else {
//		return errors.New("ctx.request empty")
//	}
//	return nil
//}

//	封装了http.ResponseWriter对外接口

//
