
curl -sX POST \
  -d "client_id=go-app" \
  -d "username=testwriteuser" \
  -d "password=test123" \
  -d "grant_type=password" \
  "http://localhost:8080/realms/myrealm/protocol/openid-connect/token" | jq '.access_token'

