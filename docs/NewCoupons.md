# NewCoupons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**ValidCharacters** | Pointer to **[]string** | Set of characters to be used when generating random part of code. Defaults to [A-Z, 0-9] (in terms of RegExp). | 
**CouponPattern** | Pointer to **string** | The pattern that will be used to generate coupon codes. The character &#x60;#&#x60; acts as a placeholder and will be replaced by a random character from the &#x60;validCharacters&#x60; set.  | 
**NumberOfCoupons** | Pointer to **int32** | The number of new coupon codes to generate for the campaign. Must be at least 1. | 
**UniquePrefix** | Pointer to **string** | A unique prefix to prepend to all generated coupons. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile | [optional] 

## Methods

### GetUsageLimit

`func (o *NewCoupons) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewCoupons) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *NewCoupons) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *NewCoupons) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetStartDate

`func (o *NewCoupons) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewCoupons) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *NewCoupons) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *NewCoupons) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *NewCoupons) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewCoupons) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *NewCoupons) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *NewCoupons) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetValidCharacters

`func (o *NewCoupons) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewCoupons) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewCoupons) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewCoupons) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetCouponPattern

`func (o *NewCoupons) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewCoupons) GetCouponPatternOk() (string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponPattern

`func (o *NewCoupons) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.

### SetCouponPattern

`func (o *NewCoupons) SetCouponPattern(v string)`

SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.

### GetNumberOfCoupons

`func (o *NewCoupons) GetNumberOfCoupons() int32`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *NewCoupons) GetNumberOfCouponsOk() (int32, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCoupons

`func (o *NewCoupons) HasNumberOfCoupons() bool`

HasNumberOfCoupons returns a boolean if a field has been set.

### SetNumberOfCoupons

`func (o *NewCoupons) SetNumberOfCoupons(v int32)`

SetNumberOfCoupons gets a reference to the given int32 and assigns it to the NumberOfCoupons field.

### GetUniquePrefix

`func (o *NewCoupons) GetUniquePrefix() string`

GetUniquePrefix returns the UniquePrefix field if non-nil, zero value otherwise.

### GetUniquePrefixOk

`func (o *NewCoupons) GetUniquePrefixOk() (string, bool)`

GetUniquePrefixOk returns a tuple with the UniquePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniquePrefix

`func (o *NewCoupons) HasUniquePrefix() bool`

HasUniquePrefix returns a boolean if a field has been set.

### SetUniquePrefix

`func (o *NewCoupons) SetUniquePrefix(v string)`

SetUniquePrefix gets a reference to the given string and assigns it to the UniquePrefix field.

### GetAttributes

`func (o *NewCoupons) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCoupons) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCoupons) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCoupons) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetRecipientIntegrationId

`func (o *NewCoupons) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *NewCoupons) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *NewCoupons) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *NewCoupons) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


