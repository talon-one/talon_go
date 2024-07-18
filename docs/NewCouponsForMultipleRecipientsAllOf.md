# NewCouponsForMultipleRecipientsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this item. | [optional] 
**RecipientsIntegrationIds** | Pointer to **[]string** | The integration IDs for recipients. | 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**CouponPattern** | Pointer to **string** | The pattern used to generate coupon codes. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### GetAttributes

`func (o *NewCouponsForMultipleRecipientsAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCouponsForMultipleRecipientsAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCouponsForMultipleRecipientsAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCouponsForMultipleRecipientsAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipientsAllOf) GetRecipientsIntegrationIds() []string`

GetRecipientsIntegrationIds returns the RecipientsIntegrationIds field if non-nil, zero value otherwise.

### GetRecipientsIntegrationIdsOk

`func (o *NewCouponsForMultipleRecipientsAllOf) GetRecipientsIntegrationIdsOk() ([]string, bool)`

GetRecipientsIntegrationIdsOk returns a tuple with the RecipientsIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipientsAllOf) HasRecipientsIntegrationIds() bool`

HasRecipientsIntegrationIds returns a boolean if a field has been set.

### SetRecipientsIntegrationIds

`func (o *NewCouponsForMultipleRecipientsAllOf) SetRecipientsIntegrationIds(v []string)`

SetRecipientsIntegrationIds gets a reference to the given []string and assigns it to the RecipientsIntegrationIds field.

### GetValidCharacters

`func (o *NewCouponsForMultipleRecipientsAllOf) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewCouponsForMultipleRecipientsAllOf) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewCouponsForMultipleRecipientsAllOf) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewCouponsForMultipleRecipientsAllOf) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetCouponPattern

`func (o *NewCouponsForMultipleRecipientsAllOf) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *NewCouponsForMultipleRecipientsAllOf) GetCouponPatternOk() (string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponPattern

`func (o *NewCouponsForMultipleRecipientsAllOf) HasCouponPattern() bool`

HasCouponPattern returns a boolean if a field has been set.

### SetCouponPattern

`func (o *NewCouponsForMultipleRecipientsAllOf) SetCouponPattern(v string)`

SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


