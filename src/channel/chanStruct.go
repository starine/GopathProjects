package _chan

/*type hchan struct {
	qcount 		uint 			//当前队列中剩余的元素个数
	dataqsize 	uint			//环形队列长度，即可以存放的元素个数 cap（）
	buf 		unsafe.Pointer 	//环形队列指针
	elemsize 	uint16 			//每个元素的大小
	closed 		uint32 			//标识关闭状态
	elemtype 	*_type 			//元素类型
	sendx 		uint 			//队列下标，指示元素写入时存放到队列中的位置
	recvx 		uint 			//队列下标，指示下一个被读取的元素在队列中的位置
	sendxq 		waitq 			//等待写消息的协程队列
	recvxq 		waitq 			//等待读消息的协程队列
	lock 		mutex  			//互斥所，chan不允许并发读写
}*/
