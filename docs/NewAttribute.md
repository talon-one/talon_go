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
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the applications where this attribute is available. | [optional] 
**SubscribedCatalogsIds** | Pointer to **[]int64** | A list of the IDs of the catalogs where this attribute is available. | [optional] 
**AllowedSubscriptions** | Pointer to **[]string** | A list of allowed subscription types for this attribute.  **Note:** This only applies to attributes associated with the &#x60;CartItem&#x60; entity.  | [optional] 

## Methods

### NewNewAttribute

`func NewNewAttribute(entity string, name string, title string, type_ string, description string, suggestions []string, editable bool, ) *NewAttribute`

NewNewAttribute instantiates a new NewAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAttributeWithDefaults

`func NewNewAttributeWithDefaults() *NewAttribute`

NewNewAttributeWithDefaults instantiates a new NewAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *NewAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *NewAttribute) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *NewAttribute) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEventType

`func (o *NewAttribute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NewAttribute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NewAttribute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *NewAttribute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *NewAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAttribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewAttribute) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *NewAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewAttribute) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *NewAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSuggestions

`func (o *NewAttribute) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *NewAttribute) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *NewAttribute) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.


### GetHasAllowedList

`func (o *NewAttribute) GetHasAllowedList() bool`

GetHasAllowedList returns the HasAllowedList field if non-nil, zero value otherwise.

### GetHasAllowedListOk

`func (o *NewAttribute) GetHasAllowedListOk() (*bool, bool)`

GetHasAllowedListOk returns a tuple with the HasAllowedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAllowedList

`func (o *NewAttribute) SetHasAllowedList(v bool)`

SetHasAllowedList sets HasAllowedList field to given value.

### HasHasAllowedList

`func (o *NewAttribute) HasHasAllowedList() bool`

HasHasAllowedList returns a boolean if a field has been set.

### GetRestrictedBySuggestions

`func (o *NewAttribute) GetRestrictedBySuggestions() bool`

GetRestrictedBySuggestions returns the RestrictedBySuggestions field if non-nil, zero value otherwise.

### GetRestrictedBySuggestionsOk

`func (o *NewAttribute) GetRestrictedBySuggestionsOk() (*bool, bool)`

GetRestrictedBySuggestionsOk returns a tuple with the RestrictedBySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBySuggestions

`func (o *NewAttribute) SetRestrictedBySuggestions(v bool)`

SetRestrictedBySuggestions sets RestrictedBySuggestions field to given value.

### HasRestrictedBySuggestions

`func (o *NewAttribute) HasRestrictedBySuggestions() bool`

HasRestrictedBySuggestions returns a boolean if a field has been set.

### GetEditable

`func (o *NewAttribute) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NewAttribute) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NewAttribute) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetSubscribedApplicationsIds

`func (o *NewAttribute) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewAttribute) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *NewAttribute) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *NewAttribute) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetSubscribedCatalogsIds

`func (o *NewAttribute) GetSubscribedCatalogsIds() []int64`

GetSubscribedCatalogsIds returns the SubscribedCatalogsIds field if non-nil, zero value otherwise.

### GetSubscribedCatalogsIdsOk

`func (o *NewAttribute) GetSubscribedCatalogsIdsOk() (*[]int64, bool)`

GetSubscribedCatalogsIdsOk returns a tuple with the SubscribedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedCatalogsIds

`func (o *NewAttribute) SetSubscribedCatalogsIds(v []int64)`

SetSubscribedCatalogsIds sets SubscribedCatalogsIds field to given value.

### HasSubscribedCatalogsIds

`func (o *NewAttribute) HasSubscribedCatalogsIds() bool`

HasSubscribedCatalogsIds returns a boolean if a field has been set.

### GetAllowedSubscriptions

`func (o *NewAttribute) GetAllowedSubscriptions() []string`

GetAllowedSubscriptions returns the AllowedSubscriptions field if non-nil, zero value otherwise.

### GetAllowedSubscriptionsOk

`func (o *NewAttribute) GetAllowedSubscriptionsOk() (*[]string, bool)`

GetAllowedSubscriptionsOk returns a tuple with the AllowedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubscriptions

`func (o *NewAttribute) SetAllowedSubscriptions(v []string)`

SetAllowedSubscriptions sets AllowedSubscriptions field to given value.

### HasAllowedSubscriptions

`func (o *NewAttribute) HasAllowedSubscriptions() bool`

HasAllowedSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


