# UpdateReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the reward. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**Status** | Pointer to **string** | The status of the reward. | 
**EligibilityConditions** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated.  Each binding pairs a name with a talang expression. The expression is evaluated once  and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. Each object defines the specific loyalty program and subledger from which points are deducted when activating the reward.  **Note:** - Objects with an &#x60;id&#x60; are updated. - Objects without an &#x60;id&#x60; are created. - Existing objects omitted from the payload are deleted.  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the customer rewards unlocked from this reward remain active, relative to when the customer reward was unlocked.  The time format is an **integer** followed by one letter indicating the time unit.  Examples: &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  This field cannot be used in combination with &#x60;expiryDate&#x60;. If both &#x60;validityDuration&#x60; and &#x60;expiryDate&#x60; are not set, customer rewards have unlimited validity.  | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the customer rewards unlocked from this reward. Customer rewards never expire if this is omitted and &#x60;validityDuration&#x60; is not set.  **Note:** - The date must be in the future. - This field cannot be used in combination with &#x60;validityDuration&#x60;.  | [optional] 

## Methods

### NewUpdateReward

`func NewUpdateReward(name string, status string, ) *UpdateReward`

NewUpdateReward instantiates a new UpdateReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRewardWithDefaults

`func NewUpdateRewardWithDefaults() *UpdateReward`

NewUpdateRewardWithDefaults instantiates a new UpdateReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReward) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateReward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateReward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateReward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateReward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateReward) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateReward) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateReward) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEligibilityConditions

`func (o *UpdateReward) GetEligibilityConditions() Rule`

GetEligibilityConditions returns the EligibilityConditions field if non-nil, zero value otherwise.

### GetEligibilityConditionsOk

`func (o *UpdateReward) GetEligibilityConditionsOk() (*Rule, bool)`

GetEligibilityConditionsOk returns a tuple with the EligibilityConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityConditions

`func (o *UpdateReward) SetEligibilityConditions(v Rule)`

SetEligibilityConditions sets EligibilityConditions field to given value.

### HasEligibilityConditions

`func (o *UpdateReward) HasEligibilityConditions() bool`

HasEligibilityConditions returns a boolean if a field has been set.

### GetRule

`func (o *UpdateReward) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UpdateReward) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UpdateReward) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *UpdateReward) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetBindings

`func (o *UpdateReward) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *UpdateReward) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *UpdateReward) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *UpdateReward) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetPointsRequired

`func (o *UpdateReward) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *UpdateReward) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *UpdateReward) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *UpdateReward) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.

### GetValidityDuration

`func (o *UpdateReward) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *UpdateReward) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *UpdateReward) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *UpdateReward) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetExpiryDate

`func (o *UpdateReward) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *UpdateReward) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *UpdateReward) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *UpdateReward) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


