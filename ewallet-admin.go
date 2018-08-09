package omisego

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type AdminAPI struct {
	*Client
}

/////////////////
// AdminSession - Resources related to admin session tokens.
/////////////////
func (a *AdminAPI) AdminLogin(reqBody AdminLoginParams) (*AuthenicationToken, error) {
	req, err := a.newRequest("POST", "/admin.login", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
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

func (a *AdminAPI) AdminLogout() error {
	req, err := a.newRequest("POST", "/me.logout", nil)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) AuthTokenSwitchAccount(reqBody AuthTokenSwitchAccountParams) (*AuthenicationToken, error) {
	req, err := a.newRequest("POST", "/auth_token.switch_account", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
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

func (a *AdminAPI) PasswordReset(reqBody AdminResetPasswordParams) error {
	req, err := a.newRequest("POST", "/admin.password_reset", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) PasswordUpdate(reqBody AdminUpdatePasswordParams) error {
	req, err := a.newRequest("POST", "/admin.password_update", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

/////////////////
// Admin - Resources related to admin users.
/////////////////
func (a *AdminAPI) AdminAll(reqBody ListParams) (*UserList, error) {
	req, err := a.newRequest("POST", "/admin.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data UserList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AdminGet(reqBody IdParam) (*User, error) {
	req, err := a.newRequest("POST", "/admin.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) InviteAccept() (*User, error) {
	req, err := a.newRequest("POST", "/invite.accept", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MeGet() (*User, error) {
	req, err := a.newRequest("POST", "/me.get", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MeUpdate(reqBody MeUpdateParams) (*User, error) {
	req, err := a.newRequest("POST", "/me.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MeUploadAvatar(reqBody UploadAvatarParams) (*User, error) {
	req, err := a.newRequest("POST", "/me.upload_avatar", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data User
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MeGetAccount() (*Account, error) {
	req, err := a.newRequest("POST", "/me.get_account", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MeGetAccounts() (*AccountList, error) {
	req, err := a.newRequest("POST", "/me.get_accounts", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccountList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// UserSession - Resources related to user session tokens.
/////////////////
func (a *AdminAPI) UserLogin(reqBody IdParam) (*AuthenicationToken, error) {
	req, err := a.newRequest("POST", "/user.login", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
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

func (a *AdminAPI) UserLogout(reqBody AuthTokenParam) error {
	req, err := a.newRequest("POST", "/user.logout", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

/////////////////
// User - Resources related to users.
/////////////////
func (a *AdminAPI) UserAll(reqBody ListParams) (*UserList, error) {
	req, err := a.newRequest("POST", "/user.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data UserList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (e *AdminAPI) UserCreate(reqBody UserParams) (*User, error) {
	req, err := e.newRequest("POST", "/user.create", reqBody)
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

func (e *AdminAPI) UserUpdate(reqBody UserParams) (*User, error) {
	req, err := e.newRequest("POST", "/user.update", reqBody)
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

func (e *AdminAPI) UserGetById(reqBody IdParam) (*User, error) {
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

func (e *AdminAPI) UserGetByProviderUserId(reqBody ProviderUserIdParam) (*User, error) {
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

func (e *AdminAPI) UserGetWalletsById(reqBody ListByIdParams) (*WalletList, error) {
	req, err := e.newRequest("POST", "/user.get_wallets", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data WalletList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *AdminAPI) UserGetWalletsByProviderUserId(reqBody ListByProviderUserIdParams) (*WalletList, error) {
	req, err := e.newRequest("POST", "/user.get_wallets", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data WalletList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (e *AdminAPI) UserGetTransactionsById(reqBody ListByIdParams) (*TransactionList, error) {
	req, err := e.newRequest("POST", "/user.get_transactions", reqBody)
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

func (e *AdminAPI) UserGetTransactionsByProviderUserId(reqBody ListByProviderUserIdParams) (*TransactionList, error) {
	req, err := e.newRequest("POST", "/user.get_transactions", reqBody)
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

func (e *AdminAPI) UserGetTransactionConsumption(reqBody ListByUserIdParams) (*TransactionComsumptionList, error) {
	req, err := e.newRequest("POST", "/user.get_transaction_consumptions", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := e.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumptionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

/////////////////
// Token - Resources related to tokens.
/////////////////
func (a *AdminAPI) TokenAll(reqBody ListParams) (*TokenList, error) {
	req, err := a.newRequest("POST", "/token.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TokenList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenGet(reqBody IdParam) (*Token, error) {
	req, err := a.newRequest("POST", "/token.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Token
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenCreate(reqBody TokenCreateParams) (*Token, error) {
	req, err := a.newRequest("POST", "/token.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Token
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenUpdate(reqBody TokenUpdateParams) (*Token, error) {
	req, err := a.newRequest("POST", "/token.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Token
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenStats(reqBody IdParam) (*TokenStats, error) {
	req, err := a.newRequest("POST", "/token.stats", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TokenStats
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenGetMints(reqBody ListParams) (*MintList, error) {
	req, err := a.newRequest("POST", "/token.get_mints", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MintList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TokenMint(reqBody TokenMintParams) (*Token, error) {
	req, err := a.newRequest("POST", "/token.mint", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Token
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// ExchangePair - Resources related to exchange pairs.
/////////////////
func (a *AdminAPI) ExchangePairAll(reqBody ListParams) (*ExchangePairList, error) {
	req, err := a.newRequest("POST", "/exchange_pair.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data ExchangePairList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) ExchangePairGet(reqBody IdParam) (*ExchangePair, error) {
	req, err := a.newRequest("POST", "/exchange_pair.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data ExchangePair
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) ExchangePairCreate(reqBody ExchangePairCreateParams) (*ExchangePairList, error) {
	req, err := a.newRequest("POST", "/exchange_pair.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data ExchangePairList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) ExchangePairUpdate(reqBody ExchangePairUpdateParams) (*ExchangePairList, error) {
	req, err := a.newRequest("POST", "/exchange_pair.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data ExchangePairList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) ExchangePairDelete(reqBody IdParam) (*ExchangePairList, error) {
	req, err := a.newRequest("POST", "/exchange_pair.delete", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data ExchangePairList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// Category - Resources related to categories.
/////////////////
func (a *AdminAPI) CategoryAll(reqBody ListParams) (*CategoryList, error) {
	req, err := a.newRequest("POST", "/category.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data CategoryList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) CategoryGet(reqBody IdParam) (*Category, error) {
	req, err := a.newRequest("POST", "/category.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Category
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) CategoryCreate(reqBody ExchangePairCreateParams) (*Category, error) {
	req, err := a.newRequest("POST", "/category.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Category
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) CategoryUpdate(reqBody ExchangePairUpdateParams) (*Category, error) {
	req, err := a.newRequest("POST", "/category.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Category
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) CategoryDelete(reqBody IdParam) (*Category, error) {
	req, err := a.newRequest("POST", "/category.delete", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Category
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// Account - Resources related to accounts.
/////////////////
func (a *AdminAPI) AccountAll(reqBody ListParams) (*AccountList, error) {
	req, err := a.newRequest("POST", "/account.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccountList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGet(reqBody IdParam) (*Account, error) {
	req, err := a.newRequest("POST", "/account.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountCreate(reqBody AccountCreateParams) (*Account, error) {
	req, err := a.newRequest("POST", "/account.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountUpdate(reqBody AccountUpdateParams) (*Account, error) {
	req, err := a.newRequest("POST", "/account.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountUploadAvatar(reqBody UploadAvatarParams) (*Account, error) {
	req, err := a.newRequest("POST", "/account.upload_avatar", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Account
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountAssignUserByUserId(reqBody AccountAssignUserByUserIdParams) error {
	req, err := a.newRequest("POST", "/account.assign_user", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) AccountAssignUserByEmail(reqBody AccountAssignUserByEmailParams) error {
	req, err := a.newRequest("POST", "/account.assign_user", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) AccountUnassignUser(reqBody AccountUnassignUserParams) error {
	req, err := a.newRequest("POST", "/account.unassign_user", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) AccountGetMembers(reqBody ListByIdParams) (*MemberList, error) {
	req, err := a.newRequest("POST", "/account.get_members", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MemberList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetUsers(reqBody ListByIdParams) (*UserList, error) {
	req, err := a.newRequest("POST", "/account.get_users", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data UserList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetDescendants(reqBody ListParams) (*AccountList, error) {
	req, err := a.newRequest("POST", "/account.get_descendants", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccountList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetWallets(reqBody ListByIdParams) (*WalletList, error) {
	req, err := a.newRequest("POST", "/account.get_wallets", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data WalletList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetTransactions(reqBody ListByIdParams) (*TransactionList, error) {
	req, err := a.newRequest("POST", "/account.get_transactions", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetTransactionRequests(reqBody ListByIdParams) (*TransactionRequestList, error) {
	req, err := a.newRequest("POST", "/account.get_transaction_requests", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequestList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccountGetTransactionConsumptions(reqBody ListByIdParams) (*TransactionComsumptionList, error) {
	req, err := a.newRequest("POST", "/account.get_transaction_requests", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumptionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// Wallet - Resources related to wallets.
/////////////////
func (a *AdminAPI) WalletAll(reqBody ListParams) (*WalletList, error) {
	req, err := a.newRequest("POST", "/wallet.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data WalletList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) WalletGet(reqBody AddressParam) (*Wallet, error) {
	req, err := a.newRequest("POST", "/wallet.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Wallet
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) WalletCreateByUserId(reqBody WalletCreateByUserIdParams) (*Wallet, error) {
	req, err := a.newRequest("POST", "/wallet.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Wallet
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) WalletCreateByAccountId(reqBody WalletCreateByAccountIdParams) (*Wallet, error) {
	req, err := a.newRequest("POST", "/wallet.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Wallet
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) WalletCreateByProviderUserId(reqBody WalletCreateByProviderUserIdParams) (*Wallet, error) {
	req, err := a.newRequest("POST", "/wallet.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Wallet
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) WalletGetTransactionConsumptions(reqBody AddressListParams) (*TransactionComsumptionList, error) {
	req, err := a.newRequest("POST", "/wallet.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumptionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// Transaction - Resources related to transactions.
/////////////////
func (a *AdminAPI) TransactionAll(reqBody ListParams) (*TransactionList, error) {
	req, err := a.newRequest("POST", "/transaction.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionGet(reqBody IdParam) (*Transaction, error) {
	req, err := a.newRequest("POST", "/transaction.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Transaction
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionCreate(reqBody TransactionCreateParams) (*Transaction, error) {
	req, err := a.newRequest("POST", "/transaction.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Transaction
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionCalculate(reqBody TransactionCalculateParams) (*TransactionCalculation, error) {
	req, err := a.newRequest("POST", "/transaction.calculate", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionCalculation
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// TransactionRequest - Resources related to the creation of transaction requests (either receiving or sending) that needs to be consumed by another user.
/////////////////
func (a *AdminAPI) TransactionRequestAll(reqBody ListParams) (*TransactionRequestList, error) {
	req, err := a.newRequest("POST", "/transaction_request.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequestList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionRequestGet(reqBody TransactionRequestGetParam) (*TransactionRequest, error) {
	req, err := a.newRequest("POST", "/transaction_request.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequest
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionRequestCreate(reqBody TransactionRequestCreateParams) (*TransactionRequest, error) {
	req, err := a.newRequest("POST", "/transaction_request.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionRequest
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionRequestConsume(reqBody TransactionRequestConsumeParams) (*TransactionComsumption, error) {
	req, err := a.newRequest("POST", "/transaction_request.consume", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionRequestGetTransactionConsumptions(reqBody TransactionRequestConsumeListParams) (*TransactionComsumptionList, error) {
	req, err := a.newRequest("POST", "/transaction_request.get_transaction_consumptions", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumptionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// TransactionConsumption - Resources related to consumption of transaction requests.
/////////////////
func (a *AdminAPI) TransactionConsumptionAll(reqBody ListParams) (*TransactionComsumptionList, error) {
	req, err := a.newRequest("POST", "/transaction_consumption.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumptionList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionConsumptionGet(reqBody IdParam) (*TransactionComsumption, error) {
	req, err := a.newRequest("POST", "/transaction_request.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionConsumptionApprove(reqBody IdParam) (*TransactionComsumption, error) {
	req, err := a.newRequest("POST", "/transaction_request.approve", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) TransactionConsumptionReject(reqBody IdParam) (*TransactionComsumption, error) {
	req, err := a.newRequest("POST", "/transaction_request.reject", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data TransactionComsumption
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// API Access - Resources related to API access.
/////////////////
func (a *AdminAPI) AccessKeyAll(reqBody ListParams) (*AccessKeyList, error) {
	req, err := a.newRequest("POST", "/access_key.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccessKeyList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) AccessKeyCreate() (*AccessKey, error) {
	req, err := a.newRequest("POST", "/access_key.create", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccessKey
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, nil
}

func (a *AdminAPI) AccessKeyUpdate(reqBody AccessKeyUpdateParams) (*AccessKey, error) {
	req, err := a.newRequest("POST", "/access_key.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data AccessKey
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (a *AdminAPI) AccessKeyDelete(reqBody IdParam) error {
	req, err := a.newRequest("POST", "/access_key.delete", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) APIKeyAll(reqBody ListParams) (*APIKeyList, error) {
	req, err := a.newRequest("POST", "/api_key.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data APIKeyList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) APIKeyCreate() (*APIKey, error) {
	req, err := a.newRequest("POST", "/api_key.create", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data APIKey
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (a *AdminAPI) APIKeyDelete(reqBody IdParam) error {
	req, err := a.newRequest("POST", "/api_key.delete", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) APIKeyUpdate(reqBody APIKeyUpdateParams) (*APIKey, error) {
	req, err := a.newRequest("POST", "/api_key.update", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data APIKey
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, nil
}

/////////////////
// Setting
/////////////////
func (a *AdminAPI) Setting() (*Setting, error) {
	req, err := a.newRequest("POST", "/settings.all", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data Setting
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, nil
}
