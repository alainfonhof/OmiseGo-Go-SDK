package omisego_test

import (
	omg "github.com/Alainy/OmiseGo-Go-SDK"
	"github.com/icrowley/fake"
	"net/http"
	"net/url"
	"testing"
)

var (
	ewalletURL = &url.URL{
		Scheme: "http",
		Host:   "localhost:4000",
		Path:   "/api",
	}
	sa = &omg.ServerAuth{
		AccessKey: "68HazVGtFNAw4rbSa7k2oz3UvSWOG6MCydXyuPmYoqg",
		SecretKey: "AVqnuxAlbOtpPIer89BjPCLTMQh_PY8g0wd_Dxd-pGU",
	}
	serverUser = omg.EWalletAPI{
		Client: &omg.Client{
			Auth:       sa,
			HttpClient: &http.Client{},
			BaseURL:    ewalletURL,
		},
	}
)

func TestUserCreate(t *testing.T) {
	body := omg.UserParams{
		ProviderUserId: fake.CharactersN(10),
		Username:       fake.UserName(),
		Metadata: map[string]interface{}{
			"first_name": fake.FirstName(),
			"last_name":  fake.LastName(),
		},
	}

	_, err := serverUser.UserCreate(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserGet(t *testing.T) {
	body := omg.ProviderUserIdParam{
		ProviderUserId: "7x1hsxeryf",
	}

	_, err := serverUser.UserGet(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserListBalances(t *testing.T) {
	body := omg.ProviderUserIdParam{
		ProviderUserId: "7x1hsxeryf",
	}

	_, err := serverUser.UserListBalances(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserCreditBalance(t *testing.T) {
	body := omg.BalanceAdjustmentParams{
		ProviderUserId: "7x1hsxeryf",
		TokenId:        "BTC",
		Amount:         100,
	}

	_, err := serverUser.UserCreditBalance(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDebitBalance(t *testing.T) {
	body := omg.BalanceAdjustmentParams{
		ProviderUserId: "7x1hsxeryf",
		TokenId:        "BTC",
		Amount:         100,
	}

	_, err := serverUser.UserDebitBalance(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionAll(t *testing.T) {
	body := omg.ListParams{
		Page:    1,
		PerPage: 10,
	}

	res, err := serverUser.TransactionAll(body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestUserListTransactions(t *testing.T) {
	body := omg.UserListTransactionsParams{
		ProviderUserId: "7x1hsxeryf",
		Page:           1,
		PerPage:        10,
	}

	_, err := serverUser.UserListTransactions(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionRequestCreate(t *testing.T) {
	body := omg.ServerCreateTransactionRequestParams{
		Type:          "send",
		TokenId:       "tok_OMG_01cbffwvj6ma9a9gg1tb24880q",
		Amount:        100,
		CorrelationId: "123",
		Address:       "2ae52683-68d8-4af6-94d7-5ed4c34ecf1a",
	}

	_, err := serverUser.TransactionRequestCreate(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionRequestGet(t *testing.T) {
	body := omg.ByIdParam{
		Id: "txr_01cbfgc8cmmyzy1cfzpqwme3ey",
	}

	_, err := serverUser.TransactionRequestGet(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionRequestConsume(t *testing.T) {
	body := omg.ServerTransactionRequestConsumeParams{
		TransactionRequestId: "txr_01cbfgcts5kqfgpqxcxn71rnbs",
		CorrelationId:        "123",
		TokenId:              "tok_OMG_01cbffwvj6ma9a9gg1tb24880q",
		Amount:               100,
		Address:              "2ae52683-68d8-4af6-94d7-5ed4c34ecf1a",
		ProviderUserId:       "7x1hsxeryf",
	}

	_, err := serverUser.TransactionRequestConsume(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionConsumptionApprove(t *testing.T) {
	body := omg.ByIdParam{
		Id: "txn_01cbfg5m2ee06kzm8tbysfmmw5",
	}

	_, err := serverUser.TransactionConsumptionApprove(body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionConsumptionReject(t *testing.T) {
	body := omg.ByIdParam{
		Id: "txn_01cbfg5m2ee06kzm8tbysfmmw5",
	}

	_, err := serverUser.TransactionConsumptionReject(body)
	if err != nil {
		t.Fatal(err)
	}
}
