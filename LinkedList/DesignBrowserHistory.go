package linkedlist

import "fmt"

type Web struct {
	val  string
	next *Web
	prev *Web
}

type BrowserHistory struct {
	cur *Web
}

func Constructor(homepage string) BrowserHistory {
	newHistory := &Web{
		val:  homepage,
		next: nil,
		prev: nil,
	}
	return BrowserHistory{
		cur: newHistory,
	}
}

func (this *BrowserHistory) Visit(url string) {
	newHistory := &Web{
		val:  url,
		next: nil,
		prev: this.cur,
	}

	this.cur.next = newHistory
	this.cur = this.cur.next
}

func (this *BrowserHistory) Back(steps int) string {
	for this.cur.prev != nil && steps > 0 {
		this.cur = this.cur.prev
		steps--
	}
	return this.cur.val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.cur.next != nil && steps > 0 {
		this.cur = this.cur.next
		steps--
	}
	return this.cur.val
}

func (this *BrowserHistory) PrintHistory() {
	idx := 0
	cur := this.cur

	// fmt.Println("===Looping===", this.curIdx)

	for cur != nil {
		fmt.Println("===Index:", (idx))

		fmt.Printf("Memory address of cur: %p\n", cur)
		fmt.Println("val", cur.val)
		// fmt.Printf("Memory address of cur next: %p\n", cur.forward)
		cur = cur.next
		idx++
	}
}

func ImplemenWebHistory() {
	//["BrowserHistory","back","back","visit","forward","visit","visit","visit","back","visit","visit","visit","back","visit","visit","visit","visit","back","visit","visit","visit","visit","visit","visit","visit","back","forward","back","forward","visit","back","visit","visit"]
	//[["icpbj.com"],[1],[10],["xbepk.com"],[8],["kls.com"],["dlkwxpf.com"],["pnep.com"],[9],["rmis.com"],["bxf.com"],["dz.com"],[2],["acuqsax.com"],["dcvo.com"],["ijbg.com"],["nlott.com"],[7],["dd.com"],["vssnq.com"],["teur.com"],["pn.com"],["szi.com"],["brhldg.com"],["yulyoqv.com"],[4],[10],[8],[5],["av.com"],[3],["okr.com"],["meli.com"]]
	obj := Constructor("icpbj.com")
	obj.Back(1)
	obj.Back(10)
	obj.Visit("xbepk.com")
	obj.Forward(8)
	obj.Visit("kls.com")
	obj.Visit("dlkwxpf.com")
	obj.Visit("pnep.com")
	obj.Back(9)
	obj.Visit("rmis.com")
	obj.Visit("bxf.com")
	obj.Visit("dz.com")
	obj.Back(2)
	obj.Visit("acuqsax.com")
	obj.Visit("dcvo.com")
	obj.Visit("ijbg.com")
	obj.Visit("nlott.com")
	obj.Back(7)
	obj.Visit("dd.com")
	obj.Visit("vssnq.com")
	obj.Visit("teur.com")
	obj.Visit("pn.com")
	obj.Visit("szi.com")
	obj.Visit("brhldg.com")
	obj.Visit("yulyoqv.com")
	obj.PrintHistory()
	obj.Back(4)
	obj.Forward(10)
	obj.Back(8)
	obj.Forward(5)
	obj.Visit("golang.com")
	obj.Visit("youtube.com")
	obj.Back(3)
	obj.Visit("fdsds.com")
	obj.Visit("dsfee.com")

	// fmt.Println(url)
	// fmt.Println(urls)
	// fmt.Println(urlB)
	// fmt.Println(urlF)
	// fmt.Println(urlD)
	// fmt.Println(urlG)
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
