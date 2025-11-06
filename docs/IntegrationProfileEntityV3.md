# IntegrationProfileEntityV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 

## Methods

### NewIntegrationProfileEntityV3

`func NewIntegrationProfileEntityV3(profileId string, ) *IntegrationProfileEntityV3`

NewIntegrationProfileEntityV3 instantiates a new IntegrationProfileEntityV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationProfileEntityV3WithDefaults

`func NewIntegrationProfileEntityV3WithDefaults() *IntegrationProfileEntityV3`

NewIntegrationProfileEntityV3WithDefaults instantiates a new IntegrationProfileEntityV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *IntegrationProfileEntityV3) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationProfileEntityV3) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *IntegrationProfileEntityV3) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


