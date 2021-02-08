FROM ubuntu:18.04
MAINTAINER suhwanggyu
COPY main .
RUN ["mkdir", "/workspace"]
RUN apt-get update && apt-get install -y locales
RUN apt-get install -y openjdk-8-jre
RUN apt-get install -y openjdk-8-jdk 
RUN echo "jdk 8  install completed"
RUN locale-gen ko_KR.UTF-8
ENV LC_ALL ko_KR.UTF-8
CMD ["/main"]
ENV PORT 3001
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV PATH $JAVA_HOME/bin:$PATH
EXPOSE 3001
