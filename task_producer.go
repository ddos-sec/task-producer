package task_producer

// TaskProducer 定义产生任务的生产者
type TaskProducer interface {

	// Run 执行，产生生产者
	Run() error

	// Stop 提前结束生产者
	Stop() error

	// Await 在异步执行时等待生产者完成
	Await() error
}
