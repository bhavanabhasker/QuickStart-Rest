# QuickStart-Rest
Quick start Rest program in golang 

##Objective 

The objective of this application is to provide the POST interface which prints the greeting given the user name.

##How to execute?

After cloning/downloading the source code,
<pre>
go run lab1.go 
</pre>


##Understanding the output 
When the user posts using the route "/hello"
and the body of the request is for example
<pre>
 { "name" : "John Doe" }
</pre>

The response from the application is as below,
<pre>
{ "greeting": "Hello,john!" }
</pre>


## Framework and other notes
1. Package gorilla/mux is used for dispatching the incoming requests to post handler. 
