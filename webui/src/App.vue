<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
        return {
			loggedin: false,
			id: null,
            username: null,
        }
    },
	methods: {
      load() {
        return load
      },
      async refresh() {
        this.loggedin = true;
        this.errormsg = null;
		if (this.$route.params.hasOwnProperty("userId") || this.$route.params.hasOwnProperty("username")) {
			// console.log(typeof this.$route.params.userId);
			this.id = parseInt(this.$route.params.userId);
        	this.username = this.$route.params.username;
			// this.loggedin = true;
			console.log("The route contains both userId and username parameters.",this.$route.params.userId);
			} else {
			  this.loggedin = false;
			  console.log("The route does not contain both userId and username parameters.");
			}
			// this.loggedin = true;
      },
          
    },
    
    mounted() {
      this.refresh()
	//   this.loggedin = false;
    }
}
</script>

<template>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
	  <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASA Photo</a>
	  <button
		class="navbar-toggler position-absolute d-md-none collapsed"
		type="button"
		data-bs-toggle="collapse"
		data-bs-target="#sidebarMenu"
		aria-controls="sidebarMenu"
		aria-expanded="false"
		aria-label="Toggle navigation"
	  >
		<span class="navbar-toggler-icon"></span>
	  </button>
	</header>
  
	<div class="container-fluid">
	  <div class="row">
		<nav
		  id="sidebarMenu"
		  class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse"
		>
		  <div class="position-sticky pt-3 sidebar-sticky">
			<div v-if="this.id!=null">
			  <h6
				class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase"
			  >
				<span>General</span>
			  </h6>
			  <ul class="nav flex-column">
				<li class="nav-item">
				  <RouterLink
					v-if="this.id!=null"
					:to="{ name: 'Stream', params: { userId: id } }"
					class="nav-link"
				  >
					<svg class="feather">
					  <use href="/feather-sprite-v4.29.0.svg#home" />
					</svg>
					Home
				  </RouterLink>
				</li>
				<li class="nav-item">
				  <RouterLink
					v-if="this.id!=null"
					:to="{ name: 'UploadPhoto', params: { userId: id } }"
					class="nav-link"
				  >
					<svg class="feather">
					  <use href="/feather-sprite-v4.29.0.svg#image" />
					</svg>
					Upload Photo
				  </RouterLink>
				</li>
				<li class="nav-item">
				  <RouterLink
					v-if="this.id!=null"
					:to="{ name: 'SearchUser' , params: { userId: id } }"
					class="nav-link"
				  >
					<svg class="feather">
					  <use href="/feather-sprite-v4.29.0.svg#search" />
					</svg>
					Search User
				  </RouterLink>
				</li>
			  </ul>
  
			  <h6
				class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase"
			  >
				<span>My Account</span>
			  </h6>
			  <ul class="nav flex-column">
				<li class="nav-item">
				  <RouterLink
					v-if="this.id!=null"
					:to="{ name: 'MyAccount', params: { userId: id, username: username } }"
					class="nav-link"
				  >
					<svg class="feather">
					  <use href="/feather-sprite-v4.29.0.svg#user" />
					</svg>
					My Account
				  </RouterLink>
				</li>
			  </ul>
			</div>
  
			
			<li class="nav-item">
			<RouterLink to="/" class="nav-link">
			  <svg class="feather">
				<use href="/feather-sprite-v4.29.0.svg#log-out" />
			  </svg>
			  Logout
			</RouterLink>
			</li>
		  </div>
		</nav>
  
		<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
		  <RouterView />
		</main>
	  </div>
	</div>
  </template>
  

<style>
</style>
