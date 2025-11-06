# MultipleAudiences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Audiences** | Pointer to [**[]MultipleAudiencesItem**](MultipleAudiencesItem.md) |  | 

## Methods

### NewMultipleAudiences

`func NewMultipleAudiences(accountId int64, audiences []MultipleAudiencesItem, ) *MultipleAudiences`

NewMultipleAudiences instantiates a new MultipleAudiences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleAudiencesWithDefaults

`func NewMultipleAudiencesWithDefaults() *MultipleAudiences`

NewMultipleAudiencesWithDefaults instantiates a new MultipleAudiences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MultipleAudiences) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MultipleAudiences) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MultipleAudiences) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAudiences

`func (o *MultipleAudiences) GetAudiences() []MultipleAudiencesItem`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *MultipleAudiences) GetAudiencesOk() (*[]MultipleAudiencesItem, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *MultipleAudiences) SetAudiences(v []MultipleAudiencesItem)`

SetAudiences sets Audiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


