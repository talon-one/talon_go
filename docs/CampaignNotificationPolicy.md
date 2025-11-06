# CampaignNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int64** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] [default to 5]

## Methods

### NewCampaignNotificationPolicy

`func NewCampaignNotificationPolicy(name string, ) *CampaignNotificationPolicy`

NewCampaignNotificationPolicy instantiates a new CampaignNotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignNotificationPolicyWithDefaults

`func NewCampaignNotificationPolicyWithDefaults() *CampaignNotificationPolicy`

NewCampaignNotificationPolicyWithDefaults instantiates a new CampaignNotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignNotificationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignNotificationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetBatchingEnabled

`func (o *CampaignNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *CampaignNotificationPolicy) GetBatchingEnabledOk() (*bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchingEnabled

`func (o *CampaignNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled sets BatchingEnabled field to given value.

### HasBatchingEnabled

`func (o *CampaignNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### GetBatchSize

`func (o *CampaignNotificationPolicy) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *CampaignNotificationPolicy) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *CampaignNotificationPolicy) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *CampaignNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


