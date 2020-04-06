# LibraryAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The name of the entity that can have this attribute. When creating or updating the entities of a given type, you can include an &#x60;attributes&#x60; object with keys corresponding to the &#x60;name&#x60; of the custom attributes for that type. | 
**Name** | Pointer to **string** | The attribute name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;region\&quot;&#x60; then you would set the region attribute by including an &#x60;attributes.region&#x60; property in your request payload.  | 
**Title** | Pointer to **string** | The human-readable name for the attribute that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Type** | Pointer to **string** | The data type of the attribute, a &#x60;time&#x60; attribute must be sent as a string that conforms to the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp format. | 
**Description** | Pointer to **string** | A description of the attribute. | 
**Presets** | Pointer to **[]string** | The presets that indicate to which industry the attribute applies to. | 
**Suggestions** | Pointer to **[]string** | Short suggestions that are used to group attributes. | 

## Methods

### GetEntity

`func (o *LibraryAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *LibraryAttribute) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *LibraryAttribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *LibraryAttribute) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetName

`func (o *LibraryAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryAttribute) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LibraryAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LibraryAttribute) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *LibraryAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LibraryAttribute) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *LibraryAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *LibraryAttribute) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetType

`func (o *LibraryAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LibraryAttribute) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LibraryAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LibraryAttribute) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *LibraryAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LibraryAttribute) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *LibraryAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *LibraryAttribute) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetPresets

`func (o *LibraryAttribute) GetPresets() []string`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *LibraryAttribute) GetPresetsOk() ([]string, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPresets

`func (o *LibraryAttribute) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresets

`func (o *LibraryAttribute) SetPresets(v []string)`

SetPresets gets a reference to the given []string and assigns it to the Presets field.

### GetSuggestions

`func (o *LibraryAttribute) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *LibraryAttribute) GetSuggestionsOk() ([]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuggestions

`func (o *LibraryAttribute) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### SetSuggestions

`func (o *LibraryAttribute) SetSuggestions(v []string)`

SetSuggestions gets a reference to the given []string and assigns it to the Suggestions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


