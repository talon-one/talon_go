# InlineResponse20059Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | Whether more pages exist after the current one. | 
**Data** | Pointer to [**[]RewardCatalogItem**](RewardCatalogItem.md) |  | 

## Methods

### NewInlineResponse20059Catalog

`func NewInlineResponse20059Catalog(hasMore bool, data []RewardCatalogItem, ) *InlineResponse20059Catalog`

NewInlineResponse20059Catalog instantiates a new InlineResponse20059Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20059CatalogWithDefaults

`func NewInlineResponse20059CatalogWithDefaults() *InlineResponse20059Catalog`

NewInlineResponse20059CatalogWithDefaults instantiates a new InlineResponse20059Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *InlineResponse20059Catalog) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InlineResponse20059Catalog) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InlineResponse20059Catalog) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *InlineResponse20059Catalog) GetData() []RewardCatalogItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineResponse20059Catalog) GetDataOk() (*[]RewardCatalogItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineResponse20059Catalog) SetData(v []RewardCatalogItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


