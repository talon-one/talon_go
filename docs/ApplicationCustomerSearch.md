# ApplicationCustomerSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Properties to match against a customer profile. All provided attributes will be exactly matched against profile attributes | [optional] 
**IntegrationIDs** | Pointer to **[]string** |  | [optional] 
**ProfileIDs** | Pointer to **[]int32** |  | [optional] 

## Methods

### GetAttributes

`func (o *ApplicationCustomerSearch) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationCustomerSearch) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *ApplicationCustomerSearch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *ApplicationCustomerSearch) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetIntegrationIDs

`func (o *ApplicationCustomerSearch) GetIntegrationIDs() []string`

GetIntegrationIDs returns the IntegrationIDs field if non-nil, zero value otherwise.

### GetIntegrationIDsOk

`func (o *ApplicationCustomerSearch) GetIntegrationIDsOk() ([]string, bool)`

GetIntegrationIDsOk returns a tuple with the IntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationIDs

`func (o *ApplicationCustomerSearch) HasIntegrationIDs() bool`

HasIntegrationIDs returns a boolean if a field has been set.

### SetIntegrationIDs

`func (o *ApplicationCustomerSearch) SetIntegrationIDs(v []string)`

SetIntegrationIDs gets a reference to the given []string and assigns it to the IntegrationIDs field.

### GetProfileIDs

`func (o *ApplicationCustomerSearch) GetProfileIDs() []int32`

GetProfileIDs returns the ProfileIDs field if non-nil, zero value otherwise.

### GetProfileIDsOk

`func (o *ApplicationCustomerSearch) GetProfileIDsOk() ([]int32, bool)`

GetProfileIDsOk returns a tuple with the ProfileIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIDs

`func (o *ApplicationCustomerSearch) HasProfileIDs() bool`

HasProfileIDs returns a boolean if a field has been set.

### SetProfileIDs

`func (o *ApplicationCustomerSearch) SetProfileIDs(v []int32)`

SetProfileIDs gets a reference to the given []int32 and assigns it to the ProfileIDs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


