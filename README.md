# go-web-dev-side-project-1
go web development side project one

This is a Golang project and includes :

- HTTP server 
- HTML dynamic templates 
- Saving and fetching contents into Postgres DB
- Middlware design following three methods; 
  - Variadic functions
  - Function chaining
  - Callbacks
  - Context package and Context trees
- Logger and defer-panic-recover middlewares 
- Dashboard of content management system "CMS"
- Service to capture usage analytics
- User Authentication service to authenticate requests
  - Username and Password login encryption using bcrypt (scrypt) package
  - Saving Sessions and Cookies using bot package (gorilla sessions) into DB
  - OAuth, JWT Token authentication using jwt-go, oauth2 (sessionless or passwordless authentication)
  - Use Google OAuth2 API to authenticate then create valid token
  - Use of sanitizers functions to capture SQL Injection or JS Scripts 
  - Use TLS over HTTP to certify HTTPS requests by checking locally generated server key but without server certificate over the browser
- Data RESTFul API 
  - JSON marhsalling, unmarshalling and streaming
  - Formatting JSON responses
- Real-time chat and notification services for visitor help 
- CI/CD automated deployment service
