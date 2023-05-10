<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
            image: null
		}
	},
	methods: {
        onFileSelected(event) {
            this.image = event.target.files[0]
        },
		async uploadPhoto() {
			this.loading = true;
			this.errormsg = null;
			try {
                let config = {
                    headers: {
                        Authorization: "Bearer " + this.$cookies.get("token"),
                        "Content-Type": "image/png"
                    }
                }
				await this.$axios.post("/Image", this.image, config);
			} catch (e) {
                
				this.errormsg = e.toString();
			}
			this.loading = false;
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
                    <div class="card-body p-5 text-center">
                        <h2 class="fw-bold mb-2 text-uppercase">Upload Photo</h2>
                            <div class="form-outline form-white mb-4">

                                <p class="text-white-50">Upload your photo!</p>
                                <input type="file" id="uploadedImage"  @change="onFileSelected"/>

                            </div>
                            <button v-if="!loading" class="btn btn-outline-light btn-lg" type="submit" @click="uploadPhoto">Upload</button>
                            <LoadingSpinner v-if="loading"></LoadingSpinner>
                        </div>        
                    </div>
                    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
                </div>
                
                </div>
            </div>
        </div>
    </template>

<style>
</style>
