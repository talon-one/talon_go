# SsoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforced** | Pointer to **bool** | An indication of whether single sign-on is enforced for the account. When enforced, users cannot use their email and password to sign in to Talon.One. It is not possible to change this to &#x60;false&#x60; after it is set to &#x60;true&#x60;.  | 
**NewAcsUrl** | Pointer to **string** | Assertion Consumer Service (ACS) URL for setting up a new SAML connection with an identity provider like Okta or Microsoft Entra ID.  | [optional] 

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

### GetNewAcsUrl

`func (o *SsoConfig) GetNewAcsUrl() string`

GetNewAcsUrl returns the NewAcsUrl field if non-nil, zero value otherwise.

### GetNewAcsUrlOk

`func (o *SsoConfig) GetNewAcsUrlOk() (string, bool)`

GetNewAcsUrlOk returns a tuple with the NewAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewAcsUrl

`func (o *SsoConfig) HasNewAcsUrl() bool`

HasNewAcsUrl returns a boolean if a field has been set.

### SetNewAcsUrl

`func (o *SsoConfig) SetNewAcsUrl(v string)`

SetNewAcsUrl gets a reference to the given string and assigns it to the NewAcsUrl field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


