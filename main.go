package go_log

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println(111)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据，需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success返回
}
