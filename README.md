# chitchat is a web server simple demo for go

## Running
Prepare build environment:

`Need to install:`  
* **docker**  
* **go**


Clone repo:
```
$ git clone https://github.com/backlash-go/chitChat.git
$ cd chitChat
$ docker build   -t  chitchat:v1 .
$ docker run -p 8281:8281 --name chitchat --detach chitchat:v1
```

Build and run:
```
$ docker build   -t  chitchat:v1 .
$ docker run -p 8281:8281 --name chitchat --detach chitchat:v1
```

Access:
```
$ curl http://host:port
```







