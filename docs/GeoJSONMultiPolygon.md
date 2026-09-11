# GeoJSONMultiPolygon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The geometry type discriminator. | 
**Coordinates** | Pointer to [**[][][][]float32**](array.md) | The shapes in this group. Each one follows the same boundary structure as a polygon. | 

## Methods

### NewGeoJSONMultiPolygon

`func NewGeoJSONMultiPolygon(type_ string, coordinates [][][][]float32, ) *GeoJSONMultiPolygon`

NewGeoJSONMultiPolygon instantiates a new GeoJSONMultiPolygon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONMultiPolygonWithDefaults

`func NewGeoJSONMultiPolygonWithDefaults() *GeoJSONMultiPolygon`

NewGeoJSONMultiPolygonWithDefaults instantiates a new GeoJSONMultiPolygon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONMultiPolygon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONMultiPolygon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONMultiPolygon) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *GeoJSONMultiPolygon) GetCoordinates() [][][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoJSONMultiPolygon) GetCoordinatesOk() (*[][][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoJSONMultiPolygon) SetCoordinates(v [][][][]float32)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


