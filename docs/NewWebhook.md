# NewWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications in which this webhook is available. An empty array means the webhook is available in &#x60;All Applications&#x60;.  | 
**Title** | Pointer to **string** | Name or title for this webhook. | 
**Description** | Pointer to **string** | A description of the webhook. | [optional] 
**Draft** | Pointer to **bool** | Indicates if the webhook is a draft. | 
**Verb** | Pointer to **string** | API method for this webhook. | 
**Url** | Pointer to **string** | API URL (supports templating using parameters) for this webhook. | 
**Headers** | Pointer to **[]string** | List of API HTTP headers for this webhook. | 
**Payload** | Pointer to **string** | API payload (supports templating using parameters) for this webhook. | [optional] 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | 
**Enabled** | Pointer to **bool** | Enables or disables webhook from showing in the Rule Builder. | 
**AuthenticationId** | Pointer to **int64** | The ID of the credential that this webhook is using. | [optional] 

## Methods

### NewNewWebhook

`func NewNewWebhook(applicationIds []int64, title string, draft bool, verb string, url string, headers []string, params []TemplateArgDef, enabled bool, ) *NewWebhook`

NewNewWebhook instantiates a new NewWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewWebhookWithDefaults

`func NewNewWebhookWithDefaults() *NewWebhook`

NewNewWebhookWithDefaults instantiates a new NewWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *NewWebhook) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewWebhook) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewWebhook) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetTitle

`func (o *NewWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewWebhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewWebhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewWebhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewWebhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDraft

`func (o *NewWebhook) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *NewWebhook) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *NewWebhook) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetVerb

`func (o *NewWebhook) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *NewWebhook) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *NewWebhook) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetUrl

`func (o *NewWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *NewWebhook) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *NewWebhook) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *NewWebhook) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetPayload

`func (o *NewWebhook) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NewWebhook) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NewWebhook) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NewWebhook) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetParams

`func (o *NewWebhook) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NewWebhook) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NewWebhook) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.


### GetEnabled

`func (o *NewWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuthenticationId

`func (o *NewWebhook) GetAuthenticationId() int64`

GetAuthenticationId returns the AuthenticationId field if non-nil, zero value otherwise.

### GetAuthenticationIdOk

`func (o *NewWebhook) GetAuthenticationIdOk() (*int64, bool)`

GetAuthenticationIdOk returns a tuple with the AuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationId

`func (o *NewWebhook) SetAuthenticationId(v int64)`

SetAuthenticationId sets AuthenticationId field to given value.

### HasAuthenticationId

`func (o *NewWebhook) HasAuthenticationId() bool`

HasAuthenticationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


