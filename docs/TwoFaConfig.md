# TwoFAConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | An indication of whether two-factor authentication is enabled for the account. | 
**RequireEverySignIn** | Pointer to **bool** | Can be &#x60;true&#x60; or &#x60;false&#x60;. - &#x60;true&#x60;: Two-factor authentication is required each time a user signs in to their Talon.One account. - &#x60;false&#x60;: Two-factor authentication is only required when a user signs in to their Talon.One account on a new device, and every 30 days after that.  | [optional] 

## Methods

### NewTwoFAConfig

`func NewTwoFAConfig(enabled bool, ) *TwoFAConfig`

NewTwoFAConfig instantiates a new TwoFAConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFAConfigWithDefaults

`func NewTwoFAConfigWithDefaults() *TwoFAConfig`

NewTwoFAConfigWithDefaults instantiates a new TwoFAConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TwoFAConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwoFAConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TwoFAConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRequireEverySignIn

`func (o *TwoFAConfig) GetRequireEverySignIn() bool`

GetRequireEverySignIn returns the RequireEverySignIn field if non-nil, zero value otherwise.

### GetRequireEverySignInOk

`func (o *TwoFAConfig) GetRequireEverySignInOk() (*bool, bool)`

GetRequireEverySignInOk returns a tuple with the RequireEverySignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEverySignIn

`func (o *TwoFAConfig) SetRequireEverySignIn(v bool)`

SetRequireEverySignIn sets RequireEverySignIn field to given value.

### HasRequireEverySignIn

`func (o *TwoFAConfig) HasRequireEverySignIn() bool`

HasRequireEverySignIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


