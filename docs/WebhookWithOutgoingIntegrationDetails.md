# WebhookWithOutgoingIntegrationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
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
**OutgoingIntegrationTemplateId** | Pointer to **int64** | Identifier of the outgoing integration template. | [optional] 
**OutgoingIntegrationTypeId** | Pointer to **int64** | Identifier of the outgoing integration type. | [optional] 
**OutgoingIntegrationTypeName** | Pointer to **string** | Name of the outgoing integration. | [optional] 

## Methods

### NewWebhookWithOutgoingIntegrationDetails

`func NewWebhookWithOutgoingIntegrationDetails(id int64, created time.Time, modified time.Time, applicationIds []int64, title string, draft bool, verb string, url string, headers []string, params []TemplateArgDef, enabled bool, ) *WebhookWithOutgoingIntegrationDetails`

NewWebhookWithOutgoingIntegrationDetails instantiates a new WebhookWithOutgoingIntegrationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithOutgoingIntegrationDetailsWithDefaults

`func NewWebhookWithOutgoingIntegrationDetailsWithDefaults() *WebhookWithOutgoingIntegrationDetails`

NewWebhookWithOutgoingIntegrationDetailsWithDefaults instantiates a new WebhookWithOutgoingIntegrationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookWithOutgoingIntegrationDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookWithOutgoingIntegrationDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *WebhookWithOutgoingIntegrationDetails) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookWithOutgoingIntegrationDetails) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *WebhookWithOutgoingIntegrationDetails) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WebhookWithOutgoingIntegrationDetails) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetApplicationIds

`func (o *WebhookWithOutgoingIntegrationDetails) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *WebhookWithOutgoingIntegrationDetails) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetTitle

`func (o *WebhookWithOutgoingIntegrationDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookWithOutgoingIntegrationDetails) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *WebhookWithOutgoingIntegrationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookWithOutgoingIntegrationDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookWithOutgoingIntegrationDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDraft

`func (o *WebhookWithOutgoingIntegrationDetails) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *WebhookWithOutgoingIntegrationDetails) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetVerb

`func (o *WebhookWithOutgoingIntegrationDetails) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *WebhookWithOutgoingIntegrationDetails) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetUrl

`func (o *WebhookWithOutgoingIntegrationDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookWithOutgoingIntegrationDetails) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *WebhookWithOutgoingIntegrationDetails) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookWithOutgoingIntegrationDetails) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.


### GetPayload

`func (o *WebhookWithOutgoingIntegrationDetails) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookWithOutgoingIntegrationDetails) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookWithOutgoingIntegrationDetails) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetParams

`func (o *WebhookWithOutgoingIntegrationDetails) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *WebhookWithOutgoingIntegrationDetails) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.


### GetEnabled

`func (o *WebhookWithOutgoingIntegrationDetails) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookWithOutgoingIntegrationDetails) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuthenticationId

`func (o *WebhookWithOutgoingIntegrationDetails) GetAuthenticationId() int64`

GetAuthenticationId returns the AuthenticationId field if non-nil, zero value otherwise.

### GetAuthenticationIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetAuthenticationIdOk() (*int64, bool)`

GetAuthenticationIdOk returns a tuple with the AuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationId

`func (o *WebhookWithOutgoingIntegrationDetails) SetAuthenticationId(v int64)`

SetAuthenticationId sets AuthenticationId field to given value.

### HasAuthenticationId

`func (o *WebhookWithOutgoingIntegrationDetails) HasAuthenticationId() bool`

HasAuthenticationId returns a boolean if a field has been set.

### GetOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateId() int64`

GetOutgoingIntegrationTemplateId returns the OutgoingIntegrationTemplateId field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTemplateIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateIdOk() (*int64, bool)`

GetOutgoingIntegrationTemplateIdOk returns a tuple with the OutgoingIntegrationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTemplateId(v int64)`

SetOutgoingIntegrationTemplateId sets OutgoingIntegrationTemplateId field to given value.

### HasOutgoingIntegrationTemplateId

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTemplateId() bool`

HasOutgoingIntegrationTemplateId returns a boolean if a field has been set.

### GetOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeId() int64`

GetOutgoingIntegrationTypeId returns the OutgoingIntegrationTypeId field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTypeIdOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeIdOk() (*int64, bool)`

GetOutgoingIntegrationTypeIdOk returns a tuple with the OutgoingIntegrationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeId(v int64)`

SetOutgoingIntegrationTypeId sets OutgoingIntegrationTypeId field to given value.

### HasOutgoingIntegrationTypeId

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeId() bool`

HasOutgoingIntegrationTypeId returns a boolean if a field has been set.

### GetOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeName() string`

GetOutgoingIntegrationTypeName returns the OutgoingIntegrationTypeName field if non-nil, zero value otherwise.

### GetOutgoingIntegrationTypeNameOk

`func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeNameOk() (*string, bool)`

GetOutgoingIntegrationTypeNameOk returns a tuple with the OutgoingIntegrationTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeName(v string)`

SetOutgoingIntegrationTypeName sets OutgoingIntegrationTypeName field to given value.

### HasOutgoingIntegrationTypeName

`func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeName() bool`

HasOutgoingIntegrationTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


