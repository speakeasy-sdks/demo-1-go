# UpdateActionRequest

A representation of the main object defined in JSON Patch specified in RFC 6902 from the IETF. The main differences are:

* Only `add`, `remove` and `replace` are supported

* `remove` can have have its scope of application applied in its `value`. e.g. `{"scope":"delta"}


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `From`             | **string*          | :heavy_minus_sign: | N/A                |
| `Op`               | **string*          | :heavy_minus_sign: | N/A                |
| `Path`             | **string*          | :heavy_minus_sign: | N/A                |
| `Value`            | *interface{}*      | :heavy_minus_sign: | N/A                |