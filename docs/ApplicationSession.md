# ApplicationSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | 
**Coupon** | Pointer to **string** | Any coupon code entered. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Discounts** | Pointer to **map[string]float32** | **API V1 only.** A map of labeled discount values, in the same currency as the session.  If you are using the V2 endpoints, refer to the &#x60;totalDiscounts&#x60; property instead.  | 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**ProfileId** | Pointer to **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**Profileintegrationid** | Pointer to **string** | Integration ID of the customer for the session. | [optional] 
**Referral** | Pointer to **string** | Any referral code entered. | 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60; or &#x60;partially_returned&#x60; 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**Total** | Pointer to **float32** | The total sum of the session before any discounts applied. | 
**TotalDiscounts** | Pointer to **float32** | The total sum of the discounts applied to this session. | 

## Methods

### GetApplicationId

`func (o *ApplicationSession) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationSession) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationSession) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationSession) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetAttributes

`func (o *ApplicationSession) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationSession) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationSession) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationSession) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetCartItems

`func (o *ApplicationSession) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *ApplicationSession) GetCartItemsOk() ([]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItems

`func (o *ApplicationSession) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItems

`func (o *ApplicationSession) SetCartItems(v []CartItem)`

SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.

### GetCoupon

`func (o *ApplicationSession) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *ApplicationSession) GetCouponOk() (string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupon

`func (o *ApplicationSession) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### SetCoupon

`func (o *ApplicationSession) SetCoupon(v string)`

SetCoupon gets a reference to the given string and assigns it to the Coupon field.

### GetCreated

`func (o *ApplicationSession) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationSession) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationSession) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationSession) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetDiscounts

`func (o *ApplicationSession) GetDiscounts() map[string]float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *ApplicationSession) GetDiscountsOk() (map[string]float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscounts

`func (o *ApplicationSession) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscounts

`func (o *ApplicationSession) SetDiscounts(v map[string]float32)`

SetDiscounts gets a reference to the given map[string]float32 and assigns it to the Discounts field.

### GetId

`func (o *ApplicationSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationSession) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationSession) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationSession) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetIntegrationId

`func (o *ApplicationSession) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ApplicationSession) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *ApplicationSession) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *ApplicationSession) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetProfileId

`func (o *ApplicationSession) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationSession) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *ApplicationSession) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *ApplicationSession) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetProfileintegrationid

`func (o *ApplicationSession) GetProfileintegrationid() string`

GetProfileintegrationid returns the Profileintegrationid field if non-nil, zero value otherwise.

### GetProfileintegrationidOk

`func (o *ApplicationSession) GetProfileintegrationidOk() (string, bool)`

GetProfileintegrationidOk returns a tuple with the Profileintegrationid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileintegrationid

`func (o *ApplicationSession) HasProfileintegrationid() bool`

HasProfileintegrationid returns a boolean if a field has been set.

### SetProfileintegrationid

`func (o *ApplicationSession) SetProfileintegrationid(v string)`

SetProfileintegrationid gets a reference to the given string and assigns it to the Profileintegrationid field.

### GetReferral

`func (o *ApplicationSession) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *ApplicationSession) GetReferralOk() (string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferral

`func (o *ApplicationSession) HasReferral() bool`

HasReferral returns a boolean if a field has been set.

### SetReferral

`func (o *ApplicationSession) SetReferral(v string)`

SetReferral gets a reference to the given string and assigns it to the Referral field.

### GetState

`func (o *ApplicationSession) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationSession) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *ApplicationSession) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *ApplicationSession) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetStoreIntegrationId

`func (o *ApplicationSession) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *ApplicationSession) GetStoreIntegrationIdOk() (string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoreIntegrationId

`func (o *ApplicationSession) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### SetStoreIntegrationId

`func (o *ApplicationSession) SetStoreIntegrationId(v string)`

SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.

### GetTotal

`func (o *ApplicationSession) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApplicationSession) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *ApplicationSession) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *ApplicationSession) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetTotalDiscounts

`func (o *ApplicationSession) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ApplicationSession) GetTotalDiscountsOk() (float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscounts

`func (o *ApplicationSession) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### SetTotalDiscounts

`func (o *ApplicationSession) SetTotalDiscounts(v float32)`

SetTotalDiscounts gets a reference to the given float32 and assigns it to the TotalDiscounts field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


