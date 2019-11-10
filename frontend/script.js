jQuery("#demo").qrcode("novak gayzão");


$(".form").on('submit', function(e){

  let userAction = async () => {
    const response = await fetch('http://localhost:3000/carteiras/add/nome=Maria', {
      method: 'GET',
      body: myBody, // string or object
      headers: {
        'Content-Type': 'application/json'
      }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
  }
  
	e.preventDefault();
	$(this).hide();
	$('.ui.card, .trans').show();
});
