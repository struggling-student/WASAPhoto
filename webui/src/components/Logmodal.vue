
<script>
export default {
	props: ["photoComments"],
	data() {
		return {
		}
	},
}
</script>

<template>
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">Comments of photo uploaded by </h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<textarea readonly style="width: 100%;" rows="25">{{ photoComments }}</textarea>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close comments</button>
				</div>
			</div>
		</div>
</template>

<style>
.modal textarea {
	font-family: "DejaVu Sans Mono", monospace;
}
</style>