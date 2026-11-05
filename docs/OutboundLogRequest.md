# OutboundLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | HTTP method of the outbound request. | 
**Url** | Pointer to **string** | Target URL of the outbound request. | 
**Headers** | Pointer to **[]string** | HTTP headers sent with the outbound request. | 
**Body** | Pointer to [**map[string]interface{}**](.md) | JSON request payload. | 

## Methods

### NewOutboundLogRequest

`func NewOutboundLogRequest(method string, url string, headers []string, body map[string]interface{}, ) *OutboundLogRequest`

NewOutboundLogRequest instantiates a new OutboundLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogRequestWithDefaults

`func NewOutboundLogRequestWithDefaults() *OutboundLogRequest`

NewOutboundLogRequestWithDefaults instantiates a new OutboundLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OutboundLogRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OutboundLogRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OutboundLogRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *OutboundLogRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OutboundLogRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OutboundLogRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *OutboundLogRequest) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OutboundLogRequest) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OutboundLogRequest) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetBody

`func (o *OutboundLogRequest) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *OutboundLogRequest) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *OutboundLogRequest) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


