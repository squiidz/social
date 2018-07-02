package graph

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph store the users extracted from the facebook_combined.txt file
type Graph struct {
	Users map[int]*User
}

// NewGraph return a initialize Graph
func NewGraph() *Graph {
	graph := Graph{Users: make(map[int]*User)}
	file, err := os.Open("facebook_combined.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanr := bufio.NewScanner(file)
	for scanr.Scan() {
		if err = graph.processEntry(scanr.Text()); err != nil {
			log.Println(err)
		}
	}
	if scanr.Err() != nil {
		log.Panicln(scanr.Err())
	}
	return &graph
}

// ProcessEntry build the friend graph
func (g *Graph) processEntry(entry string) error {
	entryToken := strings.Split(entry, " ")
	userID, err := strconv.Atoi(entryToken[0])
	if err != nil {
		return fmt.Errorf("%s is not a valid user ID", entryToken[0])
	}
	friendID, err := strconv.Atoi(entryToken[1])
	if err != nil {
		return fmt.Errorf("%s is not a valid friend ID", entryToken[1])
	}

	if _, ok := g.Users[userID]; !ok {
		newUser := New(userID)
		g.Users[userID] = newUser
	}

	if _, ok := g.Users[friendID]; !ok {
		friend := New(friendID)
		g.Users[friendID] = friend
	}

	g.Users[userID].friends = append(g.Users[userID].friends, friendID)
	g.Users[friendID].friends = append(g.Users[friendID].friends, userID)

	return nil
}

// FindRelation between two nodes
func (g *Graph) FindRelation(userFrom int, userTo int) ([]int, error) {
	var visited = make(map[int]int)
	visited[userFrom] = -1

	var queue []int
	queue = append(queue, userFrom)

	for len(queue) > 0 {
		currentUser := queue[0]
		queue = queue[1:]
		if currentUser == userTo {
			return g.backtrackPath(visited, currentUser), nil
		}

		currentUserFriends := g.Get(currentUser).friends
		for _, friendID := range currentUserFriends {
			if _, visite := visited[friendID]; !visite && !contains(queue, friendID) {
				queue = append(queue, friendID)
				visited[friendID] = currentUser
			}
		}
	}
	return nil, errors.New("Relation not found")
}

// FindCommonFriends return friends in common of two users
func (g *Graph) FindCommonFriends(userFromID int, userToID int) []int {
	var commonFriendIDs []int
	userFrom := g.Get(userFromID)
	userTo := g.Get(userToID)
	for _, friendID := range userFrom.friends {
		if ok := contains(userTo.friends, friendID); ok {
			commonFriendIDs = append(commonFriendIDs, friendID)
		}
	}
	return commonFriendIDs
}

func (g *Graph) backtrackPath(visited map[int]int, userID int) []int {
	var path = []int{userID}

	for visited[userID] != -1 {
		path = append(path, visited[userID])
		userID = visited[userID]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// Get return the user with the provided id
func (g *Graph) Get(id int) *User {
	return g.Users[id]
}

// Length return the length of the graph users map
func (g *Graph) Length() int {
	return len(g.Users)
}

// DisplayRelation display the ids in the relation in string
func DisplayRelation(relationEntries []int) string {
	var relationStr string
	for _, entry := range relationEntries {
		relationStr += strconv.Itoa(entry) + "->"
	}
	relationStr = relationStr[:len(relationStr)-2]
	return relationStr
}

func contains(userIDs []int, userID int) bool {
	for _, id := range userIDs {
		if id == userID {
			return true
		}
	}
	return false
}
