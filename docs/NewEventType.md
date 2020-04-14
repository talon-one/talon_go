# NewEventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | Pointer to **[]int32** | The IDs of the applications that are related to this entity. | 
**Title** | Pointer to **string** | The human-friendly display name for this event type. Use a short, past-tense, description of the event. | 
**Name** | Pointer to **string** | The machine-friendly canonical name for this event type. This will be used in URLs, and cannot be changed after an event type has been created. | 
**Description** | Pointer to **string** | An explanation of when the event type is triggered. Write this with a campaign manager in mind. For example:  &gt; The \&quot;Payment Accepted\&quot; event is triggered after successful processing of a payment by our payment gateway.  | 
**MimeType** | Pointer to **string** | This defines how the request payload will be parsed before your handler code is run. | 
**ExamplePayload** | Pointer to **string** | It is often helpful to include an example payload with the event type definition for documentation purposes. | [optional] 
**Schema** | Pointer to [**map[string]interface{}**](.md) | It is strongly recommended to define a JSON schema that will be used to perform structural validation of request payloads after parsing.  | [optional] 
**HandlerLanguage** | Pointer to **string** | The language of the handler code. Currently only &#x60;\&quot;talang\&quot;&#x60; is supported. | [optional] 
**Handler** | Pointer to **string** | Code that will be run after successful parsing &amp; validation of the payload for this event. This code _may_ choose to evaluate campaign rules.  | 
**Version** | Pointer to **int32** | The version of this event type. When updating an existing event type this must be **exactly** &#x60;currentVersion + 1&#x60;.  | 

## Methods

### GetApplicationIds

`func (o *NewEventType) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewEventType) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *NewEventType) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *NewEventType) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetTitle

`func (o *NewEventType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewEventType) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewEventType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewEventType) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetName

`func (o *NewEventType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewEventType) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewEventType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewEventType) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *NewEventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewEventType) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewEventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewEventType) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMimeType

`func (o *NewEventType) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *NewEventType) GetMimeTypeOk() (string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMimeType

`func (o *NewEventType) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeType

`func (o *NewEventType) SetMimeType(v string)`

SetMimeType gets a reference to the given string and assigns it to the MimeType field.

### GetExamplePayload

`func (o *NewEventType) GetExamplePayload() string`

GetExamplePayload returns the ExamplePayload field if non-nil, zero value otherwise.

### GetExamplePayloadOk

`func (o *NewEventType) GetExamplePayloadOk() (string, bool)`

GetExamplePayloadOk returns a tuple with the ExamplePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExamplePayload

`func (o *NewEventType) HasExamplePayload() bool`

HasExamplePayload returns a boolean if a field has been set.

### SetExamplePayload

`func (o *NewEventType) SetExamplePayload(v string)`

SetExamplePayload gets a reference to the given string and assigns it to the ExamplePayload field.

### GetSchema

`func (o *NewEventType) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *NewEventType) GetSchemaOk() (map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchema

`func (o *NewEventType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchema

`func (o *NewEventType) SetSchema(v map[string]interface{})`

SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.

### GetHandlerLanguage

`func (o *NewEventType) GetHandlerLanguage() string`

GetHandlerLanguage returns the HandlerLanguage field if non-nil, zero value otherwise.

### GetHandlerLanguageOk

`func (o *NewEventType) GetHandlerLanguageOk() (string, bool)`

GetHandlerLanguageOk returns a tuple with the HandlerLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandlerLanguage

`func (o *NewEventType) HasHandlerLanguage() bool`

HasHandlerLanguage returns a boolean if a field has been set.

### SetHandlerLanguage

`func (o *NewEventType) SetHandlerLanguage(v string)`

SetHandlerLanguage gets a reference to the given string and assigns it to the HandlerLanguage field.

### GetHandler

`func (o *NewEventType) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *NewEventType) GetHandlerOk() (string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandler

`func (o *NewEventType) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### SetHandler

`func (o *NewEventType) SetHandler(v string)`

SetHandler gets a reference to the given string and assigns it to the Handler field.

### GetVersion

`func (o *NewEventType) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NewEventType) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *NewEventType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *NewEventType) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


