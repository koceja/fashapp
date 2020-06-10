function require() {

    var input = document.forms["username"]["personId"].value;
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



  
