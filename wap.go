package alipay

import (
	"strings"
	"net/http"
	"net/url"
)

type PayRequest struct {
	URL string
	Form map[string][]string
}

func (this *AliPay) TradeWapPayRequest(param AliPayTradeWapPay) *PayRequest {
	req := &PayRequest{
		URL: this.apiDomain,
		Form: this.URLValues(param),
	}
	return req
}

func (this *AliPay) TradeWapPay(param AliPayTradeWapPay) (url *url.URL, err error) {
	var buf = strings.NewReader(this.URLValues(param).Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	rep, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	if err != nil {
		return nil, err
	}
	url = rep.Request.URL
	return url, err
}