# WASAPhoto [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
> Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! Directly from your PC, you can upload your photos, and they will be  visible to everyone who is following you.

</div> 
<div align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white">
  <img src="https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white">
  <img src="https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white">
  <img src="https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=F7DF1E">
  <img src="https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D">
  <img src="https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white">
  <img src="https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white">
  <img src="https://img.shields.io/badge/GIT-E44C30?style=for-the-badge&logo=git&logoColor=white">
</div

[![Screenshot-2023-01-16-at-00-10-22.png](https://i.postimg.cc/kggFZbkj/Screenshot-2023-01-16-at-00-10-22.png)](https://postimg.cc/0bhwKb5m)

As part of the Web and Software Architecture exam:
- define APIs using the OpenAPI standard.
- design and develop the server side ("backend") in Go.
- design and develop the client side ("frontend") in JavaScript.
- create a Docker container image for deployment.

## Description 
Each user will be presented with a stream of photos in reverse chronological order, with information about when it was uploaded and how many likes and comments it has. Users can place(and later remove) a "like" to photos from other users. Also, users can add comments to any photo(even those uploaded by themself). Comments can be removed by the author only.

Uses can ban other users. If Alice (user) bans Eve(user), Eve won't be able to see any information about Alice. Alice can decide to remove the ban at any moment.

Users will have their own personal profiles. The personal profile page for user detail: user's photos(in revers chronological order), their followers and following, and how many photos have been uploaded. Each user can change his/her own username, upload photos, remove photos, and follow/unfollow other users. Removal of a photo will also remove likes and comments.

User profiles can be searched via username.

The login is perfomed by specifying the username.
## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding
	
## OpenAPI
### How to view the OpenAPI documentation 
Visit the [Swagger editor](https://editor.swagger.io/) website and paste the API documentation inside `doc/api.yaml`.

## Backend
### How to build 
If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:
```
go build ./cmd/webapi/
```
If you're using the WebUI and you want to embed it into the final executable:
```
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## Frontend
### How to run (in development mode)
You can launch the backend only using:
```
go run ./cmd/webapi/
```
If you want to launch the WebUI, open a new tab and launch:
```
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

## Deployment
### How to build the images 
Backend
```
$ docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```
Frontend 
```
$ docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```
### How to run the container images
Backend
```
$ docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```
Frontend
```
$ docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```
## Website showcase
### View User profile
>User has a profile view, where he is able to change it's username, delete photos and view the followers/followings/photo count. User is also able to perform different actions on each image.

[![Screenshot-2023-01-16-at-00-01-44.png](https://i.postimg.cc/yNJx31Z6/Screenshot-2023-01-16-at-00-01-44.png)](https://postimg.cc/5HVbZMgD)
### View another user profile
>User can view the profile of another user, the view contains photos,followers/followings/photo count and the follow and ban buttons. User is also able to perform different actions on each image.

[![Screenshot-2023-01-16-at-00-02-04.png](https://i.postimg.cc/JhJmWGXX/Screenshot-2023-01-16-at-00-02-04.png)](https://postimg.cc/4KJj64B4)
## For more pictures...
check [this link](https://postimages.org/). 

#
<img src="http://ForTheBadge.com/images/badges/built-with-love.svg"> 
Last Update on : 16/01/2023
