# gRpc client server
## Steps to build application
## server
```bash
# commands shown below run at the root of the project directory
1. To Dockerize server:
  make dockerize-server
2. To compile the server side code:
    make build-server
    - To execute the binary
        ./server/cmd/userServer --port=60010
```
## client
```bash
# commands shown below run at the root of the project directory
1. To compile the client:
    make build-client

2. To execute the client binary: (client is a cli tool)
    Flags:
    Usage of ./client/cmd/userClient:
      -id int
          User ID , --id 10
      -idlist string
          User ID List , --idlist 1,2,3,4
      -server string
          grpc server addr
       -h  
          help
3. To get user by Id:
   ./client/cmd/userClient -id 1 -server "127.0.0.1:60010"

4. To get list of users by their Ids:
    ./client/cmd/userClient -idlist 1,2,3,7 -server "127.0.0.1:60010"
```
