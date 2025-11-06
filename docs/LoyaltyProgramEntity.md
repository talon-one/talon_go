# LoyaltyProgramEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramID** | Pointer to **int64** | The ID of the loyalty program that owns this entity. | 
**ProgramName** | Pointer to **string** | The integration name of the loyalty program that owns this entity. | [optional] 
**ProgramTitle** | Pointer to **string** | The Campaign Manager-displayed name of the loyalty program that owns this entity. | [optional] 

## Methods

### NewLoyaltyProgramEntity

`func NewLoyaltyProgramEntity(programID int64, ) *LoyaltyProgramEntity`

NewLoyaltyProgramEntity instantiates a new LoyaltyProgramEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramEntityWithDefaults

`func NewLoyaltyProgramEntityWithDefaults() *LoyaltyProgramEntity`

NewLoyaltyProgramEntityWithDefaults instantiates a new LoyaltyProgramEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramID

`func (o *LoyaltyProgramEntity) GetProgramID() int64`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyProgramEntity) GetProgramIDOk() (*int64, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyProgramEntity) SetProgramID(v int64)`

SetProgramID sets ProgramID field to given value.


### GetProgramName

`func (o *LoyaltyProgramEntity) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *LoyaltyProgramEntity) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *LoyaltyProgramEntity) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.

### HasProgramName

`func (o *LoyaltyProgramEntity) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### GetProgramTitle

`func (o *LoyaltyProgramEntity) GetProgramTitle() string`

GetProgramTitle returns the ProgramTitle field if non-nil, zero value otherwise.

### GetProgramTitleOk

`func (o *LoyaltyProgramEntity) GetProgramTitleOk() (*string, bool)`

GetProgramTitleOk returns a tuple with the ProgramTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramTitle

`func (o *LoyaltyProgramEntity) SetProgramTitle(v string)`

SetProgramTitle sets ProgramTitle field to given value.

### HasProgramTitle

`func (o *LoyaltyProgramEntity) HasProgramTitle() bool`

HasProgramTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


