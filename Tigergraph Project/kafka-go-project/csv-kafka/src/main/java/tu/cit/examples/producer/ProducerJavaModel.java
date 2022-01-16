package tu.cit.examples.producer;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.util.List;
import java.util.Properties;

import java.util.Properties;

public class ProducerJavaModel {
    private static final Logger logger = LogManager.getLogger();
    public static void main(String[] args)  {

        Properties props = new Properties();
        props.put(ProducerConfig.CLIENT_ID_CONFIG,"my-app-readcsv");
        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG,"localhost:9092");
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, JsonSerializer.class);


        logger.info("Producer has been created...Start sending Student Record ");

        KafkaProducer<String,userModel> producer = new KafkaProducer<String,userModel>(props);

        ReadCSV readCSV = new ReadCSV();
        List usersList = readCSV.ReadCSVFile(); //It will return the student list
        for (Object user : usersList) {
            userModel userObject = (userModel) user;
            producer.send(new ProducerRecord<String, userModel>("source",userObject.getName(),userObject));
        }
        logger.info("Producer has sent all employee records successfully...");
        producer.close();
    }
}
