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
		}
	},
	methods: {
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
            <div class="btn-toolbar mb-2 mb-md-0 ">
                        <button type="button" class="btn btn-primary " @click="">Follow</button>
                        <button type="button" class="btn btn-danger" @click="">Ban</button>
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