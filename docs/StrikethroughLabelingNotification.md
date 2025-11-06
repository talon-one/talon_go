# StrikethroughLabelingNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the strikethrough pricing notification. | [optional] 
**ValidFrom** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the strikethrough pricing update becomes valid. Set for **scheduled** strikethrough pricing updates (version: v2) only.  | [optional] 
**ApplicationId** | Pointer to **int64** | The ID of the Application to which the catalog items labels belongs. | 
**CurrentBatch** | Pointer to **int64** | The batch number of the notification. Notifications might be sent in different batches. | 
**TotalBatches** | Pointer to **int64** | The total number of batches for the notification. | 
**Trigger** | Pointer to [**StrikethroughTrigger**](StrikethroughTrigger.md) |  | 
**ChangedItems** | Pointer to [**[]StrikethroughChangedItem**](StrikethroughChangedItem.md) |  | 
**NotificationType** | Pointer to **string** | The type of the notification | 

## Methods

### NewStrikethroughLabelingNotification

`func NewStrikethroughLabelingNotification(applicationId int64, currentBatch int64, totalBatches int64, trigger StrikethroughTrigger, changedItems []StrikethroughChangedItem, notificationType string, ) *StrikethroughLabelingNotification`

NewStrikethroughLabelingNotification instantiates a new StrikethroughLabelingNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughLabelingNotificationWithDefaults

`func NewStrikethroughLabelingNotificationWithDefaults() *StrikethroughLabelingNotification`

NewStrikethroughLabelingNotificationWithDefaults instantiates a new StrikethroughLabelingNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *StrikethroughLabelingNotification) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrikethroughLabelingNotification) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StrikethroughLabelingNotification) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StrikethroughLabelingNotification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetValidFrom

`func (o *StrikethroughLabelingNotification) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *StrikethroughLabelingNotification) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *StrikethroughLabelingNotification) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *StrikethroughLabelingNotification) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetApplicationId

`func (o *StrikethroughLabelingNotification) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *StrikethroughLabelingNotification) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *StrikethroughLabelingNotification) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCurrentBatch

`func (o *StrikethroughLabelingNotification) GetCurrentBatch() int64`

GetCurrentBatch returns the CurrentBatch field if non-nil, zero value otherwise.

### GetCurrentBatchOk

`func (o *StrikethroughLabelingNotification) GetCurrentBatchOk() (*int64, bool)`

GetCurrentBatchOk returns a tuple with the CurrentBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBatch

`func (o *StrikethroughLabelingNotification) SetCurrentBatch(v int64)`

SetCurrentBatch sets CurrentBatch field to given value.


### GetTotalBatches

`func (o *StrikethroughLabelingNotification) GetTotalBatches() int64`

GetTotalBatches returns the TotalBatches field if non-nil, zero value otherwise.

### GetTotalBatchesOk

`func (o *StrikethroughLabelingNotification) GetTotalBatchesOk() (*int64, bool)`

GetTotalBatchesOk returns a tuple with the TotalBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBatches

`func (o *StrikethroughLabelingNotification) SetTotalBatches(v int64)`

SetTotalBatches sets TotalBatches field to given value.


### GetTrigger

`func (o *StrikethroughLabelingNotification) GetTrigger() StrikethroughTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *StrikethroughLabelingNotification) GetTriggerOk() (*StrikethroughTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *StrikethroughLabelingNotification) SetTrigger(v StrikethroughTrigger)`

SetTrigger sets Trigger field to given value.


### GetChangedItems

`func (o *StrikethroughLabelingNotification) GetChangedItems() []StrikethroughChangedItem`

GetChangedItems returns the ChangedItems field if non-nil, zero value otherwise.

### GetChangedItemsOk

`func (o *StrikethroughLabelingNotification) GetChangedItemsOk() (*[]StrikethroughChangedItem, bool)`

GetChangedItemsOk returns a tuple with the ChangedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedItems

`func (o *StrikethroughLabelingNotification) SetChangedItems(v []StrikethroughChangedItem)`

SetChangedItems sets ChangedItems field to given value.


### GetNotificationType

`func (o *StrikethroughLabelingNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *StrikethroughLabelingNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *StrikethroughLabelingNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


