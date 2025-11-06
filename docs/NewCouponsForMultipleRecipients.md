# NewCouponsForMultipleRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**RecipientsIntegrationIds** | Pointer to **[]string** | The integration IDs for recipients. | 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**CouponPattern** | Pointer to **string** | The pattern used to generate coupon codes. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### NewNewCouponsForMultipleRecipients

`func NewNewCouponsForMultipleRecipients(usageLimit int64, recipientsIntegrationIds []string, ) *NewCouponsForMultipleRecipients`

NewNewCouponsForMultipleRecipients instantiates a new NewCouponsForMultipleRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCouponsForMultipleRecipientsWithDefaults

`func NewNewCouponsForMultipleRecipientsWithDefaults() *NewCouponsForMultipleRecipients`

NewNewCouponsForMultipleRecipientsWithDefaults instantiates a new NewCouponsForMultipleRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageLimit

`func (o *NewCouponsForMultipleRecipients) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewCouponsForMultipleRecipients) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *NewCouponsForMultipleRecipients) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *NewCouponsForMultipleRecipients) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *NewCouponsForMultipleRecipients) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *NewCouponsForMultipleRecipients) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *NewCouponsForMultipleRecipients) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *NewCouponsForMultipleRecipients) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *NewCouponsForMultipleRecipients) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *NewCouponsForMultipleRecipients) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *NewCouponsForMultipleRecipients) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *NewCouponsForMultipleRecipients) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewCouponsForMultipleRecipients) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NewCouponsForMultipleRecipients) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NewCouponsForMultipleRecipients) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *NewCouponsForMultipleRecipients) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewCouponsForMultipleRecipients) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewCouponsForMultipleRecipients) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewCouponsForMultipleRecipients) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetAttributes

`func (o *NewCouponsForMultipleRecipients) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCouponsForMultipleRecipients) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewCouponsForMultipleRecipients) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewCouponsForMultipleRecipients) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipients) GetRecipientsIntegrationIds() []string`

GetRecipientsIntegrationIds returns the RecipientsIntegrationIds field if non-nil, zero value otherwise.

### GetRecipientsIntegrationIdsOk

`func (o *NewCouponsForMultipleRecipients) GetRecipientsIntegrationIdsOk() (*[]string, bool)`

GetRecipientsIntegrationIdsOk returns a tuple with the RecipientsIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipients) SetRecipientsIntegrationIds(v []string)`

SetRecipientsIntegrationIds sets RecipientsIntegrationIds field to given value.


### GetValidCharacters

`func (o *NewCouponsForMultipleRecipients) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewCouponsForMultipleRecipients) GetValidCharactersOk() (*[]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *NewCouponsForMultipleRecipients) SetValidCharacters(v []string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *NewCouponsForMultipleRecipients) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetCouponPattern

`func (o *NewCouponsForMultipleRecipients) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewCouponsForMultipleRecipients) GetCouponPatternOk() (*string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponPattern

`func (o *NewCouponsForMultipleRecipients) SetCouponPattern(v string)`

SetCouponPattern sets CouponPattern field to given value.

### HasCouponPattern

`func (o *NewCouponsForMultipleRecipients) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


