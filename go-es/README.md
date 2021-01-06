# Go-Es
Implementing Elasticsearch in Go with clean code

## Tools
- Go
- Elasticsearch v7
- Gorilla Mux

## TODO List
1. Create item [done!]
2. Get item by id [done!]
3. Search items [done!]
5. Update item [todo]
6. Delete item [todo]

## Create Index example
PUT `127.0.0.1:9200/items`
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

## Endpoints

- **POST `/items`**

Body:
```
{
    "title": "The dark side of the moon",
    "description": {
        "plain_text": "The Dark Side of the Moon is the eighth studio album by English rock band Pink Floyd, released on March 1st, 1973.",
        "html": "The Dark Side of the Moon is the eighth studio album by English rock band Pink Floyd, released on March 1st, 1973."
    },
    "pictures": null,
    "status": "pending"
}
```

- **GET `/items/{id}`**

- **POST `/items/search`**

Body:

```
{
    "equals": [
        {

            "field": "title",
            "value": "The dark side of the moon"
        },
        {

            "field": "status",
            "value": "pending"
        }
    ],
    "not_equals": [
        {
            "field": "title",
            "value": "test"
        }
    ]
}
```

```
{
    "like": [
        {

            "field": "title",
            "value": "The dark side of the moon"
        },
        {

            "field": "status",
            "value": "pending"
        }
    ],
    "not_like": [
        {
            "field": "title",
            "value": "test"
        }
    ]
}
```