<style>
	div{color: red;}
</style>

<template>
	<div class="main">
		<ul class="article_list">
			<li>
				<h6 class="article_listTitle">标题</h6>
				<div class="article_introduction">简介</div>
				<div class="article_listInfo">
					<ul class="article_listLabel">
						<li>javascript</li>
						<li>react</li>
					</ul>
					<span class="article_listTime">时间</span>
				</div>
			</li>
		</ul>
		<router-link to="./login">登录</router-link>
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
			axios.get("/api/test/",{
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


