# OutgoingIntegrationTemplateWithConfigurationDetails

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
**Policy** | Pointer to [**map[string]interface{}**](.md) | The outgoing integration policy specific to each integration type. | 

## Methods

### GetId

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetIntegrationType

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIntegrationType() int32`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIntegrationTypeOk() (int32, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationType

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationType

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetIntegrationType(v int32)`

SetIntegrationType gets a reference to the given int32 and assigns it to the IntegrationType field.

### GetTitle

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetPayload

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetMethod

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetMethodOk() (string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethod

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethod

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetMethod(v string)`

SetMethod gets a reference to the given string and assigns it to the Method field.

### GetRelativeUrl

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetRelativeUrlOk() (string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRelativeUrl

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasRelativeUrl() bool`

HasRelativeUrl returns a boolean if a field has been set.

### SetRelativeUrl

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetRelativeUrl(v string)`

SetRelativeUrl gets a reference to the given string and assigns it to the RelativeUrl field.

### GetHeaders

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetHeadersOk() ([]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetHeaders(v []string)`

SetHeaders gets a reference to the given []string and assigns it to the Headers field.

### GetPolicy

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPolicyOk() (map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPolicy(v map[string]interface{})`

SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


