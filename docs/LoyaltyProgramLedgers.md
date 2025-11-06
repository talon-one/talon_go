# LoyaltyProgramLedgers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of loyalty program. | 
**Title** | Pointer to **string** | Visible name of loyalty program. | 
**Name** | Pointer to **string** | Internal name of loyalty program. | 
**JoinDate** | Pointer to [**time.Time**](time.Time.md) | The date on which the customer joined the loyalty program in RFC3339.  **Note**: This is in the loyalty program&#39;s time zone.  | [optional] 
**Ledger** | Pointer to [**LedgerInfo**](LedgerInfo.md) |  | 
**SubLedgers** | Pointer to [**map[string]LedgerInfo**](LedgerInfo.md) | A map containing information about each loyalty subledger. | [optional] 

## Methods

### NewLoyaltyProgramLedgers

`func NewLoyaltyProgramLedgers(id int64, title string, name string, ledger LedgerInfo, ) *LoyaltyProgramLedgers`

NewLoyaltyProgramLedgers instantiates a new LoyaltyProgramLedgers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramLedgersWithDefaults

`func NewLoyaltyProgramLedgersWithDefaults() *LoyaltyProgramLedgers`

NewLoyaltyProgramLedgersWithDefaults instantiates a new LoyaltyProgramLedgers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoyaltyProgramLedgers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramLedgers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyProgramLedgers) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *LoyaltyProgramLedgers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgramLedgers) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LoyaltyProgramLedgers) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *LoyaltyProgramLedgers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramLedgers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoyaltyProgramLedgers) SetName(v string)`

SetName sets Name field to given value.


### GetJoinDate

`func (o *LoyaltyProgramLedgers) GetJoinDate() time.Time`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *LoyaltyProgramLedgers) GetJoinDateOk() (*time.Time, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *LoyaltyProgramLedgers) SetJoinDate(v time.Time)`

SetJoinDate sets JoinDate field to given value.

### HasJoinDate

`func (o *LoyaltyProgramLedgers) HasJoinDate() bool`

HasJoinDate returns a boolean if a field has been set.

### GetLedger

`func (o *LoyaltyProgramLedgers) GetLedger() LedgerInfo`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyProgramLedgers) GetLedgerOk() (*LedgerInfo, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *LoyaltyProgramLedgers) SetLedger(v LedgerInfo)`

SetLedger sets Ledger field to given value.


### GetSubLedgers

`func (o *LoyaltyProgramLedgers) GetSubLedgers() map[string]LedgerInfo`

GetSubLedgers returns the SubLedgers field if non-nil, zero value otherwise.

### GetSubLedgersOk

`func (o *LoyaltyProgramLedgers) GetSubLedgersOk() (*map[string]LedgerInfo, bool)`

GetSubLedgersOk returns a tuple with the SubLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgers

`func (o *LoyaltyProgramLedgers) SetSubLedgers(v map[string]LedgerInfo)`

SetSubLedgers sets SubLedgers field to given value.

### HasSubLedgers

`func (o *LoyaltyProgramLedgers) HasSubLedgers() bool`

HasSubLedgers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


