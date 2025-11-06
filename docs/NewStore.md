# NewStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the store. | 
**Description** | Pointer to **string** | The description of the store. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the store. | [optional] 
**IntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store.  **Note**: You cannot edit the &#x60;integrationId&#x60; after the store has been created.  | 

## Methods

### NewNewStore

`func NewNewStore(name string, description string, integrationId string, ) *NewStore`

NewNewStore instantiates a new NewStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewStoreWithDefaults

`func NewNewStoreWithDefaults() *NewStore`

NewNewStoreWithDefaults instantiates a new NewStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewStore) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewStore) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAttributes

`func (o *NewStore) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewStore) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewStore) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewStore) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIntegrationId

`func (o *NewStore) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewStore) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewStore) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


