package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// kafka local port
var Address = []string{"localhost:9092"}

// consume from topic source
var topic = []string{"source"}

// create list group_message
var group_message []string

// create list with sorted id
var id_sorted []string

// create list with sorted name
var name_sorted []string

// // create list with sorted continent
var continent_sorted []string

var size = 10000000

// create map and sort according to name and address and return the list
func create_sort_string_map(group_message []string, index int) []string {
	var sorted_message []string
	var count = 0
	m := make(map[string]string)
	for i := 0; i < len(group_message); i++ {
		// fmt.Println(group_message)
		s := strings.Split(group_message[i], ",")
		// fmt.Println(s)
		if index == 2 {
			m[s[index]+string(strconv.Itoa(count))] = group_message[i]
		} else {
			m[s[index]] = group_message[i]
		}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sorted_message = append(sorted_message, m[k])
	}
	return sorted_message
}

// create map and sort according to id and return the list
func create_sort_int_map(group_message []string) []string {
	var sorted_message []string
	m := make(map[int]string)
	for i := 0; i < len(group_message); i++ {
		s := strings.Split(group_message[i], ",")
		id_key := strings.Split(s[0], ":")
		intKey, err := strconv.Atoi(id_key[1])
		if err != nil {
			fmt.Println("error")
		}
		m[intKey] = group_message[i]
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		sorted_message = append(sorted_message, m[k])
	}
	return sorted_message
}

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	group_message = clusterConsumer(wg, Address, topic, "group-1")
	// fmt.Println(group_message)
	id_sorted = create_sort_int_map(group_message)
	name_sorted = create_sort_string_map(group_message, 1)
	continent_sorted = create_sort_string_map(group_message, 2)
	fmt.Println(id_sorted)
	fmt.Println(id_sorted)
	wg.Wait()
	startProducer()
}

func clusterConsumer(wg *sync.WaitGroup, brokers, topics []string, groupId string) []string {
	var empty []string
	defer wg.Done()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange

	// init consumer
	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupId, err)
		return empty
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("%s:Error: %s\n", groupId, err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
		}
	}()

	// consume messages, watch signals
	var successes int

	// run the program n times, and store the message into group_message list
Loop:
	for i := 0; i < size; i++ {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				// fmt.Println(string(msg.Value))
				group_message = append(group_message, string(msg.Value))
				successes++
				if len(group_message) == size {
					break Loop
				}
			}
		}
	}
	// fmt.Println(group_message)
	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupId, successes)
	return group_message
}

var (
	producer sarama.SyncProducer
	brokers  = []string{"localhost:9092", "localhost:9092", "localhost:9092"}
	// produce it to following topic
	topic1 = "id"
)

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	brokers := brokers
	var err error
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Printf("init producer failed : %v \n", err)
		panic(err)
	}
	fmt.Println("producer init success..")
}

func produceMsg(msg string) {
	msgx := &sarama.ProducerMessage{}
	for i := 0; i < len(id_sorted); i++ {
		msgx.Topic = topic1
		// get the message value and send it to kafka
		msgx.Value = sarama.StringEncoder(id_sorted[i])
		partition, offset, err := producer.SendMessage(msgx)
		if err != nil {
			fmt.Printf("send msg error : %s \n", err)
		}
		fmt.Printf("msg send success, message is stored in topic[%s]/partition[%d]/offset[%d] \n", topic, partition, offset)
		// after we proccssed all data, break the loop
		if i == len(id_sorted)-1 {
			break
		}
	}
}

func startProducer() {
	t := time.Now().Unix() * 1000
	msg := fmt.Sprintf("{\"timestamp\":%d }", t)
	produceMsg(msg)
}
