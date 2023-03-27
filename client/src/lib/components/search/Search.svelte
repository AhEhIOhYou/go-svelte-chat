<script lang="ts">
	import { addToContacts, deleteContact, searchUserByUserName } from '@/lib/services/api-service';

	export let userID: string;
	console.log(userID);

	let searchTerm = '';
	let results = [];

	const handleSearch = async () => {
		const dataResponse = await searchUserByUserName(searchTerm);
		console.log(dataResponse);
		results = dataResponse.users ?? [];
	};

	const handleAddContact = async (e) => {
		console.log(userID, e.target.id);
		
		const dataResponse = await addToContacts(userID, e.target.id);
		console.log(dataResponse);
	};

	const handleDeleteContact = async (e) => {
		const dataResponse = await deleteContact(userID, e.target.id);
		console.log(dataResponse);
	};
</script>

<div class="search-container">
	<div class="search__inner">
		<input class="search__input" type="search" bind:value={searchTerm} />
		<button class="search__btn" on:click={handleSearch}>Search</button>
	</div>
	<div class="search__result">
		{#if results.length > 0}
			{#each results as result}
				<div class="chat-row">
					<div class="chat-row__username">
						{result.username}
						<button id={result.id} on:click={(e) => handleAddContact(e)}> Add to contacts </button>
						<button id={result.id} on:click={(e) => handleDeleteContact(e)}>
							Delete from contacts
						</button>
					</div>
				</div>
			{/each}
		{:else}
			<div>No results</div>
		{/if}
	</div>
</div>

<style lang="scss">
	.search-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
	}

	.search__inner {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.search__input {
		padding: 0.5rem 1rem;
		border: 1px solid #ccc;
		border-radius: 0.25rem;
	}

	.search__btn {
		padding: 0.5rem 1rem;
		margin-left: 0.5rem;
		border: 1px solid #ccc;
		border-radius: 0.25rem;
	}

	.search__result {
		margin-top: 1rem;
	}
</style>
