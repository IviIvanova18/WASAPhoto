<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
        return {
			userId: this.$route.params.userId,
            username: this.$route.params.username,
        }
      },
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASA Photo</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<!-- <li class="nav-item">
							<RouterLink :to="{ name: 'Stream' }" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li> -->
						<!-- <li class="nav-item">
							<RouterLink :to="{ name: 'UploadPhoto', params: { userId: this.id } }" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
								Upload Photo
							  </RouterLink>
          					<router-link :to="{ name: 'UploadPhoto', params: { userId: this.userId }}" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
								Upload Photo</router-link>

						</li> -->
						<!-- <li class="nav-item">
							<RouterLink to="{ name : SearchUser}" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								Search User
							</RouterLink>
						</li> -->
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>My Account</span>
					</h6>
					<ul class="nav flex-column">
						<!-- <li class="nav-item">'
							<RouterLink :to="{ name: 'MyAccount', params: { userId: userId, username: username } }" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
								MyAccout
							</RouterLink>

						</li> -->
						<li class="nav-item">
							<RouterLink to="/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
								Logout 
							</RouterLink>
						</li>
					</ul>
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
