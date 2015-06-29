<!DOCTYPE html>
<html>
<head> 
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" /> 
	<link rel="Stylesheet" type="text/css" href="<% html2str .STATIC_URL_CSS%>jquery-ui.1.11.3.min.css" />
	<link rel="Stylesheet" type="text/css" href="<% html2str .STATIC_URL_CSS%>jHtmlArea.css" />
</head>
<body>
	<div id="blog_edit">
		<input type="text" id="blog_title" value="" placeholder="标题">
		<select id="blog_type"value="Node" >
			<option value="node">Node</option>
			<option value="blog">Blog</option>
			<option value="favor">Favor</option>
		</select>
		<br>
		<span id="blog_tag_list"></span>
		<input type="text" id="tag_add_text" value="" placeholder="标签">
		<input type="button" id="add_tag_button" onclick="addTag();" value="+Tag"></input>
		<br>
		<textarea id="blog_summary" cols="50" rows="3"  placeholder="摘要"></textarea>
		<textarea id="txtDefaultHtmlArea" cols="50" rows="15"  ></textarea>
		<input type="password" id="password" value="" placeholder="密码">
		<input type="button" id="commit_button" onclick="commit();" value="123543"></input>
<!-- 
		<input type="button" value="Alert HTML" onclick="alert($('#txtDefaultHtmlArea').htmlarea('html'));" />
		<input type="button" value="Change Color to Blue" onclick="$('#txtDefaultHtmlArea').htmlarea('forecolor', 'blue');" /> -->
	</div>
<script type="text/javascript" src="<% html2str .STATIC_URL_JS%>jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="<% html2str .STATIC_URL_JS%>jquery-ui.1.11.3.min.js"></script>
<script type="text/javascript" src="<% html2str .STATIC_URL_JS%>jHtmlArea-0.8.min.js"></script>
<script type="text/javascript" src="<% CUSTOM_URL_JS%>edit_blog.js"></script>
<script type="text/javascript">
</script>
</body>
</html>