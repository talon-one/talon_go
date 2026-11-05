# InlineResponse20059

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**InlineResponse20059Catalog**](inline_response_200_59_catalog.md) |  | 
**Loyalty** | Pointer to [**map[string]LoyaltyBalances**](LoyaltyBalances.md) | The customer&#39;s loyalty balances for the specified loyalty program. Returned only when &#x60;loyaltyProgramId&#x60; is provided together with &#x60;profileIntegrationId&#x60; or &#x60;loyaltyCardId&#x60;.  | [optional] 

## Methods

### NewInlineResponse20059

`func NewInlineResponse20059(catalog InlineResponse20059Catalog, ) *InlineResponse20059`

NewInlineResponse20059 instantiates a new InlineResponse20059 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20059WithDefaults

`func NewInlineResponse20059WithDefaults() *InlineResponse20059`

NewInlineResponse20059WithDefaults instantiates a new InlineResponse20059 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *InlineResponse20059) GetCatalog() InlineResponse20059Catalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *InlineResponse20059) GetCatalogOk() (*InlineResponse20059Catalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *InlineResponse20059) SetCatalog(v InlineResponse20059Catalog)`

SetCatalog sets Catalog field to given value.


### GetLoyalty

`func (o *InlineResponse20059) GetLoyalty() map[string]LoyaltyBalances`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *InlineResponse20059) GetLoyaltyOk() (*map[string]LoyaltyBalances, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *InlineResponse20059) SetLoyalty(v map[string]LoyaltyBalances)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *InlineResponse20059) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


