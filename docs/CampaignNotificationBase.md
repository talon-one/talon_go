# CampaignNotificationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of the notification | 
**TotalResultSize** | Pointer to **int64** | The total size of the result set. | 

## Methods

### NewCampaignNotificationBase

`func NewCampaignNotificationBase(notificationType string, totalResultSize int64, ) *CampaignNotificationBase`

NewCampaignNotificationBase instantiates a new CampaignNotificationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignNotificationBaseWithDefaults

`func NewCampaignNotificationBaseWithDefaults() *CampaignNotificationBase`

NewCampaignNotificationBaseWithDefaults instantiates a new CampaignNotificationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignNotificationBase) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignNotificationBase) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignNotificationBase) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignNotificationBase) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignNotificationBase) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignNotificationBase) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


