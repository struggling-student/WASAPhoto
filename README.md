# WASAPhoto

> Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! Directly from your PC, you can upload your photos, and they will be  visible to everyone who is following you.

As part of the Web and Software Architecture exam:
- define APIs using the OpenAPI standard
- design and develop the server side ("backend") in Go
- design and develop the client side ("frontend") in JavaScript
- create a Docker container image for deployment

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
## License
See [LICENSE](LICENSE).
