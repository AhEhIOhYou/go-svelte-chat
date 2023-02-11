<script lang="ts">
	export let chatlist = [];
</script>

<div class="chatlist-container">
	<div class="chatlist">
		{#each chatlist as user}
			<div id={user.userId} class="chat__item">
				{user.username}
				{#if user.online === 'Y'}
					<div class="online" />
				{:else}
					<div class="offline" />
				{/if}
			</div>
		{/each}
	</div>
</div>

<style lang="scss">
	.chatlist-container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 100%;
		width: 100%;
	}

	.chatlist {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 100%;
		width: 100%;
	}

	.chat__item {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		height: 50px;
		width: 100%;
		border-bottom: 1px solid #ccc;
	}

	.online {
		height: 10px;
		width: 10px;
		background-color: green;
		border-radius: 50%;
		margin-left: 10px;
	}

	.offline {
		height: 10px;
		width: 10px;
		background-color: red;
		border-radius: 50%;
		margin-left: 10px;
	}
</style>
