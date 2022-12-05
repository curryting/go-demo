### 1. N/A

1. route definition

- Url: /search/do
- Method: GET
- Request: `SearchReq`
- Response: `SearchReply`

2. request definition



```golang
type SearchReq struct {
	Name string `form:"name"`
}
```


3. response definition



```golang
type SearchReply struct {
	Name string `json:"name"`
	Count int `json:"count"`
}
```

### 2. N/A

1. route definition

- Url: /search/ping
- Method: GET
- Request: `-`
- Response: `-`

2. request definition



3. response definition


