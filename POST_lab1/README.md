Download the files from github

Post Installations :
go get github.com/gorilla/mux

Execution :
go run lab1.go

Open the POSTman or similar interface

Select the option POST
Enter the http://localhost:1234/hello
Enter the body 
{
    "name"  : "john"
}
The Response will be in below format
{
  "greeting": "Hello,john!"
}
