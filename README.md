# Middleware written with Go 1.22.3 using the new net/http

<p style="text-align:center;">
  <img src="https://github.com/IgorCastilhos/go_1.22_middleware/assets/101683017/13bd5fd7-8053-4f78-ae04-954d25e43dee"/>
</p>

## What is a middleware?

The term middleware is used to encompass all technologies that exist between the final application and the data providers for that application. On an HTTP server, we have middleware that intercepts the client's HTTP request and performs actions before and after.

### **Most common use cases**

1. Logger to record each request that arrives at the REST API
2. User session validation and maintaining active communication
3. User authentication
4. Custom logic for extracting request data
5. Add properties to client responses
6. Recover from a critical failure (* specific to go)

### Middleware Handler in Go

In Go, a middleware is a function that intercepts and modifies HTTP requests and responses before they are processed by the main handler. They are useful for adding functionality to your HTTP server without modifying the main app logic.




## According to the Go 1.22 Release Notes https://tip.golang.org/doc/go1.22#enhanced_routing_patterns
```

HTTP routing in the standard library is now more expressive. The patterns used by net/http.ServeMux have been enhanced to accept methods and wildcards.

Registering a handler with a method, like "POST /items/create", restricts invocations of the handler to requests with the given method. A pattern with a method takes precedence over a matching pattern without one. As a special case, registering a handler with "GET" also registers it with "HEAD".

Wildcards in patterns, like /items/{id}, match segments of the URL path. The actual segment value may be accessed by calling the Request.PathValue method. A wildcard ending in “…”, like /files/{path...}, must occur at the end of a pattern and matches all the remaining segments.

A pattern that ends in “/” matches all paths that have it as a prefix, as always. To match the exact pattern including the trailing slash, end it with {$}, as in /exact/match/{$}.

If two patterns overlap in the requests that they match, then the more specific pattern takes precedence. If neither is more specific, the patterns conflict. This rule generalizes the original precedence rules and maintains the property that the order in which patterns are registered does not matter.

This change breaks backwards compatibility in small ways, some obvious—patterns with “{” and “}” behave differently— and some less so—treatment of escaped paths has been improved. The change is controlled by a GODEBUG field named httpmuxgo121. Set httpmuxgo121=1 to restore the old behavior.

```
