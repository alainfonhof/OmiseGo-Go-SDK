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

	ErrorResponse struct {
		Code        string            `mapstructure:"code"`
		Description string            `mapstructure:"description"`
		Messages    map[string]string `mapstructure:"messages"`
	}

	AuthenicationToken struct {
		UserId              string `mapstructure:"user_id"`
		User                `mapstructure:"user"`
		MasterAdmin         bool   `mapstructure:"master_admin"`
		AuthenticationToken string `mapstructure:"authentication_token"`
		AccountId           string `mapstructure:"account_id"`
		Account             `mapstructure:"account"`
	}

	User struct {
		Id                string                 `mapstructure:"id"`
		ProviderUserId    string                 `mapstructure:"provider_user_id"`
		Username          string                 `mapstructure:"username"`
		Email             string                 `mapstructure:"email"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		Avatar            map[string]interface{} `mapstructure:"avatar"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
	}

	UserList struct {
		Data       []User `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Account struct {
		Id                string                 `mapstructure:"id"`
		ParentId          string                 `mapstructure:"parent_id"`
		Name              string                 `mapstructure:"name"`
		Description       string                 `mapstructure:"description"`
		Master            bool                   `mapstructure:"master"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
		Avatar            map[string]interface{} `mapstructure:"avatar"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
	}

	AccountList struct {
		Data       []Account `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	MintedToken struct {
		Id                string                 `mapstructure:"id"`
		Symbol            string                 `mapstructure:"symbol"`
		Name              string                 `mapstructure:"name"`
		SubunitToUnit     int                    `mapstructure:"subunit_to_unit"`
		CreatedAt         string                 `mapstructure:"created_at"`
		UpdatedAt         string                 `mapstructure:"updated_at"`
		Metadata          map[string]interface{} `mapstructure:"metadata"`
		EncryptedMetadata map[string]interface{} `mapstructure:"encrypted_metadata"`
	}

	MintedTokenList struct {
		Data       []MintedToken `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Transaction struct {
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
		Data       []Transaction `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	TransactionSource struct {
		Address     string `mapstructure:"address"`
		Amount      int    `mapstructure:"amount"`
		MintedToken `mapstructure:"minted_token"`
	}

	Exchange struct {
		Rate int `mapstructure:"rate"`
	}

	Pagination struct {
		PerPage     int  `mapstructure:"per_page"`
		CurrentPage int  `mapstructure:"current_page"`
		IsFirstPage bool `mapstructure:"is_first_page"`
		IsLastPage  bool `mapstructure:"is_last_page"`
	}

	AccessKey struct {
		Id        string `mapstructure:"id"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		AccountId string `mapstructure:"account_id"`
		CreatedAt string `mapstructure:"created_at"`
		UpdatedAt string `mapstructure:"updated_at"`
		DeletedAt string `mapstructure:"deleted_at"`
	}

	AccessKeyList struct {
		Data       []AccessKey `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	APIKey struct {
		Id        string `mapstructure:"id"`
		Key       string `mapstructure:"key"`
		OwnerApp  string `mapstructure:"owner_app"`
		AccountId string `mapstructure:"account_id"`
		CreatedAt string `mapstructure:"created_at"`
		UpdatedAt string `mapstructure:"updated_at"`
		DeletedAt string `mapstructure:"deleted_at"`
	}

	APIKeyList struct {
		Data       []APIKey `mapstructure:"data"`
		Pagination `mapstructure:"pagination"`
	}

	Address struct {
		Address  string    `mapstructure:"address"`
		Balances []Balance `mapstructure:"balances"`
	}

	AddressList struct {
		Data []Address `mapstructure:"data"`
	}

	Balance struct {
		MintedToken `mapstructure:"minted_token"`
		Amount      int `mapstructure:"amount"`
	}

	Settings struct {
		MintedTokens []MintedToken `mapstructure:"minted_tokens"`
	}

	TransactionRequest struct {
		Version             string                 `mapstructure:"version"`
		Success             bool                   `mapstructure:"success"`
		Data                map[string]interface{} `mapstructure:"data"`
		Id                  string                 `mapstructure:"id"`
		SocketTopic         string                 `mapstructure:"socket_topic"`
		Type                string                 `mapstructure:"type"`
		Amount              string                 `mapstructure:"amount"`
		Status              string                 `mapstructure:"status"`
		CorrelationId       string                 `mapstructure:"correlation_id"`
		MintedTokenId       string                 `mapstructure:"minted_token_id"`
		MintedToken         map[string]interface{} `mapstructure:"minted_token"`
		AccountId           string                 `mapstructure:"account_id"`
		UserId              string                 `mapstructure:"user_id"`
		Address             string                 `mapstructure:"address"`
		RequireConfirmation bool                   `mapstructure:"require_confirmation"`
		MaxConsumptions     int                    `mapstructure:"max_consumptions"`
		ConsumptionLifetime int                    `mapstructure:"consumption_lifetime"`
		ExpirationReason    string                 `mapstructure:"expiration_reason"`
		ExpirationDate      string                 `mapstructure:"expiration_date"`
		AllowAmountOverride bool                   `mapstructure:"allow_amount_override"`
		Metadata            map[string]interface{} `mapstructure:"id"`
		EncryptedMetadata   map[string]interface{} `mapstructure:"id"`
		CreatedAt           string                 `mapstructure:"created_at"`
		UpdatedAt           string                 `mapstructure:"updated_at"`
		ExpiredAt           string                 `mapstructure:"expired_at"`
	}

	TransactionComsumption struct {
		Version              string                 `mapstructure:"version"`
		Success              bool                   `mapstructure:"success"`
		Data                 map[string]interface{} `mapstructure:"data"`
		Id                   string                 `mapstructure:"id"`
		SocketTopic          string                 `mapstructure:"socket_topic"`
		Amount               string                 `mapstructure:"amount"`
		Status               string                 `mapstructure:"status"`
		CorrelationId        string                 `mapstructure:"correlation_id"`
		MintedTokenId        string                 `mapstructure:"minted_token_id"`
		MintedToken          map[string]interface{} `mapstructure:"minted_token"`
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
		CreatedAt            string                 `mapstructure:"created_at"`
		UpdatedAt            string                 `mapstructure:"updated_at"`
		ExpiredAt            string                 `mapstructure:"expired_at"`
		ApprovedAt           string                 `mapstructure:"approved_at"`
		RejectedAt           string                 `mapstructure:"rejected_at"`
		ConfirmedAt          string                 `mapstructure:"confirmed_at"`
		FailedAt             string                 `mapstructure:"failed_at"`
	}

	////////////
	// Request body parameters
	////////////
	ListParams struct {
		Page       int    `json:"page,omitempty"`
		PerPage    int    `json:"per_page,omitempty"`
		SearchTerm string `json:"search_term,omitempty"`
		SortBy     string `json:"sort_by,omitempty"`
		SortDir    string `json:"sort_dir,omitempty"`
	}

	ByIdParam struct {
		Id string `json:"id"`
	}

	ProviderUserIdParam struct {
		ProviderUserId string `json:"provider_user_id"`
	}
	UserParams struct {
		ProviderUserId    string                 `json:"provider_user_id"`
		Username          string                 `json:"username"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	LoginParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthTokenSwitchAccountParams struct {
		AccountId string `json:"account_id"`
	}

	PasswordResetParams struct {
		Email       string `json:"email"`
		RedirectUrl string `json:"redirect_url"`
	}

	PasswordUpdateParams struct {
		Email                string `json:"email"`
		Token                string `json:"token"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	MintedTokenCreateParams struct {
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

	MintedTokenMintParams struct {
		Id     string `json:"id"`
		Amount int    `json:"amount"`
	}

	AccountCreateParams struct {
		Name              string                 `json:"name"`
		Description       string                 `json:"description,omitempty"`
		ParentId          string                 `json:"parent_id,omitempty"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	AccountUpdateParams struct {
		Id                string                 `json:"id"`
		Name              string                 `json:"name"`
		Description       string                 `json:"description"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}

	AccountUploadAvatarParams struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
	}

	AccountListUsersParams struct {
		AccountId string `json:"account_id"`
	}

	AccountAssignUserParams struct {
		UserId      string `json:"user_id,omitempty"`
		AccountId   string `json:"account_id"`
		RoleName    string `json:"role_name"`
		RedirectUrl string `json:"redirect_url,omitempty"`
		Email       string `json:"email,omitempty"`
	}

	AccountUnassignUserParams struct {
		UserId    string `json:"user_id"`
		AccountId string `json:"account_id"`
	}

	AdminUploadAvatarParams struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
	}

	AccessKeyDeleteParams struct {
		Id        string `json:"id,omitempty"`
		AccessKey string `json:"access_key,omitempty"`
	}

	APIKeyCreateParams struct {
		OwnerApp string `json:"owner_app"`
	}

	BalanceAdjustmentParams struct {
		ProviderUserId        string                 `json:"provider_user_id"`
		TokenId               string                 `json:"token_id"`
		Amount                int                    `json:"amount"`
		AccountId             string                 `json:"account_id,omitempty"`
		BurnBalanceIdentifier string                 `json:"burn_balance_identifier,omitempty"`
		Metadata              map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata     map[string]interface{} `json:"id,omitempty"`
	}

	TransferParams struct {
		FromAddress       string                 `json:"from_address"`
		ToAddress         string                 `json:"to_address"`
		TokenId           string                 `json:"token_id"`
		Amount            int                    `json:"amount"`
		Metadata          map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata map[string]interface{} `json:"id,omitempty"`
	}

	UserListTransactionsParams struct {
		ProviderUserId string                 `json:"provider_user_id"`
		Address        string                 `json:"address"`
		Page           int                    `json:"page,omitempty"`
		PerPage        int                    `json:"per_page,omitempty"`
		SearchTerm     string                 `json:"search_term,omitempty"`
		SearchTerms    map[string]interface{} `json:"search_terms,omitempty"`
		SortBy         string                 `json:"sort_by,omitempty"`
		SortDir        string                 `json:"sort_dir,omitempty"`
	}

	ServerCreateTransactionRequestParams struct {
		Type                string                 `json:"type"`
		TokenId             string                 `json:"token_id"`
		Amount              int                    `json:"amount,omitempty"`
		CorrelationId       string                 `json:"correlation_id,omitempty"`
		AccountId           string                 `json:"account_id,omitempty"`
		ProviderUserId      string                 `json:"provider_user_id,omitempty"`
		Address             string                 `json:"address,omitempty"`
		RequireConfirmation bool                   `json:"require_confirmation,omitempty"`
		MaxConsumptions     int                    `json:"max_consumptions,omitempty"`
		ConsumptionLifetime int                    `json:"consumption_lifetime,omitempty"`
		ExpirationDate      string                 `json:"expiration_date,omitempty"`
		AllowAmountOverride bool                   `json:"allow_amount_override,omitempty"`
		Metadata            map[string]interface{} `json:"id,omitempty"`
		EncryptedMetadata   map[string]interface{} `json:"id,omitempty"`
	}

	ServerTransactionRequestConsumeParams struct {
		TransactionRequestId string                 `json:"transaction_request_id"`
		TokenId              string                 `json:"token_id,omitempty"`
		Amount               int                    `json:"amount,omitempty"`
		CorrelationId        string                 `json:"correlation_id,omitempty"`
		AccountId            string                 `json:"account_id,omitempty"`
		ProviderUserId       string                 `json:"provider_user_id,omitempty"`
		Address              string                 `json:"address,omitempty"`
		Metadata             map[string]interface{} `json:"metadata,omitempty"`
		EncryptedMetadata    map[string]interface{} `json:"encrypted_metadata,omitempty"`
	}
)

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%+v", *e)
}
