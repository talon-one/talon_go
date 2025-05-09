# ManagementKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for management key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the management key expires. | 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | The list of endpoints that can be accessed with the key | 
**AllowedApplicationIds** | Pointer to **[]int32** | A list of Application IDs that you can access with the management key. An empty or missing list means the management key can be used for all Applications in the account.  | [optional] 
**Id** | Pointer to **int32** | ID of the management key. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created it. | 
**AccountID** | Pointer to **int32** | ID of account the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the management key was created. | 
**Disabled** | Pointer to **bool** | The management key is disabled (this property is set to &#x60;true&#x60;) when the user who created the key is disabled or deleted. | [optional] 

## Methods

### GetName

`func (o *ManagementKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementKey) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ManagementKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ManagementKey) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetExpiryDate

`func (o *ManagementKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ManagementKey) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *ManagementKey) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *ManagementKey) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetEndpoints

`func (o *ManagementKey) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ManagementKey) GetEndpointsOk() ([]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndpoints

`func (o *ManagementKey) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpoints

`func (o *ManagementKey) SetEndpoints(v []Endpoint)`

SetEndpoints gets a reference to the given []Endpoint and assigns it to the Endpoints field.

### GetAllowedApplicationIds

`func (o *ManagementKey) GetAllowedApplicationIds() []int32`

GetAllowedApplicationIds returns the AllowedApplicationIds field if non-nil, zero value otherwise.

### GetAllowedApplicationIdsOk

`func (o *ManagementKey) GetAllowedApplicationIdsOk() ([]int32, bool)`

GetAllowedApplicationIdsOk returns a tuple with the AllowedApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedApplicationIds

`func (o *ManagementKey) HasAllowedApplicationIds() bool`

HasAllowedApplicationIds returns a boolean if a field has been set.

### SetAllowedApplicationIds

`func (o *ManagementKey) SetAllowedApplicationIds(v []int32)`

SetAllowedApplicationIds gets a reference to the given []int32 and assigns it to the AllowedApplicationIds field.

### GetId

`func (o *ManagementKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementKey) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ManagementKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ManagementKey) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreatedBy

`func (o *ManagementKey) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ManagementKey) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *ManagementKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *ManagementKey) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetAccountID

`func (o *ManagementKey) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ManagementKey) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *ManagementKey) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *ManagementKey) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetCreated

`func (o *ManagementKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ManagementKey) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ManagementKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ManagementKey) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetDisabled

`func (o *ManagementKey) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ManagementKey) GetDisabledOk() (bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisabled

`func (o *ManagementKey) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabled

`func (o *ManagementKey) SetDisabled(v bool)`

SetDisabled gets a reference to the given bool and assigns it to the Disabled field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


