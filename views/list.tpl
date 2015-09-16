<div id="primary" class="content-area col-md-9">
  <div id="main" class="site-main" role="main">
	<<< range $index, $blog := .Blogs >>>
	  <article class="post hentry">
		<header class="entry-header">
		  <h1 class="post-title"><a href="/<<<$blog.Name>>>.html" rel="bookmark"><<< $blog.Title >>></a></h1>
		  <div class="entry-meta">
			<time class="post-date"><i class="fa fa-clock-o"></i><<< $blog.Date >>></time>
			<span class="seperator">/</span>
			<span><i class="fa fa-user"></i> <<< $blog.Author >>></span>
	      </div>
		</header>
		<div class="entry-content">
		  <p><<< $blog.Summary >>></p>
		</div>
		<footer class="entry-footer">
		  <ul class="post-categories">
			<<< range $i, $categ := $blog.Category >>>
			  <li><a href="/category?c=<<<$categ>>>" rel="category"><<<$categ>>></a></li>
			  <<<end>>>
          </ul>
		  
		  <ul class="post-tags">
			<<<range $i, $tag:= $blog.Tags >>>
			  <li><a href="/tag?t=<<<$tag>>>" rel="tag"><<<$tag>>></a></li>
			  <<<end>>>
          </ul>
		  
		  <div class="read-more">
			<a href="/<<<$blog.Name>>>.html">阅读全文<i class="fa fa-angle-double-right "></i></a>
		  </div>
		</footer>
	  </article>
	  <<<end>>>
  </div>
</div>


