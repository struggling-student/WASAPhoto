<script>
import LogModal from "../components/Logmodal.vue";

export default {
    components: { LogModal },
    data: function () {
        return {
            errormsg: null,
            successmsg: null,
            username: localStorage.getItem('username'),
            token: localStorage.getItem('token'),
            profile: {
                requestId: 0,
                id: 0,
                username: "",
                followersCount: 0,
                followingCount: 0,
                photoCount: 0,
                followStatus: null,
                banStatus: null,
                checkIfBanned: null,
            },
            photoList: {
                requestUser: 0,
                identifier: 0,
                photos: [
                    {
                        id: 0,
                        userId: 0,
                        file: "",
                        date: "",
                        likesCount: 0,
                        commentsCount: 0,
                        likeStatus: null,
                    }
                ],
            },
            comment: "",
            follow: {
                followId: 0,
                followedId: 0,
                userId: 0,
                banStatus: 0,
            },
            ban: {
                banId: 0,
                bannedId: 0,
                userId: 0,
            },
            photoComments: {
                requestIdentifier: 0,
                photoIdentifier: 0,
                identifier: 0,
                comments: [
                    {
                        id: 0,
                        userId: 0,
                        photoId: 0,
                        photoOwner: 0,
                        ownerUsername: "",
                        comment: "",
                        username: "",
                        content: "",
                    }
                ],
            },
        }
    },
    methods: {
        async refresh() {
            await this.userProfile()
            await this.userPhotos()
        },
        async userProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/profile", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.profile = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async userPhotos() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/photo", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoList = response.data
                for (let i = 0; i < this.photoList.photos.length; i++) {
                    this.photoList.photos[i].file = 'data:image/*;base64,' + this.photoList.photos[i].file
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = "User hasen't posted any photos yet."
                    this.detailedmsg = null;
                }
            }
        },
        async banUser(username) {
            try {
                let response = await this.$axios.put("/users/" + username + "/ban/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async unbanUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/ban", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.ban = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/ban/" + this.ban.banId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async followUser(username) {
            try {
                let response = await this.$axios.put("/users/" + username + "/follow/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
                this.successmsg = "User" + username + "followed successfully"
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async unfollowUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/follow", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.follow = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/follow/" + this.follow.followId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

        },
        async sendComment(username, photoid, comment) {
            if (comment === "") {
                this.errormsg = "Emtpy comment field."
            } else {
                try {
                    let response = await this.$axios.put("/users/" + username + "/photo/" + photoid + "/comment/" + Math.floor(Math.random() * 10000), { content: comment }, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token")
                        }
                    })
                    this.clear = response.data
                    this.refresh()
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                        this.detailedmsg = null;
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                        this.detailedmsg = e.toString();
                    } else {
                        this.errormsg = e.toString();
                        this.detailedmsg = null;
                    }
                }
            }
        },
        async openLog(username, photoid) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + photoid + "/comment", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoComments = response.data;
                const modal = new bootstrap.Modal(document.getElementById('logviewer'));
                modal.show();
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async likePhoto(username, id) {
            try {
                let response = await this.$axios.put("/users/" + username + "/photo/" + id + "/like/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
        async deleteLike(username, id) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + id + "/like", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.like = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }

            try {
                let response = await this.$axios.delete("/users/" + username + "/photo/" + id + "/like/" + this.like.likeId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = e.toString();
                    this.detailedmsg = null;
                }
            }
        },
    },
    mounted() {
        this.userProfile()
        this.userPhotos()
    }
}
</script>

<template>
    <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3 sidebar-sticky">
            <h6
                class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
                <span>General</span>
            </h6>
            <ul class="nav flex-column">
                <li class="nav-item">
                    <RouterLink to="/session" class="nav-link">
                        <svg class="feather">
                            <use href="/feather-sprite-v4.29.0.svg#home" />
                        </svg>
                        Home
                    </RouterLink>
                </li>
            </ul>
        </div>
    </nav>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"
        v-if="profile.checkIfBanned == true">
        <div class="alert alert-danger " role="alert">
            <h4 class="alert-heading">Oh no...!</h4>
            <p>User @{{ profile.username }} has banned you, it means that you can't interact with this user anymore
                untill it removes the ban.</p>
            <hr>
            <p class="mb-0"></p>
        </div>

    </div>

    <div>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Profile of {{ profile.username }} </h1>
            <div v-if="profile.checkIfBanned == false" class="p-4 text-black">
                <div class="d-flex justify-content-end text-center py-1">
                    <div>
                        <p class="mb-1 h5">{{ profile.followersCount }}</p>
                        <p class="small text-muted mb-0">Followers</p>
                    </div>
                    <div class="px-3">
                        <p class="mb-1 h5">{{ profile.followingCount }}</p>
                        <p class="small text-muted mb-0">Followings</p>
                    </div>
                    <div>
                        <p class="mb-1 h5">{{ profile.photoCount }}</p>
                        <p class="small text-muted mb-0">Photos</p>
                    </div>
                </div>
            </div>
            <div class="form-group row ">
                <div class="col-md-6">
                    <button type="button" v-if="profile.followStatus == false" class="btn btn-outline-primary "
                        @click="followUser(profile.username)">Follow </button>
                    <button type="button" v-if="profile.followStatus == true" class="btn btn-primary "
                        @click="unfollowUser(profile.username)">Unfollow </button>
                </div>
                <div class="col-md-6">
                    <button type="button" v-if="profile.banStatus == false" class="btn btn-outline-danger"
                        @click="banUser(profile.username)">Ban </button>
                    <button type="button" v-if="profile.banStatus == true" class="btn btn-outline-danger"
                        @click="unbanUser(profile.username)">Unban</button>
                </div>
            </div>
        </div>

        <SuccessMsg v-if="successmsg" :msg="successmsg"></SuccessMsg>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <LogModal id="logviewer" :log="photoComments" :token="token"></LogModal>
        <div v-if="profile.checkIfBanned == false" class="row">
            <div class="col-md-4" v-for="photo in photoList.photos" :key="photo.id">
                <div class="card mb-4 shadow-sm">
                    <img class="card-img-top" :src=photo.file alt="Card image cap">
                    <div class="card-body">
                        <RouterLink :to="'/users/' + profile.username + '/view'" class="nav-link">
                            <button type="button" class="btn btn-outline-primary">{{ profile.username }}</button>
                        </RouterLink>
                        <div
                            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                        </div>
                        <div class="d-flex justify-content-between align-items-center">
                            <p class="card-text">Likes : {{ photo.likesCount }}</p>
                        </div>
                        <div class="d-flex justify-content-between align-items-center">
                            <p class="card-text">Comments : {{ photo.commentsCount }}</p>
                        </div>
                        <p class="card-text">Uploaded on : {{ photo.date }}</p>
                        <div class="input-group mb-3">
                            <input type="text" id="comment" v-model="photo.comment" class="form-control"
                                placeholder="Comment!" aria-label="Recipient's username"
                                aria-describedby="basic-addon2">
                            <div class="input-group-append">
                                <button class="btn btn-primary" type="button"
                                    @click="sendComment(profile.username, photo.id, photo.comment)">Send</button>
                            </div>
                        </div>
                        <div class="d-flex justify-content-between align-items-center">
                            <div class="btn-group">
                                <button type="button" class="btn btn-dark"
                                    @click="openLog(profile.username, photo.id)">Comments</button>
                                <button type="button" v-if="photo.likeStatus == false" class="btn btn-primary"
                                    @click="likePhoto(profile.username, photo.id)">Like</button>
                                <button type="button" v-if="photo.likeStatus == true" class="btn btn-danger"
                                    @click="deleteLike(profile.username, photo.id)">Unlike</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </div>

</template>

<style>

</style>