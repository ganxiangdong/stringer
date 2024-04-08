Copy from: golang.org/x/tools/cmd/stringer.

This tool provides the same functionality as the official one, with the addition of some IsXXX methods.

### 1. Install
```
go install github.com/ganxiangdong/stringer@main 
```

### 2. Writing generate annotation
```bash
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
```

### 3. Run stringer
```bash
go generate {your directory}
```

### 4. Generate the following code
```go
// Code generated by "stringer -type=TestOrderStatus -linecomment"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TestOrderStatusPaid-1]
	_ = x[TestOrderStatusRefund-2]
	_ = x[TestOrderStatusCancel-3]
}

const _TestOrderStatus_name = "已支付已退款已取消"

var _TestOrderStatus_index = [...]uint8{0, 9, 18, 27}

func (i TestOrderStatus) String() string {
	i -= 1
	if i < 0 || i >= TestOrderStatus(len(_TestOrderStatus_index)-1) {
		return "TestOrderStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TestOrderStatus_name[_TestOrderStatus_index[i]:_TestOrderStatus_index[i+1]]
}

// IsPaid 是否是：已支付
func (i TestOrderStatus) IsPaid() bool {
	return i == TestOrderStatusPaid
}

// IsRefund 是否是：已退款
func (i TestOrderStatus) IsRefund() bool {
	return i == TestOrderStatusRefund
}

// IsCancel 是否是：已取消
func (i TestOrderStatus) IsCancel() bool {
	return i == TestOrderStatusCancel
}

// GetValue 获取原始类型值
func (i TestOrderStatus) GetValue() int {
	return int(i)
}
```
