package main

import (
	"net/http"

	"github.com/squiidz/social/graphServer"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

func main() {
	mux := bone.New()
	mdlwr := claw.New(mw.Logger)
	graphServer := graphServer.New()

	mux.GetFunc("/distance/:userFrom/:userTo", graphServer.GetRelationHandler)
	mux.GetFunc("/friends/:userFrom/:userTo", graphServer.GetCommonFriendsHandler)

	http.ListenAndServe(":8080", mdlwr.Merge(mux))
}
