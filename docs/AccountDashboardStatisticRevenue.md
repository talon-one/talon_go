# AccountDashboardStatisticRevenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | All revenue that went through the client&#39;s shop (including purchases that didn’t trigger an effect). | 
**Influenced** | Pointer to **float32** | The revenue that was created by a purchase that triggered an effect (excluding web hooks, notifications). | 
**Datetime** | Pointer to [**time.Time**](time.Time.md) | Values aggregated for the specified date. | 

## Methods

### GetTotal

`func (o *AccountDashboardStatisticRevenue) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountDashboardStatisticRevenue) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *AccountDashboardStatisticRevenue) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *AccountDashboardStatisticRevenue) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetInfluenced

`func (o *AccountDashboardStatisticRevenue) GetInfluenced() float32`

GetInfluenced returns the Influenced field if non-nil, zero value otherwise.

### GetInfluencedOk

`func (o *AccountDashboardStatisticRevenue) GetInfluencedOk() (float32, bool)`

GetInfluencedOk returns a tuple with the Influenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInfluenced

`func (o *AccountDashboardStatisticRevenue) HasInfluenced() bool`

HasInfluenced returns a boolean if a field has been set.

### SetInfluenced

`func (o *AccountDashboardStatisticRevenue) SetInfluenced(v float32)`

SetInfluenced gets a reference to the given float32 and assigns it to the Influenced field.

### GetDatetime

`func (o *AccountDashboardStatisticRevenue) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AccountDashboardStatisticRevenue) GetDatetimeOk() (time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDatetime

`func (o *AccountDashboardStatisticRevenue) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### SetDatetime

`func (o *AccountDashboardStatisticRevenue) SetDatetime(v time.Time)`

SetDatetime gets a reference to the given time.Time and assigns it to the Datetime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


