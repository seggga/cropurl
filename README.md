# cropurl
course project in backend #1 golang

CropURL service is a tool that creates a short, unique URL. The service will redirect to the specific website definded by the customer. The main goal is to make a URL shorter and simpler.

## Examples to test application functionality
Let's say RESTAPI server's address is http://localhost:8080

1) redirect with a short ID

    Using web-browser is pretty handy. One of default redirect pairs is asdf -> http://google.com :

    http://localhost:8080/asdf

    or with cURL utility:

    **curl --request GET --head http://localhost:8080/asdf**


2) create new pair shortID -> longURL in the database

    You can use cURL command in terminal to create a new short link. For example you wish to create a short ID `bobo` linked with the address of golang.org to the database. You can add a description like 'just for fun'. The command you should use is like:

    **curl --request POST --header "Content-Type: application/json" --header "Accept: application/json" --data '{"ShortID":"bobo","longURL":"https://golang.org","description":"just for fun"}' http://localhost:8080/new-link**


3) delete a pair shortID -> longURL from the database

    You can use cURL command in terminal. For example you need to delete a short ID `asdf` from the database. The command line should be as follows.

    **curl --request POST --header "Content-Type: application/json" --header "Accept: application/json" --data '{"ShortID":"asdf"}' http://localhost:8080/delete**


