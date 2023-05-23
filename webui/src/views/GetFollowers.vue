  <script>
    export default {
      data: function() {
        return {
          errormsg: null,
          loading: false,
          followers: [],
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
            this.followers = response.data.followers; 
            console.log(this.followers);     
                        
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

		<div class="card" v-if="followers?.length === 0">
			<div class="card-body">
				<p>No followers in the database.</p>
				
			</div>
		</div>

		<div v-if="!loading" v-for="f in followers">
      {{f}}
		</div>
	</div>
</template>

