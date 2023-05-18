  <script>
    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          user: null,
        }
      },
      methods: {
        load() {
          return load
        },
        async refresh() {
          this.loading = true;
          this.errormsg = null;
          try {
            let userId = this.$route.params.userId;
            let username = this.$route.params.username;
            let apiUrl = `/users/${userId}/profile/${username}/`;
            let response = await this.$axios.get(apiUrl);
            console.log(response.data);
            this.user = response.data;        
            
          } catch (e) {
            this.errormsg = e.toString();
          }
          this.loading = false;
        },
        async followersView(){
          this.$router.push({ name: 'Followings', params: { userId: this.$route.params.userId} });
          
        },
        async followingsView(){
          this.$router.push({ name: 'Followings', params: { userId: this.$route.params.userId} });
        },
        async followUser(){

        }
        
      },
      mounted() {
        this.refresh()
      }
    }
    </script>

  
<template>
    <div class="background">
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
                      <h5 class="text-center mb-1" style="padding-bottom: 10px"> <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#at-sign"/></svg>{{ this.user?.username }} </h5>
                      <div class="d-flex justify-content-evenly rounded-3 p-2 mb-2"
                        style="background-color: #80a2d9;">
                        <div >
                          <a href="javascript:" class="text-muted mb-1" @click="followersView">Followers</a>
                          <p class="mb-0 text-center"> {{ this.user?.followersCount }} </p>
                        </div>
                        <div>
                          <a href="javascript:" class="text-muted mb-1" @click="followingsView">Followings</a>
                          <p class="mb-0 text-center"> {{ this.user?.followingCount }} </p>
                        </div>
                        <div>
                          <p class="text-muted mb-1">Photos</p>
                          <p class="mb-0 text-center"> {{ this.user?.photosCount }} </p>
                        </div>
                        <div class="d-grid gap-3">
                          <button v-if="!loading" class="btn btn-primary rounded-pill" type="submit" @click="followUser" style="background-color: #2e4a78;">Follow</button>
  
                          <LoadingSpinner v-if="loading" />
                        </div>
                      </div>  

                      </div>
                  </div>
  

  <!-- <div class="card" style="max-width: 35rem;" v-if="!loading" v-for="ph in user?.photos" >
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
  </div> -->

                </div>
              </div>
            </div>
          </div>
        </div>
        </section>
    </div>
    </template>

<style>
.background {
  background-color: #f3f3f3;
  background-image: linear-gradient(to bottom right, #80a2d9, #ffffff);
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>