# ScimSchemaResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Description** | Pointer to **string** | Human-readable description of the schema resource. | [optional] 
**Attributes** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewScimSchemaResource

`func NewScimSchemaResource() *ScimSchemaResource`

NewScimSchemaResource instantiates a new ScimSchemaResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSchemaResourceWithDefaults

`func NewScimSchemaResourceWithDefaults() *ScimSchemaResource`

NewScimSchemaResourceWithDefaults instantiates a new ScimSchemaResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimSchemaResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimSchemaResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimSchemaResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScimSchemaResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScimSchemaResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimSchemaResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimSchemaResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScimSchemaResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ScimSchemaResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimSchemaResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimSchemaResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimSchemaResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributes

`func (o *ScimSchemaResource) GetAttributes() []map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScimSchemaResource) GetAttributesOk() (*[]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScimSchemaResource) SetAttributes(v []map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScimSchemaResource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


