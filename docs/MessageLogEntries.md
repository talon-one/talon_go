# MessageLogEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** | The next value in the database.  **Note:** If this value is not present, it means that there are no more values in the database for this combination of request parameters.  | [optional] 
**Data** | Pointer to [**[]MessageLogEntry**](MessageLogEntry.md) | List of message logs. | 

## Methods

### NewMessageLogEntries

`func NewMessageLogEntries(data []MessageLogEntry, ) *MessageLogEntries`

NewMessageLogEntries instantiates a new MessageLogEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageLogEntriesWithDefaults

`func NewMessageLogEntriesWithDefaults() *MessageLogEntries`

NewMessageLogEntriesWithDefaults instantiates a new MessageLogEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *MessageLogEntries) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *MessageLogEntries) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *MessageLogEntries) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *MessageLogEntries) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetData

`func (o *MessageLogEntries) GetData() []MessageLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessageLogEntries) GetDataOk() (*[]MessageLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MessageLogEntries) SetData(v []MessageLogEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


