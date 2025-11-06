# SamlConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssertionConsumerServiceURL** | Pointer to **string** | The location where the SAML assertion is sent with a HTTP POST. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | ID of the SAML service. | 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**Issuer** | Pointer to **string** | Identity Provider Entity ID. | 
**SignOnURL** | Pointer to **string** | Single Sign-On URL. | 
**SignOutURL** | Pointer to **string** | Single Sign-Out URL. | [optional] 
**MetadataURL** | Pointer to **string** | Metadata URL. | [optional] 
**AudienceURI** | Pointer to **string** | The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.  | 
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 

## Methods

### NewSamlConnection

`func NewSamlConnection(assertionConsumerServiceURL string, accountId int64, name string, enabled bool, issuer string, signOnURL string, audienceURI string, id int64, created time.Time, ) *SamlConnection`

NewSamlConnection instantiates a new SamlConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConnectionWithDefaults

`func NewSamlConnectionWithDefaults() *SamlConnection`

NewSamlConnectionWithDefaults instantiates a new SamlConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertionConsumerServiceURL

`func (o *SamlConnection) GetAssertionConsumerServiceURL() string`

GetAssertionConsumerServiceURL returns the AssertionConsumerServiceURL field if non-nil, zero value otherwise.

### GetAssertionConsumerServiceURLOk

`func (o *SamlConnection) GetAssertionConsumerServiceURLOk() (*string, bool)`

GetAssertionConsumerServiceURLOk returns a tuple with the AssertionConsumerServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerServiceURL

`func (o *SamlConnection) SetAssertionConsumerServiceURL(v string)`

SetAssertionConsumerServiceURL sets AssertionConsumerServiceURL field to given value.


### GetAccountId

`func (o *SamlConnection) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SamlConnection) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SamlConnection) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *SamlConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlConnection) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *SamlConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlConnection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlConnection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIssuer

`func (o *SamlConnection) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlConnection) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlConnection) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSignOnURL

`func (o *SamlConnection) GetSignOnURL() string`

GetSignOnURL returns the SignOnURL field if non-nil, zero value otherwise.

### GetSignOnURLOk

`func (o *SamlConnection) GetSignOnURLOk() (*string, bool)`

GetSignOnURLOk returns a tuple with the SignOnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnURL

`func (o *SamlConnection) SetSignOnURL(v string)`

SetSignOnURL sets SignOnURL field to given value.


### GetSignOutURL

`func (o *SamlConnection) GetSignOutURL() string`

GetSignOutURL returns the SignOutURL field if non-nil, zero value otherwise.

### GetSignOutURLOk

`func (o *SamlConnection) GetSignOutURLOk() (*string, bool)`

GetSignOutURLOk returns a tuple with the SignOutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutURL

`func (o *SamlConnection) SetSignOutURL(v string)`

SetSignOutURL sets SignOutURL field to given value.

### HasSignOutURL

`func (o *SamlConnection) HasSignOutURL() bool`

HasSignOutURL returns a boolean if a field has been set.

### GetMetadataURL

`func (o *SamlConnection) GetMetadataURL() string`

GetMetadataURL returns the MetadataURL field if non-nil, zero value otherwise.

### GetMetadataURLOk

`func (o *SamlConnection) GetMetadataURLOk() (*string, bool)`

GetMetadataURLOk returns a tuple with the MetadataURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataURL

`func (o *SamlConnection) SetMetadataURL(v string)`

SetMetadataURL sets MetadataURL field to given value.

### HasMetadataURL

`func (o *SamlConnection) HasMetadataURL() bool`

HasMetadataURL returns a boolean if a field has been set.

### GetAudienceURI

`func (o *SamlConnection) GetAudienceURI() string`

GetAudienceURI returns the AudienceURI field if non-nil, zero value otherwise.

### GetAudienceURIOk

`func (o *SamlConnection) GetAudienceURIOk() (*string, bool)`

GetAudienceURIOk returns a tuple with the AudienceURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceURI

`func (o *SamlConnection) SetAudienceURI(v string)`

SetAudienceURI sets AudienceURI field to given value.


### GetId

`func (o *SamlConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlConnection) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *SamlConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SamlConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SamlConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


