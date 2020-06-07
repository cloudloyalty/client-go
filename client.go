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

const ContextValueXProcessingKey = "X-Processing-Key"

const contentType = "application/json"

type Config struct {
	HttpClient    *http.Client
	BaseURL       string
	ProcessingKey string
	OnQuery       func(req *http.Request, body []byte)
	OnReply       func(res *http.Response, body []byte)
}

type ProcessingError struct {
	error
	Code int
}

type ClientInterface interface {
	GetBalance(ctx context.Context, req *GetBalanceQuery) (*GetBalanceReply, error)
	NewClient(ctx context.Context, req *NewClientQuery) (*NewClientReply, error)
	UpdateClient(ctx context.Context, req *UpdateClientQuery) (*UpdateClientReply, error)
	CalculatePurchase(ctx context.Context, req *CalculatePurchaseQuery) (*CalculatePurchaseReply, error) // deprecated
	ApplyPurchase(ctx context.Context, req *ApplyPurchaseQuery) (*ApplyPurchaseReply, error)             // deprecated
	CalculateReturn(ctx context.Context, req *CalculateReturnQuery) (*CalculateReturnReply, error)
	ApplyReturn(ctx context.Context, req *ApplyReturnQuery) (*ApplyReturnReply, error)
	SetOrder(ctx context.Context, req *SetOrderQuery) (*SetOrderReply, error)
	ConfirmOrder(ctx context.Context, req *ConfirmOrderQuery) (*ConfirmOrderReply, error)
	CancelOrder(ctx context.Context, req *CancelOrderQuery) (*CancelOrderReply, error)
	AdjustBalance(ctx context.Context, req *AdjustBalanceQuery) (*AdjustBalanceReply, error)
	SendConfirmationCode(ctx context.Context, req *SendConfirmationCodeQuery) (*SendConfirmationCodeReply, error)
	GetHistory(ctx context.Context, req *GetHistoryQuery) (*GetHistoryReply, error)
	IssuePromocode(ctx context.Context, req *IssuePromocodeQuery) (*IssuePromocodeReply, error)
	RevertPurchase(ctx context.Context, req *RevertPurchaseQuery) (*RevertPurchaseReply, error)
	V2CalculatePurchase(ctx context.Context, req *V2CalculatePurchaseRequest) (*V2CalculatePurchaseReply, error)
	SetPurchase(ctx context.Context, req *SetPurchaseRequest) (*SetPurchaseReply, error)
	ConfirmTicket(ctx context.Context, req *ConfirmTicketRequest) (*ConfirmTicketReply, error)
	DiscardTicket(ctx context.Context, req *DiscardTicketRequest) (*DiscardTicketReply, error)
}

type Client struct {
	ClientInterface

	httpClient    *http.Client
	baseURL       string
	processingKey string
	onQuery       func(req *http.Request, body []byte)
	onReply       func(res *http.Response, body []byte)
}

func New(config *Config) ClientInterface {
	httpClient := config.HttpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		httpClient:    httpClient,
		baseURL:       strings.TrimRight(config.BaseURL, "/"),
		processingKey: config.ProcessingKey,
		onQuery:       config.OnQuery,
		onReply:       config.OnReply,
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

	if c.onQuery != nil {
		c.onQuery(httpReq, reqBody)
	}

	resp, err := c.httpClient.Do(httpReq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.onReply != nil {
		c.onReply(resp, respBody)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error code=%d", resp.StatusCode)
	}

	var serverError ErrorReply
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

func (c *Client) NewClient(ctx context.Context, req *NewClientQuery) (*NewClientReply, error) {
	respBody, err := c.request(ctx, "/new-client", req)
	if err != nil {
		return nil, err
	}
	var resp NewClientReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateClient(ctx context.Context, req *UpdateClientQuery) (*UpdateClientReply, error) {
	respBody, err := c.request(ctx, "/update-client", req)
	if err != nil {
		return nil, err
	}
	var resp UpdateClientReply
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

func (c *Client) CalculateReturn(ctx context.Context, req *CalculateReturnQuery) (*CalculateReturnReply, error) {
	respBody, err := c.request(ctx, "/calculate-return", req)
	if err != nil {
		return nil, err
	}
	var resp CalculateReturnReply
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

func (c *Client) SetOrder(ctx context.Context, req *SetOrderQuery) (*SetOrderReply, error) {
	respBody, err := c.request(ctx, "/set-order", req)
	if err != nil {
		return nil, err
	}
	var resp SetOrderReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ConfirmOrder(ctx context.Context, req *ConfirmOrderQuery) (*ConfirmOrderReply, error) {
	respBody, err := c.request(ctx, "/confirm-order", req)
	if err != nil {
		return nil, err
	}
	var resp ConfirmOrderReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderQuery) (*CancelOrderReply, error) {
	respBody, err := c.request(ctx, "/cancel-order", req)
	if err != nil {
		return nil, err
	}
	var resp CancelOrderReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) AdjustBalance(ctx context.Context, req *AdjustBalanceQuery) (*AdjustBalanceReply, error) {
	respBody, err := c.request(ctx, "/adjust-balance", req)
	if err != nil {
		return nil, err
	}
	var resp AdjustBalanceReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SendConfirmationCode(ctx context.Context, req *SendConfirmationCodeQuery) (*SendConfirmationCodeReply, error) {
	respBody, err := c.request(ctx, "/send-confirmation-code", req)
	if err != nil {
		return nil, err
	}
	var resp SendConfirmationCodeReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetHistory(ctx context.Context, req *GetHistoryQuery) (*GetHistoryReply, error) {
	respBody, err := c.request(ctx, "/get-history", req)
	if err != nil {
		return nil, err
	}
	var resp GetHistoryReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) IssuePromocode(ctx context.Context, req *IssuePromocodeQuery) (*IssuePromocodeReply, error) {
	respBody, err := c.request(ctx, "/issue-promocode", req)
	if err != nil {
		return nil, err
	}
	var resp IssuePromocodeReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RevertPurchase(ctx context.Context, req *RevertPurchaseQuery) (*RevertPurchaseReply, error) {
	respBody, err := c.request(ctx, "/revert-purchase", req)
	if err != nil {
		return nil, err
	}
	var resp RevertPurchaseReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) V2CalculatePurchase(ctx context.Context, req *V2CalculatePurchaseRequest) (*V2CalculatePurchaseReply, error) {
	respBody, err := c.request(ctx, "/v2/calculate-purchase", req)
	if err != nil {
		return nil, err
	}
	var resp V2CalculatePurchaseReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SetPurchase(ctx context.Context, req *SetPurchaseRequest) (*SetPurchaseReply, error) {
	respBody, err := c.request(ctx, "/set-purchase", req)
	if err != nil {
		return nil, err
	}
	var resp SetPurchaseReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ConfirmTicket(ctx context.Context, req *ConfirmTicketRequest) (*ConfirmTicketReply, error) {
	respBody, err := c.request(ctx, "/confirm-ticket", req)
	if err != nil {
		return nil, err
	}
	var resp ConfirmTicketReply
	if err = json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DiscardTicket(ctx context.Context, req *DiscardTicketRequest) (*DiscardTicketReply, error) {
	respBody, err := c.request(ctx, "/discard-ticket", req)
	if err != nil {
		return nil, err
	}
	var resp DiscardTicketReply
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
