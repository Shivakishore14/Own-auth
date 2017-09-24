; (function(ownAuth, window, document, undefined) {
	var widget = ownAuth.widget = ownAuth.widget || (function() {
		//console.log("loading widget");
		var base_url = "http://localhost:9000";
		var init = function(config){
			base_url = config.base_url || base_url;
		}
		var observer = new MutationObserver(function(mutations) {
			mutations.forEach(function(mutation) {
				if (mutation.type == 'childList')
					setCssForCentering();
			});
		});
		var config = { childList: true, characterData: true, subtree:true };

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

								var target = document.getElementById("own-auth-root");
								observer.observe(target, config);
								setCssForCentering();
								initializeEventListners();
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
					var xhttp = new XMLHttpRequest();
					var errorText = document.getElementById("own-auth-error-text");
					var ownAuthRoot = document.getElementById("own-auth-root");

					errorText.innerHTML = "";
					user.username = document.getElementById('ownAuthUsername').value;
					user.password = document.getElementById('ownAuthPassword').value;

					xhttp.onreadystatechange = function() {
						if (this.readyState == 4) {
							if (this.status == 200) {
								// console.log(this.responseText);
								data = JSON.parse(this.responseText);
								if (data.message.toLowerCase() === "success"){
									observer.disconnect();
									ownAuthRoot.parentNode.removeChild(ownAuthRoot);
									resolve(data.data);
								}else{
									errorText.innerHTML = data.message;
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
		var setCssForCentering = function(){
			var wrapper = document.getElementsByClassName('own-auth-wrapper')[0];
			var height = wrapper.clientHeight;
			var width = wrapper.clientWidth;
			wrapper.style["margin-top"] = (- (height/2))+"px";
			wrapper.style["margin-left"] = (- (width/2))+"px";
		}
		var initializeEventListners = function(){
			var loginText = document.getElementById("own-auth-login-footer-text");
			var signupText = document.getElementById("own-auth-signup-footer-text");
			var signUpButton = document.getElementById("ownAuthSignUpBtn");
			loginText.addEventListener("click", showLoginScreen);
			signupText.addEventListener("click", showSignUpScreen);
			signUpButton.addEventListener("click", signUp);
		}
		var showLoginScreen = function(){
			var loginWrapper = document.getElementsByClassName("own-auth-login")[0];
			var signUpWrapper = document.getElementsByClassName("own-auth-signup")[0];
			loginWrapper.style.display = "inline-block";
			signUpWrapper.style.display = "none";
			setCssForCentering();
		}
		var showSignUpScreen = function(){
			var loginWrapper = document.getElementsByClassName("own-auth-login")[0];
			var signUpWrapper = document.getElementsByClassName("own-auth-signup")[0];
			loginWrapper.style.display = "none";
			signUpWrapper.style.display = "inline-block";
			setCssForCentering();
		}
		var signUp = function(){
			var user = {}
			var xhttp = new XMLHttpRequest();
			var errorText = document.getElementById("own-auth-error-text-signup");
			var successText = document.getElementById("own-auth-success-text-signup");
			var ownAuthRoot = document.getElementById("own-auth-root");

			errorText.innerHTML = "";
			user.username = document.getElementById('ownAuthUsernameSignUp').value;
			user.password = document.getElementById('ownAuthPasswordSignUp').value;
			user.name = document.getElementById('ownAuthNameSignUp').value;
			user.phone = document.getElementById('ownPhoneSignUp').value;
			user.email = document.getElementById('ownEmailSignUp').value;

			xhttp.onreadystatechange = function() {
				if (this.readyState == 4) {
					if (this.status == 200) {
						// console.log(this.responseText);
						data = JSON.parse(this.responseText);
						if (data.message.toLowerCase() === "success"){
							resolve(data);
							successText.innerHTML = "Done";
						}else{
							errorText.innerHTML = data.message;
						}
					}else{
						errorText.innerHTML = "Server Error";
						console.log(Error("GOT status code "+this.status+" please check the netwoks tab in dev tools"));
					}
				}
			};
			xhttp.open("POST", base_url+"/api/signup", true);
			xhttp.setRequestHeader("Content-type", "text/json");
			xhttp.send(JSON.stringify(user));
		}
		return {
			loadWidget:loadWidget,
			login:login
		}
	}());
}(window._ownAuth = window._ownAuth || {}, window, document));
