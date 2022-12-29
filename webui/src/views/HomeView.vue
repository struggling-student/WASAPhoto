<script>
import LogModal from "../components/Logmodal.vue";

export default {
	components: {LogModal},
	data: function() {
		return {
			username : localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			errormsg: null,
			loading: false,
			some_data: null,
			images: null,
			image: null,
			clear: null,
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
						username: "",
						content: "",
					}
				],
			},
			comment: "",
			stream: {
				identifier: 0,
				photoStream: [
					{
						id: 0,
						userId: 0,
						username: "",
						file: "",
						date: "",
						likeCount: 0,
						commentCount: 0,
						likeStatus: null,
					}
				],
			},
			searchUserUsername: "",
			like : {
				likeId: 0,
				identifier: 0,
				photoIdentifier: 0,
				photoOwner: 0,
			}
		}
	},
	methods: {
		async refresh() {
			this.getStream()
		}, 
		async uploadFile() {
			this.images = this.$refs.file.files[0]
		
		},
		async submitFile() {
			try { 
				let response = await this.$axios.put("/users/" + this.username + "/photo/" + Math.floor(Math.random() * 10000) , this.images, {
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
		async getStream() {
			try { 
                let response = await this.$axios.get("/user/" + this.username + "/stream", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.stream = response.data
                for (let i = 0; i < this.stream.photoStream.length; i++) {
                    this.stream.photoStream[i].file = 'data:image/*;base64,' + this.stream.photoStream[i].file
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
		async SearchUser() {
			this.$router.push({path: '/users/' + this.searchUserUsername + '/view'})
		},
		async sendComment(username, photoid) {
			try { 
                let response = await this.$axios.put("/users/" + username + "/photo/" + photoid + "/comment/" + Math.floor(Math.random() * 10000), {content: this.comment}, {
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
		async likePhoto(username, id){
			try { 
                let response = await this.$axios.put("/users/" + username + "/photo/" + id + "/like/" + Math.floor(Math.random() * 10000),{}, {
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
		async deleteLike(username, id){
			try { 
				let response = await this.$axios.get("/users/" + username + "/photo/" + id + "/like", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.like = response.data
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
				let response = await this.$axios.delete("/users/" + username + "/photo/" + id + "/like/" + this.like.likeId, {
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
		this.getStream()
	}
}
</script>

<template>
	<div>	
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
				</div>
				<div class="btn-group me-2">
					<input type="file" class="btn btn-sm btn-outline-primary" @change="uploadFile" ref="file">
					<button class="btn btn-sm btn-outline-primary" @click="submitFile">Upload!</button>
				</div>
			</div>
		</div>
		<div class="input-group mb-3">
			<input type="text" id="searchUserUsername" v-model="searchUserUsername" class="form-control" placeholder="Search a user in WASAPhoto" aria-label="Recipient's username" aria-describedby="basic-addon2">
			<div class="input-group-append">
				<button class="btn btn-outline-success" type="button" @click="SearchUser">Search</button>
			</div>
		</div>

	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom"></div>
	
	<LogModal id="logviewer" :log="photoComments" :token="token"></LogModal>


    <div class="row">
        <div class="col-md-4" v-for="photo in stream.photoStream" :key="photo.id">
            <div class="card mb-4 shadow-sm">
                <img class="card-img-top" :src=photo.file alt="Card image cap">
                <div class="card-body">
					<p class="card-text">Photo uploaded by {{photo.username}}</p>
                    
                    <div class="d-flex justify-content-between align-items-center">
                        <p class="card-text">Likes : {{photo.likeCount}}</p>
                        <p class="card-text">Comments : {{photo.commentCount}}</p>
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
                            <button type="button" v-if="photo.likeStatus==false" class="btn btn-primary" @click="likePhoto(photo.username, photo.id)">Like</button>
							<button type="button" v-if="photo.likeStatus==true" class="btn btn-danger" @click="deleteLike(photo.username, photo.id)">Unlike</button>
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
