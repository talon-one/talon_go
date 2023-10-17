# NewCustomerSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**Coupon** | Pointer to **string** | Any coupon code entered. | [optional] 
**Referral** | Pointer to **string** | Any referral code entered. | [optional] 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60; or &#x60;partially_returned&#x60; 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | [optional] [default to STATE_OPEN]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | [optional] 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers).  | [optional] 
**Total** | Pointer to **float32** | The total sum of the cart in one session. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | [optional] 

## Methods

### GetProfileId

`func (o *NewCustomerSession) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *NewCustomerSession) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *NewCustomerSession) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *NewCustomerSession) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetCoupon

`func (o *NewCustomerSession) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *NewCustomerSession) GetCouponOk() (string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupon

`func (o *NewCustomerSession) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### SetCoupon

`func (o *NewCustomerSession) SetCoupon(v string)`

SetCoupon gets a reference to the given string and assigns it to the Coupon field.

### GetReferral

`func (o *NewCustomerSession) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *NewCustomerSession) GetReferralOk() (string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferral

`func (o *NewCustomerSession) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### SetReferral

`func (o *NewCustomerSession) SetReferral(v string)`

SetReferral gets a reference to the given string and assigns it to the Referral field.

### GetState

`func (o *NewCustomerSession) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewCustomerSession) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *NewCustomerSession) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *NewCustomerSession) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetCartItems

`func (o *NewCustomerSession) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *NewCustomerSession) GetCartItemsOk() ([]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItems

`func (o *NewCustomerSession) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItems

`func (o *NewCustomerSession) SetCartItems(v []CartItem)`

SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.

### GetIdentifiers

`func (o *NewCustomerSession) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *NewCustomerSession) GetIdentifiersOk() ([]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifiers

`func (o *NewCustomerSession) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiers

`func (o *NewCustomerSession) SetIdentifiers(v []string)`

SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.

### GetTotal

`func (o *NewCustomerSession) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NewCustomerSession) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *NewCustomerSession) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *NewCustomerSession) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetAttributes

`func (o *NewCustomerSession) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewCustomerSession) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewCustomerSession) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewCustomerSession) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


