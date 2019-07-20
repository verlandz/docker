# Docker
- Docker is a software that very helpful to deliver software in your machine virtualize
- You can see [here](https://www.docker.com/why-docker) for **WHY**
- It's different from virtual machine, try to read [here](https://www.upguard.com/articles/docker-vs.-vmware-how-do-they-stack-up)
- To do it, just pull the image from the registry [here](https://hub.docker.com/search?q=&type=image) then run it as a container (1 container = 1 running image)
- There's also official docs about it, try to understand [here](https://docs.docker.com/)

### Purpose
The purpose of this project is try to help people to understand the basic of **HOW TO USE DOCKER**.

### Objective
This project will make 3 container: 1 web, 1 api and 1 redis. The web will connect to api to get data, and the api will get data from redis as database. For redis image, we can download it from official `docker registry`, but for web and api image we will make it by ourself. The language to make web and api is `Golang`.

### Author Env
- OS : ubuntu 16.04 LTS
- Docker : Docker version 18.09.6, build 481bc77 (Docker Engine - Community)
- Docker-Compose : docker-compose version 1.24.0, build 0aa59064
- Golang : go version go1.12.7 linux/amd64
- Redis : redis-cli 4.0.9

### Getting Started
- [install docker](https://phoenixnap.com/kb/how-to-install-docker-on-ubuntu-18-04)
- [install docker-compose](https://docs.docker.com/compose/install/)
- [install golang](https://tecadmin.net/install-go-on-ubuntu/) (except ver 1.11.11)
- [install redis-cli](https://stackoverflow.com/questions/21795340/linux-install-redis-cli-only)
- [understand basic golang](https://golang.org/doc/)
- [understand basic redis](https://redis.io/)

Usually docker need to be super user, so try to `sudo su -` first.

`Perhaps you already aware`, why we need to install golang to our device, why not pull golang images to our docker ? This is for comparison version. So in docker we will using v1.11.11 but in device using different version from it. Also having install go in your device make a guarantee that you understand how to manage `GOPATH` and `GOROOT`.

# Walkthrough

[create an anchor](#anchors-in-markdown)
### 1. Pull golang and redis images into docker
let's pull the images from registry:
- `golang` with tag `1.11.11`, [source](https://hub.docker.com/_/golang?tab=tags), cmd: `docker pull golang:1.11.11` 
- `redis` with tag `latest`, [source](https://hub.docker.com/_/redis?tab=tags), cmd: `docker pull redis`

check the images:
```
root@201352:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
golang              1.11.11             5b555dd4804a        5 weeks ago         757MB
redis               latest              3c41ce05add9        5 weeks ago         95MB
```
you can also change the tag according to your need later on.


### 2. Run the Web in localhost
- This web is running in port `:8081`
- There's 2 service: `/` for health check and `/data` to display the data that taken from api.
- There's 2 env variable: `API_HOST` for host of API and `API_PORT` for port of API

go to **$GOPATH/src/github.com/verlandz/docker/web**<br>
cmd: `API_HOST=localhost API_PORT=8082 go run main.go`<br>
see the result in `http://localhost:8081/` and `http://localhost:8081/data`<br>

the logs will show this:
```
API_HOST : localhost
API_PORT : 8082
Running in go1.12.7
Listen and serve :8081
```
As you can see,`/data` is fail to connect due the API is not up. Let's run the API in the next step.


### 3. Run the API in localhost
- This api is running in port `:8082`
- There's 2 service: `/` for health check and `/data` to return the data that taken from redis
- There's 2 env variable: `REDIS_HOST` for host of redis and `REDIS_PORT` for port of redis

go to **$GOPATH/src/github.com/verlandz/docker/api**
cmd: `REDIS_HOST=localhost REDIS_PORT=6379 go run main.go`<br>
see the result in `http://localhost:8082/` and `http://localhost:8082/data` <br>

the logs will show this:
```
REDIS_HOST : localhost
REDIS_PORT : 6379
Running in go1.12.7
Listen and serve :8082
```
As you can see, `http://localhost:8081/data` already okay, but there's no data. Let's fill the data in the next step.


### 4. Create and run Redis container
Let's run redis image into container:
`root@201352:~# docker run -t -d --name my_redis -p 6379:6379 redis`

check redis container:
```
root@201352:~# docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                    PORTS                    NAMES
1c7810ab5d96        redis               "docker-entrypoint.s…"   6 days ago          Up 1 second               0.0.0.0:6379->6379/tcp   my_redis
```

let's fill the data, with following sample:

| ID | Name |
| -- | ---- |
| 1  | John |
| 2  | Mike |
| 3  | Sam  |

we're using `Hashes` key with key-name `user`
```
root@201352:~# redis-cli
127.0.0.1:6379> hset user 1 John
(integer) 1
127.0.0.1:6379> hset user 2 Mike
(integer) 1
127.0.0.1:6379> hset user 3 Sam
(integer) 1
127.0.0.1:6379> hgetall user
1) "1"
2) "John"
3) "2"
4) "Mike"
5) "3"
6) "Sam"
```
Now back to `http://localhost:8081/data`, those following data should be show like this
![redis](https://i.ibb.co/g4QLdFg/redis1.png)


### 5. Upload Web image to docker
- to upload it you need to create docker file under web folder.<br>
- go to **$GOPATH/src/github.com/verlandz/docker/web** <br>
- cmd: `sudo docker build --tag web:latest .`<br>
- the logs will show this:
```
Sending build context to Docker daemon  5.632kB
Step 1/8 : FROM golang:1.11.11
 ---> 5b555dd4804a
Step 2/8 : EXPOSE 8081
 ---> Using cache
 ---> 5980ad95ab7e
Step 3/8 : ENV API_HOST=localhost
 ---> Using cache
 ---> fe61e5b2165f
Step 4/8 : ENV API_PORT=8082
 ---> Using cache
 ---> 76e625f375d5
Step 5/8 : WORKDIR $GOPATH/src/github.com/verlandz/docker/web
 ---> Using cache
 ---> 6e3c8036a7d3
Step 6/8 : RUN mkdir -p $PWD
 ---> Using cache
 ---> ed92eb5ac734
Step 7/8 : COPY . $PWD
 ---> Using cache
 ---> 1ca12575cf30
Step 8/8 : CMD ["sh", "-c", "go run $PWD/main.go"]
 ---> Using cache
 ---> cf8f5bee6baf
Successfully built cf8f5bee6baf
Successfully tagged web:latest
```

### 6. Upload API image to docker
- to upload it you need to create docker file under api folder.<br>
- go to **$GOPATH/src/github.com/verlandz/docker/api** <br>
- cmd: `sudo docker build --tag api:latest .`<br>
- the logs will show this:
```
Sending build context to Docker daemon  9.728kB
Step 1/10 : FROM golang:1.11.11
 ---> 5b555dd4804a
Step 2/10 : EXPOSE 8082
 ---> Using cache
 ---> f4ee43644d13
Step 3/10 : ENV REDIS_HOST=localhost
 ---> Using cache
 ---> 75c977c9cc24
Step 4/10 : ENV REDIS_PORT=6379
 ---> Using cache
 ---> baa49dcfcf85
Step 5/10 : WORKDIR $GOPATH/src/github.com/verlandz/docker/api
 ---> Using cache
 ---> 576f2925b252
Step 6/10 : RUN mkdir -p $PWD
 ---> Using cache
 ---> 9419432f9582
Step 7/10 : COPY . $PWD
 ---> Using cache
 ---> 2bc13f161fe7
Step 8/10 : RUN go get -v -u github.com/golang/dep/cmd/dep
 ---> Using cache
 ---> e0ee766b65f8
Step 9/10 : RUN dep ensure -v -vendor-only
 ---> Using cache
 ---> 2fb5855abc11
Step 10/10 : CMD ["sh","-c","go run $PWD/main.go"]
 ---> Using cache
 ---> 0c6ed6cb0946
Successfully built 0c6ed6cb0946
Successfully tagged api:latest
```

### 7. Create and run Web + API Container
First, let's check the images that have been uploaded in step 5 and 6.
```
root@201352:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
api                 latest              0c6ed6cb0946        7 days ago          808MB
web                 latest              cf8f5bee6baf        7 days ago          757MB
```

Then try to run those images into container :
<br>***Web***<br>
`root@201352:~# docker run -t -d -e API_HOST=my_api -e API_PORT=8082 --name my_web -p 8081:8081 web
`
<br>***API***<br>
`root@201352:~# docker run -t -d -e REDIS_HOST=my_redis -e REDIS_PORT=6379 --name my_api -p 8082:8082 api
`

After that try to visit `http://localhost:8081/` and `http://localhost:8082/`, you can find the web and api is up but.. when you visit `http://localhost:8081/data` it show that web fail to connect to api and `http://localhost:8082/data` show fail to get data from redis.

The reason behind those fail is because they try to connect to their own container's localhost, therefore we need to connect those container into one network

### 8. Create and connect the network
- create the network with name my_net : `docker network create my_net` <br>
- connect to my_web : `docker network connect my_net my_web`
- connect to my_api : `docker network connect my_net my_api`
- connect to my_redis : `docker network connect my_net my_redis`
- to see whether the container already connect to desired network try to inpesct with cmd : `docker inspect conatiner_name` and see the value in field `NetworkSettings > Networks`

Now, try to visit again `http://localhost:8081/data` and `http://localhost:8082/data`, the data will show.

### 9. Storage 
Docker is capable to store data that state in container into storage, so the other container can use that data too. This section will focus on 2 types storage: `Volume` and `Bind Mount`. `Volume` save data into your docker area and `Bind Mount` save data into your file system. You can read [here](https://docs.docker.com/storage/) for further explanation.

Let's try volume :
- stop redis container `root@201352:~# docker stop my_redis`
- remove redis container `root@201352:~# docker rm my_redis`
- create volume `root@201352:~# docker volume create my_redis_data`
- run container that connect to that volume `root@201352:~# docker run -t -d --name my_redis -p 6379:6379 -v my_redis_data:/data_redis`
- try to fill data in redis-cli ex : `hset user 1 "I'm from volume"`
- connect network `root@201352:~# docker network connect my_net my_redis`
- go to `http://localhost:8081/data` or `http://localhost:8082/data`
- stop redis container `root@201352:~# docker stop my_redis`
- remove redis container `root@201352:~# docker rm my_redis`
- run container that connect to that volume `root@201352:~# docker run -t -d --name my_redis -p 6379:6379 -v my_redis_data:/data_redis`
- connect network `root@201352:~# docker network connect my_net my_redis`
- in previous, create new container will start fresh with empty data, but not in this case because we use volume
- go to `http://localhost:8081/data` or `http://localhost:8082/data`


### 10. Docker Compose
Docker compose is a tool that help you to create, run, remove, connect the container, network, etc. all together. So it's kinda a script that help you to do your thing automatically, rather than manually do a single command for each action. You can see more [here](https://docs.docker.com/compose/)

To try docker-compose let's delete existing container and network<br>
```
root@201352:~# docker stop my_api my_web my_redis
root@201352:~# docker rm my_api my_web my_redis
root@201352:~# docker network rm my_net
```

Then try to up docker-compose
- go to **$GOPATH/src/github.com/verlandz/docker** <br>
- sudo `docker-compose up -d`
the logs will show like this:
```
Creating network "my_net" with the default driver
Creating my_redis ... done
Creating my_api   ... done
Creating my_web   ... done
```
I believe you redis's data will gone missing, since the container is new. So you can set it again like step 4.

So is there's a way to store a data even the container is new ? Yes, let's move to next step for that.
