<script>
export default{
    components: {},
    data: function() {
        return {
            username: "",
            profile: {
               id: 0,
               username: "",
            },
        }
    },
    methods: {
        async login() {
            try {
                let response = await this.$axios.post("/session", {username: this.username})
                // save the response data 
                this.profile = response.data
                // save the data in the local storage
                localStorage.setItem("token", this.profile.id);
                localStorage.setItem("username", this.profile.username);
                // redirect to the home page
                this.$router.push({path: '/session'})
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
        }
    },
    mounted() {

    }

}
</script>

<template>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Welcome to WASAPhoto</h1>
    </div>
	<div class="mb-3">
        <label for="username">Username</label>
        <input type="text" id="username" v-model="username">
    </div>
    <div class="mb-3">  
       <button type="button" class="btn btn-sm btn-primary" @click="login">Login</button>
    </div>

</template>

<style>
</style>
