<div  class="body_content_label">
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
		<textarea id="blog_summary" cols="49" rows="3"  placeholder="摘要"></textarea>
		<textarea id="txtDefaultHtmlArea" cols="50" rows="25"  ></textarea>
		<input type="password" id="password" value="" placeholder="密码">
		<input type="button" id="commit_button" onclick="commit();" value="123543"></input>
	</div>
</div>