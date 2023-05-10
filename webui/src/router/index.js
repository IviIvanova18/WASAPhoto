import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import StreamView from '../views/StreamView.vue'
import SearchUser from '../views/SearchUser.vue'
import MyProfile from '../views/MyProfile.vue'
import UploadPhoto from '../views/UploadPhoto.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: LoginView, name: 'Login' },
		{ path: '/stream', component: StreamView, name: 'Stream' },
		{ path: '/newimage', component: UploadPhoto, name: 'UploadPhoto' },
		{ path: '/search/user', component: SearchUser, name: 'SearchUser' },
		{ path: '/myprofile/profile', component: MyProfile, name: 'MyProfile'}
	]
})

export default router
