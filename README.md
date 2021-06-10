# cropurl
course project in backend #1 golang

CropURL service is a tool that creates a short, unique URL. The service will redirect to the specific website definded by the customer. The main goal is to make a URL shorter and simpler.


## Choosing a router package

There are several mandatory requirements a router should meet:
    function signature should be like http.HandlerFunc
    built-in middleware
    methods support
    detailed documentation
also important criterias are:
    keeping up to date
    github rating

There was a list of routers under consideration:

-----------------------------------------------------------------
| package      | URL                                            |
| :----------- | :--------------------------------------------- |
| net/http     | https://golang.org/pkg/net/http/               |
| gorilla/mux  | https://github.com/gorilla/mux                 |
| httprouter   | https://github.com/julienschmidt/httprouter    |
| chi          | https://github.com/go-chi/chi                  |
| fasthttp     | https://github.com/valyala/fasthttp            |
| gorouter     | https://github.com/vardius/gorouter            |
| goji         | https://github.com/goji/goji                   |
| bone         | https://github.com/go-zoo/bone                 |
--------------------------------

As a result `chi` and `gorilla/mux` are the most sutable packages that fit requirements listed above.
More details about the competition are in `choose_mux.txt`.


