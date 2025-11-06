# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of item. | [optional] 
**Sku** | Pointer to **string** | Stock keeping unit of item. | 
**Quantity** | Pointer to **int64** | Number of units of this item. Due to [cart item flattening](https://docs.talon.one/docs/product/rules/understanding-cart-item-flattening), if you provide a quantity greater than 1, the item will be split in as many items as the provided quantity. This will impact the number of **per-item** effects triggered from your campaigns.  | 
**ReturnedQuantity** | Pointer to **int64** | Number of returned items, calculated internally based on returns of this item. | [optional] 
**RemainingQuantity** | Pointer to **int64** | Remaining quantity of the item, calculated internally based on returns of this item. | [optional] 
**Price** | Pointer to **float32** | Price of the item in the currency defined by your Application. This field is required if this item is not part of a [catalog](https://docs.talon.one/docs/product/account/dev-tools/managing-cart-item-catalogs). If it is part of a catalog, setting a price here overrides the price from the catalog.  | [optional] 
**Category** | Pointer to **string** | Type, group or model of the item. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Weight** | Pointer to **float32** | Weight of item in grams. | [optional] 
**Height** | Pointer to **float32** | Height of item in mm. | [optional] 
**Width** | Pointer to **float32** | Width of item in mm. | [optional] 
**Length** | Pointer to **float32** | Length of item in mm. | [optional] 
**Position** | Pointer to **float32** | Position of the Cart Item in the Cart (calculated internally). | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Use this property to set a value for the attributes of your choice. [Attributes](https://docs.talon.one/docs/dev/concepts/attributes) represent any information to attach to this cart item.  Custom _cart item_ attributes must be created in the Campaign Manager before you set them with this property.  **Note:** Any previously defined attributes that you do not include in the array will be removed.  | [optional] 
**AdditionalCosts** | Pointer to [**map[string]AdditionalCost**](AdditionalCost.md) | Use this property to set a value for the additional costs of this item, such as a shipping cost. They must be created in the Campaign Manager before you set them with this property. See [Managing additional costs](https://docs.talon.one/docs/product/account/dev-tools/managing-additional-costs).  | [optional] 
**CatalogItemID** | Pointer to **int64** | The catalog item ID. | [optional] 
**SelectedPriceType** | Pointer to **string** | The selected price type for this cart item (e.g. the price for members only). | [optional] 
**AdjustmentReferenceId** | Pointer to **string** | The reference ID of the selected price adjustment for this cart item. Only returned if the selected price resulted from a price adjustment. | [optional] 
**AdjustmentEffectiveFrom** | Pointer to [**time.Time**](time.Time.md) | The date and time from which the price adjustment is effective. Only returned if the selected price resulted from a price adjustment that contains this field. | [optional] 
**AdjustmentEffectiveUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time until which the price adjustment is effective. Only returned if the selected price resulted from a price adjustment that contains this field. | [optional] 
**Prices** | Pointer to [**map[string]PriceDetail**](PriceDetail.md) | A map of keys and values representing the price types and related price adjustment details for this cart item. The keys correspond to the &#x60;priceType&#x60; names.  | [optional] 

## Methods

### NewCartItem

`func NewCartItem(sku string, quantity int64, ) *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CartItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSku

`func (o *CartItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CartItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CartItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetQuantity

`func (o *CartItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetReturnedQuantity

`func (o *CartItem) GetReturnedQuantity() int64`

GetReturnedQuantity returns the ReturnedQuantity field if non-nil, zero value otherwise.

### GetReturnedQuantityOk

`func (o *CartItem) GetReturnedQuantityOk() (*int64, bool)`

GetReturnedQuantityOk returns a tuple with the ReturnedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedQuantity

`func (o *CartItem) SetReturnedQuantity(v int64)`

SetReturnedQuantity sets ReturnedQuantity field to given value.

### HasReturnedQuantity

`func (o *CartItem) HasReturnedQuantity() bool`

HasReturnedQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *CartItem) GetRemainingQuantity() int64`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *CartItem) GetRemainingQuantityOk() (*int64, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *CartItem) SetRemainingQuantity(v int64)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *CartItem) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *CartItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CartItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CartItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CartItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCategory

`func (o *CartItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CartItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CartItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CartItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetProduct

`func (o *CartItem) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CartItem) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CartItem) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CartItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetWeight

`func (o *CartItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CartItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CartItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CartItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetHeight

`func (o *CartItem) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CartItem) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CartItem) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CartItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *CartItem) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CartItem) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CartItem) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CartItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *CartItem) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CartItem) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *CartItem) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *CartItem) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPosition

`func (o *CartItem) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CartItem) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CartItem) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CartItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAttributes

`func (o *CartItem) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CartItem) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CartItem) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CartItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAdditionalCosts

`func (o *CartItem) GetAdditionalCosts() map[string]AdditionalCost`

GetAdditionalCosts returns the AdditionalCosts field if non-nil, zero value otherwise.

### GetAdditionalCostsOk

`func (o *CartItem) GetAdditionalCostsOk() (*map[string]AdditionalCost, bool)`

GetAdditionalCostsOk returns a tuple with the AdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCosts

`func (o *CartItem) SetAdditionalCosts(v map[string]AdditionalCost)`

SetAdditionalCosts sets AdditionalCosts field to given value.

### HasAdditionalCosts

`func (o *CartItem) HasAdditionalCosts() bool`

HasAdditionalCosts returns a boolean if a field has been set.

### GetCatalogItemID

`func (o *CartItem) GetCatalogItemID() int64`

GetCatalogItemID returns the CatalogItemID field if non-nil, zero value otherwise.

### GetCatalogItemIDOk

`func (o *CartItem) GetCatalogItemIDOk() (*int64, bool)`

GetCatalogItemIDOk returns a tuple with the CatalogItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemID

`func (o *CartItem) SetCatalogItemID(v int64)`

SetCatalogItemID sets CatalogItemID field to given value.

### HasCatalogItemID

`func (o *CartItem) HasCatalogItemID() bool`

HasCatalogItemID returns a boolean if a field has been set.

### GetSelectedPriceType

`func (o *CartItem) GetSelectedPriceType() string`

GetSelectedPriceType returns the SelectedPriceType field if non-nil, zero value otherwise.

### GetSelectedPriceTypeOk

`func (o *CartItem) GetSelectedPriceTypeOk() (*string, bool)`

GetSelectedPriceTypeOk returns a tuple with the SelectedPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPriceType

`func (o *CartItem) SetSelectedPriceType(v string)`

SetSelectedPriceType sets SelectedPriceType field to given value.

### HasSelectedPriceType

`func (o *CartItem) HasSelectedPriceType() bool`

HasSelectedPriceType returns a boolean if a field has been set.

### GetAdjustmentReferenceId

`func (o *CartItem) GetAdjustmentReferenceId() string`

GetAdjustmentReferenceId returns the AdjustmentReferenceId field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIdOk

`func (o *CartItem) GetAdjustmentReferenceIdOk() (*string, bool)`

GetAdjustmentReferenceIdOk returns a tuple with the AdjustmentReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReferenceId

`func (o *CartItem) SetAdjustmentReferenceId(v string)`

SetAdjustmentReferenceId sets AdjustmentReferenceId field to given value.

### HasAdjustmentReferenceId

`func (o *CartItem) HasAdjustmentReferenceId() bool`

HasAdjustmentReferenceId returns a boolean if a field has been set.

### GetAdjustmentEffectiveFrom

`func (o *CartItem) GetAdjustmentEffectiveFrom() time.Time`

GetAdjustmentEffectiveFrom returns the AdjustmentEffectiveFrom field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveFromOk

`func (o *CartItem) GetAdjustmentEffectiveFromOk() (*time.Time, bool)`

GetAdjustmentEffectiveFromOk returns a tuple with the AdjustmentEffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentEffectiveFrom

`func (o *CartItem) SetAdjustmentEffectiveFrom(v time.Time)`

SetAdjustmentEffectiveFrom sets AdjustmentEffectiveFrom field to given value.

### HasAdjustmentEffectiveFrom

`func (o *CartItem) HasAdjustmentEffectiveFrom() bool`

HasAdjustmentEffectiveFrom returns a boolean if a field has been set.

### GetAdjustmentEffectiveUntil

`func (o *CartItem) GetAdjustmentEffectiveUntil() time.Time`

GetAdjustmentEffectiveUntil returns the AdjustmentEffectiveUntil field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveUntilOk

`func (o *CartItem) GetAdjustmentEffectiveUntilOk() (*time.Time, bool)`

GetAdjustmentEffectiveUntilOk returns a tuple with the AdjustmentEffectiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentEffectiveUntil

`func (o *CartItem) SetAdjustmentEffectiveUntil(v time.Time)`

SetAdjustmentEffectiveUntil sets AdjustmentEffectiveUntil field to given value.

### HasAdjustmentEffectiveUntil

`func (o *CartItem) HasAdjustmentEffectiveUntil() bool`

HasAdjustmentEffectiveUntil returns a boolean if a field has been set.

### GetPrices

`func (o *CartItem) GetPrices() map[string]PriceDetail`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *CartItem) GetPricesOk() (*map[string]PriceDetail, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *CartItem) SetPrices(v map[string]PriceDetail)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *CartItem) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


