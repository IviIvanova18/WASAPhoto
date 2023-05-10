<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			otheruser: null
		}
	},
	methods: {
		async findUser() {
			this.loading = true;
			this.errormsg = null;
			try {
                let url = "/Users/" + this.$cookies.get("token") + "/Profile/" + this.otheruser
                let config = {
                    headers: {
                        Authorization: "Bearer " + this.$cookies.get("token")
                    }
                }
				let response = await this.$axios.get(url, config);
                if (this.$cookies.get("username") === this.otheruser) {
                    this.$router.push({ name: 'MyProfile'})
                } else {
                    this.$router.push({ name: 'Profile', params: { uid: this.$cookies.get('token'), username: this.otheruser }})
                }
			} catch (e) {
                if (e.response.status == 404) {
                    this.errormsg = "There is no user " + this.otheruser
                } else if (e.response.status == 401) {
                    this.$router.push({ name: 'Login'})
                } else {
                    this.errormsg = e.toString();
                }
			}
			this.loading = false;
            this.otheruser = ""
		},
	}
}
</script>

<template>
    <div>
        <div class="container py-5 h-100">
            <div class="row d-flex justify-content-center align-items-center h-100">
            <div class="col-12 col-md-8 col-lg-6 col-xl-5">
                <div class="card bg-dark text-white">
                    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
                <div class="card-body p-5 text-center">
                    <h2 class="fw-bold mb-2 text-uppercase">Search User</h2>
                    <p class="text-white-50">Please enter the name of the user you are searching!</p>
                    <!-- <form @submit.prevent="loginUser"> -->
                        <div class="form-outline form-white mb-4">
                            <input type="text" id="otheruser" class="form-control form-control-lg" v-model="otheruser" placeholder="username" required minlength="3" maxlength="16" />
                            <label class="form-label" for="otheruser">Username</label>
                        </div>
                        <button v-if="!loading" class="btn btn-outline-light btn-lg" type="submit" @click="findUser">Search</button>
                    <!-- </form> -->
                        <LoadingSpinner v-if="loading"></LoadingSpinner>
                    </div>          
                </div>
            </div>
            </div>
        </div>
    </div>
</template>

<style>
</style>
