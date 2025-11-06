# LedgerPointsEntryIntegrationAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the transaction that adds loyalty points. | 
**TransactionUUID** | Pointer to **string** | Unique identifier of the transaction in the UUID format. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Date and time the loyalty points were added. | 
**ProgramId** | Pointer to **int64** | ID of the loyalty program. | 
**CustomerSessionId** | Pointer to **string** | ID of the customer session where points were added. | [optional] 
**Name** | Pointer to **string** | Name or reason of the transaction that adds loyalty points. | 
**StartDate** | Pointer to **string** | When points become active. Possible values:   - &#x60;immediate&#x60;: Points are active immediately.   - &#x60;timestamp value&#x60;: Points become active at a given date and time.  | 
**ExpiryDate** | Pointer to **string** | Date when points expire. Possible values are:   - &#x60;unlimited&#x60;: Points have no expiration date.   - &#x60;timestamp value&#x60;: Points expire on the given date and time.  | 
**SubledgerId** | Pointer to **string** | ID of the subledger. | 
**Amount** | Pointer to **float32** | Amount of loyalty points added in the transaction. | 

## Methods

### NewLedgerPointsEntryIntegrationAPI

`func NewLedgerPointsEntryIntegrationAPI(id int64, transactionUUID string, created time.Time, programId int64, name string, startDate string, expiryDate string, subledgerId string, amount float32, ) *LedgerPointsEntryIntegrationAPI`

NewLedgerPointsEntryIntegrationAPI instantiates a new LedgerPointsEntryIntegrationAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerPointsEntryIntegrationAPIWithDefaults

`func NewLedgerPointsEntryIntegrationAPIWithDefaults() *LedgerPointsEntryIntegrationAPI`

NewLedgerPointsEntryIntegrationAPIWithDefaults instantiates a new LedgerPointsEntryIntegrationAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerPointsEntryIntegrationAPI) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerPointsEntryIntegrationAPI) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerPointsEntryIntegrationAPI) SetId(v int64)`

SetId sets Id field to given value.


### GetTransactionUUID

`func (o *LedgerPointsEntryIntegrationAPI) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LedgerPointsEntryIntegrationAPI) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LedgerPointsEntryIntegrationAPI) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *LedgerPointsEntryIntegrationAPI) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LedgerPointsEntryIntegrationAPI) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LedgerPointsEntryIntegrationAPI) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProgramId

`func (o *LedgerPointsEntryIntegrationAPI) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *LedgerPointsEntryIntegrationAPI) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *LedgerPointsEntryIntegrationAPI) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetCustomerSessionId

`func (o *LedgerPointsEntryIntegrationAPI) GetCustomerSessionId() string`

GetCustomerSessionId returns the CustomerSessionId field if non-nil, zero value otherwise.

### GetCustomerSessionIdOk

`func (o *LedgerPointsEntryIntegrationAPI) GetCustomerSessionIdOk() (*string, bool)`

GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSessionId

`func (o *LedgerPointsEntryIntegrationAPI) SetCustomerSessionId(v string)`

SetCustomerSessionId sets CustomerSessionId field to given value.

### HasCustomerSessionId

`func (o *LedgerPointsEntryIntegrationAPI) HasCustomerSessionId() bool`

HasCustomerSessionId returns a boolean if a field has been set.

### GetName

`func (o *LedgerPointsEntryIntegrationAPI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerPointsEntryIntegrationAPI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LedgerPointsEntryIntegrationAPI) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *LedgerPointsEntryIntegrationAPI) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LedgerPointsEntryIntegrationAPI) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LedgerPointsEntryIntegrationAPI) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *LedgerPointsEntryIntegrationAPI) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LedgerPointsEntryIntegrationAPI) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LedgerPointsEntryIntegrationAPI) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.


### GetSubledgerId

`func (o *LedgerPointsEntryIntegrationAPI) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *LedgerPointsEntryIntegrationAPI) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *LedgerPointsEntryIntegrationAPI) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.


### GetAmount

`func (o *LedgerPointsEntryIntegrationAPI) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerPointsEntryIntegrationAPI) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerPointsEntryIntegrationAPI) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


