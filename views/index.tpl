<% range $index, $elem := .LatestBlogs %>
<div class="body_content_label" tag_id="<% $elem.Id %>">
      <div class="body_content_article">
              <div class="article_title_div">
                      <span class="article_title_tag">[BLOG]</span>
                      <span class="article_title" art_id="<% $elem.Id  %>"> <% $elem.Title %></span>
              </div>
              <div class="article_tags">
                      <% range $index2,$elem2 := $elem.Tags %>
                      <span class="article_tag" id="<% $elem2.Id  %>"><% $elem2.Name %></span>
                      <% end %>
                      <span class="article_created">
                              <% FMT_DATETIME $elem.Created %>
                      </span>
              </div>
              <div class="article_summary" art_id="<% $elem.Id  %>">
                      <% $elem.Summary %>
              </div>
      </div>
</div>
<% end %>