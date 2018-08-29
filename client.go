package cloudloyalty_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ContextValueXProcessingKey = "X-Processing-Key"
)

const contentType = "application/json"

type Config struct {
	HttpClient    *http.Client
	BaseURL       string
	ProcessingKey string
}

type ProcessingError struct {
	error
	Code int
}

type Client struct {
	httpClient    *http.Client
	baseURL       string
	processingKey string
}

func New(config *Config) *Client {
	httpClient := config.HttpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		httpClient:    httpClient,
		baseURL:       strings.TrimRight(config.BaseURL, "/"),
		processingKey: config.ProcessingKey,
	}
}

func (c *Client) request(ctx context.Context, path string, req interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", c.baseURL+path, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	httpReq.Header.Set("X-Processing-Key", eitherKeyValue(ctx, c.processingKey))
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error code=%d", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var serverError errorReply
	if err = json.Unmarshal(respBody, &serverError); err != nil {
		return nil, err
	}
	if serverError.ErrorCode > 0 {
		return nil, &ProcessingError{
			error: errors.New(serverError.Description),
			Code:  serverError.ErrorCode,
		}
	}
	return respBody, nil
}

func (c *Client) GetBalance(ctx context.Context, req *GetBalanceQuery) (*GetBalanceReply, error) {
	respBody, err := c.request(ctx, "/get-balance", req)
	if err != nil {
		return nil, err
	}
	var resp GetBalanceReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CalculatePurchase(ctx context.Context, req *CalculatePurchaseQuery) (*CalculatePurchaseReply, error) {
	respBody, err := c.request(ctx, "/calculate-purchase", req)
	if err != nil {
		return nil, err
	}
	var resp CalculatePurchaseReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ApplyPurchase(ctx context.Context, req *ApplyPurchaseQuery) (*ApplyPurchaseReply, error) {
	respBody, err := c.request(ctx, "/apply-purchase", req)
	if err != nil {
		return nil, err
	}
	var resp ApplyPurchaseReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ApplyReturn(ctx context.Context, req *ApplyReturnQuery) (*ApplyReturnReply, error) {
	respBody, err := c.request(ctx, "/apply-return", req)
	if err != nil {
		return nil, err
	}
	var resp ApplyReturnReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func eitherKeyValue(ctx context.Context, key string) string {
	if ctx != nil {
		if value, ok := ctx.Value(ContextValueXProcessingKey).(string); ok {
			return value
		}
	}
	return key
}
