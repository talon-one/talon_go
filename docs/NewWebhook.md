# NewWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | Pointer to **[]int32** | The IDs of the Applications that are related to this entity. | 
**Title** | Pointer to **string** | Name or title for this webhook. | 
**Description** | Pointer to **string** | A description of the webhook. | [optional] 
**Verb** | Pointer to **string** | API method for this webhook. | 
**Url** | Pointer to **string** | API URL (supports templating using parameters) for this webhook. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for this webhook. | 
**Payload** | Pointer to **string** | API payload (supports templating using parameters) for this webhook. | [optional] 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | 
**Enabled** | Pointer to **bool** | Enables or disables webhook from showing in the Rule Builder. | 

## Methods

### GetApplicationIds

`func (o *NewWebhook) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewWebhook) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *NewWebhook) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *NewWebhook) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetTitle

`func (o *NewWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewWebhook) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewWebhook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewWebhook) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *NewWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewWebhook) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewWebhook) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetVerb

`func (o *NewWebhook) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *NewWebhook) GetVerbOk() (string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerb

`func (o *NewWebhook) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### SetVerb

`func (o *NewWebhook) SetVerb(v string)`

SetVerb gets a reference to the given string and assigns it to the Verb field.

### GetUrl

`func (o *NewWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewWebhook) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *NewWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *NewWebhook) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetHeaders

`func (o *NewWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NewWebhook) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *NewWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *NewWebhook) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetPayload

`func (o *NewWebhook) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NewWebhook) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *NewWebhook) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *NewWebhook) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetParams

`func (o *NewWebhook) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NewWebhook) GetParamsOk() ([]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *NewWebhook) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *NewWebhook) SetParams(v []TemplateArgDef)`

SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.

### GetEnabled

`func (o *NewWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewWebhook) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *NewWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *NewWebhook) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


