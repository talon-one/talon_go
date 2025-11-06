# OutgoingIntegrationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**Name** | Pointer to **string** | Name of the outgoing integration. | 
**Description** | Pointer to **string** | Description of the outgoing integration. | [optional] 
**Category** | Pointer to **string** | Category of the outgoing integration. | [optional] 
**DocumentationLink** | Pointer to **string** | Http link to the outgoing integration&#39;s documentation. | [optional] 

## Methods

### NewOutgoingIntegrationType

`func NewOutgoingIntegrationType(id int64, name string, ) *OutgoingIntegrationType`

NewOutgoingIntegrationType instantiates a new OutgoingIntegrationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationTypeWithDefaults

`func NewOutgoingIntegrationTypeWithDefaults() *OutgoingIntegrationType`

NewOutgoingIntegrationTypeWithDefaults instantiates a new OutgoingIntegrationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutgoingIntegrationType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingIntegrationType) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *OutgoingIntegrationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutgoingIntegrationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutgoingIntegrationType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OutgoingIntegrationType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingIntegrationType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingIntegrationType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutgoingIntegrationType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *OutgoingIntegrationType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *OutgoingIntegrationType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *OutgoingIntegrationType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *OutgoingIntegrationType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *OutgoingIntegrationType) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *OutgoingIntegrationType) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *OutgoingIntegrationType) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *OutgoingIntegrationType) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


