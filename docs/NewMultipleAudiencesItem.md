# NewMultipleAudiencesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The human-friendly display name for this audience. | 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the Applications that are connected to this audience. | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration. | [optional] 

## Methods

### NewNewMultipleAudiencesItem

`func NewNewMultipleAudiencesItem(name string, ) *NewMultipleAudiencesItem`

NewNewMultipleAudiencesItem instantiates a new NewMultipleAudiencesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMultipleAudiencesItemWithDefaults

`func NewNewMultipleAudiencesItemWithDefaults() *NewMultipleAudiencesItem`

NewNewMultipleAudiencesItemWithDefaults instantiates a new NewMultipleAudiencesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewMultipleAudiencesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewMultipleAudiencesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewMultipleAudiencesItem) SetName(v string)`

SetName sets Name field to given value.


### GetSubscribedApplicationsIds

`func (o *NewMultipleAudiencesItem) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewMultipleAudiencesItem) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *NewMultipleAudiencesItem) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *NewMultipleAudiencesItem) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetIntegrationId

`func (o *NewMultipleAudiencesItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewMultipleAudiencesItem) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewMultipleAudiencesItem) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *NewMultipleAudiencesItem) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


