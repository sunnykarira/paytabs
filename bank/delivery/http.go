package delivery

import (
	"context"
	"encoding/json"
	"github.com/paytabs/app"
	"github.com/paytabs/bank/model"
	"github.com/paytabs/bank/usecase"
	"net/http"
	"time"
)

type bankTransferDelivery struct {
	useCase    usecase.UseCase
	httpServer *app.HttpServer
}

func NewHTTPDelivery(useCase usecase.UseCase, httpServer *app.HttpServer) {
	p := &bankTransferDelivery{
		useCase:    useCase,
		httpServer: httpServer,
	}

	InitHTTPEndpoint(p)
}

func InitHTTPEndpoint(delivery *bankTransferDelivery) {
	delivery.httpServer.POST("/api/v1/fetch", delivery.FetchAccountDetails)
	delivery.httpServer.POST("/api/v1/create", delivery.CreateAccount)
	delivery.httpServer.POST("/api/v1/add", delivery.AddMoney)
	delivery.httpServer.POST("/api/v1/send", delivery.SendMoney)
}

func (b *bankTransferDelivery) FetchAccountDetails(ctx *app.Context, request *http.Request, params app.HttpParams) (app.HttpResponseBody, error) {
	var fetchRequest AccountDetailsRequest
	var err error

	err = json.NewDecoder(request.Body).Decode(&fetchRequest)
	if err != nil {
		return json.Marshal(app.BadError("BANK_101_ERROR", "Invalid or Incomplete payload data"))
	}
	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := b.useCase.FetchData(newCtx, fetchRequest.ID)
	if err != nil {
		return json.Marshal(app.BadError("BANK_101_ERROR", err.Error()))
	}
	return json.Marshal(resp)
}

func (b *bankTransferDelivery) CreateAccount(ctx *app.Context, request *http.Request, params app.HttpParams) (app.HttpResponseBody, error) {
	var fetchRequest CreateAccountRequest
	var err error

	err = json.NewDecoder(request.Body).Decode(&fetchRequest)
	if err != nil || fetchRequest.ID <= 0 || fetchRequest.Balance < 0 || fetchRequest.Location == "" || (model.AccountStatus(fetchRequest.AccountStatus) != model.AccountStatusBlocked && model.AccountStatus(fetchRequest.AccountStatus) != model.AccountStatusActive) {
		return json.Marshal(app.BadError("BANK_101_ERROR", "Invalid or Incomplete payload data"))
	}
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := b.useCase.CreateAccount(newCtx, model.Account{
		ID:            fetchRequest.ID,
		Location:      fetchRequest.Location,
		Balance:       fetchRequest.Balance,
		AccountStatus: model.AccountStatus(fetchRequest.AccountStatus),
	})
	if err != nil {
		return  json.Marshal(app.BadError("BANK_101_ERROR", err.Error()))
	}
	if resp{

		return json.Marshal(app.RespMessage("", "account created successfully"))
	}
	return json.Marshal(app.RespMessage("BANK_101_ERROR", "unable to create account"))
}

func (b *bankTransferDelivery) AddMoney(ctx *app.Context, request *http.Request, params app.HttpParams) (app.HttpResponseBody, error) {

	var fetchRequest AddMoneyRequest
	var err error

	err = json.NewDecoder(request.Body).Decode(&fetchRequest)
	if err != nil || fetchRequest.ID <= 0 || fetchRequest.Amount <= 0 {
		return json.Marshal(app.BadError("BANK_101_ERROR", "Invalid or Incomplete payload data"))
	}
	newCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := b.useCase.AddMoney(newCtx, fetchRequest.ID, fetchRequest.Amount)
	if err != nil {
		return json.Marshal(app.BadError("BANK_101_ERROR", err.Error()))
	}
	if resp{

		return json.Marshal(app.RespMessage("", "added money successfully"))
	}
	return json.Marshal(app.RespMessage("BANK_101_ERROR", "unable to add money"))
}

func (b *bankTransferDelivery) SendMoney(ctx *app.Context, request *http.Request, params app.HttpParams) (app.HttpResponseBody, error) {

	var fetchRequest SendMoneyRequest
	var err error

	err = json.NewDecoder(request.Body).Decode(&fetchRequest)
	if err != nil || fetchRequest.DestinationAccountID <= 0 || fetchRequest.SourceAccountID <= 0 || fetchRequest.Amount <= 0 {
		return json.Marshal(app.BadError("BANK_101_ERROR", "Invalid or Incomplete payload data"))
	}
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := b.useCase.SendMoney(newCtx, model.TransactionData{
		SourceAccountID:      fetchRequest.SourceAccountID,
		DestinationAccountID: fetchRequest.DestinationAccountID,
		Amount:               fetchRequest.Amount,
	})
	if err != nil {
		return json.Marshal(app.BadError("BANK_101_ERROR", err.Error()))
	}
	if resp{

		return json.Marshal(app.RespMessage("", "sent money successfully"))
	}
	return json.Marshal(app.RespMessage("BANK_101_ERROR", "unable to send money"))
}
