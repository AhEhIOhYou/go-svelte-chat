<script lang="ts">
	import type { PageData } from './$types';
	import { onDestroy, onMount } from 'svelte';

	import { getStores, navigating, page, updated } from '$app/stores';
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import Chatlist from '@/lib/components/chat-list/Chatlist.svelte';
	import { userSessionCheckHTTPRequest } from '@/lib/services/api-service';
	import {
		closeWebSocketConnection,
		connectToWebSocket,
		listenToWebSocketMessages,
		sendWebSocketMessage
	} from '@/lib/services/ws-service';

	let userDetails;

	onMount(async () => {
		userDetails = JSON.parse(localStorage.getItem('userDetails'));
		
		if (!userDetails) {
			window.location.href = '/user/login';
		} else {
			const isUserLoggedInResponse = await userSessionCheckHTTPRequest(userDetails.userID);
			if (!isUserLoggedInResponse) {
				window.location.href = '/user/login';
			}
		}

		const wsConnection = connectToWebSocket(userDetails.userID);
		if (wsConnection === null) {
			throw new Error(wsConnection.message);
		} else {
			listenToWebSocketMessages();
		}
	});

	onDestroy(() => {
		closeWebSocketConnection();
	});

	export let data: PageData;

	const logout = () => {
		// sendWebSocketMessage('mama');
		removeItemInLS('userDetails');
		closeWebSocketConnection();
		// window.location.href = '/';
	};
</script>

<div class="title">Home</div>
<div>
	ChatList
	<Chatlist />
</div>
<div>Conversation</div>
<button on:click={logout}> Logout </button>
