package omisego

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type EWalletAPI struct {
	*Client
}

/////////////////
// Session - Resources related to session tokens.
/////////////////
func (e *EWalletAPI) MeLogout() error {
	req, err := e.newRequest("POST", "/me.logout", nil)
	if err != nil {
		return err
	}

	_, err = e.do(req)
	return err
}

/////////////////
// User - Resources related to users. A user is an entity uniquely identified by its provider_user_id which is the user id in the provider database.
/////////////////
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
