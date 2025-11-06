# CreateManagementKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for management key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the management key expires. | 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | The list of endpoints that can be accessed with the key | 
**AllowedApplicationIds** | Pointer to **[]int64** | A list of Application IDs that you can access with the management key. An empty or missing list means the management key can be used for all Applications in the account.  | [optional] 

## Methods

### NewCreateManagementKey

`func NewCreateManagementKey(name string, expiryDate time.Time, endpoints []Endpoint, ) *CreateManagementKey`

NewCreateManagementKey instantiates a new CreateManagementKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManagementKeyWithDefaults

`func NewCreateManagementKeyWithDefaults() *CreateManagementKey`

NewCreateManagementKeyWithDefaults instantiates a new CreateManagementKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateManagementKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateManagementKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateManagementKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDate

`func (o *CreateManagementKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateManagementKey) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CreateManagementKey) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetEndpoints

`func (o *CreateManagementKey) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateManagementKey) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateManagementKey) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.


### GetAllowedApplicationIds

`func (o *CreateManagementKey) GetAllowedApplicationIds() []int64`

GetAllowedApplicationIds returns the AllowedApplicationIds field if non-nil, zero value otherwise.

### GetAllowedApplicationIdsOk

`func (o *CreateManagementKey) GetAllowedApplicationIdsOk() (*[]int64, bool)`

GetAllowedApplicationIdsOk returns a tuple with the AllowedApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApplicationIds

`func (o *CreateManagementKey) SetAllowedApplicationIds(v []int64)`

SetAllowedApplicationIds sets AllowedApplicationIds field to given value.

### HasAllowedApplicationIds

`func (o *CreateManagementKey) HasAllowedApplicationIds() bool`

HasAllowedApplicationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


