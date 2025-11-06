# CampaignRulesetChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of the notification | 
**TotalResultSize** | Pointer to **int64** | The total size of the result set. | 
**Data** | Pointer to [**[]CampaignRulesetChangedNotificationItem**](CampaignRulesetChangedNotificationItem.md) | A list of campaign notification data. | [optional] 

## Methods

### NewCampaignRulesetChangedNotification

`func NewCampaignRulesetChangedNotification(notificationType string, totalResultSize int64, ) *CampaignRulesetChangedNotification`

NewCampaignRulesetChangedNotification instantiates a new CampaignRulesetChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRulesetChangedNotificationWithDefaults

`func NewCampaignRulesetChangedNotificationWithDefaults() *CampaignRulesetChangedNotification`

NewCampaignRulesetChangedNotificationWithDefaults instantiates a new CampaignRulesetChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignRulesetChangedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignRulesetChangedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignRulesetChangedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignRulesetChangedNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignRulesetChangedNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignRulesetChangedNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CampaignRulesetChangedNotification) GetData() []CampaignRulesetChangedNotificationItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignRulesetChangedNotification) GetDataOk() (*[]CampaignRulesetChangedNotificationItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignRulesetChangedNotification) SetData(v []CampaignRulesetChangedNotificationItem)`

SetData sets Data field to given value.

### HasData

`func (o *CampaignRulesetChangedNotification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


