# ScimServiceProviderConfigResponseFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** | The maximum number of resources that can be returned in a single filtered query response. | [optional] 
**Supported** | Pointer to **bool** | Indicates whether the SCIM service provider supports filtering operations. | [optional] 

## Methods

### GetMaxResults

`func (o *ScimServiceProviderConfigResponseFilter) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *ScimServiceProviderConfigResponseFilter) GetMaxResultsOk() (int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxResults

`func (o *ScimServiceProviderConfigResponseFilter) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### SetMaxResults

`func (o *ScimServiceProviderConfigResponseFilter) SetMaxResults(v int32)`

SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.

### GetSupported

`func (o *ScimServiceProviderConfigResponseFilter) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ScimServiceProviderConfigResponseFilter) GetSupportedOk() (bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupported

`func (o *ScimServiceProviderConfigResponseFilter) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### SetSupported

`func (o *ScimServiceProviderConfigResponseFilter) SetSupported(v bool)`

SetSupported gets a reference to the given bool and assigns it to the Supported field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


