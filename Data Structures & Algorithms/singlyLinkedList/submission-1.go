type Node struct {
    val  int
    next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	current := ll.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			return -1
		}
		current = current.next
	}
	if current == nil {
        return -1
    }
	return current.val
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := &Node{
        val:  val,
        next: ll.head,
    }

    ll.head = newNode

    if ll.tail == nil {
        ll.tail = newNode
    }
}

func (ll *LinkedList) InsertTail(val int) {
	newNode := &Node{val: val, next: nil}

	if ll.head == nil {
        ll.head = newNode
        ll.tail = newNode
        return
    }

	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.head == nil {
    	return false
	}
	 if index == 0 {
        ll.head = ll.head.next

        if ll.head == nil {
            ll.tail = nil
        }

        return true
    }
	current := ll.head
	for i := 0; i < index - 1; i++ {
		if current.next == nil {
			return false
		}
		current = current.next
	}
	
    if current.next == nil {
        return false
    }

    if current.next == ll.tail {
        ll.tail = current
    }
	current.next = current.next.next
	return true
}

func (ll *LinkedList) GetValues() []int {
    arr := []int{}
    current := ll.head

    for current != nil {
        arr = append(arr, current.val)
        current = current.next
    }

    return arr
}
