package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// 提交参数
type Params map[string]interface{}
// 转换成URL参数
func (this Params) ToUrlQuery()(string, error){
	values, err := url.ParseQuery("")
	if err != nil{
		return "", err
	}
	for k, v := range this{
		values.Set(k, fmt.Sprintf("%v", v))
	}
	return values.Encode(), nil
}
// 转换成POST数据
func (this Params) ToPostData()*strings.Reader{
	var list []string
	for k, v := range this{
		list = append(list, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.NewReader(strings.Join(list, "&"))
}
// 转换成就是偶呢数据
func (this Params) ToJsonData()(*bytes.Buffer, error){
	json_data, err := json.Marshal(&this)
	if err != nil{
		return nil, err
	}
	return bytes.NewBuffer(json_data), nil
}