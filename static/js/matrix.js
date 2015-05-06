var mColor = new Array("#eee","#d6e685","#8cc665","#44a340","#1e6823");
var server_url = "http://127.0.0.1:8080/matrix";
$(document).ready(function() {
	for (var i = 0; i < 7; i++) {
		var divout= $('<div class="matrix_li" id="matrix_li_'+i+'" ></div>');
		for (var j = 0; j < 30; j++) {
			var str = '<div class="matrix_cell" id=matrix_cell_"' + i + '_' + j+ '" cell_h="' + i + '" cell_w="'+ j+ '" color_id="'+ (j+i)%5+ '"></div>';
			var divcell=$(str);
			divcell.css("background-color", mColor[(j+i)%5]);
			divout.append(divcell);
		};
		$("#matrix_content").append(divout);
	};
	$(".matrix_cell").click(clickMatrixCell);
	$.ajax({
		type: 'POST',
		url: server_url,
		data: {msgcode:1},
		dataType: "json",
		success: function(result){
			console.log((result.data[1][2]));
		}
	});
});
function clickMatrixCell(event){
	var w = parseInt($(this).attr("cell_w"));
	var h = parseInt($(this).attr("cell_h"));
	var c = parseInt($(this).attr("color_id"));
	console.log( (h+1) +", "+(w+1)+", "+(c));
	var new_c = (c+1)%5;
	$(this).attr('color_id', new_c);
	$(this).css("background-color", mColor[new_c]);
}
