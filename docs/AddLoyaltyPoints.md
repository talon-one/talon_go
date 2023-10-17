# AddLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points. | 
**Name** | Pointer to **string** | Name / reason for the point addition. | [optional] 
**ValidityDuration** | Pointer to **string** | The time format is either: - &#x60;immediate&#x60; or, - an **integer** followed by one letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  If passed, &#x60;validUntil&#x60; should be omitted.  | [optional] 
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | Date and time when points should expire. The value should be provided in RFC 3339 format. If passed, &#x60;validityDuration&#x60; should be omitted.  | [optional] 
**PendingDuration** | Pointer to **string** | The amount of time before the points are considered valid.  The time format is either: - &#x60;immediate&#x60; or, - an **integer** followed by one letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | [optional] 
**PendingUntil** | Pointer to [**time.Time**](time.Time.md) | Date and time after the points are considered valid. The value should be provided in RFC 3339 format. If passed, &#x60;pendingDuration&#x60; should be omitted.  | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger the points are added to. If there is no existing subledger with this ID, the subledger is created automatically. | [optional] 
**ApplicationId** | Pointer to **int32** | ID of the Application that is connected to the loyalty program. It is displayed in your Talon.One deployment URL. | [optional] 

## Methods

### GetPoints

`func (o *AddLoyaltyPoints) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *AddLoyaltyPoints) GetPointsOk() (float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoints

`func (o *AddLoyaltyPoints) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPoints

`func (o *AddLoyaltyPoints) SetPoints(v float32)`

SetPoints gets a reference to the given float32 and assigns it to the Points field.

### GetName

`func (o *AddLoyaltyPoints) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLoyaltyPoints) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AddLoyaltyPoints) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AddLoyaltyPoints) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValidityDuration

`func (o *AddLoyaltyPoints) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *AddLoyaltyPoints) GetValidityDurationOk() (string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidityDuration

`func (o *AddLoyaltyPoints) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### SetValidityDuration

`func (o *AddLoyaltyPoints) SetValidityDuration(v string)`

SetValidityDuration gets a reference to the given string and assigns it to the ValidityDuration field.

### GetValidUntil

`func (o *AddLoyaltyPoints) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AddLoyaltyPoints) GetValidUntilOk() (time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidUntil

`func (o *AddLoyaltyPoints) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntil

`func (o *AddLoyaltyPoints) SetValidUntil(v time.Time)`

SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.

### GetPendingDuration

`func (o *AddLoyaltyPoints) GetPendingDuration() string`

GetPendingDuration returns the PendingDuration field if non-nil, zero value otherwise.

### GetPendingDurationOk

`func (o *AddLoyaltyPoints) GetPendingDurationOk() (string, bool)`

GetPendingDurationOk returns a tuple with the PendingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingDuration

`func (o *AddLoyaltyPoints) HasPendingDuration() bool`

HasPendingDuration returns a boolean if a field has been set.

### SetPendingDuration

`func (o *AddLoyaltyPoints) SetPendingDuration(v string)`

SetPendingDuration gets a reference to the given string and assigns it to the PendingDuration field.

### GetPendingUntil

`func (o *AddLoyaltyPoints) GetPendingUntil() time.Time`

GetPendingUntil returns the PendingUntil field if non-nil, zero value otherwise.

### GetPendingUntilOk

`func (o *AddLoyaltyPoints) GetPendingUntilOk() (time.Time, bool)`

GetPendingUntilOk returns a tuple with the PendingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingUntil

`func (o *AddLoyaltyPoints) HasPendingUntil() bool`

HasPendingUntil returns a boolean if a field has been set.

### SetPendingUntil

`func (o *AddLoyaltyPoints) SetPendingUntil(v time.Time)`

SetPendingUntil gets a reference to the given time.Time and assigns it to the PendingUntil field.

### GetSubledgerId

`func (o *AddLoyaltyPoints) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *AddLoyaltyPoints) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *AddLoyaltyPoints) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *AddLoyaltyPoints) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetApplicationId

`func (o *AddLoyaltyPoints) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AddLoyaltyPoints) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *AddLoyaltyPoints) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *AddLoyaltyPoints) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


