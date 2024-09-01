package model

import "time"

type Order struct {
	ID         int64      `gorm:"column:id;primaryKey" json:"id"`
	UserID     int64      `gorm:"column:userId;comment:用户id" json:"userId"`
	OrderID    int64      `gorm:"column:orderId;comment:订单id" json:"orderId"`
	Subject    string     `gorm:"column:subject;comment:订单描述" json:"subject"`
	Amount     float64    `gorm:"column:amount;default:0.00;comment:支付金额" json:"amount"`
	MonthCount int32      `gorm:"column:monthCount;comment:购买时长，按月" json:"monthCount"`
	StartDate  string     `gorm:"column:startDate;comment:开始日期" json:"startDate"`
	EndDate    string     `gorm:"column:endDate;comment:结束日期" json:"endDate"`
	PayStatus  string     `gorm:"column:payStatus;comment:支付状态:待支付(pending),支付完成(paid),支付失败(fail)" json:"payStatus"`
	PaidTime   *time.Time `gorm:"column:paidTime;comment:支付完成时间" json:"paidTime"`
	FailTime   *time.Time `gorm:"column:failTime;comment:支付失败时间" json:"failTime"`
	CancelTime *time.Time `gorm:"column:cancelTime;comment:订单取消时间" json:"cancelTime"`
	Status     int8       `gorm:"column:status;default:1;comment:订单状态：1、正常；2、取消" json:"status"`
	CreatedAt  time.Time  `gorm:"column:createdAt;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"`
}

func (Order) TableName() string {
	return "order"
}
