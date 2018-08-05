[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/Alainy/OmiseGo-Go-SDK)

# Go client for OmiseGo eWallet and Admin API
Go client for all eWallet Admin API endpoints [OmiseGo](https://github.com/omisego/ewallet). It has been updated for v1.0 of the OMG eWallet (commit 6c1aca9). Keep in mind that the tests aren't done yet. Feedback is warmly welcomed. 

## Getting started
`go get github.com/Alainy/OmiseGo-Go-SDK`

```go
import (
  omg "github.com/Alainy/OmiseGo-Go-SDK"
  "net/url"
)

func main(){
  // Create a client
  adminURL := &url.URL{
    Scheme: "http",
    Host:   "localhost:4000",
    Path:   "/api/admin",
  }
  c, _ := omg.NewClient("your-accessKey", "your-secretKey", adminURL)
  adminClient := omg.AdminAPI{
    Client: c,
  }

  // Lets create a new user
  userBody := omg.UserParams{
    ProviderUserId: "provider_id",
    Username:       "good_username",
    Metadata: map[string]interface{}{
      "first_name": "Henk",
      "last_name":  "Beckwith",
    },
  }

  user, err := adminClient.UserCreate(userBody)
}
```



## Coverage 
#### Admin API
- [x] /admin.login
- [x] /me.logout
- [x] /auth_token.switch_account
- [x] /admin.reset_password
- [x] /admin.update_password
- [x] /admin.all
- [x] /admin.get
- [x] /invite.accept
- [x] /me.get
- [x] /me.update
- [x] /me.upload_avatar
- [x] /me.get_account
- [x] /me.get_accounts
- [x] /user.login
- [x] /user.logout
- [x] /user.all
- [x] /user.create
- [x] /user.update
- [x] /user.get
- [x] /user.get_wallets
- [x] /user.get_transactions
- [x] /user.get_transaction_consumptions
- [x] /token.all
- [x] /token.get
- [x] /token.create
- [x] /token.update
- [x] /token.stats
- [x] /token.get_mints
- [x] /token.mint
- [x] /exchange_pair.all
- [x] /exchange_pair.get
- [x] /exchange_pair.create
- [x] /exchange_pair.update
- [x] /exchange_pair.delete
- [x] /category.all
- [x] /category.get
- [x] /category.create
- [x] /category.update
- [x] /category.delete
- [x] /account.all
- [x] /account.get
- [x] /account.create
- [x] /account.update
- [x] /account.upload_avatar
- [x] /account.assign_user
- [x] /account.unassign_user
- [x] /account.get_members
- [x] /account.get_users
- [x] /account.get_descendants
- [x] /account.get_wallets
- [x] /account.get_transactions
- [x] /account.get_transaction_requests
- [x] /account.get_transaction_consumptions
- [x] /wallet.all
- [x] /wallet.get
- [x] /wallet.create
- [x] /wallet.get_transaction_consumptions
- [x] /transaction.all
- [x] /transaction.get
- [x] /transaction.create
- [x] /transaction.calculate
- [x] /transaction_request.all
- [x] /transaction_request.get
- [x] /transaction_request.create_request
- [x] /transaction_request.consume
- [x] /transaction_request.get_transaction_consumptions
- [x] /transaction_consumption.all
- [x] /transaction_consumption.get
- [x] /transaction_consumption.approve
- [x] /transaction_consumption.reject
- [x] /access_key.all
- [x] /access_key.create
- [x] /access_key.update
- [x] /access_key.delete
- [x] /api_key.all
- [x] /api_key.create
- [x] /api_key.update
- [x] /api_key.delete
- [x] /settings.all

#### EWallet API
- [x] /me.logout
- [x] /me.get
- [ ] /me.get_wallets
- [ ] /me.get_transactions
- [ ] /me.create_transaction
- [ ] /me.create_transaction_request
- [ ] /me.get_transaction_request
- [ ] /me.consume_transaction_request
- [ ] /me.approve_transaction_consumption
- [ ] /me.reject_transaction_consumption
- [ ] /me.get_settings
