# NewCouponsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfCoupons** | Pointer to **int32** | The number of new coupon codes to generate for the campaign. Must be at least 1. | 
**UniquePrefix** | Pointer to **string** | **DEPRECATED** To create more than 20,000 coupons in one request, use [Create coupons asynchronously](https://docs.talon.one/management-api#operation/createCouponsAsync) endpoint.  | [optional] 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this item. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile. | [optional] 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**CouponPattern** | Pointer to **string** | The pattern used to generate coupon codes. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 

## Methods

### GetNumberOfCoupons

`func (o *NewCouponsAllOf) GetNumberOfCoupons() int32`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *NewCouponsAllOf) GetNumberOfCouponsOk() (int32, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCoupons

`func (o *NewCouponsAllOf) HasNumberOfCoupons() bool`

HasNumberOfCoupons returns a boolean if a field has been set.

### SetNumberOfCoupons

`func (o *NewCouponsAllOf) SetNumberOfCoupons(v int32)`

SetNumberOfCoupons gets a reference to the given int32 and assigns it to the NumberOfCoupons field.

### GetUniquePrefix

`func (o *NewCouponsAllOf) GetUniquePrefix() string`

GetUniquePrefix returns the UniquePrefix field if non-nil, zero value otherwise.

### GetUniquePrefixOk

`func (o *NewCouponsAllOf) GetUniquePrefixOk() (string, bool)`

GetUniquePrefixOk returns a tuple with the UniquePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniquePrefix

`func (o *NewCouponsAllOf) HasUniquePrefix() bool`

HasUniquePrefix returns a boolean if a field has been set.

### SetUniquePrefix

`func (o *NewCouponsAllOf) SetUniquePrefix(v string)`

SetUniquePrefix gets a reference to the given string and assigns it to the UniquePrefix field.

### GetAttributes

`func (o *NewCouponsAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCouponsAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCouponsAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCouponsAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetRecipientIntegrationId

`func (o *NewCouponsAllOf) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *NewCouponsAllOf) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *NewCouponsAllOf) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *NewCouponsAllOf) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetValidCharacters

`func (o *NewCouponsAllOf) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewCouponsAllOf) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewCouponsAllOf) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewCouponsAllOf) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetCouponPattern

`func (o *NewCouponsAllOf) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewCouponsAllOf) GetCouponPatternOk() (string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponPattern

`func (o *NewCouponsAllOf) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.

### SetCouponPattern

`func (o *NewCouponsAllOf) SetCouponPattern(v string)`

SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.

### GetIsReservationMandatory

`func (o *NewCouponsAllOf) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *NewCouponsAllOf) GetIsReservationMandatoryOk() (bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReservationMandatory

`func (o *NewCouponsAllOf) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### SetIsReservationMandatory

`func (o *NewCouponsAllOf) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.

### GetImplicitlyReserved

`func (o *NewCouponsAllOf) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *NewCouponsAllOf) GetImplicitlyReservedOk() (bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImplicitlyReserved

`func (o *NewCouponsAllOf) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.

### SetImplicitlyReserved

`func (o *NewCouponsAllOf) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


