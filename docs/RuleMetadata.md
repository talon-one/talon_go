# RuleMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A short description of the rule. | 
**DisplayName** | Pointer to **string** | A customer-facing name for the rule. | [optional] 
**DisplayDescription** | Pointer to **string** | A customer-facing description that explains the details of the rule.   For example, this property can contain details about eligibility requirements, reward timelines, or terms and conditions.  | [optional] 
**RelatedData** | Pointer to **string** | Any additional data associated with the rule, such as an image URL, vendor name, or a content management system (CMS) ID.  | [optional] 

## Methods

### NewRuleMetadata

`func NewRuleMetadata(title string, ) *RuleMetadata`

NewRuleMetadata instantiates a new RuleMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleMetadataWithDefaults

`func NewRuleMetadataWithDefaults() *RuleMetadata`

NewRuleMetadataWithDefaults instantiates a new RuleMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDisplayName

`func (o *RuleMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RuleMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RuleMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RuleMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *RuleMetadata) GetDisplayDescription() string`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *RuleMetadata) GetDisplayDescriptionOk() (*string, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *RuleMetadata) SetDisplayDescription(v string)`

SetDisplayDescription sets DisplayDescription field to given value.

### HasDisplayDescription

`func (o *RuleMetadata) HasDisplayDescription() bool`

HasDisplayDescription returns a boolean if a field has been set.

### GetRelatedData

`func (o *RuleMetadata) GetRelatedData() string`

GetRelatedData returns the RelatedData field if non-nil, zero value otherwise.

### GetRelatedDataOk

`func (o *RuleMetadata) GetRelatedDataOk() (*string, bool)`

GetRelatedDataOk returns a tuple with the RelatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedData

`func (o *RuleMetadata) SetRelatedData(v string)`

SetRelatedData sets RelatedData field to given value.

### HasRelatedData

`func (o *RuleMetadata) HasRelatedData() bool`

HasRelatedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


