# CampaignCollectionEditedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Collection** | Pointer to [**CollectionWithoutPayload**](CollectionWithoutPayload.md) |  | 

## Methods

### NewCampaignCollectionEditedNotificationItem

`func NewCampaignCollectionEditedNotificationItem(event string, campaign Campaign, collection CollectionWithoutPayload, ) *CampaignCollectionEditedNotificationItem`

NewCampaignCollectionEditedNotificationItem instantiates a new CampaignCollectionEditedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCollectionEditedNotificationItemWithDefaults

`func NewCampaignCollectionEditedNotificationItemWithDefaults() *CampaignCollectionEditedNotificationItem`

NewCampaignCollectionEditedNotificationItemWithDefaults instantiates a new CampaignCollectionEditedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignCollectionEditedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignCollectionEditedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignCollectionEditedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignCollectionEditedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignCollectionEditedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignCollectionEditedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetRuleset

`func (o *CampaignCollectionEditedNotificationItem) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignCollectionEditedNotificationItem) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignCollectionEditedNotificationItem) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignCollectionEditedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetCollection

`func (o *CampaignCollectionEditedNotificationItem) GetCollection() CollectionWithoutPayload`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CampaignCollectionEditedNotificationItem) GetCollectionOk() (*CollectionWithoutPayload, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CampaignCollectionEditedNotificationItem) SetCollection(v CollectionWithoutPayload)`

SetCollection sets Collection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


