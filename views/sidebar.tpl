		<div id="secondary" class="widget-area main-sidebar col-md-3" role="complementary">
		  <aside class="widget widget_search sidebar-widget clearfix">
			<h3 class="widget-title">搜索</h3>
			<form class="search" action="/search" method="get">
			  <fieldset>
				<div class="text">
				  <input name="keyword" id="keyword" type="text" placeholder="Search ...">
				  <button class="fa fa-search">Search</button>
				</div>
			  </fieldset>
			</form>
		  </aside>	<aside class="widget widget_categories sidebar-widget clearfix">
			<h3 class="widget-title">分类目录</h3>
			<ul>
			  <<<range .Category>>>
				<li class="cat-item"><a href="/category?c=<<<.>>>"><<< . >>></a></li>
				<<<end>>>
			</ul>
		  </aside>
		  <aside class="widget widget_archive sidebar-widget clearfix">
			<h3 class="widget-title">文章归档</h3>
			<ul>
			  <<<range .MonthBlog>>>
				<li><a href="/archive/<<<.>>>.html"><<<.>>></a></li>
				<<<end>>>
			</ul>
		  </aside>	<aside class="widget widget_recent_entries sidebar-widget clearfix">
			<h3 class="widget-title">近期文章</h3>
			<ul>
			  <<<range $i, $blog := .LastestBlogs>>>
			  <li><a href="/blog/<<<$blog.Name>>>.html"><<<$blog.Title>>></a></li>
			  <<<end>>>
    		</ul>
		  </aside>
		  <aside class="widget widget_tag_cloud sidebar-widget clearfix">
			<h3 class="widget-title">标签</h3>
			<div class="tagcloud">
			  <<<range .Tags>>>
				<a href="/tag?t=<<<.>>>" title="<<<.>>>"><<<.>>></a>
				<<<end>>>
			</div>
		  </aside>
		  <aside class="widget widget_text sidebar-widget clearfix">
			<h3 class="widget-title">介绍</h3>
			<div class="textwidget">
			  <p>生命不息，折腾不止</p>
			</div>
		  </aside>
		</div>
