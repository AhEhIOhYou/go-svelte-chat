<script lang="ts">
	import { getItemInLS } from '@/lib/services/storage-service';
	import { connectToWebSocket } from '@/lib/services/ws-service';
	import BaseChat from '@/lib/components/base-chat/BaseChat.svelte';

	let user = null;

	async function getWsCon() {
		user = getItemInLS('userDetails');

		if (!user) {
			window.location.href = '/user/login';
		}

		return connectToWebSocket(user.ID);
	}

	let websocketConnection = getWsCon();
</script>

{#await websocketConnection}
	<div>Loading...</div>
{:then websocketConnection}
	<div class="title">Hello, {user.Name}</div>
	<BaseChat {user} {websocketConnection} />
{/await}
