package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/huweihuang/gin-api-frame/pkg/types"

	"github.com/huweihuang/gin-api-frame/pkg/handler"
)

const (
	url     = "localhost"
	reqPath = "/api/v1/instance"
	jobID   = "1xxxxxxxxxxxxx0"
)

var (
	request = &types.Instance{}
	delReq  = &types.Instance{}
)

func TestCreateInstance(t *testing.T) {
	targetUrl := fmt.Sprintf("http://%s%s", url, reqPath)
	reqStr, err := json.Marshal(request)
	if err != nil {
		t.Errorf("Failed to marshal database: %v", err)
	}

	body := bytes.NewReader(reqStr)
	req, err := http.NewRequest("PUT", targetUrl, body)
	if err != nil {
		t.Fatalf("Failed to create request, %v", err)
	}

	c, e, resp := ginTest(req)
	e.Use(handler.HandlerMiddleware)
	e.POST(reqPath, handler.CreateInstance)
	e.HandleContext(c)
	if resp.Code != http.StatusOK {
		t.Errorf("test failed: %s", resp.Body.String())
	} else {
		t.Logf("test succeed, body :[%s], code : [%d]", resp.Body.String(), resp.Code)
	}
}

func TestUpdateInstance(t *testing.T) {
	targetUrl := fmt.Sprintf("http://%s%s", url, reqPath)
	reqStr, err := json.Marshal(request)
	if err != nil {
		t.Errorf("Failed to marshal database: %v", err)
	}

	body := bytes.NewReader(reqStr)
	req, err := http.NewRequest("POST", targetUrl, body)
	if err != nil {
		t.Fatalf("Failed to create request, %v", err)
	}

	c, e, resp := ginTest(req)
	e.Use(handler.HandlerMiddleware)
	e.PUT(reqPath, handler.UpdateInstance)
	e.HandleContext(c)
	if resp.Code != http.StatusOK {
		t.Errorf("test failed: %s", resp.Body.String())
	} else {
		t.Logf("test succeed, body :[%s], code : [%d]", resp.Body.String(), resp.Code)
	}
}

func TestGetInstance(t *testing.T) {
	targetUrl := fmt.Sprintf("http://%s%s?jobID=%s", url, reqPath, jobID)
	body := strings.NewReader(`{}`)
	req, err := http.NewRequest("GET", targetUrl, body)
	if err != nil {
		t.Fatalf("Failed to create request, %v", err)
	}

	c, e, resp := ginTest(req)
	e.GET(reqPath, handler.GetInstance)
	e.HandleContext(c)
	if resp.Code != http.StatusOK {
		t.Errorf("test failed: %s", resp.Body.String())
	} else {
		t.Logf("test succeed, body :[%s], code : [%d]", resp.Body.String(), resp.Code)
	}
}

func TestDeleteInstance(t *testing.T) {
	targetUrl := fmt.Sprintf("http://%s%s", url, reqPath)
	reqStr, err := json.Marshal(delReq)
	if err != nil {
		t.Errorf("Failed to marshal database: %v", err)
	}

	body := bytes.NewReader(reqStr)
	req, err := http.NewRequest("DELETE", targetUrl, body)
	if err != nil {
		t.Fatalf("Failed to create request, %v", err)
	}

	c, e, resp := ginTest(req)
	e.DELETE(reqPath, handler.DeleteInstance)
	e.HandleContext(c)
	if resp.Code != http.StatusOK {
		t.Errorf("test failed: %s", resp.Body.String())
	} else {
		t.Logf("test succeed, body :[%s], code : [%d]", resp.Body.String(), resp.Code)
	}
}
