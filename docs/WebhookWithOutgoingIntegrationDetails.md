# WebhookWithOutgoingIntegrationDetails

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
**OutgoingIntegrationTemplateId** | Pointer to **int32** | Identifier of the outgoing integration template. | [optional] 
**OutgoingIntegrationTypeId** | Pointer to **int32** | Identifier of the outgoing integration type. | [optional] 
**OutgoingIntegrationTypeName** | Pointer to **string** | Name of the outgoing integration. | [optional] 

## Methods

### GetId

`func (o *WebhookWithOutgoingIntegrationDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *WebhookWithOutgoingIntegrationDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *WebhookWithOutgoingIntegrationDetails) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *WebhookWithOutgoingIntegrationDetails) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *WebhookWithOutgoingIntegrationDetails) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *WebhookWithOutgoingIntegrationDetails) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *WebhookWithOutgoingIntegrationDetails) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *WebhookWithOutgoingIntegrationDetails) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *WebhookWithOutgoingIntegrationDetails) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetApplicationIds

`func (o *WebhookWithOutgoingIntegrationDetails) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *WebhookWithOutgoingIntegrationDetails) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *WebhookWithOutgoingIntegrationDetails) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetTitle

`func (o *WebhookWithOutgoingIntegrationDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *WebhookWithOutgoingIntegrationDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *WebhookWithOutgoingIntegrationDetails) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *WebhookWithOutgoingIntegrationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *WebhookWithOutgoingIntegrationDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *WebhookWithOutgoingIntegrationDetails) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetVerb

`func (o *WebhookWithOutgoingIntegrationDetails) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetVerbOk() (string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerb

`func (o *WebhookWithOutgoingIntegrationDetails) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### SetVerb

`func (o *WebhookWithOutgoingIntegrationDetails) SetVerb(v string)`

SetVerb gets a reference to the given string and assigns it to the Verb field.

### GetUrl

`func (o *WebhookWithOutgoingIntegrationDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *WebhookWithOutgoingIntegrationDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *WebhookWithOutgoingIntegrationDetails) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetHeaders

`func (o *WebhookWithOutgoingIntegrationDetails) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *WebhookWithOutgoingIntegrationDetails) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *WebhookWithOutgoingIntegrationDetails) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetPayload

`func (o *WebhookWithOutgoingIntegrationDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *WebhookWithOutgoingIntegrationDetails) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *WebhookWithOutgoingIntegrationDetails) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetParams

`func (o *WebhookWithOutgoingIntegrationDetails) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetParamsOk() ([]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *WebhookWithOutgoingIntegrationDetails) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *WebhookWithOutgoingIntegrationDetails) SetParams(v []TemplateArgDef)`

SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.

### GetEnabled

`func (o *WebhookWithOutgoingIntegrationDetails) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *WebhookWithOutgoingIntegrationDetails) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *WebhookWithOutgoingIntegrationDetails) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateId() int32`

GetOutgoingIntegrationTemplateId returns the OutgoingIntegrationTemplateId field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTemplateIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateIdOk() (int32, bool)`

GetOutgoingIntegrationTemplateIdOk returns a tuple with the OutgoingIntegrationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTemplateId() bool`

HasOutgoingIntegrationTemplateId returns a boolean if a field has been set.

### SetOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTemplateId(v int32)`

SetOutgoingIntegrationTemplateId gets a reference to the given int32 and assigns it to the OutgoingIntegrationTemplateId field.

### GetOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeId() int32`

GetOutgoingIntegrationTypeId returns the OutgoingIntegrationTypeId field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTypeIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeIdOk() (int32, bool)`

GetOutgoingIntegrationTypeIdOk returns a tuple with the OutgoingIntegrationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeId() bool`

HasOutgoingIntegrationTypeId returns a boolean if a field has been set.

### SetOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeId(v int32)`

SetOutgoingIntegrationTypeId gets a reference to the given int32 and assigns it to the OutgoingIntegrationTypeId field.

### GetOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeName() string`

GetOutgoingIntegrationTypeName returns the OutgoingIntegrationTypeName field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTypeNameOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeNameOk() (string, bool)`

GetOutgoingIntegrationTypeNameOk returns a tuple with the OutgoingIntegrationTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeName() bool`

HasOutgoingIntegrationTypeName returns a boolean if a field has been set.

### SetOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeName(v string)`

SetOutgoingIntegrationTypeName gets a reference to the given string and assigns it to the OutgoingIntegrationTypeName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


