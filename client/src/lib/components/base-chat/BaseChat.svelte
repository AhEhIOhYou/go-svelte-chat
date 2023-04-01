<script lang="ts">
	import { onMount } from 'svelte';
	import { removeItemInLS } from '@/lib/services/storage-service';
	import { closeWebSocketConnection, listenToWebSocketMessages } from '@/lib/services/ws-service';
	import Search from '@/lib/components/search/Search.svelte';
	import Contacts from '@/lib/components/base-chat/contacts-list/Contacts.svelte';
	import Conversation from '@/lib/components/base-chat/conversation/Conversation.svelte';
	import Chatlist from '@/lib/components/base-chat/chatlist/Chatlist.svelte';

	let chatID = '';
	let newMessage = null;
	let root;

	export let user = null;
	export let websocketConnection = null;

	onMount(async () => {
		if (user === null || websocketConnection === null) {
			// window.location.href = '/user/login';
			console.log(websocketConnection);

			throw new Error(websocketConnection.message);
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

		const accordions = root.querySelectorAll('.accrodion__body');
		//close 0
		accordions[0].classList.remove('show');
		accordions[0].classList.add('collapsing');
		//open 1
		accordions[1].classList.add('show');
		accordions[1].classList.remove('collapsing');
	};

	const handleChatSelected = (e) => {
		chatID = e.detail.ID;
	};

	const logout = () => {
		removeItemInLS('userDetails');
		closeWebSocketConnection();
		window.location.href = '/';
	};

	const toggleAccordion = (e) => {
		const accordions = root.querySelectorAll('.accrodion__body');

		const currentAccordion = e.target.nextElementSibling;

		accordions.forEach((accordion) => {
			if (accordion !== currentAccordion) {
				accordion.classList.remove('show');
				accordion.classList.add('collapsing');
			}
		});

		if (currentAccordion.classList.contains('show')) {
			currentAccordion.classList.remove('show');
			currentAccordion.classList.add('collapsing');
		} else {
			currentAccordion.classList.add('show');
			currentAccordion.classList.remove('collapsing');
		}
	};
</script>

<div class="base-container">
	<button on:click={logout}> Logout </button>
	<div>
		Search
		{#if user != null}
			<Search userID={user.ID} />
		{/if}
	</div>
	<div class="accordion" bind:this={root}>
		<div class="accrodion__item">
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<div class="accrodion__title" on:click={(e) => toggleAccordion(e)}>Contacts</div>
			<div class="accrodion__body collapsing">
				<div class="inner">
					{#if user != null}
						<Contacts userID={user.ID} on:writeMessage={(e) => handleWriteMessage(e)} />
					{/if}
				</div>
			</div>
		</div>
		<div class="accrodion__item">
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<div class="accrodion__title" on:click={(e) => toggleAccordion(e)}>Chats</div>
			<div class="accrodion__body collapsing">
				<div class="inner">
					{#if user != null}
						<Chatlist {user} on:chat-selected={(e) => handleChatSelected(e)} />
					{/if}
					{#if user != null && chatID != ''}
						{#key chatID}
							<Conversation {user} {chatID} {newMessage} />
						{/key}
					{/if}
				</div>
			</div>
		</div>
	</div>
</div>

<style lang="scss">
	.base-container {
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	.accrodion__title {
		font-size: 20px;
		padding: 10px 15px;
		background: gray;
		cursor: pointer;
	}

	.inner {
		border: 1px solid black;
	}

	.accrodion__body.show {
		display: block;
	}

	.accrodion__body.collapsing {
		display: none;
	}
</style>
