# NewSamlConnection

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
**X509certificate** | Pointer to **string** | X.509 Certificate. | 

## Methods

### GetAccountId

`func (o *NewSamlConnection) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NewSamlConnection) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *NewSamlConnection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *NewSamlConnection) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetAudienceURI

`func (o *NewSamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *NewSamlConnection) GetAudienceURIOk() (string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceURI

`func (o *NewSamlConnection) HasAudienceURI() bool`

HasAudienceURI returns a boolean if a field has been set.

### SetAudienceURI

`func (o *NewSamlConnection) SetAudienceURI(v string)`

SetAudienceURI gets a reference to the given string and assigns it to the AudienceURI field.

### GetEnabled

`func (o *NewSamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewSamlConnection) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *NewSamlConnection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *NewSamlConnection) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIssuer

`func (o *NewSamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *NewSamlConnection) GetIssuerOk() (string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuer

`func (o *NewSamlConnection) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuer

`func (o *NewSamlConnection) SetIssuer(v string)`

SetIssuer gets a reference to the given string and assigns it to the Issuer field.

### GetMetadataURL

`func (o *NewSamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *NewSamlConnection) GetMetadataURLOk() (string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadataURL

`func (o *NewSamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### SetMetadataURL

`func (o *NewSamlConnection) SetMetadataURL(v string)`

SetMetadataURL gets a reference to the given string and assigns it to the MetadataURL field.

### GetName

`func (o *NewSamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewSamlConnection) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewSamlConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewSamlConnection) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSignOnURL

`func (o *NewSamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *NewSamlConnection) GetSignOnURLOk() (string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOnURL

`func (o *NewSamlConnection) HasSignOnURL() bool`

HasSignOnURL returns a boolean if a field has been set.

### SetSignOnURL

`func (o *NewSamlConnection) SetSignOnURL(v string)`

SetSignOnURL gets a reference to the given string and assigns it to the SignOnURL field.

### GetSignOutURL

`func (o *NewSamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *NewSamlConnection) GetSignOutURLOk() (string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOutURL

`func (o *NewSamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### SetSignOutURL

`func (o *NewSamlConnection) SetSignOutURL(v string)`

SetSignOutURL gets a reference to the given string and assigns it to the SignOutURL field.

### GetX509certificate

`func (o *NewSamlConnection) GetX509certificate() string`

GetX509certificate returns the X509certificate field if non-nil, zero value otherwise.

### GetX509certificateOk

`func (o *NewSamlConnection) GetX509certificateOk() (string, bool)`

GetX509certificateOk returns a tuple with the X509certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasX509certificate

`func (o *NewSamlConnection) HasX509certificate() bool`

HasX509certificate returns a boolean if a field has been set.

### SetX509certificate

`func (o *NewSamlConnection) SetX509certificate(v string)`

SetX509certificate gets a reference to the given string and assigns it to the X509certificate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


