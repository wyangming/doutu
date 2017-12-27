# doutu
为公众号制霸斗图提供服务
//实现div拖放功能
		$("#div_text").mousedown(function(event){
			var isMove=true;
//			var abs_x=event.pageX-$('#div_outimg').offset().left;
//			var abs_y=event.pageY-$('#div_outimg').offset().top;
			
			$(document).mousemove(function(event){
				if(!isMove){
					return
				}
				var obj=$("#div_text");
				obj.css({
                    'left': event.pageX-$('#div_outimg').offset().left,
                    'top': event.pageY-$('#div_outimg').offset().top
                });
			}).mouseup(function(){
				isMove=false;
			});
			
		});
		


document.getElementById("div_text").ondragstart=function(e){  
			var isMove=true;
	        //记录刚一拖动时，鼠标在飞机上的偏移量  
	        document.getElementById("div_text").ondrag=function(e){
	        	if(!isMove){
	        		return
	        	}
		        
				var x=event.pageX-$('#div_outimg').offset().left;
				var y=event.pageY-$('#div_outimg').offset().top;
				if (x<0||y<0){
					return
				}
		  
		  
		        document.getElementById("div_text").style.left=x+'px';  
		        document.getElementById("div_text").style.top=y+'px';  
		    };
		    document.getElementById("div_text").onmouseup=function(e){
		    	isMove=false;
		    };
	    };