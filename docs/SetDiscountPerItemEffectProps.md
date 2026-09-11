# SetDiscountPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The description of this discount. &#x60;#number&#x60; is equal to the &#x60;position&#x60; property. | 
**Value** | Pointer to **float32** | The monetary value of the effective discount applied to the item. | 
**Position** | Pointer to **float32** | The index of the item in the &#x60;cartItem&#x60; object on which this discount should be applied. | 
**SubPosition** | Pointer to **float32** | The index of the item unit in its line item. | [optional] 
**DesiredValue** | Pointer to **float32** | _(Partial discounts enabled only)_ The monetary value of the discount to be applied to the item without considering budget limitations. | [optional] 
**Scope** | Pointer to **string** | What the discount applies to. Possible values:  - &#x60;price&#x60;: discount on the price of the item. - &#x60;additionalCosts&#x60;: discount on the [additional cost](https://docs.talon.one/docs/product/account/dev-tools/manage-additional-costs) of the item. - &#x60;itemTotal&#x60;: discount on the sum of price + additional cost of the item. | [optional] 
**TotalDiscount** | Pointer to **float32** | _(Pro rata discounts only)_ The monetary value of the total effective discount | [optional] 
**DesiredTotalDiscount** | Pointer to **float32** | _(Pro rata discounts only)_ The monetary value of the total discount to be applied without considering budget limitations | [optional] 
**BundleIndex** | Pointer to **int64** | _(Discounts with bundles only)_ The position of the specific item bundle in the list of bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | _(Discounts with bundles only)_ The name of the bundle definition. | [optional] 
**TargetedItemPosition** | Pointer to **float32** | _(Discounting individual item in bundles only)_ The index of the targeted bundle item on which the applied discount is based. | [optional] 
**TargetedItemSubPosition** | Pointer to **float32** | _(Discounting individual item in bundles only)_ The sub-position of the targeted bundle item on which the applied discount is based. | [optional] 
**ExcludedFromPriceHistory** | Pointer to **bool** | When set to &#x60;true&#x60;, the applied discount is excluded from the item&#39;s price history. | [optional] 

## Methods

### NewSetDiscountPerItemEffectProps

`func NewSetDiscountPerItemEffectProps(name string, value float32, position float32, ) *SetDiscountPerItemEffectProps`

NewSetDiscountPerItemEffectProps instantiates a new SetDiscountPerItemEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDiscountPerItemEffectPropsWithDefaults

`func NewSetDiscountPerItemEffectPropsWithDefaults() *SetDiscountPerItemEffectProps`

NewSetDiscountPerItemEffectPropsWithDefaults instantiates a new SetDiscountPerItemEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetDiscountPerItemEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountPerItemEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetDiscountPerItemEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *SetDiscountPerItemEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountPerItemEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetDiscountPerItemEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetPosition

`func (o *SetDiscountPerItemEffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetDiscountPerItemEffectProps) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SetDiscountPerItemEffectProps) SetPosition(v float32)`

SetPosition sets Position field to given value.


### GetSubPosition

`func (o *SetDiscountPerItemEffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *SetDiscountPerItemEffectProps) GetSubPositionOk() (*float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosition

`func (o *SetDiscountPerItemEffectProps) SetSubPosition(v float32)`

SetSubPosition sets SubPosition field to given value.

### HasSubPosition

`func (o *SetDiscountPerItemEffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### GetDesiredValue

`func (o *SetDiscountPerItemEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountPerItemEffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *SetDiscountPerItemEffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *SetDiscountPerItemEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### GetScope

`func (o *SetDiscountPerItemEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SetDiscountPerItemEffectProps) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SetDiscountPerItemEffectProps) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SetDiscountPerItemEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *SetDiscountPerItemEffectProps) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *SetDiscountPerItemEffectProps) GetTotalDiscountOk() (*float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *SetDiscountPerItemEffectProps) SetTotalDiscount(v float32)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *SetDiscountPerItemEffectProps) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) GetDesiredTotalDiscount() float32`

GetDesiredTotalDiscount returns the DesiredTotalDiscount field if non-nil, zero value otherwise.

### GetDesiredTotalDiscountOk

`func (o *SetDiscountPerItemEffectProps) GetDesiredTotalDiscountOk() (*float32, bool)`

GetDesiredTotalDiscountOk returns a tuple with the DesiredTotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) SetDesiredTotalDiscount(v float32)`

SetDesiredTotalDiscount sets DesiredTotalDiscount field to given value.

### HasDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) HasDesiredTotalDiscount() bool`

HasDesiredTotalDiscount returns a boolean if a field has been set.

### GetBundleIndex

`func (o *SetDiscountPerItemEffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *SetDiscountPerItemEffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *SetDiscountPerItemEffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *SetDiscountPerItemEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *SetDiscountPerItemEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *SetDiscountPerItemEffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *SetDiscountPerItemEffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *SetDiscountPerItemEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetTargetedItemPosition

`func (o *SetDiscountPerItemEffectProps) GetTargetedItemPosition() float32`

GetTargetedItemPosition returns the TargetedItemPosition field if non-nil, zero value otherwise.

### GetTargetedItemPositionOk

`func (o *SetDiscountPerItemEffectProps) GetTargetedItemPositionOk() (*float32, bool)`

GetTargetedItemPositionOk returns a tuple with the TargetedItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedItemPosition

`func (o *SetDiscountPerItemEffectProps) SetTargetedItemPosition(v float32)`

SetTargetedItemPosition sets TargetedItemPosition field to given value.

### HasTargetedItemPosition

`func (o *SetDiscountPerItemEffectProps) HasTargetedItemPosition() bool`

HasTargetedItemPosition returns a boolean if a field has been set.

### GetTargetedItemSubPosition

`func (o *SetDiscountPerItemEffectProps) GetTargetedItemSubPosition() float32`

GetTargetedItemSubPosition returns the TargetedItemSubPosition field if non-nil, zero value otherwise.

### GetTargetedItemSubPositionOk

`func (o *SetDiscountPerItemEffectProps) GetTargetedItemSubPositionOk() (*float32, bool)`

GetTargetedItemSubPositionOk returns a tuple with the TargetedItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedItemSubPosition

`func (o *SetDiscountPerItemEffectProps) SetTargetedItemSubPosition(v float32)`

SetTargetedItemSubPosition sets TargetedItemSubPosition field to given value.

### HasTargetedItemSubPosition

`func (o *SetDiscountPerItemEffectProps) HasTargetedItemSubPosition() bool`

HasTargetedItemSubPosition returns a boolean if a field has been set.

### GetExcludedFromPriceHistory

`func (o *SetDiscountPerItemEffectProps) GetExcludedFromPriceHistory() bool`

GetExcludedFromPriceHistory returns the ExcludedFromPriceHistory field if non-nil, zero value otherwise.

### GetExcludedFromPriceHistoryOk

`func (o *SetDiscountPerItemEffectProps) GetExcludedFromPriceHistoryOk() (*bool, bool)`

GetExcludedFromPriceHistoryOk returns a tuple with the ExcludedFromPriceHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromPriceHistory

`func (o *SetDiscountPerItemEffectProps) SetExcludedFromPriceHistory(v bool)`

SetExcludedFromPriceHistory sets ExcludedFromPriceHistory field to given value.

### HasExcludedFromPriceHistory

`func (o *SetDiscountPerItemEffectProps) HasExcludedFromPriceHistory() bool`

HasExcludedFromPriceHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


