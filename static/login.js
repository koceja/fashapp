function require() {

    var input = document.forms["username"]["personId"].value;
    for (var i = 0; i < input.length; i++) {
        if (input[i] == "/" || input[i] == " " || input[i] == "$" || input[i] == "%" || input[i] == "&") {
            return false;
        }
    }
    if (input == "") {
        alert("Please enter a non-empty username.");  	
        return false; 
    } else {
        return true; 
    }	
  };


$("#username").submit(function() {
    return require(this);
});



  
