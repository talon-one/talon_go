# BaseSamlConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | ID of the SAML service. | 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**Issuer** | Pointer to **string** | Identity Provider Entity ID. | 
**SignOnURL** | Pointer to **string** | Single Sign-On URL. | 
**SignOutURL** | Pointer to **string** | Single Sign-Out URL. | [optional] 
**MetadataURL** | Pointer to **string** | Metadata URL. | [optional] 
**AudienceURI** | Pointer to **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | [optional] 

## Methods

### NewBaseSamlConnection

`func NewBaseSamlConnection(accountId int64, name string, enabled bool, issuer string, signOnURL string, ) *BaseSamlConnection`

NewBaseSamlConnection instantiates a new BaseSamlConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSamlConnectionWithDefaults

`func NewBaseSamlConnectionWithDefaults() *BaseSamlConnection`

NewBaseSamlConnectionWithDefaults instantiates a new BaseSamlConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BaseSamlConnection) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseSamlConnection) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BaseSamlConnection) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *BaseSamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseSamlConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseSamlConnection) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *BaseSamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseSamlConnection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseSamlConnection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIssuer

`func (o *BaseSamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *BaseSamlConnection) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *BaseSamlConnection) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSignOnURL

`func (o *BaseSamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *BaseSamlConnection) GetSignOnURLOk() (*string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnURL

`func (o *BaseSamlConnection) SetSignOnURL(v string)`

SetSignOnURL sets SignOnURL field to given value.


### GetSignOutURL

`func (o *BaseSamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *BaseSamlConnection) GetSignOutURLOk() (*string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutURL

`func (o *BaseSamlConnection) SetSignOutURL(v string)`

SetSignOutURL sets SignOutURL field to given value.

### HasSignOutURL

`func (o *BaseSamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### GetMetadataURL

`func (o *BaseSamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *BaseSamlConnection) GetMetadataURLOk() (*string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataURL

`func (o *BaseSamlConnection) SetMetadataURL(v string)`

SetMetadataURL sets MetadataURL field to given value.

### HasMetadataURL

`func (o *BaseSamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### GetAudienceURI

`func (o *BaseSamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *BaseSamlConnection) GetAudienceURIOk() (*string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceURI

`func (o *BaseSamlConnection) SetAudienceURI(v string)`

SetAudienceURI sets AudienceURI field to given value.

### HasAudienceURI

`func (o *BaseSamlConnection) HasAudienceURI() bool`

HasAudienceURI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


