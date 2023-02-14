<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import Chatlist from '@/lib/components/chatlist/Chatlist.svelte';
	import {
		userSessionCheckHTTPRequest
	} from '@/lib/services/api-service';
	import {
		closeWebSocketConnection,
		connectToWebSocket,
		listenToWebSocketMessages
	} from '@/lib/services/ws-service';
	import Conversation from '@/lib/components/conversation/Conversation.svelte';
	import Search from '@/lib/components/search/Search.svelte';

	let currentUserDetails;
	let buddyUserID;
	let chatlist = [];
	let newMessage = null;

	onMount(async () => {
		currentUserDetails = getItemInLS('userDetails');

		if (!currentUserDetails) {
			window.location.href = '/user/login';
		} else {
			const isUserLoggedInResponse = await userSessionCheckHTTPRequest(currentUserDetails.userID);
			if (!isUserLoggedInResponse) {
				window.location.href = '/user/login';
			}
		}

		const wsConnection = connectToWebSocket(currentUserDetails.userID);

		if (wsConnection === null) {
			throw new Error(wsConnection.message);
		} else {
			listenToWebSocketMessages(function (data) {
				switch (data.type) {
					case 'my-chatlist':
						chatlist = data.chatlist ?? [];
						break;
					case 'user-disconnected':
						chatlist.forEach((user, index) => {
							if (user.userId === data.chatlist.userId) {
								chatlist[index].online = 'N';
							}
						});
						break;
					case 'user-connected':
						chatlist.forEach((user, index) => {
							if (user.userId === data.chatlist.userId) {
								chatlist[index].online = 'Y';
							}
						});
						break;
					case 'new-message':
						newMessage = data;
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
	<Search />
</div>
<div>
	ChatList
	{#key chatlist}
		<Chatlist {chatlist} on:user-selected={(event) => (buddyUserID = event.detail)} />
	{/key}
</div>
<div>
	{#if buddyUserID}
		{#key buddyUserID}
			<Conversation currentUserID={currentUserDetails.userID} {buddyUserID} {newMessage} />
		{/key}
	{/if}
</div>
<button on:click={logout}> Logout </button>
