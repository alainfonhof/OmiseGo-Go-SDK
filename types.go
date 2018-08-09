package omisego

import (
	"fmt"
)

type (
	BaseResponse struct {
		Version string                 `json:"version"`
		Success bool                   `json:"success"`
		Data    map[string]interface{} `json:"data"`
	}

	UnpaginatedList struct {
		Object string        `mapstructure:"object"`
		Data   []interface{} `mapstructure:"data"`
	}

	PaginatedList struct {
		Object     string        `mapstructure:"object"`
		Data       []interface{} `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Pagination struct {
		PerPage     int  `mapstructure:"per_page"`
		CurrentPage int  `mapstructure:"current_page"`
		IsFirstPage bool `mapstructure:"is_first_page"`
		IsLastPage  bool `mapstructure:"is_last_page"`
	}

	ErrorResponse struct {
		Object      string                 `mapstructure:"object"`
		Code        string                 `mapstructure:"code"`
		Description string                 `mapstructure:"description"`
		Messages    map[string]interface{} `mapstructure:"messages"`
	}

	AuthenicationToken struct {
		Object              string `mapstructure:"object"`
		AuthenticationToken string `mapstructure:"authentication_token"`
		UserId              string `mapstructure:"user_id"`
		User                `mapstructure:"user"`
		AccountId           string `mapstructure:"account_id"`
		MasterAdmin         bool   `mapstructure:"master_admin"`
		Account             `mapstructure:"account"`
		Role                string `mapstructure:"role"`
	}

	Token struct {
		Object            string                 `mapstructure:"object"`
		Id                string                 `mapstructure:"id"`
		Symbol            string                 `mapstructure:"symbol"`
		Name              string                 `mapstructure:"name"`
		SubunitToUnit     int                    `mapstructure:"subunit_to_unit"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
	}

	TokenList struct {
		Object     string  `mapstructure:"object"`
		Data       []Token `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	TokenStats struct {
		Object      string `mapstructure:"object"`
		TokenId     string `mapstructure:"token_id"`
		Token       `mapstructure:"token"`
		TotalSupply int `mapstructure:"total_supply"`
	}

	Mint struct {
		Object        string `mapstructure:"object"`
		Id            string `mapstructure:"id"`
		Description   string `mapstructure:"description"`
		Amount        int    `mapstructure:"amount"`
		Confirmed     bool   `mapstructure:"confirmed"`
		TokenId       string `mapstructure:"token_id"`
		Token         `mapstructure:"token"`
		AccountId     string `mapstructure:"account_id"`
		Account       `mapstructure:"account"`
		TransactionId string `mapstructure:"transaction_id"`
		Transaction   `mapstructure:"transaction"`
		CreatedAt     string `mapstructure:"created_at"`
		UpdatedAt     string `mapstructure:"updated_at"`
	}

	MintList struct {
		Object     string `mapstructure:"object"`
		Data       []Mint `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	ExchangePair struct {
		Object      string  `mapstructure:"object"`
		Id          string  `mapstructure:"id"`
		Name        string  `mapstructure:"name"`
		FromTokenId string  `mapstructure:"from_token_id"`
		FromToken   Token   `mapstructure:"from_token"`
		ToTokenId   string  `mapstructure:"to_token_id"`
		ToToken     Token   `mapstructure:"to_token"`
		Rate        float64 `mapstructure:"rate"`
		CreatedAt   string  `mapstructure:"created_at"`
		UpdatedAt   string  `mapstructure:"updated_at"`
		DeletedAt   string  `mapstructure:"deleted_at"`
	}

	ExchangePairList struct {
		Object     string         `mapstructure:"object"`
		Data       []ExchangePair `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Category struct {
		Object      string          `mapstructure:"object"`
		Id          string          `mapstructure:"id"`
		Name        string          `mapstructure:"name"`
		Description string          `mapstructure:"description"`
		AccountIds  []string        `mapstructure:"account_ids"`
		Accounts    UnpaginatedList `mapstructure:"accounts"`
		CreatedAt   string          `mapstructure:"created_at"`
		UpdatedAt   string          `mapstructure:"updated_at"`
	}

	CategoryList struct {
		Object     string     `mapstructure:"object"`
		Data       []Category `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Account struct {
		Object            string                 `mapstructure:"object"`
		Id                string                 `mapstructure:"id"`
		ParentId          string                 `mapstructure:"parent_id"`
		Name              string                 `mapstructure:"name"`
		Description       string                 `mapstructure:"description"`
		Master            bool                   `mapstructure:"master"`
		Avatar            map[string]interface{} `mapstructure:"avatar"`
		CategoryIds       []string               `mapstructure:"category_ids"`
		Categories        CategoryList           `mapstructure:"categories"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
	}

	AccountList struct {
		Object     string    `mapstructure:"object"`
		Data       []Account `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Member struct {
		Object            string                 `mapstructure:"object"`
		Id                string                 `mapstructure:"id"`
		Username          string                 `mapstructure:"username"`
		ProviderUserId    string                 `mapstructure:"provider_user_id"`
		Email             string                 `mapstructure:"email"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		Avatar            map[string]interface{} `mapstructure:"avatar"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
		Role              string                 `mapstructure:"role"`
		Account           `mapstructure:"account"`
	}

	MemberList struct {
		Object     string   `mapstructure:"object"`
		Data       []Member `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	User struct {
		Object            string                 `mapstructure:"object"`
		Id                string                 `mapstructure:"id"`
		Username          string                 `mapstructure:"username"`
		ProviderUserId    string                 `mapstructure:"provider_user_id"`
		Email             string                 `mapstructure:"email"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		Avatar            map[string]interface{} `mapstructure:"avatar"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
	}

	UserList struct {
		Object     string `mapstructure:"object"`
		Data       []User `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Wallet struct {
		Object            string                 `mapstructure:"object"`
		SocketTopic       string                 `mapstructure:"socket_topic"`
		Address           string                 `mapstructure:"address"`
		Name              string                 `mapstructure:"name"`
		Identifier        string                 `mapstructure:"identifier"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		UserId            string                 `mapstructure:"user_id"`
		User              `mapstructure:"user"`
		AccountId         string `mapstructure:"account_id"`
		Account           `mapstructure:"account"`
		Balances          []Balance `mapstructure:"balances"`
	}

	WalletList struct {
		Object     string   `mapstructure:"object"`
		Data       []Wallet `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Balance struct {
		Object string `mapstructure:"object"`
		Token  `mapstructure:"token"`
		Amount int `mapstructure:"amount"`
	}

	Transaction struct {
		Object            string            `mapstructure:"object"`
		Id                string            `mapstructure:"id"`
		From              TransactionSource `mapstructure:"from"`
		To                TransactionSource `mapstructure:"to"`
		Exchange          `mapstructure:"exchange"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		Status            string                 `mapstructure:"status"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
	}

	TransactionList struct {
		Object     string        `mapstructure:"object"`
		Data       []Transaction `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	TransactionSource struct {
		Object  string  `mapstructure:"object"`
		Address string  `mapstructure:"address"`
		Amount  float64 `mapstructure:"amount"`
		Token   `mapstructure:"token"`
	}

	Exchange struct {
		Object string  `mapstructure:"object"`
		Rate   float64 `mapstructure:"rate"`
	}

	TransactionCalculation struct {
		Object       string `mapstructure:"object"`
		FromAmount   int    `mapstructure:"from_amount"`
		FromTokenId  string `mapstructure:"from_token_id"`
		ToAmount     int    `mapstructure:"to_amount"`
		ToTokenId    string `mapstructure:"to_token_id"`
		ExchangePair `mapstructure:"exchange_pair"`
		CalculatedAt string `mapstructure:"calculated_at"`
	}

	TransactionRequest struct {
		Object                 string                 `mapstructure:"object"`
		Version                string                 `mapstructure:"version"`
		Success                bool                   `mapstructure:"success"`
		Data                   map[string]interface{} `mapstructure:"data"`
		Id                     string                 `mapstructure:"id"`
		SocketTopic            string                 `mapstructure:"socket_topic"`
		Type                   string                 `mapstructure:"type"`
		Amount                 string                 `mapstructure:"amount"`
		Status                 string                 `mapstructure:"status"`
		CorrelationId          string                 `mapstructure:"correlation_id"`
		TokenId                string                 `mapstructure:"token_id"`
		Token                  map[string]interface{} `mapstructure:"token"`
		AccountId              string                 `mapstructure:"account_id"`
		UserId                 string                 `mapstructure:"user_id"`
		Address                string                 `mapstructure:"address"`
		RequireConfirmation    bool                   `mapstructure:"require_confirmation"`
		MaxConsumptions        int                    `mapstructure:"max_consumptions"`
		MaxConsumptionsPerUser int                    `mapstructure:"max_consumptions_per_user"`
		ConsumptionLifetime    int                    `mapstructure:"consumption_lifetime"`
		ExpirationReason       string                 `mapstructure:"expiration_reason"`
		ExpirationDate         string                 `mapstructure:"expiration_date"`
		AllowAmountOverride    bool                   `mapstructure:"allow_amount_override"`
		Metadata               map[string]interface{} `mapstructure:"id"`
		EncryptedMetadata      map[string]interface{} `mapstructure:"id"`
		CreatedAt              string                 `mapstructure:"created_at"`
		UpdatedAt              string                 `mapstructure:"updated_at"`
		ExpiredAt              string                 `mapstructure:"expired_at"`
	}

	TransactionRequestList struct {
		Object     string               `mapstructure:"object"`
		Data       []TransactionRequest `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	TransactionComsumption struct {
		Object               string                 `mapstructure:"object"`
		Version              string                 `mapstructure:"version"`
		Success              bool                   `mapstructure:"success"`
		Data                 map[string]interface{} `mapstructure:"data"`
		Id                   string                 `mapstructure:"id"`
		SocketTopic          string                 `mapstructure:"socket_topic"`
		Amount               string                 `mapstructure:"amount"`
		Status               string                 `mapstructure:"status"`
		CorrelationId        string                 `mapstructure:"correlation_id"`
		TokenId              string                 `mapstructure:"token_id"`
		Token                map[string]interface{} `mapstructure:"token"`
		IdempotencyToken     string                 `mapstructure:"idempotency_token"`
		TransactionId        string                 `mapstructure:"transaction_id"`
		Transaction          map[string]interface{} `mapstructure:"transaction"`
		UserId               string                 `mapstructure:"user_id"`
		User                 map[string]interface{} `mapstructure:"user"`
		AccountId            string                 `mapstructure:"account_id"`
		Account              map[string]interface{} `mapstructure:"account"`
		TransactionRequestId string                 `mapstructure:"transaction_request_id"`
		TransactionRequest   map[string]interface{} `mapstructure:"transaction_request"`
		Address              string                 `mapstructure:"address"`
		Metadata             map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata    map[string]interface{} `mapstructure:"encrypted_metadata"`
		ExpirationDate       string                 `mapstructure:"expiration_date"`
		CreatedAt            string                 `mapstructure:"created_at"`
		UpdatedAt            string                 `mapstructure:"updated_at"`
		ApprovedAt           string                 `mapstructure:"approved_at"`
		RejectedAt           string                 `mapstructure:"rejected_at"`
		ConfirmedAt          string                 `mapstructure:"confirmed_at"`
		FailedAt             string                 `mapstructure:"failed_at"`
		ExpiredAt            string                 `mapstructure:"expired_at"`
	}

	TransactionComsumptionList struct {
		Object     string                   `mapstructure:"object"`
		Data       []TransactionComsumption `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Setting struct {
		Object string  `mapstructure:"object"`
		Tokens []Token `mapstructure:"tokens"`
	}

	AccessKey struct {
		Object    string `mapstructure:"object"`
		Id        string `mapstructure:"id"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		AccountId string `mapstructure:"account_id"`
		CreatedAt string `mapstructure:"created_at"`
		UpdatedAt string `mapstructure:"updated_at"`
		DeletedAt string `mapstructure:"deleted_at"`
	}

	AccessKeyList struct {
		Object     string      `mapstructure:"object"`
		Data       []AccessKey `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	APIKey struct {
		Object    string `mapstructure:"object"`
		Id        string `mapstructure:"id"`
		Key       string `mapstructure:"key"`
		OwnerApp  string `mapstructure:"owner_app"`
		Expired   bool   `mapstructure:"expired"`
		AccountId string `mapstructure:"account_id"`
		CreatedAt string `mapstructure:"created_at"`
		UpdatedAt string `mapstructure:"updated_at"`
		DeletedAt string `mapstructure:"deleted_at"`
	}

	APIKeyList struct {
		Object     string   `mapstructure:"object"`
		Data       []APIKey `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	////////////
	// Request body parameters
	////////////
	AdminLoginParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthTokenSwitchAccountParams struct {
		AccountId string `json:"account_id"`
	}

	AdminResetPasswordParams struct {
		Email       string `json:"email"`
		RedirectUrl string `json:"redirect_url"`
	}

	AdminUpdatePasswordParams struct {
		Email                string `json:"email"`
		Token                string `json:"token"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	ListParams struct {
		Page       int    `json:"page,omitempty"`
		PerPage    int    `json:"per_page,omitempty"`
		SearchTerm string `json:"search_term,omitempty"`
		SortBy     string `json:"sort_by,omitempty"`
		SortDir    string `json:"sort_dir,omitempty"`
	}

	IdParam struct {
		Id string `json:"id"`
	}

	UserIdParam struct {
		UserId string `json:"user_id"`
	}

	ProviderUserIdParam struct {
		ProviderUserId string `json:"provider_user_id"`
	}

	ListByIdParams struct {
		Id string `json:"id"`
		ListParams
	}

	ListByProviderUserIdParams struct {
		ProviderUserId string `json:"provider_user_id"`
		ListParams
	}

	ListByUserIdParams struct {
		UserId string `json:"user_id"`
		ListParams
	}

	ListByIdAndOwnedParams struct {
		Id    string `json:"id"`
		Owned bool   `json:"owned,omitempty"`
		ListParams
	}

	MeUpdateParams struct {
		Email             string                 `json:"email,omitempty"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	UploadAvatarParams struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
	}

	AuthTokenParam struct {
		AuthToken string `json:"auth_token"`
	}

	UserParams struct {
		ProviderUserId    string                 `json:"provider_user_id"`
		Username          string                 `json:"username"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	TokenCreateParams struct {
		Name                 string                 `json:"name"`
		Symbol               string                 `json:"symbol"`
		Description          string                 `json:"description"`
		SubunitToUnit        int                    `json:"subunit_to_unit,omitempty"`
		Amount               int                    `json:"amount,omitempty"`
		IsoCode              string                 `json:"iso_code,omitempty"`
		ShortSymbol          string                 `json:"short_symbol,omitempty"`
		Subunit              string                 `json:"subunit,omitempty"`
		SymbolFirst          bool                   `json:"symbol_first,omitempty"`
		HtmlEntity           string                 `json:"html_entity,omitempty"`
		IsoNumeric           string                 `json:"iso_numeric,omitempty"`
		SmallestDenomination int                    `json:"smallest_denomination,omitempty"`
		Metadata             map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata    map[string]interface{} `json:"id,omitempty"`
	}

	TokenUpdateParams struct {
		Id                string                 `json:"id"`
		Name              string                 `json:"name,omitempty"`
		Description       string                 `json:"description,omitempty"`
		IsoCode           string                 `json:"iso_code,omitempty"`
		ShortSymbol       string                 `json:"short_symbol,omitempty"`
		SymbolFirst       bool                   `json:"symbol_first,omitempty"`
		HtmlEntity        string                 `json:"html_entity,omitempty"`
		IsoNumeric        string                 `json:"iso_numeric,omitempty"`
		Metadata          map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"id,omitempty"`
	}

	TokenMintParams struct {
		Id     string `json:"id"`
		Amount int    `json:"amount"`
	}

	ExchangePairCreateParams struct {
		FromTokenId  string  `json:"from_token_id"`
		ToTokenId    string  `json:"to_token_id"`
		Rate         float64 `json:"rate"`
		SyncOpposite bool    `json:"sync_opposite,omitempty"`
	}

	ExchangePairUpdateParams struct {
		Id   string  `json:"id"`
		Rate float64 `json:"rate,omitempty"`
	}

	CategoryCreateParams struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		AccountIds  []string `json:"account_ids"`
	}

	CategoryUpdateParams struct {
		Id          string   `json:"id"`
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"`
		AccountIds  []string `json:"account_ids,omitempty"`
	}

	AccountCreateParams struct {
		Name              string                 `json:"name"`
		Description       string                 `json:"description,omitempty"`
		ParentId          string                 `json:"parent_id,omitempty"`
		CategoryIds       []string               `json:"category_ids"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	AccountUpdateParams struct {
		Id                string                 `json:"id"`
		Name              string                 `json:"name,omitempty"`
		Description       string                 `json:"description,omitempty"`
		CategoryIds       []string               `json:"category_ids"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	AccountAssignUserByUserIdParams struct {
		UserId      string `json:"user_id"`
		AccountId   string `json:"account_id"`
		RoleName    string `json:"role_name"`
		RedirectUrl string `json:"redirect_url"`
	}

	AccountAssignUserByEmailParams struct {
		Email     string `json:"email"`
		AccountId string `json:"account_id"`
		RoleName  string `json:"role_name"`
	}

	AccountUnassignUserParams struct {
		UserId    string `json:"user_id"`
		AccountId string `json:"account_id"`
	}

	WalletCreateByUserIdParams struct {
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
		UserId     string `json:"user_id"`
	}

	WalletCreateByAccountIdParams struct {
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
		AccountId  string `json:"account_id"`
	}

	WalletCreateByProviderUserIdParams struct {
		Name           string `json:"name"`
		Identifier     string `json:"identifier"`
		ProviderUserId string `json:"provider_user_id"`
	}

	AddressParam struct {
		Address string `json:"address"`
	}

	AddressListParams struct {
		Address string `json:"address"`
		ListParams
	}

	TransactionCreateParams struct {
		IdempotencyToken      string                 `json:"idempotency_token"`
		FromAddress           string                 `json:"from_address"`
		ToAddress             string                 `json:"to_address"`
		FromAccountId         string                 `json:"from_account_id,omitempty"`
		ToAccountId           string                 `json:"to_account_id,omitempty"`
		FromUserId            string                 `json:"from_user_id,omitempty"`
		ToUserId              string                 `json:"to_user_id,omitempty"`
		FromProviderUserId    string                 `json:"from_provider_user_id,omitempty"`
		ToProviderUserId      string                 `json:"to_provider_user_id,omitempty"`
		FromTokenId           string                 `json:"from_token_id,omitempty"`
		ToTokenId             string                 `json:"to_token_id,omitempty"`
		TokenId               string                 `json:"token_id"`
		FromAmount            int                    `json:"from_amount,omitempty"`
		ToAmount              int                    `json:"to_amount,omitempty"`
		Amount                int                    `json:"amount"`
		ExchangeAccountId     string                 `json:"exchange_account_id,omitempty"`
		ExchangeWalletAddress string                 `json:"exchange_wallet_address,omitempty"`
		Metadata              map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata     map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	TransactionCalculateParams struct {
		FromAmount  int    `json:"from_amount,omitempty"`
		FromTokenId string `json:"from_token_id"`
		ToAmount    int    `json:"to_amount,omitempty"`
		ToTokenId   string `json:"to_token_id"`
	}

	TransactionRequestCreateParams struct {
		Type                   string                 `json:"type"`
		TokenId                string                 `json:"token_id"`
		Amount                 int                    `json:"amount,omitempty"`
		CorrelationId          string                 `json:"correlation_id,omitempty"`
		AccountId              string                 `json:"account_id,omitempty"`
		ProviderUserId         string                 `json:"provider_user_id,omitempty"`
		Address                string                 `json:"address,omitempty"`
		RequireConfirmation    bool                   `json:"require_confirmation,omitempty"`
		MaxConsumptions        int                    `json:"max_consumptions,omitempty"`
		MaxConsumptionsPerUser int                    `json:"max_consumptions_per_user,omitempty"`
		ConsumptionLifetime    int                    `json:"consumption_lifetime,omitempty"`
		ExpirationDate         string                 `json:"expiration_date,omitempty"`
		AllowAmountOverride    bool                   `json:"allow_amount_override,omitempty"`
		Metadata               map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata      map[string]interface{} `json:"id,omitempty"`
	}

	TransactionRequestGetParam struct {
		FormattedId string `json:"formatted_id"`
	}

	TransactionRequestConsumeParams struct {
		IdempotencyToken              string                 `json:"idempotency_token"`
		FormattedTransactionRequestId string                 `json:"formatted_transaction_request_id"`
		TokenId                       string                 `json:"token_id,omitempty"`
		Amount                        int                    `json:"amount,omitempty"`
		CorrelationId                 string                 `json:"correlation_id,omitempty"`
		AccountId                     string                 `json:"account_id,omitempty"`
		ProviderUserId                string                 `json:"provider_user_id,omitempty"`
		Address                       string                 `json:"address,omitempty"`
		Metadata                      map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata             map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	TransactionRequestConsumeListParams struct {
		FormattedTransactionRequestId string `json:"formatted_transaction_request_id"`
		ListParams
	}

	AccessKeyParam struct {
		AccessKey string `json:"access_key"`
	}

	AccessKeyUpdateParams struct {
		Id      string `json:"id"`
		Expired bool   `json:"expired"`
	}

	APIKeyUpdateParams struct {
		Id      string `json:"id"`
		Expired bool   `json:"expired"`
	}
)

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", *e)
}
