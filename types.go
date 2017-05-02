package paypalsdk

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = "https://api.sandbox.paypal.com"

	// APIBaseLive points to the live version of the API
	APIBaseLive = "https://api.paypal.com"

	// RequestNewTokenBeforeExpiresIn is used by SendWithAuth and try to get new Token when it's about to expire
	RequestNewTokenBeforeExpiresIn = 60
)

// Possible values for `no_shipping` in InputFields
//
// https://developer.paypal.com/docs/api/payment-experience/#definition-input_fields
const (
	NoShippingDisplay uint = 0
	NoShippingHide uint = 1
	NoShippingBuyerAccount uint = 2
)

// Possible values for `address_override` in InputFields
//
// https://developer.paypal.com/docs/api/payment-experience/#definition-input_fields
const (
	AddrOverrideFromFile uint = 0
	AddrOverrideFromCall uint = 1
)

// Possible values for `landing_page_type` in FlowConfig
//
// https://developer.paypal.com/docs/api/payment-experience/#definition-flow_config
const (
	LandingPageTypeBilling string = "Billing"
	LandingPageTypeLogin string = "Login"
)

type (

	Agreement struct {
		/**
		 * Identifier of the agreement.
		 */
		Id                          string `json:"id,omitempty"`

		/**
		 * State of the agreement
		 */
		State                       string `json:"state,omitempty"`

		/**
		 * Name of the agreement.
		 */
		Name                        string `json:"name,omitempty"`

		/**
		 * Description of the agreement.
		 */
		Description                 string `json:"description,omitempty"`

		/**
		 * Start date of the agreement. Date format yyyy-MM-dd z, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		 */
		StartDate                   string `json:"start_date,omitempty"`

		/**
		 * Details of the agreement.
		 */
		AgreementDetails            AgreementDetails `json:"agreement_details,omitempty"`

		/**
		 * Details of the buyer who is enrolling in this agreement. This information is gathered from execution of the approval URL.
		 */
		Payer                       Payer `json:"payer,omitempty"`

		/**
		 * Shipping address object of the agreement, which should be provided if it is different from the default address.
		 */
		ShippingAddress             Address `json:"shipping_address,omitempty"`

		/**
		 * Default merchant preferences from the billing plan are used, unless override preferences are provided here.
		 */
		OverrideMerchantPreferences MerchantPreferences `json:"override_merchant_preferences,omitempty"`

		/**
		 * Array of override_charge_model for this agreement if needed to change the default models from the billing plan.
		 */
		OverrideChargeModels        []OverrideChargeModel `json:"override_charge_models,omitempty"`

		/**
		 * Plan details for this agreement.
		 */
		Plan                        Plan `json:"plan,omitempty"`

		/**
		 * Date and time that this resource was created. Date format yyyy-MM-dd z, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		 */
		CreateTime                  string `json:"create_time,omitempty"`

		/**
		 * Date and time that this resource was updated. Date format yyyy-MM-dd z, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		 */
		UpdateTime                  string `json:"update_time,omitempty"`

		/**
		 * Payment token
		 */
		Token                       string `json:"token,omitempty"`

		/**
		 *
		 */
		Links                       []Link `json:"links,omitempty"`
	}

	AgreementDetails struct {
		/**
		* The outstanding balance for this agreement.
		*/
		OutstandingBalance AmountPayout `json:"outstanding_balance,omitempty"`

		/**
		 * Number of cycles remaining for this agreement.
		 */
		CyclesRemaining    string `json:"cycles_remaining,omitempty"`

		/**
		 * Number of cycles completed for this agreement.
		 */
		CyclesCompleted    string `json:"cycles_completed,omitempty"`

		/**
		 * The next billing date for this agreement, represented as 2014-02-19T10:00:00Z format.
		 */
		NextBillingDate    string `json:"next_billing_date,omitempty"`

		/**
		 * Last payment date for this agreement, represented as 2014-06-09T09:42:31Z format.
		 */
		LastPaymentDate    string `json:"last_payment_date,omitempty"`

		/**
		 * Last payment amount for this agreement.
		 */
		LastPaymentAmount  string `json:"last_payment_amount,omitempty"`

		/**
		 * Last payment date for this agreement, represented as 2015-02-19T10:00:00Z format.
		 */
		FinalPaymentDate   string `json:"final_payment_date,omitempty"`

		/**
		 * Total number of failed payments for this agreement.
		 */
		FailedPaymentCount string `json:"failed_payment_count,omitempty"`
	}


	// Address struct
	Address struct {
		Line1       string `json:"line1"`
		Line2       string `json:"line2,omitempty"`
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		PostalCode  string `json:"postal_code,omitempty"`
		State       string `json:"state,omitempty"`
		Phone       string `json:"phone,omitempty"`
	}

	// Amount struct
	Amount struct {
		Currency string `json:"currency"`
		Total    string `json:"total"`
	}

	// AmountPayout struct
	AmountPayout struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	}

	// Authorization struct
	Authorization struct {
		Amount                    *Amount    `json:"amount,omitempty"`
		CreateTime                *time.Time `json:"create_time,omitempty"`
		UpdateTime                *time.Time `json:"update_time,omitempty"`
		State                     string     `json:"state,omitempty"`
		ParentPayment             string     `json:"parent_payment,omitempty"`
		ID                        string     `json:"id,omitempty"`
		ValidUntil                *time.Time `json:"valid_until,omitempty"`
		Links                     []Link     `json:"links,omitempty"`
		ClearingTime              string     `json:"clearing_time,omitempty"`
		ProtectionEligibility     string     `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType string     `json:"protection_eligibility_type,omitempty"`
	}

	// BatchHeader struct
	BatchHeader struct {
		Amount            *AmountPayout      `json:"amount,omitempty"`
		Fees              *AmountPayout      `json:"fees,omitempty"`
		PayoutBatchID     string             `json:"payout_batch_id,omitempty"`
		BatchStatus       string             `json:"batch_status,omitempty"`
		TimeCreated       *time.Time         `json:"time_created,omitempty"`
		TimeCompleted     *time.Time         `json:"time_completed,omitempty"`
		SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header,omitempty"`
	}

	// Capture struct
	Capture struct {
		Amount         *Amount    `json:"amount,omitempty"`
		IsFinalCapture bool       `json:"is_final_capture"`
		CreateTime     *time.Time `json:"create_time,omitempty"`
		UpdateTime     *time.Time `json:"update_time,omitempty"`
		State          string     `json:"state,omitempty"`
		ParentPayment  string     `json:"parent_payment,omitempty"`
		ID             string     `json:"id,omitempty"`
		Links          []Link     `json:"links,omitempty"`
	}

	ChargeModels struct {
		/**
		 * Identifier of the charge model. 128 characters max.
		 */
		Id     string `json:"id,omitempty"`

		/**
		 * Type of charge model. Allowed values: `SHIPPING`, `TAX`.
		 */
		Type   string `json:"type,omitempty"`

		/**
		 * Specific amount for this charge model.
		 */
		Amount AmountPayout `json:"amount,omitempty"`
	}

	// Client represents a Paypal REST API Client
	Client struct {
		client   *http.Client
		ClientID string
		Secret   string
		APIBase  string
		Log      io.Writer // If user set log file name all requests will be logged there
		Token    *TokenResponse
	}

	// CreditCard struct
	CreditCard struct {
		ID             string   `json:"id,omitempty"`
		PayerID        string   `json:"payer_id,omitempty"`
		Number         string   `json:"number"`
		Type           string   `json:"type"`
		ExpireMonth    string   `json:"expire_month"`
		ExpireYear     string   `json:"expire_year"`
		CVV2           string   `json:"cvv2,omitempty"`
		FirstName      string   `json:"first_name,omitempty"`
		LastName       string   `json:"last_name,omitempty"`
		BillingAddress *Address `json:"billing_address,omitempty"`
		State          string   `json:"state,omitempty"`
		ValidUntil     string   `json:"valid_until,omitempty"`
	}

	// CreditCards GET /v1/vault/credit-cards
	CreditCards struct {
		Items      []CreditCard `json:"items"`
		Links      []Link       `json:"links"`
		TotalItems int          `json:"total_items"`
		TotalPages int          `json:"total_pages"`
	}

	// CreditCardToken struct
	CreditCardToken struct {
		CreditCardID string `json:"credit_card_id"`
		PayerID      string `json:"payer_id,omitempty"`
		Last4        string `json:"last4,omitempty"`
		ExpireYear   string `json:"expire_year,omitempty"`
		ExpireMonth  string `json:"expire_month,omitempty"`
	}

	// CreditCardsFilter struct
	CreditCardsFilter struct {
		PageSize int
		Page     int
	}

	// CreditCardField PATCH /v1/vault/credit-cards/credit_card_id
	CreditCardField struct {
		Operation string `json:"op"`
		Path      string `json:"path"`
		Value     string `json:"value"`
	}

	// Currency struct
	Currency struct {
		Currency string `json:"currency,omitempty"`
		Value    string `json:"value,omitempty"`
	}

	// ErrorResponse https://developer.paypal.com/docs/api/errors/
	ErrorResponse struct {
		Response        *http.Response `json:"-"`
		Name            string         `json:"name"`
		DebugID         string         `json:"debug_id"`
		Message         string         `json:"message"`
		InformationLink string         `json:"information_link"`
		Details         string         `json:"details"`
	}

	// ExecuteResponse struct
	ExecuteResponse struct {
		ID           string        `json:"id"`
		Links        []Link        `json:"links"`
		State        string        `json:"state"`
		Transactions []Transaction `json:"transactions,omitempty"`
	}

	// FundingInstrument struct
	FundingInstrument struct {
		CreditCard      *CreditCard      `json:"credit_card,omitempty"`
		CreditCardToken *CreditCardToken `json:"credit_card_token,omitempty"`
	}

	// Item struct
	Item struct {
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Price       string `json:"price"`
		Currency    string `json:"currency"`
		SKU         string `json:"sku,omitempty"`
		Description string `json:"description,omitempty"`
		Tax         string `json:"tax,omitempty"`
	}

	// ItemList struct
	ItemList struct {
		Items           []Item           `json:"items,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	}

	// Link struct
	Link struct {
		Href    string `json:"href"`
		Rel     string `json:"rel,omitempty"`
		Method  string `json:"method,omitempty"`
		Enctype string `json:"enctype,omitempty"`
	}

	MerchantPreferences struct {
		/**
	 * Identifier of the merchant_preferences. 128 characters max.
	 */
		Id                      string `json:"id"`

		/**
		 * Setup fee amount. Default is 0.
		 */
		SetupFee                AmountPayout `json:"setup_fee"`

		/**
		 * Redirect URL on cancellation of agreement request. 1000 characters max.
		 */
		CancelUrl               string `json:"cancel_url"`

		/**
		 * Redirect URL on creation of agreement request. 1000 characters max.
		 */
		ReturnUrl               string `json:"return_url"`

		/**
		 * Notify URL on agreement creation. 1000 characters max.
		 */
		NotifyUrl               string `json:"notify_url"`

		/**
		 * Total number of failed attempts allowed. Default is 0, representing an infinite number of failed attempts.
		 */
		MaxFailAttempts         string `json:"max_fail_attempts"`

		/**
		 * Allow auto billing for the outstanding amount of the agreement in the next cycle. Allowed values: `YES`, `NO`. Default is `NO`.
		 */
		AutoBillAmount          string `json:"auto_bill_amount"`

		/**
		 * Action to take if a failure occurs during initial payment. Allowed values: `CONTINUE`, `CANCEL`. Default is continue.
		 */
		InitialFailAmountAction string `json:"initial_fail_amount_action"`

		/**
		 * Payment types that are accepted for this plan.
		 */
		AcceptedPaymentType     string `json:"accepted_payment_type"`

		/**
		 * char_set for this plan.
		 */
		CharSet                 string `json:"charset"`
	}

	// Order struct
	Order struct {
		ID            string     `json:"id,omitempty"`
		CreateTime    *time.Time `json:"create_time,omitempty"`
		UpdateTime    *time.Time `json:"update_time,omitempty"`
		State         string     `json:"state,omitempty"`
		Amount        *Amount    `json:"amount,omitempty"`
		PendingReason string     `json:"pending_reason,omitempty"`
		ParentPayment string     `json:"parent_payment,omitempty"`
		Links         []Link     `json:"links,omitempty"`
	}

	OverrideChargeModel struct {
		/**
		 * ID of charge model.
		 */
		ChargeId string `json:"charge_id,omitempty"`

		/**
		 * Updated Amount to be associated with this charge model.
		 */
		Currency AmountPayout `json:"currency,omitempty"`
	}


	// Payer struct
	Payer struct {
		PaymentMethod      string              `json:"payment_method"`
		FundingInstruments []FundingInstrument `json:"funding_instruments,omitempty"`
		PayerInfo          *PayerInfo          `json:"payer_info,omitempty"`
		Status             string              `json:"payer_status,omitempty"`
	}

	// PayerInfo struct
	PayerInfo struct {
		Email           string           `json:"email,omitempty"`
		FirstName       string           `json:"first_name,omitempty"`
		LastName        string           `json:"last_name,omitempty"`
		PayerID         string           `json:"payer_id,omitempty"`
		Phone           string           `json:"phone,omitempty"`
		ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
		TaxIDType       string           `json:"tax_id_type,omitempty"`
		TaxID           string           `json:"tax_id,omitempty"`
	}

	// Payment struct
	Payment struct {
		Intent              string        `json:"intent"`
		Payer               *Payer        `json:"payer"`
		Transactions        []Transaction `json:"transactions"`
		RedirectURLs        *RedirectURLs `json:"redirect_urls,omitempty"`
		ID                  string        `json:"id,omitempty"`
		CreateTime          *time.Time    `json:"create_time,omitempty"`
		State               string        `json:"state,omitempty"`
		UpdateTime          *time.Time    `json:"update_time,omitempty"`
		ExperienceProfileID string        `json:"experience_profile_id,omitempty"`
	}

	PaymentDefinition struct {
		/**
		 * Identifier of the payment_definition. 128 characters max.
		 */
		Id                string `json:"id,omitempty"`

		/**
		 * Name of the payment definition. 128 characters max.
		 */
		Name              string `json:"name,omitempty"`

		/**
		 * Type of the payment definition. Allowed values: `TRIAL`, `REGULAR`.
		 */
		Type              string `json:"type,omitempty"`

		/**
		 * How frequently the customer should be charged.
		 */
		FrequencyInterval string `json:"frequency_interval,omitempty"`

		/**
		 * Frequency of the payment definition offered. Allowed values: `WEEK`, `DAY`, `YEAR`, `MONTH`.
		 */
		Frequency         string `json:"frequency,omitempty"`

		/**
		 * Number of cycles in this payment definition.
		 */
		Cycles            string `json:"cycles,omitempty"`

		/**
		 * Amount that will be charged at the end of each cycle for this payment definition.
		 */
		amount            AmountPayout `json:"amount_payout,omitempty"`

		/**
		 * Array of charge_models for this payment definition.
		 */
		ChargeModels      []ChargeModels `json:"charge_models,omitempty"`
	}

	// PaymentResponse structure
	PaymentResponse struct {
		ID    string `json:"id"`
		Links []Link `json:"links"`
	}

	// Payout struct
	Payout struct {
		SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header"`
		Items             []PayoutItem       `json:"items"`
	}

	// PayoutItem struct
	PayoutItem struct {
		RecipientType string        `json:"recipient_type"`
		Receiver      string        `json:"receiver"`
		Amount        *AmountPayout `json:"amount"`
		Note          string        `json:"note,omitempty"`
		SenderItemID  string        `json:"sender_item_id,omitempty"`
	}

	// PayoutItemResponse struct
	PayoutItemResponse struct {
		PayoutItemID      string        `json:"payout_item_id"`
		TransactionID     string        `json:"transaction_id"`
		TransactionStatus string        `json:"transaction_status"`
		PayoutBatchID     string        `json:"payout_batch_id,omitempty"`
		PayoutItemFee     *AmountPayout `json:"payout_item_fee,omitempty"`
		PayoutItem        *PayoutItem   `json:"payout_item"`
		TimeProcessed     *time.Time    `json:"time_processed,omitempty"`
		Links             []Link        `json:"links"`
	}

	// PayoutResponse struct
	PayoutResponse struct {
		BatchHeader *BatchHeader         `json:"batch_header"`
		Items       []PayoutItemResponse `json:"items"`
		Links       []Link               `json:"links"`
	}

	Plan struct {
		/**
		 * Identifier of the billing plan. 128 characters max.
		 */
		Id                  string `json:"id,omitempty"`

		/**
		 * Name of the billing plan. 128 characters max.
		 */
		Name                string `json:"name,omitempty"`

		/**
		 * Description of the billing plan. 128 characters max.
		 */
		Description         string `json:"fescription,omitempty"`

		/**
		 * Type of the billing plan. Allowed values: `FIXED`, `INFINITE`.
		 */
		Type                string `json:"type,omitempty"`

		/**
		 * Status of the billing plan. Allowed values: `CREATED`, `ACTIVE`, `INACTIVE`, and `DELETED`.
		 */
		State               string `json:"state,omitempty"`

		/**
		 * Time when the billing plan was created. Format YYYY-MM-DDTimeTimezone, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		 */
		CreateTime          string `json:"create_time,omitempty"`

		/**
		 * Time when this billing plan was updated. Format YYYY-MM-DDTimeTimezone, as defined in [ISO8601](http://tools.ietf.org/html/rfc3339#section-5.6).
		 */
		UpdateTime          string `json:"update_time,omitempty"`

		/**
		 * Array of payment definitions for this billing plan.
		 */
		PaymentDefinitions  []PaymentDefinition `json:"payment_definitions,omitempty"`

		/**
		 * Array of terms for this billing plan.
		 */
		Terms               []Terms `json:"terms,omitempty"`

		/**
		 * Specific preferences such as: set up fee, max fail attempts, autobill amount, and others that are configured for this billing plan.
		 */
		MerchantPreferences MerchantPreferences `json:"merchant_preferences,omitempty"`

		/**
		 *
		 */
		Links               []Link `json:"links,omitempty"`
	}

	// RedirectURLs struct
	RedirectURLs struct {
		ReturnURL string `json:"return_url,omitempty"`
		CancelURL string `json:"cancel_url,omitempty"`
	}

	// Refund struct
	Refund struct {
		ID            string     `json:"id,omitempty"`
		Amount        *Amount    `json:"amount,omitempty"`
		CreateTime    *time.Time `json:"create_time,omitempty"`
		State         string     `json:"state,omitempty"`
		CaptureID     string     `json:"capture_id,omitempty"`
		ParentPayment string     `json:"parent_payment,omitempty"`
		UpdateTime    *time.Time `json:"update_time,omitempty"`
	}

	// Related struct
	Related struct {
		Sale          *Sale          `json:"sale,omitempty"`
		Authorization *Authorization `json:"authorization,omitempty"`
		Order         *Order         `json:"order,omitempty"`
		Capture       *Capture       `json:"capture,omitempty"`
		Refund        *Refund        `json:"refund,omitempty"`
	}

	// Sale struct
	Sale struct {
		ID                        string     `json:"id,omitempty"`
		Amount                    *Amount    `json:"amount,omitempty"`
		Description               string     `json:"description,omitempty"`
		CreateTime                *time.Time `json:"create_time,omitempty"`
		State                     string     `json:"state,omitempty"`
		ParentPayment             string     `json:"parent_payment,omitempty"`
		UpdateTime                *time.Time `json:"update_time,omitempty"`
		PaymentMode               string     `json:"payment_mode,omitempty"`
		PendingReason             string     `json:"pending_reason,omitempty"`
		ReasonCode                string     `json:"reason_code,omitempty"`
		ClearingTime              string     `json:"clearing_time,omitempty"`
		ProtectionEligibility     string     `json:"protection_eligibility,omitempty"`
		ProtectionEligibilityType string     `json:"protection_eligibility_type,omitempty"`
		Links                     []Link     `json:"links,omitempty"`
	}

	// SenderBatchHeader struct
	SenderBatchHeader struct {
		EmailSubject string `json:"email_subject"`
	}

	// ShippingAddress struct
	ShippingAddress struct {
		RecipientName string `json:"recipient_name,omitempty"`
		Type          string `json:"type,omitempty"`
		Line1         string `json:"line1"`
		Line2         string `json:"line2,omitempty"`
		City          string `json:"city"`
		CountryCode   string `json:"country_code"`
		PostalCode    string `json:"postal_code,omitempty"`
		State         string `json:"state,omitempty"`
		Phone         string `json:"phone,omitempty"`
	}


	Terms struct {
		/**
		 * Identifier of the terms. 128 characters max.
		 */
		Id               string `json:"id,omitempty"`

		/**
		 * Term type. Allowed values: `MONTHLY`, `WEEKLY`, `YEARLY`.
		 */
		Type             string `json:"type,omitempty"`

		/**
		 * Max Amount associated with this term.
		 */
		MaxBillingAmount Amount `json:"max_billing_amount,omitempty"`

		/**
		 * How many times money can be pulled during this term.
		 */
		Occurrences      string `json:"occurrences,omitempty"`

		/**
		 * Amount_range associated with this term.
		 */
		AmountRange      AmountPayout `json:"amount_range,omitempty"`

		/**
		 * Buyer's ability to edit the amount in this term.
		 */
		BuyerEditable    string `json:"buyer_editable,omitempty"`
	}

	// TokenResponse is for API response for the /oauth2/token endpoint
	TokenResponse struct {
		RefreshToken string `json:"refresh_token"`
		Token        string `json:"access_token"`
		Type         string `json:"token_type"`
		ExpiresIn    int64  `json:"expires_in"`
	}

	// Transaction struct
	Transaction struct {
		Amount           *Amount   `json:"amount"`
		Description      string    `json:"description,omitempty"`
		ItemList         *ItemList `json:"item_list,omitempty"`
		InvoiceNumber    string    `json:"invoice_number,omitempty"`
		Custom           string    `json:"custom,omitempty"`
		SoftDescriptor   string    `json:"soft_descriptor,omitempty"`
		RelatedResources []Related `json:"related_resources,omitempty"`
	}

	// UserInfo struct
	UserInfo struct {
		ID              string   `json:"user_id"`
		Name            string   `json:"name"`
		GivenName       string   `json:"given_name"`
		FamilyName      string   `json:"family_name"`
		Email           string   `json:"email"`
		Verified        bool     `json:"verified,omitempty"`
		Gender          string   `json:"gender,omitempty"`
		BirthDate       string   `json:"birthdate,omitempty"`
		ZoneInfo        string   `json:"zoneinfo,omitempty"`
		Locale          string   `json:"locale,omitempty"`
		Phone           string   `json:"phone_number,omitempty"`
		Address         *Address `json:"address,omitempty"`
		VerifiedAccount bool     `json:"verified_account,omitempty"`
		AccountType     string   `json:"account_type,omitempty"`
		AgeRange        string   `json:"age_range,omitempty"`
		PayerID         string   `json:"payer_id,omitempty"`
	}

	// WebProfile represents the configuration of the payment web payment experience
	//
	// https://developer.paypal.com/docs/api/payment-experience/
	WebProfile struct {
		ID           string       `json:"id,omitempty"`
		Name         string       `json:"name"`
		Presentation Presentation `json:"presentation,omitempty"`
		InputFields  InputFields  `json:"input_fields,omitempty"`
		FlowConfig   FlowConfig   `json:"flow_config,omitempty"`
	}

	// Presentation represents the branding and locale that a customer sees on
	// redirect payments
	//
	// https://developer.paypal.com/docs/api/payment-experience/#definition-presentation
	Presentation struct {
		BrandName  string `json:"brand_name,omitempty"`
		LogoImage  string `json:"logo_image,omitempty"`
		LocaleCode string `json:"locale_code,omitempty"`
	}

	// InputFields represents the fields that are displayed to a customer on
	// redirect payments
	//
	// https://developer.paypal.com/docs/api/payment-experience/#definition-input_fields
	InputFields struct {
		AllowNote       bool `json:"allow_note,omitempty"`
		NoShipping      uint `json:"no_shipping,omitempty"`
		AddressOverride uint `json:"address_override,omitempty"`
	}

	// FlowConfig represents the general behaviour of redirect payment pages
	//
	// https://developer.paypal.com/docs/api/payment-experience/#definition-flow_config
	FlowConfig struct {
		LandingPageType   string `json:"landing_page_type,omitempty"`
		BankTXNPendingURL string `json:"bank_txn_pending_url,omitempty"`
		UserAction        string `json:"user_action,omitempty"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
