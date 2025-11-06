# CampaignDeletedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**DeletedAt** | Pointer to [**time.Time**](time.Time.md) | Time when the campaign was deleted. | 

## Methods

### NewCampaignDeletedNotificationItem

`func NewCampaignDeletedNotificationItem(event string, campaign Campaign, deletedAt time.Time, ) *CampaignDeletedNotificationItem`

NewCampaignDeletedNotificationItem instantiates a new CampaignDeletedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDeletedNotificationItemWithDefaults

`func NewCampaignDeletedNotificationItemWithDefaults() *CampaignDeletedNotificationItem`

NewCampaignDeletedNotificationItemWithDefaults instantiates a new CampaignDeletedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignDeletedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignDeletedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignDeletedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignDeletedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignDeletedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignDeletedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetDeletedAt

`func (o *CampaignDeletedNotificationItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CampaignDeletedNotificationItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CampaignDeletedNotificationItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


