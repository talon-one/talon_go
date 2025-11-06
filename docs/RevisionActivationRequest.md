# RevisionActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]int64** | The list of IDs of the users who will receive the activation request. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) | Time when the revisions are finalized after the &#x60;activate_revision&#x60; operation. The current time is used when left blank.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 

## Methods

### NewRevisionActivationRequest

`func NewRevisionActivationRequest(userIds []int64, ) *RevisionActivationRequest`

NewRevisionActivationRequest instantiates a new RevisionActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionActivationRequestWithDefaults

`func NewRevisionActivationRequestWithDefaults() *RevisionActivationRequest`

NewRevisionActivationRequestWithDefaults instantiates a new RevisionActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *RevisionActivationRequest) GetUserIds() []int64`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *RevisionActivationRequest) GetUserIdsOk() (*[]int64, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *RevisionActivationRequest) SetUserIds(v []int64)`

SetUserIds sets UserIds field to given value.


### GetActivateAt

`func (o *RevisionActivationRequest) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *RevisionActivationRequest) GetActivateAtOk() (*time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateAt

`func (o *RevisionActivationRequest) SetActivateAt(v time.Time)`

SetActivateAt sets ActivateAt field to given value.

### HasActivateAt

`func (o *RevisionActivationRequest) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


