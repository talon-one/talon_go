# FeedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the feed notification. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment this feed notification was created. | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the moment this feed notification was last updated. | 
**ArticleUrl** | Pointer to **string** | URL to the feed notification in the help center. | 
**Type** | Pointer to **string** | The type of the feed notification. | 
**Body** | Pointer to **string** | Body of the feed notification. | 

## Methods

### GetTitle

`func (o *FeedNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedNotification) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *FeedNotification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *FeedNotification) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetCreated

`func (o *FeedNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FeedNotification) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *FeedNotification) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *FeedNotification) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetUpdated

`func (o *FeedNotification) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FeedNotification) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *FeedNotification) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *FeedNotification) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetArticleUrl

`func (o *FeedNotification) GetArticleUrl() string`

GetArticleUrl returns the ArticleUrl field if non-nil, zero value otherwise.

### GetArticleUrlOk

`func (o *FeedNotification) GetArticleUrlOk() (string, bool)`

GetArticleUrlOk returns a tuple with the ArticleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArticleUrl

`func (o *FeedNotification) HasArticleUrl() bool`

HasArticleUrl returns a boolean if a field has been set.

### SetArticleUrl

`func (o *FeedNotification) SetArticleUrl(v string)`

SetArticleUrl gets a reference to the given string and assigns it to the ArticleUrl field.

### GetType

`func (o *FeedNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedNotification) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *FeedNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *FeedNotification) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetBody

`func (o *FeedNotification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FeedNotification) GetBodyOk() (string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *FeedNotification) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *FeedNotification) SetBody(v string)`

SetBody gets a reference to the given string and assigns it to the Body field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


