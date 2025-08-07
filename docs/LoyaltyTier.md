# LoyaltyTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ProgramID** | Pointer to **int32** | The ID of the loyalty program that owns this entity. | 
**ProgramName** | Pointer to **string** | The integration name of the loyalty program that owns this entity. | [optional] 
**ProgramTitle** | Pointer to **string** | The Campaign Manager-displayed name of the loyalty program that owns this entity. | [optional] 
**Name** | Pointer to **string** | The name of the tier. | 
**MinPoints** | Pointer to **float32** | The minimum amount of points required to enter the tier. | 

## Methods

### GetId

`func (o *LoyaltyTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyTier) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyTier) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyTier) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *LoyaltyTier) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyTier) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LoyaltyTier) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LoyaltyTier) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramID

`func (o *LoyaltyTier) GetProgramID() int32`

GetProgramID returns the ProgramID field if non-nil, zero value otherwise.

### GetProgramIDOk

`func (o *LoyaltyTier) GetProgramIDOk() (int32, bool)`

GetProgramIDOk returns a tuple with the ProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramID

`func (o *LoyaltyTier) HasProgramID() bool`

HasProgramID returns a boolean if a field has been set.

### SetProgramID

`func (o *LoyaltyTier) SetProgramID(v int32)`

SetProgramID gets a reference to the given int32 and assigns it to the ProgramID field.

### GetProgramName

`func (o *LoyaltyTier) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *LoyaltyTier) GetProgramNameOk() (string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramName

`func (o *LoyaltyTier) HasProgramName() bool`

HasProgramName returns a boolean if a field has been set.

### SetProgramName

`func (o *LoyaltyTier) SetProgramName(v string)`

SetProgramName gets a reference to the given string and assigns it to the ProgramName field.

### GetProgramTitle

`func (o *LoyaltyTier) GetProgramTitle() string`

GetProgramTitle returns the ProgramTitle field if non-nil, zero value otherwise.

### GetProgramTitleOk

`func (o *LoyaltyTier) GetProgramTitleOk() (string, bool)`

GetProgramTitleOk returns a tuple with the ProgramTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramTitle

`func (o *LoyaltyTier) HasProgramTitle() bool`

HasProgramTitle returns a boolean if a field has been set.

### SetProgramTitle

`func (o *LoyaltyTier) SetProgramTitle(v string)`

SetProgramTitle gets a reference to the given string and assigns it to the ProgramTitle field.

### GetName

`func (o *LoyaltyTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyTier) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyTier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyTier) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetMinPoints

`func (o *LoyaltyTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *LoyaltyTier) GetMinPointsOk() (float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinPoints

`func (o *LoyaltyTier) HasMinPoints() bool`

HasMinPoints returns a boolean if a field has been set.

### SetMinPoints

`func (o *LoyaltyTier) SetMinPoints(v float32)`

SetMinPoints gets a reference to the given float32 and assigns it to the MinPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


