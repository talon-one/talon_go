# CreateMCPKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the MCP key. | 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date the MCP key expires. | 

## Methods

### NewCreateMCPKey

`func NewCreateMCPKey(name string, expiryDate time.Time, ) *CreateMCPKey`

NewCreateMCPKey instantiates a new CreateMCPKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMCPKeyWithDefaults

`func NewCreateMCPKeyWithDefaults() *CreateMCPKey`

NewCreateMCPKeyWithDefaults instantiates a new CreateMCPKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMCPKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMCPKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMCPKey) SetName(v string)`

SetName sets Name field to given value.


### GetExpiryDate

`func (o *CreateMCPKey) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateMCPKey) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CreateMCPKey) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


