<html>
	<head>
		<style>
		 ul {
			 list-style: none;
		 }
		 .item {
			 margin-bottom: 10px;
		 }
		</style>
	</head>
	<body>
		<ul class="container">
			{{ range .Names }}
			<li class="item"><img src="http://storage.googleapis.com/hazzzit.appspot.com/{{ . }}" /></li>
			{{ end }}
		</ul>
	</body>
	<script src="/js/masonry.pkgd.min.js"></script>
	<script src="/js/imagesloaded.pkgd.min.js"></script>
	<script>
	 var container = document.querySelector(".container");
	 var imgLoaded = imagesLoaded( container );

	 imgLoaded.on( 'always', function (instance) {
		 var msnry = new Masonry( container, {
			 //columnWidth: 200,
			 itemSelector: ".item",
			 gutter: 10
		 });
	 });
	</script>
</html>
