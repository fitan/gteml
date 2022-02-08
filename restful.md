## getlist
### getList
```http request
GET http://my.api.url/posts?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"}
```
### getOne
```http request
GET http://my.api.url/posts/123
```
### getMany
```http request
GET http://my.api.url/posts?filter={"ids":[123,456,789]}
```

### getManyReference	
```http request
GET http://my.api.url/posts?filter={"author_id":345}
```
filter=[{"field": "author_id", "symbol": "gt", "val": ""},{"field": "time", "symbol": "between", "left": "", "right": ""}]

### create	
```http request
POST http://my.api.url/posts
```

### update	
```http request
PUT http://my.api.url/posts/123
```

### updateMany	
```http request
PUT http://my.api.url/posts/123
```

### delete
```http request
DELETE http://my.api.url/posts/123
```

### deleteMany
```http request
DELETE http://my.api.url/posts/123
```