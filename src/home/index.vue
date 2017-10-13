<style>
	div{color: red;}
</style>

<template>
	<div>
		{{text}}
	</div>
</template>

<script>
import axios from "axios"

export default {
	name:"home",
	data(){
		return{
			text:"Hello go And Vue!"
		}
	}
}

</script>


