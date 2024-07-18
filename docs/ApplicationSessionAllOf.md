# ApplicationSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profileintegrationid** | Pointer to **string** | Integration ID of the customer for the session. | [optional] 
**Coupon** | Pointer to **string** | Any coupon code entered. | 
**Referral** | Pointer to **string** | Any referral code entered. | 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60; or &#x60;partially_returned&#x60; 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | 
**Discounts** | Pointer to **map[string]float32** | **API V1 only.** A map of labeled discount values, in the same currency as the session.  If you are using the V2 endpoints, refer to the &#x60;totalDiscounts&#x60; property instead.  | 
**TotalDiscounts** | Pointer to **float32** | The total sum of the discounts applied to this session. | 
**Total** | Pointer to **float32** | The total sum of the session before any discounts applied. | 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this item. | [optional] 

## Methods

### GetProfileintegrationid

`func (o *ApplicationSessionAllOf) GetProfileintegrationid() string`

GetProfileintegrationid returns the Profileintegrationid field if non-nil, zero value otherwise.

### GetProfileintegrationidOk

`func (o *ApplicationSessionAllOf) GetProfileintegrationidOk() (string, bool)`

GetProfileintegrationidOk returns a tuple with the Profileintegrationid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileintegrationid

`func (o *ApplicationSessionAllOf) HasProfileintegrationid() bool`

HasProfileintegrationid returns a boolean if a field has been set.

### SetProfileintegrationid

`func (o *ApplicationSessionAllOf) SetProfileintegrationid(v string)`

SetProfileintegrationid gets a reference to the given string and assigns it to the Profileintegrationid field.

### GetCoupon

`func (o *ApplicationSessionAllOf) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *ApplicationSessionAllOf) GetCouponOk() (string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupon

`func (o *ApplicationSessionAllOf) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### SetCoupon

`func (o *ApplicationSessionAllOf) SetCoupon(v string)`

SetCoupon gets a reference to the given string and assigns it to the Coupon field.

### GetReferral

`func (o *ApplicationSessionAllOf) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *ApplicationSessionAllOf) GetReferralOk() (string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferral

`func (o *ApplicationSessionAllOf) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### SetReferral

`func (o *ApplicationSessionAllOf) SetReferral(v string)`

SetReferral gets a reference to the given string and assigns it to the Referral field.

### GetState

`func (o *ApplicationSessionAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationSessionAllOf) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *ApplicationSessionAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *ApplicationSessionAllOf) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetCartItems

`func (o *ApplicationSessionAllOf) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *ApplicationSessionAllOf) GetCartItemsOk() ([]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItems

`func (o *ApplicationSessionAllOf) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItems

`func (o *ApplicationSessionAllOf) SetCartItems(v []CartItem)`

SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.

### GetDiscounts

`func (o *ApplicationSessionAllOf) GetDiscounts() map[string]float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *ApplicationSessionAllOf) GetDiscountsOk() (map[string]float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscounts

`func (o *ApplicationSessionAllOf) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscounts

`func (o *ApplicationSessionAllOf) SetDiscounts(v map[string]float32)`

SetDiscounts gets a reference to the given map[string]float32 and assigns it to the Discounts field.

### GetTotalDiscounts

`func (o *ApplicationSessionAllOf) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ApplicationSessionAllOf) GetTotalDiscountsOk() (float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscounts

`func (o *ApplicationSessionAllOf) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### SetTotalDiscounts

`func (o *ApplicationSessionAllOf) SetTotalDiscounts(v float32)`

SetTotalDiscounts gets a reference to the given float32 and assigns it to the TotalDiscounts field.

### GetTotal

`func (o *ApplicationSessionAllOf) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApplicationSessionAllOf) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *ApplicationSessionAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *ApplicationSessionAllOf) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetAttributes

`func (o *ApplicationSessionAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationSessionAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationSessionAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationSessionAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


