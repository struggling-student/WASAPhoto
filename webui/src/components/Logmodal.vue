<script>
export default {
	props: ["log","token"],
	data() {
		return {
		}
	},

	methods: {
		async deleteComment(username, photoid, commentid ) {
			try {
				let response = await this.$axios.delete("/users/" + username + "/photo/" + photoid + "/comment/" + commentid, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				location.reload();
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
    }
}
</script>

<template>
	<div class="modal modal-xl" tabindex="-1">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">Comments</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<div class="row">
						<div class="col-md-4" v-for="comment in log.comments" :key="comment.id">
							<div class="card">
								<div class="card-body">
									<h5 class="card-title">Comment by {{ comment.username }}</h5>
									<p class="card-text">{{ comment.content }}</p>
									<button type="button" v-if="token==comment.userId"  data-bs-dismiss="modal" class="btn btn-danger" @click="deleteComment(comment.ownerUsername, comment.photoId,comment.id)">Delete</button>
								</div>
							</div>
							</div>
					</div>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
.modal textarea {
	font-family: "DejaVu Sans Mono", monospace;
}
</style>