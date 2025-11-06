# ScimServiceProviderConfigResponseBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxOperations** | Pointer to **int64** | The maximum number of individual operations that can be included in a single bulk request. | [optional] 
**MaxPayloadSize** | Pointer to **int64** | The maximum size, in bytes, of the entire payload for a bulk operation request. | [optional] 
**Supported** | Pointer to **bool** | Indicates whether the SCIM service provider supports bulk operations. | [optional] 

## Methods

### NewScimServiceProviderConfigResponseBulk

`func NewScimServiceProviderConfigResponseBulk() *ScimServiceProviderConfigResponseBulk`

NewScimServiceProviderConfigResponseBulk instantiates a new ScimServiceProviderConfigResponseBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimServiceProviderConfigResponseBulkWithDefaults

`func NewScimServiceProviderConfigResponseBulkWithDefaults() *ScimServiceProviderConfigResponseBulk`

NewScimServiceProviderConfigResponseBulkWithDefaults instantiates a new ScimServiceProviderConfigResponseBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxOperations() int64`

GetMaxOperations returns the MaxOperations field if non-nil, zero value otherwise.

### GetMaxOperationsOk

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxOperationsOk() (*int64, bool)`

GetMaxOperationsOk returns a tuple with the MaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) SetMaxOperations(v int64)`

SetMaxOperations sets MaxOperations field to given value.

### HasMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) HasMaxOperations() bool`

HasMaxOperations returns a boolean if a field has been set.

### GetMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxPayloadSize() int64`

GetMaxPayloadSize returns the MaxPayloadSize field if non-nil, zero value otherwise.

### GetMaxPayloadSizeOk

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxPayloadSizeOk() (*int64, bool)`

GetMaxPayloadSizeOk returns a tuple with the MaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) SetMaxPayloadSize(v int64)`

SetMaxPayloadSize sets MaxPayloadSize field to given value.

### HasMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) HasMaxPayloadSize() bool`

HasMaxPayloadSize returns a boolean if a field has been set.

### GetSupported

`func (o *ScimServiceProviderConfigResponseBulk) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ScimServiceProviderConfigResponseBulk) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *ScimServiceProviderConfigResponseBulk) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *ScimServiceProviderConfigResponseBulk) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


