package main

import "fmt"

type Flow struct {
	Intent Intent
	Nodes  []Node
}

type Intent interface{}

type NodeType string

const NodeTypeSimple NodeType = "SIMPLE"
const NodeTypeSubFlow NodeType = "SUB_FLOW"

type Node struct {
	Type    NodeType   `json:"type"`
	Simple  NodeSimple `json:"simple,omitempty"`
	SubFlow *Flow      `json:"flow,omitempty"`
}

type NodeSimple interface{}

func RunEffects(f *Flow) {
	RunEffectsForFlow(f)

}

func RunEffectsForFlow(f *Flow) {
	for _, node := range f.Nodes {
		node := node
		RunEffectsForNode(node)
	}
	RunEffectsForIntent(f.Intent)
}

func RunEffectsForNode(n Node) {
	switch n.Type {
	case NodeTypeSimple:
		RunEffectsForNodeSimple(n.Simple)
	case NodeTypeSubFlow:
		RunEffectsForFlow(n.SubFlow)
	}
}

func RunEffectsForNodeSimple(v interface{}) {
	if e, ok := v.(interface{ RunEffect() }); ok {
		e.RunEffect()
	}
}

func RunEffectsForIntent(v interface{}) {
	if e, ok := v.(interface{ RunEffect() }); ok {
		e.RunEffect()
	}
}

type SignupIntent struct{}

func (*SignupIntent) RunEffect() {
	fmt.Printf("POST https://myapp.com/authgear/events/user.created\n")
}

type NodeCreateUser struct {
	UserID string
}

func (n *NodeCreateUser) RunEffect() {
	fmt.Printf("INSERT INTO user (id) VALUES ('%v');\n", n.UserID)
}

type NodeCreateIdentityLoginID struct {
	UserID  string
	LoginID string
}

func (n *NodeCreateIdentityLoginID) RunEffect() {
	fmt.Printf("INSERT INTO login_id (user_id, login_id) VALUES ('%v', '%v');\n", n.UserID, n.LoginID)
}

type NodeCreatePassword struct {
	UserID string
	Hash   string
}

func (n *NodeCreatePassword) RunEffect() {
	fmt.Printf("INSERT INTO password (user_id, hash) VALUES ('%v', '%v');\n", n.UserID, n.Hash)
}

func main() {
	tree := &Flow{
		Intent: &SignupIntent{},
		Nodes: []Node{
			Node{
				Type: NodeTypeSimple,
				Simple: &NodeCreateUser{
					UserID: "user0",
				},
			},
			Node{
				Type: NodeTypeSimple,
				Simple: &NodeCreateIdentityLoginID{
					UserID:  "user0",
					LoginID: "louischan@oursky.com",
				},
			},
			Node{
				Type: NodeTypeSimple,
				Simple: &NodeCreatePassword{
					UserID: "user0",
					Hash:   "password-hash",
				},
			},
		},
	}
	RunEffects(tree)
}
