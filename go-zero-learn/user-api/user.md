### 1. "获取用户信息"

1. route definition

- Url: /user/ingo
- Method: GET
- Request: `UserReq`
- Response: `UserRes`

2. request definition



```golang
type UserReq struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type UserRes struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
```

