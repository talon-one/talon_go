# UpdateLoyaltyProgramTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the tier. | 
**Name** | Pointer to **string** | The name of the tier. | [optional] 
**MinPoints** | Pointer to **float32** | The minimum amount of points required to enter the tier. | [optional] 

## Methods

### NewUpdateLoyaltyProgramTier

`func NewUpdateLoyaltyProgramTier(id int64, ) *UpdateLoyaltyProgramTier`

NewUpdateLoyaltyProgramTier instantiates a new UpdateLoyaltyProgramTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoyaltyProgramTierWithDefaults

`func NewUpdateLoyaltyProgramTierWithDefaults() *UpdateLoyaltyProgramTier`

NewUpdateLoyaltyProgramTierWithDefaults instantiates a new UpdateLoyaltyProgramTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateLoyaltyProgramTier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoyaltyProgramTier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoyaltyProgramTier) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateLoyaltyProgramTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoyaltyProgramTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoyaltyProgramTier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoyaltyProgramTier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMinPoints

`func (o *UpdateLoyaltyProgramTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *UpdateLoyaltyProgramTier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *UpdateLoyaltyProgramTier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.

### HasMinPoints

`func (o *UpdateLoyaltyProgramTier) HasMinPoints() bool`

HasMinPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


