package requests

import (
	"io"
	"net/http"
	"net/url"
)

// 获取新的爬虫
func NewRequest() *Request {
	head := http.Header{}
	head.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
	return &Request{Client: &http.Client{}, Headers: head}
}

// 获取一个带代理的爬虫（例：socks5://127.0.0.1:1080）
func NewRequestWithProxy(proxy string) *Request {
	client := &http.Client{}
	if proxy != "" {
		p := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
		client.Transport = &http.Transport{Proxy: p}
	}
	head := http.Header{}
	head.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
	return &Request{Client: client, Headers: head}
}

// 请求
type Request struct {
	Client  *http.Client
	Headers http.Header
}

// GET请求
func (this *Request) Get(url string, params Params, charsets ...string) (*Response, error) {
	// 请求
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 参数
	if params != nil {
		request.URL.RawQuery, err = params.ToUrlQuery()
		if err != nil {
			return nil, err
		}
	}
	// 请求头
	request.Header = this.Headers
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 执行
	http_response, err := this.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return HttpToResponse(http_response, charsets...), err
}

// POST请求
func (this *Request) Post(url string, data Params, isjson bool, charsets ...string) (*Response, error) {
	// 参数
	var post_data io.Reader
	var err error
	if !isjson {
		post_data = data.ToPostData()
	} else {
		post_data, err = data.ToJsonData()
		if err != nil {
			return nil, err
		}
	}
	// 请求
	request, err := http.NewRequest(http.MethodPost, url, post_data)
	if err != nil {
		return nil, err
	}
	// 请求头
	request.Header = this.Headers
	if !isjson {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		request.Header.Set("Content-Type", "application/json")
	}
	// 执行
	http_response, err := this.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return HttpToResponse(http_response, charsets...), err
}

// HEAD请求
func (this *Request) Head(url string, charsets ...string) (*Response, error) {
	// 请求
	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	// 请求头
	request.Header = this.Headers
	// 执行
	http_response, err := this.Client.Do(request)
	if err != nil {
		return nil, err
	}
	return HttpToResponse(http_response, charsets...), err
}

// 设置代理（例：socks5://127.0.0.1:1080）
func (this *Request) SetProxy(proxy string) {
	p := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxy)
	}
	transport := &http.Transport{Proxy: p}
	this.Client.Transport = transport
}
