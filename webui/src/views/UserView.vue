<script>
export default {
	data: function() {
		return {
            profile: {
               requestId: 0,
               id: 0,
               username : "",
               followersCount: 0,
               followingCount: 0,
               photoCount: 0,
               followStatus: null,
               banStatus: null,
            },

            photoList: {
            requestUser: 0,
            identifier : 0,
            photos: [
                {
                    id: 0,
                    userId: 0,
                    file : "",
                    date: "",
                    likesCount: 0,
                    commentsCount: 0,
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
        
		}
	},
	methods: {
        async refresh() {
            await this.userProfile()
            await this.userPhotos()
        },
        async userProfile() {
            try { 
                let response = await this.$axios.get("users/" + this.$route.params.username + "/profile",{
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.profile = response.data
            } catch(e) {
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
                let response = await this.$axios.get("/users/" + this.$route.params.username  + "/photo", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoList = response.data
                for (let i = 0; i < this.photoList.photos.length; i++) {
                    this.photoList.photos[i].file = 'data:image/*;base64,' + this.photoList.photos[i].file
                }
            } catch(e) {
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

        async banUser(username) {
            try { 
                let response = await this.$axios.put("/users/" + username + "/ban/" + Math.floor(Math.random() * 10000),{}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
				this.clear = response.data
				this.refresh()
            } catch(e) {
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
            } catch(e) {
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
            } catch(e) {
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
                let response = await this.$axios.put("/users/" + username + "/follow/" + Math.floor(Math.random() * 10000),{}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
				this.clear = response.data
				this.refresh()
            } catch(e) {
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
            } catch(e) {
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
            } catch(e) {
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
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">User profile of {{ profile.username }} </h1>
    </div>
    <div class="btn-toolbar mb-2 mb-md-0 ">
                        <button type="button" v-if="profile.followStatus==false" class="btn btn-primary " @click="followUser(profile.username)">Follow {{profile.username }}</button>
                        <button type="button" v-if="profile.followStatus==true" class="btn btn-outline-danger " @click="unfollowUser(profile.username)">Unfollow</button>
                        <button type="button" v-if="profile.banStatus==false" class="btn btn-danger" @click="banUser(profile.username)">Ban{{profile.username}}</button>
                        <button type="button" v-if="profile.banStatus==true" class="btn btn-outline-danger" @click="unbanUser(profile.username)">Unban {{profile.username}}</button>
            </div>
    <div class="d-flex justify-content-between align-items-center">
        <h4>
        Followers : {{profile.followersCount}}
    </h4>
    <h4>
        Following : {{profile.followingCount}}
    </h4>
    <h4>
        Photos : {{profile.photoCount}}
    </h4>

    </div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
    <div class="row">
        <div class="col-md-4" v-for="photo in photoList.photos" :key="photo.id">
            <div class="card mb-4 shadow-sm">
                <img class="card-img-top" :src=photo.file alt="Card image cap">
                <div class="card-body">
					<p class="card-text">Photo uploaded by {{profile.username}}</p>
                    
                    <div class="d-flex justify-content-between align-items-center">
                        <p class="card-text">Likes : {{photo.likesCount}}</p>
                        <p class="card-text">Comments : {{photo.commentsCount}}</p>
                    </div>
					<p class="card-text">Photo uploaded on {{photo.date}}</p>
					<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
					
					
					<div class="input-group mb-3">
						<input type="text" id="comment" v-model="comment" class="form-control" placeholder="Comment!" aria-label="Recipient's username" aria-describedby="basic-addon2">
						<div class="input-group-append">
							<button class="btn btn-outline-success" type="button" @click="sendComment(photo.username, photo.id)">Send</button>
						</div>
					</div>

                    <div class="d-flex justify-content-between align-items-center">
                        <div class="btn-group">
                            <button type="button" class="btn btn-sm btn-outline-secondary"  @click="openLog(photo.username, photo.id)">View comments</button>
                            <button type="button" class="btn btn-outline-primary" @click="">Like Photo</button>
                        </div>
                    </div>
                    </div>
                </div>
                </div>
            </div>
</template>

<style>
</style>