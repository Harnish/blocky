package web

// IndexTmpl html template for the start page
const IndexTmpl = `<!DOCTYPE html>
<html>
	<head>
		<title>blocky</title>
		<script language="JavaScript">
		// http://stackoverflow.com/questions/12460378/how-to-get-json-from-url-in-javascript
		var getJSON = function(url, callback) {
		  var xhr = new XMLHttpRequest();
		  xhr.open('GET', url, true);
		  xhr.responseType = 'json';
		  xhr.onload = function() {
			var status = xhr.status;
			if (status == 200) {
			  callback(null, xhr.response);
			} else {
			  callback(status);
			}
		  };
		  xhr.send();
		 };
		 // http://stackoverflow.com/questions/901115/how-can-i-get-query-string-values-in-javascript
		 function getParameterByName(name, url) {
		  if (!url) url = window.location.href;
		  name = name.replace(/[\[\]]/g, "\\$&");
		  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
			  results = regex.exec(url);
		  if (!results) return null;
		  if (!results[2]) return '';
		  return decodeURIComponent(results[2].replace(/\+/g, " "));
	  }
		 function ResolveName() {
			  var name = document.getElementById('name').value;
			  var types = document.getElementsByName('type');
			  var i;
			  var mytype = 255
			  for (i=0; i < types.length; i++){
				  if (types[i].checked == true){
					  mytype = types[i].value;
				  }
			  }
			  var hosturl = window.location.host;
			  var protocol = window.location.protocol;
			  var prefixurl = protocol + "//" + hosturl;
			  document.getElementById("directurl").innerHTML = '<a href="'+ prefixurl + '/resolve?name=' + name + '&type=' + mytype + '">' + prefixurl + '/resolve?name=' + name + '&type=' + mytype + '</a><br>';
			  
			  var encodedString = window.btoa(name);
			  console.log(encodedString);
			
			  getJSON('/resolve?name=' + name,
	  function(err, data) {
		if (err != null) {
		  alert('Something went wrong: ' + err);
		} else {
		  document.getElementById("json").innerHTML = JSON.stringify(data, undefined, 2);
		}
	  });
	  }

		</script>
	</head>
	<body>
		<h1>blocky</h1>
		<hl>
		{{range .Links}}
			<a href="{{.URL}}">{{.Title}}</a> 
		{{end}}
		<hl>

		<form>
<table width=100% bgcolor="#808080">
<tr><td>DNS Lookup</td><td><input type="text" name="name" id="name"></td><td><input type="button" name="lookup" value="Resolve" OnClick=ResolveName()></td></tr>
</table>
<table>
  <tr>
    <td><input type="radio" name="type" id="type-1" value="1"> A </td>
	<td><input type="radio" name="type" id="type-28" value="28"> AAAA </td>
	<td><input type="radio" name="type" id="type-5" value="5"> CNAME </td>
	<td><input type="radio" name="type" id="type-15" value="15"> MX </td>
	<td><input type="radio" name="type" id="type-255" value="255" checked> ANY </td>
   </tr>
</table>
</form>
<hr>
Results:<br>
<pre id="json"></pre>
Short cut URL: <p id="directurl"></p>

		<p><span class="small">Version {{.Version}}</span></p> 
		</body>
	</html>`
