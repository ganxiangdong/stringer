package main

// TestOrderStatus 订单状态
//
//go:generate mystringer -type=TestOrderStatus -linecomment
type TestOrderStatus int

const (
	TestOrderStatusPaid   TestOrderStatus = iota + 1 // 已支付
	TestOrderStatusRefund                            // 已退款
	TestOrderStatusCancel                            // 已取消
)
