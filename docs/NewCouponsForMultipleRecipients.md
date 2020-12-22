# NewCouponsForMultipleRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The amount of discounts that can be given with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 
**RecipientsIntegrationIds** | Pointer to **[]string** | The integration IDs for recipients | 
**ValidCharacters** | Pointer to **[]string** | Set of characters to be used when generating random part of code. Defaults to [A-Z, 0-9] (in terms of RegExp). | [optional] 
**CouponPattern** | Pointer to **string** | The pattern that will be used to generate coupon codes. The character &#x60;#&#x60; acts as a placeholder and will be replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### GetUsageLimit

`func (o *NewCouponsForMultipleRecipients) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewCouponsForMultipleRecipients) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *NewCouponsForMultipleRecipients) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *NewCouponsForMultipleRecipients) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetDiscountLimit

`func (o *NewCouponsForMultipleRecipients) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *NewCouponsForMultipleRecipients) GetDiscountLimitOk() (float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountLimit

`func (o *NewCouponsForMultipleRecipients) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### SetDiscountLimit

`func (o *NewCouponsForMultipleRecipients) SetDiscountLimit(v float32)`

SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.

### GetStartDate

`func (o *NewCouponsForMultipleRecipients) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewCouponsForMultipleRecipients) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *NewCouponsForMultipleRecipients) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *NewCouponsForMultipleRecipients) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *NewCouponsForMultipleRecipients) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewCouponsForMultipleRecipients) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *NewCouponsForMultipleRecipients) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *NewCouponsForMultipleRecipients) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetAttributes

`func (o *NewCouponsForMultipleRecipients) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCouponsForMultipleRecipients) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCouponsForMultipleRecipients) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCouponsForMultipleRecipients) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipients) GetRecipientsIntegrationIds() []string`

GetRecipientsIntegrationIds returns the RecipientsIntegrationIds field if non-nil, zero value otherwise.

### GetRecipientsIntegrationIdsOk

`func (o *NewCouponsForMultipleRecipients) GetRecipientsIntegrationIdsOk() ([]string, bool)`

GetRecipientsIntegrationIdsOk returns a tuple with the RecipientsIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipients) HasRecipientsIntegrationIds() bool`

HasRecipientsIntegrationIds returns a boolean if a field has been set.

### SetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipients) SetRecipientsIntegrationIds(v []string)`

SetRecipientsIntegrationIds gets a reference to the given []string and assigns it to the RecipientsIntegrationIds field.

### GetValidCharacters

`func (o *NewCouponsForMultipleRecipients) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewCouponsForMultipleRecipients) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewCouponsForMultipleRecipients) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewCouponsForMultipleRecipients) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetCouponPattern

`func (o *NewCouponsForMultipleRecipients) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewCouponsForMultipleRecipients) GetCouponPatternOk() (string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponPattern

`func (o *NewCouponsForMultipleRecipients) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.

### SetCouponPattern

`func (o *NewCouponsForMultipleRecipients) SetCouponPattern(v string)`

SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


