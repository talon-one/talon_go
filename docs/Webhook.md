# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**ApplicationIds** | Pointer to **[]int32** | The IDs of the Applications that are related to this entity. The IDs of the Applications that are related to this entity. | 
**Title** | Pointer to **string** | Name or title for this webhook. | 
**Description** | Pointer to **string** | A description of the webhook. | [optional] 
**Verb** | Pointer to **string** | API method for this webhook. | 
**Url** | Pointer to **string** | API URL (supports templating using parameters) for this webhook. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for this webhook. | 
**Payload** | Pointer to **string** | API payload (supports templating using parameters) for this webhook. | [optional] 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | 
**Enabled** | Pointer to **bool** | Enables or disables webhook from showing in the Rule Builder. | 

## Methods

### GetId

`func (o *Webhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Webhook) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Webhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Webhook) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Webhook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Webhook) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *Webhook) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Webhook) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *Webhook) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *Webhook) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetApplicationIds

`func (o *Webhook) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *Webhook) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *Webhook) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *Webhook) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetTitle

`func (o *Webhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Webhook) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Webhook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Webhook) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *Webhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webhook) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Webhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Webhook) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetVerb

`func (o *Webhook) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *Webhook) GetVerbOk() (string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerb

`func (o *Webhook) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### SetVerb

`func (o *Webhook) SetVerb(v string)`

SetVerb gets a reference to the given string and assigns it to the Verb field.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetHeaders

`func (o *Webhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Webhook) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *Webhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *Webhook) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetPayload

`func (o *Webhook) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Webhook) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *Webhook) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *Webhook) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetParams

`func (o *Webhook) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Webhook) GetParamsOk() ([]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *Webhook) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *Webhook) SetParams(v []TemplateArgDef)`

SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.

### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


