# LabelTargetAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Audience** | Pointer to [**AudienceReference**](AudienceReference.md) |  | 

## Methods

### NewLabelTargetAudience

`func NewLabelTargetAudience(type_ string, audience AudienceReference, ) *LabelTargetAudience`

NewLabelTargetAudience instantiates a new LabelTargetAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelTargetAudienceWithDefaults

`func NewLabelTargetAudienceWithDefaults() *LabelTargetAudience`

NewLabelTargetAudienceWithDefaults instantiates a new LabelTargetAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LabelTargetAudience) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelTargetAudience) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelTargetAudience) SetType(v string)`

SetType sets Type field to given value.


### GetAudience

`func (o *LabelTargetAudience) GetAudience() AudienceReference`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *LabelTargetAudience) GetAudienceOk() (*AudienceReference, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *LabelTargetAudience) SetAudience(v AudienceReference)`

SetAudience sets Audience field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


