# MessageLogEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MessageLogEntry**](MessageLogEntry.md) | List of message logs. | 
**NextCursor** | Pointer to **string** | The next value in the database.  **Note:** If this value is not present, it means that there are no more values in the database for this combination of request parameters.  | [optional] 

## Methods

### GetData

`func (o *MessageLogEntries) GetData() []MessageLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MessageLogEntries) GetDataOk() ([]MessageLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *MessageLogEntries) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *MessageLogEntries) SetData(v []MessageLogEntry)`

SetData gets a reference to the given []MessageLogEntry and assigns it to the Data field.

### GetNextCursor

`func (o *MessageLogEntries) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *MessageLogEntries) GetNextCursorOk() (string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextCursor

`func (o *MessageLogEntries) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursor

`func (o *MessageLogEntries) SetNextCursor(v string)`

SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


