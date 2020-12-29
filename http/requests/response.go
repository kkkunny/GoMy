package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"my/charset"
	"net/http"
	"os"
	"strconv"
)

// 将http包中的Response转换成Response
func HttpToResponse(http_response *http.Response, charsets ...string)*Response{
	var response Response
	// 状态码
	response.Status = strconv.Itoa(http_response.StatusCode)
	// 头
	response.Headers = http_response.Header
	// 读入器
	response.reader = http_response.Body
	// 长度
	response.Length = http_response.ContentLength
	// 编码
	if len(charsets) > 0{
		response.Charset = charsets[0]
	}
	return &response
}
// 回复
type Response struct {
	Status  string  // 状态码
	Headers http.Header  // 回复头
	reader  io.ReadCloser  // 读取器
	data []byte  // 数据
	Length int64  // 数据长度
	Charset string  // 读取编码（以什么编码读取内容）
}
// 获取数据体
func (this *Response)Body()[]byte{
	if len(this.data) == 0{
		data, err := ioutil.ReadAll(this)
		if err != nil{
			return []byte{}
		}
		if this.Charset != ""{
			data, err = charset.Convert(this.Charset, "utf-8", data)
			if err != nil{
				return []byte{}
			}
		}
		this.data = data
	}
	return this.data
}
// 获取Text
func (this *Response)Text()string{
	return string(this.Body())
}
// 获取JSON
func (this *Response)Json(data interface{})error{
	body := this.Body()
	return json.Unmarshal(body, data)
}
// 获取HTML
func (this *Response)Html()(*HtmlNode, error){
	doc, err := html.Parse(bytes.NewReader(this.Body()))
	return &HtmlNode{doc}, err
}
// 保存成文件(适用于小文件)
func (this *Response)SaveToFile(path string)error{
	_, err := os.Open(path)
	if err == nil{
		return ErrFileExist
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		return err
	}
	num, err := file.Write(this.Body())
	if err != nil{
		return err
	}else if num == 0{
		return errors.New("No bytes were writed")
	}
	return file.Close()
}
// 读取数据
func (this *Response)Read(num []byte)(int, error){
	return this.reader.Read(num)
}
// 关闭连接
func (this *Response)Close()error{
	return this.reader.Close()
}