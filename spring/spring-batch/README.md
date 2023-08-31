# Spring Batch
> https://github.com/ruthetum/my-spring/tree/main/f-spring-batch

## Overview
<details>
<summary>더보기</summary>
<div markdown="1">

### 배치 프로그램
- 정해진 시간에 일괄적으로 작업을 처리하는 프로그램 (대체로 대용량 데이터를 처리)
- 서비스를 운영하는 관점에서 주기적으로 작업을 처리하기 위해 배치 프로그램 사용

#### 필요한 상황
1. 필요한 데이터를 모아서 처리해야할 때
    - ex. 월별 거래 명세서 생성
2. 일부러 지연시켜 처리할 때
    - ex. 주문한 상품을 바로 배송 처리하지 않고, 일정 시간 뒤 처리
3. 자원을 효율적으로 활용해야할 때
    - ex. 트래픽이 적은 시간 대에 서버 리소스를 활용

#### 데이터 처리 배치 프로그램
1. 각 서비스의 데이터를 데이터 웨어하우스에 저장할 때 = ETL(Extract Transform Load)
2. 아마존에서 연관 상품을 추천하는 데이터 모델을 만들 때
3. 유저 리텐션, 엑티브 상태 등 마케팅에 참고할 데이터 지표를 집계할 때
    - 유저 리텐션 : 시간이 지날수록 얼마나 많은 유저가 제품으로 다시 돌아오는지를 측정한 것

#### 서비스 배치 프로그램
1. 메세지, 이메일, 푸시 등을 발송할 때
2. 데이터를 마이그레이션할 때
3. 실패한 트랜잭션을 재처리할 때
4. 쿠폰, 포인트 등이 만료되었을 때 소진시키는 처리를 할 때
5. 월말 또는 월초에 특징 데이터를 생성할 때 (ex. 월별 거래 명세서)

</div>
</details>

## Spring Batch 기본 및 구조

<details>
<summary>더보기</summary>
<div markdown="1">

### 기본 용어

![spring batch 도메인 언어](https://user-images.githubusercontent.com/59307414/153305364-3af076aa-ca0d-4922-869c-f278be2d2c86.png)

- JoLauncher : Job을 실행시키는 컴포넌트
- Job : 배치작업
- JobRepository : Job 실행과 Job, Step을 저장
- Step : 배치 작업의 단계
- ItemReader, ItemProcesser, ItemWriter : 데이터를 읽고 처리하고 쓰는 구성

### 아키텍처

![아키텍처](https://user-images.githubusercontent.com/59307414/153305443-eb35c56c-d277-454e-850b-a3175c2f4f25.png)

- Application Layer
    - 사용자(=우리)의 코드와 구성
    - 비즈니스, 서비스 로직
    - Core, Infrastructure를 이용해 배치의 기능을 생성

- Core Layer
    - 배치 작업을 시작하고 제어하는데 필수적인 클래스
    - Job, Step, JobLauncher를 포함

- Infrastructure Layer
    - 외부와 상호작용
    - ItemReader, ItemProcesser, ItemWriter를 포함

### Job

![job](https://user-images.githubusercontent.com/59307414/153305485-b878e66b-3a69-49a6-b1cb-22c666c83eb6.png)

- 전체 배치 프로세스를 캡슐화한 도메인
- Step의 순서를 정의
- JobParameters를 받음

- Ex.
    ```java
    @Bean
    public Job footballJob() {
        return this.jobBuilderFactory.get("footballJob")
                            .start(playerLoad())            // step의 이름
                            .next(gameload())               // step의 이름
                            .next(playerSummarization())    // step의 이름
                            .build();
    }
    ```

### Step

![step](https://user-images.githubusercontent.com/59307414/153305530-0caf1493-8111-4931-ad41-724520f66cd5.png)

- 작업 처리의 단위
- Chunk 기반 스텝, Tasklet 스탭 2가지로 나뉨
    - Chunk 기반 스텝을 많이 사용
    - Tasklet 스탭은 하나의 트랜잭션 내에서 작동하고, 단순한 처리를 할 때 사용

> Chunk 기반 스텝
> ![chuck](https://user-images.githubusercontent.com/59307414/153305687-0c7a3769-c505-4651-b9ad-a904099fa8c0.png)
> - chunk 기반으로 하나의 트랜잭션에서 데이터를 처리
> - commitInterval만큼 데이터를 읽고 트랜잭션 경계 내에서 chunkSize만큼 write 진행
    >    - chunkSize : 한 트랙잭션에서 쓸 아이템의 갯수
    >    - commitInterval : reader가 한 번에 읽을 아이템의 갯수
>    - chunkSize >= commitInterval 하지만 보통 같게 맞춰서 사용하는 것이 좋음

- Ex. Chunk 기반
    ```java
    @Bean
    public Job sampleJob(JobRepository jobRepository, Step sampleStep) {
        return this.jobBuilderFactory.get("sampleJob")
                .repository(jobRepository)
                    start(sampleStep)
                    .build();
    }

    @Bean
    publuc Step sampleStep(PlatformTransactionManager transactionManager) {
        return this.stepBuilderFactory.get("sampleStep")
                .transactionManager(transactionManager)
                .<String, String>chunk(10)
                .reader(itemReader())
                .writer(itemWriter())
                .build();
    }
    ```

- Ex. TaskletStep
```java
@Bean
public Step sampleTaskletStep() {
    return this.stepBuilderFactory.get("sampleTaskletStep")
                .tasklet(myTasklet())                       
                .build();
}
```
- Tasklet 구현체를 설정. 내부에 단순한 읽기, 쓰기, 처리 로직을 모두 넣음
- RepeatStatus(반복상태)를 설정 (RepeatStatus.FINISHED)


</div>
</details>

## Spring Batch 추가 적용

<details>
<summary>더보기</summary>
<div markdown="1">

### JobParameterValidator
- 만약 추가 파라미터로 날짜를 입력해주는 경우
```java
@Slf4j
@Configuration
@AllArgsConstructor
public class AdvancedJobConfig {

    private final JobBuilderFactory jobBuilderFactory;
    private final StepBuilderFactory stepBuilderFactory;

    @Bean
    public Job advancedJob(Step advancedStep) {
        return jobBuilderFactory.get("advancedJob")
                .incrementer(new RunIdIncrementer())
                .start(advancedStep)
                .build();
    }

    @JobScope
    @Bean
    public Step advancedStep(Tasklet advancedTasklet) {
        return stepBuilderFactory.get("advancedStep")
                .tasklet(advancedTasklet)
                .build();
    }

    @StepScope
    @Bean
    public Tasklet advancedTasklet(@Value("#{jobParameters['targetDate']}") String targetDate) {
        return ((contribution, chunkContext) -> {
            LocalDate localDate = LocalDate.parse(targetDate);
            log.info("LocalDate : " + localDate);
            
            // 만약 날짜 형식이 올바르지 않다면?
            
            return RepeatStatus.FINISHED;
        });
    }
}
```
- 만약 `targetDate`로 받은 날짜 형식이 올바르지 않다면 step이 진행되는 상황에서 뒤늦게 exception이 발생
- 따라서 작업이 시작하기 전에 미리 validation을 할 수 있다면 효율적일 것이다.
- 이렇게 parameter에 대한 validation을 진행할 수 있는 게 `JobParameterValidator`

```java
# job/validator/LocalDateParameterValidator.java
@AllArgsConstructor
public class LocalDateParameterValidator implements JobParametersValidator {

    private String parameterName;

    @Override
    public void validate(JobParameters parameters) throws JobParametersInvalidException {
        String localDate = parameters.getString(parameterName);

        if (!StringUtils.hasText(localDate)) {
            throw new JobParametersInvalidException(parameterName + "가 빈 문자열이거나 존재하지 않습니다.");
        }

        try {
            LocalDate.parse(localDate);
        } catch (DateTimeParseException e) {
            throw new JobParametersInvalidException(parameterName + "의 날짜 형식이 올바르지 않습니다.");
        }
    }
}

# job/AdvancedJobConfig.java
public class AdvancedJobConfig {
    ...

    @Bean
    public Job advancedJob(Step advancedStep) {
        return jobBuilderFactory.get("advancedJob")
                .incrementer(new RunIdIncrementer())
                .validator(new LocalDateParameterValidator("targetDate"))
                .start(advancedStep)
                .build();
    }
}
```
- `validator` 설정을 통해 parameter에 대한 validation을 사전에 진행할 수 있다.

### JobExecutionListener
- 배치 작업의 상태에 따라 로직 처리가 필요한 경우
    - ex. 배치 작업이 실패하는 경우 관리자에게 이메일이나 sms 알림을 제공해야 하는 경우
- `JobExecutionListener`를 사용하자

```java
# job/AdvancedJobConfig.java
public class AdvancedJobConfig {
    
    ...
    
    @Bean
    public Job advancedJob(
            JobExecutionListener jobExecutionListener,
            Step advancedStep
    ) {
        return jobBuilderFactory.get("advancedJob")
                .incrementer(new RunIdIncrementer())
                .validator(new LocalDateParameterValidator("targetDate"))
                .listener(jobExecutionListener)
                .start(advancedStep)
                .build();
    }

    @JobScope
    @Bean
    public JobExecutionListener jobExecutionListener() {
        return new JobExecutionListener() {
            @Override
            public void beforeJob(JobExecution jobExecution) {
                log.info("[JobExecutionListenerBeforeJob] JobExecution is " + jobExecution.getStatus());
            }

            @Override
            public void afterJob(JobExecution jobExecution) {
                if (jobExecution.getStatus() == BatchStatus.FAILED) {
                    log.error("[JobExecutionListenerAfterJob] JobExecution is FAILED!!");
                    // 배치 작업이 실패했을 때 로직을 처리할 수 있다. (ex. 이메일 전송)
                }
            }
        };
    }
    
    ...
}
```

### StepExecutionListener
- `JobExecutionListener`와 동일, step 단위로 확인 가능
```java
# job/AdvancedJobConfig.java

@StepScope
@Bean
public StepExecutionListener stepExecutionListener() {
    return new StepExecutionListener() {
        @Override
        public void beforeStep(StepExecution stepExecution) {
            log.info("[StepExecutionListenerBeforeStep] StepExecution is " + stepExecution.getStatus());
        }

        @Override
        public ExitStatus afterStep(StepExecution stepExecution) {
            log.info("[StepExecutionListenerAfterStep] StepExecution is " + stepExecution.getStatus());
            return stepExecution.getExitStatus();
        }
    }
}
```

### FlatFileItemReader, ItemProcessAdapter, FlatFileItemWriter
- `FlatFileItemReader` : 파일을 읽게 해주는 ItemReader
    - chunk 기반으로 아이템들을 읽을 수 있다
    - cf.
        - https://docs.spring.io/spring-batch/docs/current/reference/html/index-single.html#flatFileItemReader
        - https://sky-h-kim.tistory.com/38

- `ItemProcessAdapter`
    - `Example 1` 처럼 바로 ItemProcessor를 적용할 수도 있고, `Example 2`처럼 별도의 Adapter를 만들어서 적용할 수도 있다.
    - Adapter를 사용하는 경우 조금 더 코드가 간단해진다.

```java
# Example 1
@JobScope
@Bean
public Step flatFileStep(
        FlatFileItemReader<PlayerDto> playerFlatFileItemReader,
        ItemProcessor<PlayerDto, PlayerSalaryDto> playerSalaryItemProcessor
    ) {
        return stepBuilderFactory.get("flatFileStep")
            .<PlayerDto, PlayerSalaryDto>chunk(5)
            .reader(playerFlatFileItemReader)
            .processor(playerSalaryItemProcessor)
            .writer(new ItemWriter<>() {
                @Override
                public void write(List<? extends PlayerSalaryDto> items) throws Exception {
                    items.forEach(System.out::println);
                }
            })
            .build();
}

@StepScope
@Bean
public ItemProcessor<PlayerDto, PlayerSalaryDto> playerSalaryItemProcessor(PlayerSalaryService playerSalaryService) {
    return new ItemProcessor<PlayerDto, PlayerSalaryDto>() {
        @Override
        public PlayerSalaryDto process(PlayerDto item) throws Exception {
            return playerSalaryService.calSalary(item);
        }
    };
}

# Example 2
@JobScope
@Bean
public Step flatFileStep(
        FlatFileItemReader<PlayerDto> playerFlatFileItemReader,
        ItemProcessorAdapter<PlayerDto, PlayerSalaryDto> playerSalaryItemProcessorAdapter
    ) {
        return stepBuilderFactory.get("flatFileStep")
            .<PlayerDto, PlayerSalaryDto>chunk(5)
            .reader(playerFlatFileItemReader)
            .processor(playerSalaryItemProcessorAdapter)
            .writer(new ItemWriter<>() {
                @Override
                public void write(List<? extends PlayerSalaryDto> items) throws Exception {
                    items.forEach(System.out::println);
                }
            })
            .build();
}

@StepScope
@Bean
public ItemProcessorAdapter<PlayerDto, PlayerSalaryDto> playerSalaryItemProcessorAdapter(PlayerSalaryService playerSalaryService) {
        ItemProcessorAdapter<PlayerDto, PlayerSalaryDto> adapter = new ItemProcessorAdapter<>();
        adapter.setTargetObject(playerSalaryService);
        adapter.setTargetMethod("calSalary");
        return adapter;
}
```

- `FlatFileItemWriter`

```java
@StepScope
@Bean
public FlatFileItemWriter<PlayerSalaryDto> playerFlatFileItemWriter() throws IOException {
    BeanWrapperFieldExtractor<PlayerSalaryDto> fieldExtractor = new BeanWrapperFieldExtractor<>();
    fieldExtractor.setNames(new String[]{"ID", "firstName", "lastName", "salary"});
    fieldExtractor.afterPropertiesSet();

    DelimitedLineAggregator<PlayerSalaryDto> lineAggregator = new DelimitedLineAggregator<>();
    lineAggregator.setDelimiter("\t");
    lineAggregator.setFieldExtractor(fieldExtractor);

    // 기존 파일 덮어쓰기
    new File("src/main/resources/sample/player-salary.txt").createNewFile();
    FileSystemResource resource = new FileSystemResource("src/main/resources/sample/player-salary.txt");

    return new FlatFileItemWriterBuilder<PlayerSalaryDto>()
            .name("playerFlatFileItemWriter")
            .resource(resource)
            .lineAggregator(lineAggregator)
            .build();
}
```

</div>
</details>

## Spring Batch 병렬 처리

<details>
<summary>더보기</summary>
<div markdown="1">

### Spring Batch에서 병렬 처리를 하는 방법 4가지
1. Multi-threaded Step (single process)
2. Parallel Steps (single process)
3. Remote Chunking of Step (multi process)
4. Partitioning (single or multi process)

- cf. https://docs.spring.io/spring-batch/docs/current/reference/html/index-single.html#multithreadedStep

### 1. Multi Threaded Step
```java
# job/parallel/MuitiThreadStepJobConfig.java
@JobScope
@Bean
public Step multiThreadStep(
        FlatFileItemReader<AmountDto> amountFileItemReader,
        ItemProcessor<AmountDto, AmountDto> amountFileItemProcessor,
        FlatFileItemWriter<AmountDto> amountFileItemWriter,
        TaskExecutor multiThreadStepTaskExecutor
) {
    return stepBuilderFactory.get("multiThreadStep")
            .<AmountDto, AmountDto>chunk(10)
            .reader(amountFileItemReader)
            .processor(amountFileItemProcessor)
            .writer(amountFileItemWriter)
            .taskExecutor(multiThreadStepTaskExecutor)
            .build();
}

@Bean
public TaskExecutor multiThreadStepTaskExecutor() {
    SimpleAsyncTaskExecutor taskExecutor = new SimpleAsyncTaskExecutor("spring-batch-task-executor");
    return taskExecutor;
}
```
- `TaskExecutor`를 이용해서 멀티 스레드 작업을 진행
- 순서가 보장되지 않고 자원에 대해 락이 걸려있으면 성능이 향상되지 않을 수 있음
- 자원 점유나 순서 보장과 관해서 자유로운 상황에서 성능을 개선해야 될 경우 사용 가능

### 2. Parallel Steps
- Step 여러 개를 동시에 실행
    - `Multi Threaded Step`은 청크 단위로 작업
- Step 자체를 하나의 스레드가 실행

```java
# job/parallel/ParallelStepJobConfig.java
public class ParallelStepJobConfig {

    private final JobBuilderFactory jobBuilderFactory;
    private final StepBuilderFactory stepBuilderFactory;

    @Bean
    public Job parallelJob(Flow splitFlow) {
        return jobBuilderFactory.get("parallelJob")
                .incrementer(new RunIdIncrementer())
                .start(splitFlow)
                .build()
                .build();
    }

    @Bean
    public Flow splitFlow(
            TaskExecutor multiThreadStepTaskExecutor,
            Flow flowAmountFileStep,
            Flow flowAnotherStep
    ) {
        return new FlowBuilder<SimpleFlow>("splitFlow")
                .split(multiThreadStepTaskExecutor)
                .add(flowAmountFileStep, flowAnotherStep)
                .build();
    }

    @Bean
    public Flow flowAmountFileStep(Step amountFileStep) {
        return new FlowBuilder<SimpleFlow>("flowAmountFileStep")
                .start(amountFileStep)
                .end();
    }

    @Bean
    public Step amountFileStep(
            FlatFileItemReader<AmountDto> amountFileItemReader,
            ItemProcessor<AmountDto, AmountDto> amountFileItemProcessor,
            FlatFileItemWriter<AmountDto> amountFileItemWriter
    ) {
        return stepBuilderFactory.get("multiThreadStep")
                .<AmountDto, AmountDto>chunk(10)
                .reader(amountFileItemReader)
                .processor(amountFileItemProcessor)
                .writer(amountFileItemWriter)
                .build();
    }

    @Bean
    public Flow flowAnotherStep(Step anotherStep) {
        return new FlowBuilder<SimpleFlow>("anotherStep")
                .start(anotherStep)
                .end();
    }

    @Bean
    public Step anotherStep() {
        return stepBuilderFactory.get("anotherStep")
                .tasklet(((contribution, chunkContext) -> {
                    Thread.sleep(500);
                    System.out.println("Another Step Completed. Thread = " + Thread.currentThread().getName());
                    return RepeatStatus.FINISHED;
                }))
                .build();
    }
}
```

### 3. Remote Chunking of Step
![remote chunking](https://docs.spring.io/spring-batch/docs/current/reference/html/images/remote-chunking.png)
- step을 다수의 process로 나누어서 작업
- worker에 channel을 통해 전달

### 4. Partitioning
![partitioning](https://docs.spring.io/spring-batch/docs/current/reference/html/images/partitioning-overview.png)
- 단일 또는 멀티 프로세스에서 실행 가능
- manager 스텝에서 만든 파티션 단위로 작업 실행

</div>
</details>

## Spring Batch 테스트

<details>
<summary>더보기</summary>
<div markdown="1">

### Mock 활용해서 서비스 테스트 코드 작성하기

```java
# test/java/com/example/springbatch/core/service/PlayerSalaryServiceTest.java
public class PlayerSalaryServiceTest {

    private PlayerSalaryService playerSalaryService;

    @BeforeEach
    public void setup() {
        playerSalaryService = new PlayerSalaryService();
    }

    @Test
    public void calSalary() {
        // given
        Year mockYear = mock(Year.class);
        when(mockYear.getValue()).thenReturn(2022);
        mockStatic(Year.class).when(Year::now).thenReturn(mockYear);

        PlayerDto mockPlayer = mock(PlayerDto.class);
        when(mockPlayer.getBirthYear()).thenReturn(1980);

        // when
        PlayerSalaryDto result = playerSalaryService.calSalary(mockPlayer);

        // then
        Assertions.assertEquals(result.getSalary(), 4200000);
    }
}
```

- Mock을 활용하기 위해 dependency 추가
    - `testImplementation 'org.mockito:mockito-inline:3.8.0'`
- mock, mockStatic을 활용해서 서비스 로직 테스트

### `AssertFile`로 파일 테스트하기

```java
# test/java/com/example/springbatch/job/FlatFileJobConfigTest.java
@SpringBootTest
@SpringBatchTest
@ExtendWith(SpringExtension.class)
@ActiveProfiles("test")
@ContextConfiguration(classes = {FlatFileJobConfig.class, BatchTestConfig.class, PlayerSalaryService.class})
public class FlatFileJobConfigTest {

    @Autowired
    private JobLauncherTestUtils jobLauncherTestUtils;

    @Test
    public void success() throws Exception {
        // given

        // when
        JobExecution execution = jobLauncherTestUtils.launchJob();

        // then
        Assertions.assertEquals(execution.getExitStatus(), ExitStatus.COMPLETED);
        AssertFile.assertFileEquals(
                new FileSystemResource("src/main/resources/sample/player-salary.txt").getFile(),
                new FileSystemResource("src/main/resources/sample/succeed-player-salary.txt").getFile()
        );
    }
}
```

- `AssertFile.assertFileEquals`를 통해 파일을 비교할 수 있다.

</div>
</details>

### Reference
- Spring batch docs : https://docs.spring.io/spring-batch/docs/current/reference/html/