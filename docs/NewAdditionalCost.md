# NewAdditionalCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The internal name used in API requests. | 
**Title** | Pointer to **string** | The human-readable name for the additional cost that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Description** | Pointer to **string** | A description of this additional cost. | 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the applications that are subscribed to this additional cost. | [optional] 
**Type** | Pointer to **string** | The type of additional cost. Possible value: - &#x60;session&#x60;: Additional cost will be added per session. - &#x60;item&#x60;: Additional cost will be added per item. - &#x60;both&#x60;: Additional cost will be added per item and session.  | [optional] [default to "session"]

## Methods

### NewNewAdditionalCost

`func NewNewAdditionalCost(name string, title string, description string, ) *NewAdditionalCost`

NewNewAdditionalCost instantiates a new NewAdditionalCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAdditionalCostWithDefaults

`func NewNewAdditionalCostWithDefaults() *NewAdditionalCost`

NewNewAdditionalCostWithDefaults instantiates a new NewAdditionalCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewAdditionalCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewAdditionalCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewAdditionalCost) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewAdditionalCost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewAdditionalCost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewAdditionalCost) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewAdditionalCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewAdditionalCost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewAdditionalCost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubscribedApplicationsIds

`func (o *NewAdditionalCost) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewAdditionalCost) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *NewAdditionalCost) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *NewAdditionalCost) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetType

`func (o *NewAdditionalCost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewAdditionalCost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewAdditionalCost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewAdditionalCost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


