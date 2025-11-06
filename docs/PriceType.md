# PriceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this price type. | [optional] 
**Name** | Pointer to **string** | The API name of the price type. This is an immutable value. | 
**Title** | Pointer to **string** | The name displayed in the Campaign Manager for the price type. | 
**Description** | Pointer to **string** | A description of the price type. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The date and time when the price type was last modified. | 
**SubscribedCatalogsIds** | Pointer to **[]int64** | A list of the IDs of the catalogs that are subscribed to this price type. | 
**TargetedAudiencesIds** | Pointer to **[]int64** | A list of the IDs of the audiences targeted by this price type. | 

## Methods

### NewPriceType

`func NewPriceType(id int64, created time.Time, name string, title string, modified time.Time, subscribedCatalogsIds []int64, targetedAudiencesIds []int64, ) *PriceType`

NewPriceType instantiates a new PriceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTypeWithDefaults

`func NewPriceTypeWithDefaults() *PriceType`

NewPriceTypeWithDefaults instantiates a new PriceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceType) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *PriceType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PriceType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PriceType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *PriceType) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PriceType) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PriceType) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PriceType) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *PriceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceType) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *PriceType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PriceType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PriceType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *PriceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PriceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PriceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PriceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModified

`func (o *PriceType) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PriceType) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PriceType) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetSubscribedCatalogsIds

`func (o *PriceType) GetSubscribedCatalogsIds() []int64`

GetSubscribedCatalogsIds returns the SubscribedCatalogsIds field if non-nil, zero value otherwise.

### GetSubscribedCatalogsIdsOk

`func (o *PriceType) GetSubscribedCatalogsIdsOk() (*[]int64, bool)`

GetSubscribedCatalogsIdsOk returns a tuple with the SubscribedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedCatalogsIds

`func (o *PriceType) SetSubscribedCatalogsIds(v []int64)`

SetSubscribedCatalogsIds sets SubscribedCatalogsIds field to given value.


### GetTargetedAudiencesIds

`func (o *PriceType) GetTargetedAudiencesIds() []int64`

GetTargetedAudiencesIds returns the TargetedAudiencesIds field if non-nil, zero value otherwise.

### GetTargetedAudiencesIdsOk

`func (o *PriceType) GetTargetedAudiencesIdsOk() (*[]int64, bool)`

GetTargetedAudiencesIdsOk returns a tuple with the TargetedAudiencesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedAudiencesIds

`func (o *PriceType) SetTargetedAudiencesIds(v []int64)`

SetTargetedAudiencesIds sets TargetedAudiencesIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


