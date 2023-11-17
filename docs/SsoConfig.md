# SsoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforced** | Pointer to **bool** | An indication of whether single sign-on is enforced for the account. When enforced, users cannot use their email and password to sign in to Talon.One. It is not possible to change this to &#x60;false&#x60; after it is set to &#x60;true&#x60;.  | 

## Methods

### GetEnforced

`func (o *SsoConfig) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *SsoConfig) GetEnforcedOk() (bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforced

`func (o *SsoConfig) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.

### SetEnforced

`func (o *SsoConfig) SetEnforced(v bool)`

SetEnforced gets a reference to the given bool and assigns it to the Enforced field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


