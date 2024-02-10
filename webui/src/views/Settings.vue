<script>
export default {
	data: function () {
		return {
			errormsg: null,
			username: "",
		}
	},

	methods:{
		async modifyUsername(){
			try{
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
					username: this.username,
				})

				this.username=""
			}catch (e){
				this.errormsg = e.toString();
			}
		},
	},

}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col d-flex justify-content-center mb-2">
				<h1>{{ this.$route.params.id }}'s Settings</h1>
			</div>
		</div>

		<div class="row ">
			
			<div class="col-12 d-flex justify-content-center">
				<p>Enter a new Username</p>
			</div>
		</div>

		<div class="row mt-2">
			<div class="col d-flex justify-content-center">
				<div class="input-group mb-3 w-25">
					<input
						type="text"
						class="form-control w-25"
						placeholder="Your new username..."
						maxlength="16"
						minlength="3"
						v-model="username"
					/>
					<div class="input-group-append">
						<button class="btn btn-outline-secondary" 
						@click="modifyUsername"
						:disabled="username === null || username.length >16 || username.length <3 || username.trim().length===0">
						Modify</button>
					</div>
				</div>
			</div>
		</div>

		<div class="row" >
			<div v-if="username.trim().length>0" class="col d-flex justify-content-center">
				Preview: {{username}} @{{ this.$route.params.id }}
			</div>
		</div>

		<div class="row">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</div>
	
</template>

<style>
</style>
