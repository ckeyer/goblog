<!DOCTYPE html>
<html>
<head>
	<link rel="Stylesheet" type="text/css" href="<% html2str .STATIC_URL_CSS%>jquery-ui.1.11.3.min.css" />
	<link rel="Stylesheet" type="text/css" href="<% html2str .STATIC_URL_CSS%>jHtmlArea.css" />
</head>
<body>
	<div id="blog_edit">
		<input type="text" name="blog_title" value="" placeholder="标题">
		<select name="blog_type" >
			<option value="node">Node</option>
			<option value="blog">Blog</option>
			<option value="favor">Favor</option>
		</select>
		<br>
		<span id="blog_tag_list"></span>
		<input type="text" id="tag_add_text" value="" placeholder="标签">
		<input type="button" id="add_tag_button" onclick="addTags();" value="+Tag"></input>
		<br>
		<textarea name="Sub" cols="50" rows="3"  placeholder="摘要"></textarea>
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
<script type="text/javascript">
var tag_list = new
$(document).ready(function() {
	console.log($('#txtDefaultHtmlArea').htmlarea('html'));
	$("#tag_add_text").keypress(function(e){
                    if (e.which == 13) {
                        $("#add_tag_button").click();
                    }
                });
	$("#password").keypress(function(e){
                    if (e.which == 13) {
                        $("#commit_button").click();
                    }
                });
});
function addTags(){
	var tag_val = $("#tag_add_text").val();
	var tag_html = '<input type="button"  class="blog_tag" id="tag_'+tag_val+'" onclick="delTag(\''+tag_val+'\'); "value="'+tag_val+'"></input>';
	$('#blog_tag_list').append($(tag_html));
	$("#tag_add_text").val("");
	return false;
};
function delTag(event){
	$("#tag_"+event).remove();
	return false;
};
function commit(){
	var tags = new Array();
	$.each($('.blog_tag'), function(index, val) {
		tags[index] = ($(val).val());
	});
	tags.add('sdf');
	console.log(tags);
	var content = $('#txtDefaultHtmlArea').val();
	console.log(content);
};
</script>
</body>
</html>