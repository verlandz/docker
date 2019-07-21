# Docker

### Introduction
- [What is docker ?](https://opensource.com/resources/what-docker) 
- [What is docker container ?](https://www.sdxcentral.com/containers/definitions/what-is-docker-container/)
- [Why docker ?](https://www.docker.com/why-docker)
- [Official Docs](https://docs.docker.com/)

### Purpose
Make a simple project with `docker` and help to learn the basics.

### Objective
This project will make 3 Containers: 1 web, 1 api and 1 redis. The web will connect to the api to get the data, and the api will get the data from redis as database. For redis image, we can download it from the official [docker registry](https://hub.docker.com/search?q=&type=image), but for web and api images, we will make it by ourself. The language to make web and api is `Golang` aka `Go`.

You can use another language than `Go`, but this project will focus on using `Go`.

### Author Env
- OS : ubuntu 16.04 LTS
- Docker : Docker version 18.09.6, build 481bc77 (Docker Engine - Community)
- Docker-Compose : docker-compose version 1.24.0, build 0aa59064
- Golang : go version go1.12.7 linux/amd64
- Redis : redis-cli 4.0.9

### Prerequisite
- [Install docker](https://phoenixnap.com/kb/how-to-install-docker-on-ubuntu-18-04)
- [Install docker-compose](https://docs.docker.com/compose/install/)
- [Install golang](https://tecadmin.net/install-go-on-ubuntu/) (except go1.11)
- [Install redis-cli](https://stackoverflow.com/questions/21795340/linux-install-redis-cli-only)
- [Basic golang](https://golang.org/doc/)
- [Basic cmd redis](https://redis.io/)

Usually docker need to be super user, so try to `sudo su -` first.

***FAQ***<br>
**Q:** Why we need to install go to our device if we can pull go image to our docker ?<br>
**A:** The reason is for comparison version. So in docker, we're using go1.11 but in local we're using another version. Also having install go in your device make a guarantee that you understand how to manage `GOPATH` and `GOROOT`, we need it to undestand for `dockerfile`.


# Walkthrough
A guideline step by step to do this project.

## 1. Run the Web in localhost
- The Web is running in port `:8081`
- There's 2 services, `/` for health check and `/data` to display the data that taken from api
- There's 2 env variables, `API_HOST` for host of api and `API_PORT` for port of api

cd `$GOPATH/src/github.com/verlandz/docker/web`<br>
cmd `API_HOST=localhost API_PORT=8082 go run main.go`<br>
```
API_HOST : localhost
API_PORT : 8082
Running in go1.12.7
Listen and serve :8081
```
![8081-home](https://i.ibb.co/VHFNhWP/8081-home.png)
![8081-data](https://i.ibb.co/R32hcqz/8081-data.png)

`/data` is fail to connect due the api is not up. Let's up the api in the next step.


## 2. Run the API in localhost
- The API is running in port `:8082`
- There's 2 services, `/` for health check and `/data` to return the data that taken from redis
- There's 2 env variables, `REDIS_HOST` for host of redis and `REDIS_PORT` for port of redis

cd `$GOPATH/src/github.com/verlandz/docker/api`<br>
cmd `REDIS_HOST=localhost REDIS_PORT=6379 go run main.go`<br>
```
REDIS_HOST : localhost
REDIS_PORT : 6379
Running in go1.12.7
Listen and serve :8082
```
![8082-home](https://i.ibb.co/9cgHCcm/8082-home.png)
![8082-data](https://i.ibb.co/xqqHZdr/8082-data.png)

`/data` is fail to connect due the redis is not up. Let's up the redis in the next step.


## 3. Run Redis in localhost
Pull the latest redis image from [here](https://hub.docker.com/_/redis?tab=tags) then run it into container.
```
root@201352:~# docker pull redis
Using default tag: latest
latest: Pulling from library/redis
f5d23c7fed46: Pull complete 
a4a5c04dafc1: Pull complete 
605bafc84bc9: Pull complete 
f07a4e35cd96: Pull complete 
17944e5e3eb7: Pull complete 
6f875a8605e0: Pull complete 
Digest: sha256:8888f6cd2509062a377e903e17777b4a6d59c92769f6807f034fa345da9eebcf
Status: Downloaded newer image for redis:latest

root@201352:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
redis               latest              598a6f110d01        9 days ago          118MB

root@201352:~# docker run -d --name my_redis -p 6379:6379 redis
6a191a7285efe06b5d2f0269e434dd9d0c1ce4b542b6deee61731ca2b9db9b14

root@201352:~# docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                    NAMES
6a191a7285ef        redis               "docker-entrypoint.s…"   5 seconds ago        Up 2 seconds        0.0.0.0:6379->6379/tcp   my_redis
```
Then fill the data in `redis-cli`
```
intern@201352:~$ redis-cli -h localhost -p 6379
localhost:6379> hset user 1 Jonathan
(integer) 1
localhost:6379> hset user 2 Mike
(integer) 1
localhost:6379> hset user 3 Ricky
(integer) 1
localhost:6379> hgetall user
1) "1"
2) "Jonathan"
3) "2"
4) "Mike"
5) "3"
6) "Ricky"
```
![8081-data-show](https://i.ibb.co/hcwMB4d/8081-data-show.png)
![8082-data-show](https://i.ibb.co/4VsTqcs/8082-data-show.png)

***FAQ***<br>
**Q:** What does it mean by `6379:6379` ?<br>
**A:** It indicate `outside:inside` port. Outside to connect to your localhost, and inside to connect to your container.<br>

**Q:** How do I know redis port in the container is 6379 ?<br>
**A:** Try to `docker inspect redis` and find `ExposedPorts` in `Config` field.<br>

**Q:** What's the diff between `docker run` and `docker create` ? <br>
**A:** `docker run` equal to `docker create` then `docker start`. Create is just create the container but not run it. <br>

**Q:** How to see my containers ? <br>
**A:** `docker ps` / `docker container ls` or `docker ps -a` / `docker container ls -a` for complete. 

**Q:** How to see my images ? <br>
**A:** `docker images` or `docker images -a` for complete.

**Q:** What `-d` for?<br>
**A:** Stand for `detach`, to run in background.


## 4. Upload Web image to docker and run it
- To upload it you need to create a dockerfile under web folder
- Don't forget to stop web that running in your localhost

cd `$GOPATH/src/github.com/verlandz/docker/web`<br>
cmd `sudo docker build --tag web:latest .`<br>
```
Sending build context to Docker daemon  5.632kB
Step 1/8 : FROM golang:1.11
1.11: Pulling from library/golang
5ae19949497e: Pull complete 
ed3d96a2798e: Pull complete 
f12136850781: Pull complete 
1a9ad5d5550b: Pull complete 
efbd5496b163: Pull complete 
b78da805a02b: Pull complete 
0cd81acaac85: Pull complete 
Digest: sha256:ae3235ae83ae21c2626eb382d6701468c1fe5a0a674bbc5a270f1a636673e7ed
Status: Downloaded newer image for golang:1.11
 ---> ee3ec9ac0398
Step 2/8 : EXPOSE 8081
 ---> Running in 84d57bb3592a
Removing intermediate container 84d57bb3592a
 ---> 4f32fa25b257
Step 3/8 : ENV API_HOST=localhost
 ---> Running in bdc0db808e8a
Removing intermediate container bdc0db808e8a
 ---> e5c9a37718b5
Step 4/8 : ENV API_PORT=8082
 ---> Running in 381ace0310d0
Removing intermediate container 381ace0310d0
 ---> 0c2cd8e5b058
Step 5/8 : WORKDIR $GOPATH/src/github.com/verlandz/docker/web
 ---> Running in 45b40194f086
Removing intermediate container 45b40194f086
 ---> 3e69b283418c
Step 6/8 : RUN mkdir -p $PWD
 ---> Running in a999f880d747
Removing intermediate container a999f880d747
 ---> 9d47ad40fb8a
Step 7/8 : COPY . $PWD
 ---> b462a32260c5
Step 8/8 : CMD ["sh", "-c", "go run $PWD/main.go"]
 ---> Running in 679177365b61
Removing intermediate container 679177365b61
 ---> 3bf32f62f510
Successfully built 3bf32f62f510
Successfully tagged web:latest
```
After that let's run into container.
```
root@201352:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
web                 latest              3bf32f62f510        10 minutes ago      796MB
golang              1.11                ee3ec9ac0398        31 hours ago        796MB
redis               latest              598a6f110d01        9 days ago          118MB

root@201352:~# docker run -d --name my_web -p 8081:8081 web
f1e420c462daf9477092fef00a684d501f4b13e3ff3af3666e5f857e1407ae68

root@201352:~# docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
f1e420c462da        web                 "sh -c 'go run $PWD/…"   4 seconds ago       Up 1 second         0.0.0.0:8081->8081/tcp   my_web
6a191a7285ef        redis               "docker-entrypoint.s…"   3 minutes ago       Up 3 minutes        0.0.0.0:6379->6379/tcp   my_redis
```
![8081-data](https://i.ibb.co/R32hcqz/8081-data.png)
<br>Let's see their logs
```
root@201352:~# docker logs my_web
API_HOST : localhost
API_PORT : 8082
Running in go1.11.12
Listen and serve :8081
requesting to http://localhost:8082/data
```
The path is correct, no error. Why is it fail to connect to API ? Because the localhost of `my_web` and `my_web` is different, that's why it doesn't connect. To solve this, we have to connect them in a `Network`.

## 5. Upload API image to docker and run it
- To upload it you need to create a dockerfile under api folder
- Don't forget to stop api that running in your localhost

cd `$GOPATH/src/github.com/verlandz/docker/api`<br>
cmd `sudo docker build --tag api:latest .`<br>
```
Sending build context to Docker daemon  9.728kB
Step 1/10 : FROM golang:1.11
 ---> ee3ec9ac0398
Step 2/10 : EXPOSE 8082
 ---> Running in e2e93652d889
Removing intermediate container e2e93652d889
 ---> 47d7fe653e32
Step 3/10 : ENV REDIS_HOST=localhost
 ---> Running in 84137f25b81b
Removing intermediate container 84137f25b81b
 ---> bea8ec247dae
Step 4/10 : ENV REDIS_PORT=6379
 ---> Running in 565151419f9c
Removing intermediate container 565151419f9c
 ---> 48de16e7c11b
Step 5/10 : WORKDIR $GOPATH/src/github.com/verlandz/docker/api
 ---> Running in 39f23a71c4fa
Removing intermediate container 39f23a71c4fa
 ---> 4d9080c0c6c9
Step 6/10 : RUN mkdir -p $PWD
 ---> Running in 3cf9fe58c2da
Removing intermediate container 3cf9fe58c2da
 ---> 6c9a04312520
Step 7/10 : COPY . $PWD
 ---> 884652d6d9fc
Step 8/10 : RUN go get -v -u github.com/golang/dep/cmd/dep
 ---> Running in dac91a6bcd67
github.com/golang/dep (download)
github.com/golang/dep/vendor/github.com/armon/go-radix
github.com/golang/dep/gps/paths
github.com/golang/dep/vendor/golang.org/x/sys/unix
github.com/golang/dep/vendor/github.com/Masterminds/semver
github.com/golang/dep/vendor/github.com/Masterminds/vcs
github.com/golang/dep/vendor/github.com/boltdb/bolt
github.com/golang/dep/vendor/github.com/golang/protobuf/proto
github.com/golang/dep/gps/pkgtree
github.com/golang/dep/vendor/github.com/pkg/errors
github.com/golang/dep/internal/fs
github.com/golang/dep/vendor/github.com/jmank88/nuts
github.com/golang/dep/vendor/github.com/nightlyone/lockfile
github.com/golang/dep/vendor/github.com/sdboyer/constext
github.com/golang/dep/vendor/golang.org/x/net/context
github.com/golang/dep/vendor/github.com/pelletier/go-toml
github.com/golang/dep/vendor/golang.org/x/sync/errgroup
github.com/golang/dep/vendor/gopkg.in/yaml.v2
github.com/golang/dep/gps/internal/pb
github.com/golang/dep/gps
github.com/golang/dep/internal/feedback
github.com/golang/dep/gps/verify
github.com/golang/dep
github.com/golang/dep/internal/importers/base
github.com/golang/dep/internal/importers/glide
github.com/golang/dep/internal/importers/glock
github.com/golang/dep/internal/importers/godep
github.com/golang/dep/internal/importers/govend
github.com/golang/dep/internal/importers/govendor
github.com/golang/dep/internal/importers/gvt
github.com/golang/dep/internal/importers/vndr
github.com/golang/dep/internal/importers
github.com/golang/dep/cmd/dep
Removing intermediate container dac91a6bcd67
 ---> 7cb4d32dbcf9
Step 9/10 : RUN dep ensure -v -vendor-only
 ---> Running in b54d44f70b7e
(1/1) Wrote gopkg.in/redis.v5@v5.2.9
Removing intermediate container b54d44f70b7e
 ---> 877510eae902
Step 10/10 : CMD ["sh","-c","go run $PWD/main.go"]
 ---> Running in c3d6e5f68031
Removing intermediate container c3d6e5f68031
 ---> b04f05a0226b
Successfully built b04f05a0226b
Successfully tagged api:latest
```
After that let's run into container.
```
root@201352:~# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
api                 latest              b04f05a0226b        48 seconds ago      848MB
web                 latest              3bf32f62f510        23 minutes ago      796MB
golang              1.11                ee3ec9ac0398        31 hours ago        796MB
redis               latest              598a6f110d01        9 days ago          118MB

root@201352:~# docker run -d --name my_api -p 8082:8082 api
a4ab302386122ba6b29eaa914578ad77fe62e1ffee2d4dbe23ce27ee49cdd4e0

root@201352:~# docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
a4ab30238612        api                 "sh -c 'go run $PWD/…"   4 seconds ago       Up 1 second         0.0.0.0:8082->8082/tcp   my_api
f1e420c462da        web                 "sh -c 'go run $PWD/…"   13 minutes ago      Up 13 minutes       0.0.0.0:8081->8081/tcp   my_web
6a191a7285ef        redis               "docker-entrypoint.s…"   16 minutes ago      Up 16 minutes       0.0.0.0:6379->6379/tcp   my_redis
```
![8082-data](https://i.ibb.co/xqqHZdr/8082-data.png)
<br>Let's see their logs
```
root@201352:~# docker logs my_api
REDIS_HOST : localhost
REDIS_PORT : 6379
Running in go1.11.12
Listen and serve :8082
2019/07/21 08:31:48 dial tcp 127.0.0.1:6379: connect: connection refused
```
The IP and port is correct, but why connection refused ? Because there's no such `127.0.0.1:6379` that running in that container. The container of `my_api` and `my_redis` has it own localhost. To solve this, we have to connect them in a `Network`.

***FAQ***<br>
**Q:** I didn't pull go1.11 into my images. Where it come from ?<br>
**A:** It come from your `dockerfile`. If dockerfile can't find the needed images, it will automatically try to pull.

**Q:** In previous, you state to use diff golang version to version comparison. How to to compare?<br>
**A:** Try to `docker logs my_web`, not like the prev log that state `Running in go1.12.7` but this one state `Running in go1.11.12`

**Q:** Why we get a lot of log when uploading api image to docker ?<br>
**A:** Because it try do `install dep` and `run dep`. It's written in dockerfile.

#### TRIVIA
- [ARG vs ENV](https://vsupalov.com/docker-arg-vs-env/)
- [RUN vs CMD vs ENTRYPOINT](https://stackoverflow.com/questions/37461868/difference-between-run-and-cmd-in-a-docker-file)
- [VAR inside CMD](https://stackoverflow.com/questions/40454470/how-can-i-use-a-variable-inside-a-dockerfile-cmd)
- [Cannot find package](https://stackoverflow.com/questions/47837149/build-docker-with-go-app-cannot-find-package)
- [VAR variance](https://stackoverflow.com/questions/18135451/what-is-the-difference-between-var-var-and-var-in-the-bash-shell)
- [Use WORKDIR value](https://stackoverflow.com/questions/37782505/is-it-possible-to-show-the-workdir-when-building-a-docker-image)
- [.dockerignore](https://docs.docker.com/engine/reference/builder/#dockerignore-file)
- [Explore docker file system](https://stackoverflow.com/questions/20813486/exploring-docker-containers-file-system)


## 6. Network
The purpose of this network is to connect container to container. To connect, the hostname not longer using `localhost` but using the `container name`. The code already designed for this case and can be handled in env var, so we must change env var in the cmd. Let's delete the previous container and build again.

```
root@201352:~# docker stop my_api my_web
my_api
my_web

root@201352:~# docker rm my_api my_web
my_api
my_web

root@201352:~# docker run -t -d -e API_HOST=my_api -e API_PORT=8082 --name my_web -p 8081:8081 web
5547bda29ebcc1d3c65989ff4b67642feb4123f832726cc297482060d3d2633f

root@201352:~# docker run -t -d -e REDIS_HOST=my_redis -e REDIS_PORT=6379 --name my_api -p 8082:8082 api
b09aeaba36f9cb3b10d920045522d4c5fdbe650ec33b4b912acac171048fa17b
```
Let's create a network and connect them.
```
root@201352:~# docker network create my_net
84bb6c32af7cd66bb438ed5939e70a5d2b72c7788d37af52639e3d27a3747390

root@201352:~# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
a76e9a1a00eb        bridge              bridge              local
b54b4beea72c        host                host                local
84bb6c32af7c        my_net              bridge              local
1040cc7c1019        none                null                local

root@201352:~# docker network connect my_net my_redis 
root@201352:~# docker network connect my_net my_api
root@201352:~# docker network connect my_net my_web
```
![8081-data-show](https://i.ibb.co/hcwMB4d/8081-data-show.png)
![8082-data-show](https://i.ibb.co/4VsTqcs/8082-data-show.png)


***FAQ***<br>
**Q:** How to check if container already connect to desired network ?<br>
**A:** Try to inspect the container `docker inspect container_name` then see the value of `Networks` field.


## 7. Storage
The container in docker is `stateless` means doesn't store the data. If you stop and start the container, the data still exist, but if you remove the container and build it again, it will start from fresh new data.

Docker is capable to store data that state in container into storage. This section will focus on 2 types storage:
- `Bind Mount` save data into your file system.
- `Volume` save data into your docker area
<br>You can read more [here](https://docs.docker.com/storage/)

Before we start, we must know `how to mounting` it, because every image has it own way. In redis, you can read [here](https://hub.docker.com/_/redis) and see this section
<br>![redis-way](https://i.ibb.co/KVpnb3w/redis-way.png)<br>
It give the example for Bind Mount `-v /docker/host/dir:/data` and for volume `-v VOLUME:/data`


> Bind Mount

This will store data in `/data/my_redis`
```
root@201352:~# docker stop my_redis
root@201352:~# docker rm my_redis
root@201352:~# docker run -d --name my_redis -p 6379:6379 -v /data/my_redis:/data redis
root@201352:~# docker network connect my_net my_redis

root@201352:~# redis-cli -h localhost -p 6379
localhost:6379> hset user 1 "This is from bind mount"
(integer) 1
localhost:6379> hgetall user
1) "1"
2) "This is from bind mount"
```
![bm](https://i.ibb.co/yfRQqfd/bindmount.png)
<br>remake the container to test
```
root@201352:~# docker stop my_redis
root@201352:~# docker rm my_redis
root@201352:~# docker run -d --name my_redis -p 6379:6379 -v /data/my_redis:/data redis
root@201352:~# docker network connect my_net my_redis
```
![bm](https://i.ibb.co/yfRQqfd/bindmount.png)

> Volume

This will store data in volume with name `my_redis_data`
```
root@201352:~# docker stop my_redis
root@201352:~# docker rm my_redis
root@201352:~# docker volume create my_redis_data 
root@201352:~# docker run -d --name my_redis -p 6379:6379 -v my_redis_data:/data redis
root@201352:~# docker network connect my_net my_redis

root@201352:~# redis-cli -h localhost -p 6379
localhost:6379> hset user 1 "This is from volume"
(integer) 1
localhost:6379> hgetall user
1) "1"
2) "This is from volume"
```
![v1](https://i.ibb.co/3skjDfc/volume1.png)
<br>remake the container to test
```
root@201352:~# docker stop my_redis
root@201352:~# docker rm my_redis
root@201352:~# docker run -d --name my_redis -p 6379:6379 -v my_redis_data:/data redis
root@201352:~# docker network connect my_net my_redis
```
![v1](https://i.ibb.co/3skjDfc/volume1.png)

***FAQ***<br>
**Q:** Which is better `Bind Mount` or `Volume`?<br>
**A:** IMO, `Volume` is better because there's feature to help backup, restore, or migrate data. But in `Bind Mount`, you must do it manually.


## 8. Docker Compose
- Tired of write command one by one manually ?
- Tired of get typo or wrong command in every action ?

Docker compose is the solution !

Docker compose is a tool that help you to run desired docker commands all together. So it's kinda a script that help you to do your thing automatically, rather than manually do a single command for each action. You can see more [here](https://docs.docker.com/compose/)

To try `docker-compose` let's delete existing container,network and volume.<br>
```
root@201352:~# docker stop my_api my_web my_redis
root@201352:~# docker rm my_api my_web my_redis
root@201352:~# docker network rm my_net
root@201352:~# docker volume rm my_redis_data
```
cd `$GOPATH/src/github.com/verlandz/docker` <br>
cmd `sudo docker-compose up -d`
```
Creating network "my_net" with the default driver
Creating volume "my_redis_data" with default driver
Creating my_redis ... done
Creating my_api   ... done
Creating my_web   ... done
```
Then fill the data in `redis-cli`
```
root@201352:~# redis-cli -h localhost -p 6379
localhost:6379> hset user 1 "This is from docker-compose"
(integer) 1
localhost:6379> hgetall user
1) "1"
2) "This is from docker-compose"
```
![dc](https://i.ibb.co/L6bDmHk/docker-compose.png)
<br>Let's test the volume
```
root@201352:~# docker stop my_redis
root@201352:~# docker rm my_redis
root@201352:~# docker run -d --name my_redis -p 6379:6379 -v my_redis_data:/data redis
root@201352:~# docker network connect my_net my_redis
```
![dc](https://i.ibb.co/L6bDmHk/docker-compose.png)
<br>Keep in mind, the above redis's container is no longer monitor by `docker-compose`, because that's a new container. You can `sudo docker-compose ps` to see what containers that currently monitor by `docker-compose`.
```
 Name             Command            State           Ports         
-------------------------------------------------------------------
my_api   sh -c go run $PWD/main.go   Up      0.0.0.0:8082->8082/tcp
my_web   sh -c go run $PWD/main.go   Up      0.0.0.0:8081->8081/tcp
```
You can `sudo docker-compose stop` to stop all containers and start it again by `sudo docker-compose start` or remove it all by `sudo docker-compose down`.

You can see more `docker-compose` commands in `sudo docker-compose help`.