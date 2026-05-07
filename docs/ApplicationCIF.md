# ApplicationCIF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The name of the Application cart item filter used in API requests. | 
**Description** | Pointer to **string** | A short description of the Application cart item filter. | [optional] 
**ActiveExpressionId** | Pointer to **int64** | The ID of the expression that the Application cart item filter uses. | [optional] 
**ModifiedBy** | Pointer to **int64** | The ID of the user who last updated the Application cart item filter. | [optional] 
**CreatedBy** | Pointer to **int64** | The ID of the user who created the Application cart item filter. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the Application cart item filter. | [optional] 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 

## Methods

### NewApplicationCIF

`func NewApplicationCIF(id int64, created time.Time, name string, applicationId int64, ) *ApplicationCIF`

NewApplicationCIF instantiates a new ApplicationCIF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCIFWithDefaults

`func NewApplicationCIFWithDefaults() *ApplicationCIF`

NewApplicationCIFWithDefaults instantiates a new ApplicationCIF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCIF) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCIF) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCIF) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ApplicationCIF) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationCIF) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationCIF) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *ApplicationCIF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCIF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCIF) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationCIF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCIF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCIF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationCIF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActiveExpressionId

`func (o *ApplicationCIF) GetActiveExpressionId() int64`

GetActiveExpressionId returns the ActiveExpressionId field if non-nil, zero value otherwise.

### GetActiveExpressionIdOk

`func (o *ApplicationCIF) GetActiveExpressionIdOk() (*int64, bool)`

GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExpressionId

`func (o *ApplicationCIF) SetActiveExpressionId(v int64)`

SetActiveExpressionId sets ActiveExpressionId field to given value.

### HasActiveExpressionId

`func (o *ApplicationCIF) HasActiveExpressionId() bool`

HasActiveExpressionId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ApplicationCIF) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ApplicationCIF) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ApplicationCIF) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ApplicationCIF) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApplicationCIF) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationCIF) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApplicationCIF) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApplicationCIF) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModified

`func (o *ApplicationCIF) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ApplicationCIF) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ApplicationCIF) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ApplicationCIF) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationCIF) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationCIF) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationCIF) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


