# RulesetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Internal ID of this entity. | [optional] [readonly] 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | [optional] [readonly] 
**UserId** | Pointer to **int64** | The ID of the user that created this ruleset. | [optional] [readonly] 
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | [optional] [readonly] 
**TemplateId** | Pointer to **int64** | The ID of the campaign template that owns this entity. | [optional] [readonly] 
**ActivatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp indicating when this ruleset was activated. | [optional] [readonly] 
**PromotionRules** | Pointer to [**[]RuleV2**](RuleV2.md) | Set of promotion rules. | 
**StrikethroughRules** | Pointer to [**[]RuleV2**](RuleV2.md) | Set of strikethrough rules. | [optional] 
**Selectors** | Pointer to [**[]Selector**](Selector.md) | Variable bindings of type selector. | [optional] [readonly] 
**Bundles** | Pointer to [**[]Bundle**](Bundle.md) | Variable bindings of type bundle. | [optional] [readonly] 
**Parameters** | Pointer to [**[]TemplateParameter**](TemplateParameter.md) | Variable bindings of type template parameter. | [optional] [readonly] 

## Methods

### NewRulesetV2

`func NewRulesetV2(promotionRules []RuleV2, ) *RulesetV2`

NewRulesetV2 instantiates a new RulesetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetV2WithDefaults

`func NewRulesetV2WithDefaults() *RulesetV2`

NewRulesetV2WithDefaults instantiates a new RulesetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RulesetV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RulesetV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RulesetV2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RulesetV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *RulesetV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RulesetV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RulesetV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RulesetV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUserId

`func (o *RulesetV2) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RulesetV2) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RulesetV2) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RulesetV2) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCampaignId

`func (o *RulesetV2) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *RulesetV2) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *RulesetV2) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *RulesetV2) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetTemplateId

`func (o *RulesetV2) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RulesetV2) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RulesetV2) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *RulesetV2) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetActivatedAt

`func (o *RulesetV2) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *RulesetV2) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *RulesetV2) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *RulesetV2) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetPromotionRules

`func (o *RulesetV2) GetPromotionRules() []RuleV2`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *RulesetV2) GetPromotionRulesOk() (*[]RuleV2, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *RulesetV2) SetPromotionRules(v []RuleV2)`

SetPromotionRules sets PromotionRules field to given value.


### GetStrikethroughRules

`func (o *RulesetV2) GetStrikethroughRules() []RuleV2`

GetStrikethroughRules returns the StrikethroughRules field if non-nil, zero value otherwise.

### GetStrikethroughRulesOk

`func (o *RulesetV2) GetStrikethroughRulesOk() (*[]RuleV2, bool)`

GetStrikethroughRulesOk returns a tuple with the StrikethroughRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikethroughRules

`func (o *RulesetV2) SetStrikethroughRules(v []RuleV2)`

SetStrikethroughRules sets StrikethroughRules field to given value.

### HasStrikethroughRules

`func (o *RulesetV2) HasStrikethroughRules() bool`

HasStrikethroughRules returns a boolean if a field has been set.

### GetSelectors

`func (o *RulesetV2) GetSelectors() []Selector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *RulesetV2) GetSelectorsOk() (*[]Selector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *RulesetV2) SetSelectors(v []Selector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *RulesetV2) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetBundles

`func (o *RulesetV2) GetBundles() []Bundle`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *RulesetV2) GetBundlesOk() (*[]Bundle, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *RulesetV2) SetBundles(v []Bundle)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *RulesetV2) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetParameters

`func (o *RulesetV2) GetParameters() []TemplateParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RulesetV2) GetParametersOk() (*[]TemplateParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RulesetV2) SetParameters(v []TemplateParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RulesetV2) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


