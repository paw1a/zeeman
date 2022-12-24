<script>
    import Controls from "./Controls.svelte";
    import View from "./View.svelte";
</script>

<main class="flex-shrink-0" role="main">
	<div class="container">
		<h1 class="mt-5">Эффект Зеемана</h1>
		<p class="lead">Эффект расщепления линий атомных спектров в присутствии статического магнитного поля</p>
		<div class="row">
			<View/>
			<Controls/>
		</div>
	</div>
</main>