# BestPriorPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | sku | 
**ObservedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the best price was observed. | 
**ContextId** | Pointer to **string** | The context ID of the context active at the time of observation.  | 
**Price** | Pointer to **float32** | Price of the item. | 
**Metadata** | Pointer to [**BestPriorPriceMetadata**](BestPriorPriceMetadata.md) |  | 
**Target** | Pointer to [**map[string]interface{}**](.md) |  | 

## Methods

### NewBestPriorPrice

`func NewBestPriorPrice(sku string, observedAt time.Time, contextId string, price float32, metadata BestPriorPriceMetadata, target map[string]interface{}, ) *BestPriorPrice`

NewBestPriorPrice instantiates a new BestPriorPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorPriceWithDefaults

`func NewBestPriorPriceWithDefaults() *BestPriorPrice`

NewBestPriorPriceWithDefaults instantiates a new BestPriorPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *BestPriorPrice) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *BestPriorPrice) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *BestPriorPrice) SetSku(v string)`

SetSku sets Sku field to given value.


### GetObservedAt

`func (o *BestPriorPrice) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *BestPriorPrice) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *BestPriorPrice) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetContextId

`func (o *BestPriorPrice) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *BestPriorPrice) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *BestPriorPrice) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetPrice

`func (o *BestPriorPrice) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BestPriorPrice) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BestPriorPrice) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetMetadata

`func (o *BestPriorPrice) GetMetadata() BestPriorPriceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BestPriorPrice) GetMetadataOk() (*BestPriorPriceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BestPriorPrice) SetMetadata(v BestPriorPriceMetadata)`

SetMetadata sets Metadata field to given value.


### GetTarget

`func (o *BestPriorPrice) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BestPriorPrice) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BestPriorPrice) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


