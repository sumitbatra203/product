
TOKEN=$(./gettoken.sh)
curl -H "Authorization: Bearer $TOKEN" http://localhost:8090/products

