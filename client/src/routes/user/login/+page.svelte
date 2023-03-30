<script lang="ts">
	import Login from '@/lib/components/auth/login/Login.svelte';
	import { loginHTTPRequest } from '@/lib/services/api-service';
	import { setItemInLS } from '@/lib/services/storage-service';

	async function handleLogin({ detail: { username, password } }) {
		const userDetails = await loginHTTPRequest(username, password);

		if (userDetails) {
			setItemInLS('userDetails', { 
				ID: userDetails.user.id,
				Name: userDetails.user.username,
			});
			window.location.href = '/home';
		} else {
			new Error('Something went wrong');
		}
	}
</script>

<Login on:login={handleLogin} />
