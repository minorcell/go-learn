<script setup>
import { VPButton } from 'vitepress/theme'
</script>

<template>
  <div class="NotFoundContainer">
    <div class="glitch-wrapper">
      <div class="glitch" data-text="404">404</div>
    </div>
    <h1 class="Title">
      <span class="prompt">></span> NIL POINTER DEREFERENCE<span class="cursor">_</span>
    </h1>
    <p class="Message">
      你访问的页面似乎指向了一个空地址 (nil)。在 Go 的世界里，这意味着一次无效的内存寻址；在这里，它可能是一个失效的链接或一个拼写错误的 URL。
    </p>
    <div class="Actions">
      <VPButton
        tag="a"
        size="big"
        theme="brand"
        text="返回首页"
        href="/"
      />
      <VPButton
        tag="a"
        size="big"
        theme="alt"
        text="报告问题"
        href="https://github.com/minorcell/go-learn/issues/new"
      />
    </div>
  </div>
</template>

<style scoped>
.NotFoundContainer {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 1.5rem;
  height: calc(100vh - var(--vp-nav-height));
  box-sizing: border-box;
  text-align: center;
}

.glitch-wrapper {
  position: relative;
  margin-bottom: 2rem;
}

.glitch {
  position: relative;
  font-size: 8rem;
  font-weight: 700;
  line-height: 1;
  color: var(--vp-c-text-1);
  z-index: 1;
  animation: glitch-shake 2.5s infinite linear;
}

.glitch::before,
.glitch::after {
  content: attr(data-text);
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: var(--vp-c-bg);
  overflow: hidden;
}

.glitch::before {
  left: 2px;
  text-shadow: -2px 0 var(--vp-c-brand-1);
  clip-path: rect(44px, 450px, 56px, 0);
  animation: glitch-anim-1 2s infinite linear alternate-reverse;
}

.glitch::after {
  left: -2px;
  text-shadow: -2px 0 var(--vp-c-red-1), 2px 2px var(--vp-c-brand-1);
  clip-path: rect(85px, 450px, 90px, 0);
  animation: glitch-anim-2 3s infinite linear alternate-reverse;
}

.Title {
  font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, Courier, monospace;
  font-size: 1.75rem;
  font-weight: 600;
  color: var(--vp-c-text-2);
  margin: 0;
  letter-spacing: 0.05em;
}

.prompt {
  color: var(--vp-c-brand-1);
  margin-right: 0.5rem;
}

.cursor {
  color: var(--vp-c-brand-1);
  animation: blink 1s steps(1) infinite;
}

.Message {
  font-size: 1.1rem;
  color: var(--vp-c-text-2);
  margin: 1.5rem auto;
  max-width: 480px;
  line-height: 1.6;
}

.Actions {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-top: 2rem;
}

@keyframes blink {
  50% {
    opacity: 0;
  }
}

@keyframes glitch-shake {
  0%, 100% { transform: skewX(0); }
  25% { transform: skewX(2deg); }
  50% { transform: skewX(-2deg); }
  75% { transform: skewX(1deg); }
}

@keyframes glitch-anim-1 {
  0% { clip-path: rect(42px, 9999px, 44px, 0); }
  25% { clip-path: rect(5px, 9999px, 80px, 0); }
  50% { clip-path: rect(70px, 9999px, 20px, 0); }
  75% { clip-path: rect(30px, 9999px, 90px, 0); }
  100% { clip-path: rect(50px, 9999px, 10px, 0); }
}

@keyframes glitch-anim-2 {
  0% { clip-path: rect(85px, 9999px, 90px, 0); }
  25% { clip-path: rect(25px, 9999px, 70px, 0); }
  50% { clip-path: rect(50px, 9999px, 60px, 0); }
  75% { clip-path: rect(10px, 9999px, 95px, 0); }
  100% { clip-path: rect(80px, 9999px, 40px, 0); }
}

@media (max-width: 768px) {
  .glitch {
    font-size: 6rem;
  }
  .Title {
    font-size: 1.25rem;
  }
}

@media (max-width: 480px) {
  .glitch {
    font-size: 4rem;
  }
  .Title {
    font-size: 1rem;
  }
  .Message {
    font-size: 1rem;
  }
  .Actions {
    flex-direction: column;
    width: 100%;
    max-width: 280px;
  }
  .Actions :deep(.VPButton) {
    width: 100%;
    margin: 0;
  }
}
</style> 