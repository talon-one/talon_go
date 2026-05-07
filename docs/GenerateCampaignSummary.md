# GenerateCampaignSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignID** | Pointer to **int64** | ID of a campaign. | 
**RulesetID** | Pointer to **int64** | ID of a ruleset. | 
**Currency** | Pointer to **string** | Currency for the campaign. | 

## Methods

### NewGenerateCampaignSummary

`func NewGenerateCampaignSummary(campaignID int64, rulesetID int64, currency string, ) *GenerateCampaignSummary`

NewGenerateCampaignSummary instantiates a new GenerateCampaignSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCampaignSummaryWithDefaults

`func NewGenerateCampaignSummaryWithDefaults() *GenerateCampaignSummary`

NewGenerateCampaignSummaryWithDefaults instantiates a new GenerateCampaignSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignID

`func (o *GenerateCampaignSummary) GetCampaignID() int64`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *GenerateCampaignSummary) GetCampaignIDOk() (*int64, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *GenerateCampaignSummary) SetCampaignID(v int64)`

SetCampaignID sets CampaignID field to given value.


### GetRulesetID

`func (o *GenerateCampaignSummary) GetRulesetID() int64`

GetRulesetID returns the RulesetID field if non-nil, zero value otherwise.

### GetRulesetIDOk

`func (o *GenerateCampaignSummary) GetRulesetIDOk() (*int64, bool)`

GetRulesetIDOk returns a tuple with the RulesetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetID

`func (o *GenerateCampaignSummary) SetRulesetID(v int64)`

SetRulesetID sets RulesetID field to given value.


### GetCurrency

`func (o *GenerateCampaignSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GenerateCampaignSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GenerateCampaignSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


