package omisego

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type EWalletAPI struct {
	*Client
}

/////////////////
// Session
/////////////////
func (e *EWalletAPI) Login(reqBody ProviderUserIdParam) (*AuthenicationToken, error) {
	req, err := e.newRequest("POST", "/login", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data AuthenicationToken
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (e *EWalletAPI) Logout() error {
	req, err := e.newRequest("POST", "/logout", nil)
	if err != nil {
		return err
	}

	_, err = e.do(req)
	return err
}

/////////////////
// User
/////////////////
func (e *EWalletAPI) UserCreate(reqBody UserParams) (*User, error) {
	req, err := e.newRequest("POST", "/user.create", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) UserUpdate(reqBody UserParams) (*User, error) {
	req, err := e.newRequest("POST", "/user.update", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) UserGet(reqBody ProviderUserIdParam) (*User, error) {
	req, err := e.newRequest("POST", "/user.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) MeGet() (*User, error) {
	req, err := e.newRequest("POST", "/me.get", nil)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Balance
/////////////////
func (e *EWalletAPI) UserListBalances(reqBody ProviderUserIdParam) (*AddressList, error) {
	req, err := e.newRequest("POST", "/user.list_balances", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data AddressList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) UserCreditBalance(reqBody BalanceAdjustmentParams) (*AddressList, error) {
	req, err := e.newRequest("POST", "/user.credit_balance", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data AddressList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) UserDebitBalance(reqBody BalanceAdjustmentParams) (*AddressList, error) {
	req, err := e.newRequest("POST", "/user.debit_balance", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data AddressList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) Transfer(reqBody TransferParams) (*AddressList, error) {
	req, err := e.newRequest("POST", "/transfer", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data AddressList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Settings
/////////////////
func (e *EWalletAPI) GetSettings() (*Settings, error) {
	req, err := e.newRequest("POST", "/get_settings", nil)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data Settings
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Transaction
/////////////////
func (e *EWalletAPI) TransactionAll(reqBody ListParams) (*TransactionList, error) {
	req, err := e.newRequest("POST", "/transaction.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) UserListTransactions(reqBody UserListTransactionsParams) (*TransactionList, error) {
	req, err := e.newRequest("POST", "/user.list_transactions", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Transaction Request
/////////////////
func (e *EWalletAPI) TransactionRequestCreate(reqBody ServerCreateTransactionRequestParams) (*TransactionRequest, error) {
	req, err := e.newRequest("POST", "/transaction_request.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequest
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) TransactionRequestGet(reqBody ByIdParam) (*TransactionRequest, error) {
	req, err := e.newRequest("POST", "/transaction_request.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequest
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) TransactionRequestConsume(reqBody ServerTransactionRequestConsumeParams) (*TransactionComsumption, error) {
	req, err := e.newRequest("POST", "/transaction_request.consume", reqBody)
	req.Header.Set("Idempotency-Token", NewIdempotencyToken())
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Transaction Consumption
/////////////////
func (e *EWalletAPI) TransactionConsumptionApprove(reqBody ByIdParam) (*TransactionComsumption, error) {
	req, err := e.newRequest("POST", "/transaction_consumption.approve", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *EWalletAPI) TransactionConsumptionReject(reqBody ByIdParam) (*TransactionComsumption, error) {
	req, err := e.newRequest("POST", "/transaction_consumption.reject", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}
