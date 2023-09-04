# Hadoop
## Apache Hadoop

![](https://techvidvan.com/tutorials/wp-content/uploads/sites/2/2019/11/Hadoop-Ecosystem.jpg)

- Hadoop Common
- Hadoop Distributed File System (HDFS)
- Hadoop YARN
- Hadoop MapReduce

## HDFS
Hadoop Distributed File System (based on Google File System, GFS)
- Serves as the distributed file system for most tools in the Hadoop ecosystem
- Scalability for large data sets
- Reliability to cope with hardware failures

HDFS good for
- Large files

Not good for
- Lots of small files
- Low latency access (because Disk I/O)

### HDFS Architecture
> Master/Slave design
> - Master node
>    - Single NameNode for managing metadata
> - Slave node
>    - Multiple DataNode for storing data
> - Other
>    - Secondary NameNode as a backup

![](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-hdfs/images/hdfsarchitecture.png)

NameNode
- keeps the metadata, the name, location and directory

DataNode
- provides storage for blocks of data

### HDFS Blocks
> Replication of Blocks for fault tolerance


![](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-hdfs/images/hdfsdatanodes.png)

![](https://www.softwaretestinghelp.com/wp-content/qa/uploads/2019/12/Block-Replication.jpg)

HDFS files are divided into blocks
- It's the basic unit of read/write
- Default block size is 128MB
- Hence makes HDFS good for large files, not good for small files

HDFS blocks are replicated multiple times
- One block stored at multiple location, also at different racks (usually 3 times)
- This makes HDFS storage fault tolerance and faster to read

## MapReduce
> Data processing model designed to process large amounts of data in a distributed/parallel computing environment
> 
> When large data comes in, the data is divided into blocks of a certain size and a Map Task and a Reduce Task are performed for each block.

![](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FbqdJru%2FbtqVyUYM4qo%2FAAXd4ZPEQawzek9nT5uGpk%2Fimg.png)

Simple programming paradigm for the Hadoop ecosystem

Traditional parallel programming requires expertise of different parallel programming paradigms

The Map Task and the Reduce Task use the Key-Value structure for input and output. Map refers to an operation of grouping processed data in the form of (Key, Value)
- The Map returns data in the form of a Key, Value in the form of a List.
- The Reduce removes and merges data with duplicate key values from data processed with Map, and extracts the desired data.