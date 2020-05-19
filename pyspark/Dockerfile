FROM jupyter/pyspark-notebook

USER root

# PySpark graphframe package variables
ENV GRAPHFRAMES_VERSION=0.8.0 \
  SPARK_VERSION=2.4 \
  SCALA_VERSION=2.11

# Graphframes full package name
ENV GRAPHFRAMES_PACKAGE=${GRAPHFRAMES_VERSION}-spark${SPARK_VERSION}-s_${SCALA_VERSION}

# Install PySpark graphframes jars
RUN ${SPARK_HOME}/bin/pyspark --packages graphframes:graphframes:${GRAPHFRAMES_PACKAGE}
RUN wget http://dl.bintray.com/spark-packages/maven/graphframes/graphframes/${GRAPHFRAMES_PACKAGE}/graphframes-${GRAPHFRAMES_PACKAGE}.jar -qO ${SPARK_HOME}/jars/graphframes.jar

USER $NB_UID

# Install python graphframes package
RUN pip install graphframes
