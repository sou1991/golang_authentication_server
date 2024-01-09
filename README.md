# golang_authentication_server

alpineのコンテナ接続
`dc exec auth /bin/ash`

### client check

```
curl http://localhost:8080/outh2/v1/auth?client_id=abcde&response_type=code
```

### auth
```
curl http://localhost:8080/outh2/v1/auth/login \
    -i \
    -H "Content-Type: application/json" \
    -X "POST" \
    -d '{"email": "auth@example.com","title": "password"}'
```

### 記事書いた
https://zenn.dev/sou1991/articles/876cb2ca79a539
