package net

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

type RequestParams struct {
	URL            string
	LocalIP        string
	ConnectTimeout time.Duration
	TotalTimeout   time.Duration
}

func NewRequestParamsWithLocalIP(url, localIP string) *RequestParams {
	return &RequestParams{
		URL:            url,
		LocalIP:        localIP,
		ConnectTimeout: 5 * time.Second,
		TotalTimeout:   10 * time.Second,
	}
}

func NewRequestParamsWithLocalIPAndTimeout(url, localIP string, connnectTimeout, totalTimeout int) *RequestParams {
	return &RequestParams{
		URL:            url,
		LocalIP:        localIP,
		ConnectTimeout: time.Duration(connnectTimeout) * time.Second,
		TotalTimeout:   time.Duration(totalTimeout) * time.Second,
	}
}

func Request(r *RequestParams) (*http.Response, error) {
	localAddr, err := net.ResolveIPAddr("ip", r.LocalIP)
	if err != nil {
		return nil, err
	}

	localTCPAddr := net.TCPAddr{
		IP: localAddr.IP,
	}

	d := net.Dialer{
		LocalAddr: &localTCPAddr,
		Timeout:   r.ConnectTimeout,
		Deadline:  time.Now().Add(r.TotalTimeout),
	}

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial:  d.Dial,
	}

	webclient := &http.Client{
		Transport: tr,
		// 不跟随跳转
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: r.TotalTimeout,
	}
	defer webclient.CloseIdleConnections()

	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		return nil, err
	}

	res, err := webclient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func RequestMultiConcurrentlyForFirstOKResponse(paramsList []*RequestParams, timeout time.Duration) (ret *http.Response, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	length := len(paramsList)
	okChan := make(chan bool, length)
	for _, params := range paramsList {
		p := params
		p.TotalTimeout = timeout
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				return
			default:
				resp, err := Request(p)
				if err == nil && resp.StatusCode == http.StatusOK {
					ret = resp
					okChan <- true
					return
				}
				okChan <- false
			}
		}(ctx)
	}

	for range paramsList {
		r := <-okChan
		if r == true {
			return
		}
	}
	err = errors.New("no any response")
	return
}
