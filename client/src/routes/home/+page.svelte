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

	let currentUserDetails;
	let userID = '';

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

{#if userID != ''}
	<Contacts userID={userID} />
{/if}

<button on:click={logout}> Logout </button>
