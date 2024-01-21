# RPC api 

## Methods
### Reserve products
__scripts_examples/reserve_codes.sh__
```
curl -X POST \
   http://localhost:8088/rpc \
   -H 'content-type: application/json charset=UTF-8' \
   -d '{
   "method": "ProductService.ReserveProducts",
   "params": [{"codes": ["code-4"]}],
   "id": "1"
}'
```
Result

```
{"result":"reserved","error":null,"id":"1"}
```

### Release products
__scripts_examples/release_codes.sh__
```
curl -X POST \
   http://localhost:8088/rpc \
   -H 'content-type: application/json charset=UTF-8' \
   -d '{
   "method": "ProductService.ReleaseProducts",
   "params": [{"codes": ["code-4"]}],
   "id": "1"
}'
```
Result
```
{"result":"released","error":null,"id":"1"}
```

### Available products on storage
__scripts_examples/available_products.sh__
```
curl -X POST \
   http://localhost:8088/rpc \
   -H 'content-type: application/json charset=UTF-8' \
   -d '{
   "method": "ProductService.AvailableProducts",
   "params": [{"StorageId": "1"}],
   "id": "1"
}'
```
Result
```
{
  "result": [
    {
      "Name": "keyboard",
      "Size": "10x10",
      "Code": "code-1",
      "Amount": 2
    },
    {
      "Name": "headphones",
      "Size": "10x10",
      "Code": "code-2",
      "Amount": 4
    },
    {
      "Name": "mouse",
      "Size": "10x10",
      "Code": "code-3",
      "Amount": 3
    },
    {
      "Name": "usb-hub",
      "Size": "10x10",
      "Code": "code-4",
      "Amount": 8
    }
  ],
  "error": null,
  "id": "1"
}
```
## Launch

``` 
make up 
```
## Migrate
``` 
make migrate 
```