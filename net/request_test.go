package net

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	tests := []struct {
		name    string
		args    *RequestParams
		wantErr bool
	}{
		{name: "http", args: NewRequestParamsWithLocalIP("http://www.baidu.com/", ""), wantErr: false},
		{name: "https", args: NewRequestParamsWithLocalIP("https://www.baidu.com/", ""), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Request(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(got.StatusCode == 200) {
				t.Errorf("Request() = %v, want %v", got, 200)
			}
		})
	}
}

func TestRequestMultiConcurrentlyForFirstOKResponse(t *testing.T) {
	paramsList := []*RequestParams{}
	urls := []string{"https://1.1.1.1/", "http://2.2.2.2/", "http://3.3.3.3/", "https://www.baidu.com/"}
	for _, url := range urls {
		paramsList = append(paramsList,
			NewRequestParamsWithLocalIP(url, ""))
	}
	resp, err := RequestMultiConcurrentlyForFirstOKResponse(paramsList, 5*time.Second)
	assert.NoError(t, err)
	t.Log(resp.Request.URL, resp.StatusCode)
}

func TestRequestMultiConcurrentlyForFirstOKResponseAllFail(t *testing.T) {
	paramsList := []*RequestParams{}
	urls := []string{"https://1.2.2.2/", "http://1.2.2.3/", "http://1.3.3.3/", "http://wesong.top/302"}
	for _, url := range urls {
		paramsList = append(paramsList,
			NewRequestParamsWithLocalIP(url, ""))
	}
	_, err := RequestMultiConcurrentlyForFirstOKResponse(paramsList, 5*time.Second)
	assert.NotEqual(t, err, nil)
}
