# WASAPhoto

> Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! Directly from your PC, you can upload your photos, and they will be  visible to everyone who is following you.

As part of the Web and Software Architecture exame, i will :
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

## OpenAPI

## Backend
## Frontend
## Deployment