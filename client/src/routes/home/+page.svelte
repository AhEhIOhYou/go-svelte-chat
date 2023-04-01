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

<div class="title">Home</div>
{#await websocketConnection}
	<div>Laoding...</div>
{:then websocketConnection}
	<BaseChat {user} {websocketConnection} />
{/await}
