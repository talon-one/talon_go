# WebhookLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID reference of the webhook request | 
**IntegrationRequestUuid** | Pointer to **string** | UUID reference of the integration request linked to this webhook request | 
**WebhookId** | Pointer to **int32** | ID of the webhook that triggered the request | 
**ApplicationId** | Pointer to **int32** | ID of the application that triggered the webhook | [optional] 
**Url** | Pointer to **string** | Target url of request | 
**Request** | Pointer to **string** | Request message | 
**Response** | Pointer to **string** | Response message | [optional] 
**Status** | Pointer to **int32** | HTTP status code of response | [optional] 
**RequestTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of request | 
**ResponseTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of response | [optional] 

## Methods

### GetId

`func (o *WebhookLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookLogEntry) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *WebhookLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *WebhookLogEntry) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIntegrationRequestUuid

`func (o *WebhookLogEntry) GetIntegrationRequestUuid() string`

GetIntegrationRequestUuid returns the IntegrationRequestUuid field if non-nil, zero value otherwise.

### GetIntegrationRequestUuidOk

`func (o *WebhookLogEntry) GetIntegrationRequestUuidOk() (string, bool)`

GetIntegrationRequestUuidOk returns a tuple with the IntegrationRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationRequestUuid

`func (o *WebhookLogEntry) HasIntegrationRequestUuid() bool`

HasIntegrationRequestUuid returns a boolean if a field has been set.

### SetIntegrationRequestUuid

`func (o *WebhookLogEntry) SetIntegrationRequestUuid(v string)`

SetIntegrationRequestUuid gets a reference to the given string and assigns it to the IntegrationRequestUuid field.

### GetWebhookId

`func (o *WebhookLogEntry) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookLogEntry) GetWebhookIdOk() (int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhookId

`func (o *WebhookLogEntry) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookId

`func (o *WebhookLogEntry) SetWebhookId(v int32)`

SetWebhookId gets a reference to the given int32 and assigns it to the WebhookId field.

### GetApplicationId

`func (o *WebhookLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *WebhookLogEntry) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *WebhookLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *WebhookLogEntry) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUrl

`func (o *WebhookLogEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookLogEntry) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *WebhookLogEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *WebhookLogEntry) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### GetRequest

`func (o *WebhookLogEntry) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WebhookLogEntry) GetRequestOk() (string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequest

`func (o *WebhookLogEntry) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequest

`func (o *WebhookLogEntry) SetRequest(v string)`

SetRequest gets a reference to the given string and assigns it to the Request field.

### GetResponse

`func (o *WebhookLogEntry) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookLogEntry) GetResponseOk() (string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponse

`func (o *WebhookLogEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponse

`func (o *WebhookLogEntry) SetResponse(v string)`

SetResponse gets a reference to the given string and assigns it to the Response field.

### GetStatus

`func (o *WebhookLogEntry) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookLogEntry) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *WebhookLogEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *WebhookLogEntry) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.

### GetRequestTime

`func (o *WebhookLogEntry) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *WebhookLogEntry) GetRequestTimeOk() (time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestTime

`func (o *WebhookLogEntry) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### SetRequestTime

`func (o *WebhookLogEntry) SetRequestTime(v time.Time)`

SetRequestTime gets a reference to the given time.Time and assigns it to the RequestTime field.

### GetResponseTime

`func (o *WebhookLogEntry) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WebhookLogEntry) GetResponseTimeOk() (time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseTime

`func (o *WebhookLogEntry) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### SetResponseTime

`func (o *WebhookLogEntry) SetResponseTime(v time.Time)`

SetResponseTime gets a reference to the given time.Time and assigns it to the ResponseTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


