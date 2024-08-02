# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | ID of application associated with change. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Entity** | Pointer to **string** | API endpoint on which the change was initiated. | 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**ManagementKeyId** | Pointer to **int32** | ID of management key used to perform changes. | [optional] 
**New** | Pointer to [**map[string]interface{}**](.md) | Resource after the change occurred. | [optional] 
**Old** | Pointer to [**map[string]interface{}**](.md) | Resource before the change occurred. | [optional] 
**UserId** | Pointer to **int32** | The ID of the user associated with this entity. | 

## Methods

### GetApplicationId

`func (o *Change) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Change) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Change) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Change) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCreated

`func (o *Change) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Change) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Change) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Change) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetEntity

`func (o *Change) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Change) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *Change) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *Change) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetId

`func (o *Change) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Change) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Change) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Change) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetManagementKeyId

`func (o *Change) GetManagementKeyId() int32`

GetManagementKeyId returns the ManagementKeyId field if non-nil, zero value otherwise.

### GetManagementKeyIdOk

`func (o *Change) GetManagementKeyIdOk() (int32, bool)`

GetManagementKeyIdOk returns a tuple with the ManagementKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementKeyId

`func (o *Change) HasManagementKeyId() bool`

HasManagementKeyId returns a boolean if a field has been set.

### SetManagementKeyId

`func (o *Change) SetManagementKeyId(v int32)`

SetManagementKeyId gets a reference to the given int32 and assigns it to the ManagementKeyId field.

### GetNew

`func (o *Change) GetNew() map[string]interface{}`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Change) GetNewOk() (map[string]interface{}, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNew

`func (o *Change) HasNew() bool`

HasNew returns a boolean if a field has been set.

### SetNew

`func (o *Change) SetNew(v map[string]interface{})`

SetNew gets a reference to the given map[string]interface{} and assigns it to the New field.

### GetOld

`func (o *Change) GetOld() map[string]interface{}`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *Change) GetOldOk() (map[string]interface{}, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOld

`func (o *Change) HasOld() bool`

HasOld returns a boolean if a field has been set.

### SetOld

`func (o *Change) SetOld(v map[string]interface{})`

SetOld gets a reference to the given map[string]interface{} and assigns it to the Old field.

### GetUserId

`func (o *Change) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Change) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *Change) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *Change) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


