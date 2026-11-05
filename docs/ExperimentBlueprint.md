# ExperimentBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this experiment blueprint. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this experiment blueprint. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this experiment blueprint. | 
**Title** | Pointer to **string** | A short description of the experiment blueprint. | 
**Description** | Pointer to **string** | A longer, more detailed description of the experiment blueprint. | [optional] 
**Category** | Pointer to **string** | Category used to group experiment blueprints. | 
**Source** | Pointer to **string** | Indicates whether the experiment blueprint is custom or shipped by Talon.One. | 
**Variants** | Pointer to [**[]ExperimentBlueprintVariant**](ExperimentBlueprintVariant.md) | Array of experiment variant templates in this experiment blueprint. Variants only contain name, weight, ruleset and the primary flag. | 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Array of cart item filter templates in this experiment blueprint. Cart item filters only contain name (no description, as description is at the experiment blueprint level). | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the experiment blueprint was created. | 
**CreatedBy** | Pointer to **int64** | ID of the user who created the experiment blueprint. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the experiment blueprint was last updated. | [optional] 
**ModifiedBy** | Pointer to **int64** | ID of the user who last updated the experiment blueprint. | [optional] 

## Methods

### NewExperimentBlueprint

`func NewExperimentBlueprint(id int64, accountId int64, applicationId int64, title string, category string, source string, variants []ExperimentBlueprintVariant, cartItemFilters []CartItemFilterTemplate, created time.Time, createdBy int64, ) *ExperimentBlueprint`

NewExperimentBlueprint instantiates a new ExperimentBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentBlueprintWithDefaults

`func NewExperimentBlueprintWithDefaults() *ExperimentBlueprint`

NewExperimentBlueprintWithDefaults instantiates a new ExperimentBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExperimentBlueprint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentBlueprint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentBlueprint) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *ExperimentBlueprint) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExperimentBlueprint) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExperimentBlueprint) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetApplicationId

`func (o *ExperimentBlueprint) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ExperimentBlueprint) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ExperimentBlueprint) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetTitle

`func (o *ExperimentBlueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExperimentBlueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExperimentBlueprint) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ExperimentBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ExperimentBlueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExperimentBlueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExperimentBlueprint) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSource

`func (o *ExperimentBlueprint) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExperimentBlueprint) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExperimentBlueprint) SetSource(v string)`

SetSource sets Source field to given value.


### GetVariants

`func (o *ExperimentBlueprint) GetVariants() []ExperimentBlueprintVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentBlueprint) GetVariantsOk() (*[]ExperimentBlueprintVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentBlueprint) SetVariants(v []ExperimentBlueprintVariant)`

SetVariants sets Variants field to given value.


### GetCartItemFilters

`func (o *ExperimentBlueprint) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *ExperimentBlueprint) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *ExperimentBlueprint) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.


### GetCreated

`func (o *ExperimentBlueprint) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExperimentBlueprint) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExperimentBlueprint) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *ExperimentBlueprint) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExperimentBlueprint) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExperimentBlueprint) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetModified

`func (o *ExperimentBlueprint) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ExperimentBlueprint) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ExperimentBlueprint) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ExperimentBlueprint) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ExperimentBlueprint) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ExperimentBlueprint) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ExperimentBlueprint) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ExperimentBlueprint) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


