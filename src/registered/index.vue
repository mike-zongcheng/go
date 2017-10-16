<style>
	div{color: red;}
</style>

<template>
	<div>
		<div><input type="text" name="" v-model="name"></div>
		<div><input type="password" name=""  v-model="password"></div>
		<button type="button" @click="login">注册</button>
	</div>
</template>

<script>
import axios from "axios"

export default {
	name:"login",
	data(){
		return{
			name:null,
			password:null
		}
	},
	methods:{
		login(){
			if(!this.name){
				alert("请填写用户名")
				return
			}if(!this.password){
				alert("请填写密码")
				return
			}
			axios.get("/api/registered/",{
				params:{
					name:this.name,
					password:this.password
				}}
			).then(function(result){
				window.location.href = "/#/test"
			})
		}
	},
	created(){
		
	}
}

</script>


