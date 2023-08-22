package org.example;

import org.apache.kafka.clients.producer.*;
import org.apache.kafka.common.serialization.StringSerializer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Properties;

public class ProducerWithAsyncCallback {
    private final static String TOPIC_NAME = "test";
    private final static String BOOTSTRAP_SERVERS = "localhost:29092";

    public static void main(String[] args) {

        Properties configs = new Properties();
        configs.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, BOOTSTRAP_SERVERS);
        configs.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        configs.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());

        KafkaProducer<String, String> producer = new KafkaProducer<>(configs);

        ProducerRecord<String, String> record = new ProducerRecord<>(TOPIC_NAME, "Pangyo", "Pangyo");
        producer.send(record, new ProducerCallback());

        producer.flush();
        producer.close();
    }

    public static class ProducerCallback implements Callback {
        private final static Logger logger = LoggerFactory.getLogger(ProducerCallback.class);

        @Override
        public void onCompletion(RecordMetadata recordMetadata, Exception e) {
            if (e != null)
                logger.error(e.getMessage(), e);
            else
                logger.info(recordMetadata.toString());
        }
    }
}
