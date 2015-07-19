$ ->
	console.log "fuck"
	$('#div_commit').click ->
		console.log "hello"
		$.post(  
			"http://localhost/test" 
			new_title: "hello"  
			id: "where"  
			-> log "hello"
			'json' 
		) 

log  (s)->
	console.log s