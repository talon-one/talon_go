# RuleMetadataEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A short description of the rule. | 
**DisplayName** | Pointer to **string** | A customer-facing name for the rule. | [optional] 
**DisplayDescription** | Pointer to **string** | A customer-facing description that explains the details of the rule.   For example, this property can contain details about eligibility requirements, reward timelines, or terms and conditions.  | [optional] 
**RelatedData** | Pointer to **string** | Any additional data associated with the rule, such as an image URL, vendor name, or a content management system (CMS) ID.  | [optional] 
**Eligibility** | Pointer to [**[]RuleEligibility**](RuleEligibility.md) |  | 

## Methods

### NewRuleMetadataEligibility

`func NewRuleMetadataEligibility(title string, eligibility []RuleEligibility, ) *RuleMetadataEligibility`

NewRuleMetadataEligibility instantiates a new RuleMetadataEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleMetadataEligibilityWithDefaults

`func NewRuleMetadataEligibilityWithDefaults() *RuleMetadataEligibility`

NewRuleMetadataEligibilityWithDefaults instantiates a new RuleMetadataEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleMetadataEligibility) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleMetadataEligibility) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleMetadataEligibility) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisplayName

`func (o *RuleMetadataEligibility) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RuleMetadataEligibility) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RuleMetadataEligibility) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RuleMetadataEligibility) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *RuleMetadataEligibility) GetDisplayDescription() string`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *RuleMetadataEligibility) GetDisplayDescriptionOk() (*string, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *RuleMetadataEligibility) SetDisplayDescription(v string)`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *RuleMetadataEligibility) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetRelatedData

`func (o *RuleMetadataEligibility) GetRelatedData() string`

GetRelatedData returns the RelatedData field if non-nil, zero value otherwise.

### GetRelatedDataOk

`func (o *RuleMetadataEligibility) GetRelatedDataOk() (*string, bool)`

GetRelatedDataOk returns a tuple with the RelatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedData

`func (o *RuleMetadataEligibility) SetRelatedData(v string)`

SetRelatedData sets RelatedData field to given value.

### HasRelatedData

`func (o *RuleMetadataEligibility) HasRelatedData() bool`

HasRelatedData returns a boolean if a field has been set.

### GetEligibility

`func (o *RuleMetadataEligibility) GetEligibility() []RuleEligibility`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *RuleMetadataEligibility) GetEligibilityOk() (*[]RuleEligibility, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *RuleMetadataEligibility) SetEligibility(v []RuleEligibility)`

SetEligibility sets Eligibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


