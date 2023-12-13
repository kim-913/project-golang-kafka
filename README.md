# Golang-kafka project

#### User Guide

### Before using this project, you need to install Python, Golan, Java, and Kafka to your local computer. Specific instruction please follow google.

**Step1**, you need to downloa mysql and create schema called project-tigergraph and create a table called project
```
CREATE TABLE `project`
(`id` INTEGER DEFAULT NULL,
`name` VARCHAR(15) DEFAULT NULL,
`address` VARCHAR(20) DEFAULT NULL,
`Continent` VARCHAR(20) DEFAULT NULL);
```
**Step2**, run csvPython.py inside pythonGenerate directory. Install mysql package and import mysql.connector. Remember to change host, user, and password to your local PC setup. In such way, we can generat a 10million rows of data insid a csv file tha can be exported through MYSQL command ``` SELECT * FROM project ````.
<br />
**Step3**, create java maven project called csv-kafka. In this folder, the csv data generated above in step 2 is stored in the data directory which allows us the read from. Here, we need buiild indepedencies inside the pom.xml file and run maven-lifecyle-clean that automatically install the dependencies to the project.
<br />
**Step4**, create a kafka topic using by firstly startine zookeepar and kafka, and build a topic according to your need. Here, I created a topic called source.
<br />
**Step5**, Run ProducerJavaModel.java inside csv-kafka/src/main an store the csv data to kafka topic called source.
<br />
**Step6**, create kafka topic where we need to store the sorted data into.
<br />
**Step7**, run go-kafka project. In this go-project, we firstly need to install "github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster" by running command of go install "github.com/Shopify/sarama". In such way, we can connect our local go project local kafka. After this, you can simply run go-kafka/main/main.go. Remember that the topic name inside the consume group needs to match the topic we store the data in, here it's "source". You can also create different topic name inside kafka and switch the name of topic insid the producer group according to the topic you want the data to produce to. 
<br />
**Conlusion**: The sorted data will be correctly produced in topics called id numerically, name alphabetically, and continent alphabetically

## Project overview
##### Divided the whole pipline into four parts.

### 1. Create MYSQL table called `project`

### 2. Generate 10million random data of the following form:

* id: integer number within 32-bit range

* name: names are strings with the English character only, length ranging from 10-15 

* address: addresses are strings with a mixture of numbers, characters, and space, length ranging from 15-20

* Continent: one value from the following values {“North America”, “Asia”, “South America”, “Europe”, “Africa”, “Australia”}

Runtime: O(n)
SpaceComplexity: O(n)

#### to minimize space, instead of creating a giant array, I pushed each row of data to MYSQL individually. The runtime of generating and pushing the data to csv file is around 15min.

### 3. Create Java maven project that firstly create connection between local host with Apache Kafka.

* I divided the maven project into four different parts that reads the csv generated in step 1 and 2 and which belongs to a class called userModel. In this model, I designed four APIs that matches the required form for the csv.

* Following the step, the ProducerJavaModel class producers between local csv and kafka topic called source. 

Runtime: O(n)
SpaceComplexity: O(n)

### 4. Golang project that reads the data inside the source topic in kafka and sort accodingly to :

* id(numerically) 
* name(alphabetically)
* continent(alphabetically)
* and produce the sorted data to three different topics(id, name, continent).
 
 #### Consumer:
* In the program, I created connections between golang and kafka by creating one consumer group that reads csv data and stores in a global variable called 
group_message. In such way, I deployed a sorting algorithm that builds a map such that take are ID, name, and Continent as keys respectively and values for the row data inside the csv. Following along, I sorted the map according to the key and append the values into differernt arrays.

#### Producer:
* As sorted in three different arrays called id_sorted, name_sorted, and continent_sorted, I appened the sorted data as message to kafka and produced into three differnt topics called id, name, and continent. 
 
 Runtime: O(nlogn)
 SpaceComplexity: O(n)

The whole pipline runs within 2GB and 4cores.

#### The optimization I've done is
* Firstly, when generating the data, to save memory, I didn't use giant arrays and append when generation is done and instead I used single values. 
* Secondly, I improved my efficiency to the golang-kafka project by not listing out the map but only the sorted arrays.
 
#### Improvements:
* If I were given more time, I would study deep about multi thread and instead of running the data seperately, I would run the whole pipline together.
