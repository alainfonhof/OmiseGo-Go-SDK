# Go client for OmiseGo eWallet and Admin API
Go client for all admin and server endpoints [OmiseGo](https://github.com/omisego/ewallet). Keep in mind that the tests aren't done yet and will return fails. Feedback is warmly welcomed. 

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
    Path:   "/admin/api",
  }
  c, _ := omg.NewClient("apiKeyId", "apiKey", adminURL)
  adminClient := omg.AdminAPI{
    Client: c,
  }

  // Login to authenticate the client with UserAuth
  body := omg.LoginParams{
    Email:    "your@email.com",
    Password: "pwd",
  }
  authToken, err := adminClient.Login(body)
  
  // Create your access key for ServerAuth in the EWallet API
  accessKey, err := adminClient.AccessKeyCreate()
  
  // Now we can create a client for the EWallet API
  ewalletURL := &url.URL{
    Scheme: "http",
    Host:   "localhost:4000",
    Path:   "/api",
  }
  sa := &omg.ServerAuth{
    AccessKey: accessKey.AccessKey,
    SecretKey: accessKey.SecretKey,
  }
  serverUser := omg.EWalletAPI{
  Client: &omg.Client{
      Auth:       sa,
      HttpClient: &http.Client{},
      BaseURL:    ewalletURL,
    },
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

  user, err := serverUser.UserCreate(userBody)
}
```



## Coverage 
#### Admin API
- [x] /login
- [x] /logout
- [x] /auth_token.switch_account
- [x] /password.update
- [x] /minted_token.all
- [x] /minted_token.get
- [x] /minted_token.create
- [x] /minted_token.mint
- [x] /account.all
- [x] /account.get
- [x] /account.create
- [x] /account.update
- [x] /account.upload_avatar
- [x] /account.list_users
- [x] /account.assign_user
- [x] /account.unassign_user
- [x] /user.all
- [x] /user.get
- [x] /me.get
- [x] /me.get_account
- [x] /me.get_accounts
- [x] /invite.accept
- [x] /admin.all
- [x] /admin.get
- [x] /admin.upload_avatar
- [x] /transaction.all
- [x] /transaction.get
- [x] /access_key.all
- [x] /access_key.create
- [x] /access_key.delete
- [x] /api_key.all
- [x] /api_key.create
- [x] /api_key.delete

#### EWallet API
- [x] /login
- [x] /logout
- [x] /user.create
- [x] /user.update
- [x] /user.get
- [x] /me.get
- [x] /user.list_balances
- [ ] /me.list_balances
- [x] /user.credit_balance
- [x] /user.debit_balance
- [x] /transfer
- [x] /get_settings
- [ ] /me.get_settings
- [x] /transaction.all
- [x] /user.list_transactions
- [ ] /me.list_transactions
- [x] /transaction_request.create
- [x] /transaction_request.get
- [x] /transaction_request.consume
- [ ] /me.create_transaction_request
- [ ] /me.get_transaction_request
- [ ] /me.consume_transaction_request
- [x] /transaction_consumption.approve
- [x] /transaction_consumption.reject
- [ ] /me.approve_transaction_consumption
- [ ] /me.reject_transaction_consumption
