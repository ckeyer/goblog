<div class="body_content_label" tag_id="<% .Article.Id %>">
	<div class="body_content_article">
		<div class="article_title_div">
			<span class="article_title_tag">[BLOG]</span>
			<span class="article_title" art_id="<% .Article.Id  %>"> <% .Article.Title %></span>
		</div>

		<div class="article_tags">
			<% range $index2,$elem2 :=  .Article.Tags %>
			<span class="article_tag" id="<% $elem2.Id  %>"><% $elem2.Name %></span>
			<% end %>
			<span class="article_created">
			<% FMT_DATETIME .Article.Created %>
			</span>
		</div>
		<div class="article_content">
			<% str2html .Article.Content %>
		</div>
	</div>
</div>