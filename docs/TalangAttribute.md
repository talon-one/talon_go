# TalangAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The name of the entity of the attribute. | [optional] 
**Name** | Pointer to **string** | The attribute name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;region\&quot;&#x60; then you would set the region attribute by including an &#x60;attributes.region&#x60; property in your request payload.  | 
**Title** | Pointer to **string** | The name of the attribute that is displayed to the user in the Campaign Manager. | [optional] 
**Type** | Pointer to **string** | The talang type of the attribute. | 
**Description** | Pointer to **string** | A description of the attribute. | [optional] 
**Visible** | Pointer to **bool** | Indicates whether the attribute is visible in the UI or not. | [default to true]
**Kind** | Pointer to **string** | Indicate the kind of the attribute. | 
**CampaignsCount** | Pointer to **int64** | The number of campaigns that refer to the attribute. | 
**ExampleValue** | Pointer to **[]string** | Examples of values that can be assigned to the attribute. | [optional] 

## Methods

### NewTalangAttribute

`func NewTalangAttribute(name string, type_ string, visible bool, kind string, campaignsCount int64, ) *TalangAttribute`

NewTalangAttribute instantiates a new TalangAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTalangAttributeWithDefaults

`func NewTalangAttributeWithDefaults() *TalangAttribute`

NewTalangAttributeWithDefaults instantiates a new TalangAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *TalangAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TalangAttribute) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *TalangAttribute) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *TalangAttribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetName

`func (o *TalangAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TalangAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TalangAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *TalangAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TalangAttribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TalangAttribute) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TalangAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *TalangAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TalangAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TalangAttribute) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *TalangAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TalangAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TalangAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TalangAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisible

`func (o *TalangAttribute) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TalangAttribute) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TalangAttribute) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetKind

`func (o *TalangAttribute) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TalangAttribute) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TalangAttribute) SetKind(v string)`

SetKind sets Kind field to given value.


### GetCampaignsCount

`func (o *TalangAttribute) GetCampaignsCount() int64`

GetCampaignsCount returns the CampaignsCount field if non-nil, zero value otherwise.

### GetCampaignsCountOk

`func (o *TalangAttribute) GetCampaignsCountOk() (*int64, bool)`

GetCampaignsCountOk returns a tuple with the CampaignsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignsCount

`func (o *TalangAttribute) SetCampaignsCount(v int64)`

SetCampaignsCount sets CampaignsCount field to given value.


### GetExampleValue

`func (o *TalangAttribute) GetExampleValue() []string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *TalangAttribute) GetExampleValueOk() (*[]string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleValue

`func (o *TalangAttribute) SetExampleValue(v []string)`

SetExampleValue sets ExampleValue field to given value.

### HasExampleValue

`func (o *TalangAttribute) HasExampleValue() bool`

HasExampleValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


