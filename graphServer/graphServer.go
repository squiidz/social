package graphServer

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/squiidz/social/graph"

	"github.com/go-zoo/bone"
)

// GraphServer provide handler with access to a graph without delcaring a global one
type GraphServer struct {
	graph *graph.Graph
}

// New return a initialize GraphServer
func New() GraphServer {
	return GraphServer{graph: graph.NewGraph()}
}

// GetRelationHandler return the relation between 2 users if it exist
func (gs GraphServer) GetRelationHandler(rw http.ResponseWriter, req *http.Request) {
	params := bone.GetAllValues(req)

	userFrom, err := strconv.Atoi(params["userFrom"])
	if err != nil {
		rw.Write([]byte("invalid user id"))
		return
	}

	userTo, err := strconv.Atoi(params["userTo"])
	if err != nil {
		rw.Write([]byte("invalid user id"))
		return
	}

	relationIDs, err := gs.graph.FindRelation(userFrom, userTo)
	if err != nil {
		http.NotFound(rw, req)
		return
	}

	relation := struct {
		Distance int   `json:"distance"`
		Path     []int `json:"path"`
	}{len(relationIDs), relationIDs}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(relation)
}

// GetCommonFriendsHandler return the friends in common between 2 users
func (gs GraphServer) GetCommonFriendsHandler(rw http.ResponseWriter, req *http.Request) {
	params := bone.GetAllValues(req)

	userFrom, err := strconv.Atoi(params["userFrom"])
	if err != nil {
		rw.Write([]byte("invalid user id"))
		return
	}

	userTo, err := strconv.Atoi(params["userTo"])
	if err != nil {
		rw.Write([]byte("invalid user id"))
		return
	}

	commonFriends := gs.graph.FindCommonFriends(userFrom, userTo)
	if len(commonFriends) == 0 {
		http.NotFound(rw, req)
		return
	}

	common := struct {
		Friends []int `json:"friends"`
	}{commonFriends}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(common)
}
