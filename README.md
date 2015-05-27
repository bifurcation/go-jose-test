# Go JOSE Testing

This is a simple test of [go-jose](https://github.com/square/go-jose).  The go-jose snippet in `jose-validate.go` is identical to the JWS validation step in [boulder](https://github.com/letsencrypt/boulder/), so if your JWS generation code passes validation with `jose-validate`, then it should work fine with boulder.

## Quickstart

```
> make && make test
go build jose-validate.go
javac -classpath jose4j-0.4.2.jar JwsTest.java
java -classpath jose4j-0.4.2.jar:commons-logging-1.2.jar:. JwsTest \
		2>/dev/null \
		| ./jose-validate
parsing: OK
verification: OK
```


## Contents

* jose-validate.go - JWS verifier based on go-jose
* jose.html - JS/WebCrypto JWS verifier (requires Firefox for ES6)
* JwtTest.java - JWS generator based on jose4j

The outputs of the Go and JS versions should be consistent.
