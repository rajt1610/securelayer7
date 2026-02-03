# CRUD REST API in Go

Simple in-memory CRUD REST API.

Run:

```bash
cd /home/sl7-go/securelayer7/crud-golang
go run .
```

Endpoints:

- `GET /items` — list items
- `POST /items` — create item (JSON: `name`, `description`)
- `GET /items/{id}` — get item
- `PUT /items/{id}` — update item (JSON: `name`, `description`)
- `DELETE /items/{id}` — delete item

Example create:

```bash
curl -s -X POST -H "Content-Type: application/json" -d '{"name":"foo","description":"desc"}' http://localhost:8080/items
```

Database:

- The server uses SQLite and will create a `crud.db` file in the working directory.
- To reset data, stop the server and remove `crud.db`.

