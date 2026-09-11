# RewardUnlockRejection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A human-readable summary of why the reward unlock was rejected. | 
**RuleFailureReasons** | Pointer to [**[]RuleFailureReason**](RuleFailureReason.md) | The reasons why the reward could not be unlocked. | 

## Methods

### NewRewardUnlockRejection

`func NewRewardUnlockRejection(message string, ruleFailureReasons []RuleFailureReason, ) *RewardUnlockRejection`

NewRewardUnlockRejection instantiates a new RewardUnlockRejection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardUnlockRejectionWithDefaults

`func NewRewardUnlockRejectionWithDefaults() *RewardUnlockRejection`

NewRewardUnlockRejectionWithDefaults instantiates a new RewardUnlockRejection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RewardUnlockRejection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RewardUnlockRejection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RewardUnlockRejection) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRuleFailureReasons

`func (o *RewardUnlockRejection) GetRuleFailureReasons() []RuleFailureReason`

GetRuleFailureReasons returns the RuleFailureReasons field if non-nil, zero value otherwise.

### GetRuleFailureReasonsOk

`func (o *RewardUnlockRejection) GetRuleFailureReasonsOk() (*[]RuleFailureReason, bool)`

GetRuleFailureReasonsOk returns a tuple with the RuleFailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleFailureReasons

`func (o *RewardUnlockRejection) SetRuleFailureReasons(v []RuleFailureReason)`

SetRuleFailureReasons sets RuleFailureReasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


