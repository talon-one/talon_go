# ManagementKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for management key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the management key expires. | 
**Endpoints** | Pointer to [**[]Endpoint**](Endpoint.md) | The list of endpoints that can be accessed with the key | 
**AllowedApplicationIds** | Pointer to **[]int64** | A list of Application IDs that you can access with the management key. An empty or missing list means the management key can be used for all Applications in the account.  | [optional] 
**Id** | Pointer to **int64** | ID of the management key. | 
**CreatedBy** | Pointer to **int64** | ID of the user who created it. | 
**AccountID** | Pointer to **int64** | ID of account the key is used for. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The date the management key was created. | 
**Disabled** | Pointer to **bool** | The management key is disabled (this property is set to &#x60;true&#x60;) when the user who created the key is disabled or deleted. | [optional] 
**LastUsed** | Pointer to [**time.Time**](time.Time.md) | The last time the management key was used. | [optional] 

## Methods

### NewManagementKey

`func NewManagementKey(name string, expiryDate time.Time, endpoints []Endpoint, id int64, createdBy int64, accountID int64, created time.Time, ) *ManagementKey`

NewManagementKey instantiates a new ManagementKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementKeyWithDefaults

`func NewManagementKeyWithDefaults() *ManagementKey`

NewManagementKeyWithDefaults instantiates a new ManagementKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDate

`func (o *ManagementKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ManagementKey) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ManagementKey) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetEndpoints

`func (o *ManagementKey) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ManagementKey) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ManagementKey) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.


### GetAllowedApplicationIds

`func (o *ManagementKey) GetAllowedApplicationIds() []int64`

GetAllowedApplicationIds returns the AllowedApplicationIds field if non-nil, zero value otherwise.

### GetAllowedApplicationIdsOk

`func (o *ManagementKey) GetAllowedApplicationIdsOk() (*[]int64, bool)`

GetAllowedApplicationIdsOk returns a tuple with the AllowedApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApplicationIds

`func (o *ManagementKey) SetAllowedApplicationIds(v []int64)`

SetAllowedApplicationIds sets AllowedApplicationIds field to given value.

### HasAllowedApplicationIds

`func (o *ManagementKey) HasAllowedApplicationIds() bool`

HasAllowedApplicationIds returns a boolean if a field has been set.

### GetId

`func (o *ManagementKey) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementKey) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementKey) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ManagementKey) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ManagementKey) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ManagementKey) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetAccountID

`func (o *ManagementKey) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *ManagementKey) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *ManagementKey) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.


### GetCreated

`func (o *ManagementKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ManagementKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ManagementKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDisabled

`func (o *ManagementKey) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ManagementKey) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ManagementKey) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ManagementKey) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLastUsed

`func (o *ManagementKey) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ManagementKey) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ManagementKey) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *ManagementKey) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


