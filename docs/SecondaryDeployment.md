# SecondaryDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Name** | Pointer to **string** | The name of the deployment. Used as subdomain, e.g. experimental.your-company.europe-west1.talon.one | 
**UserId** | Pointer to **int64** | The ID of the user who created the deployment. | 
**Status** | Pointer to **string** | The status of the deployment. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the deployment was created. | 
**ActiveAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the deployment became active. | [optional] 
**FailedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the deployment failed. | [optional] 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the deployment was deleted. | [optional] 

## Methods

### NewSecondaryDeployment

`func NewSecondaryDeployment(id int64, name string, userId int64, status string, createdAt time.Time, ) *SecondaryDeployment`

NewSecondaryDeployment instantiates a new SecondaryDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryDeploymentWithDefaults

`func NewSecondaryDeploymentWithDefaults() *SecondaryDeployment`

NewSecondaryDeploymentWithDefaults instantiates a new SecondaryDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecondaryDeployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecondaryDeployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecondaryDeployment) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *SecondaryDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecondaryDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecondaryDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *SecondaryDeployment) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SecondaryDeployment) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SecondaryDeployment) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *SecondaryDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecondaryDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecondaryDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *SecondaryDeployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecondaryDeployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecondaryDeployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActiveAt

`func (o *SecondaryDeployment) GetActiveAt() time.Time`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *SecondaryDeployment) GetActiveAtOk() (*time.Time, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *SecondaryDeployment) SetActiveAt(v time.Time)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *SecondaryDeployment) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### GetFailedAt

`func (o *SecondaryDeployment) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *SecondaryDeployment) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *SecondaryDeployment) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *SecondaryDeployment) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *SecondaryDeployment) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SecondaryDeployment) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SecondaryDeployment) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SecondaryDeployment) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


