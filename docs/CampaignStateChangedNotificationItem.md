# CampaignStateChangedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**OldState** | Pointer to **string** | The campaign&#39;s old state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;archived&#39;]  | 
**NewState** | Pointer to **string** | The campaign&#39;s new state. Can be one of the following: [&#39;running&#39;, &#39;disabled&#39;, &#39;scheduled&#39;, &#39;expired&#39;, &#39;archived&#39;]  | 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Placeholders** | Pointer to [**[]PlaceholderDetails**](PlaceholderDetails.md) | The current details of the [placeholders](https://docs.talon.one/docs/product/campaigns/templates/create-templates#use-placeholders) in the campaign. | [optional] 

## Methods

### NewCampaignStateChangedNotificationItem

`func NewCampaignStateChangedNotificationItem(event string, campaign Campaign, oldState string, newState string, ) *CampaignStateChangedNotificationItem`

NewCampaignStateChangedNotificationItem instantiates a new CampaignStateChangedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStateChangedNotificationItemWithDefaults

`func NewCampaignStateChangedNotificationItemWithDefaults() *CampaignStateChangedNotificationItem`

NewCampaignStateChangedNotificationItemWithDefaults instantiates a new CampaignStateChangedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignStateChangedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignStateChangedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignStateChangedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignStateChangedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignStateChangedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignStateChangedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetOldState

`func (o *CampaignStateChangedNotificationItem) GetOldState() string`

GetOldState returns the OldState field if non-nil, zero value otherwise.

### GetOldStateOk

`func (o *CampaignStateChangedNotificationItem) GetOldStateOk() (*string, bool)`

GetOldStateOk returns a tuple with the OldState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldState

`func (o *CampaignStateChangedNotificationItem) SetOldState(v string)`

SetOldState sets OldState field to given value.


### GetNewState

`func (o *CampaignStateChangedNotificationItem) GetNewState() string`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *CampaignStateChangedNotificationItem) GetNewStateOk() (*string, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *CampaignStateChangedNotificationItem) SetNewState(v string)`

SetNewState sets NewState field to given value.


### GetRuleset

`func (o *CampaignStateChangedNotificationItem) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignStateChangedNotificationItem) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignStateChangedNotificationItem) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignStateChangedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetPlaceholders

`func (o *CampaignStateChangedNotificationItem) GetPlaceholders() []PlaceholderDetails`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *CampaignStateChangedNotificationItem) GetPlaceholdersOk() (*[]PlaceholderDetails, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *CampaignStateChangedNotificationItem) SetPlaceholders(v []PlaceholderDetails)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *CampaignStateChangedNotificationItem) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


