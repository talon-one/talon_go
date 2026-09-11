# NewIntegrationHubCoupons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**ApplicationId** | Pointer to **int64** | The ID of the Application the coupons will belong to. | 
**CampaignId** | Pointer to **int64** | The ID of the Campaign the coupons will belong to. | 
**BatchId** | Pointer to **string** | An identifier for the batch of coupons being created. | 
**NumberOfCoupons** | Pointer to **int64** | The number of new coupon codes to generate for the campaign. Must be at least 1. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**CouponPattern** | Pointer to **string** | The pattern used to generate coupon codes. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile. | [optional] 
**SupportRequestId** | Pointer to **int64** | The identifier of the support request to link to the coupon creation. The request must exist and not yet be processed. | [optional] 
**SupportRequestNote** | Pointer to **string** | A note recorded when the linked support request is approved or rejected. Applied when &#x60;supportRequestId&#x60; is provided. | [optional] 

## Methods

### NewNewIntegrationHubCoupons

`func NewNewIntegrationHubCoupons(usageLimit int64, applicationId int64, campaignId int64, batchId string, numberOfCoupons int64, ) *NewIntegrationHubCoupons`

NewNewIntegrationHubCoupons instantiates a new NewIntegrationHubCoupons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewIntegrationHubCouponsWithDefaults

`func NewNewIntegrationHubCouponsWithDefaults() *NewIntegrationHubCoupons`

NewNewIntegrationHubCouponsWithDefaults instantiates a new NewIntegrationHubCoupons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageLimit

`func (o *NewIntegrationHubCoupons) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewIntegrationHubCoupons) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *NewIntegrationHubCoupons) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *NewIntegrationHubCoupons) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *NewIntegrationHubCoupons) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *NewIntegrationHubCoupons) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *NewIntegrationHubCoupons) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *NewIntegrationHubCoupons) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *NewIntegrationHubCoupons) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *NewIntegrationHubCoupons) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *NewIntegrationHubCoupons) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *NewIntegrationHubCoupons) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewIntegrationHubCoupons) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NewIntegrationHubCoupons) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NewIntegrationHubCoupons) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *NewIntegrationHubCoupons) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewIntegrationHubCoupons) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewIntegrationHubCoupons) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewIntegrationHubCoupons) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetLimits

`func (o *NewIntegrationHubCoupons) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewIntegrationHubCoupons) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *NewIntegrationHubCoupons) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *NewIntegrationHubCoupons) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetApplicationId

`func (o *NewIntegrationHubCoupons) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NewIntegrationHubCoupons) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NewIntegrationHubCoupons) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCampaignId

`func (o *NewIntegrationHubCoupons) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *NewIntegrationHubCoupons) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *NewIntegrationHubCoupons) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetBatchId

`func (o *NewIntegrationHubCoupons) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *NewIntegrationHubCoupons) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *NewIntegrationHubCoupons) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetNumberOfCoupons

`func (o *NewIntegrationHubCoupons) GetNumberOfCoupons() int64`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *NewIntegrationHubCoupons) GetNumberOfCouponsOk() (*int64, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCoupons

`func (o *NewIntegrationHubCoupons) SetNumberOfCoupons(v int64)`

SetNumberOfCoupons sets NumberOfCoupons field to given value.


### GetAttributes

`func (o *NewIntegrationHubCoupons) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewIntegrationHubCoupons) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewIntegrationHubCoupons) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewIntegrationHubCoupons) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetValidCharacters

`func (o *NewIntegrationHubCoupons) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewIntegrationHubCoupons) GetValidCharactersOk() (*[]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *NewIntegrationHubCoupons) SetValidCharacters(v []string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *NewIntegrationHubCoupons) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetCouponPattern

`func (o *NewIntegrationHubCoupons) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewIntegrationHubCoupons) GetCouponPatternOk() (*string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponPattern

`func (o *NewIntegrationHubCoupons) SetCouponPattern(v string)`

SetCouponPattern sets CouponPattern field to given value.

### HasCouponPattern

`func (o *NewIntegrationHubCoupons) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.

### GetIsReservationMandatory

`func (o *NewIntegrationHubCoupons) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *NewIntegrationHubCoupons) GetIsReservationMandatoryOk() (*bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReservationMandatory

`func (o *NewIntegrationHubCoupons) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory sets IsReservationMandatory field to given value.

### HasIsReservationMandatory

`func (o *NewIntegrationHubCoupons) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### GetImplicitlyReserved

`func (o *NewIntegrationHubCoupons) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *NewIntegrationHubCoupons) GetImplicitlyReservedOk() (*bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitlyReserved

`func (o *NewIntegrationHubCoupons) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved sets ImplicitlyReserved field to given value.

### HasImplicitlyReserved

`func (o *NewIntegrationHubCoupons) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *NewIntegrationHubCoupons) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *NewIntegrationHubCoupons) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *NewIntegrationHubCoupons) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *NewIntegrationHubCoupons) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetSupportRequestId

`func (o *NewIntegrationHubCoupons) GetSupportRequestId() int64`

GetSupportRequestId returns the SupportRequestId field if non-nil, zero value otherwise.

### GetSupportRequestIdOk

`func (o *NewIntegrationHubCoupons) GetSupportRequestIdOk() (*int64, bool)`

GetSupportRequestIdOk returns a tuple with the SupportRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestId

`func (o *NewIntegrationHubCoupons) SetSupportRequestId(v int64)`

SetSupportRequestId sets SupportRequestId field to given value.

### HasSupportRequestId

`func (o *NewIntegrationHubCoupons) HasSupportRequestId() bool`

HasSupportRequestId returns a boolean if a field has been set.

### GetSupportRequestNote

`func (o *NewIntegrationHubCoupons) GetSupportRequestNote() string`

GetSupportRequestNote returns the SupportRequestNote field if non-nil, zero value otherwise.

### GetSupportRequestNoteOk

`func (o *NewIntegrationHubCoupons) GetSupportRequestNoteOk() (*string, bool)`

GetSupportRequestNoteOk returns a tuple with the SupportRequestNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestNote

`func (o *NewIntegrationHubCoupons) SetSupportRequestNote(v string)`

SetSupportRequestNote sets SupportRequestNote field to given value.

### HasSupportRequestNote

`func (o *NewIntegrationHubCoupons) HasSupportRequestNote() bool`

HasSupportRequestNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


