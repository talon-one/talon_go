# ApplicationCIFExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CartItemFilterId** | Pointer to **int64** | The ID of the Application cart item filter. | [optional] 
**CreatedBy** | Pointer to **int64** | The ID of the user who created the Application cart item filter. | [optional] 
**Expression** | Pointer to **[]map[string]interface{}** | Arbitrary additional JSON data associated with the Application cart item filter. | [optional] 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 

## Methods

### NewApplicationCIFExpression

`func NewApplicationCIFExpression(id int64, created time.Time, applicationId int64, ) *ApplicationCIFExpression`

NewApplicationCIFExpression instantiates a new ApplicationCIFExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCIFExpressionWithDefaults

`func NewApplicationCIFExpressionWithDefaults() *ApplicationCIFExpression`

NewApplicationCIFExpressionWithDefaults instantiates a new ApplicationCIFExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCIFExpression) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCIFExpression) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCIFExpression) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *ApplicationCIFExpression) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationCIFExpression) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationCIFExpression) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCartItemFilterId

`func (o *ApplicationCIFExpression) GetCartItemFilterId() int64`

GetCartItemFilterId returns the CartItemFilterId field if non-nil, zero value otherwise.

### GetCartItemFilterIdOk

`func (o *ApplicationCIFExpression) GetCartItemFilterIdOk() (*int64, bool)`

GetCartItemFilterIdOk returns a tuple with the CartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilterId

`func (o *ApplicationCIFExpression) SetCartItemFilterId(v int64)`

SetCartItemFilterId sets CartItemFilterId field to given value.

### HasCartItemFilterId

`func (o *ApplicationCIFExpression) HasCartItemFilterId() bool`

HasCartItemFilterId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApplicationCIFExpression) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationCIFExpression) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApplicationCIFExpression) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApplicationCIFExpression) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExpression

`func (o *ApplicationCIFExpression) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ApplicationCIFExpression) GetExpressionOk() (*[]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ApplicationCIFExpression) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ApplicationCIFExpression) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationCIFExpression) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationCIFExpression) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationCIFExpression) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


