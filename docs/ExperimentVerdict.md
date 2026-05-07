# ExperimentVerdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WinnerVariantName** | Pointer to **string** | The name of the winning variant. If no variant shows a statistically significant advantage on key business metrics, return &#39;Inconclusive&#39;. | 
**VerdictSummary** | Pointer to **string** | A one-sentence summary of the outcome, including the key metric and confidence level that led to the decision. | 
**KeyFindings** | Pointer to **[]string** | A bullet point stating the most important finding, including the metric, the percentage change, and the confidence. | 
**AiConfidenceLevel** | Pointer to **string** | Your confidence in this overall verdict, from 0 to 100. | 
**Recommendation** | Pointer to **string** | A short, actionable recommendation based on the findings. If inconclusive, suggest running the test longer. If there is a clear winner, recommend promoting it. | 

## Methods

### NewExperimentVerdict

`func NewExperimentVerdict(winnerVariantName string, verdictSummary string, keyFindings []string, aiConfidenceLevel string, recommendation string, ) *ExperimentVerdict`

NewExperimentVerdict instantiates a new ExperimentVerdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVerdictWithDefaults

`func NewExperimentVerdictWithDefaults() *ExperimentVerdict`

NewExperimentVerdictWithDefaults instantiates a new ExperimentVerdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWinnerVariantName

`func (o *ExperimentVerdict) GetWinnerVariantName() string`

GetWinnerVariantName returns the WinnerVariantName field if non-nil, zero value otherwise.

### GetWinnerVariantNameOk

`func (o *ExperimentVerdict) GetWinnerVariantNameOk() (*string, bool)`

GetWinnerVariantNameOk returns a tuple with the WinnerVariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerVariantName

`func (o *ExperimentVerdict) SetWinnerVariantName(v string)`

SetWinnerVariantName sets WinnerVariantName field to given value.


### GetVerdictSummary

`func (o *ExperimentVerdict) GetVerdictSummary() string`

GetVerdictSummary returns the VerdictSummary field if non-nil, zero value otherwise.

### GetVerdictSummaryOk

`func (o *ExperimentVerdict) GetVerdictSummaryOk() (*string, bool)`

GetVerdictSummaryOk returns a tuple with the VerdictSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdictSummary

`func (o *ExperimentVerdict) SetVerdictSummary(v string)`

SetVerdictSummary sets VerdictSummary field to given value.


### GetKeyFindings

`func (o *ExperimentVerdict) GetKeyFindings() []string`

GetKeyFindings returns the KeyFindings field if non-nil, zero value otherwise.

### GetKeyFindingsOk

`func (o *ExperimentVerdict) GetKeyFindingsOk() (*[]string, bool)`

GetKeyFindingsOk returns a tuple with the KeyFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFindings

`func (o *ExperimentVerdict) SetKeyFindings(v []string)`

SetKeyFindings sets KeyFindings field to given value.


### GetAiConfidenceLevel

`func (o *ExperimentVerdict) GetAiConfidenceLevel() string`

GetAiConfidenceLevel returns the AiConfidenceLevel field if non-nil, zero value otherwise.

### GetAiConfidenceLevelOk

`func (o *ExperimentVerdict) GetAiConfidenceLevelOk() (*string, bool)`

GetAiConfidenceLevelOk returns a tuple with the AiConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfidenceLevel

`func (o *ExperimentVerdict) SetAiConfidenceLevel(v string)`

SetAiConfidenceLevel sets AiConfidenceLevel field to given value.


### GetRecommendation

`func (o *ExperimentVerdict) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ExperimentVerdict) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ExperimentVerdict) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


