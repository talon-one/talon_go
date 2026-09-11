# CheckTierBlockTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the tier. | 
**Name** | Pointer to **string** | The display name of the tier. | 
**MinPoints** | Pointer to **float32** | The minimum amount of points required to enter the tier. | 
**UpperLimit** | Pointer to **float32** |  | [optional] 

## Methods

### NewCheckTierBlockTier

`func NewCheckTierBlockTier(id int64, name string, minPoints float32, ) *CheckTierBlockTier`

NewCheckTierBlockTier instantiates a new CheckTierBlockTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTierBlockTierWithDefaults

`func NewCheckTierBlockTierWithDefaults() *CheckTierBlockTier`

NewCheckTierBlockTierWithDefaults instantiates a new CheckTierBlockTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckTierBlockTier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTierBlockTier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTierBlockTier) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CheckTierBlockTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckTierBlockTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckTierBlockTier) SetName(v string)`

SetName sets Name field to given value.


### GetMinPoints

`func (o *CheckTierBlockTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *CheckTierBlockTier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *CheckTierBlockTier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.


### GetUpperLimit

`func (o *CheckTierBlockTier) GetUpperLimit() float32`

GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.

### GetUpperLimitOk

`func (o *CheckTierBlockTier) GetUpperLimitOk() (*float32, bool)`

GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLimit

`func (o *CheckTierBlockTier) SetUpperLimit(v float32)`

SetUpperLimit sets UpperLimit field to given value.

### HasUpperLimit

`func (o *CheckTierBlockTier) HasUpperLimit() bool`

HasUpperLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


