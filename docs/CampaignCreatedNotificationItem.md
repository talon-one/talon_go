# CampaignCreatedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | 
**Ruleset** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 
**Placeholders** | Pointer to [**[]PlaceholderDetails**](PlaceholderDetails.md) | The current details of the [placeholders](https://docs.talon.one/docs/product/campaigns/templates/create-templates#use-placeholders) in the campaign. | [optional] 
**EvaluationPosition** | Pointer to [**CampaignEvaluationPosition**](CampaignEvaluationPosition.md) |  | 

## Methods

### NewCampaignCreatedNotificationItem

`func NewCampaignCreatedNotificationItem(event string, campaign Campaign, evaluationPosition CampaignEvaluationPosition, ) *CampaignCreatedNotificationItem`

NewCampaignCreatedNotificationItem instantiates a new CampaignCreatedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreatedNotificationItemWithDefaults

`func NewCampaignCreatedNotificationItemWithDefaults() *CampaignCreatedNotificationItem`

NewCampaignCreatedNotificationItemWithDefaults instantiates a new CampaignCreatedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignCreatedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignCreatedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignCreatedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignCreatedNotificationItem) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignCreatedNotificationItem) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignCreatedNotificationItem) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.


### GetRuleset

`func (o *CampaignCreatedNotificationItem) GetRuleset() Ruleset`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignCreatedNotificationItem) GetRulesetOk() (*Ruleset, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignCreatedNotificationItem) SetRuleset(v Ruleset)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignCreatedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetPlaceholders

`func (o *CampaignCreatedNotificationItem) GetPlaceholders() []PlaceholderDetails`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *CampaignCreatedNotificationItem) GetPlaceholdersOk() (*[]PlaceholderDetails, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *CampaignCreatedNotificationItem) SetPlaceholders(v []PlaceholderDetails)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *CampaignCreatedNotificationItem) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### GetEvaluationPosition

`func (o *CampaignCreatedNotificationItem) GetEvaluationPosition() CampaignEvaluationPosition`

GetEvaluationPosition returns the EvaluationPosition field if non-nil, zero value otherwise.

### GetEvaluationPositionOk

`func (o *CampaignCreatedNotificationItem) GetEvaluationPositionOk() (*CampaignEvaluationPosition, bool)`

GetEvaluationPositionOk returns a tuple with the EvaluationPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPosition

`func (o *CampaignCreatedNotificationItem) SetEvaluationPosition(v CampaignEvaluationPosition)`

SetEvaluationPosition sets EvaluationPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


