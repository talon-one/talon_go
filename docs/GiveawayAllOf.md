# GiveawayAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The code value of this giveaway. | 
**PoolId** | Pointer to **int32** | The ID of the pool to return giveaway codes from. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the giveaway becomes valid. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the giveaway becomes invalid. | [optional] 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this giveaway. | [optional] 
**Used** | Pointer to **bool** | Indicates whether this giveaway code was given before. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this giveaway. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The third-party integration ID of the customer profile that was awarded the giveaway, if the giveaway was awarded. | [optional] 
**ProfileId** | Pointer to **int32** | The internal ID of the customer profile that was awarded the giveaway, if the giveaway was awarded and an internal ID exists. | [optional] 

## Methods

### GetCode

`func (o *GiveawayAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GiveawayAllOf) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *GiveawayAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *GiveawayAllOf) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetPoolId

`func (o *GiveawayAllOf) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GiveawayAllOf) GetPoolIdOk() (int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoolId

`func (o *GiveawayAllOf) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### SetPoolId

`func (o *GiveawayAllOf) SetPoolId(v int32)`

SetPoolId gets a reference to the given int32 and assigns it to the PoolId field.

### GetStartDate

`func (o *GiveawayAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GiveawayAllOf) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *GiveawayAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *GiveawayAllOf) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetEndDate

`func (o *GiveawayAllOf) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GiveawayAllOf) GetEndDateOk() (time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndDate

`func (o *GiveawayAllOf) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDate

`func (o *GiveawayAllOf) SetEndDate(v time.Time)`

SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.

### GetAttributes

`func (o *GiveawayAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiveawayAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *GiveawayAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *GiveawayAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetUsed

`func (o *GiveawayAllOf) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GiveawayAllOf) GetUsedOk() (bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsed

`func (o *GiveawayAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### SetUsed

`func (o *GiveawayAllOf) SetUsed(v bool)`

SetUsed gets a reference to the given bool and assigns it to the Used field.

### GetImportId

`func (o *GiveawayAllOf) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *GiveawayAllOf) GetImportIdOk() (int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportId

`func (o *GiveawayAllOf) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportId

`func (o *GiveawayAllOf) SetImportId(v int32)`

SetImportId gets a reference to the given int32 and assigns it to the ImportId field.

### GetProfileIntegrationId

`func (o *GiveawayAllOf) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *GiveawayAllOf) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *GiveawayAllOf) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *GiveawayAllOf) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetProfileId

`func (o *GiveawayAllOf) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GiveawayAllOf) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *GiveawayAllOf) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *GiveawayAllOf) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


