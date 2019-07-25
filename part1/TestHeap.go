package part1

import (
	"container/heap"
	"fmt"
	"github.com/xtario/xtar/core/types"
	"math/rand"
	"strconv"
	"time"
)

const defaultLeng = 100

// the BidName which has been bidden correspond to a HeapNode to record the position of the BidName in Heap
type HeapNode struct {
	BidName string
	Amount  uint64
}

type Bid struct {
	BidderAddress types.Address
	Amount        uint64
	BidTime       uint32 //time by block_height
	IsClosed      bool
}

type BidHeap struct {
	HeapNodes  []*HeapNode
	LengthNode *HeapNode
}

func (bidHeap *BidHeap) Copy(dst, src *HeapNode) {
	dst.Amount = src.Amount
	dst.BidName = src.BidName
}

var bidHeap *BidHeap

func (bidHeap *BidHeap) Len() int {
	return int(bidHeap.LengthNode.Amount)
}

func (bidHeap *BidHeap) Less(i, j int) bool {
	return bidHeap.HeapNodes[i].Amount > bidHeap.HeapNodes[j].Amount ||
		 (bidHeap.HeapNodes[i].Amount == bidHeap.HeapNodes[j].Amount && i>j)

}

func (bidHeap *BidHeap) Swap(i, j int) {
	tmp := &HeapNode{
	}
	bidHeap.Copy(tmp, bidHeap.HeapNodes[i])
	bidHeap.Copy(bidHeap.HeapNodes[i], bidHeap.HeapNodes[j])
	bidHeap.Copy(bidHeap.HeapNodes[j], tmp)
}

func (bidHeap *BidHeap) Push(x interface{}) {
	heapNode := x.(*HeapNode)
	bidHeap.HeapNodes = append(bidHeap.HeapNodes, heapNode)
	bidHeap.LengthNode.Amount++
}

func (bidHeap *BidHeap) Pop() interface{} {

	n := bidHeap.LengthNode.Amount
	if n == 0 {
		return nil
	}

	old := bidHeap.HeapNodes
	res := old[n-1]
	bidHeap.HeapNodes = old[0 : n-1]
	bidHeap.LengthNode.Amount--
	return res
}

func newBidHeap() (*BidHeap) {
	lengthNode := &HeapNode{
		BidName: "Length",
		Amount:  defaultLeng,
	}
	return &BidHeap{
		make([]*HeapNode, defaultLeng),
		lengthNode,
	}
}

func TestHeap() {

	bid := Bid{}
	fmt.Println(bid)

	bidHeap := newBidHeap()

	for i := 0; i < defaultLeng; i++ {
		node := &HeapNode{
			Amount: uint64(rand.Intn(20)),
		}

		if i < 10 {
			node.BidName = "time0" + strconv.Itoa(i) + "  amount:"
		} else {
			node.BidName = "time" + strconv.Itoa(i) + "  amount:"
		}
		bidHeap.HeapNodes[i] = node
	}

	heap.Init(bidHeap)

	start := time.Now()
	node := &HeapNode{
		strconv.Itoa(defaultLeng) + " : hyz",
		uint64(100000111),
	}
	heap.Push(bidHeap, node)
	end := time.Now()

	fmt.Println(end.Sub(start))

	n := defaultLeng/2 - 1
	for i := 0; i < n; i++ {
		if bidHeap.HeapNodes[i].Amount < bidHeap.HeapNodes[i*2+1].Amount || bidHeap.HeapNodes[i].Amount < bidHeap.HeapNodes[i*2+2].Amount {
			fmt.Println("wrong heap" + strconv.Itoa(i))
		}
	}

	for i := 0; i < 100; i++ {
		fmt.Println(heap.Pop(bidHeap).(*HeapNode))
	}
}
