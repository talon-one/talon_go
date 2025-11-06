# LoyaltyTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ProgramID** | Pointer to **int64** | The ID of the loyalty program that owns this entity. | 
**ProgramName** | Pointer to **string** | The integration name of the loyalty program that owns this entity. | [optional] 
**ProgramTitle** | Pointer to **string** | The Campaign Manager-displayed name of the loyalty program that owns this entity. | [optional] 
**Name** | Pointer to **string** | The name of the tier. | 
**MinPoints** | Pointer to **float32** | The minimum amount of points required to enter the tier. | 

## Methods

### NewLoyaltyTier

`func NewLoyaltyTier(id int64, created time.Time, programID int64, name string, minPoints float32, ) *LoyaltyTier`

NewLoyaltyTier instantiates a new LoyaltyTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyTierWithDefaults

`func NewLoyaltyTierWithDefaults() *LoyaltyTier`

NewLoyaltyTierWithDefaults instantiates a new LoyaltyTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyTier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyTier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyTier) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *LoyaltyTier) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyTier) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyTier) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramID

`func (o *LoyaltyTier) GetProgramID() int64`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyTier) GetProgramIDOk() (*int64, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramID

`func (o *LoyaltyTier) SetProgramID(v int64)`

SetProgramID sets ProgramID field to given value.


### GetProgramName

`func (o *LoyaltyTier) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *LoyaltyTier) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *LoyaltyTier) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.

### HasProgramName

`func (o *LoyaltyTier) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### GetProgramTitle

`func (o *LoyaltyTier) GetProgramTitle() string`

GetProgramTitle returns the ProgramTitle field if non-nil, zero value otherwise.

### GetProgramTitleOk

`func (o *LoyaltyTier) GetProgramTitleOk() (*string, bool)`

GetProgramTitleOk returns a tuple with the ProgramTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramTitle

`func (o *LoyaltyTier) SetProgramTitle(v string)`

SetProgramTitle sets ProgramTitle field to given value.

### HasProgramTitle

`func (o *LoyaltyTier) HasProgramTitle() bool`

HasProgramTitle returns a boolean if a field has been set.

### GetName

`func (o *LoyaltyTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyTier) SetName(v string)`

SetName sets Name field to given value.


### GetMinPoints

`func (o *LoyaltyTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *LoyaltyTier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *LoyaltyTier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


