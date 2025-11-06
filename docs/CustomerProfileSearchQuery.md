# CustomerProfileSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Properties to match against a customer profile. All provided attributes will be exactly matched against profile attributes. | [optional] 
**IntegrationIDs** | Pointer to **[]string** |  | [optional] 
**ProfileIDs** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCustomerProfileSearchQuery

`func NewCustomerProfileSearchQuery() *CustomerProfileSearchQuery`

NewCustomerProfileSearchQuery instantiates a new CustomerProfileSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProfileSearchQueryWithDefaults

`func NewCustomerProfileSearchQueryWithDefaults() *CustomerProfileSearchQuery`

NewCustomerProfileSearchQueryWithDefaults instantiates a new CustomerProfileSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CustomerProfileSearchQuery) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerProfileSearchQuery) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerProfileSearchQuery) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CustomerProfileSearchQuery) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIntegrationIDs

`func (o *CustomerProfileSearchQuery) GetIntegrationIDs() []string`

GetIntegrationIDs returns the IntegrationIDs field if non-nil, zero value otherwise.

### GetIntegrationIDsOk

`func (o *CustomerProfileSearchQuery) GetIntegrationIDsOk() (*[]string, bool)`

GetIntegrationIDsOk returns a tuple with the IntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIDs

`func (o *CustomerProfileSearchQuery) SetIntegrationIDs(v []string)`

SetIntegrationIDs sets IntegrationIDs field to given value.

### HasIntegrationIDs

`func (o *CustomerProfileSearchQuery) HasIntegrationIDs() bool`

HasIntegrationIDs returns a boolean if a field has been set.

### GetProfileIDs

`func (o *CustomerProfileSearchQuery) GetProfileIDs() []int64`

GetProfileIDs returns the ProfileIDs field if non-nil, zero value otherwise.

### GetProfileIDsOk

`func (o *CustomerProfileSearchQuery) GetProfileIDsOk() (*[]int64, bool)`

GetProfileIDsOk returns a tuple with the ProfileIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIDs

`func (o *CustomerProfileSearchQuery) SetProfileIDs(v []int64)`

SetProfileIDs sets ProfileIDs field to given value.

### HasProfileIDs

`func (o *CustomerProfileSearchQuery) HasProfileIDs() bool`

HasProfileIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


