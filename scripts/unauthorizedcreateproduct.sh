
TOKEN=$(./gettoken.sh)
curl -X POST -H "Authorization: Bearer $TOKEN" http://localhost:8090/products -d "{\"name\": \"laptop\",\"price\": 100}"

