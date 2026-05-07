# SSOConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforced** | Pointer to **bool** | An indication of whether single sign-on is enforced for the account. When enforced, users cannot use their email and password to sign in to Talon.One. It is not possible to change this to &#x60;false&#x60; after it is set to &#x60;true&#x60;.  | 
**NewAcsUrl** | Pointer to **string** | Assertion Consumer Service (ACS) URL for setting up a new SAML connection with an identity provider like Okta or Microsoft Entra ID.  | [optional] 

## Methods

### NewSSOConfig

`func NewSSOConfig(enforced bool, ) *SSOConfig`

NewSSOConfig instantiates a new SSOConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSOConfigWithDefaults

`func NewSSOConfigWithDefaults() *SSOConfig`

NewSSOConfigWithDefaults instantiates a new SSOConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforced

`func (o *SSOConfig) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *SSOConfig) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *SSOConfig) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.


### GetNewAcsUrl

`func (o *SSOConfig) GetNewAcsUrl() string`

GetNewAcsUrl returns the NewAcsUrl field if non-nil, zero value otherwise.

### GetNewAcsUrlOk

`func (o *SSOConfig) GetNewAcsUrlOk() (*string, bool)`

GetNewAcsUrlOk returns a tuple with the NewAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAcsUrl

`func (o *SSOConfig) SetNewAcsUrl(v string)`

SetNewAcsUrl sets NewAcsUrl field to given value.

### HasNewAcsUrl

`func (o *SSOConfig) HasNewAcsUrl() bool`

HasNewAcsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


