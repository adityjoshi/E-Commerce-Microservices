# Finally Coding Time 

Now it's the time too code 💅🏻


the project strucutre will initially look like this 

```
E-Commerce Project  
  |-- main.go                # Entry point of the application  
  |-- go.mod                 # Go module file  
  |-- docker-compose.yml      # Docker configuration for services  

```

# Let's create our docker file
inside docker-compose.yml write
```

services:
  postgres:
    image: postgres:latest
    container_name: postgres_db
    restart: always
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: mydb
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    container_name: zookeeper
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    container_name: kafka_broker
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
    volumes:
      - kafka_data:/var/lib/kafka/data

volumes:
  postgres_data:
  kafka_data:

```
The above will help us to setup the kafka and postgres without installing them locally in our machine.

