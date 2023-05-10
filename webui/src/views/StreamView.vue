<script>
	export default {
		data: function() {
			return {
				errormsg: null,
				loading: false,
				photoes: null,
				wantstocomment: false
			}
		},
		emits: ["refresh"],
		methods: {
			async refresh() {
				this.loading = true;
				this.errormsg = null;
				try {
					let url = "/Users/" + this.$cookies.get("token") + "/Stream"
					let config = {
						headers: {
							Authorization: "Bearer " + this.$cookies.get("token")
						}
					}
					let response = await this.$axios.get(url, config)
					this.photoes = response.data;
				} catch (e) {
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
		<div>
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">Stream</h1>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2">
						<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
							Refresh
						</button>
					</div>
				</div>
			</div>
	
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<section>
		
	<section class="row d-flex justify-content-center align-items-center">
		<div class="card bg-light" style="max-width: 35rem; padding-top: 5px;" >
		<div class="card" style="max-width: 35rem;" v-if="!loading" v-for="ph in photoes">
		<div class="card-body">
			<div class="d-flex mb-3">
			<div>
				<a href="" class="text-dark mb-0">
				<strong>{{ ph.username }}</strong>
				</a>
			</div>
			</div>
		</div>
		<div class="bg-image hover-overlay ripple rounded-0" data-mdb-ripple-color="light">
			<ImageWithAuthorization :id="ph.id"></ImageWithAuthorization>
			<a href="#!">
			<div class="mask" style="background-color: rgba(251, 251, 251, 0.2)"></div>
			</a>
		</div>
		<div class="card-body">
			<div class="d-flex justify-content-between mb-3">
			<div>
				<a href="">
					<span>{{ ph.no_likes }}</span>
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#thumbs-up"/></svg>
				</a>
			</div>
			<div>
				<a href="" class="text-muted"> {{ ph.no_comments }} </a>
			</div>
			</div>
			<div class="d-flex justify-content-between text-center border-top border-bottom mb-4">
			<LikeButton :id="ph.id" :is_liked="ph.is_liked" @doRefresh="refresh"></LikeButton>
			<RouterLink :to="{ name: 'CommentPage', params: { id: ph.id, username: ph.username }}"><button type="submit" class="btn btn-light btn-lg">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-square"/></svg> Comment
			</button></RouterLink>

		</div>
			
		</div>
		</div>
	</div>
	</section>
</section>	
		</div>
	</template>
	
	<style>
	</style>
	