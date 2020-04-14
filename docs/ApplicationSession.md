# ApplicationSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**ProfileId** | Pointer to **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**IntegrationId** | Pointer to **string** | The ID used for this entity in the application system. | 
**Profileintegrationid** | Pointer to **string** | Integration ID of the customer for the session. | [optional] 
**Coupon** | Pointer to **string** | Any coupon code entered. | 
**Referral** | Pointer to **string** | Any referral code entered. | 
**State** | Pointer to **string** | Indicating if the customer session is in progress (\&quot;open\&quot;), \&quot;closed\&quot;, or \&quot;cancelled\&quot;. | 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | 
**Discounts** | Pointer to **map[string]float32** | A map of labelled discount values, in the same currency as the session. | 
**Total** | Pointer to **float32** | The total sum of the session before any discounts applied. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 

## Methods

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


