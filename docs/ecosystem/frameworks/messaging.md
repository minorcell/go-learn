# 消息队列客户端

> 异步解耦是现代应用架构的核心。选择消息队列需要根据业务场景权衡性能、可靠性和复杂度。

## 为什么需要消息队列？

**同步 vs 异步** 的架构差异：

- **解耦服务**：发送方无需关心接收方的可用性
- **削峰填谷**：缓解流量高峰对下游系统的冲击
- **提升用户体验**：耗时操作异步处理，快速响应
- **系统韧性**：消息持久化，避免数据丢失

**适用场景：** 订单处理、消息通知、数据同步、日志收集

## 主流消息队列对比

| 特性 | Redis | RabbitMQ | Kafka | NATS | NSQ |
|------|-------|----------|-------|------|-----|
| **性能** | 极高 | 中等 | 极高 | 高 | 高 |
| **可靠性** | 中等 | 高 | 高 | 中等 | 中等 |
| **复杂度** | 低 | 中等 | 高 | 低 | 低 |
| **持久化** | 有限 | 完整 | 完整 | 可选 | 本地 |
| **适用场景** | 缓存+简单队列 | 企业消息 | 大数据流 | 云原生 | 分布式队列 |

---

## Redis：高性能轻量级选择

**为什么选择Redis？**

- 性能极高，单实例10万+QPS
- 部署简单，运维成本低
- 丰富的数据结构支持
- 适合轻量级消息场景

**适用场景：** 简单异步任务、缓存失效通知、实时计数

::: details Redis 基础队列实现
```go
package main

import (
    "context"
    "encoding/json"
    "time"
    "github.com/go-redis/redis/v8"
)

type RedisQueue struct {
    client *redis.Client
    queue  string
}

type Message struct {
    ID        string    `json:"id"`
    Type      string    `json:"type"`
    Payload   interface{} `json:"payload"`
    Timestamp time.Time `json:"timestamp"`
}

func NewRedisQueue(addr, password, queue string) *RedisQueue {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       0,
    })
    
    return &RedisQueue{
        client: rdb,
        queue:  queue,
    }
}

// 发送消息
func (rq *RedisQueue) Send(ctx context.Context, msgType string, payload interface{}) error {
    msg := Message{
        ID:        generateID(),
        Type:      msgType,
        Payload:   payload,
        Timestamp: time.Now(),
    }
    
    data, err := json.Marshal(msg)
    if err != nil {
        return err
    }
    
    return rq.client.LPush(ctx, rq.queue, data).Err()
}

// 接收消息（阻塞）
func (rq *RedisQueue) Receive(ctx context.Context) (*Message, error) {
    result, err := rq.client.BRPop(ctx, 0, rq.queue).Result()
    if err != nil {
        return nil, err
    }
    
    var msg Message
    err = json.Unmarshal([]byte(result[1]), &msg)
    return &msg, err
}
```
:::

::: details RabbitMQ 企业级解决方案
```go
package main

import (
    "encoding/json"
    "github.com/streadway/amqp"
)

type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }
    
    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    
    return &RabbitMQ{
        conn:    conn,
        channel: ch,
    }, nil
}

func (rmq *RabbitMQ) Publish(queue string, message interface{}) error {
    body, err := json.Marshal(message)
    if err != nil {
        return err
    }
    
    return rmq.channel.Publish(
        "",    // exchange
        queue, // routing key
        false, // mandatory
        false, // immediate
        amqp.Publishing{
            ContentType:  "application/json",
            Body:         body,
            DeliveryMode: amqp.Persistent, // 持久化
        },
    )
}
```
:::

::: details Kafka 大数据流处理
```go
package main

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
    writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
    return &KafkaProducer{
        writer: &kafka.Writer{
            Addr:     kafka.TCP(brokers...),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }
}

func (kp *KafkaProducer) Send(ctx context.Context, key string, value interface{}) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    
    return kp.writer.WriteMessages(ctx, kafka.Message{
        Key:   []byte(key),
        Value: data,
    })
}
```
:::

---

## 选择建议

### 🚀 **高性能场景**
- **首选：** Redis（简单队列）、Kafka（大数据）

### 🏢 **企业级应用**
- **首选：** RabbitMQ（可靠性）

### ☁️ **云原生架构**
- **首选：** NATS（轻量级）

---

## 常见问题

**Q: 如何保证消息不丢失？**
A: 生产者确认 + 消息持久化 + 消费者手动ACK

**Q: 如何处理消息积压？**
A: 增加消费者实例、优化处理逻辑、考虑批量处理

# Kafka：分布式流式处理平台的Go实践

> Kafka是LinkedIn开源的分布式流处理平台，基于分区日志模型提供高吞吐、可持久化的消息传递。本文从工程实践角度分析Kafka的Go客户端使用、可靠性机制和生产环境最佳实践。

## 队列模型简介与使用场景

### 核心概念

Kafka采用**分布式分区日志**模型，与传统消息队列显著不同：

- **Topic**: 消息类别，类似数据库表
- **Partition**: Topic的物理分片，保证分区内消息顺序
- **Offset**: 分区内消息的唯一标识，单调递增
- **Consumer Group**: 消费者组，实现负载均衡和故障转移

### 通信模型特点

| 特性 | Kafka表现 | 工程意义 |
|------|----------|----------|
| **持久化** | 磁盘存储，可配置保留期 | 支持消息重放和审计 |
| **顺序性** | 分区内严格有序 | 适合事件流处理 |
| **吞吐量** | 百万级消息/秒 | 处理大规模数据流 |
| **可扩展性** | 水平扩展分区 | 支持大规模集群 |

### 适用场景

**✅ 强烈推荐**：
- 大数据流处理（点击流、日志聚合）
- 事件溯源架构（Event Sourcing）
- 微服务间异步通信（解耦、削峰）
- 实时数据管道（ETL、CDC）

**❌ 不推荐**：
- 简单任务队列（Redis、RabbitMQ更适合）
- 低延迟要求（<1ms）的场景
- 小规模应用（运维复杂度高）

---

## Go客户端初始化与连接

### 主流客户端对比

| 客户端 | 维护方 | 特点 | 适用场景 |
|--------|--------|------|----------|
| **segmentio/kafka-go** | Segment | 简洁、性能好 | 大多数场景首选 |
| **Shopify/sarama** | Shopify | 功能全、配置多 | 复杂业务需求 |
| **confluentinc/confluent-kafka-go** | Confluent | 官方、C binding | 追求极致性能 |

本文主要使用`segmentio/kafka-go`，它在简洁性和性能间取得了良好平衡。

### 基础连接配置

::: details Kafka客户端初始化和配置
```go
package main

import (
    "context"
    "crypto/tls"
    "time"
    
    "github.com/segmentio/kafka-go"
    "github.com/segmentio/kafka-go/sasl/plain"
    "github.com/segmentio/kafka-go/sasl/scram"
)

// 基础配置
type KafkaConfig struct {
    Brokers   []string
    Username  string
    Password  string
    Topic     string
    GroupID   string
    EnableTLS bool
}

// 生产者配置
func NewKafkaWriter(config KafkaConfig) *kafka.Writer {
    dialer := &kafka.Dialer{
        Timeout:   10 * time.Second,
        DualStack: true,
    }
    
    // SASL认证配置
    if config.Username != "" {
        mechanism := plain.Mechanism{
            Username: config.Username,
            Password: config.Password,
        }
        dialer.SASLMechanism = mechanism
    }
    
    // TLS配置
    if config.EnableTLS {
        dialer.TLS = &tls.Config{
            MinVersion: tls.VersionTLS12,
        }
    }
    
    return &kafka.Writer{
        Addr:         kafka.TCP(config.Brokers...),
        Topic:        config.Topic,
        Balancer:     &kafka.Hash{}, // 按key哈希分区
        Dialer:       dialer,
        WriteTimeout: 10 * time.Second,
        ReadTimeout:  10 * time.Second,
        
        // 批处理配置
        BatchSize:    100,               // 批次大小
        BatchTimeout: 10 * time.Millisecond, // 批次超时
        
        // 可靠性配置
        RequiredAcks: kafka.RequireAll, // 等待所有副本确认
        Async:        false,            // 同步发送
    }
}

// 消费者配置
func NewKafkaReader(config KafkaConfig) *kafka.Reader {
    dialer := &kafka.Dialer{
        Timeout:   10 * time.Second,
        DualStack: true,
    }
    
    // 认证配置（同生产者）
    if config.Username != "" {
        mechanism := plain.Mechanism{
            Username: config.Username,
            Password: config.Password,
        }
        dialer.SASLMechanism = mechanism
    }
    
    if config.EnableTLS {
        dialer.TLS = &tls.Config{
            MinVersion: tls.VersionTLS12,
        }
    }
    
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers:         config.Brokers,
        Topic:           config.Topic,
        GroupID:         config.GroupID,
        Dialer:          dialer,
        
        // 消费配置
        MinBytes:        1,              // 最小读取字节数
        MaxBytes:        10e6,           // 最大读取字节数 (10MB)
        MaxWait:         1 * time.Second, // 最大等待时间
        
        // 分区分配策略
        GroupBalancers: []kafka.GroupBalancer{
            &kafka.RoundRobinGroupBalancer{}, // 轮询分配
        },
        
        // offset管理
        StartOffset: kafka.LastOffset, // 从最新位置开始消费
    })
}

// 连接健康检查
func CheckKafkaConnection(brokers []string) error {
    conn, err := kafka.Dial("tcp", brokers[0])
    if err != nil {
        return err
    }
    defer conn.Close()
    
    // 获取集群信息
    brokers, err := conn.Brokers()
    if err != nil {
        return err
    }
    
    log.Printf("Connected to Kafka cluster with %d brokers", len(brokers))
    return nil
}
```
:::

### 环境适配配置

::: details 不同环境的配置策略
```go
// 开发环境配置
func NewDevelopmentConfig() KafkaConfig {
    return KafkaConfig{
        Brokers:   []string{"localhost:9092"},
        Topic:     "dev-events",
        GroupID:   "dev-consumer-group",
        EnableTLS: false,
    }
}

// 生产环境配置
func NewProductionConfig() KafkaConfig {
    return KafkaConfig{
        Brokers: []string{
            "kafka-1.prod.company.com:9092",
            "kafka-2.prod.company.com:9092", 
            "kafka-3.prod.company.com:9092",
        },
        Username:  os.Getenv("KAFKA_USERNAME"),
        Password:  os.Getenv("KAFKA_PASSWORD"),
        Topic:     "prod-events",
        GroupID:   "prod-consumer-group",
        EnableTLS: true,
    }
}

// 云服务配置 (如Confluent Cloud)
func NewCloudConfig() KafkaConfig {
    return KafkaConfig{
        Brokers: []string{
            "pkc-xxx.us-east-1.aws.confluent.cloud:9092",
        },
        Username:  os.Getenv("CONFLUENT_API_KEY"),
        Password:  os.Getenv("CONFLUENT_API_SECRET"),
        EnableTLS: true,
    }
}
```
:::

---

## 消息发布与消费示例

### 消息发布模式

::: details 生产者实现和消息发送
```go
package producer

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    
    "github.com/segmentio/kafka-go"
)

type EventProducer struct {
    writer *kafka.Writer
}

type Event struct {
    ID        string                 `json:"id"`
    Type      string                 `json:"type"`
    Source    string                 `json:"source"`
    Data      map[string]interface{} `json:"data"`
    Timestamp time.Time              `json:"timestamp"`
}

func NewEventProducer(config KafkaConfig) *EventProducer {
    return &EventProducer{
        writer: NewKafkaWriter(config),
    }
}

// 单条消息发送
func (p *EventProducer) SendEvent(ctx context.Context, key string, event Event) error {
    event.Timestamp = time.Now()
    
    data, err := json.Marshal(event)
    if err != nil {
        return fmt.Errorf("marshal event: %w", err)
    }
    
    message := kafka.Message{
        Key:   []byte(key),
        Value: data,
        Headers: []kafka.Header{
            {Key: "event-type", Value: []byte(event.Type)},
            {Key: "source", Value: []byte(event.Source)},
        },
    }
    
    // 同步发送，确保可靠性
    err = p.writer.WriteMessages(ctx, message)
    if err != nil {
        return fmt.Errorf("write message: %w", err)
    }
    
    return nil
}

// 批量消息发送
func (p *EventProducer) SendEventsBatch(ctx context.Context, events []Event) error {
    messages := make([]kafka.Message, len(events))
    
    for i, event := range events {
        event.Timestamp = time.Now()
        data, err := json.Marshal(event)
        if err != nil {
            return fmt.Errorf("marshal event %d: %w", i, err)
        }
        
        messages[i] = kafka.Message{
            Key:   []byte(event.ID), // 使用事件ID作为分区键
            Value: data,
            Headers: []kafka.Header{
                {Key: "event-type", Value: []byte(event.Type)},
                {Key: "batch-index", Value: []byte(fmt.Sprintf("%d", i))},
            },
        }
    }
    
    return p.writer.WriteMessages(ctx, messages...)
}

// 异步发送（高吞吐场景）
func (p *EventProducer) SendEventAsync(ctx context.Context, key string, event Event) <-chan error {
    errCh := make(chan error, 1)
    
    go func() {
        defer close(errCh)
        err := p.SendEvent(ctx, key, event)
        if err != nil {
            errCh <- err
        }
    }()
    
    return errCh
}

// 带重试的发送
func (p *EventProducer) SendEventWithRetry(ctx context.Context, key string, event Event, maxRetries int) error {
    var lastErr error
    
    for attempt := 0; attempt <= maxRetries; attempt++ {
        err := p.SendEvent(ctx, key, event)
        if err == nil {
            return nil
        }
        
        lastErr = err
        if attempt < maxRetries {
            backoff := time.Duration(attempt+1) * 100 * time.Millisecond
            select {
            case <-time.After(backoff):
                continue
            case <-ctx.Done():
                return ctx.Err()
            }
        }
    }
    
    return fmt.Errorf("failed after %d attempts: %w", maxRetries, lastErr)
}

func (p *EventProducer) Close() error {
    return p.writer.Close()
}
```
:::

### 消息消费模式

::: details 消费者实现和并发控制
```go
package consumer

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "sync"
    "time"
    
    "github.com/segmentio/kafka-go"
)

type EventConsumer struct {
    reader   *kafka.Reader
    handlers map[string]EventHandler
    workers  int
}

type EventHandler func(ctx context.Context, event Event) error

func NewEventConsumer(config KafkaConfig, workers int) *EventConsumer {
    return &EventConsumer{
        reader:   NewKafkaReader(config),
        handlers: make(map[string]EventHandler),
        workers:  workers,
    }
}

// 注册事件处理器
func (c *EventConsumer) RegisterHandler(eventType string, handler EventHandler) {
    c.handlers[eventType] = handler
}

// 单线程消费
func (c *EventConsumer) ConsumeMessages(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        
        message, err := c.reader.ReadMessage(ctx)
        if err != nil {
            log.Printf("Error reading message: %v", err)
            continue
        }
        
        if err := c.processMessage(ctx, message); err != nil {
            log.Printf("Error processing message: %v", err)
            // 根据业务需求决定是否继续或重试
        }
    }
}

// 多线程并发消费
func (c *EventConsumer) ConsumeMessagesParallel(ctx context.Context) error {
    msgChan := make(chan kafka.Message, c.workers*2)
    var wg sync.WaitGroup
    
    // 启动worker goroutines
    for i := 0; i < c.workers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            c.messageWorker(ctx, workerID, msgChan)
        }(i)
    }
    
    // 消息读取循环
    go func() {
        defer close(msgChan)
        for {
            select {
            case <-ctx.Done():
                return
            default:
            }
            
            message, err := c.reader.ReadMessage(ctx)
            if err != nil {
                log.Printf("Error reading message: %v", err)
                continue
            }
            
            select {
            case msgChan <- message:
            case <-ctx.Done():
                return
            }
        }
    }()
    
    wg.Wait()
    return ctx.Err()
}

func (c *EventConsumer) messageWorker(ctx context.Context, workerID int, msgChan <-chan kafka.Message) {
    for {
        select {
        case message, ok := <-msgChan:
            if !ok {
                return
            }
            
            start := time.Now()
            err := c.processMessage(ctx, message)
            duration := time.Since(start)
            
            if err != nil {
                log.Printf("Worker %d failed to process message: %v (duration: %v)", 
                    workerID, err, duration)
            } else {
                log.Printf("Worker %d processed message successfully (duration: %v)", 
                    workerID, duration)
            }
            
        case <-ctx.Done():
            return
        }
    }
}

func (c *EventConsumer) processMessage(ctx context.Context, message kafka.Message) error {
    var event Event
    if err := json.Unmarshal(message.Value, &event); err != nil {
        return fmt.Errorf("unmarshal message: %w", err)
    }
    
    // 查找对应的处理器
    handler, exists := c.handlers[event.Type]
    if !exists {
        log.Printf("No handler for event type: %s", event.Type)
        return nil // 忽略未知事件类型
    }
    
    // 执行处理逻辑
    if err := handler(ctx, event); err != nil {
        return fmt.Errorf("handle event %s: %w", event.Type, err)
    }
    
    return nil
}

// 手动提交offset
func (c *EventConsumer) ConsumeWithManualCommit(ctx context.Context) error {
    for {
        message, err := c.reader.FetchMessage(ctx)
        if err != nil {
            log.Printf("Error fetching message: %v", err)
            continue
        }
        
        // 处理消息
        if err := c.processMessage(ctx, message); err != nil {
            log.Printf("Error processing message: %v", err)
            // 可以选择跳过或重试
            continue
        }
        
        // 手动提交offset
        if err := c.reader.CommitMessages(ctx, message); err != nil {
            log.Printf("Error committing message: %v", err)
        }
    }
}

func (c *EventConsumer) Close() error {
    return c.reader.Close()
}
```
:::

---

## 可靠性机制

### ACK机制和幂等性

::: details 消息确认和幂等处理
```go
// 生产者可靠性配置
func NewReliableProducer(config KafkaConfig) *kafka.Writer {
    return &kafka.Writer{
        Addr:     kafka.TCP(config.Brokers...),
        Topic:    config.Topic,
        Balancer: &kafka.Hash{},
        
        // 可靠性配置
        RequiredAcks: kafka.RequireAll,    // 等待所有副本确认
        Async:       false,                // 同步发送
        
        // 重试配置
        MaxAttempts: 3,
        BatchTimeout: 100 * time.Millisecond,
        
        // 幂等性配置
        Idempotent: true,                  // 启用幂等性
    }
}

// 幂等消费处理
type IdempotentConsumer struct {
    consumer     *EventConsumer
    processedIDs sync.Map // 使用内存缓存，生产环境建议用Redis
}

func (ic *IdempotentConsumer) ProcessEvent(ctx context.Context, event Event) error {
    // 检查是否已处理
    if _, exists := ic.processedIDs.Load(event.ID); exists {
        log.Printf("Event %s already processed, skipping", event.ID)
        return nil
    }
    
    // 执行业务逻辑
    err := ic.doBusinessLogic(ctx, event)
    if err != nil {
        return err
    }
    
    // 记录处理状态
    ic.processedIDs.Store(event.ID, time.Now())
    return nil
}

func (ic *IdempotentConsumer) doBusinessLogic(ctx context.Context, event Event) error {
    // 具体业务处理逻辑
    switch event.Type {
    case "user.created":
        return ic.handleUserCreated(ctx, event)
    case "order.placed":
        return ic.handleOrderPlaced(ctx, event)
    default:
        return fmt.Errorf("unknown event type: %s", event.Type)
    }
}
```
:::

### 重试和死信队列

::: details 重试机制和错误处理
```go
package retry

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"
    
    "github.com/segmentio/kafka-go"
)

type RetryConfig struct {
    MaxRetries      int
    InitialInterval time.Duration
    MaxInterval     time.Duration
    Multiplier      float64
}

type RetryableConsumer struct {
    mainConsumer  *EventConsumer
    retryProducer *EventProducer
    dlqProducer   *EventProducer
    retryConfig   RetryConfig
}

type MessageWithMetadata struct {
    OriginalMessage kafka.Message `json:"original_message"`
    RetryCount      int           `json:"retry_count"`
    FirstFailedAt   time.Time     `json:"first_failed_at"`
    LastFailedAt    time.Time     `json:"last_failed_at"`
    ErrorMessage    string        `json:"error_message"`
}

func NewRetryableConsumer(
    mainConsumer *EventConsumer,
    retryProducer *EventProducer,
    dlqProducer *EventProducer,
    retryConfig RetryConfig,
) *RetryableConsumer {
    return &RetryableConsumer{
        mainConsumer:  mainConsumer,
        retryProducer: retryProducer,
        dlqProducer:   dlqProducer,
        retryConfig:   retryConfig,
    }
}

func (rc *RetryableConsumer) ProcessWithRetry(ctx context.Context, message kafka.Message) error {
    var metadata MessageWithMetadata
    
    // 尝试解析重试元数据
    if err := json.Unmarshal(message.Value, &metadata); err != nil {
        // 首次处理的消息
        metadata = MessageWithMetadata{
            OriginalMessage: message,
            RetryCount:      0,
            FirstFailedAt:   time.Now(),
        }
    }
    
    // 处理消息
    err := rc.mainConsumer.processMessage(ctx, metadata.OriginalMessage)
    if err == nil {
        return nil // 处理成功
    }
    
    // 处理失败，判断是否需要重试
    metadata.RetryCount++
    metadata.LastFailedAt = time.Now()
    metadata.ErrorMessage = err.Error()
    
    if metadata.RetryCount <= rc.retryConfig.MaxRetries {
        // 发送到重试队列
        return rc.sendToRetryQueue(ctx, metadata)
    } else {
        // 发送到死信队列
        return rc.sendToDLQ(ctx, metadata)
    }
}

func (rc *RetryableConsumer) sendToRetryQueue(ctx context.Context, metadata MessageWithMetadata) error {
    // 计算重试延迟
    interval := rc.calculateRetryInterval(metadata.RetryCount)
    
    // 添加延迟标识
    retryEvent := Event{
        ID:   fmt.Sprintf("retry-%s-%d", metadata.OriginalMessage.Key, metadata.RetryCount),
        Type: "retry",
        Data: map[string]interface{}{
            "metadata":  metadata,
            "retry_at": time.Now().Add(interval),
        },
    }
    
    log.Printf("Sending message to retry queue (attempt %d/%d)", 
        metadata.RetryCount, rc.retryConfig.MaxRetries)
    
    return rc.retryProducer.SendEvent(ctx, string(metadata.OriginalMessage.Key), retryEvent)
}

func (rc *RetryableConsumer) sendToDLQ(ctx context.Context, metadata MessageWithMetadata) error {
    dlqEvent := Event{
        ID:   fmt.Sprintf("dlq-%s", metadata.OriginalMessage.Key),
        Type: "dead_letter",
        Data: map[string]interface{}{
            "metadata": metadata,
            "reason":   "max_retries_exceeded",
        },
    }
    
    log.Printf("Sending message to DLQ after %d failed attempts", metadata.RetryCount)
    
    return rc.dlqProducer.SendEvent(ctx, string(metadata.OriginalMessage.Key), dlqEvent)
}

func (rc *RetryableConsumer) calculateRetryInterval(retryCount int) time.Duration {
    interval := rc.retryConfig.InitialInterval
    
    for i := 1; i < retryCount; i++ {
        interval = time.Duration(float64(interval) * rc.retryConfig.Multiplier)
        if interval > rc.retryConfig.MaxInterval {
            interval = rc.retryConfig.MaxInterval
            break
        }
    }
    
    return interval
}

// 重试队列消费者
func (rc *RetryableConsumer) ConsumeRetryQueue(ctx context.Context) error {
    // 注册重试事件处理器
    rc.mainConsumer.RegisterHandler("retry", func(ctx context.Context, event Event) error {
        metadataData, ok := event.Data["metadata"]
        if !ok {
            return fmt.Errorf("missing metadata in retry event")
        }
        
        var metadata MessageWithMetadata
        metadataBytes, _ := json.Marshal(metadataData)
        if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
            return fmt.Errorf("unmarshal retry metadata: %w", err)
        }
        
        // 检查是否到了重试时间
        retryAtData := event.Data["retry_at"]
        retryAtStr, ok := retryAtData.(string)
        if ok {
            retryAt, err := time.Parse(time.RFC3339, retryAtStr)
            if err == nil && time.Now().Before(retryAt) {
                // 还未到重试时间，重新发送到重试队列
                return rc.sendToRetryQueue(ctx, metadata)
            }
        }
        
        // 执行重试处理
        return rc.ProcessWithRetry(ctx, metadata.OriginalMessage)
    })
    
    return rc.mainConsumer.ConsumeMessages(ctx)
}
```
:::

---

## 性能与实践建议

### 性能优化配置

::: details 生产环境性能优化
```go
// 高吞吐量生产者配置
func NewHighThroughputProducer(config KafkaConfig) *kafka.Writer {
    return &kafka.Writer{
        Addr:     kafka.TCP(config.Brokers...),
        Topic:    config.Topic,
        Balancer: &kafka.Hash{},
        
        // 批处理优化
        BatchSize:    1000,                    // 增大批次大小
        BatchTimeout: 10 * time.Millisecond,  // 减少批次超时
        BatchBytes:   1048576,                 // 1MB批次大小
        
        // 异步发送
        Async: true,
        
        // 压缩配置
        Compression: kafka.Snappy, // 使用压缩减少网络传输
        
        // 连接复用
        WriteTimeout: 30 * time.Second,
        ReadTimeout:  30 * time.Second,
    }
}

// 高性能消费者配置
func NewHighPerformanceConsumer(config KafkaConfig) *kafka.Reader {
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers: config.Brokers,
        Topic:   config.Topic,
        GroupID: config.GroupID,
        
        // 批量读取优化
        MinBytes: 10e3,  // 10KB
        MaxBytes: 10e6,  // 10MB
        MaxWait:  100 * time.Millisecond,
        
        // 预取配置
        QueueCapacity: 1000, // 增加队列容量
        
        // 分区分配优化
        GroupBalancers: []kafka.GroupBalancer{
            &kafka.RoundRobinGroupBalancer{},
        },
    })
}

// 性能监控
type KafkaMetrics struct {
    MessagesSent     int64
    MessagesReceived int64
    BytesSent        int64
    BytesReceived    int64
    ProduceLatency   time.Duration
    ConsumeLatency   time.Duration
}

func (m *KafkaMetrics) RecordProduceLatency(start time.Time) {
    m.ProduceLatency = time.Since(start)
    atomic.AddInt64(&m.MessagesSent, 1)
}

func (m *KafkaMetrics) RecordConsumeLatency(start time.Time) {
    m.ConsumeLatency = time.Since(start)
    atomic.AddInt64(&m.MessagesReceived, 1)
}
```
:::

### 可观测性集成

::: details 监控和链路追踪
```go
package observability

import (
    "context"
    "time"
    
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

var (
    // Prometheus指标
    kafkaProduceTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "kafka_produce_messages_total",
            Help: "Total number of messages produced to Kafka",
        },
        []string{"topic", "status"},
    )
    
    kafkaProduceDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "kafka_produce_duration_seconds",
            Help:    "Time spent producing messages to Kafka",
            Buckets: prometheus.DefBuckets,
        },
        []string{"topic"},
    )
    
    kafkaConsumeTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "kafka_consume_messages_total",
            Help: "Total number of messages consumed from Kafka",
        },
        []string{"topic", "group", "status"},
    )
)

// 带监控的生产者
type ObservableProducer struct {
    producer *EventProducer
    tracer   trace.Tracer
}

func NewObservableProducer(producer *EventProducer) *ObservableProducer {
    return &ObservableProducer{
        producer: producer,
        tracer:   otel.Tracer("kafka-producer"),
    }
}

func (op *ObservableProducer) SendEvent(ctx context.Context, key string, event Event) error {
    // 开始链路追踪
    ctx, span := op.tracer.Start(ctx, "kafka.produce")
    defer span.End()
    
    // 记录开始时间
    start := time.Now()
    
    // 发送消息
    err := op.producer.SendEvent(ctx, key, event)
    
    // 记录指标
    duration := time.Since(start)
    kafkaProduceDuration.WithLabelValues(event.Type).Observe(duration.Seconds())
    
    status := "success"
    if err != nil {
        status = "error"
        span.RecordError(err)
    }
    kafkaProduceTotal.WithLabelValues(event.Type, status).Inc()
    
    return err
}

// 带监控的消费者
type ObservableConsumer struct {
    consumer *EventConsumer
    tracer   trace.Tracer
    topic    string
    groupID  string
}

func NewObservableConsumer(consumer *EventConsumer, topic, groupID string) *ObservableConsumer {
    return &ObservableConsumer{
        consumer: consumer,
        tracer:   otel.Tracer("kafka-consumer"),
        topic:    topic,
        groupID:  groupID,
    }
}

func (oc *ObservableConsumer) ProcessEvent(ctx context.Context, event Event) error {
    // 开始链路追踪
    ctx, span := oc.tracer.Start(ctx, "kafka.consume")
    defer span.End()
    
    // 记录开始时间
    start := time.Now()
    
    // 处理消息
    var err error
    if handler, exists := oc.consumer.handlers[event.Type]; exists {
        err = handler(ctx, event)
    }
    
    // 记录指标
    status := "success"
    if err != nil {
        status = "error"
        span.RecordError(err)
    }
    
    kafkaConsumeTotal.WithLabelValues(oc.topic, oc.groupID, status).Inc()
    
    return err
}

// 健康检查
func (oc *ObservableConsumer) HealthCheck(ctx context.Context) error {
    // 检查消费者连接状态
    stats := oc.consumer.reader.Stats()
    
    if stats.Messages == 0 && time.Since(stats.StartTime) > 5*time.Minute {
        return fmt.Errorf("no messages consumed for 5 minutes")
    }
    
    return nil
}
```
:::

### 常见问题和解决方案

**1. 消费延迟过高**
- 增加消费者并发数
- 优化消息处理逻辑
- 考虑增加分区数

**2. 消息堆积**
- 监控consumer lag
- 扩展消费者组实例
- 优化批处理大小

**3. 重复消费**
- 实现幂等性处理
- 使用唯一消息ID
- 检查offset提交逻辑

**4. 数据丢失**
- 配置合适的ACK级别
- 启用生产者重试
- 实现消费者错误处理

---

## 与其他MQ对比

### Kafka vs RabbitMQ

| 维度 | Kafka | RabbitMQ | 分析 |
|------|-------|----------|------|
| **吞吐量** | 百万级/秒 | 万级/秒 | Kafka适合大数据场景 |
| **延迟** | 毫秒级 | 微秒级 | RabbitMQ低延迟更优 |
| **持久化** | 强持久化 | 可选持久化 | Kafka天然支持消息重放 |
| **复杂度** | 高 | 中等 | RabbitMQ学习成本更低 |
| **顺序性** | 分区内有序 | 队列级有序 | 各有适用场景 |

### Kafka vs NATS

| 维度 | Kafka | NATS | 分析 |
|------|-------|------|------|
| **部署复杂度** | 高（需Zookeeper/KRaft） | 低（单二进制） | NATS更适合云原生 |
| **消息持久化** | 强持久化 | 有限持久化 | Kafka更适合事件溯源 |
| **性能** | 高吞吐 | 低延迟 | 看业务需求选择 |
| **生态** | 丰富 | 精简 | Kafka企业级工具更多 |

---

## 小结与推荐

### 选择Kafka的决策因素

**强烈推荐场景：**
- 大数据流处理（日志聚合、点击流分析）
- 事件驱动架构（Event Sourcing、CQRS）
- 微服务解耦（异步通信、削峰填谷）
- 需要消息重放能力的系统

**技术选型决策树：**

```
消息队列需求评估
├── 吞吐量 > 10万/秒 → ✅ 选择Kafka
├── 需要消息重放 → ✅ 选择Kafka  
├── 事件流处理 → ✅ 选择Kafka
├── 简单任务队列 → 考虑Redis/RabbitMQ
└── 低延迟要求 → 考虑NATS/RabbitMQ
```

### 实施建议

**生产环境最佳实践：**
1. **集群规划**：至少3个broker，奇数个分区
2. **监控体系**：监控consumer lag、磁盘使用率
3. **数据备份**：配置合适的副本因子
4. **版本管理**：使用schema registry管理消息格式

**避免常见陷阱：**
- 不要过度分区（分区数 ≈ 消费者数）
- 避免大消息（>1MB考虑拆分）
- 合理设置保留策略避免磁盘爆满
- 实现proper的错误处理和重试机制

Kafka作为分布式流处理的事实标准，在大数据和事件驱动架构中具有不可替代的地位。对于有大规模数据处理需求的系统，Kafka是最佳选择。
