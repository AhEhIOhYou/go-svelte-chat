<script lang="ts">
	import { onMount } from 'svelte';
	import { getItemInLS, removeItemInLS } from '@/lib/services/storage-service';
	import { getUserByID } from '@/lib/services/api-service';
	import {
		closeWebSocketConnection,
		connectToWebSocket,
		listenToWebSocketMessages
	} from '@/lib/services/ws-service';
	import Search from '@/lib/components/search/Search.svelte';
	import Contacts from '@/lib/components/contacts-list/Contacts.svelte';
	import Conversation from '@/lib/components/conversation/Conversation.svelte';

	let user = null;
	let chatID = '';
	let newMessage = null;

	onMount(async () => {
		user = getItemInLS('userDetails');

		if (!user) {
			window.location.href = '/user/login';
		} else {
			const isUserLoggedInResponse = await getUserByID(user.ID);
			if (!isUserLoggedInResponse) {
				window.location.href = '/user/login';
			}
		}

		const wsConnection = connectToWebSocket(user.ID);
		console.log(wsConnection);

		if (wsConnection === null) {
			throw new Error(wsConnection.message);
		} else {
			listenToWebSocketMessages(function (data) {
				switch (data.name) {
					case 'connected':
						console.log(data);
						break;
					case 'disconnected':
						console.log(data);
						break;
					case 'message':
						newMessage = data.payload;
						break;
					default:
						console.log('default');
						console.log(data);
				}
			});
		}
	});

	const handleWriteMessage = (e) => {
		chatID = e.detail;
	};

	const logout = () => {
		removeItemInLS('userDetails');
		closeWebSocketConnection();
		window.location.href = '/';
	};
</script>

<div class="title">Home</div>
<button on:click={logout}> Logout </button>
<div>
	Search
	{#if user != null}
		<Search userID={user.ID} />
	{/if}
</div>

{#if user != null}
	<Contacts userID={user.ID} on:writeMessage={(e) => handleWriteMessage(e)} />
{/if}
{#if user != null && chatID != ''}
	{#key chatID}
		<Conversation {user} {chatID} {newMessage} />
	{/key}
{/if}
