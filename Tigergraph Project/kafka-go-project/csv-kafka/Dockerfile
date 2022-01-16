FROM maven:3.5.4-jdk-8-alpine as maven
COPY ./pom.xml ./pom.xml
COPY ./src ./src
RUN mvn dependency:go-offline -B
RUN mvn package
FROM openjdk:8u171-jre-alpine
WORKDIR /kafka-go-project
COPY --from=maven target/csv-kafka-1.0-SNAPSHOT.jar /kafka-go-project/javaProject.jar
CMD ["java", "-jar", "./javaProject.jar"]