<script lang="ts">
	import '../styles/styles.scss';
</script>

<main>
	<div class="size-ruler">
		<slot />
	</div>
</main>

<style lang="scss">
	.size-ruler {
		position: relative;
		margin: 50px auto 60px;
		border-radius: 2px;
		min-height: calc(100vh - 228px);
		max-width: 1200px;
	}
</style>