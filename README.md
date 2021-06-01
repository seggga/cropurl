# cropurl
course project in backend #1 golang

CropURL service is a tool that creates a short, unique URL. The service will redirect to the specific website definded by the customer. The main goal is to make a URL shorter and simpler.


Open API

the service implements functions as follows:
1) user authorisation 
    - /user/login       - POST
    - /user/logout      - POST
2) create short URL that corresponds to specified long URL
    - /new_linknks      - POST
3) manage stored short URLs
    - /links/{shortURL} - GET       - view the statistics
    - /links/{shortURL} - DELETE    - delete short URL
4) redirect users request to the long URL
    - /{shortURL}       - GET