# TalangAttributeAllOf

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

`func (o *TalangAttributeAllOf) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TalangAttributeAllOf) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *TalangAttributeAllOf) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *TalangAttributeAllOf) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetName

`func (o *TalangAttributeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TalangAttributeAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TalangAttributeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TalangAttributeAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *TalangAttributeAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TalangAttributeAllOf) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *TalangAttributeAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *TalangAttributeAllOf) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetType

`func (o *TalangAttributeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TalangAttributeAllOf) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *TalangAttributeAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *TalangAttributeAllOf) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *TalangAttributeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TalangAttributeAllOf) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TalangAttributeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TalangAttributeAllOf) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetVisible

`func (o *TalangAttributeAllOf) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TalangAttributeAllOf) GetVisibleOk() (bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisible

`func (o *TalangAttributeAllOf) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisible

`func (o *TalangAttributeAllOf) SetVisible(v bool)`

SetVisible gets a reference to the given bool and assigns it to the Visible field.

### GetKind

`func (o *TalangAttributeAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TalangAttributeAllOf) GetKindOk() (string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKind

`func (o *TalangAttributeAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKind

`func (o *TalangAttributeAllOf) SetKind(v string)`

SetKind gets a reference to the given string and assigns it to the Kind field.

### GetCampaignsCount

`func (o *TalangAttributeAllOf) GetCampaignsCount() int32`

GetCampaignsCount returns the CampaignsCount field if non-nil, zero value otherwise.

### GetCampaignsCountOk

`func (o *TalangAttributeAllOf) GetCampaignsCountOk() (int32, bool)`

GetCampaignsCountOk returns a tuple with the CampaignsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignsCount

`func (o *TalangAttributeAllOf) HasCampaignsCount() bool`

HasCampaignsCount returns a boolean if a field has been set.

### SetCampaignsCount

`func (o *TalangAttributeAllOf) SetCampaignsCount(v int32)`

SetCampaignsCount gets a reference to the given int32 and assigns it to the CampaignsCount field.

### GetExampleValue

`func (o *TalangAttributeAllOf) GetExampleValue() []string`

GetExampleValue returns the ExampleValue field if non-nil, zero value otherwise.

### GetExampleValueOk

`func (o *TalangAttributeAllOf) GetExampleValueOk() ([]string, bool)`

GetExampleValueOk returns a tuple with the ExampleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExampleValue

`func (o *TalangAttributeAllOf) HasExampleValue() bool`

HasExampleValue returns a boolean if a field has been set.

### SetExampleValue

`func (o *TalangAttributeAllOf) SetExampleValue(v []string)`

SetExampleValue gets a reference to the given []string and assigns it to the ExampleValue field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


