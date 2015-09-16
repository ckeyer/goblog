<div id="primary" class="content-area single col-md-9">
  <div id="main" class="site-main" role="main">
	<article class="post hentry">
	  <header class="entry-header">
		<h1 class="post-title"><a href="/<<<.Blog.Name>>>.html" rel="bookmark"><<<.Blog.Title>>></a></h1>
		<div class="entry-meta">
		  <time class="post-date"><i class="fa fa-clock-o"></i><<<.Blog.Date>>></time>
		  <span class="seperator">/</span>
	      <span><i class="fa fa-user"></i> <<<.Blog.Author>>></span>
	    </div><!-- .entry-meta -->
	  </header><!-- .entry-header -->
	  <div class="entry-content">
		<<<str2html .BContent>>>
	  </div><!-- .entry-content -->
	  <footer class="entry-footer">
		<ul class="post-categories">
          <<<range .Blog.Category>>>
			<li><a href="/category?c=<<<.>>>" rel="category"><<<.>>></a></li>
			<<<end>>>
        </ul>
		
		<ul class="post-tags">
		  <<<range .Blog.Tags>>>
			<li><a href="/tag?t=<<<.>>>" rel="tag"><<<.>>></a></li>
			<<<end>>>
        </ul>
		
	  </footer><!-- .entry-footer -->
	</article><!-- #post-## -->
	<<<template "Duoshuo">>>
  </div>
  <!-- #main -->
</div>
