# NewCustomerSessionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | [optional] 
**CouponCodes** | Pointer to **[]string** | Any coupon codes entered. | [optional] 
**ReferralCode** | Pointer to **string** | Any referral code entered. | [optional] 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;, after which valid transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60;  For more information, see [Entites](/docs/dev/concepts/entities#customer-session).  | [optional] [default to STATE_OPEN]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | All items the customer will be purchasing in this session | [optional] 
**AdditionalCosts** | Pointer to [**map[string]AdditionalCost**](AdditionalCost.md) | Any costs associated with the session that can not be explicitly attributed to cart items. Examples include shipping costs and service fees. | [optional] 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts.  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | [optional] 

## Methods

### GetProfileId

`func (o *NewCustomerSessionV2) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *NewCustomerSessionV2) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *NewCustomerSessionV2) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *NewCustomerSessionV2) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetCouponCodes

`func (o *NewCustomerSessionV2) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *NewCustomerSessionV2) GetCouponCodesOk() ([]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponCodes

`func (o *NewCustomerSessionV2) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### SetCouponCodes

`func (o *NewCustomerSessionV2) SetCouponCodes(v []string)`

SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.

### GetReferralCode

`func (o *NewCustomerSessionV2) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *NewCustomerSessionV2) GetReferralCodeOk() (string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCode

`func (o *NewCustomerSessionV2) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### SetReferralCode

`func (o *NewCustomerSessionV2) SetReferralCode(v string)`

SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.

### GetState

`func (o *NewCustomerSessionV2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewCustomerSessionV2) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *NewCustomerSessionV2) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *NewCustomerSessionV2) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetCartItems

`func (o *NewCustomerSessionV2) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *NewCustomerSessionV2) GetCartItemsOk() ([]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItems

`func (o *NewCustomerSessionV2) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItems

`func (o *NewCustomerSessionV2) SetCartItems(v []CartItem)`

SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.

### GetAdditionalCosts

`func (o *NewCustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *NewCustomerSessionV2) GetAdditionalCostsOk() (map[string]AdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCosts

`func (o *NewCustomerSessionV2) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### SetAdditionalCosts

`func (o *NewCustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost)`

SetAdditionalCosts gets a reference to the given map[string]AdditionalCost and assigns it to the AdditionalCosts field.

### GetIdentifiers

`func (o *NewCustomerSessionV2) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *NewCustomerSessionV2) GetIdentifiersOk() ([]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifiers

`func (o *NewCustomerSessionV2) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiers

`func (o *NewCustomerSessionV2) SetIdentifiers(v []string)`

SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.

### GetAttributes

`func (o *NewCustomerSessionV2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCustomerSessionV2) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCustomerSessionV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCustomerSessionV2) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


