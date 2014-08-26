# Pnthr Examples

We're aiming to build both a client and server in Golang for pnthr. Currently the Server is mid-refactoring to be a more feature rich package. Follow the following instructions to get up and running:

## Create the Mongo Database

Pnthr Server takes two parameters currently for MongoDB: a host uri and database name. For these examples, we will need to create a database called `pnthr` with a collection called `instances` with this single record:

```
{
  "_id": ObjectId("53a1c59f6239370002000000"),
  "name": "Test for Ruby Gem Spec",
  "description": "password is 'password'",
  "password": "5f4dcc3b5aa765d61d8327deb882cf99",
  "secret": "aa88906ffcf6c59aaf5908d3900f21a6",
  "user_id": ObjectId("535df0646636640002000000"),
  "updated_at": new Date(1403110815507),
  "created_at": new Date(1403110815507)
}
```

## Get the Package

In terminal, we'll install the package from the repo like usual:
`go get github.com/pnthr/pnthr`



## Run the server

To run the server locally:
`DB_NAME=pnthr go run pnthr-server.go`

You should see either any errors or a success along the lines of `Listening for connections...`. You can now run a suite of ruby tests (until the go client is ready) by downloading the [pnthr gem](https://github.com/pnthr/pnthr-gem) (instructions are on that repo's readme).

Happy encrypting!
