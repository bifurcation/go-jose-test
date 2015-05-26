# Go JOSE

This is a simple test of [go-jose](https://github.com/square/go-jose).

```
go get github.com/square/go-jose
go build jose-validate.go
./jose-validate
# Paste in JWS
# validation: OK or validation: FAIL
```

This input should return OK:

```
{
  "header": {
    "alg": "RS256",
    "jwk": {
      "e": "AQAB",
      "kty": "RSA",
      "n": "tSwgy3ORGvc7YJI9B2qqkelZRUC6F1S5NwXFvM4w5-M0TsxbFsH5UH6adigV0jzsDJ5imAechcSoOhAh9POceCbPN1sTNwLpNbOLiQQ7RD5mY_pSUHWXNmS9R4NZ3t2fQAzPeW7jOfF0LKuJRGkekx6tXP1uSnNibgpJULNc4208dgBaCHo3mvaE2HV2GmVl1yxwWX5QZZkGQGjNDZYnjFfa2DKVvFs0QbAk21ROm594kAxlRlMMrvqlf24Eq4ERO0ptzpZgm_3j_e4hGRD39gJS7kAzK-j2cacFQ5Qi2Y6wZI2p-FCq_wiYsfEAIkATPBiLKl_6d_Jfcvs_impcXQ"
    }
  },
  "payload": "Zm9vCg",
  "signature": "hRt2eYqBd_MyMRNIh8PEIACoFtmBi7BHTLBaAhpSU6zyDAFdEBaX7us4VB9Vo1afOL03Q8iuoRA0AT4akdV_mQTAQ_jhTcVOAeXPr0tB8b8Q11UPQ0tXJYmU4spAW2SapJIvO50ntUaqU05kZd0qw8-noH1Lja-aNnU-tQII4iYVvlTiRJ5g8_CADsvJqOk6FcHuo2mG643TRnhkAxUtazvHyIHeXMxydMMSrpwUwzMtln4ZJYBNx4QGEq6OhpAD_VSp-w8Lq5HOwGQoNs0bPxH1SGrArt67LFQBfjlVr94E1sn26p4vigXm83nJdNhWAMHHE9iV67xN-r29LT-FjA"
}
```
