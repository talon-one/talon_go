# CampaignEditedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of the notification | 
**TotalResultSize** | Pointer to **int64** | The total size of the result set. | 
**Data** | Pointer to [**[]CampaignEditedNotificationItem**](CampaignEditedNotificationItem.md) | A list of campaign notification data. | [optional] 

## Methods

### NewCampaignEditedNotification

`func NewCampaignEditedNotification(notificationType string, totalResultSize int64, ) *CampaignEditedNotification`

NewCampaignEditedNotification instantiates a new CampaignEditedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEditedNotificationWithDefaults

`func NewCampaignEditedNotificationWithDefaults() *CampaignEditedNotification`

NewCampaignEditedNotificationWithDefaults instantiates a new CampaignEditedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignEditedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignEditedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignEditedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignEditedNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignEditedNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignEditedNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CampaignEditedNotification) GetData() []CampaignEditedNotificationItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignEditedNotification) GetDataOk() (*[]CampaignEditedNotificationItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignEditedNotification) SetData(v []CampaignEditedNotificationItem)`

SetData sets Data field to given value.

### HasData

`func (o *CampaignEditedNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


