package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"moul.io/http2curl/v2"
)

func DoRequest(method string, rawUrl string, headers map[string]string, body io.Reader) ([]byte, int, http.Header, error) {
	req, err := http.NewRequest(method, rawUrl, body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	// 非 200 的 Code 交给外层业务去判断
	if resp.StatusCode != http.StatusOK {
		return data, resp.StatusCode, resp.Header, nil
	}
	// 其余空包进行报错拦截
	if len(data) == 0 {
		return nil, 0, nil, errors.Errorf("empty text, url: %s", rawUrl)
	}
	return data, resp.StatusCode, resp.Header, nil
}

func ProxyRequest(method string, rawUrl string, headers map[string]string, body io.Reader) ([]byte, int, http.Header, error) {
	req, err := http.NewRequest(method, rawUrl, body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	trans := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
	httpClient := &http.Client{
		Transport: trans,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	// 非 200 的 Code 交给外层业务去判断
	if resp.StatusCode != http.StatusOK {
		return data, resp.StatusCode, resp.Header, nil
	}
	// 其余空包进行报错拦截
	if len(data) == 0 {
		return nil, 0, nil, errors.Errorf("empty text, url: %s", rawUrl)
	}
	return data, resp.StatusCode, resp.Header, nil
}

func DoCurl(method string, rawUrl string, headers map[string]string, body io.Reader) ([]byte, int, http.Header, error) {
	req, err := http.NewRequest(method, rawUrl, body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	{
		curlCmd, err := http2curl.GetCurlCommand(req)
		if err != nil {
			return nil, 0, nil, errors.WithStack(err)
		}
		fmt.Println(curlCmd.String())
		return []byte{}, http.StatusInternalServerError, nil, errors.New("debug")
	}
}
