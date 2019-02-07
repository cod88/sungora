package core

import (
	"bytes"
	"net/http"
	"net/url"
)

// ContraFace is an interface to uniform all controller handler.
type Controller interface {
	Init(w http.ResponseWriter, r *http.Request) (err error)
	GET()
	POST()
	PUT()
	DELETE()
	OPTIONS()
	Response()
}

// Контроллер для реализации api запросов в формате json
type ControllerApi struct {
	RW      *requestResponse
	Session *sessionTyp
	Data    interface{}
}

func (c *ControllerApi) Init(w http.ResponseWriter, r *http.Request) (err error) {
	c.RW = NewRW(r, w)
	// request parameter "application/x-www-form-urlencoded"
	c.RW.RequestParams, _ = url.ParseQuery(r.URL.Query().Encode())
	if err = r.ParseForm(); err != nil {
		return
	}
	for i, v := range r.Form {
		c.RW.RequestParams[i] = v
	}
	// initialization session
	var token string
	if 0 < Config.SessionTimeout {
		if token, err = c.RW.CookieGet(Config.ServiceName); err != nil {
			return
		}
		if token == "" {
			token = NewRandomString(10)
			c.RW.CookieSet(Config.ServiceName, token)
		}
		c.Session = GetSession(token)
	}
	return
}

func (c *ControllerApi) GET() {
	return
}
func (c *ControllerApi) POST() {
	return
}
func (c *ControllerApi) PUT() {
	return
}
func (c *ControllerApi) DELETE() {
	return
}
func (c *ControllerApi) OPTIONS() {
	return
}

func (c *ControllerApi) Response() {
	if c.RW.isResponse {
		return
	}
	c.RW.ResponseJson(c.Data, c.RW.Status)
}

// Контроллер для реализации html страниц
type ControllerHtml struct {
	RW            *requestResponse
	Session       *sessionTyp
	Variables     map[string]interface{} // Variable (по умолчанию пустой)
	Functions     map[string]interface{} // html/template.FuncMap (по умолчанию пустой)
	TplController string
	TplLayout     string
}

func (c *ControllerHtml) Init(w http.ResponseWriter, r *http.Request) (err error) {
	c.RW = NewRW(r, w)
	// request parameter "application/x-www-form-urlencoded"
	c.RW.RequestParams, _ = url.ParseQuery(r.URL.Query().Encode())
	if err = r.ParseForm(); err != nil {
		return
	}
	for i, v := range r.Form {
		c.RW.RequestParams[i] = v
	}
	// initialization session
	var token string
	if 0 < Config.SessionTimeout {
		if token, err = c.RW.CookieGet(Config.ServiceName); err != nil {
			return
		}
		if token == "" {
			token = NewRandomString(10)
			c.RW.CookieSet(Config.ServiceName, token)
		}
		c.Session = GetSession(token)
	}
	//
	c.Functions = make(map[string]interface{})
	c.Variables = make(map[string]interface{})
	c.TplLayout = Config.DirWww + "/layout/index.html"
	c.TplController = Config.DirWww + "/controllers"
	return
}

func (c *ControllerHtml) GET() {
	return
}
func (c *ControllerHtml) POST() {
	return
}
func (c *ControllerHtml) PUT() {
	return
}
func (c *ControllerHtml) DELETE() {
	return
}
func (c *ControllerHtml) OPTIONS() {
	return
}

func (c *ControllerHtml) Response() {
	if c.RW.isResponse {
		return
	}
	var err error
	// шаблон контроллера
	var buf bytes.Buffer
	if buf, err = TplCompilation(c.TplController, c.Functions, c.Variables); err != nil {
		c.RW.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	if c.TplLayout == "" {
		c.RW.ResponseHtml(buf.String(), 200)
		return
	}
	// шаблон макета
	c.Variables["Content"] = buf.String()
	if buf, err = TplCompilation(c.TplLayout, c.Functions, c.Variables); err != nil {
		c.RW.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	c.RW.ResponseHtml(buf.String(), 200)
	return
}
