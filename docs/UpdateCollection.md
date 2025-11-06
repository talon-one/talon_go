# UpdateCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the purpose of this collection. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the Applications where this collection is enabled. | [optional] 

## Methods

### NewUpdateCollection

`func NewUpdateCollection() *UpdateCollection`

NewUpdateCollection instantiates a new UpdateCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionWithDefaults

`func NewUpdateCollectionWithDefaults() *UpdateCollection`

NewUpdateCollectionWithDefaults instantiates a new UpdateCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *UpdateCollection) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *UpdateCollection) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *UpdateCollection) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *UpdateCollection) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


