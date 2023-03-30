<script lang="ts">
	import { getContacts, getDialog } from '@/lib/services/api-service';
	import { createEventDispatcher } from 'svelte';
	import { onMount } from 'svelte';

	const dispatch = createEventDispatcher();

	export let userID = null;
	let contacts = [];

	onMount(async () => {
		if (!userID) return;
		let contactsRes = await getContacts(userID);
		contacts = contactsRes.contacts;
	});

	async function handleWriteMessage(e) {
		let dialogRes = await getDialog([userID, e.target.id]);
		dispatch('writeMessage', dialogRes.chat.id);
	}
</script>

<div class="contacts">
	<div class="title">Contacts</div>
	<div class="list">
		{#if !contacts}
			<div class="contact">
				<div class="name">No contacts</div>
			</div>
		{:else}
			{#each contacts as contact}
				<div class="contact">
					<div class="name">{contact.contactUsername}</div>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<div id={contact.contactUserID} class="new-message" on:click={(e) => handleWriteMessage(e)}>message</div>
				</div>
			{/each}
		{/if}
	</div>
</div>

<style lang="scss">
	.contacts {
		display: flex;
		flex-direction: column;
		height: 100%;
		width: 100%;

		.title {
			font-size: 1.5rem;
			font-weight: 600;
			padding: 1rem;
		}

		.list {
			flex: 1;
			overflow-y: auto;

			.contact {
				display: flex;
				flex-direction: row;
				justify-content: space-between;
				align-items: center;
				padding: 1rem;
				border-bottom: 1px solid #ccc;

				.name {
					font-size: 1.2rem;
					font-weight: 600;
				}

				.new-message {
					cursor: pointer;
					font-size: 1.2rem;
					font-weight: 600;
					color: #007bff;
				}
			}
		}
	}
</style>
