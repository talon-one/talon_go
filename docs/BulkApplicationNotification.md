# BulkApplicationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**Data** | Pointer to [**[]ApplicationNotification**](ApplicationNotification.md) |  | 

## Methods

### NewBulkApplicationNotification

`func NewBulkApplicationNotification(totalResultSize int64, data []ApplicationNotification, ) *BulkApplicationNotification`

NewBulkApplicationNotification instantiates a new BulkApplicationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkApplicationNotificationWithDefaults

`func NewBulkApplicationNotificationWithDefaults() *BulkApplicationNotification`

NewBulkApplicationNotificationWithDefaults instantiates a new BulkApplicationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *BulkApplicationNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *BulkApplicationNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *BulkApplicationNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *BulkApplicationNotification) GetData() []ApplicationNotification`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkApplicationNotification) GetDataOk() (*[]ApplicationNotification, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkApplicationNotification) SetData(v []ApplicationNotification)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


