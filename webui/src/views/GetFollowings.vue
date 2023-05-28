<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			followings: [],
			userId: null,
			username: null,
		};
	},
	methods: {
		load() {
			return load;
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				this.userId = this.$route.params.userId;
				this.username = this.$route.params.username;

				let apiUrl = `/users/${this.userId}/profile/${this.username}/`;
				let response = await this.$axios.get(apiUrl, {
					headers: {
						Authorization:
							"Bearer " + localStorage.getItem("token"),
					},
				});
				this.followings = response.data.followings;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async gotoAccount(id, username) {
			this.$router.push({
				name: "MyAccount",
				params: { userId: id, username: username },
			});
		},
	},

	mounted() {
		this.refresh();
	},
};
</script>

<template>
	<div>
		<h1>Followings list</h1>

		<ErrorMsg
			class="error-container"
			v-if="errormsg"
			:msg="errormsg"
		></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="followings?.length === 0">
			<div class="card-body">
				<p>No followings in the database.</p>
			</div>
		</div>

		<div v-if="!loading" v-for="f in followings">
			<a
				href="javascript:"
				class="text-muted mb-1 larger-text"
				style="text-decoration: none"
				@click="gotoAccount(this.userId, f)"
				>{{ f }}</a
			>
		</div>
	</div>
</template>
