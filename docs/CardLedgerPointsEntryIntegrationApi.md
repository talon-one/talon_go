# CardLedgerPointsEntryIntegrationApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the transaction that adds loyalty points. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty card points were added. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**CustomerProfileID** | Pointer to **string** | Integration ID of the customer profile linked to the card. | [optional] 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where points were added. | [optional] 
**Name** | Pointer to **string** | Name or reason of the transaction that adds loyalty points. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are active immediately.   - &#x60;timestamp value&#x60;: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date and time.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added in the transaction. | 

## Methods

### GetId

`func (o *CardLedgerPointsEntryIntegrationApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CardLedgerPointsEntryIntegrationApi) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CardLedgerPointsEntryIntegrationApi) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CardLedgerPointsEntryIntegrationApi) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CardLedgerPointsEntryIntegrationApi) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CardLedgerPointsEntryIntegrationApi) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramId

`func (o *CardLedgerPointsEntryIntegrationApi) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *CardLedgerPointsEntryIntegrationApi) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *CardLedgerPointsEntryIntegrationApi) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetCustomerProfileID

`func (o *CardLedgerPointsEntryIntegrationApi) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetCustomerProfileIDOk() (string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerProfileID

`func (o *CardLedgerPointsEntryIntegrationApi) HasCustomerProfileID() bool`

HasCustomerProfileID returns a boolean if a field has been set.

### SetCustomerProfileID

`func (o *CardLedgerPointsEntryIntegrationApi) SetCustomerProfileID(v string)`

SetCustomerProfileID gets a reference to the given string and assigns it to the CustomerProfileID field.

### GetCustomerSessionId

`func (o *CardLedgerPointsEntryIntegrationApi) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *CardLedgerPointsEntryIntegrationApi) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *CardLedgerPointsEntryIntegrationApi) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetName

`func (o *CardLedgerPointsEntryIntegrationApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CardLedgerPointsEntryIntegrationApi) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CardLedgerPointsEntryIntegrationApi) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartDate

`func (o *CardLedgerPointsEntryIntegrationApi) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *CardLedgerPointsEntryIntegrationApi) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *CardLedgerPointsEntryIntegrationApi) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetExpiryDate

`func (o *CardLedgerPointsEntryIntegrationApi) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *CardLedgerPointsEntryIntegrationApi) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *CardLedgerPointsEntryIntegrationApi) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetSubledgerId

`func (o *CardLedgerPointsEntryIntegrationApi) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *CardLedgerPointsEntryIntegrationApi) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *CardLedgerPointsEntryIntegrationApi) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetAmount

`func (o *CardLedgerPointsEntryIntegrationApi) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerPointsEntryIntegrationApi) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *CardLedgerPointsEntryIntegrationApi) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *CardLedgerPointsEntryIntegrationApi) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


