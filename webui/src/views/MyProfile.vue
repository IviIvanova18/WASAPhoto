<script>
	export default {
		data: function() {
			return {
				errormsg: null,
				loading: false,
                user: null,
                lastDeleted: 0
			}
		},
		methods: {
            async refresh() {
                this.refresh2()
            },
			async refresh2() {
                let data = null
				this.loading = true;
                this.errormsg = null;
                try {
                    let url = "/Users/" + this.$cookies.get("token") + "/Profile/" + this.$cookies.get("username")
                    let config = {
                        headers: {
                            Authorization: "Bearer " + this.$cookies.get("token")
                        }
                    }
                    let response = await this.$axios.get(url, config);
                    data = null
                    console.log("data1", data)
                    data = response.data
                    console.log("data2", data)
                    data.nofollowers = data.followers.length
                    data.nofollowings = data.followings.length
                   
                    let index = 0
                    data.photoesDetails = []
                    
                    while (index < data.nophotoes) {
                        console.log("MAMA", data.photoes[index])
                        data.photoesDetails.push(await this.getPhotoesDetails(data.photoes[index]))
                        let aux = data.photoesDetails[index].date.toString()
                        data.photoesDetails[index].date = aux.substring(0, 9) + " " + aux.substring(11, 19)
                        index = index + 1
                    }
                    this.user = data
                    console.log("this is user", data)
                } catch (e) {
                        this.errormsg = e.toString();
                }
                this.loading = false;
            },
            async getPhotoesDetails(id) {
                this.loading = true;
                this.errormsg = null;
                try {
                    let url = "/Images/" + id + "/Details"
                    let config = {
                        headers: {
                            Authorization: "Bearer " + this.$cookies.get("token")
                        }
                    }  
                    let response = await this.$axios.get(url, config);
                    this.loading = false;
                    return response.data
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
            },
            async deletePhoto(id) {
                this.lastDeleted = id
                this.loading = true;
                this.errormsg = null;
                try {
                    let url = "/Images/" + id
                    let config = {
                        headers: {
                            Authorization: "Bearer " + this.$cookies.get("token")
                        }
                    }  
                    await this.$axios.delete(url, config);
                    
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
                this.$router.push({name: 'MyProfile'})
                this.refresh()
            }
        },
        mounted() {
            this.refresh()
        }
    }
	</script>


<template>
    <div>
  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <section class="vh-100">
        <div class="container py-5 h-100">
          <div class="row d-flex justify-content-center align-items-center h-100">
            <div class="col col-md-9 col-lg-7 col-xl-5">
              <div class="card" style="border-radius: 15px;">
                
                <div class="card-body p-4">
                  <div class="d-flex text-black">
                    <div class="flex-shrink-0">

                    </div>
                    <div class="flex-grow-1 ms-3">
                      <h5 class="text-center mb-1" style="padding-bottom: 10px"> <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#at-sign"/></svg>{{ this.$cookies.get("username") }} </h5>
                      <div class="d-flex justify-content-evenly rounded-3 p-2 mb-2"
                        style="background-color: #efefef;">
                        <div >
                          <p class="text-muted mb-1">Followers</p>
                          <p class="mb-0 text-center"> {{ this.user?.nofollowers }} </p>
                        </div>
                        <div>
                          <p class="text-muted mb-1">Followings</p>
                          <p class="mb-0 text-center"> {{ this.user?.nofollowings }} </p>
                        </div>
                        <div>
                          <p class="text-muted mb-1">Photoes</p>
                          <p class="mb-0 text-center"> {{ this.user?.nophotoes }} </p>
                        </div>
                      </div>  

                      </div>
                  </div>
<!--  -->
  <div class="card" style="max-width: 35rem;" v-if="!loading" v-for="ph in user?.photoesDetails" >
  <div class="card-body">
    <div class="d-flex mb-3">
    <div>
      <span class="text-dark mb-0">
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center">
        <span class="d-flex justify-content-end flex-wrap flex-md-nowrap align-items-center">
      <strong><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#at-sign"/></svg>{{ ph?.username }}</strong>
      </span>
      <span  style="padding-left:20rem">
      <button type="button" class="btn btn-outline-danger border-white" @click="deletePhoto(ph?.id)"><span>
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x"/></svg></span></button>
    </span>
    </div>
      
        <div>
      <small class="text-muted">{{ ph?.date }}</small>
        </div>
      </span>

    </div>
    </div>
  </div>
  <div class="bg-image hover-overlay ripple rounded-0" data-mdb-ripple-color="light">
    <ImageWithAuthorization v-if="lastDeleted != ph?.id" :id="ph?.id"></ImageWithAuthorization>
    <a href="#!">
    <div class="mask" style="background-color: rgba(251, 251, 251, 0.2)"></div>
    </a>
  </div>
  <div class="card-body">
    <div class="d-flex justify-content-between mb-3">
    <div>
      <a href="">
        <span>{{ ph?.no_likes }}</span>
        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#thumbs-up"/></svg>
      </a>
    </div>
    <div>
      <a href="" class="text-muted"> {{ ph?.no_comments }} </a>
    </div>
    </div>
    <div class="d-flex justify-content-between text-center border-top border-bottom mb-4">
    <LikeButton :id="ph?.id" :is_liked="ph?.is_liked" @doRefresh="refresh"></LikeButton>
    <RouterLink :to="{ name: 'CommentPage', params: { id: ph.id, username: ph.username }}"><button type="submit" class="btn btn-light btn-lg">
      <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg> Comment
    </button></RouterLink>

  </div>
    
  </div>
  </div>
  <!--  -->

                </div>
              </div>
            </div>
          </div>
        </div>
        </section>
    </div>
    </template>

<style>
</style>
