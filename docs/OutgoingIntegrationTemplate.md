# OutgoingIntegrationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**IntegrationType** | Pointer to **int32** | Unique ID of outgoing integration type. | 
**Title** | Pointer to **string** | The title of the integration template. | 
**Description** | Pointer to **string** | The description of the specific outgoing integration template. | 
**Payload** | Pointer to **string** | The API payload (supports templating using parameters) for this integration template. | 
**Method** | Pointer to **string** | API method for this webhook. | 
**RelativeUrl** | Pointer to **string** | The relative URL corresponding to each integration template. | 
**Headers** | Pointer to **[]string** | The list of HTTP headers for this integration template. | 

## Methods

### GetId

`func (o *OutgoingIntegrationTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationTemplate) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *OutgoingIntegrationTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *OutgoingIntegrationTemplate) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetIntegrationType

`func (o *OutgoingIntegrationTemplate) GetIntegrationType() int32`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *OutgoingIntegrationTemplate) GetIntegrationTypeOk() (int32, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationType

`func (o *OutgoingIntegrationTemplate) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationType

`func (o *OutgoingIntegrationTemplate) SetIntegrationType(v int32)`

SetIntegrationType gets a reference to the given int32 and assigns it to the IntegrationType field.

### GetTitle

`func (o *OutgoingIntegrationTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutgoingIntegrationTemplate) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *OutgoingIntegrationTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *OutgoingIntegrationTemplate) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *OutgoingIntegrationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingIntegrationTemplate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *OutgoingIntegrationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *OutgoingIntegrationTemplate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetPayload

`func (o *OutgoingIntegrationTemplate) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutgoingIntegrationTemplate) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *OutgoingIntegrationTemplate) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *OutgoingIntegrationTemplate) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetMethod

`func (o *OutgoingIntegrationTemplate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutgoingIntegrationTemplate) GetMethodOk() (string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethod

`func (o *OutgoingIntegrationTemplate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethod

`func (o *OutgoingIntegrationTemplate) SetMethod(v string)`

SetMethod gets a reference to the given string and assigns it to the Method field.

### GetRelativeUrl

`func (o *OutgoingIntegrationTemplate) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *OutgoingIntegrationTemplate) GetRelativeUrlOk() (string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRelativeUrl

`func (o *OutgoingIntegrationTemplate) HasRelativeUrl() bool`

HasRelativeUrl returns a boolean if a field has been set.

### SetRelativeUrl

`func (o *OutgoingIntegrationTemplate) SetRelativeUrl(v string)`

SetRelativeUrl gets a reference to the given string and assigns it to the RelativeUrl field.

### GetHeaders

`func (o *OutgoingIntegrationTemplate) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutgoingIntegrationTemplate) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *OutgoingIntegrationTemplate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *OutgoingIntegrationTemplate) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


