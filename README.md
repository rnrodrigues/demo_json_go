## Go Demo API

A tiny demo writted in GO that show how to use JSON responses in an API.


### Usage

Clone this repo:
```
git clone https://github.com/rnrodrigues/demo_json_go.git
```

Run command to compile:
```
CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o server
```

Run commands below to build and run docker image:
```
sudo docker build -t mytinyapi .
sudo docker run --rm -it -p 8080:8080 mytinyapi
```
That's it!
