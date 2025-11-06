# Giveaway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Code** | Pointer to **string** | The code value of this giveaway. | 
**PoolId** | Pointer to **int64** | The ID of the pool to return giveaway codes from. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the giveaway becomes valid. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the giveaway becomes invalid. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this giveaway. | [optional] 
**Used** | Pointer to **bool** | Indicates whether this giveaway code was given before. | [optional] 
**ImportId** | Pointer to **int64** | The ID of the Import which created this giveaway. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The third-party integration ID of the customer profile that was awarded the giveaway, if the giveaway was awarded. | [optional] 
**ProfileId** | Pointer to **int64** | The internal ID of the customer profile that was awarded the giveaway, if the giveaway was awarded and an internal ID exists. | [optional] 

## Methods

### NewGiveaway

`func NewGiveaway(id int64, created time.Time, code string, poolId int64, ) *Giveaway`

NewGiveaway instantiates a new Giveaway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayWithDefaults

`func NewGiveawayWithDefaults() *Giveaway`

NewGiveawayWithDefaults instantiates a new Giveaway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Giveaway) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Giveaway) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Giveaway) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Giveaway) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Giveaway) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Giveaway) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCode

`func (o *Giveaway) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Giveaway) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Giveaway) SetCode(v string)`

SetCode sets Code field to given value.


### GetPoolId

`func (o *Giveaway) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Giveaway) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Giveaway) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetStartDate

`func (o *Giveaway) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Giveaway) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Giveaway) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Giveaway) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Giveaway) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Giveaway) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Giveaway) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Giveaway) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAttributes

`func (o *Giveaway) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Giveaway) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Giveaway) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Giveaway) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUsed

`func (o *Giveaway) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Giveaway) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Giveaway) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Giveaway) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetImportId

`func (o *Giveaway) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Giveaway) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *Giveaway) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *Giveaway) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetProfileIntegrationId

`func (o *Giveaway) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *Giveaway) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *Giveaway) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.

### HasProfileIntegrationId

`func (o *Giveaway) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### GetProfileId

`func (o *Giveaway) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Giveaway) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Giveaway) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Giveaway) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


