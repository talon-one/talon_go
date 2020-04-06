# LoyaltyMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Joined** | Pointer to [**time.Time**](time.Time.md) | The moment in which the loyalty program was joined. | [optional] 
**LoyaltyProgramId** | Pointer to **int32** | The ID of the loyalty program belonging to this entity. | 

## Methods

### GetJoined

`func (o *LoyaltyMembership) GetJoined() time.Time`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *LoyaltyMembership) GetJoinedOk() (time.Time, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJoined

`func (o *LoyaltyMembership) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### SetJoined

`func (o *LoyaltyMembership) SetJoined(v time.Time)`

SetJoined gets a reference to the given time.Time and assigns it to the Joined field.

### GetLoyaltyProgramId

`func (o *LoyaltyMembership) GetLoyaltyProgramId() int32`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LoyaltyMembership) GetLoyaltyProgramIdOk() (int32, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyProgramId

`func (o *LoyaltyMembership) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### SetLoyaltyProgramId

`func (o *LoyaltyMembership) SetLoyaltyProgramId(v int32)`

SetLoyaltyProgramId gets a reference to the given int32 and assigns it to the LoyaltyProgramId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


