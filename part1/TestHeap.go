package part1

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)


const defaultLeng = 10000000
// the BidName which has been bidden correspond to a HeapNode to record the position of the BidName in Heap
type HeapNode struct {
	BidName string
	Amount  uint64
}

type BidHeap struct {
	HeapNodes  []*HeapNode
	LengthNode *HeapNode
}

func (bidHeap *BidHeap) Copy(dst, src *HeapNode){
	dst.Amount = src.Amount
	dst.BidName = src.BidName
}

var bidHeap *BidHeap

func (bidHeap *BidHeap) Len() int {
	return int(bidHeap.LengthNode.Amount)
}

func (bidHeap *BidHeap) Less(i, j int) bool {
	return bidHeap.HeapNodes[i].Amount > bidHeap.HeapNodes[j].Amount
}

func (bidHeap *BidHeap) Swap(i, j int) {
	tmp := &HeapNode{
	}
	bidHeap.Copy(tmp, bidHeap.HeapNodes[i])
	bidHeap.Copy(bidHeap.HeapNodes[i], bidHeap.HeapNodes[j])
	bidHeap.Copy(bidHeap.HeapNodes[j], tmp)
}

func (bidHeap *BidHeap) Push(heapNode *HeapNode) {
	bidHeap.HeapNodes = append(bidHeap.HeapNodes, heapNode)
	bidHeap.LengthNode.Amount++
}


func (bidHeap *BidHeap) Pop()  *HeapNode{

	n := bidHeap.LengthNode.Amount
	bidHeap.Swap(0, int(n-1))

	old := bidHeap.HeapNodes
	res := old[n-1]
	bidHeap.HeapNodes = old[0 : n-1]

	return res
}

func newBidHeap() (*BidHeap) {
	lengthNode := &HeapNode{
		BidName:"Length",
		Amount:defaultLeng,
	}
	return &BidHeap{
		make([]*HeapNode, defaultLeng),
		lengthNode,
	}
}

func TestHeap() {
	bidHeap:=newBidHeap()

	for i:=0; i<defaultLeng; i++{
		node := &HeapNode{
			strconv.Itoa(defaultLeng) + " : hyz",
			uint64(rand.Intn(100000)),
		}
		bidHeap.HeapNodes[i] = node
	}

	sort.Sort(bidHeap)




	start:= time.Now()
	node := &HeapNode{
		strconv.Itoa(defaultLeng) + " : hyz",
		uint64(rand.Intn(100000)),
	}
	bidHeap.Push(node)
	end:=time.Now()



	fmt.Println(end.Sub(start))


}