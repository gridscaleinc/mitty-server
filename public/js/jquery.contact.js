(function($) {
	"use strict";

jQuery(document).ready(function(){
	$('#cform').submit(function(){

		var action = $(this).attr('action');

		$("#message").slideUp(750,function() {
		$('#message').hide();

 		$('#submit')
			//.before('<img src="/img/fancybox_loading.gif" class="contact-loader" />')
			.attr('disabled','disabled');

		$.post(action, {
			name: $('#name').val(),
			email: $('#email').val(),
			comment: $('#comment').val(),
		},
			function(data){
				document.getElementById('message').innerHTML = "Email Sent Successfully.";
				$('#message').slideDown('slow');
				$('#cform img.contact-loader').fadeOut('slow',function(){$(this).remove()});
				$('#submit').removeAttr('disabled');
				$('#cform').slideUp('slow');
			}
		);

		});

		return false;

	});

});

}(jQuery));
