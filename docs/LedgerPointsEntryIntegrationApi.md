# LedgerPointsEntryIntegrationApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of loyalty points added in the transaction. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty points were added. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where points were added. | [optional] 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date and time.  | 
**Id** | Pointer to **int32** | ID of the transaction that adds loyalty points. | 
**Name** | Pointer to **string** | Name or reason of the transaction that adds loyalty points. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are active immediately.   - &#x60;timestamp value&#x60;: Points become active at a given date and time.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 

## Methods

### GetAmount

`func (o *LedgerPointsEntryIntegrationApi) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerPointsEntryIntegrationApi) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *LedgerPointsEntryIntegrationApi) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *LedgerPointsEntryIntegrationApi) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetCreated

`func (o *LedgerPointsEntryIntegrationApi) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerPointsEntryIntegrationApi) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *LedgerPointsEntryIntegrationApi) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *LedgerPointsEntryIntegrationApi) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCustomerSessionId

`func (o *LedgerPointsEntryIntegrationApi) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LedgerPointsEntryIntegrationApi) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *LedgerPointsEntryIntegrationApi) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *LedgerPointsEntryIntegrationApi) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetExpiryDate

`func (o *LedgerPointsEntryIntegrationApi) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerPointsEntryIntegrationApi) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *LedgerPointsEntryIntegrationApi) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *LedgerPointsEntryIntegrationApi) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetId

`func (o *LedgerPointsEntryIntegrationApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerPointsEntryIntegrationApi) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LedgerPointsEntryIntegrationApi) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LedgerPointsEntryIntegrationApi) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *LedgerPointsEntryIntegrationApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerPointsEntryIntegrationApi) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LedgerPointsEntryIntegrationApi) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LedgerPointsEntryIntegrationApi) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetProgramId

`func (o *LedgerPointsEntryIntegrationApi) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LedgerPointsEntryIntegrationApi) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *LedgerPointsEntryIntegrationApi) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *LedgerPointsEntryIntegrationApi) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetStartDate

`func (o *LedgerPointsEntryIntegrationApi) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LedgerPointsEntryIntegrationApi) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *LedgerPointsEntryIntegrationApi) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *LedgerPointsEntryIntegrationApi) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetSubledgerId

`func (o *LedgerPointsEntryIntegrationApi) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LedgerPointsEntryIntegrationApi) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *LedgerPointsEntryIntegrationApi) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *LedgerPointsEntryIntegrationApi) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


