import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/LoginView.vue'
// import StreamView from '../views/StreamView.vue'
import MyAccount from '../views/MyAccount.vue'
import Followings from '../views/GetFollowings.vue'
// import UploadPhoto from '../views/UploadPhoto.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: Login, name: 'Login' },
		// { path: '/stream', component: StreamView, name: 'Stream' },
		// { path: '/newphoto', component: UploadPhoto, name: 'UploadPhoto' },
		{ path: '/users/:userId/profile/:username', component: MyAccount, name: 'MyAccount'},
		{ path: '/users/:userId/following', component: Followings, name: 'Followings'}

		
	]
})

export default router
