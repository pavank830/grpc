package user

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

type Server struct {
	UnimplementedUserDetailsServer
}

// User info
var (
	UserInfoList = map[int]*UserInfo{
		1: {
			Id:      1,
			Fname:   "Steve",
			City:    "LA",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
		2: {
			Id:      2,
			Fname:   "Denver",
			City:    "Berlin",
			Phone:   1212343456,
			Height:  5.1,
			Married: false,
		},
		3: {
			Id:      3,
			Fname:   "Akash",
			City:    "Sydney",
			Phone:   3467091255,
			Height:  5.3,
			Married: false,
		},
		4: {
			Id:      4,
			Fname:   "John",
			City:    "Delhi",
			Phone:   2001567265,
			Height:  5.6,
			Married: true,
		},
	}
)

func getUserInfo(id int64) (*UserInfo, error) {
	if val, ok := UserInfoList[int(id)]; ok {
		return val, nil
	}
	return &UserInfo{}, fmt.Errorf("user not found with id: %d", id)
}

func getUserList(idList []int64) ([]*UserInfo, error) {
	userList := []*UserInfo{}
	var err error
	for _, id := range idList {
		user, err := getUserInfo(id)
		if err != nil {
			continue
		}
		userList = append(userList, user)
	}
	if len(userList) == 0 {
		err = fmt.Errorf("no users found for the requested ids: %v", idList)
	}
	return userList, err
}

// GetUserByID - get user info for given user ID.
func (*Server) GetUserByID(ctx context.Context, req *UserReq) (*UserResp, error) {
	resp := &UserResp{}
	if req == nil {
		return nil, errors.New("request is nil")
	}
	if req.Id <= 0 {
		resp.Code = int32(codes.InvalidArgument)
		resp.Message = "ID should be a positive value"
		return resp, nil
	}
	userInfo, err := getUserInfo(req.Id)
	if err != nil {
		resp.Code = int32(codes.NotFound)
		resp.Message = err.Error()
		return resp, err
	}
	resp.User = userInfo
	resp.Code = int32(codes.OK)
	return resp, err
}

// ListUsersByID - list users by IDs
func (*Server) ListUsersByID(ctx context.Context, in *UserListReq) (*UserListResp, error) {
	resp := &UserListResp{}
	if in == nil {
		return nil, errors.New("request is nil")
	}

	if len(in.UserIDList) == 0 {
		resp.Code = int32(codes.InvalidArgument)
		resp.Message = "ID list is empty"
		return resp, nil
	}
	userList, err := getUserList(in.UserIDList)
	if err != nil {
		resp.Code = int32(codes.NotFound)
		resp.Message = err.Error()
		return resp, err
	}
	resp.Code = int32(codes.OK)
	resp.UserList = userList
	return resp, nil
}
