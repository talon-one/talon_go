# AsyncCouponCreationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | The batch ID that all coupons created by the request will have. | 

## Methods

### NewAsyncCouponCreationResponse

`func NewAsyncCouponCreationResponse(batchId string, ) *AsyncCouponCreationResponse`

NewAsyncCouponCreationResponse instantiates a new AsyncCouponCreationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncCouponCreationResponseWithDefaults

`func NewAsyncCouponCreationResponseWithDefaults() *AsyncCouponCreationResponse`

NewAsyncCouponCreationResponseWithDefaults instantiates a new AsyncCouponCreationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *AsyncCouponCreationResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *AsyncCouponCreationResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *AsyncCouponCreationResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


