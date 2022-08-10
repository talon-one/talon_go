# NewAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The name of the entity that can have this attribute. When creating or updating the entities of a given type, you can include an &#x60;attributes&#x60; object with keys corresponding to the &#x60;name&#x60; of the custom attributes for that type. | 
**EventType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The attribute name that will be used in API requests and Talang. E.g. if &#x60;name &#x3D;&#x3D; \&quot;region\&quot;&#x60; then you would set the region attribute by including an &#x60;attributes.region&#x60; property in your request payload. | 
**Title** | Pointer to **string** | The human-readable name for the attribute that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Type** | Pointer to **string** | The data type of the attribute, a &#x60;time&#x60; attribute must be sent as a string that conforms to the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp format. | 
**Description** | Pointer to **string** | A description of this attribute. | 
**Suggestions** | Pointer to **[]string** | A list of suggestions for the attribute. | 
**HasAllowedList** | Pointer to **bool** | Whether or not this attribute has an allowed list of values associated with it. | [optional] [default to false]
**RestrictedBySuggestions** | Pointer to **bool** | Whether or not this attribute&#39;s value is restricted by suggestions (&#x60;suggestions&#x60; property) or by an allowed list of value (&#x60;hasAllowedList&#x60; property).  | [optional] [default to false]
**Editable** | Pointer to **bool** | Whether or not this attribute can be edited. | 
**SubscribedApplicationsIds** | Pointer to **[]int32** | A list of the IDs of the applications where this attribute is available. | [optional] 
**SubscribedCatalogsIds** | Pointer to **[]int32** | A list of the IDs of the catalogs where this attribute is available. | [optional] 
**AllowedSubscriptions** | Pointer to **[]string** | A list of allowed subscription types for this attribute.  **Note:** This only applies to attributes associated with the &#x60;CartItem&#x60; entity.  | [optional] 

## Methods

### GetEntity

`func (o *NewAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *NewAttribute) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *NewAttribute) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *NewAttribute) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetEventType

`func (o *NewAttribute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NewAttribute) GetEventTypeOk() (string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventType

`func (o *NewAttribute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventType

`func (o *NewAttribute) SetEventType(v string)`

SetEventType gets a reference to the given string and assigns it to the EventType field.

### GetName

`func (o *NewAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAttribute) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewAttribute) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *NewAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAttribute) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewAttribute) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetType

`func (o *NewAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewAttribute) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *NewAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *NewAttribute) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *NewAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewAttribute) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewAttribute) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSuggestions

`func (o *NewAttribute) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *NewAttribute) GetSuggestionsOk() ([]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuggestions

`func (o *NewAttribute) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### SetSuggestions

`func (o *NewAttribute) SetSuggestions(v []string)`

SetSuggestions gets a reference to the given []string and assigns it to the Suggestions field.

### GetHasAllowedList

`func (o *NewAttribute) GetHasAllowedList() bool`

GetHasAllowedList returns the HasAllowedList field if non-nil, zero value otherwise.

### GetHasAllowedListOk

`func (o *NewAttribute) GetHasAllowedListOk() (bool, bool)`

GetHasAllowedListOk returns a tuple with the HasAllowedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAllowedList

`func (o *NewAttribute) HasHasAllowedList() bool`

HasHasAllowedList returns a boolean if a field has been set.

### SetHasAllowedList

`func (o *NewAttribute) SetHasAllowedList(v bool)`

SetHasAllowedList gets a reference to the given bool and assigns it to the HasAllowedList field.

### GetRestrictedBySuggestions

`func (o *NewAttribute) GetRestrictedBySuggestions() bool`

GetRestrictedBySuggestions returns the RestrictedBySuggestions field if non-nil, zero value otherwise.

### GetRestrictedBySuggestionsOk

`func (o *NewAttribute) GetRestrictedBySuggestionsOk() (bool, bool)`

GetRestrictedBySuggestionsOk returns a tuple with the RestrictedBySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRestrictedBySuggestions

`func (o *NewAttribute) HasRestrictedBySuggestions() bool`

HasRestrictedBySuggestions returns a boolean if a field has been set.

### SetRestrictedBySuggestions

`func (o *NewAttribute) SetRestrictedBySuggestions(v bool)`

SetRestrictedBySuggestions gets a reference to the given bool and assigns it to the RestrictedBySuggestions field.

### GetEditable

`func (o *NewAttribute) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NewAttribute) GetEditableOk() (bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEditable

`func (o *NewAttribute) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditable

`func (o *NewAttribute) SetEditable(v bool)`

SetEditable gets a reference to the given bool and assigns it to the Editable field.

### GetSubscribedApplicationsIds

`func (o *NewAttribute) GetSubscribedApplicationsIds() []int32`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewAttribute) GetSubscribedApplicationsIdsOk() ([]int32, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplicationsIds

`func (o *NewAttribute) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### SetSubscribedApplicationsIds

`func (o *NewAttribute) SetSubscribedApplicationsIds(v []int32)`

SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.

### GetSubscribedCatalogsIds

`func (o *NewAttribute) GetSubscribedCatalogsIds() []int32`

GetSubscribedCatalogsIds returns the SubscribedCatalogsIds field if non-nil, zero value otherwise.

### GetSubscribedCatalogsIdsOk

`func (o *NewAttribute) GetSubscribedCatalogsIdsOk() ([]int32, bool)`

GetSubscribedCatalogsIdsOk returns a tuple with the SubscribedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedCatalogsIds

`func (o *NewAttribute) HasSubscribedCatalogsIds() bool`

HasSubscribedCatalogsIds returns a boolean if a field has been set.

### SetSubscribedCatalogsIds

`func (o *NewAttribute) SetSubscribedCatalogsIds(v []int32)`

SetSubscribedCatalogsIds gets a reference to the given []int32 and assigns it to the SubscribedCatalogsIds field.

### GetAllowedSubscriptions

`func (o *NewAttribute) GetAllowedSubscriptions() []string`

GetAllowedSubscriptions returns the AllowedSubscriptions field if non-nil, zero value otherwise.

### GetAllowedSubscriptionsOk

`func (o *NewAttribute) GetAllowedSubscriptionsOk() ([]string, bool)`

GetAllowedSubscriptionsOk returns a tuple with the AllowedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedSubscriptions

`func (o *NewAttribute) HasAllowedSubscriptions() bool`

HasAllowedSubscriptions returns a boolean if a field has been set.

### SetAllowedSubscriptions

`func (o *NewAttribute) SetAllowedSubscriptions(v []string)`

SetAllowedSubscriptions gets a reference to the given []string and assigns it to the AllowedSubscriptions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


