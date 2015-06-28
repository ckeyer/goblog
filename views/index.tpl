<!DOCTYPE html>
<html>

<head>
      <!--<meta http-equiv="refresh" content="0;url=https://www.ckeyer.com/">-->
      <meta http-equiv="content-type" content="text/html; charset=utf-8" />
      <meta name="keywords" content="" />
      <meta name="description" content="" />
      <link rel="stylesheet" type="text/css" href="<% CUSTOM_URL_CSS %>matrix.css" />
      <link rel="shortcut icon" href="<% CUSTOM_URL_IMG%>i_logo1.png" >
      <title><% .PageTitle %></title>
      <link rel="stylesheet" href="<% CUSTOM_URL_CSS%>style.css" media="screen" type="text/css" />

</head>

<body>

      <div class="left-menu">
            <div class="logo"><i class="fa fa-align-justify"></i>
                  <div><h2> Ckeyer</h2></div>
                  <div>
                        <font size="3">Man, just luv techoligy.</font>
                </div>
        </div>
        <div class="accordion">
              <div class="sectionlabel">
                    <input type="radio" name="accordion-1" id="sectionlabel-1" checked="checked"/>
                    <label for="sectionlabel-1"><span class="sectionlabel-title">博客</span></label>
                    <div class="content">
                          <ul>
                                <% range $index, $elem := .BlogsMonth %>
                                <li class="tag_label" onclick="console.log('<%$elem.Month%>');">
                                      <i class="fa fa-inbox"></i>
                                      <span><% $elem.Month %></span>
                                      (<span><% $elem.BlogCount %></span>)
                              </li>
                              <% end %>
                      </ul>
              </div>
      </div>

      <div class="sectionlabel">
            <input type="radio" name="accordion-1" id="sectionlabel-2" value="toggle"/>
            <label for="sectionlabel-2"><span>标签</span></label>
            <div class="content">
                  <ul>
                        <ul>
                              <% range $index, $elem := .BlogsTag %>
                              <li class="tag_label" tag_id="<% $elem.Id %>">
                                    <i class="fa fa-inbox"></i>
                                    <span><% $elem.Name %></span>
                                    (<span><% $elem.BlogCount %></span>)
                            </li>
                            <% end %>
                    </ul>
            </div>
    </div>

    <div class="sectionlabel">
      <input type="radio" name="accordion-1" id="sectionlabel-3" value="toggle"/>
      <label for="sectionlabel-3"><span>收藏夹 </span></label>
      <div class="content">
            <ul>
                  <ul>
                        <% range $index, $elem := .BlogsTag %>
                        <li class="tag_label" tag_id="<% $elem.Id %>">
                              <i class="fa fa-inbox"></i>
                              <span><% $elem.Name %></span>
                              (<span><% $elem.BlogCount %></span>)
                      </li>
                      <% end %>
              </ul>
      </div>
</div>

<div class="sectionlabel">
    <input type="radio" name="accordion-1" id="sectionlabel-4" value="toggle"/>
    <label for="sectionlabel-4"><span>工具</span></label>
    <div class="content">
          <ul>
                <li onclick="window.location.href='/chat';"><i class="fa fa-coffee"></i><span>聊天室</span></li>
                <li onclick="window.location.href='http://earth.ckeyer.com/';"><i class="fa fa-coffee"></i><span>地球</span></li>
                <li onclick="window.location.href='http://ntop.ckeyer.com/';"><i class="fa fa-coffee"></i><span>流量监控</span></li>
                <li onclick="window.location.href='/admin/blog';"><i class="fa fa-coffee"></i><span>管理</span></li>
        </ul>
</div>
</div>

<div class="sectionlabel">
    <input type="radio" name="accordion-1" id="sectionlabel-5" value="toggle"/>
    <label for="sectionlabel-5"onclick="clic('/admin');"><span>留言</span></label>
    <div class="content">
          给我留言
  </div>
</div>

<div class="sectionlabel">
    <input type="radio" name="accordion-1" id="sectionlabel-6" value="toggle"/>
    <label for="sectionlabel-6"><span>关于</span></label>
    <div class="content">
          <h3>Github.</h3>
          <a href="https://github.com/ckeyer" target="_blank">https://github.com/ckeyer</a> 
          <h3>Email.</h3>
          <a href="Mailto:me@ckeyer.com">me@ckeyer.com</a> <br>
          <h4><a href="/about/me.html">More...</a> </h4>

  </div>
</div>
</div>
</div>

<div id ="body_content" style="text-align:center;clear:both">
        <div id="body_content_1" class="body_content_label">
                <div id="matrix_content"></div>
        </div>

        <div id="body_content_2">
                  <% range $index, $elem := .LatestBlogs %>
                  <div class="body_content_label" tag_id="<% $elem.Id %>">
                        <div class="body_content_article">
                                <div class="article_title_div">
                                        <span class="article_title_tag">[BLOG]</span>
                                        <span class="article_title"> <% $elem.Title %></span>
                                </div>
                                <div class="article_tags">
                                        <% range $index2,$elem2 := $elem.Tags %>
                                        <span class="article_tag" id="<% $elem2.Id  %>"><% $elem2.Name %></span>
                                        <% end %>
                                        <span class="article_created">
                                                <% FMT_DATETIME $elem.Created %>
                                        </span>
                                </div>
                                <div class="article_summary">
                                        <% $elem.Summary %>
                                </div>
                        </div>
                </div>
                <% end %>
        </div>

        <div class="body_content_label">
                
                  helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>546<br>546 <br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>514<br>514<br>514<br>514 <br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>
                <br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>helloasdfabr<br>asdf
                </div>

        </div>

<canvas id="body_bg" >A drawing of something...</canvas>

<script type="text/javascript" src="<% STATIC_URL_JS %>jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="<% CUSTOM_URL_JS %>matrix.js"></script>
<script src="<% STATIC_URL_JS %>modernizr.js"></script>
<script src='<% STATIC_URL_JS %>dat.gui.min.js'></script>
<script src='<% STATIC_URL_JS %>toxiclibs.min.js'></script>
<script src='<% STATIC_URL_JS %>animitter.min.js'></script>
<script src="<% STATIC_URL_JS %>bg_index.js"></script>
<script type="text/javascript" charset="utf-8" >
    function clic(new_url){
          var title = document.title;
          var url = window.location.pathname;
          var state = { title: title,url:url};
          bg_app.reset();
          console.log(state);
          window.history.pushState(state, document.title, new_url);
  };

</script>
</body>
</html>
