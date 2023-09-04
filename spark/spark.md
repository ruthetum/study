# Apache Spark
- Fast and general purpose cluster computing system (faster than Hadoop MapReduce)
- 10x (on disk) - 100x (in memory) faster
  - Hadoop structures cause Disk I/O each time a request is made
  - The concept of Spark is to process this Disk I/O as In-memory
- Provides high level APIs in Java, Scala, Python/R
- Integration with Hadoop and its ecosystem and can read existing data

## MapReduce vs Spark
![image](https://github.com/ruthetum/study/assets/59307414/b4944255-7d13-4099-9125-c880f93f0f40)

![image](https://github.com/ruthetum/study/assets/59307414/ae6d84c8-92b1-4d68-a434-8ee9c108ca39)

Spark helps overcomes this limitation and optimises processing speed
- Hadoop MapReduce writes most of the intermediate results to **disk**, which is a slow process
- Spark achieves this by minimising disk read/write operations for intermediate results, storing these **in memory** and perform disk operations only when essential

![](https://images.velog.io/images/king3456/post/c1f7ba13-b240-4f3a-b7cc-af41734ca897/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA%202021-04-19%20%E1%84%8B%E1%85%A9%E1%84%92%E1%85%AE%209.17.10.png)

Spark supports data analysis, machine learning, graphs, streaming data, etc.

## Architecture

![](https://www.edureka.co/blog/wp-content/uploads/2018/09/Picture6-2.png)

