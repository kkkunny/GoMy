package web

import (
	"GoMy/re"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 上传文件
type uploadFile struct {
	file   multipart.File
	header *multipart.FileHeader
}

// 获取文件名
func (this *uploadFile) GetFileName() string {
	dis := this.header.Header.Get("Content-Disposition")
	result := re.FindAll("filename=\"(.*?)\"", dis)
	if len(result) > 0 && len(result[0]) >= 1 {
		return result[0][1]
	}
	return ""
}

// 获取内容
func (this *uploadFile) GetData() []byte {
	data, err := ioutil.ReadAll(this.file)
	if err != nil {
		return []byte{}
	}
	return data
}

// 获取reader
func (this *uploadFile) GetReadCloser() io.ReadCloser {
	return this.file
}

// 将write和request转化成context
func NewContext(w http.ResponseWriter, r *http.Request, params map[string]string) *Context {
	return &Context{
		writer:      w,
		req:         r,
		routeParams: params,
	}
}

// 请求
type Context struct {
	writer      http.ResponseWriter
	req         *http.Request
	routeParams map[string]string
}

// 获取请求方式
func (this *Context) GetMethod() string {
	return this.req.Method
}

// 获取请求url
func (this *Context) GetUrl() *url.URL {
	return this.req.URL
}

// 获取全部请求头
func (this *Context) GetReqHeaders() http.Header {
	return this.req.Header
}

// 获取请求头
func (this *Context) GetOneReqHeader(key string) string {
	return this.req.Header.Get(key)
}

// 获取客户端ip
func (this *Context) GetRequestIp() string {
	// 反向代理
	xForwardedFor := this.GetOneReqHeader("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(this.GetOneReqHeader("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	// 直接请求
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(this.req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// 获取数据
func (this *Context) GetData() []byte {
	datas, err := ioutil.ReadAll(this.req.Body)
	if err != nil {
		return []byte{}
	}
	this.req.Body = ioutil.NopCloser(bytes.NewReader(datas))
	return datas
}

// 获取url参数
func (this *Context) GetUrlParam(key string) string {
	return this.req.URL.Query().Get(key)
}

// 获取post数据(multipart/form-data)
func (this *Context) getPostDataFromData(key string) string {
	return this.req.PostFormValue(key)
}

// 获取post数据(application/x-www-form-urlencoded)
func (this *Context) getPostDataXFormUrlencoded(key string) string {
	u, err := url.Parse("/test?" + string(this.GetData()))
	if err == nil && u != nil {
		return u.Query().Get(key)
	}
	return ""
}

// 获取POST数据
func (this *Context) GetPostData(key string) string {
	// 获取编码
	var contentType = "multipart/form-data"
	encode := this.GetOneReqHeader("Content-Type")
	if len(encode) > 0 {
		if spl := strings.Split(encode, ";"); len(spl) > 0 {
			contentType = spl[0]
		}
	}
	// 获取
	switch contentType {
	case "multipart/form-data":
		return this.getPostDataFromData(key)
	case "application/x-www-form-urlencoded":
		return this.getPostDataXFormUrlencoded(key)
	default:
		return ""
	}
}

// 获取路由参数
func (this *Context) GetRouteParam(key string) string {
	value, ok := this.routeParams[key]
	if ok {
		return value
	} else {
		return ""
	}
}

// 获取上传文件
func (this *Context) GetFile(key string) *uploadFile {
	file := new(uploadFile)
	f, header, err := this.req.FormFile(key)
	if err != nil {
		return nil
	}
	file.file = f
	file.header = header
	return file
}

// 获取json
func (this *Context) GetJson(data interface{}) error {
	return json.Unmarshal(this.GetData(), data)
}

// 获取Cookie
func (this *Context) GetCookie(name string) *http.Cookie {
	cookie, err := this.req.Cookie(name)
	if err != nil {
		return nil
	} else {
		return cookie
	}
}

// 设置响应头
func (this *Context) SetOneHeader(key string, value string) {
	this.writer.Header().Set(key, value)
}

// 设置Cookie
func (this *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(this.writer, cookie)
}

// 删除Cookie
func (this *Context) DeleteCookie(name string) {
	cookie := &http.Cookie{
		Name:       name,
		Value:      "",
		Path:       "/",
		Domain:     "",
		Expires:    time.Now().Add(-100 * time.Hour),
		RawExpires: "",
		MaxAge:     -1,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	this.SetCookie(cookie)
}

// 修改返回状态码,修改完后不允许修改头部
func (this *Context) SetStatusCode(code int) {
	this.writer.WriteHeader(code)
}

// 返回字符串
func (this *Context) ReturnString(code int, msg string) error {
	fmt.Printf("返回数据:[%s]\n", msg)
	this.SetStatusCode(code)
	_, err := io.WriteString(this.writer, msg)
	return err
}

// 返回404
func (this *Context) ReturnNotFound() error {
	fmt.Printf("返回数据:[%s]\n", "404 Page not found")
	return this.ReturnString(http.StatusNotFound, "404 Page not found")
}

// 请求方式不支持
func (this *Context) ReturnMethodAllowed() error {
	fmt.Printf("返回数据:[%s]\n", "405 Request method is not allowed")
	return this.ReturnString(http.StatusMethodNotAllowed, "405 Request method is not allowed")
}

// 返回json
func (this *Context) ReturnJson(code int, data interface{}) error {
	this.SetStatusCode(code)
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Printf("返回数据:[%s]\n", string(js))
	_, err = this.writer.Write(js)
	return err
}

// 返回文件
func (this *Context) ReturnFile(reader io.Reader, filename string) error {
	this.SetOneHeader("Content-Type", "application/octet-stream")
	this.SetOneHeader("Content-Disposition", "attachment;filename="+filename)
	datas, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Printf("返回数据:[文件: %s, 大小: %d]\n", filename, len(datas))
	_, err = this.writer.Write(datas)
	return err
}

// 返回模板
func (this *Context) ReturnTemplate(data interface{}, path ...string) error {
	temp, err := template.ParseFiles(path...)
	if err != nil {
		return err
	}
	fmt.Printf("返回数据:[模板: %v]\n", path)
	return temp.Execute(this.writer, data)
}

// 重定向
func (this *Context) Redirect(url string) {
	http.Redirect(this.writer, this.req, url, http.StatusMovedPermanently)
}
