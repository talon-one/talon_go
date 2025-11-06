# TimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | Pointer to **int64** | The achievement ends and resets in this month.  **Note**: Only applicable if the period is set to &#x60;Y&#x60;.  | [optional] 
**DayOfMonth** | Pointer to **int64** | The achievement ends and resets on this day of the month.  **Note**: Only applicable if the period is set to &#x60;Y&#x60; or &#x60;M&#x60;.  | [optional] 
**DayOfWeek** | Pointer to **int64** | The achievement ends and resets on this day of the week. &#x60;1&#x60; represents &#x60;Monday&#x60; and &#x60;7&#x60; represents &#x60;Sunday&#x60;.  **Note**: Only applicable if the period is set to &#x60;W&#x60;.  | [optional] 
**Hour** | Pointer to **int64** | The achievement ends and resets at this hour. | 
**Minute** | Pointer to **int64** | The achievement ends and resets at this minute. | 
**Second** | Pointer to **int64** | The achievement ends and resets at this second. | 

## Methods

### NewTimePoint

`func NewTimePoint(hour int64, minute int64, second int64, ) *TimePoint`

NewTimePoint instantiates a new TimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePointWithDefaults

`func NewTimePointWithDefaults() *TimePoint`

NewTimePointWithDefaults instantiates a new TimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *TimePoint) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *TimePoint) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *TimePoint) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *TimePoint) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *TimePoint) GetDayOfMonth() int64`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *TimePoint) GetDayOfMonthOk() (*int64, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *TimePoint) SetDayOfMonth(v int64)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *TimePoint) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *TimePoint) GetDayOfWeek() int64`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TimePoint) GetDayOfWeekOk() (*int64, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TimePoint) SetDayOfWeek(v int64)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TimePoint) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *TimePoint) GetHour() int64`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *TimePoint) GetHourOk() (*int64, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *TimePoint) SetHour(v int64)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *TimePoint) GetMinute() int64`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *TimePoint) GetMinuteOk() (*int64, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *TimePoint) SetMinute(v int64)`

SetMinute sets Minute field to given value.


### GetSecond

`func (o *TimePoint) GetSecond() int64`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *TimePoint) GetSecondOk() (*int64, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *TimePoint) SetSecond(v int64)`

SetSecond sets Second field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


