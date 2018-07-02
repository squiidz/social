package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/squiidz/social/user"

	"github.com/go-zoo/bone"
)

var graph = user.NewGraph()

// GetRelationHandler return the relation between 2 users if it exist
func GetRelationHandler(rw http.ResponseWriter, req *http.Request) {
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

	relationIDs, err := graph.FindRelation(userFrom, userTo)
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

// GetCommonFriendHandler return the friends in common between 2 users
func GetCommonFriendHandler(rw http.ResponseWriter, req *http.Request) {
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

	commonFriends := graph.FindCommonFriends(userFrom, userTo)
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
