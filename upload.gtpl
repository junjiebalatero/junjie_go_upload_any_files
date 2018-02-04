<!DOCTYPE html>
<html lang="en">
<head>
  <title>Junjie Website</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>

<nav class="navbar navbar-inverse">
  <div class="container-fluid">

    <!-- logo -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#mainNavBar">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a href="#" class="navbar-brand">Junjie</a>     
    </div>

    <!-- Menu Item -->
    <div class="collpase navbar-collapse" id="mainNavBar">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Home</a></li>
        <li><a href="#">About</a></li>
        <li><a href="#">Contact</a></li>

        <!-- drop down menu -->
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">My Profile <span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Friends</a></li>
            <li><a href="#">Photos</a></li>
            <li><a href="#">Settings</a></li>           
          </ul>         
        </li>
      </ul>

        <!--right align -->
        <ul class="nav navbar-nav navbar-right">
          <li><a href="#">Logout</a></li>         
        </ul>
      </ul>
    </div>
    
  </div>
  
</nav>
<form enctype="multipart/form-data" action="http://192.168.1.9:9090/upload" method="post">
          {{/* 1. File input */}}
          <input type="file" name="uploadfile" />
 
          {{/* 2. Submit button */}}
          <input type="submit" value="upload file" />
      </form><br>

  <button onclick="myFunction()">back to home page</button>

<script>
function myFunction() {
    location.href = "http://127.0.0.1";
}
</script>

 
  </body>
  </html>