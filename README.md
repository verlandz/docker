# DOCKER
- Docker is a software that very helpful to deliver software in your machine virtualize.
- You can see [here](https://www.docker.com/why-docker) for **WHY**
- It's different from virtual machine, try to read [here](https://www.upguard.com/articles/docker-vs.-vmware-how-do-they-stack-up)
- To do it, just pull the image from the registry [here](https://hub.docker.com/search?q=&type=image) then run it as a container (1 container = 1 running image)
- There's also official docs about it, try to understand [here](https://docs.docker.com/)

### Purpose
The purpose of this project is try to help people to understand the basic of **HOW TO USE DOCKER**.

### Objective
This project will make 3 container: 1 web, 1 api and 1 redis. The web will connect to api to get data, and the api will get data from redis as database. For redis image, we can download it from official `docker registry`, but for web and api image we will make it by ourself. The language to make web and api is `Golang`.

Sorry for people who don't understand `Golang`, but it might be a good chance to learn docker and golang at the same time.
And I'm using ubuntu during the project, so it might be a little different to another OS.

### Author Env
- OS : ubuntu 16.04 LTS
- Docker : Docker version 18.09.6, build 481bc77 (Docker Engine - Community)
- Golang : go version go1.12.7 linux/amd64
- Redis : redis-cli 4.0.9

for simplify try to install `docker community` edition

### Getting Started
- [install docker](https://phoenixnap.com/kb/how-to-install-docker-on-ubuntu-18-04)
- [install docker-compose](https://docs.docker.com/compose/install/)
- [install golang](https://tecadmin.net/install-go-on-ubuntu/) (except ver 1.11.11)
- [install redis-cli](https://stackoverflow.com/questions/21795340/linux-install-redis-cli-only)
- [understand basic golang](https://golang.org/doc/)
- [understand basic redis](https://redis.io/)

Usually docker need to be super user, so try to `sudo su -` first.<br>
`Perhaps you already aware`, why we need to install golang to our device, why not pull golang images to our docker ? This is for comparison version. So in docker we will using v1.11.11 but in device using different version from it. Also having install go in your device make a guarantee that you understand how to manage `GOPATH` and `GOROOT`.

# GUIDE

### 1. Pull golang and redis
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


### 2. Create Web
- This web is running in port `:8081`
- There's 2 service: `/` for health check and `/data` to display the data that taken from api.
- There's 2 env variable: `API_HOST` for host of API and `API_PORT`, for port of API

cmd: **$GOPATH/src/github.com/verlandz/docker/web$** `API_HOST=localhost API_PORT=8082 go run main.go`<br>
see result in `http://localhost:8081/` and `http://localhost:8081/data`<br>

the logs will show something like this:
```
API_HOST : localhost
API_PORT : 8082
Running in go1.12.7
Listen and serve :8081
```
As you can see,`/data` is fail to connect due the API is not up. Let's up the API in the next step.


### 3. Create API
- This api is running in port `:8082`
- There's 2 service: `/` for health check and `/data` o return the data that taken from redis
- There's 2 env variable: `REDIS_HOST` for host of redis and `REDIS_PORT` for port of redis

cmd: **$GOPATH/src/github.com/verlandz/docker/api$** `REDIS_HOST=localhost REDIS_PORT=6379 go run main.go`<br>
see result in `http://localhost:8082/` and `http://localhost:8082/data` <br>

the logs will show something like this:
```
REDIS_HOST : localhost
REDIS_PORT : 6379
Running in go1.12.7
Listen and serve :8082
```

As you can see, `http://localhost:8081/data` already okay, but there's no data. Let's fill the data is next step.

### 4. Create Redis
Let's run the redis:
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
Now back to `http://localhost:8081/data`, those following data should be show like [this](https://ibb.co/3RV9vYh) 
