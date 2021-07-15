# cropurl
course project in backend #1 golang

CropURL service is a tool that links a short URL to a long unreadable and untypable by hand link. The service will redirect to the specific website definded by the customer. The main goal is to make a URL shorter and simpler to pass.\


### chi package (https://github.com/go-chi/chi) is used as http multiplexer  


<details>
<summary>Choosing a router package</summary>
<p>
There are several mandatory requirements a router should meet:

  * function signature should be like http.HandlerFunc
  * built-in middleware
  * methods support
  * detailed documentation  
  
also important criterias are:  

  * keeping up to date
  * github rating

There was a list of routers under consideration:

-----------------------------------------------------------------
| number | package      | URL                                            |
| :------| :----------- | :--------------------------------------------- |
| 1      | net/http     | https://golang.org/pkg/net/http/               |
| 2      | gorilla/mux  | https://github.com/gorilla/mux                 |
| 3      | httprouter   | https://github.com/julienschmidt/httprouter    |
| 4      | chi          | https://github.com/go-chi/chi                  |
| 5      | fasthttp     | https://github.com/valyala/fasthttp            |
| 6      | gorouter     | https://github.com/vardius/gorouter            |
| 7      | goji         | https://github.com/goji/goji                   |
| 8      | bone         | https://github.com/go-zoo/bone                 |
--------------------------------

As a result `chi` and `gorilla/mux` are the most sutable packages that fit requirements listed above.
More details about the competition are listed below.

2, 4, 6, 7 meet mandatory requirements
according to "important things" routers are prioritized as folows 4, 2, 6, 7

1) net/http - https://golang.org/pkg/net/http/ 
- standard library
- detailed documentation
- no methods support

2) gorilla/mux - https://github.com/gorilla/mux
- stars       - 14500		
- last commit - 22 Aug 2020	
- functions signature 
- detailed documentation
- middleware support
- methods support, query-parameters, URL-prefixes, ...

3) httprouter - https://github.com/julienschmidt/httprouter	
- last commit - 21 Sep 2020
- proprietary functions signature
- there is a limitation on partial path match (short url links should be set as a query parameter)

4) chi - https://github.com/go-chi/chi
- srars - 9500
- last commit - 30 Apr 2021
- functions signature
- detailed documentation
- middleware support
- methods support

5) fasthttp - https://github.com/valyala/fasthttp
- stars - 15200
- proprietary functions signature
- last commit - 1 Jun 2021
- proprietary functions signature

6) gorouter - https://github.com/vardius/gorouter
- stars - 95
- last commit - 27 Nov 2020
- functions signature
- detailed documentation
- middleware support,
- methods support

7) goji - https://github.com/goji/goji
- stars - 	862
- last commit - 27 Jan 2019
- functions signature
- detailed documentation
- middleware support,
- methods support

8) bone - https://github.com/go-zoo/bone
- stars` - 1300	
- last commit - 17 Apr 2019
- functions signature,
- pure documentation
- methods support  
</p>
</details>


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




