# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
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
**EventTypeId** | Pointer to **int64** |  | [optional] 

## Methods

### NewAttribute

`func NewAttribute(id int64, created time.Time, accountId int64, entity string, name string, title string, type_ string, description string, suggestions []string, editable bool, ) *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attribute) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribute) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Attribute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Attribute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Attribute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *Attribute) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Attribute) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Attribute) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetEntity

`func (o *Attribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Attribute) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Attribute) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEventType

`func (o *Attribute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Attribute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Attribute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Attribute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetName

`func (o *Attribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Attribute) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Attribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Attribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Attribute) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attribute) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *Attribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attribute) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSuggestions

`func (o *Attribute) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *Attribute) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *Attribute) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.


### GetHasAllowedList

`func (o *Attribute) GetHasAllowedList() bool`

GetHasAllowedList returns the HasAllowedList field if non-nil, zero value otherwise.

### GetHasAllowedListOk

`func (o *Attribute) GetHasAllowedListOk() (*bool, bool)`

GetHasAllowedListOk returns a tuple with the HasAllowedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAllowedList

`func (o *Attribute) SetHasAllowedList(v bool)`

SetHasAllowedList sets HasAllowedList field to given value.

### HasHasAllowedList

`func (o *Attribute) HasHasAllowedList() bool`

HasHasAllowedList returns a boolean if a field has been set.

### GetRestrictedBySuggestions

`func (o *Attribute) GetRestrictedBySuggestions() bool`

GetRestrictedBySuggestions returns the RestrictedBySuggestions field if non-nil, zero value otherwise.

### GetRestrictedBySuggestionsOk

`func (o *Attribute) GetRestrictedBySuggestionsOk() (*bool, bool)`

GetRestrictedBySuggestionsOk returns a tuple with the RestrictedBySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedBySuggestions

`func (o *Attribute) SetRestrictedBySuggestions(v bool)`

SetRestrictedBySuggestions sets RestrictedBySuggestions field to given value.

### HasRestrictedBySuggestions

`func (o *Attribute) HasRestrictedBySuggestions() bool`

HasRestrictedBySuggestions returns a boolean if a field has been set.

### GetEditable

`func (o *Attribute) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Attribute) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Attribute) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetSubscribedApplicationsIds

`func (o *Attribute) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *Attribute) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *Attribute) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *Attribute) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetSubscribedCatalogsIds

`func (o *Attribute) GetSubscribedCatalogsIds() []int64`

GetSubscribedCatalogsIds returns the SubscribedCatalogsIds field if non-nil, zero value otherwise.

### GetSubscribedCatalogsIdsOk

`func (o *Attribute) GetSubscribedCatalogsIdsOk() (*[]int64, bool)`

GetSubscribedCatalogsIdsOk returns a tuple with the SubscribedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedCatalogsIds

`func (o *Attribute) SetSubscribedCatalogsIds(v []int64)`

SetSubscribedCatalogsIds sets SubscribedCatalogsIds field to given value.

### HasSubscribedCatalogsIds

`func (o *Attribute) HasSubscribedCatalogsIds() bool`

HasSubscribedCatalogsIds returns a boolean if a field has been set.

### GetAllowedSubscriptions

`func (o *Attribute) GetAllowedSubscriptions() []string`

GetAllowedSubscriptions returns the AllowedSubscriptions field if non-nil, zero value otherwise.

### GetAllowedSubscriptionsOk

`func (o *Attribute) GetAllowedSubscriptionsOk() (*[]string, bool)`

GetAllowedSubscriptionsOk returns a tuple with the AllowedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubscriptions

`func (o *Attribute) SetAllowedSubscriptions(v []string)`

SetAllowedSubscriptions sets AllowedSubscriptions field to given value.

### HasAllowedSubscriptions

`func (o *Attribute) HasAllowedSubscriptions() bool`

HasAllowedSubscriptions returns a boolean if a field has been set.

### GetEventTypeId

`func (o *Attribute) GetEventTypeId() int64`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *Attribute) GetEventTypeIdOk() (*int64, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *Attribute) SetEventTypeId(v int64)`

SetEventTypeId sets EventTypeId field to given value.

### HasEventTypeId

`func (o *Attribute) HasEventTypeId() bool`

HasEventTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


