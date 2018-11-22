package consistenthashing

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
1. create ring
2. add nodes
3. add keys
4. drop a node
5. check for key distribution
**/

var (
	node1id = "node1"
	node2id = "node2"
	node3id = "node3"
)

func TestAddNode(t *testing.T) {
	Convey("Given empty ring", t, func() {
		Convey("I should add node", func() {
			r := NewRing()
			r.AddNode(node1id)

			So(r.Nodes.Len(), ShouldEqual, 1)

			Convey("Node should have hash id", func() {
				So(r.Nodes[0].HashID, ShouldHaveSameTypeAs, uint32(0))
			})

		})

		Convey("Add nodes to a new ring and sort", func() {
			r := NewRing()
			r.AddNode(node1id)
			r.AddNode(node2id)

			So(r.Nodes.Len(), ShouldEqual, 2)

			for i, value := range r.Nodes {
				fmt.Println(i)
				fmt.Println(value)
			}

			node1hash := hashID(node1id)
			node2hash := hashID(node2id)

			So(node1hash, ShouldBeGreaterThan, node2hash)

			So(r.Nodes[0].ID, ShouldEqual, node2id)
			So(r.Nodes[1].ID, ShouldEqual, node1id)

			So(r.Nodes[0].HashID, ShouldEqual, node2hash)

		})
	})
}

func TestRemoveNode(t *testing.T) {
	Convey("Lets first create ring and nodes", t, func() {
		r := NewRing()
		r.AddNode(node1id)
		r.AddNode(node2id)
		r.AddNode(node3id)

		So(r.Nodes.Len(), ShouldEqual, 3)

		Convey("Remove a node", func() {
			err := r.RemoveNode(node2id)
			So(err, ShouldBeNil)

			So(r.Nodes.Len(), ShouldEqual, 2)

			So(r.Nodes[0].ID, ShouldEqual, node3id)
			So(r.Nodes[1].ID, ShouldEqual, node1id)

		})
	})
}

func TestGeneralFlow(t *testing.T) {
	Convey("Create ring and add nodes", t, func() {
		r := NewRing()
		r.AddNode(node1id)
		r.AddNode(node2id)
		r.AddNode(node3id)

		So(r.Nodes.Len(), ShouldEqual, 3)

		for i, value := range r.Nodes {
			fmt.Println(i)
			fmt.Println(value)
		}

		fmt.Println("==================")

		Convey("Lets remove node 2", func() {
			err := r.RemoveNode(node3id)
			So(err, ShouldBeNil)

			for i, value := range r.Nodes {
				fmt.Println(i)
				fmt.Println(value)
			}

			insertNode := "ms"
			fmt.Println(hashID(insertNode))
			insertnode := r.Get(insertNode)
			So(insertnode, ShouldEqual, node1id)
		})
	})
}
