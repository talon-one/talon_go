# OutboundMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** | Cursor for the next page of results. Omitted when there are no more results. | [optional] 
**Data** | Pointer to [**[]OutboundMessage**](OutboundMessage.md) | List of outbound messages. | 

## Methods

### NewOutboundMessages

`func NewOutboundMessages(data []OutboundMessage, ) *OutboundMessages`

NewOutboundMessages instantiates a new OutboundMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessagesWithDefaults

`func NewOutboundMessagesWithDefaults() *OutboundMessages`

NewOutboundMessagesWithDefaults instantiates a new OutboundMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *OutboundMessages) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OutboundMessages) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OutboundMessages) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OutboundMessages) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetData

`func (o *OutboundMessages) GetData() []OutboundMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundMessages) GetDataOk() (*[]OutboundMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundMessages) SetData(v []OutboundMessage)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


