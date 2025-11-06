# CampaignNotificationItemBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 

## Methods

### NewCampaignNotificationItemBase

`func NewCampaignNotificationItemBase(event string, ) *CampaignNotificationItemBase`

NewCampaignNotificationItemBase instantiates a new CampaignNotificationItemBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignNotificationItemBaseWithDefaults

`func NewCampaignNotificationItemBaseWithDefaults() *CampaignNotificationItemBase`

NewCampaignNotificationItemBaseWithDefaults instantiates a new CampaignNotificationItemBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignNotificationItemBase) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignNotificationItemBase) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignNotificationItemBase) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


