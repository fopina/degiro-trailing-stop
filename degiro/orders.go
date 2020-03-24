package degiro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *APIClient) orderURL(path string, intAccount int) string {
	return fmt.Sprintf(
		"%s%s;jsessionid=%s?intAccount=%d&sessionId=%s",
		c.config.TradingURL, path, c.config.SessionID,
		intAccount, c.config.SessionID,
	)
}

func (c *APIClient) CheckOrder(intAccount int, update Order) (*CheckOrderInfo, error) {
	jsonData, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.orderURL("v5/checkOrder", intAccount),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %v - %v", resp.Status, string(body))
	}

	var info CheckOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	return &info.Data, nil
}

func (c *APIClient) CreateOrder(order string, intAccount int, update Order) (*CreateOrderInfo, error) {
	jsonData, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.orderURL("v5/order/"+order, intAccount),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %v - %v", resp.Status, string(body))
	}

	var info CreateOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	return &info.Data, nil
}

func (c *APIClient) UpdateOrder(order string, intAccount int, update Order) error {
	json, err := json.Marshal(update)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		c.orderURL("v5/order/"+order, intAccount),
		bytes.NewBuffer(json),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status %v - %v", resp.Status, string(body))
	}
	return nil
}

func (c *APIClient) DeleteOrder(order string, intAccount int) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		c.orderURL("v5/order/"+order, intAccount),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status %v - %v", resp.Status, string(body))
	}
	return nil
}
