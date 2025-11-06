# NewMessageTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]string** | List of API HTTP headers for the given message. | [optional] 
**Verb** | Pointer to **string** | API method for this message. | 
**Url** | Pointer to **string** | API URL for the given message. | 
**Payload** | Pointer to **string** | API payload of this message. | [optional] 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications in which this webhook is available. An empty array means the webhook is available in &#x60;All Applications&#x60;.  | [optional] 

## Methods

### NewNewMessageTest

`func NewNewMessageTest(verb string, url string, ) *NewMessageTest`

NewNewMessageTest instantiates a new NewMessageTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMessageTestWithDefaults

`func NewNewMessageTestWithDefaults() *NewMessageTest`

NewNewMessageTestWithDefaults instantiates a new NewMessageTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *NewMessageTest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NewMessageTest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NewMessageTest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *NewMessageTest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetVerb

`func (o *NewMessageTest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *NewMessageTest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *NewMessageTest) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetUrl

`func (o *NewMessageTest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewMessageTest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewMessageTest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPayload

`func (o *NewMessageTest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NewMessageTest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NewMessageTest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NewMessageTest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetParams

`func (o *NewMessageTest) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NewMessageTest) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NewMessageTest) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.

### HasParams

`func (o *NewMessageTest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetApplicationIds

`func (o *NewMessageTest) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewMessageTest) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewMessageTest) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.

### HasApplicationIds

`func (o *NewMessageTest) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


