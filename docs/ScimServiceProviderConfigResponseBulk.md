# ScimServiceProviderConfigResponseBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxOperations** | Pointer to **int32** | The maximum number of individual operations that can be included in a single bulk request. | [optional] 
**MaxPayloadSize** | Pointer to **int32** | The maximum size, in bytes, of the entire payload for a bulk operation request. | [optional] 
**Supported** | Pointer to **bool** | Indicates whether the SCIM service provider supports bulk operations. | [optional] 

## Methods

### GetMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxOperations() int32`

GetMaxOperations returns the MaxOperations field if non-nil, zero value otherwise.

### GetMaxOperationsOk

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxOperationsOk() (int32, bool)`

GetMaxOperationsOk returns a tuple with the MaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) HasMaxOperations() bool`

HasMaxOperations returns a boolean if a field has been set.

### SetMaxOperations

`func (o *ScimServiceProviderConfigResponseBulk) SetMaxOperations(v int32)`

SetMaxOperations gets a reference to the given int32 and assigns it to the MaxOperations field.

### GetMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxPayloadSize() int32`

GetMaxPayloadSize returns the MaxPayloadSize field if non-nil, zero value otherwise.

### GetMaxPayloadSizeOk

`func (o *ScimServiceProviderConfigResponseBulk) GetMaxPayloadSizeOk() (int32, bool)`

GetMaxPayloadSizeOk returns a tuple with the MaxPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) HasMaxPayloadSize() bool`

HasMaxPayloadSize returns a boolean if a field has been set.

### SetMaxPayloadSize

`func (o *ScimServiceProviderConfigResponseBulk) SetMaxPayloadSize(v int32)`

SetMaxPayloadSize gets a reference to the given int32 and assigns it to the MaxPayloadSize field.

### GetSupported

`func (o *ScimServiceProviderConfigResponseBulk) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ScimServiceProviderConfigResponseBulk) GetSupportedOk() (bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupported

`func (o *ScimServiceProviderConfigResponseBulk) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### SetSupported

`func (o *ScimServiceProviderConfigResponseBulk) SetSupported(v bool)`

SetSupported gets a reference to the given bool and assigns it to the Supported field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


