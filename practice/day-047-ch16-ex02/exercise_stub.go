package main

type OrderInfo struct {
	OrderCode   rune
	Amount      int
	OrderNumber uint16
	Items       []string
	IsReady     bool
}

type SmallOrderInfo struct {
	IsReady     bool
	OrderNumber uint16
	OrderCode   rune
	Amount      int
	Items       []string
}
