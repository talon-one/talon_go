# EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
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

### GetId

`func (o *EventType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventType) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *EventType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *EventType) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *EventType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventType) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *EventType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *EventType) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationIds

`func (o *EventType) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *EventType) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *EventType) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *EventType) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetTitle

`func (o *EventType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventType) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *EventType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *EventType) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetName

`func (o *EventType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventType) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *EventType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *EventType) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *EventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventType) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *EventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *EventType) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMimeType

`func (o *EventType) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *EventType) GetMimeTypeOk() (string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMimeType

`func (o *EventType) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeType

`func (o *EventType) SetMimeType(v string)`

SetMimeType gets a reference to the given string and assigns it to the MimeType field.

### GetExamplePayload

`func (o *EventType) GetExamplePayload() string`

GetExamplePayload returns the ExamplePayload field if non-nil, zero value otherwise.

### GetExamplePayloadOk

`func (o *EventType) GetExamplePayloadOk() (string, bool)`

GetExamplePayloadOk returns a tuple with the ExamplePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExamplePayload

`func (o *EventType) HasExamplePayload() bool`

HasExamplePayload returns a boolean if a field has been set.

### SetExamplePayload

`func (o *EventType) SetExamplePayload(v string)`

SetExamplePayload gets a reference to the given string and assigns it to the ExamplePayload field.

### GetSchema

`func (o *EventType) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *EventType) GetSchemaOk() (map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchema

`func (o *EventType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchema

`func (o *EventType) SetSchema(v map[string]interface{})`

SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.

### GetHandlerLanguage

`func (o *EventType) GetHandlerLanguage() string`

GetHandlerLanguage returns the HandlerLanguage field if non-nil, zero value otherwise.

### GetHandlerLanguageOk

`func (o *EventType) GetHandlerLanguageOk() (string, bool)`

GetHandlerLanguageOk returns a tuple with the HandlerLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandlerLanguage

`func (o *EventType) HasHandlerLanguage() bool`

HasHandlerLanguage returns a boolean if a field has been set.

### SetHandlerLanguage

`func (o *EventType) SetHandlerLanguage(v string)`

SetHandlerLanguage gets a reference to the given string and assigns it to the HandlerLanguage field.

### GetHandler

`func (o *EventType) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *EventType) GetHandlerOk() (string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandler

`func (o *EventType) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### SetHandler

`func (o *EventType) SetHandler(v string)`

SetHandler gets a reference to the given string and assigns it to the Handler field.

### GetVersion

`func (o *EventType) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventType) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *EventType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *EventType) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


