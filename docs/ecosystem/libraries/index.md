# Go第三方库：深度选型与工程实践

欢迎来到Go语言的第三方库（Libraries）模块。本模块并非一个简单的库列表，而是一系列深度剖析和选型指南，旨在帮助有经验的Go工程师在构建复杂系统时，就关键技术栈做出明智的决策。

我们摒弃了千篇一律的介绍，为每一个核心领域都撰写了具有独特风格和视角的深度文章。

## 核心主题

<div class="vp-doc">
  <ul>
    <li>
      <a href="./http-clients.md">HTTP客户端</a> - <span class="custom-tag">实战对比与选型指南</span>
      <p>在标准库`net/http`、性能王者`fasthttp`和便捷封装`resty`之间如何选择？本文通过性能基准、API对比和场景分析，为你提供清晰的决策树。</p>
    </li>
    <li>
      <a href="./serialization.md">序列化</a> - <span class="custom-tag">深度技术剖析</span>
      <p>深入`JSON`, `Protobuf`, `MessagePack`和`Gob`的对决。我们剖析了它们的性能、体积、跨语言能力和使用场景，助你为微服务和数据存储选择最优协议。</p>
    </li>
    <li>
      <a href="./security.md">安全</a> - <span class="custom-tag">威胁建模与防御实践</span>
      <p>以攻击者的视角审视你的应用。本文围绕OWASP Top 10，通过`bcrypt`密码哈希、参数化查询、`html/template`防XSS等，提供可落地的Go安全编码实践。</p>
    </li>
    <li>
      <a href="./logging.md">日志</a> - <span class="custom-tag">工程治理与可观测性</span>
      <p>日志是生产系统的眼睛。本文对比了官方新标准`slog`、性能王者`zerolog`和功能强大的`zap`，并探讨了结构化日志、日志分级和上下文传递等工程最佳实践。</p>
    </li>
    <li>
      <a href="./configuration.md">配置管理</a> - <span class="custom-tag">系统设计与演进</span>
      <p>从简单的环境变量到动态配置中心，本文探讨了配置管理的演进之路，并提供了业界标准`Viper`的完整使用指南，助你构建灵活、可控的配置体系。</p>
    </li>
  </ul>
</div> 