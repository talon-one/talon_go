# UpdateAudienceMembershipBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | The action to perform. | 
**Profile** | Pointer to **string** | The customer profile to add or remove from the audience. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Audience** | Pointer to [**UpdateAudienceMembershipBlockAudience**](UpdateAudienceMembershipBlock_audience.md) |  | 

## Methods

### NewUpdateAudienceMembershipBlock

`func NewUpdateAudienceMembershipBlock(type_ string, operator string, profile string, audience UpdateAudienceMembershipBlockAudience, ) *UpdateAudienceMembershipBlock`

NewUpdateAudienceMembershipBlock instantiates a new UpdateAudienceMembershipBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAudienceMembershipBlockWithDefaults

`func NewUpdateAudienceMembershipBlockWithDefaults() *UpdateAudienceMembershipBlock`

NewUpdateAudienceMembershipBlockWithDefaults instantiates a new UpdateAudienceMembershipBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAudienceMembershipBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAudienceMembershipBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAudienceMembershipBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAudienceMembershipBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateAudienceMembershipBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAudienceMembershipBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAudienceMembershipBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *UpdateAudienceMembershipBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAudienceMembershipBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAudienceMembershipBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAudienceMembershipBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateAudienceMembershipBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateAudienceMembershipBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateAudienceMembershipBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetProfile

`func (o *UpdateAudienceMembershipBlock) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateAudienceMembershipBlock) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateAudienceMembershipBlock) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAudience

`func (o *UpdateAudienceMembershipBlock) GetAudience() UpdateAudienceMembershipBlockAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateAudienceMembershipBlock) GetAudienceOk() (*UpdateAudienceMembershipBlockAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateAudienceMembershipBlock) SetAudience(v UpdateAudienceMembershipBlockAudience)`

SetAudience sets Audience field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


