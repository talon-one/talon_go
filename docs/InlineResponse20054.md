# InlineResponse20054

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | 
**Data** | Pointer to [**[]AchievementProgressWithDefinition**](AchievementProgressWithDefinition.md) |  | 

## Methods

### NewInlineResponse20054

`func NewInlineResponse20054(hasMore bool, data []AchievementProgressWithDefinition, ) *InlineResponse20054`

NewInlineResponse20054 instantiates a new InlineResponse20054 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20054WithDefaults

`func NewInlineResponse20054WithDefaults() *InlineResponse20054`

NewInlineResponse20054WithDefaults instantiates a new InlineResponse20054 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20054) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20054) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20054) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20054) GetData() []AchievementProgressWithDefinition`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20054) GetDataOk() (*[]AchievementProgressWithDefinition, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20054) SetData(v []AchievementProgressWithDefinition)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


