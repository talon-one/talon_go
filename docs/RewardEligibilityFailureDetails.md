# RewardEligibilityFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCode** | Pointer to **string** | A code identifying why the customer is not eligible for the reward. | 
**ConditionIndex** | Pointer to **int64** | The index of the eligibility condition that the customer did not meet. Only applicable when &#x60;failureCode&#x60; is &#x60;CONDITION_NOT_MET&#x60;. | [optional] 

## Methods

### NewRewardEligibilityFailureDetails

`func NewRewardEligibilityFailureDetails(failureCode string, ) *RewardEligibilityFailureDetails`

NewRewardEligibilityFailureDetails instantiates a new RewardEligibilityFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardEligibilityFailureDetailsWithDefaults

`func NewRewardEligibilityFailureDetailsWithDefaults() *RewardEligibilityFailureDetails`

NewRewardEligibilityFailureDetailsWithDefaults instantiates a new RewardEligibilityFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCode

`func (o *RewardEligibilityFailureDetails) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *RewardEligibilityFailureDetails) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *RewardEligibilityFailureDetails) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### GetConditionIndex

`func (o *RewardEligibilityFailureDetails) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RewardEligibilityFailureDetails) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *RewardEligibilityFailureDetails) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *RewardEligibilityFailureDetails) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


