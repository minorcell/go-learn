# Mermaid图表示例

本页面演示了如何使用script标签方式嵌入Mermaid图表，避免HTML属性中的特殊字符问题。

## 流程图示例

<script type="text/plain" id="flowchart-example">
graph TD
    A[开始] --> B{输入验证}
    B -->|有效| C[处理数据]
    B -->|无效| D[显示错误]
    C --> E[计算结果]
    E --> F[显示结果]
    D --> G[重新输入]
    F --> H[结束]
    G --> B
    
    style A fill:#e1f5fe,stroke:#01579b
    style B fill:#fff3e0,stroke:#e65100
    style C fill:#e8f5e8,stroke:#2e7d32
    style D fill:#ffebee,stroke:#c62828
    style E fill:#f3e5f5,stroke:#7b1fa2
    style F fill:#e0f2f1,stroke:#00695c
    style G fill:#fff8e1,stroke:#f57f17
    style H fill:#fce4ec,stroke:#ad1457
</script>

<MermaidDiagram id="flowchart-example" />

## 序列图示例

<script type="text/plain" id="sequence-diagram">
sequenceDiagram
    participant U as 用户
    participant UI as 界面
    participant P as 解析器
    participant C as 计算器
    
    U->>UI: 输入 "2 + 3 * 4"
    UI->>P: 解析表达式
    
    Note over P: 词法分析
    P->>P: tokenize()
    
    Note over P: 语法分析
    P->>P: parseExpression()
    P->>P: parseTerm()
    P->>P: parseFactor()
    
    P->>C: 计算结果
    C->>UI: 返回 14
    UI->>U: 显示结果
    
    Note over U,C: 计算完成
</script>

<MermaidDiagram id="sequence-diagram" />

## 状态图示例

<script type="text/plain" id="state-diagram">
stateDiagram-v2
    [*] --> 等待输入
    等待输入 --> 解析中 : 用户输入
    解析中 --> 计算中 : 解析成功
    解析中 --> 错误处理 : 解析失败
    计算中 --> 显示结果 : 计算成功
    计算中 --> 错误处理 : 计算失败
    显示结果 --> 等待输入 : 继续
    错误处理 --> 等待输入 : 重试
    等待输入 --> [*] : 退出
</script>

<MermaidDiagram id="state-diagram" />

## 类图示例

<script type="text/plain" id="class-diagram">
classDiagram
    class Calculator {
        -parser: ExpressionParser
        -history: HistoryManager
        +calculate(expression: string) float64
        +getHistory() []HistoryEntry
        +clearHistory()
    }
    
    class ExpressionParser {
        -tokens: []Token
        -position: int
        +parse(expression: string) float64
        +tokenize(expression: string) []Token
        -parseExpression() float64
        -parseTerm() float64
        -parseFactor() float64
    }
    
    class HistoryManager {
        -entries: []HistoryEntry
        -maxSize: int
        +add(expression: string, result: float64)
        +getAll() []HistoryEntry
        +clear()
    }
    
    class HistoryEntry {
        +expression: string
        +result: float64
        +timestamp: time.Time
    }
    
    Calculator --> ExpressionParser
    Calculator --> HistoryManager
    HistoryManager --> HistoryEntry
</script>

<MermaidDiagram id="class-diagram" />

## Git流程图示例

<script type="text/plain" id="git-diagram">
gitgraph
    commit id: "初始化"
    branch feature
    checkout feature
    commit id: "添加解析器"
    commit id: "添加计算器"
    checkout main
    commit id: "更新文档"
    merge feature
    commit id: "发布v1.0"
    branch hotfix
    checkout hotfix
    commit id: "修复除零错误"
    checkout main
    merge hotfix
    commit id: "发布v1.1"
</script>

<MermaidDiagram id="git-diagram" />

## 饼图示例

<script type="text/plain" id="pie-chart">
pie title 计算器功能使用统计
    "基础运算" : 45
    "括号表达式" : 25
    "历史查询" : 15
    "帮助命令" : 10
    "错误处理" : 5
</script>

<MermaidDiagram id="pie-chart" />

## 甘特图示例

<script type="text/plain" id="gantt-chart">
gantt
    title 计算器项目开发计划
    dateFormat  YYYY-MM-DD
    section 设计阶段
    需求分析           :active, design1, 2024-01-01, 3d
    架构设计           :design2, after design1, 2d
    界面设计           :design3, after design2, 2d
    section 开发阶段
    核心解析器         :dev1, after design3, 5d
    计算引擎           :dev2, after dev1, 3d
    用户界面           :dev3, after dev2, 3d
    section 测试阶段
    单元测试           :test1, after dev3, 3d
    集成测试           :test2, after test1, 2d
    用户测试           :test3, after test2, 2d
</script>

<MermaidDiagram id="gantt-chart" />

## 使用说明

### Script标签方式的优点

1. **避免特殊字符问题**：不需要在HTML属性中转义特殊字符
2. **更好的可读性**：代码格式清晰，易于维护
3. **语法高亮支持**：在编辑器中可以获得更好的语法高亮
4. **向后兼容**：仍然支持原有的`code`属性方式

### 使用方法

```html
<!-- 1. 定义Mermaid代码 -->
<script type="text/plain" id="unique-diagram-id">
graph TD
    A[开始] --> B[结束]
    
    style A fill:#e1f5fe
    style B fill:#f3e5f5
</script>

<!-- 2. 渲染图表 -->
<MermaidDiagram id="unique-diagram-id" />
```

### 注意事项

- `id`必须唯一
- `type="text/plain"`防止浏览器执行脚本
- 模板字符串中可以包含任何特殊字符
- 保持向后兼容，仍支持`code`属性方式 