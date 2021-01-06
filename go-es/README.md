# Go-Es
Implementing Elasticsearch in Go with clean code

## Tools
- Go
- Elasticsearch v7
- Gorilla Mux

## Elastic Index
PUT 127.0.0.1:9200/items
```
{
    "settings": {
        "index": {
            "number_of_shards": 4,
            "number_of_replicas": 2
        }
    }
}
```

