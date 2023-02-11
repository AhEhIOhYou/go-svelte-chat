<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import Chatlist from '@/lib/components/chatlist/Chatlist.svelte';
	import { userSessionCheckHTTPRequest } from '@/lib/services/api-service';
	import {
		closeWebSocketConnection,
		connectToWebSocket,
		listenToWebSocketMessages
	} from '@/lib/services/ws-service';

	let userDetails;
	let chatlist = [];

	onMount(async () => {
		userDetails = getItemInLS('userDetails');
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
	ChatList
	{#key chatlist}
		<Chatlist {chatlist} />
	{/key}
</div>
<div>Conversation</div>
<button on:click={logout}> Logout </button>
