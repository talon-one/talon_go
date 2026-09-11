# AdditionalCostReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal identifier of the additional cost. | 
**Name** | Pointer to **string** | The additional cost name as used in API requests. | 
**Title** | Pointer to **string** | The human-readable title of the additional cost. | [optional] 

## Methods

### NewAdditionalCostReference

`func NewAdditionalCostReference(id int64, name string, ) *AdditionalCostReference`

NewAdditionalCostReference instantiates a new AdditionalCostReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalCostReferenceWithDefaults

`func NewAdditionalCostReferenceWithDefaults() *AdditionalCostReference`

NewAdditionalCostReferenceWithDefaults instantiates a new AdditionalCostReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdditionalCostReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdditionalCostReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdditionalCostReference) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AdditionalCostReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalCostReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalCostReference) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AdditionalCostReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdditionalCostReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdditionalCostReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdditionalCostReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


