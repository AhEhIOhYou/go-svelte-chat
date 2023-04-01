<script lang="ts">
	import { getChatsByUser } from '@/lib/services/api-service';
	import { createEventDispatcher } from 'svelte';
	import { onMount } from 'svelte';

	const dispatch = createEventDispatcher();

	export let user = null;

	let chatslist = [];

	onMount(async () => {
		let chatsRes = await getChatsByUser(user.ID);
		chatslist = chatsRes.chats;
	});

	function handleClick(e) {
		dispatch('chat-selected', {
			ID: e.target.id
		});
	}
</script>

<div class="chatlist-container">
	<div class="chatlist">
		{#if !chatslist}
			<div>No Chats</div>
		{:else}
			{#each chatslist as chat}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div class="chat__item" id={chat.id} on:click={(e) => handleClick(e)}>
					{chat.name}
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
		}
	}
</style>
