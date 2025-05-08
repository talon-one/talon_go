# RevisionActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]int32** | The list of IDs of the users who will receive the activation request. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) | Time when the revisions are finalized after the &#x60;activate_revision&#x60; operation. The current time is used when left blank.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 

## Methods

### GetUserIds

`func (o *RevisionActivationRequest) GetUserIds() []int32`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *RevisionActivationRequest) GetUserIdsOk() ([]int32, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserIds

`func (o *RevisionActivationRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIds

`func (o *RevisionActivationRequest) SetUserIds(v []int32)`

SetUserIds gets a reference to the given []int32 and assigns it to the UserIds field.

### GetActivateAt

`func (o *RevisionActivationRequest) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *RevisionActivationRequest) GetActivateAtOk() (time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivateAt

`func (o *RevisionActivationRequest) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.

### SetActivateAt

`func (o *RevisionActivationRequest) SetActivateAt(v time.Time)`

SetActivateAt gets a reference to the given time.Time and assigns it to the ActivateAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


