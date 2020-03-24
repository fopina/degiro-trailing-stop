package degiro

type Config struct {
	ClientID       int    `json:"clientId"`
	SessionID      string `json:"sessionId"`
	TradingURL     string `json:"tradingUrl"`
	PaURL          string `json:"paUrl"`
	TaskManagerURL string `json:"taskManagerUrl"`
}

type ClientInfo struct {
	IntAccount int `json:"intAccount"`
}

type ClientInfoResponse struct {
	Data ClientInfo `json:"data"`
}

type CurrencyPair struct {
	Id    int    `json:"id"`
	Price string `json:"price"`
}

type CashFund struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ProductIds []int  `json:"productIds"`
}

type AccountInfo struct {
	ClientID            int                     `json:"clientId"`
	BaseCurrency        string                  `json:"baseCurrency"`
	CurrencyPairs       map[string]CurrencyPair `json:"currencyPairs"`
	Margintype          string                  `json:"marginType"`
	CashFunds           map[string][]CashFund   `json:"cashFunds"`
	CompensationCapping float64                 `json:"compensationCapping"`
}

type AccountInfoResponse struct {
	Data AccountInfo `json:"data"`
}

type TransactionFee struct {
	Id       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type CheckOrderInfo struct {
	ConfirmationID  string           `json:"confirmationId"`
	TransactionFees []TransactionFee `json:"transactionFees"`
}

type CheckOrderResponse struct {
	Data CheckOrderInfo `json:"data"`
}

type CreateOrderInfo struct {
	OrderID string `json:"orderId"`
}

type CreateOrderResponse struct {
	Data CreateOrderInfo `json:"data"`
}

type OrderAction string
type OrderType int
type OrderTime int

const (
	OrderActionBuy     OrderAction = "BUY"
	OrderActionSell    OrderAction = "SELL"
	OrderTypeLimit     OrderType   = 0
	OrderTypeStopLimit OrderType   = 1
	OrderTypeMarket    OrderType   = 2
	OrderTypeStoploss  OrderType   = 3
	OrderTimeDay       OrderTime   = 1
	OrderTimeGTC       OrderTime   = 3
)

type Order struct {
	BuySell   OrderAction `json:"buySell"`
	OrderType OrderType   `json:"orderType"`
	ProductID string      `json:"productId"`
	TimeType  OrderTime   `json:"timeType"`
	Size      int         `json:"size"`
	Price     int         `json:"price"`
	StopPrice int         `json:"stopPrice,omitempty"`
}
