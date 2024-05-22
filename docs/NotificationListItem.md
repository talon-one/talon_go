# NotificationListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | Pointer to **int32** | The ID of the notification. | 
**NotificationName** | Pointer to **string** | The name of the notification. | 
**EntityId** | Pointer to **int32** | The ID of the entity to which this notification belongs. For example, in case of a loyalty notification, this value is the ID of the loyalty program.  | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | 

## Methods

### GetNotificationId

`func (o *NotificationListItem) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotificationListItem) GetNotificationIdOk() (int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationId

`func (o *NotificationListItem) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### SetNotificationId

`func (o *NotificationListItem) SetNotificationId(v int32)`

SetNotificationId gets a reference to the given int32 and assigns it to the NotificationId field.

### GetNotificationName

`func (o *NotificationListItem) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *NotificationListItem) GetNotificationNameOk() (string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationName

`func (o *NotificationListItem) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### SetNotificationName

`func (o *NotificationListItem) SetNotificationName(v string)`

SetNotificationName gets a reference to the given string and assigns it to the NotificationName field.

### GetEntityId

`func (o *NotificationListItem) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *NotificationListItem) GetEntityIdOk() (int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntityId

`func (o *NotificationListItem) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityId

`func (o *NotificationListItem) SetEntityId(v int32)`

SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.

### GetEnabled

`func (o *NotificationListItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationListItem) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *NotificationListItem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *NotificationListItem) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


