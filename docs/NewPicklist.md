# NewPicklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of allowed values in the picklist. If the type &#x60;time&#x60; is chosen, it must be an RFC3339 timestamp string. | 
**Values** | Pointer to **[]string** | The list of allowed values provided by this picklist. | 

## Methods

### NewNewPicklist

`func NewNewPicklist(type_ string, values []string, ) *NewPicklist`

NewNewPicklist instantiates a new NewPicklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewPicklistWithDefaults

`func NewNewPicklistWithDefaults() *NewPicklist`

NewNewPicklistWithDefaults instantiates a new NewPicklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NewPicklist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewPicklist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewPicklist) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *NewPicklist) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *NewPicklist) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *NewPicklist) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


