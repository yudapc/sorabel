# sorabel

This project for Toko Ijah.



### Requirements:
- Go version go1.12.6 darwin/amd64
- Glide
- sqlite

### How to run in development

#### Run:

Please checkout this project at `https://github.com/yudapc/sorabel`. And make sure you have been installed glide on your local machine.

Install dependencies:
`glide install`

Run the app:
`go run app.go`

### How to run in production

#### Run:

Please checkout this project at `https://github.com/yudapc/sorabel`. And make sure you have been installed glide on your local machine.

#### Install dependencies:

     `glide install`

---
#### Compile this project:
* OSX:

     `make sorabel-osx`

* Linux:

     `make sorabel-linux`

---
#### Run binary the sorabel app:

     `./sorabel-osx`
---

### ROUTES

#### List Endpoint Items / product:

```
GET /items
GET /items/:id
POST /items
PUT /items/:id
DELETE /items/:id
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
```

