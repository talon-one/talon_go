# ChangeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | ID of application associated with change. | [optional] 
**Entity** | Pointer to **string** | API endpoint on which the change was initiated. | 
**Old** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Resource before the change occurred. | [optional] 
**New** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Resource after the change occurred. | [optional] 
**ManagementKeyId** | Pointer to **int32** | ID of management key used to perform changes. | [optional] 

## Methods

### GetApplicationId

`func (o *ChangeAllOf) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ChangeAllOf) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ChangeAllOf) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ChangeAllOf) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetEntity

`func (o *ChangeAllOf) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ChangeAllOf) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *ChangeAllOf) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *ChangeAllOf) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetOld

`func (o *ChangeAllOf) GetOld() map[string]map[string]interface{}`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *ChangeAllOf) GetOldOk() (map[string]map[string]interface{}, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOld

`func (o *ChangeAllOf) HasOld() bool`

HasOld returns a boolean if a field has been set.

### SetOld

`func (o *ChangeAllOf) SetOld(v map[string]map[string]interface{})`

SetOld gets a reference to the given map[string]map[string]interface{} and assigns it to the Old field.

### GetNew

`func (o *ChangeAllOf) GetNew() map[string]map[string]interface{}`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *ChangeAllOf) GetNewOk() (map[string]map[string]interface{}, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNew

`func (o *ChangeAllOf) HasNew() bool`

HasNew returns a boolean if a field has been set.

### SetNew

`func (o *ChangeAllOf) SetNew(v map[string]map[string]interface{})`

SetNew gets a reference to the given map[string]map[string]interface{} and assigns it to the New field.

### GetManagementKeyId

`func (o *ChangeAllOf) GetManagementKeyId() int32`

GetManagementKeyId returns the ManagementKeyId field if non-nil, zero value otherwise.

### GetManagementKeyIdOk

`func (o *ChangeAllOf) GetManagementKeyIdOk() (int32, bool)`

GetManagementKeyIdOk returns a tuple with the ManagementKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementKeyId

`func (o *ChangeAllOf) HasManagementKeyId() bool`

HasManagementKeyId returns a boolean if a field has been set.

### SetManagementKeyId

`func (o *ChangeAllOf) SetManagementKeyId(v int32)`

SetManagementKeyId gets a reference to the given int32 and assigns it to the ManagementKeyId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


