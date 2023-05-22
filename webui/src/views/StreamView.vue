  <script>
    import './styles.css'; 

    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          userId: this.$route.params.userId,
          photos:[],
          comments:{},
          userLikes:{},
          newComment: null,
        }
      },
      methods: {
        load() {
          return load
        },
        requirePhoto(path) {
          return `../src/assets/images/${path}`;
        },
        async refresh() {
          this.loading = true;
          this.errormsg = null;
          try {
            let userId = this.$route.params.userId;
            let apiUrl = `/users/${userId}/stream/`;                         
            let response = await this.$axios.get(apiUrl);
            this.photos = response.data;  
            for (const photo of this.photos) {
              
              const commentsResponse = await this.$axios.get(`/photos/${photo.id}/comments/`);         
              this.comments[photo.id] = commentsResponse.data;
              const likesResponse = await this.$axios.get(`/photos/${photo.id}/likes/`);         
              this.userLikes[photo.id] = likesResponse.data.map(p => p.idUser).includes(photo.id_user)
            
            }

          } catch (e) {
            this.errormsg = e.toString();
          }
          this.loading = false;
        },
        async likePhoto(photoId, userId){
            this.loading = true;
			      this.errormsg = null
			      try {
              if (this.userLikes[photoId]) {
                await this.$axios.delete(`/photos/${photoId}/likes/${this.userId}/`) 
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
		    }
        
      },
      
      mounted() {
        this.refresh()
      }
    }
    </script>

<template>
  <div class="background">
    <div class="container py-5">
      <div class="row justify-content-center align-items-start">
        <div class="col-md-8">
          <div class="photo-card bg-white text-dark rounded-5 mb-4" v-for="photo in photos" :key="photo.id">
            <img :src="requirePhoto(photo.path)" class="photo-img" />
            <div class="card-body">
              <div class="d-flex justify-content-between align-items-center mb-2">
                <div class="likes">
                  <button class="btn btn-link" :class="{ 'text-danger': userLikes[photo.id], 'text-dark': !userLikes[photo.id] }" @click="likePhoto(photo.id, photo.id_user)">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                  </button>
                  {{ photo.likes }} likes
                </div>
                <div class="comments" v-if="comments[photo.id]">
                  {{ comments[photo.id].length }} comments
                </div>
              </div>
              <div class="comment-section">
                <div class="comment" v-for="comment in comments[photo.id]" :key="comment.id">
                  <strong>{{ comment.username }}</strong> {{ comment.comment }}
                  <button class="btn btn-link text-danger m1-auto" @click="deleteComment(photo.id, comment.id)">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash"/></svg>
                  </button>
                </div>
                
                <div class="mt-3">
                  <input type="text" v-model="newComment" class="form-control" placeholder="Add a comment..." @keyup.enter="addComment(photo.id)" />
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
    
    

