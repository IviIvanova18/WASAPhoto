import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import MyAccount from '../views/MyAccount.vue'
import Followings from '../views/GetFollowings.vue'
import UploadPhoto from '../views/UploadPhotoView.vue'
import SearchUser from '../views/SearchUser.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: Login, name: 'Login' },
		{ path: '/users/:userId/stream', component: StreamView, name: 'Stream' },
		{ path: '/users/:userId/photos', component: UploadPhoto, name: 'UploadPhoto' },
		{ path: '/users/:userId/profile/:username', component: MyAccount, name: 'MyAccount'},
		{ path: '/users/:userId/followings', component: Followings, name: 'Followings'},
		{ path: '/users/:userId/followers', component: Followings, name: 'Followings'},
		{ path: '/:userId/search', component: SearchUser, name: 'SearchUser'}
		
	]
})

export default router
