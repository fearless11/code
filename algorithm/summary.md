
<!-- TOC -->

- [数据结构与算法](#数据结构与算法)
    - [术语](#术语)
    - [数据结构](#数据结构)
        - [数组](#数组)
        - [集合](#集合)
        - [有序数组](#有序数组)
        - [链表](#链表)
        - [散列表](#散列表)
        - [栈](#栈)
        - [队列](#队列)
        - [二叉树](#二叉树)
        - [平衡二叉树](#平衡二叉树)
        - [堆](#堆)
        - [B树](#b树)
        - [红黑树](#红黑树)
        - [图](#图)
    - [算法](#算法)
        - [查找](#查找)
            - [线性查找](#线性查找)
            - [二分查找](#二分查找)
            - [快速选择](#快速选择)
        - [排序](#排序)
            - [冒泡排序](#冒泡排序)
            - [选择排序](#选择排序)
            - [插入排序](#插入排序)
            - [快速排序](#快速排序)
            - [希尔排序](#希尔排序)
            - [归并排序](#归并排序)
            - [堆排序](#堆排序)
        - [其他](#其他)
            - [动态规划](#动态规划)
    - [总结](#总结)
        - [时间复杂度](#时间复杂度)
        - [空间复杂度](#空间复杂度)
        - [空间换时间](#空间换时间)
        - [递归](#递归)
        - [分而治之](#分而治之)

<!-- /TOC -->
# 数据结构与算法

- 研究数据以什么形式存储，如何用一种流程方法更高效的利用内存和获得更快的执行速度。

## 术语

- 数据结构：数据以什么结构存储到内存，为了更好的操作（读取、查找、插入、删除）。

- 算法：解决某个问题的一套流程。

- 时间复杂度：完成一次操作所需要的步数。

- 空间复杂度：完成一次操作所需要的内存。

- 大O计法：一个函数的增长率的上限。在此是随着数据量的增加所操作步数的趋势。

- O(1)：不管数据量如何变化，算法的步数恒定。

- 数据结构有哪些？
  
   `数组、有序数组、集合、散列表、栈、队列
   链表、双向链表、树、二叉树、堆、B树、红黑树、2-3-4树、图`

- 算法有哪些？
  
   ```bash
   查找：线性查找、二分查找、快速选择
   遍历：广度优先搜索、深度优先搜索、Dijkstra算法最短路径
   排序：冒泡排序、选择排序、插入排序、快速排序、希尔排序、归并排序、堆排序
   动态规划
   ```

- 递归：自身调用自身。用于无法预估计算深度的问题。


## 数据结构

### 数组

- 内存：分配的内存地址是连续的格子
- 分析
  
  ```bash
  读取：知道内存地址，直接取。1次
  查找：最坏情况遍历数组所有格子没找到。N次
  插入：最坏情况数组开头插入，N次移动，一次插入。N+1次
  删除：最坏情况数组开头删除，一次删除，N-1次移动。N次
  ```

  | 数组        | 读取 | 查找 | 插入 | 删除 |
  | ----------- | ---- | ---- | ---- | ---- |
  | 操作步数/次 | 1    | N    | N+1  | N    |
  | 最坏时间复杂度  | O(1) | O(N) | O(N) | O(N) |

### 集合
- 内存：基于数组内存地址是连续格子。
- 特点：存储的数据不允许重复
- 分析

  ```bash
  读取：知道内存地址，直接取。1次
  查找：最坏情况遍历集合所有格子没找到。N次
  插入：比较是否存在后头部插入。N次比较，N次移动，一次插入。2N+1
  删除：最坏情况集合开头删除，一次删除，N-1次移动。N次
  ```

  | 数组           | 读取 | 查找 | 插入 | 删除 |
  | -------------- | ---- | ---- | ---- |   ---- |
  | 操作步数/次    | 1    | N    | 2N+1 |   N    |
  | 最坏时间复杂度 | O(1) | O(N) | O(N) | O(N)   |

### 有序数组
- 内存：分配的内存地址是连续的格子
- 特点：存取的数据是有序的
- 分析

  ```bash
  读取：知道内存地址，直接取。1次
  插入：比较大小后插入。N次查找，N次移动，一次插入。2N+1
  删除：最坏情况开头删除，一次删除，N-1次移动。N次
  线性查找：最坏情况遍历有序数组没找到，N次。
  二分查找：数据有序，一分为二。每次查找后排出一半元素。logN次。 
  （如：4个元素是2*2，所以4以2为底的对数 log4=2 ）
  ```

  | 有序数组       | 读取 | 线性查找 | 二分查找        | 插入 | 删除 |
  | -------------- | ---- | -------- | --------------- | ---- | ---- |
  | 操作步数/次    | 1    | N        | logN（2为底数） | 2N+1 | N    |
  | 最坏时间复杂度 | O(1) | O(N)     | O(logN)         | O(N) | O(N) |

- 场景

  ```bash
  查找很多，插入相对教少的情况。
  ```

### 链表
- 内存：分配的内存地址可以是不连续的格子
- 特点：结点 = 1个保存数据 + 1个保存下个结点的内存地址；头指针 —> 指向第一个结点
- 分析

  ```bash
  读取：最坏情况读取链表最后一个值，N次找到内存地址。N次
  查找：最坏情况合遍历链表没找到。N次
  插入：最坏情况链表末尾插入，N次查找内存地址，一次插入结点。N+1次
  删除：最坏情况链表末尾删除，N次查找内存地址，1次删除。N+1次
   ```
   
  | 链表           | 读取 | 查找 | 插入 | 删除 |
  | -------------- | ---- | ---- | ---- |   ---- |
  | 操作步数/次    | N    | N    | N+1  | N+1  |
  | 最坏时间复杂度 | O(N) | O(N) | O(N) | O(N)   |

- 场景

   ```bash
   高效遍历单个列表并删除其中多个元素。
   
   问题：删除电子邮件列表中无效的格式地址？
   考虑：
   1. 不管用数组还是链表，要检查每个元素都需要N步
   2. 删除时，数组需要花额外N步左移后面的数据，而链表只用1步。假设10000个数据，最坏情况下数组1步删除9999步左移，而链表时1步删除不移动。
   ```
   
### 散列表
- 理解：为数据先设置编号，编号存放到一起，但编号对应的数据存放到其他地方。
- 内存：连续内存地址存编码键key，非连续的存值value
- 特点
  ```bash
  格子：分配的key的连续内存空间长度
  散列：将字符串转化为数字串的过程，为key编号
  冲突：经过散列函数后数字串相同
  负载因子：数据量/格子数
  理想负载因子为0.7，计算机科学家研究的黄金法则，每增加7个元素就增加10个格子存key
  
  影响散列表效率因素：
  1. 负载因子。 数据量少，格子大，冲突少，效率高
  2. 散列函数。决定数据是否分散存放到格子，越分散，冲  突少，效率高

  解决冲突方案：
  分离链接：存放value的是一个数组。
  ```
- 分析
  
  ```bash
  没有冲突情况下，读取、查找、插入、删除都是1
  有冲突情况下，查找和删除取决于value的存储结构
  ```
  | 没有冲突时的散列表 | 读取 | 查找 | 插入 | 删 除 |
  | -------------------- | ---- | ---- |   ---- | ---- |
  | 操作步数/次          | 1    | 1    | 1      | 1    |
  | 最坏时间复杂度       | O(1) | O(1) | O(1)   | O(1) |

- 场景
  
  ```bash
  一般编程语言自带： 映射、字典、关联数组等
  检查一个数据存在性。
  ```
  
### 栈
- 内存：操作受限的数组或链表
- 特点

  ```bash
  栈顶：栈的末尾
  栈底：栈的开头
  出栈：读取数据
  压栈：插入数据
  对数据操作：读取、插入、删除都在栈顶
  LIFO：后进先出，最后入栈的元素会最先出栈
  ```
- 场景

  ```bash
  用于处理临时数据。
  递归调用核心思路。

  问题：检测一行代码里写的括号，包括圆括号、花括号、方括号对不对？
  思考：
  1. 如果读到的字符不是任意字符则忽略，继续下一个
  2. 如果读到左括号，则入栈
  3. 如果读到右括号，比较栈顶元素分析
     a. 如果栈里没有任何元素，则第2类语法错误（没有左括号有右括号 1+2]*3 ）
     b. 如果栈里有元素但不匹配，则第3类语法错误（有左括号但右括号不匹配[1+2}*3 ）
     c. 如果栈顶是匹配的左括号，则正常将其弹出，继续第1步
  4. 如果一行代码读完，栈里还有元素，则第1类语法错误（有左括号没有右括号[{1+2} ）
  ```

### 队列
- 内存：操作受限的数组或链表
- 特点

  ```bash
  入队：只能在末尾插入数据
  出队：只能在读取开头的数据
  只能移除开头的数据

  FIFO：先进先出
  ```
- 分析

  ```bash
  双向链表：
  结点 = 1个保存数据 + 1个保存上节点的内存地址 + 1个保存下节点的内存地址
  ```
  | 双向链表的队列 | 读取 |  插入 | 删除 |
  | -------------------- | ---- |   ---- | ---- |
  | 操作步数/次          | 1     | 1      | 1    |
  | 最坏时间复杂度       | O(1)  | O(1)   | O(1) |
    
- 场景
  
  ```bash
  处理临时数据
  打印机的作业设置
  ```
  
### 二叉树
- 内存：基于链表
- 特点
  ```bash
  结点 = 一个保存数据值的格子 + 一个存放左结点内存地址的格子 + 一个存放右结点内存地址的格子
  优点：有序数组优势在于二分查找，链表优势在于数据插入和删除。但有序数组插入较慢，链表查找较慢。故有二叉树，既保持有序，又快速查找、插入和删除。
  
  二叉树需遵守的规定：
  1. 每个结点的子结点数量可为0，1，2
  2. 如果有两个子结点，则其中一个子结点的值必须小于父结点，另一个子结点必须大于父结点
  ```
- 分析
  
  ```bash
  查找：每一次操作把结点一分为二，可排除一半结点 logN次
  插入：先查找位置后一步插入 logN+1次
  删除：先查找后少量额外的步骤去处理悬空子节点 logN+次
  ```
    
  | 二叉树         | 读取 | 查找    | 插入    | 删除    |
  | -------------- | ---- | ------- | ------- | ------- |
  | 操作步数/次    | N    | logN    | logN+1  | logN+   |
  | 最坏时间复杂度 | O(N) | O(logN) | O(logN) | O(logN) |
  
  ```bash
  前序遍历
  中序遍历
    1. 此结点有左子结点，则在左子结点上调用自身
    2. 访问此结点
    3. 此结点有右子结点，则在右子结点上调用自身
  后序遍历
  ```

  ```bash
  删除
  1. 如果删除结点没有子结点，则直接删除
  2. 如果删除结点有一个子结点，删除后子结点替换
  3. 如果删除结点有两个子结点，则将该结点替换成其后继结点。后继结点是被删除结点下子结点中大于删除结点最小的那个。如果后继节点有右结点，则作为后继结点的父结点的右结点。
  ```
- 场景

  ```bash
  存储修改有序数据。

  书目维护应用：
  1. 书名依照字母打印（有序，遍历）
  2. 持续更新书目（插入&删除）
  3. 搜索书名（查找）
  ```

### 平衡二叉树
  
- 特点
  ```bash
  目的是为了避免一般二叉树查找最坏的情况，只有左子树。
  
  规定：
  1. 每个结点的子结点数量可为0，1，2
  2. 如果有两个子结点，则其中一个子结点的值必须小于父结点，另一个子结点必须大于父结点
  3. 空树或者每个结点左右子树的高度差最多为1
  ```

### 堆
- 特点
  ```bash
  堆性质：
  1. 每个结点的值大于等于（或小于等于）其每个子结点的值
  2. 堆属于完全二叉树

  大顶堆： 每个结点大于等于其每个子结点的值
  小顶堆： 每个结点大于等于其每个子结点的值
  ```

### B树
- 特点
  ```bash
  解决数据过多时，二叉树存储数据后深度过深的问题
  性质：
  多个左右子结点  
  ```

### 红黑树
- 特点
  ```bash
  一种自平衡二叉查找树
  ```

### 图
- 内存：很多方式，其一是用散列表实现
- 特点
  
  ```bash
  关系型数据的数据结构
  ```
- 分析
  ```bash
  广度优先搜索：需要使用队列。需要有出队，还有访问
  深度优先搜索
  ```
## 算法

### 查找

#### 线性查找
- 定义：遍历逐个格子查找
- 分析
  ```golang
   // 时间复杂度O(N)
   func linearSearch() {
   	num := 7
   	numbers := [5]int{1, 4, 5, 6, 7}
   
   	for i := 0; i < len(numbers); i++ {
   		if num == numbers[i] {
   			fmt.Printf("find the number %v", num)
   			break
   		}
   	}
   }
  ```
#### 二分查找
- 定义：针对有序的数组进行一分为二的查找，每次操作后可排除一半的数据
- 步骤
  ```bash
  条件：有序数据
  
  1. 三个指针指向开头、中间、结尾
  2. 如果查找值小于中间指针值，则结尾指针等于中间指针减一，同时重新计算中间指针
  3. 如果查找值大于中间指针值，则开头指针等于中间指针加一，同时重新计算中间指针
  4. 重复1,2,3直到结束
  ```
- 分析
 ```golang
 // 关键： 三个额外的指针，分别是向开头、结尾、中间
 
 // 时间复杂度O(logN)
 // 空间复杂度O(1)
func binarySearch() {
	num := 5
	numbers := [5]int{1, 3, 4, 5, 7}

	// 上下边界
	lowerBound := 0
	upperBound := len(numbers) - 1
	for lowerBound <= upperBound {
		// 取中间元素
		midpoint := (lowerBound + upperBound) / 2
		midvalue := numbers[midpoint]
		if num < midvalue {
			upperBound = midpoint - 1
		} else if num > midvalue {
			lowerBound = midpoint + 1
		} else if num == midvalue {
			fmt.Printf("the position numbers[%v] find the number  %v\n", midpoint, midvalue)
			break
		}
	}
}
 ```

#### 快速选择
- 场景
  ```bash
  问题：一个无序的数组，不需要排序，找出里面第10小的值？
  思考：
  方案一：先将数组排序，然后到对应格子去找。即使用快排时间复杂度O(NlogN)
  方案二: 结合快排的分区+二分查找思想，先取数据一个值作为轴，将数据小于轴的放左边，大于轴的放右边。看要选择的数据比值大还是小，这样可以剔除掉一半元素。下一次分区只需要操作上一次分出的一半区域数据。操作步数：N+(N/2)+(N/4)+(N/8)+...+2大于为2N步,故时间复杂度O(N)
  ```

### 排序
#### 冒泡排序
 
- 理解：每次把最大的数冒到右边直到排序结束，一次轮回：多次比较&多次交换。

  两个指针，指向相邻两个位置，两两比较后确保交换大的在后，两个指针右移直到最大数保存在最后。循环前面的操作，判断排序好可以用趟数也可以通过一次轮回没有交换元素确定。

- 步骤
  ```bash
  # 每次轮回后都会冒一个最大的数到末尾
  # 1、2、3步为一个轮回
  # 考虑：多少个轮回 或者 一个轮回没有交换数据作为结束
  1. 对于数组执行两个指针，指向开头的两个元素
  2. 比较大小，左边大于右边就互换位置，否则什么都不做
  3. 将两个指针右移一格，重复1和2直到数组末尾，第一个最大的数已经到冒泡到数组末尾
  4. 重复1,2,3步直到无需交换。
  ```
- 代码
  ```golang
  // 关键：比较 & 交换
  // 最坏情况：一个长度为5的逆序数组，比较4+3+2+1=10次，交换4+3+2+1=10次。步数为20次，约为5*5=25次
  // 时间复杂度：O(N^2)

  func sortBubble() {
  	numbers := []int{3, 6, 7, 2, 4}
  	unsortedIndex := len(numbers) - 1
  	// 一个轮回中没有发生任何交换则数组已经排序完成
  	sorted := false
  
  	for sorted != true {
  		sorted = true
  		// 比较交换：前后两个值大小，将大的方后面
  		for i := 0; i < unsortedIndex; i++ {
  			if numbers[i] > numbers[i+1] {
  				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
  				sorted = false
  			}
  		}
  		unsortedIndex = unsortedIndex - 1
  	}
  
  	fmt.Println("bubble sort", numbers)
  }
  ```
  
#### 选择排序

- 理解：每次选择最小的值放到左边直到结束。 一次轮回：多次比较 && 一次交换
  
  取出左边的值，与其余值比较，最小的索引保存下来，然后取最小值索引的数据放到左边。依次遍历结束。

- 步骤
  ```bash
  思路：每一次将检测的最小值移动到检测的起点，直到数组比较完
  第一次：从左到右遍历一次，找出最小值的索引，将其第一个元素交换
  轮回：按照第一次检查，每次与本次检测的起点值交换
  ```
- 代码
  ```golang
  // 比较冒泡，优点减少了交换的次数
  // 最坏情况：一个长度为5的逆序数组，比较4+3+2+1=10次，交换4次。步数为14次，约为(5*5)/2=25次。
  // 操作次数：N^2/2 
  // 时间复杂度： O(N^2)

  func sortSelect() {
  	numbers := []int{3, 6, 7, 2, 4}
  
  	for i := 0; i < len(numbers); i++ {
  		// 默认每次起点值为最小值
  		var lowerNumberIndex = i
  		// 比较：起点后面的值与最小值比较
  		for j := i + 1; j < len(numbers); j++ {
  			if numbers[j] < numbers[lowerNumberIndex] {
  				lowerNumberIndex = j
  			}
  		}
  
  		// 交换：最小值与本次检测起点交换
  		if lowerNumberIndex != i {
  			var temp = numbers[i]
  			numbers[i] = numbers[lowerNumberIndex]
  			numbers[lowerNumberIndex] = temp
  		}
  	}
  
  	fmt.Println("select sort", numbers)
  } 
  ```
  
#### 插入排序

- 理解：每次取出一个值，然后与左边数据比较找到合适的位置。

  取数，比较，平移，插入。从取第2个元素到临时值，与左边比较，如果左边数据大则左边数据右移，直到左边没数据或左边值小时在当前位置插入临时值。

- 步骤
  ```bash
  1. 移出：在第一轮里，暂时将索引值移走，用一个临时变量来保存它。使得索引留下一个空隙。
  2. 比较平移：空隙左侧的值与临时变量比较，如果空隙左侧的值大于临时变量的值，则左边的值右移一格。随着右移，空隙会左移，如果遇到比临时变量小的值，或者空隙已经到了数组的最左端就结束平移。
  3. 插入：临时值插入到当前空隙。
  4. 重复1-3步，直到数组完全有序
    
  操作步数：
  比较：在最坏的情况，数组完全逆序时，临时值要与左侧所有值比较，则1+2+3+4+…+(N-1)大约为N^2/2
  平移：完全逆序时，有多少次比较就有多少次平移。N^2/2
  移出：N-1次移出临时值
  插入：N-1次插入临时值
  所以： N^2+2N-2步
  
  时间复杂度： O(N^2)
  ```
- 代码
  ```golang
  // 时间复杂度: O(N^2)
  func sortInsert() {
  	numbers := []int{3, 6, 7, 2, 4}
  
  	for i := 1; i < len(numbers); i++ {
  		// 一次轮回
  		// 移出
  		position := i
  		temp := numbers[i]
  		// 比较 : 与左边值比较，左边大则左边数据右移，左边小则当前位置插入临时值，或者左边没有数据插入临时值。
  		for position > 0 && numbers[position-1] > temp {
  			// 平移
  			numbers[position] = numbers[position-1]
  			position = position - 1
  		}
  		// 插入
  		numbers[position] = temp
  	}
  }
  ```

#### 快速排序

- 理解：基准情形，分而治之的思想。

  基准情形：右边值的指针小于等于左边值的指针。
  先找一个值，判断值合适的位置存放；再对值左右两边的分区数据进行同样操作。分别找值以及值合适的位置存放，一直达到基准情形结束。

- 步骤
  ```bash
  关键：轴、左指针、右指针

  1. 分区：数组中随便选择一个值，以此为轴，将比它小的值放到左边，比它大的放到右边。以数组最右边的值为轴分析
  2. 左指针逐个格子向右移动，当遇到大于或等于轴的值时，就停下来
  3. 右指针逐个格子向左移动，当遇到小于或等于轴的值时，就停下来
  4. 左右指针都停下，就交换左右指针的值
  5. 重复2，3，4步，直到两指针重合，或左指针移动到右指针的右边
  6. 将轴与左指针指向的值交换位置
  ```

- 代码
  
  ```golang
  // 时间复杂度：
  //   最好 O(NlogN) 平均 O(NlogN)  最坏 O(N^2)
  func sortQuick(array []int, leftIndex int, rightIndex int) {
  	// 基准情形
  	if rightIndex <= leftIndex {
  		return
  	}
  	// 分而治之
  	// 一次分区后轴的位置
  	pivotPosition := sortPartition(array, leftIndex, rightIndex)
  	// 左侧
  	sortQuick(array, leftIndex, pivotPosition-1)
  	// 右侧
  	sortQuick(array, pivotPosition+1, rightIndex)
  }
  
  func sortPartition(array []int, leftPointer int, rightPointer int) int {
  	pivotPosition := rightPointer
  	pivot := array[pivotPosition]
  
  	rightPointer = rightPointer - 1
  	for {
      // 左边比较：直到碰到大于的值
  		for array[leftPointer] < pivot {
  			leftPointer = leftPointer + 1
  		}
      // 右边比较：直到碰到小于的值
  		for array[rightPointer] > pivot {
  			rightPointer = rightPointer - 1
  		}
  
  		if leftPointer > rightPointer {
  			break
  		} else {
  			// 交换左右指针的值
  			array[leftPointer], array[rightPointer] = array[rightPointer], array[leftPointer]
  		}
  	}
  
  	// 值放到合适位置
  	array[leftPointer], array[pivotPosition] = array[pivotPosition], array[leftPointer]
  	return leftPointer
  }
  ```

#### 希尔排序
- 步骤
- 分析
- 代码
#### 归并排序
- 步骤
- 分析
- 代码
#### 堆排序
- 步骤
- 分析
- 代码
### 其他
- 步骤
- 分析
- 代码
#### 动态规划
- 定义
  
  ```bash
  动态规划：把一个复杂的问题转化为一个分阶段逐步递推的过程，从简单的初始状态一步一步递推，最终到复杂的问题的最优解。
  ```
- 分析
- 代码

## 总结

### 时间复杂度

### 空间复杂度

### 空间换时间

- 场景
  ```bash
  问题：检查数组中是否有重复元素？
  思考：
  方案一：嵌套两个for，时间复杂度O(N^2)
  方案二：用数组值作为映射中的key，不存在key对应的值为1，存在则是重复元素
  ```

- 代码
  ```golang
  // 时间复杂度：O(N^2)
  // 空间复杂度: 0
  func hasDuplicateValue1() {
  	array := []int{1, 3, 5, 7, 9, 3}
  
  	for i := 0; i < len(array); i++ {
  		for j := 0; j < len(array); j++ {
  			// 自身不与自身比较
  			if i != j && array[i] == array[j] {
  				fmt.Println("has duplicate value", array[i])
  				return
  			}
  		}
  	}
  }
  
  // 时间复杂度： O(N)
  // 空间复杂度： O(N)
  func hasDuplicateValue2() {
  	array := []int{1, 3, 5, 7, 9, 3}
  	mapArray := make(map[int]bool)
  
  	for i := 0; i < len(array); i++ {
  		if _, exist := (mapArray[array[i]]); !exist {
  			mapArray[array[i]] = true
  		} else {
  			fmt.Println("has duplicate value", array[i])
  		}
  	}
  }
  ```

### 递归

- 关键
  ```bash
  基准情形：什么条件跳出递归结束。
  适用：需要重复调用自身，无法预估深度的问题。
  优点：改变算法的实现方式；提高算法效率
  问题：栈溢出。尾递归。
  ```
- 分析
  ```bash
  由问题倒推，最终状态的前一步状态是什么，前前一步状态是什么，是否相同。
  遍历文件系统
  ```

### 分而治之
  ```bash
  挑选一个作为标准，然后一分为二操作
  ```