package client

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/grpc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//Trigger --
func Trigger(id int, idListStr, serverAddr string) {
	var err error
	var idList []int64
	if id != 0 {
		getUserByID(serverAddr, id)
	}
	if idListStr != "" {
		idList, err = checkInputData(idListStr)
		if err != nil {
			log.Fatalf("user id list input data is invalid")
		}
		getUserList(serverAddr, idList)
	}
}

func checkInputData(idListString string) ([]int64, error) {
	var idList []int64
	idStrArr := strings.Split(idListString, ",")
	for _, i := range idStrArr {
		id, err := strconv.Atoi(i)
		if err != nil {
			return idList, err
		}
		idList = append(idList, int64(id))
	}
	return idList, nil
}

func establishConnection(serverAddr string) (*grpc.ClientConn, user.UserDetailsClient) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	c := user.NewUserDetailsClient(conn)
	return conn, c
}

func getUserByID(serverAddr string, userID int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, userClient := establishConnection(serverAddr)
	if conn == nil {
		log.Fatal("Connection is lost.")
	}
	defer conn.Close()
	resp, err := userClient.GetUserByID(ctx, &user.UserReq{Id: int64(userID)})
	if err != nil {
		log.Printf("Failed to get UserInfo for a userID: %d, err: %v", userID, err)
	}
	log.Printf("Server resp: %+v\n", resp)

}
func getUserList(serverAddr string, idList []int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, userClient := establishConnection(serverAddr)
	if conn == nil {
		log.Fatal("Connection is lost.")
	}
	defer conn.Close()
	resp, err := userClient.ListUsersByID(ctx, &user.UserListReq{UserIDList: idList})
	if err != nil {
		log.Printf("Failed to get UserList for list of userID: %v, err: %v", idList, err)
	}
	log.Printf("Server resp: %+v\n", resp)
}
