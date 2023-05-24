  <script>
    import '../styles.css'; 

    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          user: null,
          userId: null,
          username:null,
          comments:{},
          likes:{},
          userLikes:{},
          followTag:false,
          banTag:false,
          newComment: null,
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
            this.userId = this.$route.params.userId;
            this.username = this.$route.params.username;
            let apiUrl = `/users/${this.$route.params.userId}/profile/${this.$route.params.username}/`;
            let response = await this.$axios.get(apiUrl);
            this.user = response.data;  
            for (const photo of this.user.idPhotos) {
              const commentsResponse = await this.$axios.get(`/photos/${photo}/comments/`);         
              this.comments[photo] = commentsResponse.data;
              const likesResponse = await this.$axios.get(`/photos/${photo}/likes/`); 
              this.likes[photo] = likesResponse.data.length;               
              this.userLikes[photo] = likesResponse.data.map(p => p.idUser).includes(parseInt(this.userId));
            }

          } catch (e) {
            if (e.response.status == 404) {
              this.errormsg = "You can not acces user" + this.username
            } else if (e.response.status == 401) {
              this.$router.push({ name: 'Login'})
            } else {
              this.errormsg = e.toString();
            }
          }
          this.loading = false;
        },
        async likePhoto(photoId, userId){
            this.loading = true;
			      this.errormsg = null
			      try {
              if (this.userLikes[photoId]) {
                await this.$axios.delete(`/photos/${photoId}/likes/${userId}/`) 
                this.userLikes[photoId] = false
              } else {
                await this.$axios.put(`/photos/${photoId}/likes/${userId}/`)
                this.userLikes[photoId] = true
              }
              await this.refresh();
			      } catch (e) {
			      	this.errormsg = e.toString();
			      }
			      this.loading = false;
        },
        async deleteComment(photoId,commentId){
          this.loading = true;
			    this.errormsg = null;
			    try {
			    	await this.$axios.delete(`/photos/${photoId}/comments/${commentId}/`);
			    	await this.refresh();
			    } catch (e) {
			    	this.errormsg = e.toString();
			    }
			    this.loading = false;
        },
        addComment: async function (photoId){
			    this.loading = true;
			    this.errormsg = null;
			    try {
			    	await this.$axios.post(`/photos/${photoId}/comments/`, {
              idUser: parseInt(this.userId),
			    		comment: this.newComment,
			    	});
            await this.refresh();
            this.newComment = "";
			    } catch (e) {
			    	this.errormsg = e.toString();
			    }
			    this.loading = false;
		    },
        async followersView(id,username){
         
          this.$router.push({ name: 'Followers', params: { userId: id, username: username} });
        },
        async followingsView(id,username){
         
          this.$router.push({ name: 'Followings', params: { userId: id, username: username} });
        },
        
        async followUser(id, followedUserId) {
          let response = await this.$axios.get(`/users/${id}/`);
          let follower = response.data.username;           
          this.followTag = this.user.followers.includes(follower);
			    
          this.loading = true;
			    this.errormsg = null;
			    try {
            if (this.followTag) {
              await this.$axios.delete(`/users/${id}/following/${followedUserId}/`) 
              this.followTag = false
            } else {
              await this.$axios.put(`/users/${id}/following/${followedUserId}/`)
              this.followTag = true
            }
            await this.refresh();
			    } catch (e) {
			    	this.errormsg = e.toString();
			    }
			    this.loading = false;
		    },
        async banUser(id, bannedUserId) {
          let response = await this.$axios.get(`/users/${bannedUserId}/`);
          let bannedUser = response.data.username; 
          response = await this.$axios.get(`/users/${id}/banned/`);
          let banned = response.data; 		    
          this.loading = true;
			    this.errormsg = null;
			    try {
            if (this.banTag) {
              await this.$axios.delete(`/users/${id}/banned/${bannedUserId}/`) 
              this.banTag = false
            } else {
              await this.$axios.put(`/users/${id}/banned/${bannedUserId}/`)
              this.banTag = true
            }
            await this.refresh();
			    } catch (e) {
			    	this.errormsg = e.toString();
			    }
			    this.loading = false;
		    },
        requirePhoto(idPhoto) {
          
          return `../src/assets/images/${this.user.photos[this.user.idPhotos.indexOf(idPhoto)]}`;
        },
        async deletePhoto(photo){
          this.loading = true;
			    this.errormsg = null;
			    try {
			    	await this.$axios.delete(`/photos/${photo}/`);
			    	await this.refresh();
			    } catch (e) {
			    	this.errormsg = e.toString();
			    }
			    this.loading = false;
          await this.refresh();


        }
      },
      mounted() {
        this.refresh()
      }
    }
    </script>

    <template>
      <div class="background">
        <div class="col-md-3 col-lg-2 empty-col"></div>
        <div class="col-md-9 col-lg-10">
          <div class="d-flex justify-content-center align-items-center h-70">
            <div class="card2 bg-white text-dark rounded-5">
              <div class="card-body p-5 text-center">
                <h2 class="fw-bold mb-4 text-uppercase">@{{ this.user?.username }}</h2>
                <div class="info">
                  <div>
                    <a href="javascript:" class="text-muted mb-1 larger-text" style="text-decoration: none;" @click="followersView(this.userId,this.user?.username)">Followers</a>
                    <p class="mb-0 text-center"> {{ this.user?.followersCount }} </p>
                  </div>
                  <div>
                    <a href="javascript:" class="text-muted mb-1 larger-text" style="text-decoration: none;" @click="followingsView(this.userId,this.user?.username)">Followings</a>
                    <p class="mb-0 text-center"> {{ this.user?.followingCount }} </p>
                  </div>
                  <div>
                    <p class="text-muted mb-1 larger-text">Photos</p>
                    <p class="mb-0 text-center"> {{ this.user?.photosCount }} </p>
                  </div>
                </div>
                <div v-if="parseInt(this.user?.id) !== parseInt(this.userId)" class="gap-3 ">
                  <button v-if="!loading" class="btn btn-primary btn-block rounded-pill larger-text" type="submit" @click="banUser(this.userId, this.user?.id)" style="background-color: #d10606f5;"> {{ banTag ? 'Unban' : 'Ban' }}</button>
                  <button v-if="!loading" class="btn btn-primary btn-block rounded-pill larger-text" type="submit" @click="followUser(this.userId, this.user?.id)" style="background-color: #2e4a78;"> {{ followTag ? 'Unfollow' : 'Follow' }}</button>
                  <LoadingSpinner v-if="loading" />
                </div>              
              </div>
            </div>
          </div>
          <div v-if="parseInt(this.user?.id) === parseInt(this.userId)" class="d-flex justify-content-center align-items-center h-80">
            <router-link :to="{ name: 'UploadPhoto', params: { userId: this.userId }}" class="btn btn-primary rounded-pill larger-text" style="background-color: #2e4a78;">Upload Photo</router-link>
            <router-link :to="{ name: 'SearchUser', params: { userId: this.userId }}" class="btn btn-primary rounded-pill larger-text" style="background-color: #2e4a78;">Search User</router-link>
            
          </div> 
          
          <div class="photo-grid d-flex justify-content-center align-items-center h-80">
            <div class="col-9 col-sm-6 col-md-4 col-lg-3 mb-5" v-for="photo in this.user?.idPhotos" :key="photo.id">
              <div class="photo-card bg-white text-dark rounded-5 position-relative">
                <div class="dropdown position-absolute top-0 end-0 p-2">
                  <button class="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#more-vertical"/></svg>
                  </button>
                  <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                    <li><button class="dropdown-item" @click="deletePhoto(photo)">Delete Photo</button></li>
                  </ul>
                </div>
                <img :src="requirePhoto(photo)" class="photo-img" />
                <div class="card-body">
                  <div class="d-flex justify-content-between align-items-center mb-2">
                    <div class="likes">
                      <button class="btn btn-link" :class="{ 'text-danger': userLikes[photo], 'text-dark': !userLikes[photo] }" @click="likePhoto(photo, this.user?.id)">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                      </button>
                      {{ this.likes[photo] }} likes
                    </div>
                    <div class="comments" v-if="comments[photo]">
                      <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#comment"/></svg>
                      {{ comments[photo].length }} comments
                    </div>
                  </div>
                  <div class="comment-section">
                    <div class="comment" v-for="comment in comments[photo]" :key="comment.id">
                      <router-link :to="{ name: 'MyAccount', params: { username: comment.username, userId: this.userId } }" style="text-decoration: none; color: black; font-weight: bold;">
                        {{ comment.username }}
                      </router-link> {{ comment.comment }}
                      <button v-if="parseInt(comment.idUser) === parseInt(this.userId)" class="btn btn-link text-danger m1-auto" @click="deleteComment(photo, comment.id)">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash"/></svg>
                      </button>
                    </div>
                    
                    <div class="mt-3">
                      <input type="text" v-model="newComment" class="form-control" placeholder="Add a comment..." @keyup.enter="addComment(photo)" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div> 
        </div>
        
        <ErrorMsg class="error-container" v-if="errormsg" :msg="errormsg" />
      </div>
    </template>
    
    <style>
    </style>
    
    

