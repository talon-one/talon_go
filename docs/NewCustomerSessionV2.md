# NewCustomerSessionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int32** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**CouponCodes** | Pointer to **[]string** | Any coupon codes entered.  **Important - for requests only**:  - If you [create a coupon budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a coupon code by the time you close it. - In requests where &#x60;dry&#x3D;false&#x60;, providing an empty array discards any previous coupons. To avoid this, provide &#x60;\&quot;couponCodes\&quot;: null&#x60; or omit the parameter entirely.  | [optional] 
**ReferralCode** | Pointer to **string** | Any referral code entered.  **Important - for requests only**:  - If you [create a referral budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign, ensure the session contains a referral code by the time you close it. - In requests where &#x60;dry&#x3D;false&#x60;, providing an empty value discards the previous referral code. To avoid this, provide &#x60;\&quot;referralCode\&quot;: null&#x60; or omit the parameter entirely.  | [optional] 
**LoyaltyCards** | Pointer to **[]string** | Identifier of a loyalty card. | [optional] 
**State** | Pointer to **string** | Indicates the current state of the session. Sessions can be created as &#x60;open&#x60; or &#x60;closed&#x60;. The state transitions are:  1. &#x60;open&#x60; → &#x60;closed&#x60; 2. &#x60;open&#x60; → &#x60;cancelled&#x60; 3. Either:    - &#x60;closed&#x60; → &#x60;cancelled&#x60; (**only** via [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2)) or    - &#x60;closed&#x60; → &#x60;partially_returned&#x60; (**only** via [Return cart items](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/returnCartItems))    - &#x60;closed&#x60; → &#x60;open&#x60; (**only** via [Reopen customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/reopenCustomerSession)) 4. &#x60;partially_returned&#x60; → &#x60;cancelled&#x60;  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions).  | [optional] [default to STATE_OPEN]
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | The items to add to this session. **Do not exceed 1000 items** and ensure the sum of all cart item&#39;s &#x60;quantity&#x60; **does not exceed 10.000** per request.  | [optional] 
**AdditionalCosts** | Pointer to [**map[string]AdditionalCost**](AdditionalCost.md) | Use this property to set a value for the additional costs of this session, such as a shipping cost.  They must be created in the Campaign Manager before you set them with this property. See [Managing additional costs](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs).  | [optional] 
**Identifiers** | Pointer to **[]string** | Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers).  **Important**: Ensure the session contains an identifier by the time you close it if: - You [create a unique identifier budget](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets/#budget-types) for your campaign. - Your campaign has [coupons](https://docs.talon.one/docs/product/campaigns/coupons/coupon-page-overview).  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Use this property to set a value for the attributes of your choice. Attributes represent any information to attach to your session, like the shipping city.  You can use [built-in attributes](https://docs.talon.one/docs/dev/concepts/attributes#built-in-attributes) or [custom ones](https://docs.talon.one/docs/dev/concepts/attributes#custom-attributes). Custom attributes must be created in the Campaign Manager before you set them with this property.  | [optional] 

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

### GetStoreIntegrationId

`func (o *NewCustomerSessionV2) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *NewCustomerSessionV2) GetStoreIntegrationIdOk() (string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoreIntegrationId

`func (o *NewCustomerSessionV2) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### SetStoreIntegrationId

`func (o *NewCustomerSessionV2) SetStoreIntegrationId(v string)`

SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.

### GetEvaluableCampaignIds

`func (o *NewCustomerSessionV2) GetEvaluableCampaignIds() []int32`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *NewCustomerSessionV2) GetEvaluableCampaignIdsOk() ([]int32, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluableCampaignIds

`func (o *NewCustomerSessionV2) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### SetEvaluableCampaignIds

`func (o *NewCustomerSessionV2) SetEvaluableCampaignIds(v []int32)`

SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.

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

### GetLoyaltyCards

`func (o *NewCustomerSessionV2) GetLoyaltyCards() []string`

GetLoyaltyCards returns the LoyaltyCards field if non-nil, zero value otherwise.

### GetLoyaltyCardsOk

`func (o *NewCustomerSessionV2) GetLoyaltyCardsOk() ([]string, bool)`

GetLoyaltyCardsOk returns a tuple with the LoyaltyCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyCards

`func (o *NewCustomerSessionV2) HasLoyaltyCards() bool`

HasLoyaltyCards returns a boolean if a field has been set.

### SetLoyaltyCards

`func (o *NewCustomerSessionV2) SetLoyaltyCards(v []string)`

SetLoyaltyCards gets a reference to the given []string and assigns it to the LoyaltyCards field.

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


