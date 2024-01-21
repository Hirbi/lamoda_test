curl -X POST \
   http://localhost:8088/rpc \
   -H 'content-type: application/json charset=UTF-8' \
   -d '{
   "method": "ProductService.ReleaseProducts",
   "params": [{"codes": ["code-4"]}],
   "id": "1"
   }'