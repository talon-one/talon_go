# PlaceholderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the placeholder. | 
**Type** | Pointer to **string** | The type of the value assigned to the placeholder. | 
**Value** | Pointer to **[]map[string]interface{}** | The current value of the placeholder. | 

## Methods

### NewPlaceholderDetails

`func NewPlaceholderDetails(name string, type_ string, value []map[string]interface{}, ) *PlaceholderDetails`

NewPlaceholderDetails instantiates a new PlaceholderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceholderDetailsWithDefaults

`func NewPlaceholderDetailsWithDefaults() *PlaceholderDetails`

NewPlaceholderDetailsWithDefaults instantiates a new PlaceholderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlaceholderDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaceholderDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaceholderDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PlaceholderDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceholderDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceholderDetails) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *PlaceholderDetails) GetValue() []map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlaceholderDetails) GetValueOk() (*[]map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlaceholderDetails) SetValue(v []map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


