<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let chatlist = [];

	function handleClick(event) {
		const userId = event.target.id;
		dispatch('user-selected', userId);
	}
</script>

<div class="chatlist-container">
	<div class="chatlist">
		{#if chatlist.length === 0}
			<div class="chat__item">No chats</div>
		{:else}
			{#each chatlist as user}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div id={user.userId} class="chat__item" on:click={handleClick}>
					{user.username}
					{#if user.online === 'Y'}
						<div class="online" />
					{:else}
						<div class="offline" />
					{/if}
				</div>
			{/each}
		{/if}
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

		.chatlist {
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			height: 100%;
			width: 100%;

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
		}
	}
</style>
