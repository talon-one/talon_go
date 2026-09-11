# ExcludePriceObservationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int64** | A list of historical price IDs to exclude from best prior price calculation. Must contain between 1 and 1000 IDs. All IDs must be valid &#x60;id&#x60; values obtained from the [Get summary of price history](https://docs.talon.one/management-api#tag/Catalogs/operation/priceHistory.responses.200.history) endpoint, must belong to the specified Application, and must not already be excluded from best prior price calculation.  | 
**Reason** | Pointer to **string** | The reason for excluding these historical price IDs. Applies to all IDs in the batch.  | 

## Methods

### NewExcludePriceObservationsRequest

`func NewExcludePriceObservationsRequest(ids []int64, reason string, ) *ExcludePriceObservationsRequest`

NewExcludePriceObservationsRequest instantiates a new ExcludePriceObservationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludePriceObservationsRequestWithDefaults

`func NewExcludePriceObservationsRequestWithDefaults() *ExcludePriceObservationsRequest`

NewExcludePriceObservationsRequestWithDefaults instantiates a new ExcludePriceObservationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ExcludePriceObservationsRequest) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ExcludePriceObservationsRequest) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ExcludePriceObservationsRequest) SetIds(v []int64)`

SetIds sets Ids field to given value.


### GetReason

`func (o *ExcludePriceObservationsRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ExcludePriceObservationsRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ExcludePriceObservationsRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


