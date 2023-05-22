<script>
  import './styles.css'; 

  export default {
    data: function() {
      return {
        errormsg: null,
        loading: false,
        photo: null,
        
      }
    },
    methods: {
      load() {
        return load
      },
      // async refresh() {
      //   this.loading = true;
      //   this.errormsg = null;
      //   await this.$axios.refresh();
      // },
      uploadPhoto: async function (){
        this.loading = true;
        this.errormsg = null;
        try{
          let userId = this.$route.params.userId;
          await this.$axios.post(`/users/${userId}/photos/`,{
            path : this.photo,
          });
          this.$router.push({ name: 'Stream', params: { userId: userId } });
          
          
        }catch(e){
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
    <div class="container py-5">
      <div class="row justify-content-center align-items-center">
        <div class="col-md-6">
          <div class="card bg-white text-dark rounded-3">
            <div class="card-body p-5 text-center">
              <h2 class="fw-bold mb-4 text-uppercase">Upload Photo</h2>
              <p class="text-muted">Please select your photo.</p>
              <div class="form-group">
                <input
                  type="file"
                  id="photo"
                  class="form-control form-control-lg rounded-pill mb-3"
                  accept="image/*"
                  required
                />
              </div>
              <div class="d-grid gap-3">
                <button class="btn btn-primary rounded-pill" type="submit" @click="uploadPhoto" style="background-color: #2e4a78;">
                  Upload
                </button>
                <LoadingSpinner v-if="loading" />
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
      

    