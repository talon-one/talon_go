# InlineResponse20040

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]WebhookWithOutgoingIntegrationDetails**](WebhookWithOutgoingIntegrationDetails.md) |  | 

## Methods

### NewInlineResponse20040

`func NewInlineResponse20040(totalResultSize int64, data []WebhookWithOutgoingIntegrationDetails, ) *InlineResponse20040`

NewInlineResponse20040 instantiates a new InlineResponse20040 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040WithDefaults

`func NewInlineResponse20040WithDefaults() *InlineResponse20040`

NewInlineResponse20040WithDefaults instantiates a new InlineResponse20040 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *InlineResponse20040) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *InlineResponse20040) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *InlineResponse20040) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *InlineResponse20040) GetData() []WebhookWithOutgoingIntegrationDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20040) GetDataOk() (*[]WebhookWithOutgoingIntegrationDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20040) SetData(v []WebhookWithOutgoingIntegrationDetails)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


