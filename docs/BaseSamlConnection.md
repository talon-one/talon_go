# BaseSamlConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**AudienceURI** | Pointer to **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | [optional] 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**Issuer** | Pointer to **string** | Identity Provider Entity ID. | 
**MetadataURL** | Pointer to **string** | Metadata URL. | [optional] 
**Name** | Pointer to **string** | ID of the SAML service. | 
**SignOnURL** | Pointer to **string** | Single Sign-On URL. | 
**SignOutURL** | Pointer to **string** | Single Sign-Out URL. | [optional] 

## Methods

### GetAccountId

`func (o *BaseSamlConnection) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseSamlConnection) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *BaseSamlConnection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *BaseSamlConnection) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetAudienceURI

`func (o *BaseSamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *BaseSamlConnection) GetAudienceURIOk() (string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceURI

`func (o *BaseSamlConnection) HasAudienceURI() bool`

HasAudienceURI returns a boolean if a field has been set.

### SetAudienceURI

`func (o *BaseSamlConnection) SetAudienceURI(v string)`

SetAudienceURI gets a reference to the given string and assigns it to the AudienceURI field.

### GetEnabled

`func (o *BaseSamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseSamlConnection) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *BaseSamlConnection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *BaseSamlConnection) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIssuer

`func (o *BaseSamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *BaseSamlConnection) GetIssuerOk() (string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuer

`func (o *BaseSamlConnection) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuer

`func (o *BaseSamlConnection) SetIssuer(v string)`

SetIssuer gets a reference to the given string and assigns it to the Issuer field.

### GetMetadataURL

`func (o *BaseSamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *BaseSamlConnection) GetMetadataURLOk() (string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadataURL

`func (o *BaseSamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### SetMetadataURL

`func (o *BaseSamlConnection) SetMetadataURL(v string)`

SetMetadataURL gets a reference to the given string and assigns it to the MetadataURL field.

### GetName

`func (o *BaseSamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseSamlConnection) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *BaseSamlConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *BaseSamlConnection) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSignOnURL

`func (o *BaseSamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *BaseSamlConnection) GetSignOnURLOk() (string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOnURL

`func (o *BaseSamlConnection) HasSignOnURL() bool`

HasSignOnURL returns a boolean if a field has been set.

### SetSignOnURL

`func (o *BaseSamlConnection) SetSignOnURL(v string)`

SetSignOnURL gets a reference to the given string and assigns it to the SignOnURL field.

### GetSignOutURL

`func (o *BaseSamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *BaseSamlConnection) GetSignOutURLOk() (string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOutURL

`func (o *BaseSamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### SetSignOutURL

`func (o *BaseSamlConnection) SetSignOutURL(v string)`

SetSignOutURL gets a reference to the given string and assigns it to the SignOutURL field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


