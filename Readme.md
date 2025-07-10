# Build and push Docker image
docker build -t myregistry/go-app:latest .
docker push myregistry/go-app:latest

# Apply Kubernetes manifests
kubectl apply -f deploy/keycloak.yaml
kubectl apply -f deploy/go-app.yaml

# Access Keycloak (port-forward)
kubectl port-forward svc/keycloak 8080:8080

# Test
1. Get a Keycloak token:
curl -X POST \
  -d "client_id=go-app" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "username=testuser" \
  -d "password=test123" \
  -d "grant_type=password" \
  "http://localhost:8080/realms/myrealm/protocol/openid-connect/token"

2. Use the token in API requests:

curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/products
