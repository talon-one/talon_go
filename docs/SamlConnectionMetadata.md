# SamlConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ID of the SAML service. | 
**Enabled** | Pointer to **bool** | Determines if this SAML connection active. | 
**AccountId** | Pointer to **float32** |  | 
**MetadataDocument** | Pointer to **string** | Identity Provider metadata XML document. | 

## Methods

### NewSamlConnectionMetadata

`func NewSamlConnectionMetadata(name string, enabled bool, accountId float32, metadataDocument string, ) *SamlConnectionMetadata`

NewSamlConnectionMetadata instantiates a new SamlConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConnectionMetadataWithDefaults

`func NewSamlConnectionMetadataWithDefaults() *SamlConnectionMetadata`

NewSamlConnectionMetadataWithDefaults instantiates a new SamlConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlConnectionMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlConnectionMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlConnectionMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *SamlConnectionMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlConnectionMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlConnectionMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAccountId

`func (o *SamlConnectionMetadata) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SamlConnectionMetadata) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SamlConnectionMetadata) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetMetadataDocument

`func (o *SamlConnectionMetadata) GetMetadataDocument() string`

GetMetadataDocument returns the MetadataDocument field if non-nil, zero value otherwise.

### GetMetadataDocumentOk

`func (o *SamlConnectionMetadata) GetMetadataDocumentOk() (*string, bool)`

GetMetadataDocumentOk returns a tuple with the MetadataDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataDocument

`func (o *SamlConnectionMetadata) SetMetadataDocument(v string)`

SetMetadataDocument sets MetadataDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


