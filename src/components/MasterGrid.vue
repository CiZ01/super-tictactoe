<script>

export default {
	name: 'MasterGrid',
	data() {
		return {
			masterGrid:{

			}
		}
	},
	methods: {
		// methods
	},


}

</script>


<template>
	<div id="container-mg" class="">
		<div class="row">
			<div class="col-md-12">
				<h1>Master Grid</h1>
			</div>
		</div>
		<div class="row">
			<div class="col-md-12">
				<router-view></router-view>
			</div>
		</div>
	</div>
</template>


<style>

</style>