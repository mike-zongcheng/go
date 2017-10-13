import Vue from "vue"
import VueRouter from "vue-router"
import Index from "./home/index.vue"

Vue.use(VueRouter)

const router = new VueRouter({
	routes:[
		{
			path:"/",
			components:{default:Index},
			name:"home"
		}
	]
})


 export default router





