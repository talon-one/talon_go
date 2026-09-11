# InlineResponse20056

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**InlineResponse20056Catalog**](inline_response_200_56_catalog.md) |  | 
**Loyalty** | Pointer to [**map[string]LoyaltyBalances**](LoyaltyBalances.md) | The customer&#39;s loyalty balances for the specified loyalty program. Returned only when &#x60;loyaltyProgramId&#x60; is provided together with &#x60;profileIntegrationId&#x60; or &#x60;loyaltyCardId&#x60;.  | [optional] 

## Methods

### NewInlineResponse20056

`func NewInlineResponse20056(catalog InlineResponse20056Catalog, ) *InlineResponse20056`

NewInlineResponse20056 instantiates a new InlineResponse20056 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20056WithDefaults

`func NewInlineResponse20056WithDefaults() *InlineResponse20056`

NewInlineResponse20056WithDefaults instantiates a new InlineResponse20056 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *InlineResponse20056) GetCatalog() InlineResponse20056Catalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *InlineResponse20056) GetCatalogOk() (*InlineResponse20056Catalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *InlineResponse20056) SetCatalog(v InlineResponse20056Catalog)`

SetCatalog sets Catalog field to given value.


### GetLoyalty

`func (o *InlineResponse20056) GetLoyalty() map[string]LoyaltyBalances`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *InlineResponse20056) GetLoyaltyOk() (*map[string]LoyaltyBalances, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *InlineResponse20056) SetLoyalty(v map[string]LoyaltyBalances)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *InlineResponse20056) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


