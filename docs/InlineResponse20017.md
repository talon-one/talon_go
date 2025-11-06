# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]LoyaltyProgramTransaction**](LoyaltyProgramTransaction.md) |  | 

## Methods

### NewInlineResponse20017

`func NewInlineResponse20017(hasMore bool, data []LoyaltyProgramTransaction, ) *InlineResponse20017`

NewInlineResponse20017 instantiates a new InlineResponse20017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017WithDefaults

`func NewInlineResponse20017WithDefaults() *InlineResponse20017`

NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20017) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20017) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20017) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20017) GetData() []LoyaltyProgramTransaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20017) GetDataOk() (*[]LoyaltyProgramTransaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20017) SetData(v []LoyaltyProgramTransaction)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


