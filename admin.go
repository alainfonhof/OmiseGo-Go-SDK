package omisego

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type AdminAPI struct {
	*Client
}

/////////////////
// Session
/////////////////
func (a *AdminAPI) Login(reqBody LoginParams) (*AuthenicationToken, error) {
	req, err := a.newRequest("POST", "/login", reqBody)
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

	// Log the admin in with new authentication
	a.Auth = &AdminUserAuth{
		ApiKey:        a.Auth.(*AdminClientAuth).ApiKey,
		ApiKeyId:      a.Auth.(*AdminClientAuth).ApiKeyId,
		UserId:        data.UserId,
		UserAuthToken: data.AuthenticationToken,
	}

	return &data, err
}

func (a *AdminAPI) Logout() error {
	req, err := a.newRequest("POST", "/logout", nil)
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

func (a *AdminAPI) PasswordReset(reqBody PasswordResetParams) error {
	req, err := a.newRequest("POST", "/password.reset", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

func (a *AdminAPI) PasswordUpdate(reqBody PasswordUpdateParams) error {
	req, err := a.newRequest("POST", "/password.update", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}

/////////////////
// Minted Token
/////////////////
func (a *AdminAPI) MintedTokenAll(reqBody ListParams) (*MintedTokenList, error) {
	req, err := a.newRequest("POST", "/minted_token.all", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MintedTokenList
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MintedTokenGet(reqBody ByIdParam) (*MintedToken, error) {
	req, err := a.newRequest("POST", "/minted_token.get", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MintedToken
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MintedTokenCreate(reqBody MintedTokenCreateParams) (*MintedToken, error) {
	req, err := a.newRequest("POST", "/minted_token.create", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MintedToken
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

func (a *AdminAPI) MintedTokenMint(reqBody MintedTokenMintParams) (*MintedToken, error) {
	req, err := a.newRequest("POST", "/minted_token.mint", reqBody)
	if err != nil {
		return nil, err
	}

	res, err := a.do(req)
	if err != nil {
		return nil, err
	}

	var data MintedToken
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}

	return &data, err
}

/////////////////
// Account
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

func (a *AdminAPI) AccountGet(reqBody ByIdParam) (*Account, error) {
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

func (a *AdminAPI) AccountUploadAvatar(reqBody AccountUploadAvatarParams) (*Account, error) {
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

func (a *AdminAPI) AccountListUsers(reqBody AccountListUsersParams) (*UserList, error) {
	req, err := a.newRequest("POST", "/account.list_users", reqBody)
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

func (a *AdminAPI) AccountAssignUser(reqBody AccountAssignUserParams) error {
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

/////////////////
// User
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

func (a *AdminAPI) UserGet(reqBody ByIdParam) (*User, error) {
	req, err := a.newRequest("POST", "/user.get", reqBody)
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

/////////////////
// Admin
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

func (a *AdminAPI) AdminGet(reqBody ByIdParam) (*User, error) {
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

func (a *AdminAPI) AdminUploadAvatar(reqBody AdminUploadAvatarParams) (*User, error) {
	req, err := a.newRequest("POST", "/admin.upload_avatar", reqBody)
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

/////////////////
// Transaction
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

func (a *AdminAPI) TransactionGet(reqBody ByIdParam) (*Transaction, error) {
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

/////////////////
// API Access
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
	var data AccessKey
	err = mapstructure.Decode(res.Data, &data)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong with decoding %+v to %T", res.Data, data)
	}
	return &data, nil
}

func (a *AdminAPI) AccessKeyDelete(reqBody AccessKeyDeleteParams) error {
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

func (a *AdminAPI) APIKeyCreate(reqBody APIKeyCreateParams) (*APIKey, error) {
	req, err := a.newRequest("POST", "/api_key.create", reqBody)
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

func (a *AdminAPI) APIKeyDelete(reqBody ByIdParam) error {
	req, err := a.newRequest("POST", "/api_key.delete", reqBody)
	if err != nil {
		return err
	}

	_, err = a.do(req)
	return err
}
