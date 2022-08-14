# NewManagementKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for management key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the management key expires. | 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | The list of endpoints that can be accessed with the key | 
**Id** | Pointer to **int32** | ID of the management key. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created it. | 
**AccountID** | Pointer to **int32** | ID of account the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the management key was created. | 
**Key** | Pointer to **string** | The management key. | 

## Methods

### GetName

`func (o *NewManagementKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewManagementKey) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewManagementKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewManagementKey) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetExpiryDate

`func (o *NewManagementKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewManagementKey) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *NewManagementKey) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *NewManagementKey) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetEndpoints

`func (o *NewManagementKey) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *NewManagementKey) GetEndpointsOk() ([]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndpoints

`func (o *NewManagementKey) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpoints

`func (o *NewManagementKey) SetEndpoints(v []Endpoint)`

SetEndpoints gets a reference to the given []Endpoint and assigns it to the Endpoints field.

### GetId

`func (o *NewManagementKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewManagementKey) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *NewManagementKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *NewManagementKey) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreatedBy

`func (o *NewManagementKey) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NewManagementKey) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *NewManagementKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *NewManagementKey) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetAccountID

`func (o *NewManagementKey) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *NewManagementKey) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *NewManagementKey) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *NewManagementKey) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetCreated

`func (o *NewManagementKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NewManagementKey) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *NewManagementKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *NewManagementKey) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetKey

`func (o *NewManagementKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewManagementKey) GetKeyOk() (string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKey

`func (o *NewManagementKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKey

`func (o *NewManagementKey) SetKey(v string)`

SetKey gets a reference to the given string and assigns it to the Key field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


