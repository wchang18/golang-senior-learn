package platform_pay

import (
	"fmt"
	"github.com/thoas/go-funk"
	"time"
)

type PlatformServer struct {
}

func NewPlatformServer() *PlatformServer {
	return &PlatformServer{}
}

func (p *PlatformServer) GetPayCode(req *GetPlatformCodeReq) (res *GetPlatformCodeRes, err error) {
	res = &GetPlatformCodeRes{
		TradeNo:  req.TradeNo,
		PayUrl:   p.GetPayUrl(req.PlatformId),
		ExpireAt: time.Now().Add(time.Hour).Format("2006-01-02 15:04:05"),
	}
	return
}

func (p *PlatformServer) GetPayUrl(platformId int32) (url string) {
	no := funk.RandomString(12)
	if platformId == 1 {
		url += fmt.Sprintf("https://qr.alipay.com/%s", no)
	} else if platformId == 2 {
		url += fmt.Sprintf("weixin://wxpay/bizpayurl?pr=%s", no)
	}
	return
}

type GetPlatformCodeReq struct {
	PlatformId int32   `json:"platformId"`
	Amount     float32 `json:"amount"`
	Subject    string  `json:"subject"`
	TradeNo    string  `json:"tradeNo"`
}

type GetPlatformCodeRes struct {
	TradeNo  string `json:"tradeNo"`
	PayUrl   string `json:"payUrl"`
	ExpireAt string `json:"expireAt"`
}
