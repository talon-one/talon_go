# TwoFaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | An indication of whether two-factor authentication is enabled for the account. | 
**RequireEverySignIn** | Pointer to **bool** | Can be &#x60;true&#x60; or &#x60;false&#x60;. - &#x60;true&#x60;: Two-factor authentication is required each time a user signs in to their Talon.One account. - &#x60;false&#x60;: Two-factor authentication is only required when a user signs in to their Talon.One account on a new device, and every 30 days after that.  | [optional] 

## Methods

### GetEnabled

`func (o *TwoFaConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TwoFaConfig) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *TwoFaConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *TwoFaConfig) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetRequireEverySignIn

`func (o *TwoFaConfig) GetRequireEverySignIn() bool`

GetRequireEverySignIn returns the RequireEverySignIn field if non-nil, zero value otherwise.

### GetRequireEverySignInOk

`func (o *TwoFaConfig) GetRequireEverySignInOk() (bool, bool)`

GetRequireEverySignInOk returns a tuple with the RequireEverySignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequireEverySignIn

`func (o *TwoFaConfig) HasRequireEverySignIn() bool`

HasRequireEverySignIn returns a boolean if a field has been set.

### SetRequireEverySignIn

`func (o *TwoFaConfig) SetRequireEverySignIn(v bool)`

SetRequireEverySignIn gets a reference to the given bool and assigns it to the RequireEverySignIn field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


