# AccountDashboardStatisticApiCalls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Total number of API calls received. | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### GetTotal

`func (o *AccountDashboardStatisticApiCalls) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticApiCalls) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *AccountDashboardStatisticApiCalls) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *AccountDashboardStatisticApiCalls) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetDatetime

`func (o *AccountDashboardStatisticApiCalls) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticApiCalls) GetDatetimeOk() (time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatetime

`func (o *AccountDashboardStatisticApiCalls) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### SetDatetime

`func (o *AccountDashboardStatisticApiCalls) SetDatetime(v time.Time)`

SetDatetime gets a reference to the given time.Time and assigns it to the Datetime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


