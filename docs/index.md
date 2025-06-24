---
layout: home

hero:
  name: "Gopher成长指南"
  text: "全栈式Go语言学习路径"
  tagline: 从语言基础到工程实践，从生态认知到职业成长的完整指南
  image:
    src: /logo.png
    alt: Gopher成长指南
  actions:
    - theme: alt
      text: 开始学习
      link: /learn/
    - theme: alt
      text: 工程实践
      link: /practice/
    - theme: alt
      text: 生态探索
      link: /ecosystem/
    - theme: brand
      text: GitHub
      link: https://github.com/minorcell/go-learn

---

<ModuleShowcase />

<script setup>
import ModuleShowcase from './.vitepress/components/ModuleShowcase.vue'
</script>