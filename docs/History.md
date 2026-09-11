# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the historical price. | 
**ObservedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the price was observed. | 
**ContextIds** | Pointer to **[]string** | The identifiers of the relevant context at the time the price was observed. Includes the context IDs of any price adjustments and of the campaigns that influenced the final price.  | 
**Price** | Pointer to **float32** | Price of the item. | 
**Metadata** | Pointer to [**BestPriorPriceMetadata**](BestPriorPriceMetadata.md) |  | 
**Target** | Pointer to [**map[string]interface{}**](.md) |  | 
**ExcludedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the historical price ID was excluded. | [optional] 
**ExclusionReason** | Pointer to **string** | The reason for excluding this historical price ID. | [optional] 

## Methods

### NewHistory

`func NewHistory(id int64, observedAt time.Time, contextIds []string, price float32, metadata BestPriorPriceMetadata, target map[string]interface{}, ) *History`

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


### GetContextIds

`func (o *History) GetContextIds() []string`

GetContextIds returns the ContextIds field if non-nil, zero value otherwise.

### GetContextIdsOk

`func (o *History) GetContextIdsOk() (*[]string, bool)`

GetContextIdsOk returns a tuple with the ContextIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextIds

`func (o *History) SetContextIds(v []string)`

SetContextIds sets ContextIds field to given value.


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


### GetExcludedAt

`func (o *History) GetExcludedAt() time.Time`

GetExcludedAt returns the ExcludedAt field if non-nil, zero value otherwise.

### GetExcludedAtOk

`func (o *History) GetExcludedAtOk() (*time.Time, bool)`

GetExcludedAtOk returns a tuple with the ExcludedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAt

`func (o *History) SetExcludedAt(v time.Time)`

SetExcludedAt sets ExcludedAt field to given value.

### HasExcludedAt

`func (o *History) HasExcludedAt() bool`

HasExcludedAt returns a boolean if a field has been set.

### GetExclusionReason

`func (o *History) GetExclusionReason() string`

GetExclusionReason returns the ExclusionReason field if non-nil, zero value otherwise.

### GetExclusionReasonOk

`func (o *History) GetExclusionReasonOk() (*string, bool)`

GetExclusionReasonOk returns a tuple with the ExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionReason

`func (o *History) SetExclusionReason(v string)`

SetExclusionReason sets ExclusionReason field to given value.

### HasExclusionReason

`func (o *History) HasExclusionReason() bool`

HasExclusionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


