### 1. N/A

1. route definition

- Url: /user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginReply`

2. request definition



```golang
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginReply struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 2. N/A

1. route definition

- Url: /user/search
- Method: GET
- Request: `SearchReq`
- Response: `SearchReply`

2. request definition



```golang
type SearchReq struct {
	Name string `json:"name"`
}
```


3. response definition



```golang
type SearchReply struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Number string `json:"number"`
	Password string `json:"password"`
}
```

