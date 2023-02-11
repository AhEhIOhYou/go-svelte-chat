<script lang="ts">
	import Regisrtation from '@/lib/components/auth/regisrtation/Regisrtation.svelte';
	import { registerHTTPRequest } from '@/lib/services/api-service';
	import { setItemInLS } from '@/lib/services/storage-service';
	import type { PageData } from '../../regisrtation/$types';

	export let data: PageData;

	async function handleReg({ detail: { username, password } }) {
		const userDetails = await registerHTTPRequest(username, password);
		if (userDetails) {
			setItemInLS('userDetails', { userID: userDetails.userId });
			window.location.href = '/home';
		} else {
			new Error('Something went wrong');
		}
	}
</script>

<Regisrtation on:regisrtation={handleReg} />
