package main

import (
	"net/http"
	"strconv"

	"github.com/squiidz/social/graphServer"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

func main() {
	mux := bone.New()
	mdlwr := claw.New(mw.Logger, mw.Zipper)
	mux.RegisterValidatorFunc("isNum", isNum)
	graphServer := graphServer.New()

	mux.GetFunc("/distance/:userFrom|isNum/:userTo|isNum", graphServer.GetRelationHandler)
	mux.GetFunc("/friends/:userFrom|isNum/:userTo|isNum", graphServer.GetCommonFriendsHandler)

	http.ListenAndServe(":8080", mdlwr.Merge(mux))
}

func isNum(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
