# ApplicationSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**ProfileId** | Pointer to **int64** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 
**Profileintegrationid** | Pointer to **string** | Integration ID of the customer for the session. | [optional] 
**Coupon** | Pointer to **string** | Any coupon code entered. | 
**Referral** | Pointer to **string** | Any referral code entered. | 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60; or &#x60;partially_returned&#x60; 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | 
**Discounts** | Pointer to **map[string]float32** | **API V1 only.** A map of labeled discount values, in the same currency as the session.  If you are using the V2 endpoints, refer to the &#x60;totalDiscounts&#x60; property instead.  | 
**TotalDiscounts** | Pointer to **float32** | The total sum of the discounts applied to this session.  **Note:** If more than one session is returned, this value is displayed as &#x60;0&#x60;.  | 
**Total** | Pointer to **float32** | The total sum of the session before any discounts applied. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 

## Methods

### NewApplicationSession

`func NewApplicationSession(id int64, created time.Time, integrationId string, applicationId int64, coupon string, referral string, state string, cartItems []CartItem, discounts map[string]float32, totalDiscounts float32, total float32, ) *ApplicationSession`

NewApplicationSession instantiates a new ApplicationSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSessionWithDefaults

`func NewApplicationSessionWithDefaults() *ApplicationSession`

NewApplicationSessionWithDefaults instantiates a new ApplicationSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationSession) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ApplicationSession) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationSession) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationSession) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetIntegrationId

`func (o *ApplicationSession) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ApplicationSession) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ApplicationSession) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetStoreIntegrationId

`func (o *ApplicationSession) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *ApplicationSession) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *ApplicationSession) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *ApplicationSession) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationSession) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationSession) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationSession) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *ApplicationSession) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationSession) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ApplicationSession) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ApplicationSession) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileintegrationid

`func (o *ApplicationSession) GetProfileintegrationid() string`

GetProfileintegrationid returns the Profileintegrationid field if non-nil, zero value otherwise.

### GetProfileintegrationidOk

`func (o *ApplicationSession) GetProfileintegrationidOk() (*string, bool)`

GetProfileintegrationidOk returns a tuple with the Profileintegrationid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileintegrationid

`func (o *ApplicationSession) SetProfileintegrationid(v string)`

SetProfileintegrationid sets Profileintegrationid field to given value.

### HasProfileintegrationid

`func (o *ApplicationSession) HasProfileintegrationid() bool`

HasProfileintegrationid returns a boolean if a field has been set.

### GetCoupon

`func (o *ApplicationSession) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *ApplicationSession) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *ApplicationSession) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.


### GetReferral

`func (o *ApplicationSession) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *ApplicationSession) GetReferralOk() (*string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *ApplicationSession) SetReferral(v string)`

SetReferral sets Referral field to given value.


### GetState

`func (o *ApplicationSession) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationSession) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationSession) SetState(v string)`

SetState sets State field to given value.


### GetCartItems

`func (o *ApplicationSession) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *ApplicationSession) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *ApplicationSession) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.


### GetDiscounts

`func (o *ApplicationSession) GetDiscounts() map[string]float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *ApplicationSession) GetDiscountsOk() (*map[string]float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *ApplicationSession) SetDiscounts(v map[string]float32)`

SetDiscounts sets Discounts field to given value.


### GetTotalDiscounts

`func (o *ApplicationSession) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *ApplicationSession) GetTotalDiscountsOk() (*float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *ApplicationSession) SetTotalDiscounts(v float32)`

SetTotalDiscounts sets TotalDiscounts field to given value.


### GetTotal

`func (o *ApplicationSession) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApplicationSession) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApplicationSession) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetAttributes

`func (o *ApplicationSession) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationSession) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationSession) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationSession) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


