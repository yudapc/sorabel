# sorabel

This project for solution Toko Ijah.

Board this project https://trello.com/b/DLofC4BT/sorabel

### Requirements:

- Go version go1.12.6 darwin/amd64
- Glide
- sqlite
- This app run on PORT 8000


### Checkout source code

`git clone https://github.com/yudapc/sorabel.git`

### Install dependencies:

`glide install`

### Run in development:

`go run app.go`

### Added permission execute:

OSX: `chmod +x start-osx.sh`
Linux: `chmod +x start-linux.sh`

### Run with script:

OSX: `./start-osx.sh`
Linux: `./start-linux.sh`

### Compile this project:

OSX: `make sorabel-osx`
Linux: `make sorabel-linux`

#### Run binary the sorabel app:

OSX: `./sorabel-osx`
Linux: `./sorabel-linux`

### POSTMAN COLLECTION

You can import file collections to postman for testing each endpoint with sample payload:

https://github.com/yudapc/sorabel/blob/master/Sorabel.postman_collection.json

---

### ROUTES

#### List Endpoint Items / product:

```
GET /items
GET /items/:id
POST /items
PUT /items/:id
DELETE /items/:id
POST /items/import
GET /items/export
```

---

#### List Endpoint Purchase:

```
GET /purchases
GET /purchases/:id
GET /purchases/:id/items
POST /purchases
PUT /purchases/:id
DELETE /purchases/:id
POST /purchases/import
GET /purchases/export
```

---

#### List Endpoint Sales:

```
GET /sales
GET /sales/:id
GET /sales/:id/items
POST /sales
PUT /sales/:id
DELETE /sales/:id
POST /sales/import
GET /sales/export
```
