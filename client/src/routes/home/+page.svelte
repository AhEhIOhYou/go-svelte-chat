<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import Chatlist from '@/lib/components/chatlist/Chatlist.svelte';
	import { getUserByID } from '@/lib/services/api-service';
	import {
		closeWebSocketConnection,
		connectToWebSocket,
		listenToWebSocketMessages
	} from '@/lib/services/ws-service';
	import Conversation from '@/lib/components/conversation/Conversation.svelte';
	import Search from '@/lib/components/search/Search.svelte';

	let currentUserDetails;
	let userID = '';
	let buddyUserID;
	let chatlist = [];
	let newMessage = null;

	onMount(async () => {
		currentUserDetails = getItemInLS('userDetails');
		userID = currentUserDetails.userID;

		if (!currentUserDetails) {
			window.location.href = '/user/login';
		} else {
			const isUserLoggedInResponse = await getUserByID(currentUserDetails.userID);
			if (!isUserLoggedInResponse) {
				window.location.href = '/user/login';
			}
		}

		const wsConnection = connectToWebSocket(currentUserDetails.userID);
		console.log(wsConnection);

		if (wsConnection === null) {
			throw new Error(wsConnection.message);
		} else {
			listenToWebSocketMessages(function (data) {
				switch (data.type) {
					case 'connected':
						console.log(data);
						break;
					case 'disconnected':
						console.log(data);
						break;
					case 'message':
						console.log(data);
						break;
					default:
						console.log(data);
				}
			});
		}
	});

	onDestroy(() => {
		closeWebSocketConnection();
	});

	const logout = () => {
		removeItemInLS('userDetails');
		closeWebSocketConnection();
		window.location.href = '/';
	};
</script>

<div class="title">Home</div>
<div>
	Search
	{#if userID != ''}
			<Search userID={userID} />
	{/if}
</div>
<div>
	ChatList
	<!-- {#key chatlist}
		<Chatlist {chatlist} on:user-selected={(event) => (buddyUserID = event.detail)} />
	{/key} -->
</div>
<div>
	<!-- {#if buddyUserID}
		{#key buddyUserID}
			<Conversation currentUserID={currentUserDetails.userID} {buddyUserID} {newMessage} />
		{/key} -->
	<!-- {/if} -->
</div>
<button on:click={logout}> Logout </button>
