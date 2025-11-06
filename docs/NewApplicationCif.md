# NewApplicationCIF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Application cart item filter used in API requests. | 
**Description** | Pointer to **string** | A short description of the Application cart item filter. | [optional] 
**ActiveExpressionId** | Pointer to **int64** | The ID of the expression that the Application cart item filter uses. | [optional] 
**ModifiedBy** | Pointer to **int64** | The ID of the user who last updated the Application cart item filter. | [optional] 
**CreatedBy** | Pointer to **int64** | The ID of the user who created the Application cart item filter. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the Application cart item filter. | [optional] 

## Methods

### NewNewApplicationCIF

`func NewNewApplicationCIF(name string, ) *NewApplicationCIF`

NewNewApplicationCIF instantiates a new NewApplicationCIF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApplicationCIFWithDefaults

`func NewNewApplicationCIFWithDefaults() *NewApplicationCIF`

NewNewApplicationCIFWithDefaults instantiates a new NewApplicationCIF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewApplicationCIF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApplicationCIF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewApplicationCIF) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewApplicationCIF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewApplicationCIF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewApplicationCIF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewApplicationCIF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActiveExpressionId

`func (o *NewApplicationCIF) GetActiveExpressionId() int64`

GetActiveExpressionId returns the ActiveExpressionId field if non-nil, zero value otherwise.

### GetActiveExpressionIdOk

`func (o *NewApplicationCIF) GetActiveExpressionIdOk() (*int64, bool)`

GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExpressionId

`func (o *NewApplicationCIF) SetActiveExpressionId(v int64)`

SetActiveExpressionId sets ActiveExpressionId field to given value.

### HasActiveExpressionId

`func (o *NewApplicationCIF) HasActiveExpressionId() bool`

HasActiveExpressionId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *NewApplicationCIF) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *NewApplicationCIF) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *NewApplicationCIF) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *NewApplicationCIF) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *NewApplicationCIF) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NewApplicationCIF) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NewApplicationCIF) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *NewApplicationCIF) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModified

`func (o *NewApplicationCIF) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NewApplicationCIF) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NewApplicationCIF) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NewApplicationCIF) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


