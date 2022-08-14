# UserFeedNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdate** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the last request for this list. | 
**Notifications** | Pointer to [**[]FeedNotification**](FeedNotification.md) | List of all notifications to notify the user about. | 

## Methods

### GetLastUpdate

`func (o *UserFeedNotifications) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserFeedNotifications) GetLastUpdateOk() (time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdate

`func (o *UserFeedNotifications) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdate

`func (o *UserFeedNotifications) SetLastUpdate(v time.Time)`

SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.

### GetNotifications

`func (o *UserFeedNotifications) GetNotifications() []FeedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UserFeedNotifications) GetNotificationsOk() ([]FeedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotifications

`func (o *UserFeedNotifications) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### SetNotifications

`func (o *UserFeedNotifications) SetNotifications(v []FeedNotification)`

SetNotifications gets a reference to the given []FeedNotification and assigns it to the Notifications field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


