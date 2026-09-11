# RewardEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | Pointer to **bool** | Indicates whether the customer is eligible for the reward. | 
**Details** | Pointer to [**[]RewardEligibilityFailureDetails**](RewardEligibilityFailureDetails.md) | The reasons the customer is not eligible for the reward. Empty when &#x60;passed&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### NewRewardEligibility

`func NewRewardEligibility(passed bool, ) *RewardEligibility`

NewRewardEligibility instantiates a new RewardEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardEligibilityWithDefaults

`func NewRewardEligibilityWithDefaults() *RewardEligibility`

NewRewardEligibilityWithDefaults instantiates a new RewardEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *RewardEligibility) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *RewardEligibility) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *RewardEligibility) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetDetails

`func (o *RewardEligibility) GetDetails() []RewardEligibilityFailureDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RewardEligibility) GetDetailsOk() (*[]RewardEligibilityFailureDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RewardEligibility) SetDetails(v []RewardEligibilityFailureDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RewardEligibility) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


