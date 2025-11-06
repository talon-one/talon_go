# UpdateApplicationCIF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the Application cart item filter. | [optional] 
**ActiveExpressionId** | Pointer to **int64** | The ID of the expression that the Application cart item filter uses. | [optional] 
**ModifiedBy** | Pointer to **int64** | The ID of the user who last updated the Application cart item filter. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the Application cart item filter. | [optional] 

## Methods

### NewUpdateApplicationCIF

`func NewUpdateApplicationCIF() *UpdateApplicationCIF`

NewUpdateApplicationCIF instantiates a new UpdateApplicationCIF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationCIFWithDefaults

`func NewUpdateApplicationCIFWithDefaults() *UpdateApplicationCIF`

NewUpdateApplicationCIFWithDefaults instantiates a new UpdateApplicationCIF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateApplicationCIF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplicationCIF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApplicationCIF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApplicationCIF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActiveExpressionId

`func (o *UpdateApplicationCIF) GetActiveExpressionId() int64`

GetActiveExpressionId returns the ActiveExpressionId field if non-nil, zero value otherwise.

### GetActiveExpressionIdOk

`func (o *UpdateApplicationCIF) GetActiveExpressionIdOk() (*int64, bool)`

GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExpressionId

`func (o *UpdateApplicationCIF) SetActiveExpressionId(v int64)`

SetActiveExpressionId sets ActiveExpressionId field to given value.

### HasActiveExpressionId

`func (o *UpdateApplicationCIF) HasActiveExpressionId() bool`

HasActiveExpressionId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *UpdateApplicationCIF) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *UpdateApplicationCIF) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *UpdateApplicationCIF) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *UpdateApplicationCIF) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModified

`func (o *UpdateApplicationCIF) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *UpdateApplicationCIF) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *UpdateApplicationCIF) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *UpdateApplicationCIF) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


