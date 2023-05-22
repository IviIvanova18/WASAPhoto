  <script>
    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          followings: [],
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
            this.followings = response.data.followings; 
            console.log(this.followings);     
                        
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
	<div >
			<h1>Followings list</h1>

		<ErrorMsg class="error-container" v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="followings?.length === 0">
			<div class="card-body">
				<p>No followings in the database.</p>
				
			</div>
		</div>

		<div v-if="!loading" v-for="f in followings">
      {{f}}
		</div>
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