package promo

import (
	"context"

	"github.com/iqbal-maxiapp/durianpay-go-sdk/common"
)

const (
	ROUTE_GROUP string = "/merchants/promos"
)

type Promo struct {
	*common.Agent
}

func (p *Promo) CreatePromo(ctx context.Context, req *CreatePromoRequestPayload) (res *CreatePromoResponse, err error) {
	res = &CreatePromoResponse{}
	err = p.Agent.Call(ctx, "POST", ROUTE_GROUP, req, res)
	return
}

func (p *Promo) FetchPromo(ctx context.Context) (res *FetchPromoResponse, err error) {
	res = &FetchPromoResponse{}
	err = p.Agent.Call(ctx, "GET", ROUTE_GROUP, nil, res)
	return
}

func (p *Promo) FetchPromoById(ctx context.Context, promoId string) (res *FetchPromoByIdResponse, err error) {
	res = &FetchPromoByIdResponse{}
	err = p.Agent.Call(ctx, "GET", ROUTE_GROUP+"/"+promoId, nil, res)
	return
}

func (p *Promo) UpdatePromo(ctx context.Context, promoId string, req *UpdatePromoRequestPayload) (res *UpdatePromoResponse, err error) {
	res = &UpdatePromoResponse{}
	err = p.Agent.Call(ctx, "PATCH", ROUTE_GROUP+"/"+promoId, req, res)
	return
}

func (p *Promo) DeletePromo(ctx context.Context, promoId string) (res *DeletePromoResponse, err error) {
	res = &DeletePromoResponse{}
	err = p.Agent.Call(ctx, "DELETE", ROUTE_GROUP+"/"+promoId, nil, res)
	return
}
