# CardLedgerTransactionLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty card transaction occurred. | 
**ProgramId** | Pointer to **int32** | ID of the loyalty program. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | [optional] 
**SessionId** | Pointer to **int32** | The **internal** ID of the session.  | [optional] 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points become active from the given date.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Id** | Pointer to **int32** | ID of the loyalty ledger entry. | 

## Methods

### GetCreated

`func (o *CardLedgerTransactionLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerTransactionLogEntry) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CardLedgerTransactionLogEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CardLedgerTransactionLogEntry) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetProgramId

`func (o *CardLedgerTransactionLogEntry) GetProgramId() int32`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerTransactionLogEntry) GetProgramIdOk() (int32, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgramId

`func (o *CardLedgerTransactionLogEntry) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramId

`func (o *CardLedgerTransactionLogEntry) SetProgramId(v int32)`

SetProgramId gets a reference to the given int32 and assigns it to the ProgramId field.

### GetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifierOk() (string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardIdentifier

`func (o *CardLedgerTransactionLogEntry) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### SetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) SetCardIdentifier(v string)`

SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.

### GetApplicationId

`func (o *CardLedgerTransactionLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CardLedgerTransactionLogEntry) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CardLedgerTransactionLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CardLedgerTransactionLogEntry) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetSessionId

`func (o *CardLedgerTransactionLogEntry) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CardLedgerTransactionLogEntry) GetSessionIdOk() (int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *CardLedgerTransactionLogEntry) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *CardLedgerTransactionLogEntry) SetSessionId(v int32)`

SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.

### GetCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *CardLedgerTransactionLogEntry) GetCustomerSessionIdOk() (string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### SetCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) SetCustomerSessionId(v string)`

SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.

### GetType

`func (o *CardLedgerTransactionLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardLedgerTransactionLogEntry) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CardLedgerTransactionLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CardLedgerTransactionLogEntry) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetName

`func (o *CardLedgerTransactionLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerTransactionLogEntry) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CardLedgerTransactionLogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CardLedgerTransactionLogEntry) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStartDate

`func (o *CardLedgerTransactionLogEntry) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerTransactionLogEntry) GetStartDateOk() (string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *CardLedgerTransactionLogEntry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *CardLedgerTransactionLogEntry) SetStartDate(v string)`

SetStartDate gets a reference to the given string and assigns it to the StartDate field.

### GetExpiryDate

`func (o *CardLedgerTransactionLogEntry) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerTransactionLogEntry) GetExpiryDateOk() (string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *CardLedgerTransactionLogEntry) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *CardLedgerTransactionLogEntry) SetExpiryDate(v string)`

SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.

### GetSubledgerId

`func (o *CardLedgerTransactionLogEntry) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerTransactionLogEntry) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *CardLedgerTransactionLogEntry) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *CardLedgerTransactionLogEntry) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetAmount

`func (o *CardLedgerTransactionLogEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerTransactionLogEntry) GetAmountOk() (float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAmount

`func (o *CardLedgerTransactionLogEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmount

`func (o *CardLedgerTransactionLogEntry) SetAmount(v float32)`

SetAmount gets a reference to the given float32 and assigns it to the Amount field.

### GetId

`func (o *CardLedgerTransactionLogEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerTransactionLogEntry) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CardLedgerTransactionLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CardLedgerTransactionLogEntry) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


