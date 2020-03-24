package degiro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type APIClient struct {
	Username   string
	Password   string
	URL        string
	httpClient *http.Client
	config     Config
}

func NewAPIClient(username, password string) *APIClient {
	return NewAPIClientWithURL(username, password, "https://trader.degiro.nl/")
}

func NewAPIClientWithURL(username, password, url string) *APIClient {
	c := APIClient{Username: username, Password: password, URL: url}
	jar, _ := cookiejar.New(nil)
	c.httpClient = &http.Client{Jar: jar, Timeout: 10 * time.Second}
	return &c
}

func (c *APIClient) Login() error {
	req := map[string]string{
		"username":           c.Username,
		"password":           c.Password,
		"isPassCodeReset":    "false",
		"isRedirectToMobile": "false",
	}
	jsonValue, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Post(c.URL+"login/secure/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status %v - %v", resp.Status, string(body))
	}
	return nil
}

func (c *APIClient) GetConfig() error {
	resp, err := c.httpClient.Get(c.URL + "login/secure/config")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status %v - %v", resp.Status, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&c.config)
	if err != nil {
		return err
	}
	return nil
}

func (c *APIClient) GetClientInfo() (*ClientInfo, error) {
	resp, err := c.httpClient.Get(c.config.PaURL + "client?sessionId=" + c.config.SessionID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %v - %v", resp.Status, string(body))
	}

	var info ClientInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	return &info.Data, nil
}

func (c *APIClient) GetAccountInfo(intAccount int) (*AccountInfo, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(
		"%sv5/account/info/%d;jsessionid=%s",
		c.config.TradingURL, intAccount,
		c.config.SessionID,
	))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %v - %v", resp.Status, string(body))
	}

	var info AccountInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	return &info.Data, nil
}
