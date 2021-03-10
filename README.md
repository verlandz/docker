# Docker

### Introduction
- [What is docker ?](https://opensource.com/resources/what-docker) 
- [What is docker container ?](https://www.sdxcentral.com/containers/definitions/what-is-docker-container/)
- [Why docker ?](https://www.docker.com/why-docker)
- [Official Docs](https://docs.docker.com/)

### Purpose
Make a simple project with `docker` and help to learn the basics.

### Objective
This project will make 2 containers: 1 app & 1 datastore(redis). The app will get data from redis and show it in the response. For redis image, we can download it from the official [docker registry](https://hub.docker.com/_/redis), but for app image, we will make it by ourself. The language to make the app is `Golang` aka `Go`. 

***You can use another language than `Go`, but this project will focus on using `Go`.***

### Author Setup
- OS: Ubuntu 16.04.6 LTS
- Docker: Docker version 18.09.7, build 2d0083d (Docker Engine - Community)
- Docker-Compose: docker-compose version 1.24.1, build 4667896b
- Docker-Machine: docker-machine version 0.16.0, build 702c267f
- Golang: go version go1.13 linux/amd64

### Docker Images
- [redis:6.0.9](https://hub.docker.com/layers/redis/library/redis/6.0.9/images/sha256-e4b1fffb060afd6f31955f7af1ac7e68270fdc3c4c798ec3a93def617c68f481?context=explore)
- [go:1.11](https://hub.docker.com/layers/golang/library/golang/1.11/images/sha256-cdb2c594a968289dcb9b7f6d3ec31820f9c9dc5687dd62ce8f34e923bd39a2b3?context=explore)
- [verlandz/docker-app:1.0](https://hub.docker.com/repository/docker/verlandz/docker-app) -> we'll create this

***It's a common pratice to not use latest tag***

### Prerequisite
- [Install docker](https://phoenixnap.com/kb/how-to-install-docker-on-ubuntu-18-04)
- [Install docker-compose](https://docs.docker.com/compose/install/)
- [Install docker-machine](https://docs.docker.com/machine/install-machine/)
- [Install golang](https://tecadmin.net/install-go-on-ubuntu/) (except go1.11)
- [Basic golang](https://golang.org/doc/)
- [Basic cmd redis](https://redis.io/)

Usually docker need to be super user, so try to `sudo su -` first.

***FAQ***<br>
**Q:** Why we need to install go to our device if we can pull go image to our docker ?<br>
**A:** The reason is for comparison version. So in docker, we're using go1.11 but in local we're using another version. Also having install go in your device make a guarantee that you understand how to manage `GOPATH` and `GOROOT`, we need it to undestand for `dockerfile`.

# Walkthrough
You can follow from this [docs](https://github.com/verlandz/docker/tree/master/docs)<br>
***If you don't see `root@verlandz` in sample cmd, that's mean it's not running on the `root`***

a
a
