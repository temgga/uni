package heap

// ----- Leftist Heap-------

type LNode struct {
	key         int
	rank        int
	left, right *LNode
}

type LeftistHeap struct {
	root *LNode //корень дерева
	size int    //сколько элементов
}

// вспомогательная функция
// если узел пустой, его ранг = 0, иначе возвращаем его ранг
func lRank(n *LNode) int {
	if n == nil {
		return 0
	}
	return n.rank
}

// функция объединяет две кучи
// если одна куда пустая return вторую
func lMerge(a, b *LNode) *LNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	//выбираем наименьш корень
	if a.key > b.key {
		a, b = b, a
	}
	a.right = lMerge(a.right, b) //рекурсивный merge
	//мы засовываем b в правое поддерево а

	//балансируем, левое всегда должно быть "тяжелее"
	if lRank(a.left) < lRank(a.right) {
		a.left, a.right = a.right, a.left
	}
	a.rank = lRank(a.right) + 1 //обновляем ранк
	return a                    //возвращаем корень
}

// ф-ия достает минимум
func (h *LeftistHeap) ExtractMin() int {
	min := h.root.key //запоминаем значение
	h.root = lMerge(h.root.left, h.root.right)
	//удаление корня, просто сливаем их
	h.size-- //уменьшаем размер на 1
	return min
}

// уменьшаем значение узла
func (h *LeftistHeap) DecreaseKey(node *LNode, newKey int) {
	node.key = newKey             //меняем ключ
	h.root = lMerge(h.root, node) //заново merge'им
}

// ----- Skew Heap ------

type SNode struct {
	key         int    //значение
	left, right *SNode //дети
}

type SkewHeap struct {
	root *SNode //корень
	size int    //кол-во элементов
}

func sMerge(a, b *SNode) *SNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.key > b.key {
		a, b = b, a
	}
	a.left, a.right = sMerge(a.right, b), a.left
	//новый лефт результат мерджа
	return a
}

func (h *SkewHeap) Insert(key int) {
	node := &SNode{key: key}      //новый узелу
	h.root = sMerge(h.root, node) //мердж с одним элементом
	h.size++
}

// берем значение корня, потому что мин всегда в корне
func (h *SkewHeap) GetMin() int {
	return h.root.key
}

func (h *SkewHeap) ExtractMin() int {
	min := h.root.key
	h.root = sMerge(h.root.left, h.root.right)
	h.size--
	return min
}

func (h *SkewHeap) DecreaseKey(node *SNode, newKey int) {
	node.key = newKey
	h.root = sMerge(h.root, node)
}
