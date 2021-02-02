# LoyaltyProgramLedgers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of loyalty program | 
**Title** | Pointer to **string** | Visible name of loyalty program | 
**Name** | Pointer to **string** | Internal name of loyalty program | 
**Ledger** | Pointer to [**LoyaltyProgramBalance**](LoyaltyProgramBalance.md) |  | 
**SubLedgers** | Pointer to [**map[string]LoyaltyProgramBalance**](LoyaltyProgramBalance.md) | A map containing a list of all loyalty subledger balances | [optional] 

## Methods

### GetId

`func (o *LoyaltyProgramLedgers) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramLedgers) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyProgramLedgers) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyProgramLedgers) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetTitle

`func (o *LoyaltyProgramLedgers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgramLedgers) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *LoyaltyProgramLedgers) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *LoyaltyProgramLedgers) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetName

`func (o *LoyaltyProgramLedgers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramLedgers) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyProgramLedgers) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyProgramLedgers) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetLedger

`func (o *LoyaltyProgramLedgers) GetLedger() LoyaltyProgramBalance`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyProgramLedgers) GetLedgerOk() (LoyaltyProgramBalance, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLedger

`func (o *LoyaltyProgramLedgers) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### SetLedger

`func (o *LoyaltyProgramLedgers) SetLedger(v LoyaltyProgramBalance)`

SetLedger gets a reference to the given LoyaltyProgramBalance and assigns it to the Ledger field.

### GetSubLedgers

`func (o *LoyaltyProgramLedgers) GetSubLedgers() map[string]LoyaltyProgramBalance`

GetSubLedgers returns the SubLedgers field if non-nil, zero value otherwise.

### GetSubLedgersOk

`func (o *LoyaltyProgramLedgers) GetSubLedgersOk() (map[string]LoyaltyProgramBalance, bool)`

GetSubLedgersOk returns a tuple with the SubLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgers

`func (o *LoyaltyProgramLedgers) HasSubLedgers() bool`

HasSubLedgers returns a boolean if a field has been set.

### SetSubLedgers

`func (o *LoyaltyProgramLedgers) SetSubLedgers(v map[string]LoyaltyProgramBalance)`

SetSubLedgers gets a reference to the given map[string]LoyaltyProgramBalance and assigns it to the SubLedgers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


