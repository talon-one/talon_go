# CustomerSessionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int64** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**CouponCodes** | Pointer to **[]string** | Any coupon codes entered.  **Important - for requests only**:  - If you [create a coupon budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a coupon code by the time you close it. - In requests where &#x60;dry&#x3D;false&#x60;, providing an empty array discards any previous coupons. To avoid this, omit the parameter entirely.  | [optional] 
**ReferralCode** | Pointer to **string** | Any referral code entered.  **Important - for requests only**:  - If you [create a referral budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a referral code by the time you close it. - In requests where &#x60;dry&#x3D;false&#x60;, providing an empty value discards the previous referral code. To avoid this, omit the parameter entirely.  | [optional] 
**LoyaltyCards** | Pointer to **[]string** | Identifier of a loyalty card. | [optional] 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. Either:    - &#x60;closed&#x60; → &#x60;cancelled&#x60; (**only** via [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2)) or    - &#x60;closed&#x60; → &#x60;partially_returned&#x60; (**only** via [Return cart items](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/returnCartItems))    - &#x60;closed&#x60; → &#x60;open&#x60; (**only** via [Reopen customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/reopenCustomerSession)) 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | [default to "open"]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | The items to add to this session. **Do not exceed 1000 items** and ensure the sum of all cart item&#39;s &#x60;quantity&#x60; **does not exceed 10.000** per request.  | 
**AdditionalCosts** | Pointer to [**map[string]AdditionalCost**](AdditionalCost.md) | Use this property to set a value for the additional costs of this session, such as a shipping cost.  They must be created in the Campaign Manager before you set them with this property. See [Managing additional costs](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs).  | [optional] 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers).  **Important**: Ensure the session contains an identifier by the time you close it if: - You [create a unique identifier budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign. - Your campaign has [coupons](https://docs.talon.one/docs/product/campaigns/coupons/coupon-page-overview). - We recommend passing an anonymized (hashed) version of the identifier value.  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Use this property to set a value for the attributes of your choice. Attributes represent any information to attach to your session, like the shipping city.  You can use [built-in attributes](https://docs.talon.one/docs/dev/concepts/attributes#built-in-attributes) or [custom ones](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes). Custom attributes must be created in the Campaign Manager before you set them with this property.  | 
**FirstSession** | Pointer to **bool** | Indicates whether this is the first session for the customer&#39;s profile. It&#39;s always &#x60;true&#x60; for anonymous sessions. | 
**Total** | Pointer to **float32** | The total value of cart items and additional costs in the session, before any discounts are applied. | 
**CartItemTotal** | Pointer to **float32** | The total value of cart items, before any discounts are applied. | 
**AdditionalCostTotal** | Pointer to **float32** | The total value of additional costs, before any discounts are applied. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received on this session. | 

## Methods

### NewCustomerSessionV2

`func NewCustomerSessionV2(id int64, created time.Time, integrationId string, applicationId int64, profileId string, state string, cartItems []CartItem, attributes map[string]interface{}, firstSession bool, total float32, cartItemTotal float32, additionalCostTotal float32, updated time.Time, ) *CustomerSessionV2`

NewCustomerSessionV2 instantiates a new CustomerSessionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSessionV2WithDefaults

`func NewCustomerSessionV2WithDefaults() *CustomerSessionV2`

NewCustomerSessionV2WithDefaults instantiates a new CustomerSessionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerSessionV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerSessionV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerSessionV2) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CustomerSessionV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerSessionV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerSessionV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetIntegrationId

`func (o *CustomerSessionV2) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerSessionV2) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerSessionV2) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetApplicationId

`func (o *CustomerSessionV2) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CustomerSessionV2) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CustomerSessionV2) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileId

`func (o *CustomerSessionV2) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CustomerSessionV2) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CustomerSessionV2) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetStoreIntegrationId

`func (o *CustomerSessionV2) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *CustomerSessionV2) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *CustomerSessionV2) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *CustomerSessionV2) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetEvaluableCampaignIds

`func (o *CustomerSessionV2) GetEvaluableCampaignIds() []int64`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *CustomerSessionV2) GetEvaluableCampaignIdsOk() (*[]int64, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluableCampaignIds

`func (o *CustomerSessionV2) SetEvaluableCampaignIds(v []int64)`

SetEvaluableCampaignIds sets EvaluableCampaignIds field to given value.

### HasEvaluableCampaignIds

`func (o *CustomerSessionV2) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### GetCouponCodes

`func (o *CustomerSessionV2) GetCouponCodes() []string`

GetCouponCodes returns the CouponCodes field if non-nil, zero value otherwise.

### GetCouponCodesOk

`func (o *CustomerSessionV2) GetCouponCodesOk() (*[]string, bool)`

GetCouponCodesOk returns a tuple with the CouponCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodes

`func (o *CustomerSessionV2) SetCouponCodes(v []string)`

SetCouponCodes sets CouponCodes field to given value.

### HasCouponCodes

`func (o *CustomerSessionV2) HasCouponCodes() bool`

HasCouponCodes returns a boolean if a field has been set.

### GetReferralCode

`func (o *CustomerSessionV2) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *CustomerSessionV2) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *CustomerSessionV2) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *CustomerSessionV2) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### GetLoyaltyCards

`func (o *CustomerSessionV2) GetLoyaltyCards() []string`

GetLoyaltyCards returns the LoyaltyCards field if non-nil, zero value otherwise.

### GetLoyaltyCardsOk

`func (o *CustomerSessionV2) GetLoyaltyCardsOk() (*[]string, bool)`

GetLoyaltyCardsOk returns a tuple with the LoyaltyCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCards

`func (o *CustomerSessionV2) SetLoyaltyCards(v []string)`

SetLoyaltyCards sets LoyaltyCards field to given value.

### HasLoyaltyCards

`func (o *CustomerSessionV2) HasLoyaltyCards() bool`

HasLoyaltyCards returns a boolean if a field has been set.

### GetState

`func (o *CustomerSessionV2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerSessionV2) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerSessionV2) SetState(v string)`

SetState sets State field to given value.


### GetCartItems

`func (o *CustomerSessionV2) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CustomerSessionV2) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CustomerSessionV2) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.


### GetAdditionalCosts

`func (o *CustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *CustomerSessionV2) GetAdditionalCostsOk() (*map[string]AdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCosts

`func (o *CustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost)`

SetAdditionalCosts sets AdditionalCosts field to given value.

### HasAdditionalCosts

`func (o *CustomerSessionV2) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### GetIdentifiers

`func (o *CustomerSessionV2) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *CustomerSessionV2) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *CustomerSessionV2) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *CustomerSessionV2) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetAttributes

`func (o *CustomerSessionV2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerSessionV2) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerSessionV2) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetFirstSession

`func (o *CustomerSessionV2) GetFirstSession() bool`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *CustomerSessionV2) GetFirstSessionOk() (*bool, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSession

`func (o *CustomerSessionV2) SetFirstSession(v bool)`

SetFirstSession sets FirstSession field to given value.


### GetTotal

`func (o *CustomerSessionV2) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomerSessionV2) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CustomerSessionV2) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetCartItemTotal

`func (o *CustomerSessionV2) GetCartItemTotal() float32`

GetCartItemTotal returns the CartItemTotal field if non-nil, zero value otherwise.

### GetCartItemTotalOk

`func (o *CustomerSessionV2) GetCartItemTotalOk() (*float32, bool)`

GetCartItemTotalOk returns a tuple with the CartItemTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemTotal

`func (o *CustomerSessionV2) SetCartItemTotal(v float32)`

SetCartItemTotal sets CartItemTotal field to given value.


### GetAdditionalCostTotal

`func (o *CustomerSessionV2) GetAdditionalCostTotal() float32`

GetAdditionalCostTotal returns the AdditionalCostTotal field if non-nil, zero value otherwise.

### GetAdditionalCostTotalOk

`func (o *CustomerSessionV2) GetAdditionalCostTotalOk() (*float32, bool)`

GetAdditionalCostTotalOk returns a tuple with the AdditionalCostTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCostTotal

`func (o *CustomerSessionV2) SetAdditionalCostTotal(v float32)`

SetAdditionalCostTotal sets AdditionalCostTotal field to given value.


### GetUpdated

`func (o *CustomerSessionV2) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomerSessionV2) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CustomerSessionV2) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


