# PriceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this price type. | [optional] 
**Name** | Pointer to **string** | The API name of the price type. This is an immutable value. | 
**Title** | Pointer to **string** | The title of the price type. | 
**Description** | Pointer to **string** | The description of the price type. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The date and time when the price type was last modified. | 
**SubscribedCatalogsIds** | Pointer to **[]int32** | A list of the IDs of the catalogs that are subscribed to this price type. | 
**TargetedAudiencesIds** | Pointer to **[]int32** | A list of the IDs of the audiences that are targeted by this price type. | 

## Methods

### GetId

`func (o *PriceType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceType) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *PriceType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *PriceType) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *PriceType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PriceType) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *PriceType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *PriceType) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetAccountId

`func (o *PriceType) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PriceType) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *PriceType) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *PriceType) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetName

`func (o *PriceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceType) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *PriceType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *PriceType) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *PriceType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PriceType) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *PriceType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *PriceType) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *PriceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PriceType) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *PriceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *PriceType) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetModified

`func (o *PriceType) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PriceType) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *PriceType) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *PriceType) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetSubscribedCatalogsIds

`func (o *PriceType) GetSubscribedCatalogsIds() []int32`

GetSubscribedCatalogsIds returns the SubscribedCatalogsIds field if non-nil, zero value otherwise.

### GetSubscribedCatalogsIdsOk

`func (o *PriceType) GetSubscribedCatalogsIdsOk() ([]int32, bool)`

GetSubscribedCatalogsIdsOk returns a tuple with the SubscribedCatalogsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedCatalogsIds

`func (o *PriceType) HasSubscribedCatalogsIds() bool`

HasSubscribedCatalogsIds returns a boolean if a field has been set.

### SetSubscribedCatalogsIds

`func (o *PriceType) SetSubscribedCatalogsIds(v []int32)`

SetSubscribedCatalogsIds gets a reference to the given []int32 and assigns it to the SubscribedCatalogsIds field.

### GetTargetedAudiencesIds

`func (o *PriceType) GetTargetedAudiencesIds() []int32`

GetTargetedAudiencesIds returns the TargetedAudiencesIds field if non-nil, zero value otherwise.

### GetTargetedAudiencesIdsOk

`func (o *PriceType) GetTargetedAudiencesIdsOk() ([]int32, bool)`

GetTargetedAudiencesIdsOk returns a tuple with the TargetedAudiencesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedAudiencesIds

`func (o *PriceType) HasTargetedAudiencesIds() bool`

HasTargetedAudiencesIds returns a boolean if a field has been set.

### SetTargetedAudiencesIds

`func (o *PriceType) SetTargetedAudiencesIds(v []int32)`

SetTargetedAudiencesIds gets a reference to the given []int32 and assigns it to the TargetedAudiencesIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


