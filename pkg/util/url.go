package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

// URL调用封装工具
func CallURL(method, url, path string, header map[string]string, data interface{}) (
	statusCode int, body []byte, err error) {
	log.Logger.Infof("call url http://%s%s, method: %s, body: %s", url, path, method, PrintObjectJson(data))
	params, err := encodeData(data)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to encode data, err: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("http://%s%s", url, path), params)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to new request, %v", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to send http request, err: %v", err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read body, err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		if len(body) == 0 {
			return resp.StatusCode, nil, fmt.Errorf("call url error %s", http.StatusText(resp.StatusCode))
		}
		return resp.StatusCode, body, fmt.Errorf("call url error: %s", bytes.TrimSpace(body))
	}

	return resp.StatusCode, body, nil
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		if _, err := params.Write(buf); err != nil {
			return nil, err
		}
	}
	return params, nil
}
