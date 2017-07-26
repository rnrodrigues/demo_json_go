## Go Demo API ##

A tiny demo written in GO that show how to use JSON responses in an API.


### Usage ###

Clone this repo:
```
git clone https://github.com/rnrodrigues/demo_json_go.git
```

Run command on terminal to compile:
```
CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o server
```

or run shell script:
```
sh build.sh
```

Run commands below to build and run docker image:
```
sudo docker build -t mytinyapi .
sudo docker run --rm -it -p 8080:8080 mytinyapi
```
That's it!
