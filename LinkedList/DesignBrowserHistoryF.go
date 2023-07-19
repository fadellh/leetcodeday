package linkedlist

// import "fmt"

// type Web struct {
// 	val     string
// 	forward *Web
// 	back    *Web
// }

// type BrowserHistory struct {
// 	head     *Web
// 	tail     *Web
// 	curIdx   int
// 	backStep int
// }

// func Constructor(homepage string) BrowserHistory {
// 	newHistory := &Web{
// 		val:     homepage,
// 		forward: nil,
// 		back:    nil,
// 	}
// 	return BrowserHistory{
// 		head:     newHistory,
// 		curIdx:   0,
// 		backStep: 0,
// 	}
// }

// func (this *BrowserHistory) Visit(url string) {
// 	cur := this.head

// 	idx := 0
// 	for cur != nil {
// 		// if cur.forward
// 		if idx == this.curIdx {
// 			newHistory := &Web{
// 				val:     url,
// 				forward: nil,
// 				back:    cur,
// 			}
// 			cur.forward = newHistory
// 			this.curIdx++
// 			if this.curIdx == 1 {
// 				this.backStep = 1
// 			}
// 			return
// 		}
// 		cur = cur.forward
// 		idx++
// 	}
// }

// func (this *BrowserHistory) Back(steps int) string {
// 	cur := this.head
// 	if steps == 4 || steps == 7 {
// 		fmt.Println("BACK SA", this.backStep, this.curIdx)
// 	}
// 	this.backStep += steps
// 	curIdx := this.curIdx
// 	curIdx -= steps
// 	if curIdx < 0 {
// 		curIdx = 0
// 		this.backStep = this.curIdx
// 	}

// 	idx := 0
// 	for cur != nil {
// 		if curIdx == idx {
// 			this.curIdx = idx
// 			fmt.Println("BACK ===", cur.val)
// 			return cur.val
// 		}
// 		cur = cur.forward
// 		idx++
// 	}

// 	return ""
// }

// func (this *BrowserHistory) Forward(steps int) string {
// 	cur := this.head
// 	curIdx := this.curIdx
// 	diffHistory := steps - this.backStep
// 	fmt.Println("FORW STEP", this.backStep, curIdx)

// 	idx := 0
// 	for cur != nil {
// 		if diffHistory > 0 {
// 			if (curIdx + diffHistory) == idx {
// 				this.curIdx = idx
// 				fmt.Println("===", cur.val)
// 				return cur.val
// 			}
// 		}
// 		if (curIdx + steps) == idx {
// 			this.curIdx = idx
// 			fmt.Println("===", cur.val)

// 			return cur.val
// 		}

// 		cur = cur.forward
// 		idx++
// 		if cur.forward == nil {
// 			fmt.Println("===", cur.val)

// 			return cur.val
// 		}
// 	}

// 	return ""
// }

// func (this *BrowserHistory) PrintHistory() {
// 	idx := 0
// 	cur := this.head

// 	fmt.Println("===Looping===", this.curIdx)

// 	for cur != nil {
// 		fmt.Println("===Index:", (idx))

// 		fmt.Printf("Memory address of cur: %p\n", cur)
// 		fmt.Println("val", cur.val)
// 		fmt.Printf("Memory address of cur next: %p\n", cur.forward)
// 		cur = cur.forward
// 		idx++
// 	}
// }

// func ImplemenWebHistory() {
// 	//["BrowserHistory","back","back","visit","forward","visit","visit","visit","back","visit","visit","visit","back","visit","visit","visit","visit","back","visit","visit","visit","visit","visit","visit","visit","back","forward","back","forward","visit","back","visit","visit"]
// 	//[["icpbj.com"],[1],[10],["xbepk.com"],[8],["kls.com"],["dlkwxpf.com"],["pnep.com"],[9],["rmis.com"],["bxf.com"],["dz.com"],[2],["acuqsax.com"],["dcvo.com"],["ijbg.com"],["nlott.com"],[7],["dd.com"],["vssnq.com"],["teur.com"],["pn.com"],["szi.com"],["brhldg.com"],["yulyoqv.com"],[4],[10],[8],[5],["av.com"],[3],["okr.com"],["meli.com"]]
// 	obj := Constructor("icpbj.com")
// 	obj.Back(1)
// 	obj.Back(10)
// 	obj.Visit("xbepk.com")
// 	obj.Forward(8)
// 	obj.Visit("kls.com")
// 	obj.Visit("dlkwxpf.com")
// 	obj.Visit("pnep.com")
// 	obj.Back(9)
// 	obj.Visit("rmis.com")
// 	obj.Visit("bxf.com")
// 	obj.Visit("dz.com")
// 	obj.Back(2)
// 	obj.Visit("acuqsax.com")
// 	obj.Visit("dcvo.com")
// 	obj.Visit("ijbg.com")
// 	obj.Visit("nlott.com")
// 	obj.Back(7)
// 	obj.Visit("dd.com")
// 	obj.Visit("vssnq.com")
// 	obj.Visit("teur.com")
// 	obj.Visit("pn.com")
// 	obj.Visit("szi.com")
// 	obj.Visit("brhldg.com")
// 	obj.Visit("yulyoqv.com")
// 	obj.PrintHistory()
// 	obj.Back(4)
// 	obj.Forward(10)
// 	obj.Back(8)
// 	obj.Forward(5)
// 	obj.Visit("golang.com")
// 	obj.Visit("youtube.com")
// 	obj.Back(3)
// 	obj.Visit("fdsds.com")
// 	obj.Visit("dsfee.com")

// 	// fmt.Println(url)
// 	// fmt.Println(urls)
// 	// fmt.Println(urlB)
// 	// fmt.Println(urlF)
// 	// fmt.Println(urlD)
// 	// fmt.Println(urlG)
// }

// /**
//  * Your BrowserHistory object will be instantiated and called as such:
//  * obj := Constructor(homepage);
//  * obj.Visit(url);
//  * param_2 := obj.Back(steps);
//  * param_3 := obj.Forward(steps);
//  */
