# CampaignCreatedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of the notification | 
**TotalResultSize** | Pointer to **int64** | The total size of the result set. | 
**Data** | Pointer to [**[]CampaignCreatedNotificationItem**](CampaignCreatedNotificationItem.md) | A list of campaign notification data. | [optional] 

## Methods

### NewCampaignCreatedNotification

`func NewCampaignCreatedNotification(notificationType string, totalResultSize int64, ) *CampaignCreatedNotification`

NewCampaignCreatedNotification instantiates a new CampaignCreatedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreatedNotificationWithDefaults

`func NewCampaignCreatedNotificationWithDefaults() *CampaignCreatedNotification`

NewCampaignCreatedNotificationWithDefaults instantiates a new CampaignCreatedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignCreatedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignCreatedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignCreatedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignCreatedNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignCreatedNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignCreatedNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CampaignCreatedNotification) GetData() []CampaignCreatedNotificationItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignCreatedNotification) GetDataOk() (*[]CampaignCreatedNotificationItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignCreatedNotification) SetData(v []CampaignCreatedNotificationItem)`

SetData sets Data field to given value.

### HasData

`func (o *CampaignCreatedNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


