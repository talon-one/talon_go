# CustomerSessionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID for this entity sent to and used in the Talon.One system. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**ProfileId** | Pointer to **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | 
**CouponCodes** | Pointer to **[]string** | Any coupon codes entered. | [optional] 
**ReferralCode** | Pointer to **string** | Any referral code entered. | [optional] 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;, after which valid transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. &#x60;closed&#x60; → &#x60;cancelled&#x60;  For more information, see [Entites](/docs/dev/concepts/entities#customer-session).  | [default to STATE_OPEN]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | All items the customer will be purchasing in this session | 
**AdditionalCosts** | Pointer to [**map[string]AdditionalCost**](AdditionalCost.md) | Any costs associated with the session that can not be explicitly attributed to cart items. Examples include shipping costs and service fees. | [optional] 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts.  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | 
**FirstSession** | Pointer to **bool** | Indicates whether this is the first session for the customer&#39;s profile. Will always be true for anonymous sessions. | 
**Total** | Pointer to **float32** | The total sum of cart-items, as well as additional costs, before any discounts applied | 
**CartItemTotal** | Pointer to **float32** | The total sum of cart-items before any discounts applied | 
**AdditionalCostTotal** | Pointer to **float32** | The total sum of additional costs before any discounts applied | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received on this session | 

## Methods

### GetIntegrationId

`func (o *CustomerSessionV2) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerSessionV2) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *CustomerSessionV2) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *CustomerSessionV2) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetCreated

`func (o *CustomerSessionV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerSessionV2) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CustomerSessionV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CustomerSessionV2) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *CustomerSessionV2) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CustomerSessionV2) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CustomerSessionV2) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CustomerSessionV2) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetProfileId

`func (o *CustomerSessionV2) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CustomerSessionV2) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *CustomerSessionV2) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *CustomerSessionV2) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetCouponCodes

`func (o *CustomerSessionV2) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *CustomerSessionV2) GetCouponCodesOk() ([]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponCodes

`func (o *CustomerSessionV2) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### SetCouponCodes

`func (o *CustomerSessionV2) SetCouponCodes(v []string)`

SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.

### GetReferralCode

`func (o *CustomerSessionV2) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *CustomerSessionV2) GetReferralCodeOk() (string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCode

`func (o *CustomerSessionV2) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### SetReferralCode

`func (o *CustomerSessionV2) SetReferralCode(v string)`

SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.

### GetState

`func (o *CustomerSessionV2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerSessionV2) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *CustomerSessionV2) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *CustomerSessionV2) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetCartItems

`func (o *CustomerSessionV2) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CustomerSessionV2) GetCartItemsOk() ([]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItems

`func (o *CustomerSessionV2) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItems

`func (o *CustomerSessionV2) SetCartItems(v []CartItem)`

SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.

### GetAdditionalCosts

`func (o *CustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *CustomerSessionV2) GetAdditionalCostsOk() (map[string]AdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCosts

`func (o *CustomerSessionV2) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### SetAdditionalCosts

`func (o *CustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost)`

SetAdditionalCosts gets a reference to the given map[string]AdditionalCost and assigns it to the AdditionalCosts field.

### GetIdentifiers

`func (o *CustomerSessionV2) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *CustomerSessionV2) GetIdentifiersOk() ([]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifiers

`func (o *CustomerSessionV2) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiers

`func (o *CustomerSessionV2) SetIdentifiers(v []string)`

SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.

### GetAttributes

`func (o *CustomerSessionV2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSessionV2) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CustomerSessionV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CustomerSessionV2) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetFirstSession

`func (o *CustomerSessionV2) GetFirstSession() bool`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *CustomerSessionV2) GetFirstSessionOk() (bool, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstSession

`func (o *CustomerSessionV2) HasFirstSession() bool`

HasFirstSession returns a boolean if a field has been set.

### SetFirstSession

`func (o *CustomerSessionV2) SetFirstSession(v bool)`

SetFirstSession gets a reference to the given bool and assigns it to the FirstSession field.

### GetTotal

`func (o *CustomerSessionV2) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomerSessionV2) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *CustomerSessionV2) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *CustomerSessionV2) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetCartItemTotal

`func (o *CustomerSessionV2) GetCartItemTotal() float32`

GetCartItemTotal returns the CartItemTotal field if non-nil, zero value otherwise.

### GetCartItemTotalOk

`func (o *CustomerSessionV2) GetCartItemTotalOk() (float32, bool)`

GetCartItemTotalOk returns a tuple with the CartItemTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemTotal

`func (o *CustomerSessionV2) HasCartItemTotal() bool`

HasCartItemTotal returns a boolean if a field has been set.

### SetCartItemTotal

`func (o *CustomerSessionV2) SetCartItemTotal(v float32)`

SetCartItemTotal gets a reference to the given float32 and assigns it to the CartItemTotal field.

### GetAdditionalCostTotal

`func (o *CustomerSessionV2) GetAdditionalCostTotal() float32`

GetAdditionalCostTotal returns the AdditionalCostTotal field if non-nil, zero value otherwise.

### GetAdditionalCostTotalOk

`func (o *CustomerSessionV2) GetAdditionalCostTotalOk() (float32, bool)`

GetAdditionalCostTotalOk returns a tuple with the AdditionalCostTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCostTotal

`func (o *CustomerSessionV2) HasAdditionalCostTotal() bool`

HasAdditionalCostTotal returns a boolean if a field has been set.

### SetAdditionalCostTotal

`func (o *CustomerSessionV2) SetAdditionalCostTotal(v float32)`

SetAdditionalCostTotal gets a reference to the given float32 and assigns it to the AdditionalCostTotal field.

### GetUpdated

`func (o *CustomerSessionV2) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomerSessionV2) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *CustomerSessionV2) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *CustomerSessionV2) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


