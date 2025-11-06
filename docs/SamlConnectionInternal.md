# SamlConnectionInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ID of the SAML service. | 
**MetadataDocument** | Pointer to **string** | Identity Provider metadata XML document. | 

## Methods

### NewSamlConnectionInternal

`func NewSamlConnectionInternal(name string, metadataDocument string, ) *SamlConnectionInternal`

NewSamlConnectionInternal instantiates a new SamlConnectionInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConnectionInternalWithDefaults

`func NewSamlConnectionInternalWithDefaults() *SamlConnectionInternal`

NewSamlConnectionInternalWithDefaults instantiates a new SamlConnectionInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlConnectionInternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlConnectionInternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlConnectionInternal) SetName(v string)`

SetName sets Name field to given value.


### GetMetadataDocument

`func (o *SamlConnectionInternal) GetMetadataDocument() string`

GetMetadataDocument returns the MetadataDocument field if non-nil, zero value otherwise.

### GetMetadataDocumentOk

`func (o *SamlConnectionInternal) GetMetadataDocumentOk() (*string, bool)`

GetMetadataDocumentOk returns a tuple with the MetadataDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataDocument

`func (o *SamlConnectionInternal) SetMetadataDocument(v string)`

SetMetadataDocument sets MetadataDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


