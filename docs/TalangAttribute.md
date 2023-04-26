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
**CampaignsCount** | Pointer to **int32** | The number of campaigns that refer to the attribute. | 
**ExampleValue** | Pointer to **[]string** | Examples of values that can be assigned to the attribute. | [optional] 

## Methods

### GetEntity

`func (o *TalangAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TalangAttribute) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *TalangAttribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *TalangAttribute) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetName

`func (o *TalangAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TalangAttribute) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TalangAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TalangAttribute) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *TalangAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TalangAttribute) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *TalangAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *TalangAttribute) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetType

`func (o *TalangAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TalangAttribute) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *TalangAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *TalangAttribute) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *TalangAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TalangAttribute) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TalangAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TalangAttribute) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetVisible

`func (o *TalangAttribute) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TalangAttribute) GetVisibleOk() (bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisible

`func (o *TalangAttribute) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisible

`func (o *TalangAttribute) SetVisible(v bool)`

SetVisible gets a reference to the given bool and assigns it to the Visible field.

### GetKind

`func (o *TalangAttribute) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TalangAttribute) GetKindOk() (string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKind

`func (o *TalangAttribute) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKind

`func (o *TalangAttribute) SetKind(v string)`

SetKind gets a reference to the given string and assigns it to the Kind field.

### GetCampaignsCount

`func (o *TalangAttribute) GetCampaignsCount() int32`

GetCampaignsCount returns the CampaignsCount field if non-nil, zero value otherwise.

### GetCampaignsCountOk

`func (o *TalangAttribute) GetCampaignsCountOk() (int32, bool)`

GetCampaignsCountOk returns a tuple with the CampaignsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsCount

`func (o *TalangAttribute) HasCampaignsCount() bool`

HasCampaignsCount returns a boolean if a field has been set.

### SetCampaignsCount

`func (o *TalangAttribute) SetCampaignsCount(v int32)`

SetCampaignsCount gets a reference to the given int32 and assigns it to the CampaignsCount field.

### GetExampleValue

`func (o *TalangAttribute) GetExampleValue() []string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *TalangAttribute) GetExampleValueOk() ([]string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExampleValue

`func (o *TalangAttribute) HasExampleValue() bool`

HasExampleValue returns a boolean if a field has been set.

### SetExampleValue

`func (o *TalangAttribute) SetExampleValue(v []string)`

SetExampleValue gets a reference to the given []string and assigns it to the ExampleValue field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


