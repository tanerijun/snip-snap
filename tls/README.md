`cd` into this folder and run this command to generate certificates

```sh
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Update the path accordingly if your Go installation folder is different
