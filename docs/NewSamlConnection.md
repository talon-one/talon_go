# NewSamlConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X509certificate** | Pointer to **string** | X.509 Certificate. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | ID of the SAML service. | 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**Issuer** | Pointer to **string** | Identity Provider Entity ID. | 
**SignOnURL** | Pointer to **string** | Single Sign-On URL. | 
**SignOutURL** | Pointer to **string** | Single Sign-Out URL. | [optional] 
**MetadataURL** | Pointer to **string** | Metadata URL. | [optional] 
**AudienceURI** | Pointer to **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | [optional] 

## Methods

### NewNewSamlConnection

`func NewNewSamlConnection(x509certificate string, accountId int64, name string, enabled bool, issuer string, signOnURL string, ) *NewSamlConnection`

NewNewSamlConnection instantiates a new NewSamlConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewSamlConnectionWithDefaults

`func NewNewSamlConnectionWithDefaults() *NewSamlConnection`

NewNewSamlConnectionWithDefaults instantiates a new NewSamlConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX509certificate

`func (o *NewSamlConnection) GetX509certificate() string`

GetX509certificate returns the X509certificate field if non-nil, zero value otherwise.

### GetX509certificateOk

`func (o *NewSamlConnection) GetX509certificateOk() (*string, bool)`

GetX509certificateOk returns a tuple with the X509certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509certificate

`func (o *NewSamlConnection) SetX509certificate(v string)`

SetX509certificate sets X509certificate field to given value.


### GetAccountId

`func (o *NewSamlConnection) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NewSamlConnection) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NewSamlConnection) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *NewSamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewSamlConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewSamlConnection) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *NewSamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewSamlConnection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewSamlConnection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIssuer

`func (o *NewSamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *NewSamlConnection) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *NewSamlConnection) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSignOnURL

`func (o *NewSamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *NewSamlConnection) GetSignOnURLOk() (*string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnURL

`func (o *NewSamlConnection) SetSignOnURL(v string)`

SetSignOnURL sets SignOnURL field to given value.


### GetSignOutURL

`func (o *NewSamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *NewSamlConnection) GetSignOutURLOk() (*string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutURL

`func (o *NewSamlConnection) SetSignOutURL(v string)`

SetSignOutURL sets SignOutURL field to given value.

### HasSignOutURL

`func (o *NewSamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### GetMetadataURL

`func (o *NewSamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *NewSamlConnection) GetMetadataURLOk() (*string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataURL

`func (o *NewSamlConnection) SetMetadataURL(v string)`

SetMetadataURL sets MetadataURL field to given value.

### HasMetadataURL

`func (o *NewSamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### GetAudienceURI

`func (o *NewSamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *NewSamlConnection) GetAudienceURIOk() (*string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceURI

`func (o *NewSamlConnection) SetAudienceURI(v string)`

SetAudienceURI sets AudienceURI field to given value.

### HasAudienceURI

`func (o *NewSamlConnection) HasAudienceURI() bool`

HasAudienceURI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


