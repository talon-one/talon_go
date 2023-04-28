# SamlConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssertionConsumerServiceURL** | Pointer to **string** | The location where the SAML assertion is sent with a HTTP POST. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | ID of the SAML service. | 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**Issuer** | Pointer to **string** | Identity Provider Entity ID. | 
**SignOnURL** | Pointer to **string** | Single Sign-On URL. | 
**SignOutURL** | Pointer to **string** | Single Sign-Out URL. | [optional] 
**MetadataURL** | Pointer to **string** | Metadata URL. | [optional] 
**AudienceURI** | Pointer to **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 

## Methods

### GetAssertionConsumerServiceURL

`func (o *SamlConnection) GetAssertionConsumerServiceURL() string`

GetAssertionConsumerServiceURL returns the AssertionConsumerServiceURL field if non-nil, zero value otherwise.

### GetAssertionConsumerServiceURLOk

`func (o *SamlConnection) GetAssertionConsumerServiceURLOk() (string, bool)`

GetAssertionConsumerServiceURLOk returns a tuple with the AssertionConsumerServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssertionConsumerServiceURL

`func (o *SamlConnection) HasAssertionConsumerServiceURL() bool`

HasAssertionConsumerServiceURL returns a boolean if a field has been set.

### SetAssertionConsumerServiceURL

`func (o *SamlConnection) SetAssertionConsumerServiceURL(v string)`

SetAssertionConsumerServiceURL gets a reference to the given string and assigns it to the AssertionConsumerServiceURL field.

### GetAccountId

`func (o *SamlConnection) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SamlConnection) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *SamlConnection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *SamlConnection) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetName

`func (o *SamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlConnection) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SamlConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SamlConnection) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetEnabled

`func (o *SamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlConnection) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *SamlConnection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *SamlConnection) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIssuer

`func (o *SamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlConnection) GetIssuerOk() (string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuer

`func (o *SamlConnection) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuer

`func (o *SamlConnection) SetIssuer(v string)`

SetIssuer gets a reference to the given string and assigns it to the Issuer field.

### GetSignOnURL

`func (o *SamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *SamlConnection) GetSignOnURLOk() (string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOnURL

`func (o *SamlConnection) HasSignOnURL() bool`

HasSignOnURL returns a boolean if a field has been set.

### SetSignOnURL

`func (o *SamlConnection) SetSignOnURL(v string)`

SetSignOnURL gets a reference to the given string and assigns it to the SignOnURL field.

### GetSignOutURL

`func (o *SamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *SamlConnection) GetSignOutURLOk() (string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignOutURL

`func (o *SamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### SetSignOutURL

`func (o *SamlConnection) SetSignOutURL(v string)`

SetSignOutURL gets a reference to the given string and assigns it to the SignOutURL field.

### GetMetadataURL

`func (o *SamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *SamlConnection) GetMetadataURLOk() (string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetadataURL

`func (o *SamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### SetMetadataURL

`func (o *SamlConnection) SetMetadataURL(v string)`

SetMetadataURL gets a reference to the given string and assigns it to the MetadataURL field.

### GetAudienceURI

`func (o *SamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *SamlConnection) GetAudienceURIOk() (string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudienceURI

`func (o *SamlConnection) HasAudienceURI() bool`

HasAudienceURI returns a boolean if a field has been set.

### SetAudienceURI

`func (o *SamlConnection) SetAudienceURI(v string)`

SetAudienceURI gets a reference to the given string and assigns it to the AudienceURI field.

### GetId

`func (o *SamlConnection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlConnection) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *SamlConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *SamlConnection) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *SamlConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SamlConnection) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *SamlConnection) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *SamlConnection) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


