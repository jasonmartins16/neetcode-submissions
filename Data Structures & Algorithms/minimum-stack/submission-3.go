type MinStack struct {
    st []int
    minst []int
}

func Constructor() MinStack {
    return MinStack{
        st: []int{},
        minst: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.st = append(this.st, val)
    if len(this.minst) == 0 || val <= this.minst[len(this.minst) - 1] {
        this.minst = append(this.minst, val)
    }
}

func (this *MinStack) Pop() {
    if len(this.st) == 0 {
        return
    }
    topval := this.st[len(this.st) - 1]
    minval := this.minst[len(this.minst) - 1]
    if topval == minval {
        this.minst = this.minst[:len(this.minst) - 1]
    }

    this.st = this.st[:len(this.st) - 1]
}

func (this *MinStack) Top() int {
    return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
    return this.minst[len(this.minst) - 1]
}
