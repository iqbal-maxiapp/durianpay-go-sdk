package durianpay

import (
	"github.com/iqbal-maxiapp/durianpay-go-sdk/common"

	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/disbursement"
	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/order"
	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/payment"
	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/promo"
	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/refund"
	"github.com/iqbal-maxiapp/durianpay-go-sdk/modules/settlement"
)

type Modules struct {
	Disbursement disbursement.Disbursement
	Order        order.Order
	Payment      payment.Payment
	Promo        promo.Promo
	Settlement   settlement.Settlement
	Refund       refund.Refund
}

func NewClient(config *common.ClientConfig) *Modules {
	agent := common.NewAgent(config)

	return &Modules{
		Disbursement: disbursement.Disbursement{Agent: agent},
		Order:        order.Order{Agent: agent},
		Payment:      payment.Payment{Agent: agent},
		Promo:        promo.Promo{Agent: agent},
		Settlement:   settlement.Settlement{Agent: agent},
		Refund:       refund.Refund{Agent: agent},
	}
}
