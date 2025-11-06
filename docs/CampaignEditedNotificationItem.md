# CampaignEditedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**OldCampaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 

## Methods

### NewCampaignEditedNotificationItem

`func NewCampaignEditedNotificationItem(event string, campaign Campaign, oldCampaign Campaign, ) *CampaignEditedNotificationItem`

NewCampaignEditedNotificationItem instantiates a new CampaignEditedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEditedNotificationItemWithDefaults

`func NewCampaignEditedNotificationItemWithDefaults() *CampaignEditedNotificationItem`

NewCampaignEditedNotificationItemWithDefaults instantiates a new CampaignEditedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignEditedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignEditedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignEditedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignEditedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignEditedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignEditedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetOldCampaign

`func (o *CampaignEditedNotificationItem) GetOldCampaign() Campaign`

GetOldCampaign returns the OldCampaign field if non-nil, zero value otherwise.

### GetOldCampaignOk

`func (o *CampaignEditedNotificationItem) GetOldCampaignOk() (*Campaign, bool)`

GetOldCampaignOk returns a tuple with the OldCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCampaign

`func (o *CampaignEditedNotificationItem) SetOldCampaign(v Campaign)`

SetOldCampaign sets OldCampaign field to given value.


### GetRuleset

`func (o *CampaignEditedNotificationItem) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignEditedNotificationItem) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignEditedNotificationItem) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignEditedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


