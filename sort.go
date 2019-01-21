package collections

// 冒泡排序：稳定
// 平均 O(n^2)	最好 O(n) 最坏 O(n^2)
// https://upload.wikimedia.org/wikipedia/commons/3/37/Bubble_sort_animation.gif
func BubbleSort(items []int) {
	if len(items) < 2 {
		return
	}
	n := len(items)
	index := n
	for index > 0 {
		n = index
		index = 0
		for i := 1; i < n; i++ {
			if items[i-1] > items[i] {
				items[i-1], items[i] = items[i], items[i-1]
				// 表示 i 后面的数都已经排序好了
				// 如 3 2 4 5 6 7 8 9，在 3 和 2 交换之后 无需和后面进行交换
				// 就代表着 3 之后的数都是有序的，那后面的数就不用排了
				index = i
			}
		}
	}
}

// 直接插入排序：稳定
// 平均 O(n^2) 最好 O(n) 最差 O(n^2)
// https://upload.wikimedia.org/wikipedia/commons/2/25/Insertion_sort_animation.gif
func InsertSort(items []int) {
	length := len(items)
	if length < 2 {
		return
	}

	for i := 1; i < length; i++ {
		// 如果右边的数小于左边的数 表示这个数的需要移动到左边
		if items[i] < items[i-1] {
			x := items[i]
			var j int
			// i-1 是原来的有序区 需要将 i 放入到原来的有序区中 形成一个更大的有序区
			// 2   3   4   5   1   6
			//             j   j+1      => j+1 = j
			// 2   3   4   1   5   6
			//         j   j+1          => j+1 = j
			// ...                      => j == -1
			// 1   2   3   4   5   6    => items[j+1] = x
			for j = i - 1; j >= 0 && items[j] > x; j-- {
				// 将比 x 大的数玩后面移动一位
				items[j+1] = items[j]
			}
			// 将 x 放入到正确的位置中 因为前面有 -- 操作，所以这里是 j+1
			items[j+1] = x
		}
	}
}

// 快速排序：不稳定
// 平均 O(nlogn) 最好 O(nlogn) 最坏 O(n^2)
// https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif
func QuickSort(items []int) {
	if len(items) < 2 {
		return
	}
	quickSort(items, 0, len(items)-1)
}

// 递归执行快速排序
func quickSort(items []int, l, r int) {
	// 当左右相遇的时候 递归出口
	if l < r {
		i, j := l, r
		// 轴点 用于划分左右两部分
		pivot := items[i]
		for i < j {
			// 从右往左找 找到第一个小于 轴点 的数
			for i < j && items[j] >= pivot {
				j--
			}
			// 如果找到的话 `填坑`
			if i < j {
				items[i] = items[j]
				i++
			}
			// 从左往右找 找到第一个大于 轴点 的数
			for i < j && items[i] < pivot {
				i++
			}
			// 如果找到的话 `填坑`
			if i < j {
				items[j] = items[i]
				j--
			}
		}
		items[i] = pivot         // 将轴点放到左右两边分界处
		quickSort(items, l, i-1) // 递归排序左半部分
		quickSort(items, i+1, r) // 递归排序由半部分
	}
}

// 希尔排序 不稳定
// 平均 O(nlogn) 最好 O(n) 最坏 O(n^2)
// https://upload.wikimedia.org/wikipedia/commons/d/d8/Sorting_shellsort_anim.gif
func ShellSort(items []int) {
	length := len(items)
	if length < 2 {
		return
	}

	// 增量为 gap，每次更新为 gap /= 2
	for gap := length / 2; gap > 0; gap /= 2 {
		// 假设数组为 49, 38, 65, 55, 26, 13, 27, 49, 97, 4 长度为 10
		// 第一轮的时候 gap=5 分为
		// 13 49 - 27 38 - 49 65 - 97 55 - 4 26
		// 总共需要 5 次直接插入排序 排序完后的顺序
		// 13 49 - 27 38 - 49 65 - 55 97 - 4 26

		// 第二轮的时候 gap=2 分为
		// 13 27 49 55 4 - 49 38 65 97 26
		// 总共需要 2 次直接插入排序 排序完后的顺序
		// 4 13 27 49 55 - 26 38 49 65 97

		// 第三轮的时候 gap=1 分为
		// 4 13 27 49 55 26 38 49 65 97
		// 总共需要 1 次直接插入排序 排序完后的顺序
		// 4 13 26 27 38 49 49 55 65 97
		for j := gap; j < length; j++ {
			if items[j] < items[j-gap] {
				x := items[j]
				k := j - gap
				// items[k] 为`原先在左边的数` 可理解为原来的有序区
				// x 为即将放入到有序区的值，x < items[k]，所以要把它往左边放
				for k >= 0 && items[k] > x {
					// 把 x 往左边放就意味着要把原先右边的值右移
					items[k+gap] = items[k]
					k -= gap
				}
				// 将 x 放入到正确的位置
				items[k+gap] = x
			}
		}
	}
}

// 堆排序：不稳定
// 平均 O(nlogn) 最好 O(nlogn) 最坏 O(nlogn)
// https://upload.wikimedia.org/wikipedia/commons/1/1b/Sorting_heapsort_anim.gif
func HeapSort(items []int) {
	length := len(items)
	if length < 2 {
		return
	}

	// 初始化大顶堆 从最后一个父节点开始往前调整
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(items, i, length-1)
	}

	// 进行堆排序 将最后一个叶子节点与根节点交换
	// 调整后的最后的的叶子节点即为最大值 适合求最大的 k 个数
	// 如 i > 0 => i > 3 则可以只算出最大的 3 个数
	for i := length - 1; i > 0; i-- {
		items[0], items[i] = items[i], items[0]
		adjustHeap(items, 0, i-1)
	}
}

// 调整堆 使其保持堆结构
func adjustHeap(items []int, start, end int) {
	dad := start     // 父节点
	son := dad*2 + 1 // 左子节点
	for son <= end {
		// son+1 为右子节点 如果右节点 > 左节点 son++ 取子节点中较大的一个
		if son+1 <= end && items[son] < items[son+1] {
			son++
		}
		// 如果符合大顶堆结构 直接跳出
		if items[dad] > items[son] {
			return
		}
		// 讲父节点与子节点中较大的节点进行交换
		items[dad], items[son] = items[son], items[dad]
		// 在此修正子节点堆结构
		dad = son
		son = dad*2 + 1
	}
}

// 归并排序：稳定
// 平均 O(nlogn) 最好 O(nlogn) 最差 O(nlogn)
// https://upload.wikimedia.org/wikipedia/commons/c/c5/Merge_sort_animation2.gif
func MergeSort(items []int) {
	length := len(items)
	if length < 2 {
		return
	}
	res := make([]int, length)
	mergeSort(items, 0, length-1, res)
}

// 递归执行归并排序及合并
func mergeSort(items []int, first, last int, res []int) {
	// 递归出口
	if first < last {
		mid := (first + last) / 2                // 计算出切分左右部分的边界点
		mergeSort(items, first, mid, res)        // 递归排序排序左半部分
		mergeSort(items, mid+1, last, res)       // 递归排序排序右半部分
		mergeArray(items, first, mid, last, res) // 合并左右数组
	}
}

// 将两个有序数组按序合并为一个
func mergeArray(items []int, first, mid, last int, res []int) {
	i, j := first, mid+1
	leftLen, rightLen := mid, last

	var k int

	// 左半部分 first, mid
	// 右半部分 mid+1  last
	for i <= leftLen && j <= rightLen {
		// 如果左边的值大 则将左边的值放入到 res 中
		if items[i] <= items[j] {
			res[k] = items[i]
			k++
			i++
			// 如果右边的值大 则将右边的值放入到 res 中
		} else {
			res[k] = items[j]
			k++
			j++
		}
	}

	// 右边数组遍历结束了 将左边剩下的值放入 res 中
	for i <= leftLen {
		res[k] = items[i]
		k++
		i++
	}

	// 左边数组遍历结束了 将右边剩下的值放入 res 中
	for j <= rightLen {
		res[k] = items[j]
		k++
		j++
	}

	// 这里是重点 将合并后的数据放到`原 items`中
	// 也就说放入后从 first+i 到 k（leftLen+rightLen）中是有序的
	for i := 0; i < k; i++ {
		items[first+i] = res[i]
	}
}
