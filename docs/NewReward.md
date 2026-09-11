# NewReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the reward. | 
**ApiName** | Pointer to **string** | A unique identifier used to reference the reward in API integrations. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications this reward is connected to.   **Note**: Currently, a reward can only be connected to one Application.  | 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox reward. Rewards of a given type can only be connected to Applications of the same type. | 
**EligibilityConditions** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated. Each binding pairs a name with a talang expression. The expression is evaluated once and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. Each object defines the specific loyalty program and subledger from which points are deducted when activating the reward.  **Note:** When creating a reward, the &#x60;id&#x60; of each entry is ignored and a new entry is always created.  | [optional] 

## Methods

### NewNewReward

`func NewNewReward(name string, apiName string, applicationIds []int64, sandbox bool, ) *NewReward`

NewNewReward instantiates a new NewReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRewardWithDefaults

`func NewNewRewardWithDefaults() *NewReward`

NewNewRewardWithDefaults instantiates a new NewReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewReward) SetName(v string)`

SetName sets Name field to given value.


### GetApiName

`func (o *NewReward) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *NewReward) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *NewReward) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetDescription

`func (o *NewReward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewReward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewReward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewReward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *NewReward) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewReward) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewReward) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetSandbox

`func (o *NewReward) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewReward) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewReward) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetEligibilityConditions

`func (o *NewReward) GetEligibilityConditions() Rule`

GetEligibilityConditions returns the EligibilityConditions field if non-nil, zero value otherwise.

### GetEligibilityConditionsOk

`func (o *NewReward) GetEligibilityConditionsOk() (*Rule, bool)`

GetEligibilityConditionsOk returns a tuple with the EligibilityConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityConditions

`func (o *NewReward) SetEligibilityConditions(v Rule)`

SetEligibilityConditions sets EligibilityConditions field to given value.

### HasEligibilityConditions

`func (o *NewReward) HasEligibilityConditions() bool`

HasEligibilityConditions returns a boolean if a field has been set.

### GetRule

`func (o *NewReward) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *NewReward) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *NewReward) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *NewReward) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetBindings

`func (o *NewReward) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *NewReward) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *NewReward) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *NewReward) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetPointsRequired

`func (o *NewReward) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *NewReward) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *NewReward) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *NewReward) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


