package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"golang-senior-learn/chapter3/mall/order/internal/logic"
	"golang-senior-learn/chapter3/mall/order/internal/svc"
	"golang-senior-learn/chapter3/mall/order/internal/types"
)

func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
