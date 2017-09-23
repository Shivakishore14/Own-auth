; (function(ownAuth, window, document, undefined) {
	var widget = ownAuth.widget = ownAuth.widget || (function() {
		//console.log("loading widget");
		var base_url = "http://localhost:9000";
		var isInitialized = false;
		var init = function(config){
			base_url = config.base_url || base_url;
			isInitialized = true;
		}
		var loadWidget = function(){
			var promise = new Promise(function(resolve, reject) {
				var xhttp = new XMLHttpRequest();
				xhttp.onreadystatechange = function() {
					if (this.readyState == 4) {
						if (this.status == 200) {
							wrapper = document.getElementsByTagName('html');
							content = document.createElement('div')
							content.innerHTML = this.responseText;
							if (wrapper.length != 0){
								wrapper[0].appendChild(content);
								resolve("done");
							}else{
								//console.log("Error cannot find `html` tag");
								reject(Error("Error cannot find `html` tag"));
							}
						}else{
							reject(Error("GOT status code "+this.status+" please check the netwoks tab in dev tools"));
						}
					}
				};
				xhttp.open("GET", "widget.html", true);
				xhttp.send();
			});
			return promise
		}
		var login = function(){
			button = document.getElementById("ownAuthLoginBtn");
			var promise = new Promise(function(resolve, reject) {
				button.addEventListener("click", function(){
					var user = {}
					user.username = document.getElementById('ownAuthUsername').value;
					user.password = document.getElementById('ownAuthPassword').value;
					var xhttp = new XMLHttpRequest();
					xhttp.onreadystatechange = function() {
						if (this.readyState == 4) {
							if (this.status == 200) {
								console.log(this.responseText);
								data = JSON.parse(this.responseText);
								//console.log(data);
								if (data.message.toLowerCase() === "success"){
									resolve(data.data);
								}else{
									reject(Error(data.message));
								}
							}else{
								reject(Error("GOT status code "+this.status+" please check the netwoks tab in dev tools"));
							}
						}
					};
					xhttp.open("POST", base_url+"/api/login", true);
					xhttp.setRequestHeader("Content-type", "text/json");
					xhttp.send(JSON.stringify(user));
				});
			});
			return promise;
		}
		return {
			loadWidget:loadWidget,
			login:login
		}
	}());
}(window._ownAuth = window._ownAuth || {}, window, document));
