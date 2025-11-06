# UpdateCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of this cart item catalog. | [optional] 
**Name** | Pointer to **string** | Name of this cart item catalog. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the applications that are subscribed to this catalog. | [optional] 

## Methods

### NewUpdateCatalog

`func NewUpdateCatalog() *UpdateCatalog`

NewUpdateCatalog instantiates a new UpdateCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogWithDefaults

`func NewUpdateCatalogWithDefaults() *UpdateCatalog`

NewUpdateCatalogWithDefaults instantiates a new UpdateCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *UpdateCatalog) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *UpdateCatalog) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *UpdateCatalog) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *UpdateCatalog) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


