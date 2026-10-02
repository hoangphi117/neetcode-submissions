type DynamicArray struct {
	data []int
	size int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	arr := make([]int,0,capacity)
	return &DynamicArray{
		data: arr,
		size: 0,
		capacity:capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size == da.capacity {
		da.resize()
	}
	da.data = append(da.data, n)
	da.size++
}

func (da *DynamicArray) Popback() int {
	res := da.data[da.size - 1]
	da.data = da.data[:da.size-1]
	da.size--
	return res
}

func (da *DynamicArray) resize() {
	da.capacity *= 2
	newArr := make([]int, da.size, da.capacity)
	copy(newArr, da.data[:da.size])
	da.data = newArr
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
