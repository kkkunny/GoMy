package charset

import (
	"bytes"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// 通过数据获取编码
func GetCharsetByData(data []byte)encoding.Encoding{
	encode, _, _ := charset.DetermineEncoding(data, "")
	return encode
}
// 通过名字获取编码
func GetCharsetByName(name string)(encoding.Encoding, error){
	encode, err := ianaindex.MIB.Encoding(name)
	if err != nil{
		return encode, err
	}else if encode == nil{
		return encode, errors.New("No this encode")
	}
	return encode, nil
}
// 转换
func Convert(src, dst string, data []byte)([]byte, error){
	if dst == src{
		return data, nil
	}else{
		// 转换成UTF-8
		src_charset, err := GetCharsetByName(src)
		if err != nil{
			return []byte{}, err
		}
		result, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data), src_charset.NewDecoder()))
		if err != nil{
			return []byte{}, err
		}
		// 转成成其他编码
		if dst != "utf-8" && dst != "UTF-8"{
			dst_charset, err := GetCharsetByName(dst)
			if err != nil{
				return []byte{}, err
			}
			return ioutil.ReadAll(transform.NewReader(bytes.NewReader(result), dst_charset.NewEncoder()))
		}
		return result, nil
	}
}