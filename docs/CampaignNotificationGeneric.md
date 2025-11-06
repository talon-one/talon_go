# CampaignNotificationGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of the notification | 
**TotalResultSize** | Pointer to **int64** | The total size of the result set. | 
**Data** | Pointer to **[]map[string]interface{}** | A list of campaign notification data. | 

## Methods

### NewCampaignNotificationGeneric

`func NewCampaignNotificationGeneric(notificationType string, totalResultSize int64, data []map[string]interface{}, ) *CampaignNotificationGeneric`

NewCampaignNotificationGeneric instantiates a new CampaignNotificationGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignNotificationGenericWithDefaults

`func NewCampaignNotificationGenericWithDefaults() *CampaignNotificationGeneric`

NewCampaignNotificationGenericWithDefaults instantiates a new CampaignNotificationGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignNotificationGeneric) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignNotificationGeneric) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignNotificationGeneric) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignNotificationGeneric) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignNotificationGeneric) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignNotificationGeneric) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CampaignNotificationGeneric) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignNotificationGeneric) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignNotificationGeneric) SetData(v []map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


