##Run

mkdir api

cd api 

###clone repo

go mod init api

go mod tidy

go run .

##OR

./main


#Curl

##get
curl http://localhost:8080/albumbs

##get by id
curl http://localhost:8080/albumbs/4


##post
curl http://localhost:8080/albumbs  --include  --header "Content-Type: application/json"  --request "POST"  --data '{"id": 4, "title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

##update by id
curl http://localhost:8080/albumbs/1  --include  --header "Content-Type
: application/json"  --request "PUT"  --data '{"id": 1, "title": "The Modern Sound 
of Betty Carter","artist": "Betty Carter","price": 49.99}'

##Delete by id
curl-X DELETE http://localhost:8080/albumbs/3