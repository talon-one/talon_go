# CustomerSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**Coupon** | Pointer to **string** | Any coupon code entered. | 
**Referral** | Pointer to **string** | Any referral code entered. | 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60; or &#x60;partially_returned&#x60; 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | [default to "open"]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | Serialized JSON representation. | 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers).  | [optional] 
**Total** | Pointer to **float32** | The total sum of the cart in one session. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | 
**FirstSession** | Pointer to **bool** | Indicates whether this is the first session for the customer&#39;s profile. Will always be true for anonymous sessions. | 
**Discounts** | Pointer to **map[string]float32** | A map of labelled discount values, values will be in the same currency as the application associated with the session. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received on this session. | 

## Methods

### NewCustomerSession

`func NewCustomerSession(integrationId string, created time.Time, applicationId int64, profileId string, coupon string, referral string, state string, cartItems []CartItem, total float32, attributes map[string]interface{}, firstSession bool, discounts map[string]float32, updated time.Time, ) *CustomerSession`

NewCustomerSession instantiates a new CustomerSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSessionWithDefaults

`func NewCustomerSessionWithDefaults() *CustomerSession`

NewCustomerSessionWithDefaults instantiates a new CustomerSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *CustomerSession) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerSession) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerSession) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetCreated

`func (o *CustomerSession) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerSession) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerSession) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *CustomerSession) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CustomerSession) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CustomerSession) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *CustomerSession) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CustomerSession) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CustomerSession) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetCoupon

`func (o *CustomerSession) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *CustomerSession) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *CustomerSession) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.


### GetReferral

`func (o *CustomerSession) GetReferral() string`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *CustomerSession) GetReferralOk() (*string, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *CustomerSession) SetReferral(v string)`

SetReferral sets Referral field to given value.


### GetState

`func (o *CustomerSession) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerSession) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerSession) SetState(v string)`

SetState sets State field to given value.


### GetCartItems

`func (o *CustomerSession) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CustomerSession) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CustomerSession) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.


### GetIdentifiers

`func (o *CustomerSession) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *CustomerSession) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *CustomerSession) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *CustomerSession) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetTotal

`func (o *CustomerSession) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomerSession) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CustomerSession) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetAttributes

`func (o *CustomerSession) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSession) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerSession) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetFirstSession

`func (o *CustomerSession) GetFirstSession() bool`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *CustomerSession) GetFirstSessionOk() (*bool, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSession

`func (o *CustomerSession) SetFirstSession(v bool)`

SetFirstSession sets FirstSession field to given value.


### GetDiscounts

`func (o *CustomerSession) GetDiscounts() map[string]float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *CustomerSession) GetDiscountsOk() (*map[string]float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *CustomerSession) SetDiscounts(v map[string]float32)`

SetDiscounts sets Discounts field to given value.


### GetUpdated

`func (o *CustomerSession) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomerSession) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CustomerSession) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


