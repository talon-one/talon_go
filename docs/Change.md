# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**UserId** | Pointer to **int64** | The ID of the user associated with this entity. | 
**ApplicationId** | Pointer to **int64** | ID of application associated with change. | [optional] 
**Entity** | Pointer to **string** | API endpoint on which the change was initiated. | 
**Old** | Pointer to [**map[string]interface{}**](.md) | Resource before the change occurred. | [optional] 
**New** | Pointer to [**map[string]interface{}**](.md) | Resource after the change occurred. | [optional] 
**ManagementKeyId** | Pointer to **int64** | ID of management key used to perform changes. | [optional] 

## Methods

### NewChange

`func NewChange(id int64, created time.Time, userId int64, entity string, ) *Change`

NewChange instantiates a new Change object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWithDefaults

`func NewChangeWithDefaults() *Change`

NewChangeWithDefaults instantiates a new Change object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Change) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Change) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Change) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Change) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Change) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Change) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUserId

`func (o *Change) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Change) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Change) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetApplicationId

`func (o *Change) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Change) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Change) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Change) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEntity

`func (o *Change) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Change) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Change) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetOld

`func (o *Change) GetOld() map[string]interface{}`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *Change) GetOldOk() (*map[string]interface{}, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *Change) SetOld(v map[string]interface{})`

SetOld sets Old field to given value.

### HasOld

`func (o *Change) HasOld() bool`

HasOld returns a boolean if a field has been set.

### GetNew

`func (o *Change) GetNew() map[string]interface{}`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Change) GetNewOk() (*map[string]interface{}, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *Change) SetNew(v map[string]interface{})`

SetNew sets New field to given value.

### HasNew

`func (o *Change) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetManagementKeyId

`func (o *Change) GetManagementKeyId() int64`

GetManagementKeyId returns the ManagementKeyId field if non-nil, zero value otherwise.

### GetManagementKeyIdOk

`func (o *Change) GetManagementKeyIdOk() (*int64, bool)`

GetManagementKeyIdOk returns a tuple with the ManagementKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementKeyId

`func (o *Change) SetManagementKeyId(v int64)`

SetManagementKeyId sets ManagementKeyId field to given value.

### HasManagementKeyId

`func (o *Change) HasManagementKeyId() bool`

HasManagementKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


