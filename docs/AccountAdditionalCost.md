# AccountAdditionalCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | The internal name used in API requests. | 
**Title** | Pointer to **string** | The human-readable name for the additional cost that will be shown in the Campaign Manager. Like &#x60;name&#x60;, the combination of entity and title must also be unique. | 
**Description** | Pointer to **string** | A description of this additional cost. | 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the applications that are subscribed to this additional cost. | [optional] 
**Type** | Pointer to **string** | The type of additional cost. Possible value: - &#x60;session&#x60;: Additional cost will be added per session. - &#x60;item&#x60;: Additional cost will be added per item. - &#x60;both&#x60;: Additional cost will be added per item and session.  | [optional] [default to "session"]

## Methods

### NewAccountAdditionalCost

`func NewAccountAdditionalCost(id int64, created time.Time, accountId int64, name string, title string, description string, ) *AccountAdditionalCost`

NewAccountAdditionalCost instantiates a new AccountAdditionalCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAdditionalCostWithDefaults

`func NewAccountAdditionalCostWithDefaults() *AccountAdditionalCost`

NewAccountAdditionalCostWithDefaults instantiates a new AccountAdditionalCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountAdditionalCost) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountAdditionalCost) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountAdditionalCost) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *AccountAdditionalCost) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountAdditionalCost) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountAdditionalCost) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *AccountAdditionalCost) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountAdditionalCost) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountAdditionalCost) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *AccountAdditionalCost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountAdditionalCost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountAdditionalCost) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AccountAdditionalCost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountAdditionalCost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountAdditionalCost) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AccountAdditionalCost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountAdditionalCost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountAdditionalCost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubscribedApplicationsIds

`func (o *AccountAdditionalCost) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *AccountAdditionalCost) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *AccountAdditionalCost) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *AccountAdditionalCost) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetType

`func (o *AccountAdditionalCost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountAdditionalCost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountAdditionalCost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountAdditionalCost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


