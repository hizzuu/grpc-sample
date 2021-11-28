// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/hizzuu/grpc-example-bff/gen/pb"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserPayload struct {
	User *pb.User `json:"user"`
}

type GetCurrentUserPayload struct {
	User *pb.User `json:"user"`
}

type GetUserPayload struct {
	User *pb.User `json:"user"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	User  *pb.User  `json:"user"`
	Token *pb.Token `json:"token"`
}

type PageInfo struct {
	HasNextPage bool    `json:"hasNextPage"`
	EndCursor   *string `json:"endCursor"`
}
