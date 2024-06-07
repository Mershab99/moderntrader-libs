# Strategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyId** | Pointer to **string** | strategyId (UUID) from DB | [optional] 
**StrategyName** | Pointer to **string** | Name of Strategy | [optional] 
**Tags** | Pointer to **[]string** | Tags on this Strategy | [optional] 

## Methods

### NewStrategy

`func NewStrategy() *Strategy`

NewStrategy instantiates a new Strategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyWithDefaults

`func NewStrategyWithDefaults() *Strategy`

NewStrategyWithDefaults instantiates a new Strategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategyId

`func (o *Strategy) GetStrategyId() string`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *Strategy) GetStrategyIdOk() (*string, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *Strategy) SetStrategyId(v string)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *Strategy) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyName

`func (o *Strategy) GetStrategyName() string`

GetStrategyName returns the StrategyName field if non-nil, zero value otherwise.

### GetStrategyNameOk

`func (o *Strategy) GetStrategyNameOk() (*string, bool)`

GetStrategyNameOk returns a tuple with the StrategyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyName

`func (o *Strategy) SetStrategyName(v string)`

SetStrategyName sets StrategyName field to given value.

### HasStrategyName

`func (o *Strategy) HasStrategyName() bool`

HasStrategyName returns a boolean if a field has been set.

### GetTags

`func (o *Strategy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Strategy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Strategy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Strategy) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


