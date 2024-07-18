# EffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**Id** | Pointer to **int32** | The id of the referral code that was redeemed. | 
**RejectionReason** | Pointer to **string** | The reason why this referral code was rejected. | 
**ConditionIndex** | Pointer to **int32** | The index of the condition that caused the rejection of the referral. | [optional] 
**EffectIndex** | Pointer to **int32** | The index of the effect that caused the rejection of the referral. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**CampaignExclusionReason** | Pointer to **string** | The reason why the campaign was not applied. | [optional] 
**ProfileId** | Pointer to **int32** | The internal ID of the customer profile. | 
**Name** | Pointer to **string** | The name / description of this discount | 
**Scope** | Pointer to **string** | The scope of the rolled back discount - For a discount per session, it can be one of &#x60;cartItems&#x60;, &#x60;additionalCosts&#x60; or &#x60;sessionTotal&#x60; - For a discount per item, it can be one of &#x60;price&#x60;, &#x60;additionalCosts&#x60; or &#x60;itemTotal&#x60;  | [optional] 
**DesiredValue** | Pointer to **float32** | Only with [partial discounts enabled](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#partial-discounts). Represents the monetary value of the discount to be applied to additional discount without considering budget limitations.  | [optional] 
**Position** | Pointer to **float32** | The index of the item in the cart item list containing the additional cost to be discounted. | 
**SubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the sub position indicates which item the discount applies to.  | [optional] 
**TotalDiscount** | Pointer to **float32** | The total discount given if this effect is a result of a prorated discount. | [optional] 
**DesiredTotalDiscount** | Pointer to **float32** | The original total discount to give if this effect is a result of a prorated discount. | [optional] 
**BundleIndex** | Pointer to **int32** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 
**TargetedItemPosition** | Pointer to **float32** | The index of the targeted bundle item on which the applied discount is based. | [optional] 
**TargetedItemSubPosition** | Pointer to **float32** | The sub-position of the targeted bundle item on which the applied discount is based.  | [optional] 
**AdditionalCostId** | Pointer to **int32** | The ID of the additional cost. | 
**AdditionalCost** | Pointer to **string** | The name of the additional cost. | 
**WebhookId** | Pointer to **float32** | The ID of the webhook that was triggered. | 
**WebhookName** | Pointer to **string** | The name of the webhook that was triggered. | 
**ProgramId** | Pointer to **int32** | The ID of the loyalty program where these points were reimbursed. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program where these points were reimbursed. | 
**RecipientIntegrationId** | Pointer to **string** | The integration ID of the profile that will be awarded the giveaway. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will be valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Date after which the reimbursed points will expire. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of &#39;addition&#39; entries added to the ledger as the &#x60;deductLoyaltyPoints&#x60; effect is rolled back. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart item list to which the custom effect is applied. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with quantity &gt; 1, the sub position indicates to which item unit the custom effect is applied.  | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 
**RuleTitle** | Pointer to **string** | The title of the rule that triggered the tier upgrade. | 
**PreviousTierName** | Pointer to **string** | The name of the tier from which the user was upgraded. | [optional] 
**NewTierName** | Pointer to **string** | The name of the tier to which the user has been upgraded. | 
**Sku** | Pointer to **string** | SKU of the item that needs to be added. | 
**NotificationType** | Pointer to **string** | The type of notification that should be shown (e.g. error/warning/info). | 
**Title** | Pointer to **string** | Title of the notification. | 
**Body** | Pointer to **string** | Body of the notification. | 
**Path** | Pointer to **string** | The exact path of the attribute that was updated. | 
**Description** | Pointer to **string** | Description of the product bundle. | 
**BundleAttributes** | Pointer to **[]string** | The cart item attributes that determined which items are being bundled together. | 
**ItemsIndices** | Pointer to **[]float32** | The indices in the cart items array of the bundled items. | 
**PoolId** | Pointer to **int32** | The ID of the giveaways pool the code will be taken from. | 
**PoolName** | Pointer to **string** | The name of the giveaways pool the code will be taken from. | 
**GiveawayId** | Pointer to **int32** | The internal ID for the giveaway that was awarded. | 
**Code** | Pointer to **string** | The giveaway code that was awarded. | 
**Message** | Pointer to **string** | The error message. | 
**EffectId** | Pointer to **int32** | The ID of the custom effect that was triggered. | 
**Payload** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | The JSON payload of the custom effect. | 
**CouponValue** | Pointer to **string** | The value of the coupon currently on scope. | 
**ProfileIntegrationId** | Pointer to **string** | The ID of the customer profile in the third-party integration platform. | 
**IsNewReservation** | Pointer to **bool** | Indicates whether this is a new coupon reservation or not. | 
**AudienceId** | Pointer to **int32** | The internal ID of the audience. | [optional] 
**AudienceName** | Pointer to **string** | The name of the audience. | [optional] 
**AchievementId** | Pointer to **int32** | The internal ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int32** | The internal ID of the achievement progress tracker. | 
**Delta** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement is increased. | 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 
**IsJustCompleted** | Pointer to **bool** | Indicates if the customer has completed the achievement in the current session. | 
**DecreaseProgressBy** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement is decreased. | 
**CurrentProgress** | Pointer to **float32** | The current progress of the customer in the achievement. | 

## Methods

### GetValue

`func (o *EffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *EffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *EffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetId

`func (o *EffectProps) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EffectProps) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *EffectProps) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *EffectProps) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetRejectionReason

`func (o *EffectProps) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *EffectProps) GetRejectionReasonOk() (string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRejectionReason

`func (o *EffectProps) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReason

`func (o *EffectProps) SetRejectionReason(v string)`

SetRejectionReason gets a reference to the given string and assigns it to the RejectionReason field.

### GetConditionIndex

`func (o *EffectProps) GetConditionIndex() int32`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *EffectProps) GetConditionIndexOk() (int32, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionIndex

`func (o *EffectProps) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### SetConditionIndex

`func (o *EffectProps) SetConditionIndex(v int32)`

SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.

### GetEffectIndex

`func (o *EffectProps) GetEffectIndex() int32`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *EffectProps) GetEffectIndexOk() (int32, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectIndex

`func (o *EffectProps) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### SetEffectIndex

`func (o *EffectProps) SetEffectIndex(v int32)`

SetEffectIndex gets a reference to the given int32 and assigns it to the EffectIndex field.

### GetDetails

`func (o *EffectProps) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EffectProps) GetDetailsOk() (string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *EffectProps) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *EffectProps) SetDetails(v string)`

SetDetails gets a reference to the given string and assigns it to the Details field.

### GetCampaignExclusionReason

`func (o *EffectProps) GetCampaignExclusionReason() string`

GetCampaignExclusionReason returns the CampaignExclusionReason field if non-nil, zero value otherwise.

### GetCampaignExclusionReasonOk

`func (o *EffectProps) GetCampaignExclusionReasonOk() (string, bool)`

GetCampaignExclusionReasonOk returns a tuple with the CampaignExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignExclusionReason

`func (o *EffectProps) HasCampaignExclusionReason() bool`

HasCampaignExclusionReason returns a boolean if a field has been set.

### SetCampaignExclusionReason

`func (o *EffectProps) SetCampaignExclusionReason(v string)`

SetCampaignExclusionReason gets a reference to the given string and assigns it to the CampaignExclusionReason field.

### GetProfileId

`func (o *EffectProps) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EffectProps) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *EffectProps) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *EffectProps) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetName

`func (o *EffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *EffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *EffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetScope

`func (o *EffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EffectProps) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *EffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *EffectProps) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### GetDesiredValue

`func (o *EffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *EffectProps) GetDesiredValueOk() (float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredValue

`func (o *EffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### SetDesiredValue

`func (o *EffectProps) SetDesiredValue(v float32)`

SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.

### GetPosition

`func (o *EffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EffectProps) GetPositionOk() (float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *EffectProps) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *EffectProps) SetPosition(v float32)`

SetPosition gets a reference to the given float32 and assigns it to the Position field.

### GetSubPosition

`func (o *EffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *EffectProps) GetSubPositionOk() (float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubPosition

`func (o *EffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### SetSubPosition

`func (o *EffectProps) SetSubPosition(v float32)`

SetSubPosition gets a reference to the given float32 and assigns it to the SubPosition field.

### GetTotalDiscount

`func (o *EffectProps) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *EffectProps) GetTotalDiscountOk() (float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscount

`func (o *EffectProps) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscount

`func (o *EffectProps) SetTotalDiscount(v float32)`

SetTotalDiscount gets a reference to the given float32 and assigns it to the TotalDiscount field.

### GetDesiredTotalDiscount

`func (o *EffectProps) GetDesiredTotalDiscount() float32`

GetDesiredTotalDiscount returns the DesiredTotalDiscount field if non-nil, zero value otherwise.

### GetDesiredTotalDiscountOk

`func (o *EffectProps) GetDesiredTotalDiscountOk() (float32, bool)`

GetDesiredTotalDiscountOk returns a tuple with the DesiredTotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredTotalDiscount

`func (o *EffectProps) HasDesiredTotalDiscount() bool`

HasDesiredTotalDiscount returns a boolean if a field has been set.

### SetDesiredTotalDiscount

`func (o *EffectProps) SetDesiredTotalDiscount(v float32)`

SetDesiredTotalDiscount gets a reference to the given float32 and assigns it to the DesiredTotalDiscount field.

### GetBundleIndex

`func (o *EffectProps) GetBundleIndex() int32`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *EffectProps) GetBundleIndexOk() (int32, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleIndex

`func (o *EffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### SetBundleIndex

`func (o *EffectProps) SetBundleIndex(v int32)`

SetBundleIndex gets a reference to the given int32 and assigns it to the BundleIndex field.

### GetBundleName

`func (o *EffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *EffectProps) GetBundleNameOk() (string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleName

`func (o *EffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### SetBundleName

`func (o *EffectProps) SetBundleName(v string)`

SetBundleName gets a reference to the given string and assigns it to the BundleName field.

### GetTargetedItemPosition

`func (o *EffectProps) GetTargetedItemPosition() float32`

GetTargetedItemPosition returns the TargetedItemPosition field if non-nil, zero value otherwise.

### GetTargetedItemPositionOk

`func (o *EffectProps) GetTargetedItemPositionOk() (float32, bool)`

GetTargetedItemPositionOk returns a tuple with the TargetedItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedItemPosition

`func (o *EffectProps) HasTargetedItemPosition() bool`

HasTargetedItemPosition returns a boolean if a field has been set.

### SetTargetedItemPosition

`func (o *EffectProps) SetTargetedItemPosition(v float32)`

SetTargetedItemPosition gets a reference to the given float32 and assigns it to the TargetedItemPosition field.

### GetTargetedItemSubPosition

`func (o *EffectProps) GetTargetedItemSubPosition() float32`

GetTargetedItemSubPosition returns the TargetedItemSubPosition field if non-nil, zero value otherwise.

### GetTargetedItemSubPositionOk

`func (o *EffectProps) GetTargetedItemSubPositionOk() (float32, bool)`

GetTargetedItemSubPositionOk returns a tuple with the TargetedItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedItemSubPosition

`func (o *EffectProps) HasTargetedItemSubPosition() bool`

HasTargetedItemSubPosition returns a boolean if a field has been set.

### SetTargetedItemSubPosition

`func (o *EffectProps) SetTargetedItemSubPosition(v float32)`

SetTargetedItemSubPosition gets a reference to the given float32 and assigns it to the TargetedItemSubPosition field.

### GetAdditionalCostId

`func (o *EffectProps) GetAdditionalCostId() int32`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *EffectProps) GetAdditionalCostIdOk() (int32, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCostId

`func (o *EffectProps) HasAdditionalCostId() bool`

HasAdditionalCostId returns a boolean if a field has been set.

### SetAdditionalCostId

`func (o *EffectProps) SetAdditionalCostId(v int32)`

SetAdditionalCostId gets a reference to the given int32 and assigns it to the AdditionalCostId field.

### GetAdditionalCost

`func (o *EffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *EffectProps) GetAdditionalCostOk() (string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCost

`func (o *EffectProps) HasAdditionalCost() bool`

HasAdditionalCost returns a boolean if a field has been set.

### SetAdditionalCost

`func (o *EffectProps) SetAdditionalCost(v string)`

SetAdditionalCost gets a reference to the given string and assigns it to the AdditionalCost field.

### GetWebhookId

`func (o *EffectProps) GetWebhookId() float32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *EffectProps) GetWebhookIdOk() (float32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhookId

`func (o *EffectProps) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookId

`func (o *EffectProps) SetWebhookId(v float32)`

SetWebhookId gets a reference to the given float32 and assigns it to the WebhookId field.

### GetWebhookName

`func (o *EffectProps) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *EffectProps) GetWebhookNameOk() (string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhookName

`func (o *EffectProps) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### SetWebhookName

`func (o *EffectProps) SetWebhookName(v string)`

SetWebhookName gets a reference to the given string and assigns it to the WebhookName field.

### GetProgramId

`func (o *EffectProps) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *EffectProps) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *EffectProps) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *EffectProps) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetSubLedgerId

`func (o *EffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *EffectProps) GetSubLedgerIdOk() (string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerId

`func (o *EffectProps) HasSubLedgerId() bool`

HasSubLedgerId returns a boolean if a field has been set.

### SetSubLedgerId

`func (o *EffectProps) SetSubLedgerId(v string)`

SetSubLedgerId gets a reference to the given string and assigns it to the SubLedgerId field.

### GetRecipientIntegrationId

`func (o *EffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *EffectProps) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *EffectProps) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *EffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetStartDate

`func (o *EffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EffectProps) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *EffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *EffectProps) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *EffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *EffectProps) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *EffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *EffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetTransactionUUID

`func (o *EffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *EffectProps) GetTransactionUUIDOk() (string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactionUUID

`func (o *EffectProps) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.

### SetTransactionUUID

`func (o *EffectProps) SetTransactionUUID(v string)`

SetTransactionUUID gets a reference to the given string and assigns it to the TransactionUUID field.

### GetCartItemPosition

`func (o *EffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *EffectProps) GetCartItemPositionOk() (float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemPosition

`func (o *EffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### SetCartItemPosition

`func (o *EffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.

### GetCartItemSubPosition

`func (o *EffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *EffectProps) GetCartItemSubPositionOk() (float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemSubPosition

`func (o *EffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### SetCartItemSubPosition

`func (o *EffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.

### GetCardIdentifier

`func (o *EffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *EffectProps) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *EffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *EffectProps) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.

### GetRuleTitle

`func (o *EffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *EffectProps) GetRuleTitleOk() (string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleTitle

`func (o *EffectProps) HasRuleTitle() bool`

HasRuleTitle returns a boolean if a field has been set.

### SetRuleTitle

`func (o *EffectProps) SetRuleTitle(v string)`

SetRuleTitle gets a reference to the given string and assigns it to the RuleTitle field.

### GetPreviousTierName

`func (o *EffectProps) GetPreviousTierName() string`

GetPreviousTierName returns the PreviousTierName field if non-nil, zero value otherwise.

### GetPreviousTierNameOk

`func (o *EffectProps) GetPreviousTierNameOk() (string, bool)`

GetPreviousTierNameOk returns a tuple with the PreviousTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviousTierName

`func (o *EffectProps) HasPreviousTierName() bool`

HasPreviousTierName returns a boolean if a field has been set.

### SetPreviousTierName

`func (o *EffectProps) SetPreviousTierName(v string)`

SetPreviousTierName gets a reference to the given string and assigns it to the PreviousTierName field.

### GetNewTierName

`func (o *EffectProps) GetNewTierName() string`

GetNewTierName returns the NewTierName field if non-nil, zero value otherwise.

### GetNewTierNameOk

`func (o *EffectProps) GetNewTierNameOk() (string, bool)`

GetNewTierNameOk returns a tuple with the NewTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewTierName

`func (o *EffectProps) HasNewTierName() bool`

HasNewTierName returns a boolean if a field has been set.

### SetNewTierName

`func (o *EffectProps) SetNewTierName(v string)`

SetNewTierName gets a reference to the given string and assigns it to the NewTierName field.

### GetSku

`func (o *EffectProps) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EffectProps) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *EffectProps) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *EffectProps) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetNotificationType

`func (o *EffectProps) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *EffectProps) GetNotificationTypeOk() (string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationType

`func (o *EffectProps) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### SetNotificationType

`func (o *EffectProps) SetNotificationType(v string)`

SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.

### GetTitle

`func (o *EffectProps) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EffectProps) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *EffectProps) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *EffectProps) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetBody

`func (o *EffectProps) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EffectProps) GetBodyOk() (string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *EffectProps) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *EffectProps) SetBody(v string)`

SetBody gets a reference to the given string and assigns it to the Body field.

### GetPath

`func (o *EffectProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EffectProps) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *EffectProps) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *EffectProps) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### GetDescription

`func (o *EffectProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EffectProps) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *EffectProps) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *EffectProps) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetBundleAttributes

`func (o *EffectProps) GetBundleAttributes() []string`

GetBundleAttributes returns the BundleAttributes field if non-nil, zero value otherwise.

### GetBundleAttributesOk

`func (o *EffectProps) GetBundleAttributesOk() ([]string, bool)`

GetBundleAttributesOk returns a tuple with the BundleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleAttributes

`func (o *EffectProps) HasBundleAttributes() bool`

HasBundleAttributes returns a boolean if a field has been set.

### SetBundleAttributes

`func (o *EffectProps) SetBundleAttributes(v []string)`

SetBundleAttributes gets a reference to the given []string and assigns it to the BundleAttributes field.

### GetItemsIndices

`func (o *EffectProps) GetItemsIndices() []float32`

GetItemsIndices returns the ItemsIndices field if non-nil, zero value otherwise.

### GetItemsIndicesOk

`func (o *EffectProps) GetItemsIndicesOk() ([]float32, bool)`

GetItemsIndicesOk returns a tuple with the ItemsIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItemsIndices

`func (o *EffectProps) HasItemsIndices() bool`

HasItemsIndices returns a boolean if a field has been set.

### SetItemsIndices

`func (o *EffectProps) SetItemsIndices(v []float32)`

SetItemsIndices gets a reference to the given []float32 and assigns it to the ItemsIndices field.

### GetPoolId

`func (o *EffectProps) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *EffectProps) GetPoolIdOk() (int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoolId

`func (o *EffectProps) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### SetPoolId

`func (o *EffectProps) SetPoolId(v int32)`

SetPoolId gets a reference to the given int32 and assigns it to the PoolId field.

### GetPoolName

`func (o *EffectProps) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *EffectProps) GetPoolNameOk() (string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoolName

`func (o *EffectProps) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolName

`func (o *EffectProps) SetPoolName(v string)`

SetPoolName gets a reference to the given string and assigns it to the PoolName field.

### GetGiveawayId

`func (o *EffectProps) GetGiveawayId() int32`

GetGiveawayId returns the GiveawayId field if non-nil, zero value otherwise.

### GetGiveawayIdOk

`func (o *EffectProps) GetGiveawayIdOk() (int32, bool)`

GetGiveawayIdOk returns a tuple with the GiveawayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGiveawayId

`func (o *EffectProps) HasGiveawayId() bool`

HasGiveawayId returns a boolean if a field has been set.

### SetGiveawayId

`func (o *EffectProps) SetGiveawayId(v int32)`

SetGiveawayId gets a reference to the given int32 and assigns it to the GiveawayId field.

### GetCode

`func (o *EffectProps) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EffectProps) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *EffectProps) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *EffectProps) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetMessage

`func (o *EffectProps) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EffectProps) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *EffectProps) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *EffectProps) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetEffectId

`func (o *EffectProps) GetEffectId() int32`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *EffectProps) GetEffectIdOk() (int32, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectId

`func (o *EffectProps) HasEffectId() bool`

HasEffectId returns a boolean if a field has been set.

### SetEffectId

`func (o *EffectProps) SetEffectId(v int32)`

SetEffectId gets a reference to the given int32 and assigns it to the EffectId field.

### GetPayload

`func (o *EffectProps) GetPayload() map[string]map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EffectProps) GetPayloadOk() (map[string]map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *EffectProps) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *EffectProps) SetPayload(v map[string]map[string]interface{})`

SetPayload gets a reference to the given map[string]map[string]interface{} and assigns it to the Payload field.

### GetCouponValue

`func (o *EffectProps) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *EffectProps) GetCouponValueOk() (string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponValue

`func (o *EffectProps) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### SetCouponValue

`func (o *EffectProps) SetCouponValue(v string)`

SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.

### GetProfileIntegrationId

`func (o *EffectProps) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *EffectProps) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *EffectProps) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *EffectProps) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetIsNewReservation

`func (o *EffectProps) GetIsNewReservation() bool`

GetIsNewReservation returns the IsNewReservation field if non-nil, zero value otherwise.

### GetIsNewReservationOk

`func (o *EffectProps) GetIsNewReservationOk() (bool, bool)`

GetIsNewReservationOk returns a tuple with the IsNewReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsNewReservation

`func (o *EffectProps) HasIsNewReservation() bool`

HasIsNewReservation returns a boolean if a field has been set.

### SetIsNewReservation

`func (o *EffectProps) SetIsNewReservation(v bool)`

SetIsNewReservation gets a reference to the given bool and assigns it to the IsNewReservation field.

### GetAudienceId

`func (o *EffectProps) GetAudienceId() int32`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *EffectProps) GetAudienceIdOk() (int32, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceId

`func (o *EffectProps) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### SetAudienceId

`func (o *EffectProps) SetAudienceId(v int32)`

SetAudienceId gets a reference to the given int32 and assigns it to the AudienceId field.

### GetAudienceName

`func (o *EffectProps) GetAudienceName() string`

GetAudienceName returns the AudienceName field if non-nil, zero value otherwise.

### GetAudienceNameOk

`func (o *EffectProps) GetAudienceNameOk() (string, bool)`

GetAudienceNameOk returns a tuple with the AudienceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceName

`func (o *EffectProps) HasAudienceName() bool`

HasAudienceName returns a boolean if a field has been set.

### SetAudienceName

`func (o *EffectProps) SetAudienceName(v string)`

SetAudienceName gets a reference to the given string and assigns it to the AudienceName field.

### GetAchievementId

`func (o *EffectProps) GetAchievementId() int32`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *EffectProps) GetAchievementIdOk() (int32, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementId

`func (o *EffectProps) HasAchievementId() bool`

HasAchievementId returns a boolean if a field has been set.

### SetAchievementId

`func (o *EffectProps) SetAchievementId(v int32)`

SetAchievementId gets a reference to the given int32 and assigns it to the AchievementId field.

### GetAchievementName

`func (o *EffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *EffectProps) GetAchievementNameOk() (string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementName

`func (o *EffectProps) HasAchievementName() bool`

HasAchievementName returns a boolean if a field has been set.

### SetAchievementName

`func (o *EffectProps) SetAchievementName(v string)`

SetAchievementName gets a reference to the given string and assigns it to the AchievementName field.

### GetProgressTrackerId

`func (o *EffectProps) GetProgressTrackerId() int32`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *EffectProps) GetProgressTrackerIdOk() (int32, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgressTrackerId

`func (o *EffectProps) HasProgressTrackerId() bool`

HasProgressTrackerId returns a boolean if a field has been set.

### SetProgressTrackerId

`func (o *EffectProps) SetProgressTrackerId(v int32)`

SetProgressTrackerId gets a reference to the given int32 and assigns it to the ProgressTrackerId field.

### GetDelta

`func (o *EffectProps) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *EffectProps) GetDeltaOk() (float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDelta

`func (o *EffectProps) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDelta

`func (o *EffectProps) SetDelta(v float32)`

SetDelta gets a reference to the given float32 and assigns it to the Delta field.

### GetTarget

`func (o *EffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EffectProps) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *EffectProps) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *EffectProps) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetIsJustCompleted

`func (o *EffectProps) GetIsJustCompleted() bool`

GetIsJustCompleted returns the IsJustCompleted field if non-nil, zero value otherwise.

### GetIsJustCompletedOk

`func (o *EffectProps) GetIsJustCompletedOk() (bool, bool)`

GetIsJustCompletedOk returns a tuple with the IsJustCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsJustCompleted

`func (o *EffectProps) HasIsJustCompleted() bool`

HasIsJustCompleted returns a boolean if a field has been set.

### SetIsJustCompleted

`func (o *EffectProps) SetIsJustCompleted(v bool)`

SetIsJustCompleted gets a reference to the given bool and assigns it to the IsJustCompleted field.

### GetDecreaseProgressBy

`func (o *EffectProps) GetDecreaseProgressBy() float32`

GetDecreaseProgressBy returns the DecreaseProgressBy field if non-nil, zero value otherwise.

### GetDecreaseProgressByOk

`func (o *EffectProps) GetDecreaseProgressByOk() (float32, bool)`

GetDecreaseProgressByOk returns a tuple with the DecreaseProgressBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDecreaseProgressBy

`func (o *EffectProps) HasDecreaseProgressBy() bool`

HasDecreaseProgressBy returns a boolean if a field has been set.

### SetDecreaseProgressBy

`func (o *EffectProps) SetDecreaseProgressBy(v float32)`

SetDecreaseProgressBy gets a reference to the given float32 and assigns it to the DecreaseProgressBy field.

### GetCurrentProgress

`func (o *EffectProps) GetCurrentProgress() float32`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *EffectProps) GetCurrentProgressOk() (float32, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentProgress

`func (o *EffectProps) HasCurrentProgress() bool`

HasCurrentProgress returns a boolean if a field has been set.

### SetCurrentProgress

`func (o *EffectProps) SetCurrentProgress(v float32)`

SetCurrentProgress gets a reference to the given float32 and assigns it to the CurrentProgress field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


