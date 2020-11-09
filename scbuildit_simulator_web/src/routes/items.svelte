<svelte:head>
	<title>Item Catalog, lur!</title>
</svelte:head>

<script>
	import Item from "../components/Item.svelte"
	import rawMaterials from '../datas/raw-materials.js';
	import buildingSupplies from '../datas/building-supplies.js';
</script>
                
<div class="tm-col tm-col-big">
	<div class="bg-white tm-block h-100">
		<h2 class="tm-block-title">Raw Materials</h2>

		<ol class="tm-list-group">
			{#each rawMaterials as item}
			<li class="tm-list-group-item">
				<Item item="{item}" />
			</li>
			{/each}
		</ol>
	</div>
</div>
                
<div class="tm-col tm-col-big">
	<div class="bg-white tm-block h-100">
		<h2 class="tm-block-title">Raw Materials</h2>

		<ol class="tm-list-group">
			{#each buildingSupplies as item}
			<li class="tm-list-group-item">
				<Item item="{item}" />
				<ol class="tm-list-group">
					{#each item.materials as material}
					<li class="tm-list-group-item">
						<Item item="{material}" />
					</li>
					{/each}
				</ol>
			</li>
			{/each}
		</ol>
	</div>
</div>