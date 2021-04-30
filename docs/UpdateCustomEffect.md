# UpdateCustomEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this effect. | 
**Title** | Pointer to **string** | The title of this effect. | 
**Payload** | Pointer to **string** | The JSON payload of this effect. | 
**Description** | Pointer to **string** | The description of this effect. | [optional] 
**Enabled** | Pointer to **bool** | Determines if this effect is active. | 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that this effect is enabled for | [optional] 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions | [optional] 

## Methods

### GetName

`func (o *UpdateCustomEffect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomEffect) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateCustomEffect) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateCustomEffect) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *UpdateCustomEffect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateCustomEffect) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *UpdateCustomEffect) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *UpdateCustomEffect) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetPayload

`func (o *UpdateCustomEffect) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UpdateCustomEffect) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *UpdateCustomEffect) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *UpdateCustomEffect) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetDescription

`func (o *UpdateCustomEffect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomEffect) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateCustomEffect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateCustomEffect) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetEnabled

`func (o *UpdateCustomEffect) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCustomEffect) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *UpdateCustomEffect) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *UpdateCustomEffect) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetSubscribedApplicationsIds

`func (o *UpdateCustomEffect) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *UpdateCustomEffect) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *UpdateCustomEffect) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *UpdateCustomEffect) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.

### GetParams

`func (o *UpdateCustomEffect) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateCustomEffect) GetParamsOk() ([]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *UpdateCustomEffect) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *UpdateCustomEffect) SetParams(v []TemplateArgDef)`

SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


