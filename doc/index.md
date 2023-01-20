# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [calculator/v1/calculator.proto](#calculator_v1_calculator-proto)
    - [Error](#calculator-v1-Error)
    - [Operation](#calculator-v1-Operation)
  
- [services/calculate/v1/calculate.proto](#services_calculate_v1_calculate-proto)
    - [AdditionRequest](#services-calculate-v1-AdditionRequest)
    - [AdditionResponse](#services-calculate-v1-AdditionResponse)
    - [AdditionsRequest](#services-calculate-v1-AdditionsRequest)
    - [AdditionsResponse](#services-calculate-v1-AdditionsResponse)
    - [AdditionsResponse.ErrorsEntry](#services-calculate-v1-AdditionsResponse-ErrorsEntry)
    - [DivisionRequest](#services-calculate-v1-DivisionRequest)
    - [DivisionResponse](#services-calculate-v1-DivisionResponse)
  
    - [CalculateService](#services-calculate-v1-CalculateService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="calculator_v1_calculator-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## calculator/v1/calculator.proto


 


<a name="calculator-v1-Error"></a>

### Error


| Name | Number | Description |
| ---- | ------ | ----------- |
| ERROR_UNSPECIFIED | 0 |  |
| ERROR_WRONG_OPERATION | 1 |  |
| ERROR_DIVISION_BY_ZERO | 2 |  |



<a name="calculator-v1-Operation"></a>

### Operation


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATION_UNSPECIFIED | 0 |  |
| OPERATION_ADDITION | 1 |  |
| OPERATION_SUBTRACTION | 2 |  |
| OPERATION_MULTIPLICATION | 3 |  |
| OPERATION_DIVISION | 4 |  |


 

 

 



<a name="services_calculate_v1_calculate-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## services/calculate/v1/calculate.proto



<a name="services-calculate-v1-AdditionRequest"></a>

### AdditionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [calculator.v1.Operation](#calculator-v1-Operation) |  |  |
| operator_first | [double](#double) |  |  |
| operator_second | [double](#double) |  |  |






<a name="services-calculate-v1-AdditionResponse"></a>

### AdditionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [string](#string) |  |  |






<a name="services-calculate-v1-AdditionsRequest"></a>

### AdditionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [calculator.v1.Operation](#calculator-v1-Operation) |  |  |
| operator_first | [double](#double) |  |  |
| operator_second | [double](#double) |  |  |






<a name="services-calculate-v1-AdditionsResponse"></a>

### AdditionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [double](#double) |  |  |
| errors | [AdditionsResponse.ErrorsEntry](#services-calculate-v1-AdditionsResponse-ErrorsEntry) | repeated |  |






<a name="services-calculate-v1-AdditionsResponse-ErrorsEntry"></a>

### AdditionsResponse.ErrorsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [calculator.v1.Error](#calculator-v1-Error) |  |  |






<a name="services-calculate-v1-DivisionRequest"></a>

### DivisionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [calculator.v1.Operation](#calculator-v1-Operation) |  |  |
| operator_first | [double](#double) |  |  |
| operator_second | [double](#double) |  |  |






<a name="services-calculate-v1-DivisionResponse"></a>

### DivisionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [double](#double) |  |  |





 

 

 


<a name="services-calculate-v1-CalculateService"></a>

### CalculateService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Addition | [AdditionRequest](#services-calculate-v1-AdditionRequest) | [AdditionResponse](#services-calculate-v1-AdditionResponse) |  |
| Division | [DivisionRequest](#services-calculate-v1-DivisionRequest) | [DivisionResponse](#services-calculate-v1-DivisionResponse) |  |
| Additions | [AdditionsRequest](#services-calculate-v1-AdditionsRequest) stream | [AdditionsResponse](#services-calculate-v1-AdditionsResponse) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

