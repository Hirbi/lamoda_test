curl -X POST \
   http://localhost:8088/rpc \
   -H 'content-type: application/json charset=UTF-8' \
   -d '{
   "method": "ProductService.AvailableProducts",
   "params": [{"StorageId": "1"}],
   "id": "1"
   }'