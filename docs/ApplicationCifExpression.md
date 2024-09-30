# ApplicationCifExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CartItemFilterId** | Pointer to **int32** | The ID of the Application cart item filter. | [optional] 
**CreatedBy** | Pointer to **int32** | The ID of the user who created the Application cart item filter. | [optional] 
**Expression** | Pointer to [**[]interface{}**](map[string]interface{}.md) | Arbitrary additional JSON data associated with the Application cart item filter. | [optional] 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 

## Methods

### GetId

`func (o *ApplicationCifExpression) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCifExpression) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationCifExpression) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationCifExpression) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *ApplicationCifExpression) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationCifExpression) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationCifExpression) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationCifExpression) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCartItemFilterId

`func (o *ApplicationCifExpression) GetCartItemFilterId() int32`

GetCartItemFilterId returns the CartItemFilterId field if non-nil, zero value otherwise.

### GetCartItemFilterIdOk

`func (o *ApplicationCifExpression) GetCartItemFilterIdOk() (int32, bool)`

GetCartItemFilterIdOk returns a tuple with the CartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemFilterId

`func (o *ApplicationCifExpression) HasCartItemFilterId() bool`

HasCartItemFilterId returns a boolean if a field has been set.

### SetCartItemFilterId

`func (o *ApplicationCifExpression) SetCartItemFilterId(v int32)`

SetCartItemFilterId gets a reference to the given int32 and assigns it to the CartItemFilterId field.

### GetCreatedBy

`func (o *ApplicationCifExpression) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationCifExpression) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *ApplicationCifExpression) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *ApplicationCifExpression) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetExpression

`func (o *ApplicationCifExpression) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ApplicationCifExpression) GetExpressionOk() ([]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *ApplicationCifExpression) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *ApplicationCifExpression) SetExpression(v []interface{})`

SetExpression gets a reference to the given []map[string]interface{} and assigns it to the Expression field.

### GetApplicationId

`func (o *ApplicationCifExpression) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationCifExpression) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationCifExpression) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationCifExpression) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


