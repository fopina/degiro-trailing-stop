package degiro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	Username   string
	Password   string
	URL        string
	httpClient *http.Client
}

func NewAPIClient(username, password string) *APIClient {
	return NewAPIClientWithURL(username, password, "https://trader.degiro.nl/")
}

func NewAPIClientWithURL(username, password, url string) *APIClient {
	c := APIClient{Username: username, Password: password, URL: url}
	c.httpClient = &http.Client{}
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
	fmt.Println(string(jsonValue))
	resp, err := c.httpClient.Post(c.URL+"login/secure/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
