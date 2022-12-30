import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import UserView from '../views/UserView.vue'
import ErorrView from '../views/404view.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/session', component: HomeView},
		{path: '/users/:username/profile', component: ProfileView},
		{path: '/users/:username/view', component: UserView},
		{path: '/:catchAll(.*)', component: ErorrView}
	]
})

export default router
