# APIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Short description of the problem. | 
**Details** | Pointer to **string** | Longer description of this specific instance of the problem. | [optional] 
**Source** | Pointer to [**ErrorSource**](ErrorSource.md) |  | 

## Methods

### NewAPIError

`func NewAPIError(title string, source ErrorSource, ) *APIError`

NewAPIError instantiates a new APIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIErrorWithDefaults

`func NewAPIErrorWithDefaults() *APIError`

NewAPIErrorWithDefaults instantiates a new APIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *APIError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *APIError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *APIError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetails

`func (o *APIError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *APIError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *APIError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *APIError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSource

`func (o *APIError) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *APIError) GetSourceOk() (*ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *APIError) SetSource(v ErrorSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


