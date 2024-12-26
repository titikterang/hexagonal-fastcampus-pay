## Generate JWK from self signed key

```shell
# generate private key
openssl genrsa -out ./files/keypair/jwtRSA256-private.pem 2048
openssl genrsa -out ./files/keypair/jwtRSA256-refresh-private.pem 2048
```

```shell
# Extract the public key from the private key
openssl rsa -in ./files/keypair/jwtRSA256-private.pem -pubout -outform PEM -out ./files/keypair/jwtRSA256-public.pem
openssl rsa -in ./files/keypair/jwtRSA256-refresh-private.pem -pubout -outform PEM -out ./files/keypair/jwtRSA256-refresh-public.pem
```

```shell
# install this cli tool to generate jwk 
# https://github.com/jphastings/jwker
go install github.com/jphastings/jwker/cmd/jwker@latest
```


```shell
# Convert the format of the public key from PEM to JWK
echo "generate public jwks to ./cmd/membership/jwks/public.jwk.json"
cat ./files/keypair/jwtRSA256-public.pem | jwker > ./files/jwks/public.jwk.json
echo "generate public jwks to ./cmd/membership/jwks/private.jwk.json"
cat ./files/keypair/jwtRSA256-private.pem | jwker > ./files/jwks/private.jwk.json
```


```shell
# Generate KID random string, use this value as 'kid' field value on public.jwk.json
openssl rand -hex 20
```
