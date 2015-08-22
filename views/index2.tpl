<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<meta name="keywords" content="" />
	<meta name="description" content="" />
	<!-- <link href="< CUSTOM_URL_CSS%>default.css" rel="stylesheet" type="text/css" /> -->
	<!-- <link href="< CUSTOM_URL_CSS%>home.css" rel="stylesheet" type="text/css" /> -->
	<!-- <link rel="stylesheet" type="text/css" href="< CUSTOM_URL_CSS %>matrix.css" /> -->
	<!-- <link rel="shortcut icon" href="< CUSTOM_URL_IMG%>i_logo1.png" > -->
	<title><% .PageTitle %></title>
</head>
<body>
<div id="header">
	<div id="logo">
		<!-- <h1><img href="/" src="<STATIC_URL_IMG%>ckeyer.png" alt="ckeyer" /></h1> -->
		<h2><a href="/">Man, I just luv technology...... </a></h2>
		<!-- <h2><a href="/">O ever youthful, O ever weeping. </a></h2> -->
	</div>
	<div id="menu">
		<ul>
			<li class="first"><a href="/blog">博  客</a></li>
			<li><a href="/photo">相  册</a></li>
			<li><a href="/favorite">聊  天</a></li>
			<li><a href="/contact">留  言</a></li>
			<li><a href="/about">关  于</a></li>
			<li><a></a></li>
		</ul>
	</div>
</div>
<div id="content">
	<div class="colOne">
		<% range $index, $elem := .Blogs %>
		<div class="art_area" id="art_content_<% $index %>">
		  	<div class="article" id="article_<% $index %>">
		  		<div class="art_title" id="art_title_<% $index %>">
		  			<a href="/blog/<% $elem.ID %>"><% $elem.Title|DECODEBASE64 %></a>
		  		</div>
		  		<div class="art_time"><% $elem.CreatedTime %> &nbsp&nbsp
		  			<span> 
					<% range $ind, $ele := $elem.Tags %>
						<a href="/tag/<%$ele.ID%>" class="art_label"><%$ele.Name%></a>&nbsp
					<% end %>
		  			</span>	
	  			</div>
		  		<div class="art_summary">
		  			<% $elem.Summary|DECODEBASE64 %>
		  		</div>
		  		<div class="read_more">
		  			<a href="/blog/<% $elem.ID %>" title="">阅读全文</a>
		  		</div>
		  		<div class="art_content"></div>

		  			<!-- <hr class="art_separate"> -->
			</div>
		</div>
		<% end %>
	<div id="matrix_content"></div>
	</div>

	<div id="colTwo">
		<ul>
			<li>
				<h2>Archives</h2>
				<ul>
					<li><a href="#">December 2014</a></li>
					<li><a href="#">November 2014</a></li>
					<li><a href="#">October 2014</a></li>
					<li><a href="#">September 2014</a></li>
					<li><a href="#">August 2014</a></li>
				</ul>
			</li>
			<li>
				<h2>Categories</h2>
				<ul>
					<% range $index, $elem := .HotTags %>
						<li><a href="/tag/<%$elem.ID%>" ><%$elem.Name%></a> (<%$elem.ArtCount%>)</li>
					<%end%>
				</ul>
			</li>

			<li id="Friend_site_Link">
				<h2>Friend Site Link </h2>
				<ul>
					<!-- <li><a href="http://ys.cjstudio.org/" target="_blank">严申的个人博客</a></li>
					<li><a href="http://www.originate.com/">Originate</a></li>
					<li><a href="http://programmer.csdn.net/programmer.html">程序员杂志</a></li>
					<li><a href="http://www.litrin.net/">开源小站</a></li>
					<li><a href="http://lusongsong.com/">卢松松博客</a></li> -->
					<li><a href="https://gowalker.org/">gowalker.org</a></li>
				</ul>
			</li>
		</ul>
		<div style="clear: both;">&nbsp;</div>
	</div>
</div>
<div id="footer"><br>
	<p>Copyright &copy; 2014 lab204. Designed by <a href="#"><strong>Lab204-CJStudio</strong></a></p>
</div>
<script type="text/javascript" src="< STATIC_URL_JS %>jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="< CUSTOM_URL_JS %>matrix.js"></script>
</body>
</html>
