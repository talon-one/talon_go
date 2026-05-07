# MCPKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the MCP key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the MCP key expires. | 
**Id** | Pointer to **int64** | ID of the MCP key. | 
**CreatedBy** | Pointer to **int64** | ID of the user who created it. | 
**AccountID** | Pointer to **int64** | ID of account the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the MCP key was created. | 
**Disabled** | Pointer to **bool** | The MCP key is disabled (this property is set to &#x60;true&#x60;) when the user who created the key is disabled or deleted. | [optional] 
**LastUsed** | Pointer to [**time.Time**](time.Time.md) | The last time the MCP key was used. | [optional] 

## Methods

### NewMCPKey

`func NewMCPKey(name string, expiryDate time.Time, id int64, createdBy int64, accountID int64, created time.Time, ) *MCPKey`

NewMCPKey instantiates a new MCPKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPKeyWithDefaults

`func NewMCPKeyWithDefaults() *MCPKey`

NewMCPKeyWithDefaults instantiates a new MCPKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MCPKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDate

`func (o *MCPKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MCPKey) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MCPKey) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetId

`func (o *MCPKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPKey) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *MCPKey) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MCPKey) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MCPKey) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetAccountID

`func (o *MCPKey) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *MCPKey) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *MCPKey) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.


### GetCreated

`func (o *MCPKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MCPKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MCPKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDisabled

`func (o *MCPKey) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MCPKey) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MCPKey) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MCPKey) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLastUsed

`func (o *MCPKey) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *MCPKey) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *MCPKey) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *MCPKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


