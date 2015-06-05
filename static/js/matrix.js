var mColor = new Array("#eee","#d6e685","#8cc665","#44a340","#1e6823");
var server_url = "/v1/matrix";
var MATRIX_H = 7,MATRIX_W = 30;
$(document).ready(function() {
	console.log(server_url);
	for (var i = 0; i < MATRIX_H; i++) {
		var divout= $('<div class="matrix_li" id="matrix_li_'+i+'" ></div>');
		for (var j = 0; j < MATRIX_W; j++) {
			var str = '<div class="matrix_cell" \
				id="matrix_cell_' + i + '_' + j+ '" cell_h="' + i + '" cell_w="'+ j+ '"  \
				color_id="'+ (j+i+1)%5+ '"></div>';
			var divcell=$(str);
			divcell.css("background-color", mColor[(j+i+1)%5]);
			divout.append(divcell);
		};
		$("#matrix_content").append(divout);
	};
	$(".matrix_cell").click(clickMatrixCell);
	$.ajax({
		type: 'POST',
		url: server_url,
		data: {"code":1},
		dataType: "json",
		success: function(result){
			clearMatrixAll(result.data);
		}
	});
});
function clearMatrixAll(data){
	for (var i = 0; i < MATRIX_H; i++) {
		for (var j = 0; j < MATRIX_W; j++) {
			updateMatrixColor(i,j,data[i][j]);
		};
	};
}
// 更新方格上的颜色状态
function updateMatrixColor(h,w,c){
	var div = $('#matrix_cell_' + h + '_' + w);
	div.css("background-color", mColor[c]);
	div.attr('color_id', c);
}
// 方格上的点击事件
function clickMatrixCell(cell_id){
	var w = parseInt($(this).attr("cell_w"));
	var h = parseInt($(this).attr("cell_h"));
	var c = parseInt($(this).attr("color_id"));
	console.log( (h+1) +", "+(w+1)+", "+(c));
	var new_c = (c+1)%5;
	updateMatrixColor(h,w,new_c);

	$.ajax({
		type: 'POST',
		url: server_url,
		data: {"code":2,"h":h,"w":w,"val":new_c },
		dataType: "json",
		success: function(result){
			console.log(result);
		}
	});
}