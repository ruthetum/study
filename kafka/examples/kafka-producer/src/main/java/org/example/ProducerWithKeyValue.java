package org.example;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.serialization.StringSerializer;

import java.util.Properties;

public class ProducerWithKeyValue {
    private final static String TOPIC_NAME = "test";
    private final static String BOOTSTRAP_SERVERS = "localhost:29092";

    public static void main(String[] args) {

        Properties config = new Properties();
        config.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, BOOTSTRAP_SERVERS);
        config.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        config.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());

        KafkaProducer<String, String> producer = new KafkaProducer<>(config);

        ProducerRecord<String, String> record = new ProducerRecord<>(TOPIC_NAME, "Seoul", "Seoul");
        producer.send(record);

        ProducerRecord<String, String> record2 = new ProducerRecord<>(TOPIC_NAME, "Busan", "Busan");
        producer.send(record2);

        producer.flush();
        producer.close();
    }
}
