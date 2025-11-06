# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]AchievementProgress**](AchievementProgress.md) |  | 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002(totalResultSize int64, data []AchievementProgress, ) *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *InlineResponse2002) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse2002) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse2002) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *InlineResponse2002) GetData() []AchievementProgress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse2002) GetDataOk() (*[]AchievementProgress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse2002) SetData(v []AchievementProgress)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


