package pay4go

import (
	"net/http"
)

type Service struct {
	channels map[string]PayChannel
}

func NewService() *Service {
	var s = &Service{}
	s.channels = make(map[string]PayChannel)
	return s
}

func (this *Service) RegisterChannel(c PayChannel) {
	if c != nil {
		this.channels[c.Identifier()] = c
	}
}

func (this *Service) RemoveChannel(channel string) {
	delete(this.channels, channel)
}

func (this *Service) CreatePayment(channel string, order *Order) (url string, err error) {
	var p = this.channels[channel]
	if p == nil {
		return "", ErrUnknownChannel
	}
	return p.CreateTradeOrder(order)
}

func (this *Service) GetTrade(channel string, tradeNo string) (result *Trade, err error) {
	var p = this.channels[channel]
	if p == nil {
		return nil, ErrUnknownChannel
	}
	return p.GetTrade(tradeNo)
}

func (this *Service) GetTradeWithOrderNo(channel string, orderNo string) (result *Trade, err error) {
	var p = this.channels[channel]
	if p == nil {
		return nil, ErrUnknownChannel
	}
	return p.GetTradeWithOrderNo(orderNo)
}

func (this *Service) ReturnRequestHandler(req *http.Request) (result *Trade, err error) {
	req.ParseForm()

	var channel = req.FormValue("channel")
	var p = this.channels[channel]
	if p == nil {
		return nil, ErrUnknownChannel
	}
	return p.ReturnRequestHandler(req)
}

func (this *Service) NotifyRequestHandler(req *http.Request) (result *Notification, err error) {
	req.ParseForm()

	var channel = req.FormValue("channel")
	var p = this.channels[channel]
	if p == nil {
		return nil, ErrUnknownChannel
	}
	return p.NotifyRequestHandler(req)
}
