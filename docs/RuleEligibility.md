# RuleEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | Pointer to **bool** | Indicates whether the customer was eligible for the rule in the current session, based on whether all of the rule&#39;s conditions were met. | 
**CouponCode** | Pointer to **string** | The coupon code used to check a customer&#39;s eligibility for the rule in the current session, if applicable. | [optional] 
**Details** | Pointer to [**RuleEligibilityFailureDetails**](RuleEligibilityFailureDetails.md) |  | [optional] 

## Methods

### NewRuleEligibility

`func NewRuleEligibility(passed bool, ) *RuleEligibility`

NewRuleEligibility instantiates a new RuleEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEligibilityWithDefaults

`func NewRuleEligibilityWithDefaults() *RuleEligibility`

NewRuleEligibilityWithDefaults instantiates a new RuleEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *RuleEligibility) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *RuleEligibility) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *RuleEligibility) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetCouponCode

`func (o *RuleEligibility) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *RuleEligibility) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *RuleEligibility) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *RuleEligibility) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetDetails

`func (o *RuleEligibility) GetDetails() RuleEligibilityFailureDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RuleEligibility) GetDetailsOk() (*RuleEligibilityFailureDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RuleEligibility) SetDetails(v RuleEligibilityFailureDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RuleEligibility) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


