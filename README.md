# REST API Server in Golang

- This API server provides endpoints to create,read,update & delete sample articles.
  

## To Start API Server
```$ git clone https://github.com/piyush1146115/Go-REST-API-with-CLI.git```

```$ cd Go-REST-API-with-CLI```

```$ go install```

```$ Go-REST-API-with-CLI version``` [print the version of the api server]

```$ Go-REST-API-with-CLI startserver```  [run the api server]

## Command to run unit test for API endpoints
```$ cd api```

```$ go test```

## Commands to run API server in docker container
```shell
$ docker build -t <image_name> .
$ docker run -p 8080:8080 <given_image_name> # to start the server with default config
```


## Data Model

- User Model
```
type Article struct {
	Id      string `json: "Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
```


## Available API Endpoints

|  Method | API Endpoint | Description |
|---|---|---|
|GET| /article/{id} | Return an article having the given id| 
|GET| /articles | Returns all the articles| 
|POST| /article | Add a new article | 
|PUT| /article/{id} | Update an article with the given ID| 
|DELETE| /article/{id} | Delete the article with the respective ID| 

## Available Flags

| Flag | Shorthand | Default value | Example | Description
|---|---|---|---|---|
|port|p|8080| Go-REST-API-with-CLI start --port=8090 | Start API server in the given port otherwise in default port

