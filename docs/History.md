# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the historical price. | 
**ObservedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the price was observed. | 
**ContextId** | Pointer to **string** | Identifier of the relevant context at the time the price was observed (e.g. summer sale).  | 
**Price** | Pointer to **float32** | Price of the item. | 
**Metadata** | Pointer to [**BestPriorPriceMetadata**](BestPriorPriceMetadata.md) |  | 
**Target** | Pointer to [**map[string]interface{}**](.md) |  | 

## Methods

### NewHistory

`func NewHistory(id int64, observedAt time.Time, contextId string, price float32, metadata BestPriorPriceMetadata, target map[string]interface{}, ) *History`

NewHistory instantiates a new History object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryWithDefaults

`func NewHistoryWithDefaults() *History`

NewHistoryWithDefaults instantiates a new History object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *History) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *History) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *History) SetId(v int64)`

SetId sets Id field to given value.


### GetObservedAt

`func (o *History) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *History) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *History) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetContextId

`func (o *History) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *History) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *History) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetPrice

`func (o *History) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *History) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *History) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetMetadata

`func (o *History) GetMetadata() BestPriorPriceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *History) GetMetadataOk() (*BestPriorPriceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *History) SetMetadata(v BestPriorPriceMetadata)`

SetMetadata sets Metadata field to given value.


### GetTarget

`func (o *History) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *History) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *History) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


