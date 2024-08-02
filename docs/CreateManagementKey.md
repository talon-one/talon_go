# CreateManagementKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedApplicationIds** | Pointer to **[]int32** | A list of Application IDs that you can access with the management key. An empty or missing list means the management key can be used for all Applications in the account.  | [optional] 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | The list of endpoints that can be accessed with the key | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the management key expires. | 
**Name** | Pointer to **string** | Name for management key. | 

## Methods

### GetAllowedApplicationIds

`func (o *CreateManagementKey) GetAllowedApplicationIds() []int32`

GetAllowedApplicationIds returns the AllowedApplicationIds field if non-nil, zero value otherwise.

### GetAllowedApplicationIdsOk

`func (o *CreateManagementKey) GetAllowedApplicationIdsOk() ([]int32, bool)`

GetAllowedApplicationIdsOk returns a tuple with the AllowedApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedApplicationIds

`func (o *CreateManagementKey) HasAllowedApplicationIds() bool`

HasAllowedApplicationIds returns a boolean if a field has been set.

### SetAllowedApplicationIds

`func (o *CreateManagementKey) SetAllowedApplicationIds(v []int32)`

SetAllowedApplicationIds gets a reference to the given []int32 and assigns it to the AllowedApplicationIds field.

### GetEndpoints

`func (o *CreateManagementKey) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateManagementKey) GetEndpointsOk() ([]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndpoints

`func (o *CreateManagementKey) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpoints

`func (o *CreateManagementKey) SetEndpoints(v []Endpoint)`

SetEndpoints gets a reference to the given []Endpoint and assigns it to the Endpoints field.

### GetExpiryDate

`func (o *CreateManagementKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateManagementKey) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *CreateManagementKey) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *CreateManagementKey) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetName

`func (o *CreateManagementKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateManagementKey) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CreateManagementKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CreateManagementKey) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


