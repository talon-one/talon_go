# OutgoingIntegrationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**IntegrationType** | Pointer to **int64** | Unique ID of outgoing integration type. | 
**Title** | Pointer to **string** | The title of the integration template. | 
**Description** | Pointer to **string** | The description of the specific outgoing integration template. | 
**Payload** | Pointer to **string** | The API payload (supports templating using parameters) for this integration template. | 
**Method** | Pointer to **string** | API method for this webhook. | 
**RelativeUrl** | Pointer to **string** | The relative URL corresponding to each integration template. | 
**Headers** | Pointer to **[]string** | The list of HTTP headers for this integration template. | 

## Methods

### NewOutgoingIntegrationTemplate

`func NewOutgoingIntegrationTemplate(id int64, integrationType int64, title string, description string, payload string, method string, relativeUrl string, headers []string, ) *OutgoingIntegrationTemplate`

NewOutgoingIntegrationTemplate instantiates a new OutgoingIntegrationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationTemplateWithDefaults

`func NewOutgoingIntegrationTemplateWithDefaults() *OutgoingIntegrationTemplate`

NewOutgoingIntegrationTemplateWithDefaults instantiates a new OutgoingIntegrationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutgoingIntegrationTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingIntegrationTemplate) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegrationType

`func (o *OutgoingIntegrationTemplate) GetIntegrationType() int64`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *OutgoingIntegrationTemplate) GetIntegrationTypeOk() (*int64, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *OutgoingIntegrationTemplate) SetIntegrationType(v int64)`

SetIntegrationType sets IntegrationType field to given value.


### GetTitle

`func (o *OutgoingIntegrationTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutgoingIntegrationTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutgoingIntegrationTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OutgoingIntegrationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingIntegrationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingIntegrationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *OutgoingIntegrationTemplate) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutgoingIntegrationTemplate) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OutgoingIntegrationTemplate) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetMethod

`func (o *OutgoingIntegrationTemplate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutgoingIntegrationTemplate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OutgoingIntegrationTemplate) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRelativeUrl

`func (o *OutgoingIntegrationTemplate) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *OutgoingIntegrationTemplate) GetRelativeUrlOk() (*string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeUrl

`func (o *OutgoingIntegrationTemplate) SetRelativeUrl(v string)`

SetRelativeUrl sets RelativeUrl field to given value.


### GetHeaders

`func (o *OutgoingIntegrationTemplate) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutgoingIntegrationTemplate) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutgoingIntegrationTemplate) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


