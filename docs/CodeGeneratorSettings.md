# CodeGeneratorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code.  | 
**CouponPattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | 

## Methods

### NewCodeGeneratorSettings

`func NewCodeGeneratorSettings(validCharacters []string, couponPattern string, ) *CodeGeneratorSettings`

NewCodeGeneratorSettings instantiates a new CodeGeneratorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeGeneratorSettingsWithDefaults

`func NewCodeGeneratorSettingsWithDefaults() *CodeGeneratorSettings`

NewCodeGeneratorSettingsWithDefaults instantiates a new CodeGeneratorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidCharacters

`func (o *CodeGeneratorSettings) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *CodeGeneratorSettings) GetValidCharactersOk() (*[]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *CodeGeneratorSettings) SetValidCharacters(v []string)`

SetValidCharacters sets ValidCharacters field to given value.


### GetCouponPattern

`func (o *CodeGeneratorSettings) GetCouponPattern() string`

GetCouponPattern returns the CouponPattern field if non-nil, zero value otherwise.

### GetCouponPatternOk

`func (o *CodeGeneratorSettings) GetCouponPatternOk() (*string, bool)`

GetCouponPatternOk returns a tuple with the CouponPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponPattern

`func (o *CodeGeneratorSettings) SetCouponPattern(v string)`

SetCouponPattern sets CouponPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


