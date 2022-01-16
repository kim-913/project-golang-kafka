# tigergraph-project

## Divided the whole pipline into four parts.

### 1. Create MYSQL table called `project`

### 2. Generate 10million random data of the following form:

⋅⋅* id: integer number within 32-bit range

⋅⋅* name: names are strings with the English character only, length ranging from 10-15 

⋅⋅* address: addresses are strings with a mixture of numbers, characters, and space, length ranging from 15-20

⋅⋅* Continent: one value from the following values {“North America”, “Asia”, “South America”, “Europe”, “Africa”, “Australia”}

Runtime: O(n)
SpaceComplexity: O(n)

#### to minimize space, instead of creating a giant array, I pushed each row of data to MYSQL individually. The runtime of generating and pushing the data to csv file is around 15min.

### 3. Create Java maven project that firstly create connection between local host with Apache Kafka.

⋅⋅* I divided the maven project into four different parts that reads the csv generated in step 1 and 2 and which belongs to a class called userModel. In this model, I designed four APIs that matches the required form for the csv.

⋅⋅* Following the step, the ProducerJavaModel class producers between local csv and kafka topic called source. 

Runtime: O(n)
SpaceComplexity: O(n)

#### 4. Golang project that reads the data inside the source topic in kafka and sort accodingly to :

 ⋅⋅* id(numerically) 
 ⋅⋅* name(alphabetically)
 ⋅⋅* continent(alphabetically)
 ⋅⋅* and produce the sorted data to three different topics(id, name, continent).
 
 #### Consumer:
  ⋅⋅* In the program, I created connections between golang and kafka by creating one consumer group that reads csv data and stores in a global variable called 
group_message. In such way, I deployed a sorting algorithm that builds a map such that take are ID, name, and Continent as keys respectively and values for the row data inside the csv. Following along, I sorted the map according to the key and append the values into differernt arrays.

#### Producer:
 ⋅⋅* As sorted in three different arrays called id_sorted, name_sorted, and continent_sorted, I appened the sorted data as message to kafka and produced into three differnt topics called id, name, and continent. 
 
 Runtime: O(nlogn)
 SpaceComplexity: O(n)

The whole pipline runs within 2GB and 4cores.

#### The optimization I've done is
 ⋅⋅* Firstly, when generating the data, to save memory, I didn't use giant arrays and append when generation is done and instead I used single values. 
 ⋅⋅* Secondly, I improved my efficiency to the golang-kafka project by not listing out the map but only the sorted arrays.
 
#### Improvements:
 ⋅⋅* If I were given more time, I would study deep about multi thread and instead of running the data seperately, I would run the whole pipline together.
