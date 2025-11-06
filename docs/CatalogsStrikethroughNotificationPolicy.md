# CatalogsStrikethroughNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**AheadOfDaysTrigger** | Pointer to **int64** | The number of days in advance that strikethrough pricing updates should be sent. | [optional] 
**BatchSize** | Pointer to **int64** | The required size of each batch of data. | [optional] [default to 1000]

## Methods

### NewCatalogsStrikethroughNotificationPolicy

`func NewCatalogsStrikethroughNotificationPolicy(name string, ) *CatalogsStrikethroughNotificationPolicy`

NewCatalogsStrikethroughNotificationPolicy instantiates a new CatalogsStrikethroughNotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogsStrikethroughNotificationPolicyWithDefaults

`func NewCatalogsStrikethroughNotificationPolicyWithDefaults() *CatalogsStrikethroughNotificationPolicy`

NewCatalogsStrikethroughNotificationPolicyWithDefaults instantiates a new CatalogsStrikethroughNotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogsStrikethroughNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogsStrikethroughNotificationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogsStrikethroughNotificationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetAheadOfDaysTrigger

`func (o *CatalogsStrikethroughNotificationPolicy) GetAheadOfDaysTrigger() int64`

GetAheadOfDaysTrigger returns the AheadOfDaysTrigger field if non-nil, zero value otherwise.

### GetAheadOfDaysTriggerOk

`func (o *CatalogsStrikethroughNotificationPolicy) GetAheadOfDaysTriggerOk() (*int64, bool)`

GetAheadOfDaysTriggerOk returns a tuple with the AheadOfDaysTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAheadOfDaysTrigger

`func (o *CatalogsStrikethroughNotificationPolicy) SetAheadOfDaysTrigger(v int64)`

SetAheadOfDaysTrigger sets AheadOfDaysTrigger field to given value.

### HasAheadOfDaysTrigger

`func (o *CatalogsStrikethroughNotificationPolicy) HasAheadOfDaysTrigger() bool`

HasAheadOfDaysTrigger returns a boolean if a field has been set.

### GetBatchSize

`func (o *CatalogsStrikethroughNotificationPolicy) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *CatalogsStrikethroughNotificationPolicy) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *CatalogsStrikethroughNotificationPolicy) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *CatalogsStrikethroughNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


