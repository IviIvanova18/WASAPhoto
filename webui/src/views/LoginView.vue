<script>
    export default {
        data: function() {
            return {
                errormsg: null,
                loading: false,
                username: null,
            }
        },
        methods: {
            async loginUser() {
                this.loading = true;
                this.errormsg = null;
                try {
                    let response = await this.$axios.post("/session", {
                        username: this.username
                    }, {
                        "Content-Type": "application/json"
                    });
                    this.$cookies.set("token", response.data.id)
                    this.$cookies.set("username", response.data.username)
                    this.$router.push({ name: 'MyProfile'})
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
                        <h2 class="fw-bold mb-2 text-uppercase">Login</h2>
                        <p class="text-white-50">Please enter your username!</p>
                            <div class="form-outline form-white mb-4">
                                <input type="text" id="username" class="form-control form-control-lg" v-model="username" placeholder="username" required minlength="3" maxlength="16" />
                                <label class="form-label" for="username">Username</label>
                            </div>
                            <button v-if="!loading" class="btn btn-outline-light btn-lg" type="submit" @click="loginUser">Login</button>
                            <LoadingSpinner v-if="loading"></LoadingSpinner>
                        </div>          
                    </div>
                </div>
                </div>
            </div>
            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        </div>
    </template>
    
    <style>
    </style>
    