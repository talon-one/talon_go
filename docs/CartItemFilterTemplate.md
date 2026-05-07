# CartItemFilterTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Application cart item filter. | 
**Expression** | Pointer to **[]map[string]interface{}** | The Talang expression for the cart item filter. | 

## Methods

### NewCartItemFilterTemplate

`func NewCartItemFilterTemplate(name string, expression []map[string]interface{}, ) *CartItemFilterTemplate`

NewCartItemFilterTemplate instantiates a new CartItemFilterTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemFilterTemplateWithDefaults

`func NewCartItemFilterTemplateWithDefaults() *CartItemFilterTemplate`

NewCartItemFilterTemplateWithDefaults instantiates a new CartItemFilterTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CartItemFilterTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItemFilterTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartItemFilterTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *CartItemFilterTemplate) GetExpression() []map[string]interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CartItemFilterTemplate) GetExpressionOk() (*[]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CartItemFilterTemplate) SetExpression(v []map[string]interface{})`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


