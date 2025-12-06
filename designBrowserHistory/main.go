type BrowserHistory struct {
	head *Node
	len  int
}

type Node struct {
	HomePage string
	Down     *Node
	Up       *Node
}

func Constructor(homepage string) BrowserHistory {
	newList := &Node{HomePage: homepage}

	return BrowserHistory{head: newList, len: 1}
}

func (this *BrowserHistory) Visit(url string) {
	// new()
	//

	curr := this.head
//	for curr.Up != nil {
//		curr = curr.Up
//	}
    curr.Up = nil
	tmp := &Node{HomePage: url, Down: curr}
	curr.Up = tmp
	curr = curr.Up
	this.head = curr

	return
}

func (this *BrowserHistory) Back(steps int) string {

	curr := this.head
	for curr.Down != nil && steps != 0 {
		curr = curr.Down
		steps--
	}
	this.head = curr

	return this.head.HomePage
}

func (this *BrowserHistory) Forward(steps int) string {

	curr := this.head
	for curr.Up != nil && steps != 0 {
		curr = curr.Up
		steps--
	}
	this.head = curr
	return this.head.HomePage
}
