---
title: "编程笔记：Raft 共识算法的简化实现"
description: "一份硬核的技术笔记，记录了对照 Raft 论文，用 Go 语言从零开始实现一个简化版 Raft 共识协议的心路历程与关键挑战。"
---

# 编程笔记：Raft 共识算法的简化实现

> **免责声明**: 这是一份学习笔记，旨在记录将 Raft 论文思想转化为代码的过程。此实现仅为教学目的，省略了许多生产环境所需的重要特性（如持久化、成员变更、日志压缩等），**请勿用于生产环境**。

## 背景

Raft 是一种分布式共识算法，用于在分布式系统中达成一致。它通过选举领导者、复制日志和提交日志来实现一致性。你可以在这里查看 Raft 论文：[Raft: In Search of an Understandable Consensus Algorithm](https://raft.github.io/raft.pdf)，或者这里查看 Raft 论文的[中文翻译](https://github.com/maemual/raft-zh_cn)。

## 实现

## 阶段一：理解并定义状态

实现任何算法的第一步都是精确地定义其状态。反复阅读了 Raft 论文中的 Figure 2，它像一张蓝图一样清晰地列出了所有节点需要维护的状态。可以将其翻译为 Go 的数据结构。

`raft.go`:
```go
type State string

const (
    Follower  State = "follower"
    Candidate State = "candidate"
    Leader    State = "leader"
)

// LogEntry 代表一条日志记录
type LogEntry struct {
    Term    int
    Command interface{}
}

// RaftNode 是单个 Raft 节点的核心结构
type RaftNode struct {
    mu sync.Mutex // 用于保护节点状态的互斥锁

    id    int           // 节点 ID
    peers map[int]string // 其他节点的地址

    state State // 节点当前状态 (Follower, Candidate, Leader)
    term  int   // 当前任期

    // 所有节点上都存在的持久化状态
    log      []LogEntry // 日志条目
    votedFor int        // 在当前任期内投票给的候选人 ID

    // 所有节点上都存在的易失性状态
    commitIndex int // 已知的最大的已经被提交的日志条目的索引值
    lastApplied int // 已经被应用到状态机的最高的日志条目的索引值

    // 领导者节点上存在的易失性状态 (选举后会重新初始化)
    nextIndex  map[int]int // 对于每一台服务器，发送到该服务器的下一个日志条目的索引
    matchIndex map[int]int // 对于每一台服务器，已知的已经复制到该服务器的最高日志条目的索引

    // ... 其他辅助字段，如计时器、通道等
}
```
特别注意区分了**持久化状态**（理论上需要写入稳定存储）和**易失性状态**。这个 `RaftNode` 结构体将是我们整个实现的核心。

## 阶段二：实现领导者选举

Raft 通过一种心跳机制来触发领导者选举。如果一个跟随者在一段时间（选举超时）内没有收到来自领导者的心跳，它就认为领导者可能已经下线，并开始一轮新的选举。

### 2.1 选举超时

每个节点设计了一个循环，在其中监控选举计时器。这个计时器的时长是随机的，以避免所有节点同时发起选举导致"选票瓜分"。

```go
func (n *RaftNode) runElectionTimer() {
    timeout := randomElectionTimeout()
    n.electionTimer = time.NewTimer(timeout)

    for {
        select {
        case <-n.electionTimer.C:
            // 计时器触发，开始选举
            n.startElection()
            // 重置计时器
            n.electionTimer.Reset(randomElectionTimeout())
        case <- n.heartbeatC:
             // 收到了领导者的心跳，重置计时器
            n.electionTimer.Reset(randomElectionTimeout())
        // ...
        }
    }
}
```

### 2.2 发起选举 (`startElection`)

当选举计时器触发时，节点会：
1.  **状态转换**：将自己的状态从 `Follower` 变为 `Candidate`。
2.  **增加任期**：将 `n.term` 加一。
3.  **为自己投票**：将 `n.votedFor` 设置为自己的 ID，并初始化票数为 1。
4.  **广播投票请求**：并行地向所有其他节点发送 `RequestVote` RPC。

`RequestVoteArgs` 和 `RequestVoteReply` 的结构也严格遵循论文：
```go
type RequestVoteArgs struct {
    Term         int // 候选人的任期号
    CandidateID  int // 请求选票的候选人的 ID
    LastLogIndex int // 候选人的最后日志条目的索引值
    LastLogTerm  int // 候选人最后日志条目的任期号
}

type RequestVoteReply struct {
    Term        int  // 当前任期号，以便于候选人去更新自己的任期号
    VoteGranted bool // 候选人赢得了此张选票时为 true
}
```

### 2.3 处理投票请求 (`RequestVote` RPC Handler)

当一个节点收到 `RequestVote` RPC 时，它的决策逻辑是关键：

1.  如果请求的 `Term` 小于自己的 `Term`，直接拒绝投票。
2.  如果请求的 `Term` 大于自己的 `Term`，则无条件将自己转换为 `Follower`，并更新自己的 `Term`。
3.  在任期相同的情况下，它会检查自己是否已经投过票 (`votedFor == -1` 或 `votedFor == args.CandidateID`)。
4.  **安全性检查**：最重要的一步！它必须确保候选人的日志至少和自己一样"新"。Raft 定义"新"的规则是：比较最后一条日志的任期号，任期号大的日志更新；如果任期号相同，则日志更长的更新。

```go
// isLogUpToDate 检查候选人的日志是否至少和本节点一样新
func (n *RaftNode) isLogUpToDate(candidateLastLogIndex, candidateLastLogTerm int) bool {
    myLastLogTerm := n.log[len(n.log)-1].Term
    myLastLogIndex := len(n.log) - 1
    if candidateLastLogTerm > myLastLogTerm {
        return true
    }
    if candidateLastLogTerm == myLastLogTerm && candidateLastLogIndex >= myLastLogIndex {
        return true
    }
    return false
}
```
只有通过所有这些检查，节点才会投出赞成票。

## 阶段三：实现日志复制

一旦一个候选人获得了超过半数节点的选票，它就成为新的领导者。领导者的首要职责就是管理日志复制，确保所有跟随者的日志与自己保持一致。

### 3.1 领导者的心跳 (`AppendEntries` RPC)

领导者会周期性地向所有跟随者发送 `AppendEntries` RPC。在没有新日志需要发送时，这个 RPC 就是一个空的心跳，用于告知大家"我还活着"，从而阻止跟随者发起新的选举。

`AppendEntriesArgs` 的设计是 Raft 保证日志一致性的核心：
```go
type AppendEntriesArgs struct {
    Term         int        // 领导人的任期号
    LeaderID     int        // 领导人 ID
    PrevLogIndex int        // 新日志条目紧随之前的索引值
    PrevLogTerm  int        // PrevLogIndex 条目的任期号
    Entries      []LogEntry // 准备存储的日志条目（表示心跳时为空）
    LeaderCommit int        // 领导人已经提交的日志的索引值
}
```

### 3.2 跟随者的一致性检查

当一个跟随者收到 `AppendEntries` RPC 时，它会执行一个至关重要的一致性检查：
**它会检查本地日志中 `PrevLogIndex` 位置的日志条目，看其任期号是否与 `PrevLogTerm` 匹配。**

-   **如果匹配**：说明到 `PrevLogIndex` 为止的日志是一致的。跟随者可以安全地接收并附加 `Entries` 中的新日志（如果存在冲突，会先删除本地冲突的日志及其后的所有日志）。
-   **如果不匹配**：说明日志在 `PrevLogIndex` 处出现了分歧。跟随者会拒绝这次 RPC，并返回 `false`。

### 3.3 领导者的回溯

当领导者收到一个因为日志不一致而被拒绝的 `AppendEntries` 回复时，它就知道自己为该跟随者记录的 `nextIndex` 太超前了。于是它会：
1.  将该跟随者的 `nextIndex` 减一。
2.  在下一次心跳时，重新发送 `AppendEntries` RPC，这次 RPC 会包含更早的日志条目。

这个回溯过程会一直持续，直到领导者找到与跟随者日志一致的那个点。一旦找到，领导者就可以开始向该跟随者同步后续的所有日志，最终使集群的日志达成一致。

### 3.4 提交日志

当领导者发现，对于某个日志索引 `N`，已经有超过半数的跟随者的 `matchIndex >= N`，并且 `log[N].Term == currentTerm` 时，领导者就可以认为索引 `N` 的日志已经被"提交"（committed）。它会将这个信息通过后续 `AppendEntries` RPC 中的 `LeaderCommit` 字段广播给所有跟随者，这样跟随者也可以相应地更新自己的 `commitIndex`，并将已提交的日志应用到它们各自的状态机中。

## 最终思考

从零实现 Raft 是一次非常有益的思维锻炼。它迫使我将抽象的理论具象化为代码逻辑，并在过程中处理各种边界情况。虽然这个实现是简化的，但它捕捉了 Raft 算法最核心的两个部分：如何通过选举就"谁是领导者"达成共识，以及如何通过日志复制就"操作序列"达成共识。这正是分布式系统可靠性的基石。
