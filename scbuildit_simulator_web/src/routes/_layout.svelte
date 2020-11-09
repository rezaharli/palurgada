<script>
	import Nav from '../components/Nav.svelte';

	export let segment;
</script>

<svelte:head>
	<!--

    Template 2108 Dashboard

	http://www.tooplate.com/view/2108-dashboard

    -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,600" />
    <!-- https://fonts.google.com/specimen/Open+Sans -->
    <link rel="stylesheet" href="css/fontawesome.min.css" />
    <!-- https://fontawesome.com/ -->
    <link rel="stylesheet" href="css/fullcalendar.min.css" />
    <!-- https://fullcalendar.io/ -->
    <link rel="stylesheet" href="css/bootstrap.min.css" />
    <!-- https://getbootstrap.com/ -->
	<link rel="stylesheet" href="css/tooplate.css" />
</svelte:head>

<div class="row">
	<div class="col-12">
		<Nav {segment} />
	</div>
</div>

<main>
	<div class="row tm-content-row tm-mt-big">
		<slot></slot>
	</div>

	<footer class="row tm-mt-small">
		<div class="col-12 font-weight-light">
			<p class="d-inline-block tm-bg-black text-white py-2 px-4">
				Copyright &copy; 2018 Admin Dashboard . Created by
				<a rel="nofollow" href="https://www.tooplate.com" class="text-white tm-footer-link">Tooplate</a>
			</p>
		</div>
	</footer>
</main>