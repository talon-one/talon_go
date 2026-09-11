# EventV3ReferralEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferralCode** | Pointer to **string** | The referral code submitted with the event. The endpoint does not validate the code, and submitting a code does not redeem it. Use the \&quot;Referral code is valid\&quot; condition in the Rule Builder to validate and redeem the code, or \&quot;Referral code is valid (without redemption)\&quot; to validate without redeeming.  | [optional] 

## Methods

### NewEventV3ReferralEntity

`func NewEventV3ReferralEntity() *EventV3ReferralEntity`

NewEventV3ReferralEntity instantiates a new EventV3ReferralEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventV3ReferralEntityWithDefaults

`func NewEventV3ReferralEntityWithDefaults() *EventV3ReferralEntity`

NewEventV3ReferralEntityWithDefaults instantiates a new EventV3ReferralEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferralCode

`func (o *EventV3ReferralEntity) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *EventV3ReferralEntity) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *EventV3ReferralEntity) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *EventV3ReferralEntity) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


