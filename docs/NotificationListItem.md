# NotificationListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | Pointer to **int64** | The ID of the notification. | 
**NotificationName** | Pointer to **string** | The name of the notification. | 
**EntityId** | Pointer to **int64** | The ID of the entity to which this notification belongs. For example, in case of a loyalty notification, this value is the ID of the loyalty program.  | 
**Enabled** | Pointer to **bool** | Indicates whether the notification is activated. | 

## Methods

### NewNotificationListItem

`func NewNotificationListItem(notificationId int64, notificationName string, entityId int64, enabled bool, ) *NotificationListItem`

NewNotificationListItem instantiates a new NotificationListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListItemWithDefaults

`func NewNotificationListItemWithDefaults() *NotificationListItem`

NewNotificationListItemWithDefaults instantiates a new NotificationListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationId

`func (o *NotificationListItem) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotificationListItem) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotificationListItem) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationName

`func (o *NotificationListItem) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *NotificationListItem) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *NotificationListItem) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.


### GetEntityId

`func (o *NotificationListItem) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *NotificationListItem) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *NotificationListItem) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.


### GetEnabled

`func (o *NotificationListItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationListItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationListItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


