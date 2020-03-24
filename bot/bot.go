package bot

import (
	"fmt"

	"github.com/fopina/degiro-trailing-stop/degiro"
)

type botState struct {
	apiClient   *degiro.APIClient
	clientInfo  *degiro.ClientInfo
	accountInfo *degiro.AccountInfo
}

func NewBot(username, password string) *botState {
	api := degiro.NewAPIClient(username, password)
	b := botState{apiClient: api}
	return &b
}

func (b *botState) Login() error {
	err := b.apiClient.Login()
	if err != nil {
		return err
	}
	err = b.apiClient.GetConfig()
	if err != nil {
		return err
	}
	b.clientInfo, err = b.apiClient.GetClientInfo()
	if err != nil {
		return err
	}
	b.accountInfo, err = b.apiClient.GetAccountInfo(b.clientInfo.IntAccount)
	return err
}

func (b *botState) Test() error {
	o := degiro.Order{
		BuySell: degiro.OrderActionBuy, OrderType: degiro.OrderTypeStopLimit,
		ProductID: "1153605", TimeType: degiro.OrderTimeGTC, Size: 1,
		Price: 441, StopPrice: 440,
	}
	d, err := b.apiClient.CheckOrder(
		b.clientInfo.IntAccount,
		o,
	)
	if err != nil {
		return err
	}
	o2, err := b.apiClient.CreateOrder(d.ConfirmationID, b.clientInfo.IntAccount, o)
	if err != nil {
		return err
	}
	fmt.Println(o)
	err = b.apiClient.DeleteOrder(o2.OrderID, b.clientInfo.IntAccount)
	return err
}
