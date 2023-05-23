  <script>
    import './styles.css'; 

    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          user: null,
          userId: null,
          username:null,
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
            // location.reload();
            this.userId = this.$route.params.userId;
            // console.log(this.userId);
            this.username = this.$route.params.username;
            // console.log(this.username);

            let apiUrl = `/users/${this.$route.params.userId}/profile/${this.$route.params.username}/`;
            let response = await this.$axios.get(apiUrl);
            this.user = response.data;  
            // console.log(response.data);      
            
            // console.log(typeof this.user.id); 
            // console.log(typeof userId);        

          } catch (e) {
            this.errormsg = e.toString();
          }
          this.loading = false;
        },
        async followersView(){
          //TO_DO
          this.$router.push({ name: 'Followings', params: { userId: this.user.id, username: this.username} });
        },
        async followingsView(){
          //TO_DO
          this.$router.push({ name: 'Followers', params: { userId: this.user.id, username: this.username} });
        },
        
        async followUser(id, followedUserId) {
			      this.loading = true;
			      this.errormsg = null
            console.log(id);
            console.log(followedUserId);
            let url = `/users/${id}/following/${followedUserId}/`;
			      try {
			      	await this.$axios.put(url);
			      } catch (e) {
			      	this.errormsg = e.toString();
			      }
			      this.loading = false;
		    },
        async banUser(id, bannedUserId) {
			      this.loading = true;
			      this.errormsg = null;
            let url = `/users/${id}/banned/${bannedUserId}/`;
			      try {
			      	await this.$axios.put(url);
			      } catch (e) {
			      	this.errormsg = e.toString();
			      }
			      this.loading = false;
		    },
      },
      mounted() {
        this.refresh()
      }
    }
    </script>

    <template>
      <div class="background">
        <div class="d-flex justify-content-center align-items-center h-70">
          <div class="card2 bg-white text-dark rounded-5">
            <div class="card-body p-5 text-center">
              <h2 class="fw-bold mb-4 text-uppercase">@{{ this.user?.username }}</h2>
              <div class="info">
                <div >
                  <a href="javascript:" class="text-muted mb-1 larger-text" style="text-decoration: none;" @click="followersView">Followers</a>
                  <p class="mb-0 text-center"> {{ this.user?.followersCount }} </p>
                </div>
                <div>
                  <a href="javascript:" class="text-muted mb-1 larger-text" style="text-decoration: none;" @click="followingsView">Followings</a>
                  <p class="mb-0 text-center"> {{ this.user?.followingCount }} </p>
                </div>
                <div>
                  <p class="text-muted mb-1 larger-text">Photos</p>
                  <p class="mb-0 text-center"> {{ this.user?.photosCount }} </p>
                </div>
                <div v-if="parseInt(this.user?.id) !== parseInt(this.userId)" class="d-grid gap-3 ">
                  <button v-if="!loading" class="btn btn-primary rounded-pill larger-text" type="submit" @click="banUser(this.userId, this.user?.id)" style="background-color: #d10606f5;">Ban</button>
                  <button v-if="!loading" class="btn btn-primary rounded-pill larger-text" type="submit" @click="followUser(this.userId, this.user?.id)" style="background-color: #2e4a78;">Follow</button>
                  <LoadingSpinner v-if="loading" />
                </div> 
              </div>
            </div>
          </div>
        </div>
        <div v-if="this.user?.photos && this.user.photos.length === 0 && parseInt(this.user?.id) === parseInt(this.userId)" class="d-flex justify-content-center align-items-center h-80">
          <router-link :to="{ name: 'UploadPhoto', params: { userId: this.userId }}" class="btn btn-primary rounded-pill larger-text" style="background-color: #2e4a78;">Upload Photo</router-link>
          <!-- <button v-if="!loading" class="btn btn-primary rounded-pill" type="submit" @click="uploadPhoto" style="background-color: #2e4a78;">Add Photo</button> -->
        </div> 
        
        <!-- <div v-else class="photo-grid">
          <div class="photo-card-user" v-for="photo in this.user.photos">
            <img :src="requirePhoto(photo)" class="photo-img" />
          </div>
        </div>  -->
        <ErrorMsg class="error-container" v-if="errormsg" :msg="errormsg" />
      </div>
    </template>
    
    <style>
    </style>
    
    

