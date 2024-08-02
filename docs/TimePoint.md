# TimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The achievement ends and resets on this day of the month.  **Note**: Only applicable if the period is set to &#x60;Y&#x60; or &#x60;M&#x60;.  | [optional] 
**DayOfWeek** | Pointer to **int32** | The achievement ends and resets on this day of the week. &#x60;1&#x60; represents &#x60;Monday&#x60; and &#x60;7&#x60; represents &#x60;Sunday&#x60;.  **Note**: Only applicable if the period is set to &#x60;W&#x60;.  | [optional] 
**Hour** | Pointer to **int32** | The achievement ends and resets at this hour. | 
**Minute** | Pointer to **int32** | The achievement ends and resets at this minute. | 
**Month** | Pointer to **int32** | The achievement ends and resets in this month.  **Note**: Only applicable if the period is set to &#x60;Y&#x60;.  | [optional] 
**Second** | Pointer to **int32** | The achievement ends and resets at this second. | 

## Methods

### GetDayOfMonth

`func (o *TimePoint) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *TimePoint) GetDayOfMonthOk() (int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDayOfMonth

`func (o *TimePoint) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonth

`func (o *TimePoint) SetDayOfMonth(v int32)`

SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.

### GetDayOfWeek

`func (o *TimePoint) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TimePoint) GetDayOfWeekOk() (int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDayOfWeek

`func (o *TimePoint) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeek

`func (o *TimePoint) SetDayOfWeek(v int32)`

SetDayOfWeek gets a reference to the given int32 and assigns it to the DayOfWeek field.

### GetHour

`func (o *TimePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *TimePoint) GetHourOk() (int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHour

`func (o *TimePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHour

`func (o *TimePoint) SetHour(v int32)`

SetHour gets a reference to the given int32 and assigns it to the Hour field.

### GetMinute

`func (o *TimePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *TimePoint) GetMinuteOk() (int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinute

`func (o *TimePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### SetMinute

`func (o *TimePoint) SetMinute(v int32)`

SetMinute gets a reference to the given int32 and assigns it to the Minute field.

### GetMonth

`func (o *TimePoint) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *TimePoint) GetMonthOk() (int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonth

`func (o *TimePoint) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonth

`func (o *TimePoint) SetMonth(v int32)`

SetMonth gets a reference to the given int32 and assigns it to the Month field.

### GetSecond

`func (o *TimePoint) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *TimePoint) GetSecondOk() (int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecond

`func (o *TimePoint) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### SetSecond

`func (o *TimePoint) SetSecond(v int32)`

SetSecond gets a reference to the given int32 and assigns it to the Second field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


