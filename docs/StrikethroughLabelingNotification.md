# StrikethroughLabelingNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the strikethrough pricing notification. | [optional] 
**ValidFrom** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the strikethrough pricing update becomes valid. Set for **scheduled** strikethrough pricing updates (version: v2) only.  | [optional] 
**ApplicationId** | Pointer to **int32** | The ID of the Application to which the catalog items labels belongs. | 
**CurrentBatch** | Pointer to **int32** | The batch number of the notification. Notifications might be sent in different batches. | 
**TotalBatches** | Pointer to **int32** | The total number of batches for the notification. | 
**Trigger** | Pointer to [**StrikethroughTrigger**](StrikethroughTrigger.md) |  | 
**ChangedItems** | Pointer to [**[]StrikethroughChangedItem**](StrikethroughChangedItem.md) |  | 

## Methods

### GetVersion

`func (o *StrikethroughLabelingNotification) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrikethroughLabelingNotification) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *StrikethroughLabelingNotification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *StrikethroughLabelingNotification) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### GetValidFrom

`func (o *StrikethroughLabelingNotification) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *StrikethroughLabelingNotification) GetValidFromOk() (time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidFrom

`func (o *StrikethroughLabelingNotification) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFrom

`func (o *StrikethroughLabelingNotification) SetValidFrom(v time.Time)`

SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.

### GetApplicationId

`func (o *StrikethroughLabelingNotification) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *StrikethroughLabelingNotification) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *StrikethroughLabelingNotification) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *StrikethroughLabelingNotification) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetCurrentBatch

`func (o *StrikethroughLabelingNotification) GetCurrentBatch() int32`

GetCurrentBatch returns the CurrentBatch field if non-nil, zero value otherwise.

### GetCurrentBatchOk

`func (o *StrikethroughLabelingNotification) GetCurrentBatchOk() (int32, bool)`

GetCurrentBatchOk returns a tuple with the CurrentBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentBatch

`func (o *StrikethroughLabelingNotification) HasCurrentBatch() bool`

HasCurrentBatch returns a boolean if a field has been set.

### SetCurrentBatch

`func (o *StrikethroughLabelingNotification) SetCurrentBatch(v int32)`

SetCurrentBatch gets a reference to the given int32 and assigns it to the CurrentBatch field.

### GetTotalBatches

`func (o *StrikethroughLabelingNotification) GetTotalBatches() int32`

GetTotalBatches returns the TotalBatches field if non-nil, zero value otherwise.

### GetTotalBatchesOk

`func (o *StrikethroughLabelingNotification) GetTotalBatchesOk() (int32, bool)`

GetTotalBatchesOk returns a tuple with the TotalBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalBatches

`func (o *StrikethroughLabelingNotification) HasTotalBatches() bool`

HasTotalBatches returns a boolean if a field has been set.

### SetTotalBatches

`func (o *StrikethroughLabelingNotification) SetTotalBatches(v int32)`

SetTotalBatches gets a reference to the given int32 and assigns it to the TotalBatches field.

### GetTrigger

`func (o *StrikethroughLabelingNotification) GetTrigger() StrikethroughTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *StrikethroughLabelingNotification) GetTriggerOk() (StrikethroughTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrigger

`func (o *StrikethroughLabelingNotification) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### SetTrigger

`func (o *StrikethroughLabelingNotification) SetTrigger(v StrikethroughTrigger)`

SetTrigger gets a reference to the given StrikethroughTrigger and assigns it to the Trigger field.

### GetChangedItems

`func (o *StrikethroughLabelingNotification) GetChangedItems() []StrikethroughChangedItem`

GetChangedItems returns the ChangedItems field if non-nil, zero value otherwise.

### GetChangedItemsOk

`func (o *StrikethroughLabelingNotification) GetChangedItemsOk() ([]StrikethroughChangedItem, bool)`

GetChangedItemsOk returns a tuple with the ChangedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangedItems

`func (o *StrikethroughLabelingNotification) HasChangedItems() bool`

HasChangedItems returns a boolean if a field has been set.

### SetChangedItems

`func (o *StrikethroughLabelingNotification) SetChangedItems(v []StrikethroughChangedItem)`

SetChangedItems gets a reference to the given []StrikethroughChangedItem and assigns it to the ChangedItems field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


