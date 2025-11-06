# CardLedgerTransactionLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionUUID** | Pointer to **string** | Unique identifier of the transaction in the UUID format. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty card transaction occurred. | 
**ProgramId** | Pointer to **int64** | ID of the loyalty program. | 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | [optional] 
**SessionId** | Pointer to **int64** | The **internal** ID of the session.  | [optional] 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where the transaction occurred. | [optional] 
**Type** | Pointer to **string** | Type of transaction. Possible values:   - &#x60;addition&#x60;: Signifies added points.   - &#x60;subtraction&#x60;: Signifies deducted points.  | 
**Name** | Pointer to **string** | Name or reason of the loyalty ledger transaction. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are immediately active.   - a timestamp value: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points become active from the given date.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added or deducted in the transaction. | 
**Id** | Pointer to **int64** | ID of the loyalty ledger entry. | 

## Methods

### NewCardLedgerTransactionLogEntry

`func NewCardLedgerTransactionLogEntry(transactionUUID string, created time.Time, programId int64, cardIdentifier string, type_ string, name string, startDate string, expiryDate string, subledgerId string, amount float32, id int64, ) *CardLedgerTransactionLogEntry`

NewCardLedgerTransactionLogEntry instantiates a new CardLedgerTransactionLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardLedgerTransactionLogEntryWithDefaults

`func NewCardLedgerTransactionLogEntryWithDefaults() *CardLedgerTransactionLogEntry`

NewCardLedgerTransactionLogEntryWithDefaults instantiates a new CardLedgerTransactionLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionUUID

`func (o *CardLedgerTransactionLogEntry) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *CardLedgerTransactionLogEntry) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *CardLedgerTransactionLogEntry) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *CardLedgerTransactionLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CardLedgerTransactionLogEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CardLedgerTransactionLogEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramId

`func (o *CardLedgerTransactionLogEntry) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CardLedgerTransactionLogEntry) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *CardLedgerTransactionLogEntry) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *CardLedgerTransactionLogEntry) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *CardLedgerTransactionLogEntry) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetApplicationId

`func (o *CardLedgerTransactionLogEntry) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CardLedgerTransactionLogEntry) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CardLedgerTransactionLogEntry) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CardLedgerTransactionLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetSessionId

`func (o *CardLedgerTransactionLogEntry) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CardLedgerTransactionLogEntry) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CardLedgerTransactionLogEntry) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CardLedgerTransactionLogEntry) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *CardLedgerTransactionLogEntry) GetCustomerSessionIdOk() (*string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) SetCustomerSessionId(v string)`

SetCustomerSessionId sets CustomerSessionId field to given value.

### HasCustomerSessionId

`func (o *CardLedgerTransactionLogEntry) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### GetType

`func (o *CardLedgerTransactionLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardLedgerTransactionLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardLedgerTransactionLogEntry) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CardLedgerTransactionLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardLedgerTransactionLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardLedgerTransactionLogEntry) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CardLedgerTransactionLogEntry) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardLedgerTransactionLogEntry) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardLedgerTransactionLogEntry) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *CardLedgerTransactionLogEntry) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CardLedgerTransactionLogEntry) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CardLedgerTransactionLogEntry) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetSubledgerId

`func (o *CardLedgerTransactionLogEntry) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *CardLedgerTransactionLogEntry) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *CardLedgerTransactionLogEntry) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetAmount

`func (o *CardLedgerTransactionLogEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardLedgerTransactionLogEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardLedgerTransactionLogEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *CardLedgerTransactionLogEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardLedgerTransactionLogEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardLedgerTransactionLogEntry) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


