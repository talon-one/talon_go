# LoyaltyCardBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfCards** | Pointer to **int64** | Number of loyalty cards in the batch. | 
**BatchId** | Pointer to **string** | ID of the loyalty card batch. | [optional] 
**Status** | Pointer to **string** | Status of the loyalty cards in the batch. | [optional] [default to "active"]
**CardCodeSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 

## Methods

### NewLoyaltyCardBatch

`func NewLoyaltyCardBatch(numberOfCards int64, ) *LoyaltyCardBatch`

NewLoyaltyCardBatch instantiates a new LoyaltyCardBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardBatchWithDefaults

`func NewLoyaltyCardBatchWithDefaults() *LoyaltyCardBatch`

NewLoyaltyCardBatchWithDefaults instantiates a new LoyaltyCardBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfCards

`func (o *LoyaltyCardBatch) GetNumberOfCards() int64`

GetNumberOfCards returns the NumberOfCards field if non-nil, zero value otherwise.

### GetNumberOfCardsOk

`func (o *LoyaltyCardBatch) GetNumberOfCardsOk() (*int64, bool)`

GetNumberOfCardsOk returns a tuple with the NumberOfCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCards

`func (o *LoyaltyCardBatch) SetNumberOfCards(v int64)`

SetNumberOfCards sets NumberOfCards field to given value.


### GetBatchId

`func (o *LoyaltyCardBatch) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *LoyaltyCardBatch) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *LoyaltyCardBatch) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *LoyaltyCardBatch) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetStatus

`func (o *LoyaltyCardBatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoyaltyCardBatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoyaltyCardBatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoyaltyCardBatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCardCodeSettings

`func (o *LoyaltyCardBatch) GetCardCodeSettings() CodeGeneratorSettings`

GetCardCodeSettings returns the CardCodeSettings field if non-nil, zero value otherwise.

### GetCardCodeSettingsOk

`func (o *LoyaltyCardBatch) GetCardCodeSettingsOk() (*CodeGeneratorSettings, bool)`

GetCardCodeSettingsOk returns a tuple with the CardCodeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCodeSettings

`func (o *LoyaltyCardBatch) SetCardCodeSettings(v CodeGeneratorSettings)`

SetCardCodeSettings sets CardCodeSettings field to given value.

### HasCardCodeSettings

`func (o *LoyaltyCardBatch) HasCardCodeSettings() bool`

HasCardCodeSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


