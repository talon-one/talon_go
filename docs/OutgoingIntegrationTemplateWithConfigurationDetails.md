# OutgoingIntegrationTemplateWithConfigurationDetails

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
**Policy** | Pointer to [**map[string]interface{}**](.md) | The outgoing integration policy specific to each integration type. | 

## Methods

### NewOutgoingIntegrationTemplateWithConfigurationDetails

`func NewOutgoingIntegrationTemplateWithConfigurationDetails(id int64, integrationType int64, title string, description string, payload string, method string, relativeUrl string, headers []string, policy map[string]interface{}, ) *OutgoingIntegrationTemplateWithConfigurationDetails`

NewOutgoingIntegrationTemplateWithConfigurationDetails instantiates a new OutgoingIntegrationTemplateWithConfigurationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationTemplateWithConfigurationDetailsWithDefaults

`func NewOutgoingIntegrationTemplateWithConfigurationDetailsWithDefaults() *OutgoingIntegrationTemplateWithConfigurationDetails`

NewOutgoingIntegrationTemplateWithConfigurationDetailsWithDefaults instantiates a new OutgoingIntegrationTemplateWithConfigurationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegrationType

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIntegrationType() int64`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIntegrationTypeOk() (*int64, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetIntegrationType(v int64)`

SetIntegrationType sets IntegrationType field to given value.


### GetTitle

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPayload

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetMethod

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRelativeUrl

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetRelativeUrl() string`

GetRelativeUrl returns the RelativeUrl field if non-nil, zero value otherwise.

### GetRelativeUrlOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetRelativeUrlOk() (*string, bool)`

GetRelativeUrlOk returns a tuple with the RelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeUrl

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetRelativeUrl(v string)`

SetRelativeUrl sets RelativeUrl field to given value.


### GetHeaders

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetPolicy

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


