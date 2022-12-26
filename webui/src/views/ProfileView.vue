<script>
export default{
    components: {},
    data: function() {
        return {
            username : localStorage.getItem('username'),
            newUsername : "",
            profile: {
               requestId: 0,
               id: 0,
               username : "",
               followersCount: 0,
               followingCount: 0,
               photoCount: 0,
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
           user: {
               id: 0,
               username: "",
            },
        }
    },
    methods: {
        async userProfile() {
            try { 
                let response = await this.$axios.get("/users/" + this.username + "/profile", {
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
                let response = await this.$axios.get("/users/" + this.username + "/photo", {
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
        async deletePhoto(photoid) {
            try { 
                let response = await this.$axios.delete("/users/" + this.username + "/photo/" + photoid, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
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
        async getComments(){

        },
        async changeName() {
            try { 
                let response = await this.$axios.put("/user/" + this.username + "/setusername", {username: this.newUsername}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.user= response.data
                localStorage.setItem("username", this.user.username);
                this.$router.push({path: '/users/' + this.user.username + '/profile'})
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
        async logout() {

        }
    },
    mounted() {
        this.userProfile()
        this.userPhotos()
    }

}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">User profile of {{username}} </h1>
        <div class="btn-toolbar mb-2 mb-md-0">
                    <button type="button" class="btn btn-outline-danger" @click="logout">
						Logout
					</button>
        </div>
    </div>
    <div class="input-group mb-3">
			<input type="text" id="newUsername" v-model="newUsername" class="form-control" placeholder="Insert a new username for your profile..." aria-label="Recipient's username" aria-describedby="basic-addon2">
			<div class="input-group-append">
				<button class="btn btn-outline-success" type="button" @click="changeName" >Change Username</button>
			</div>
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
                    <p class="card-text">Photo uploaded on {{photo.date}}</p>
                    <div class="d-flex justify-content-between align-items-center">
                        <p class="card-text">Likes : {{photo.likesCount}}</p>
                        <p class="card-text">Comments : {{photo.commentsCount}}</p>
                    </div>
                    <div class="d-flex justify-content-between align-items-center">
                        <div class="btn-group">
                            <button type="button" class="btn btn-sm btn-outline-secondary">View photo</button>
                            <button type="button" class="btn btn-sm btn-outline-secondary">View comments</button>
                            <button type="button" class="btn btn-sm btn btn-outline-danger" @click="deletePhoto(photo.id)">Delete Photo</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>


</template>

<style>
.card-img-top {
    width: 100%;
    height: 15vw;
    object-fit: cover;
}
</style>
