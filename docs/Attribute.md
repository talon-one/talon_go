# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Entity** | Pointer to **string** | The name of the entity that can have this attribute. When creating or updating the entities of a given type, you can include an &#x60;attributes&#x60; object with keys corresponding to the &#x60;name&#x60; of the custom attributes for that type. | 
**EventType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The attribute name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;region\&quot;&#x60; then you would set the region attribute by including an &#x60;attributes.region&#x60; property in your request payload. | 
**Title** | Pointer to **string** | The human-readable name for the attribute that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Type** | Pointer to **string** | The data type of the attribute, a &#x60;time&#x60; attribute must be sent as a string that conforms to the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp format. | 
**Description** | Pointer to **string** | A description of this attribute. | 
**Suggestions** | Pointer to **[]string** | A list of suggestions for the attribute. | 
**Editable** | Pointer to **bool** | Whether or not this attribute can be edited. | 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications that are subscribed to this attribute | [optional] 

## Methods

### GetId

`func (o *Attribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Attribute) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Attribute) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Attribute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Attribute) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Attribute) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Attribute) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetAccountId

`func (o *Attribute) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Attribute) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Attribute) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Attribute) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetEntity

`func (o *Attribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Attribute) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *Attribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *Attribute) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetEventType

`func (o *Attribute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Attribute) GetEventTypeOk() (string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventType

`func (o *Attribute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventType

`func (o *Attribute) SetEventType(v string)`

SetEventType gets a reference to the given string and assigns it to the EventType field.

### GetName

`func (o *Attribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attribute) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Attribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Attribute) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *Attribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Attribute) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Attribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Attribute) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetType

`func (o *Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Attribute) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Attribute) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *Attribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Attribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Attribute) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSuggestions

`func (o *Attribute) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *Attribute) GetSuggestionsOk() ([]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuggestions

`func (o *Attribute) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### SetSuggestions

`func (o *Attribute) SetSuggestions(v []string)`

SetSuggestions gets a reference to the given []string and assigns it to the Suggestions field.

### GetEditable

`func (o *Attribute) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Attribute) GetEditableOk() (bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEditable

`func (o *Attribute) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditable

`func (o *Attribute) SetEditable(v bool)`

SetEditable gets a reference to the given bool and assigns it to the Editable field.

### GetSubscribedApplicationsIds

`func (o *Attribute) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *Attribute) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *Attribute) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *Attribute) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


