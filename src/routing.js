import Vue from "vue"
import VueRouter from "vue-router"

import Index from "./home/index.vue"
import Test from "./test/index.vue"
import Login from "./login/index.vue"
import registered from "./registered/index.vue"

Vue.use(VueRouter)

const router = new VueRouter({
	routes:[
		{
			path:"/",
			components:{default:Index},
			name:"Home"
		},
		{
			path:"/test",
			components:{default:Test},
			name:"home"
		},
		{
			path:"/login",
			components:{default:Login},
			name:"login"
		},
		{
			path:"/registered",
			components:{default:registered},
			name:"login"
		}
	]
})


 export default router





