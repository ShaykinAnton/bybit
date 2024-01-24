package bybit

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

// SubscribeExecution :
func (s *V5WebsocketPrivateService) SubscribeExecution(
	f func(V5WebsocketPrivateExecutionResponse) error,
) (func() error, error) {
	key := V5WebsocketPrivateParamKey{
		Topic: V5WebsocketPrivateTopicExecution,
	}
	if err := s.addParamExecutionFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{V5WebsocketPrivateTopicExecution},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return nil, err
	}
	return func() error {
		param := struct {
			Op   string        `json:"op"`
			Args []interface{} `json:"args"`
		}{
			Op:   "unsubscribe",
			Args: []interface{}{V5WebsocketPrivateTopicExecution},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamExecutionFunc(key)
		return nil
	}, nil
}

// V5WebsocketPrivateExecutionResponse :
type V5WebsocketPrivateExecutionResponse struct {
	ID           string                            `json:"id"`
	Topic        V5WebsocketPrivateTopic           `json:"topic"`
	CreationTime int64                             `json:"creationTime"`
	Data         []V5WebsocketPrivateExecutionData `json:"data"`
}

// V5WebsocketPrivateExecutionData :
type V5WebsocketPrivateExecutionData struct {
	Symbol          SymbolV5   `json:"symbol"`
	OrderID         string     `json:"orderId"`
	OrderLinkID     string     `json:"orderLinkId"`
	Side            Side       `json:"side"`
	OrderPrice      string     `json:"orderPrice"`
	OrderQty        string     `json:"orderQty"`
	LeavesQty       string     `json:"leavesQty"`
	OrderType       OrderType  `json:"orderType"`
	StopOrderType   string     `json:"stopOrderType"`
	ExecFee         string     `json:"execFee"`
	ExecID          string     `json:"execId"`
	ExecPrice       string     `json:"execPrice"`
	ExecQty         string     `json:"execQty"`
	ExecType        ExecTypeV5 `json:"execType"`
	ExecValue       string     `json:"execValue"`
	ExecTime        string     `json:"execTime"`
	IsMaker         bool       `json:"isMaker"`
	FeeRate         string     `json:"feeRate"`
	TradeIv         string     `json:"tradeIv"`
	MarkIv          string     `json:"markIv"`
	MarkPrice       string     `json:"markPrice"`
	IndexPrice      string     `json:"indexPrice"`
	UnderlyingPrice string     `json:"underlyingPrice"`
	BlockTradeID    string     `json:"blockTradeId"`
	ClosedSize      string     `json:"closedSize"`
	Seq             int64      `json:"seq"`
}

// Key :
func (r *V5WebsocketPrivateExecutionResponse) Key() V5WebsocketPrivateParamKey {
	return V5WebsocketPrivateParamKey{
		Topic: r.Topic,
	}
}

// addParamExecutionFunc :
func (s *V5WebsocketPrivateService) addParamExecutionFunc(param V5WebsocketPrivateParamKey, f func(V5WebsocketPrivateExecutionResponse) error) error {
	if _, exist := s.paramExecutionMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramExecutionMap[param] = f
	return nil
}

// removeParamExecutionFunc :
func (s *V5WebsocketPrivateService) removeParamExecutionFunc(key V5WebsocketPrivateParamKey) {
	delete(s.paramExecutionMap, key)
}

// retrieveExecutionFunc :
func (s *V5WebsocketPrivateService) retrieveExecutionFunc(key V5WebsocketPrivateParamKey) (func(V5WebsocketPrivateExecutionResponse) error, error) {
	f, exist := s.paramExecutionMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
