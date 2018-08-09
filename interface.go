package omisego

// EwalletAdminAPI provides an interface to enable mocking the
// omg.AdminAPI service client's API operations.
// This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations.

type EwalletAdminAPI interface {
	AdminLogin(reqBody AdminLoginParams) (*AuthenicationToken, error)
	AdminLogout() error
	AuthTokenSwitchAccount(reqBody AuthTokenSwitchAccountParams) (*AuthenicationToken, error)
	PasswordReset(reqBody AdminResetPasswordParams) error
	PasswordUpdate(reqBody AdminUpdatePasswordParams) error
	AdminAll(reqBody ListParams) (*UserList, error)
	AdminGet(reqBody IdParam) (*User, error)
	InviteAccept() (*User, error)
	MeGet() (*User, error)
	MeUpdate(reqBody MeUpdateParams) (*User, error)
	MeUploadAvatar(reqBody UploadAvatarParams) (*User, error)
	MeGetAccount() (*Account, error)
	MeGetAccounts() (*AccountList, error)
	UserLogin(reqBody IdParam) (*AuthenicationToken, error)
	UserLogout(reqBody AuthTokenParam) error
	UserAll(reqBody ListParams) (*UserList, error)
	UserCreate(reqBody UserParams) (*User, error)
	UserUpdate(reqBody UserParams) (*User, error)
	UserGetById(reqBody IdParam) (*User, error)
	UserGetByProviderUserId(reqBody ProviderUserIdParam) (*User, error)
	UserGetWalletsById(reqBody ListByIdParams) (*WalletList, error)
	UserGetWalletsByProviderUserId(reqBody ListByProviderUserIdParams) (*WalletList, error)
	UserGetTransactionsById(reqBody ListByIdParams) (*TransactionList, error)
	UserGetTransactionsByProviderUserId(reqBody ListByProviderUserIdParams) (*TransactionList, error)
	UserGetTransactionConsumption(reqBody ListByUserIdParams) (*TransactionComsumptionList, error)
	TokenAll(reqBody ListParams) (*TokenList, error)
	TokenGet(reqBody IdParam) (*Token, error)
	TokenCreate(reqBody TokenCreateParams) (*Token, error)
	TokenUpdate(reqBody TokenUpdateParams) (*Token, error)
	TokenStats(reqBody IdParam) (*TokenStats, error)
	TokenGetMints(reqBody ListParams) (*MintList, error)
	TokenMint(reqBody TokenMintParams) (*Token, error)
	ExchangePairAll(reqBody ListParams) (*ExchangePairList, error)
	ExchangePairGet(reqBody IdParam) (*ExchangePair, error)
	ExchangePairCreate(reqBody ExchangePairCreateParams) (*ExchangePairList, error)
	ExchangePairUpdate(reqBody ExchangePairUpdateParams) (*ExchangePairList, error)
	ExchangePairDelete(reqBody IdParam) (*ExchangePairList, error)
	CategoryAll(reqBody ListParams) (*CategoryList, error)
	CategoryGet(reqBody IdParam) (*Category, error)
	CategoryCreate(reqBody ExchangePairCreateParams) (*Category, error)
	CategoryUpdate(reqBody ExchangePairUpdateParams) (*Category, error)
	CategoryDelete(reqBody IdParam) (*Category, error)
	AccountAll(reqBody ListParams) (*AccountList, error)
	AccountGet(reqBody IdParam) (*Account, error)
	AccountCreate(reqBody AccountCreateParams) (*Account, error)
	AccountUpdate(reqBody AccountUpdateParams) (*Account, error)
	AccountUploadAvatar(reqBody UploadAvatarParams) (*Account, error)
	AccountAssignUserByUserId(reqBody AccountAssignUserByUserIdParams) error
	AccountAssignUserByEmail(reqBody AccountAssignUserByEmailParams) error
	AccountUnassignUser(reqBody AccountUnassignUserParams) error
	AccountGetMembers(reqBody ListByIdParams) (*MemberList, error)
	AccountGetUsers(reqBody ListByIdParams) (*UserList, error)
	AccountGetDescendants(reqBody ListParams) (*AccountList, error)
	AccountGetWallets(reqBody ListByIdParams) (*WalletList, error)
	AccountGetTransactions(reqBody ListByIdParams) (*TransactionList, error)
	AccountGetTransactionRequests(reqBody ListByIdParams) (*TransactionRequestList, error)
	AccountGetTransactionConsumptions(reqBody ListByIdParams) (*TransactionComsumptionList, error)
	WalletAll(reqBody ListParams) (*WalletList, error)
	WalletGet(reqBody AddressParam) (*Wallet, error)
	WalletCreateByUserId(reqBody WalletCreateByUserIdParams) (*Wallet, error)
	WalletCreateByAccountId(reqBody WalletCreateByAccountIdParams) (*Wallet, error)
	WalletCreateByProviderUserId(reqBody WalletCreateByProviderUserIdParams) (*Wallet, error)
	WalletGetTransactionConsumptions(reqBody AddressListParams) (*TransactionComsumptionList, error)
	TransactionAll(reqBody ListParams) (*TransactionList, error)
	TransactionGet(reqBody IdParam) (*Transaction, error)
	TransactionCreate(reqBody TransactionCreateParams) (*Transaction, error)
	TransactionCalculate(reqBody TransactionCalculateParams) (*TransactionCalculation, error)
	TransactionRequestAll(reqBody ListParams) (*TransactionRequestList, error)
	TransactionRequestGet(reqBody TransactionRequestGetParam) (*TransactionRequest, error)
	TransactionRequestCreate(reqBody TransactionRequestCreateParams) (*TransactionRequest, error)
	TransactionRequestConsume(reqBody TransactionRequestConsumeParams) (*TransactionComsumption, error)
	TransactionRequestGetTransactionConsumptions(reqBody TransactionRequestConsumeListParams) (*TransactionComsumptionList, error)
	TransactionConsumptionAll(reqBody ListParams) (*TransactionComsumptionList, error)
	TransactionConsumptionGet(reqBody IdParam) (*TransactionComsumption, error)
	TransactionConsumptionApprove(reqBody IdParam) (*TransactionComsumption, error)
	TransactionConsumptionReject(reqBody IdParam) (*TransactionComsumption, error)
	AccessKeyAll(reqBody ListParams) (*AccessKeyList, error)
	AccessKeyCreate() (*AccessKey, error)
	AccessKeyUpdate(reqBody AccessKeyUpdateParams) (*AccessKey, error)
	AccessKeyDelete(reqBody IdParam) error
	APIKeyAll(reqBody ListParams) (*APIKeyList, error)
	APIKeyCreate() (*APIKey, error)
	APIKeyDelete(reqBody IdParam) error
	APIKeyUpdate(reqBody APIKeyUpdateParams) (*APIKey, error)
	Setting() (*Setting, error)
}
