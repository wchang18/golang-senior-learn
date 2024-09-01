package model

import "time"

type OrderPay struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID    int64      `gorm:"column:orderId;comment:订单id" json:"orderId"`
	PlatformID int32      `gorm:"column:platformId;comment:支付方式：1支付宝，2微信" json:"platformId"`
	PlatformNo string     `gorm:"column:platformNo;comment:第三方交易号" json:"platformNo"`
	TradeNo    string     `gorm:"column:tradeNo;comment:交易号" json:"tradeNo"`
	ExpireAt   string     `gorm:"column:expireAt;comment:过期时间" json:"expireAt"`
	PayQrCode  string     `gorm:"column:payQrCode;comment:二维码" json:"payQrCode"`
	PayStatus  int8       `gorm:"column:payStatus;comment:支付状态:支付状态:1:待支付 2:交易成功 3:交易关闭" json:"payStatus"`
	PaidTime   *time.Time `gorm:"column:paidTime;comment:支付完成时间" json:"paidTime"`
	FailTime   *time.Time `gorm:"column:failTime;comment:支付失败时间" json:"failTime"`
	Status     int8       `gorm:"column:status;default:1;comment:订单状态：1、正常；2、取消" json:"status"`
	CreatedAt  time.Time  `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
}

func (OrderPay) TableName() string {
	return "order_pay"
}
