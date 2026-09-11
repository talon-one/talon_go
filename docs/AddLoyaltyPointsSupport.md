# AddLoyaltyPointsSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points. | 
**Name** | Pointer to **string** | Name / reason for the point addition. | [optional] 
**ValidityDuration** | Pointer to **string** | The time format is either: - &#x60;unlimited&#x60; or, - an **integer** followed by one letter indicating the time unit.  Examples: &#x60;unlimited&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  If passed, &#x60;validUntil&#x60; should be omitted.  | [optional] 
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | Date and time when points should expire. The value should be provided in RFC 3339 format. If passed, &#x60;validityDuration&#x60; should be omitted.  | [optional] 
**PendingDuration** | Pointer to **string** | The amount of time before the points are considered valid.  The time format is either: - &#x60;immediate&#x60; or, - &#x60;on_action&#x60; or, - an **integer** followed by one letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;, &#x60;on_action&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | [optional] 
**PendingUntil** | Pointer to [**time.Time**](time.Time.md) | Date and time after the points are considered valid. The value should be provided in RFC 3339 format. If passed, &#x60;pendingDuration&#x60; should be omitted.  | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger the points are added to. If there is no existing subledger with this ID, the subledger is created automatically. | [optional] 
**ApplicationId** | Pointer to **int64** | ID of the Application that is connected to the loyalty program. It is displayed in your Talon.One deployment URL. | [optional] 
**SupportRequestId** | Pointer to **int64** | ID of the support request to approve. When provided by an admin, the points are added on behalf of the support user who created the request. | [optional] 
**ProcessingNote** | Pointer to **string** | Note from the admin approving the support request. Stored as the processing note on the support request record. This is only used when a supportRequestId is passed. | [optional] 

## Methods

### NewAddLoyaltyPointsSupport

`func NewAddLoyaltyPointsSupport(points float32, ) *AddLoyaltyPointsSupport`

NewAddLoyaltyPointsSupport instantiates a new AddLoyaltyPointsSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLoyaltyPointsSupportWithDefaults

`func NewAddLoyaltyPointsSupportWithDefaults() *AddLoyaltyPointsSupport`

NewAddLoyaltyPointsSupportWithDefaults instantiates a new AddLoyaltyPointsSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *AddLoyaltyPointsSupport) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *AddLoyaltyPointsSupport) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *AddLoyaltyPointsSupport) SetPoints(v float32)`

SetPoints sets Points field to given value.


### GetName

`func (o *AddLoyaltyPointsSupport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLoyaltyPointsSupport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLoyaltyPointsSupport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddLoyaltyPointsSupport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidityDuration

`func (o *AddLoyaltyPointsSupport) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *AddLoyaltyPointsSupport) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *AddLoyaltyPointsSupport) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *AddLoyaltyPointsSupport) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetValidUntil

`func (o *AddLoyaltyPointsSupport) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AddLoyaltyPointsSupport) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *AddLoyaltyPointsSupport) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *AddLoyaltyPointsSupport) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetPendingDuration

`func (o *AddLoyaltyPointsSupport) GetPendingDuration() string`

GetPendingDuration returns the PendingDuration field if non-nil, zero value otherwise.

### GetPendingDurationOk

`func (o *AddLoyaltyPointsSupport) GetPendingDurationOk() (*string, bool)`

GetPendingDurationOk returns a tuple with the PendingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDuration

`func (o *AddLoyaltyPointsSupport) SetPendingDuration(v string)`

SetPendingDuration sets PendingDuration field to given value.

### HasPendingDuration

`func (o *AddLoyaltyPointsSupport) HasPendingDuration() bool`

HasPendingDuration returns a boolean if a field has been set.

### GetPendingUntil

`func (o *AddLoyaltyPointsSupport) GetPendingUntil() time.Time`

GetPendingUntil returns the PendingUntil field if non-nil, zero value otherwise.

### GetPendingUntilOk

`func (o *AddLoyaltyPointsSupport) GetPendingUntilOk() (*time.Time, bool)`

GetPendingUntilOk returns a tuple with the PendingUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUntil

`func (o *AddLoyaltyPointsSupport) SetPendingUntil(v time.Time)`

SetPendingUntil sets PendingUntil field to given value.

### HasPendingUntil

`func (o *AddLoyaltyPointsSupport) HasPendingUntil() bool`

HasPendingUntil returns a boolean if a field has been set.

### GetSubledgerId

`func (o *AddLoyaltyPointsSupport) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *AddLoyaltyPointsSupport) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *AddLoyaltyPointsSupport) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.

### HasSubledgerId

`func (o *AddLoyaltyPointsSupport) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### GetApplicationId

`func (o *AddLoyaltyPointsSupport) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AddLoyaltyPointsSupport) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AddLoyaltyPointsSupport) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AddLoyaltyPointsSupport) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetSupportRequestId

`func (o *AddLoyaltyPointsSupport) GetSupportRequestId() int64`

GetSupportRequestId returns the SupportRequestId field if non-nil, zero value otherwise.

### GetSupportRequestIdOk

`func (o *AddLoyaltyPointsSupport) GetSupportRequestIdOk() (*int64, bool)`

GetSupportRequestIdOk returns a tuple with the SupportRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestId

`func (o *AddLoyaltyPointsSupport) SetSupportRequestId(v int64)`

SetSupportRequestId sets SupportRequestId field to given value.

### HasSupportRequestId

`func (o *AddLoyaltyPointsSupport) HasSupportRequestId() bool`

HasSupportRequestId returns a boolean if a field has been set.

### GetProcessingNote

`func (o *AddLoyaltyPointsSupport) GetProcessingNote() string`

GetProcessingNote returns the ProcessingNote field if non-nil, zero value otherwise.

### GetProcessingNoteOk

`func (o *AddLoyaltyPointsSupport) GetProcessingNoteOk() (*string, bool)`

GetProcessingNoteOk returns a tuple with the ProcessingNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingNote

`func (o *AddLoyaltyPointsSupport) SetProcessingNote(v string)`

SetProcessingNote sets ProcessingNote field to given value.

### HasProcessingNote

`func (o *AddLoyaltyPointsSupport) HasProcessingNote() bool`

HasProcessingNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


