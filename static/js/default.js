var  matrix, article ; 

//  Github的Commit矩阵
matrix=(function(){
	var  proto={ 
		mColor : new Array("#eee","#d6e685","#8cc665","#44a340","#1e6823"),
		server_url : "/matrix",
		MATRIX_H : 7,
		MATRIX_W : 30,

		// 更新方格上的颜色状态
		updateMatrixColor:function (h,w,c){
			var div = $('#matrix_cell_' + h + '_' + w);
			div.css("background-color", proto.mColor[c]);
			div.attr('color_id', c);
		},
		clearMatrixAll:function (data){
			for (var i = 0; i < proto.MATRIX_H; i++) {
				for (var j = 0; j < proto.MATRIX_W; j++) {
					proto.updateMatrixColor(i,j,data[i][j]);
				};
			};
		},
		// 方格上的点击事件
		clickMatrixCell:function (cell_id){
			var w = parseInt($(this).attr("cell_w"));
			var h = parseInt($(this).attr("cell_h"));
			var c = parseInt($(this).attr("color_id"));
			// console.log( (h+1) +", "+(w+1)+", "+(c));
			var new_c = (c+1)%5;
			proto.updateMatrixColor(h,w,new_c);

			$.ajax({
				type: 'POST',
				url: proto.server_url,
				data: {"code":2,"h":h,"w":w,"val":new_c },
				dataType: "json",
				success: function(result){
					console.log("Download your use my life");
				}
			});
		},
		//$(document).ready(
		init:function() {
			for (var i = 0; i < proto.MATRIX_H; i++) {
				var divout= $('<div class="matrix_li" id="matrix_li_'+i+'" ></div>');
				for (var j = 0; j < proto.MATRIX_W; j++) {
					var str = '<div class="matrix_cell" \
					id="matrix_cell_' + i + '_' + j+ '" cell_h="' + i + '" cell_w="'+ j+ '"  \
					color_id="'+ (j+i+1)%5+ '"></div>';
					var divcell=$(str);
					divcell.css("background-color", proto.mColor[(j+i+1)%5]);
					divout.append(divcell);
				};
				$("#matrix_content").append(divout);
			};
			$(".matrix_cell").click(proto.clickMatrixCell);
			$.ajax({
				type: 'POST',
				url: proto.server_url,
				data: {"code":1},
				dataType: "json",
				success: function(result){
					proto.clearMatrixAll(result.data);
				}
			});
		}
	}
	return proto;
})();

article =(function(){
	var proto ={
		server_url : "/blog",
		init:function(){
			$(".article_title").click(proto.clickArticleTitle);
			$(".article_summary").click(proto.clickArticleTitle);
		},
		clickArticleTitle:function(){
			var id = parseInt($(this).attr("art_id"));
			proto.getArticle(id);
		},
		getArticle:function(id){
			$.ajax({
				type: 'POST',
				url: proto.server_url,
				data: {"code":1,"id":id },
				dataType: "json",
				success: function(result){
					console.log("Success");
					proto.showArticle(result.data);
				}
			});
		},
		showArticle:function(data){
			console.log("hello");
			console.log(data);
		}
	}
	return proto;
})();


main = function(){
	article.init();
	matrix.init();
};

$(document).ready(main());