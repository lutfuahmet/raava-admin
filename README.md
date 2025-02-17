## Raava Admin
Forked from [qor/admin](https://github.com/qor/admin)

Instantly create a beautiful, cross platform, configurable Admin Interface and API for managing your data in minutes.

[![GoDoc](https://godoc.org/github.com/qor/admin?status.svg)](https://godoc.org/github.com/qor/admin)
[![Build Status](https://travis-ci.com/qor/admin.svg?branch=master)](https://travis-ci.com/qor/admin)

**For security issues, please send us an email to security@getqor.com and give us time to respond BEFORE posting as an issue or reporting on public forums.**

## Features

- Generate Admin Interface for managing data
- RESTFul JSON API
- Association handling
- Search and filtering
- Actions/Batch Actions
- Authentication and Authorization
- Extendability

## Quick Start

```go
package main

import (
  "fmt"
  "net/http"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  admin "github.com/TykTechnologies/raava-admin"
)

// Create a GORM-backend model
type User struct {
  gorm.Model
  Name string
}

// Create another GORM-backend model
type Product struct {
  gorm.Model
  Name        string
  Description string
}

func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Product{})

  // Initialize
  Admin := admin.New(&admin.AdminConfig{DB: DB})

  // Allow to use Admin to manage User, Product
  Admin.AddResource(&User{})
  Admin.AddResource(&Product{})

  // initialize an HTTP request multiplexer
  mux := http.NewServeMux()

  // Mount admin interface to mux
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}
```

`go run main.go` and visit `localhost:9000/admin` to see the result!

## Live DEMO

* Live Demo [http://demo.getqor.com/admin](http://demo.getqor.com/admin)
* Source Code of Live Demo [https://github.com/qor/qor-example](https://github.com/qor/qor-example)

## Documentation

<https://doc.getqor.com/admin>

To print all registered routes
```go
// adm is a QOR admin instance
adm.GetRouter().PrintRoutes()
```

## ViewPath Note

QOR was developed before go mod was introduced. So it still support go path while finding its template files. The priority is

1. check vendor, if not found
2. check $GOPATH/src/github.com/TykTechnologies/raava-admin/views, if still not found
3. load view path from $GOPATH/pkg/mod/github.com/TykTechnologies/raava-admin@v0.x/views. the version would be detected automatically by you go.sum file

So if you want to use the template under the pkg/mod, make sure $GOPATH/src/github.com/TykTechnologies/raava-admin is absent.

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).

# Frontend Development
Requires [Node.js](https://nodejs.org/) and [Gulp](http://gulpjs.com/) for building frontend files

```
npm install && npm install -g gulp
```

* Watch SCSS/JavaScript changes (hot-reload): `gulp`
* Build Release files: `gulp release`

# Update Raava to use the latest version
We need the latest commit hash from `raava-admin` repo. You can get it from the GitHub repo, or:

### From a local copy
We need to have the latest version of the `master` branch
```bash
# make sure we are in the correct branch
$ git checkout master
# get the lastest update
$ git pull
# display the hash from the las commit 
$ git log -1 --format='%H'
```

Now that we got the commit hash, we need to move to the local `raava` folder and do this
```bash
$ go get github.com/TykTechnologies/raava-admin@<commit-hash-here>
```