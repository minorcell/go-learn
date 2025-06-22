<template>
  <div class="mermaid-wrapper">
    <div ref="mermaidRef" class="mermaid-content"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'

const props = defineProps<{
  code: string
}>()

const mermaidRef = ref<HTMLElement>()

onMounted(async () => {
  await nextTick()
  
  try {
    const mermaid = await import('mermaid')
    
    // 配置Mermaid
    mermaid.default.initialize({
      startOnLoad: false,
      theme: 'default',
      themeVariables: {
        primaryColor: '#00ADD8',
        primaryTextColor: '#ffffff',
        primaryBorderColor: '#0099CC',
        lineColor: '#00ADD8',
        sectionBkgColor: '#f0f9ff',
        altSectionBkgColor: '#e0f4f7',
        gridColor: '#cccccc',
        secondaryColor: '#5DC9E2',
        tertiaryColor: '#ffffff',
        background: '#ffffff',
        mainBkg: '#00ADD8',
        secondBkg: '#5DC9E2',
        tertiaryBkg: '#ffffff'
      }
    })
    
    if (mermaidRef.value) {
      const id = `mermaid-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`
      const { svg } = await mermaid.default.render(id, props.code)
      mermaidRef.value.innerHTML = svg
    }
  } catch (error) {
    console.error('Mermaid渲染失败:', error)
    if (mermaidRef.value) {
      mermaidRef.value.innerHTML = `<pre class="mermaid-error">${props.code}</pre>`
    }
  }
})
</script>

<style scoped>
.mermaid-wrapper {
  margin: 2rem 0;
  text-align: center;
}

.mermaid-content {
  background: var(--vp-c-bg-soft);
  border-radius: 8px;
  border: 1px solid var(--vp-c-border);
  padding: 1rem;
  overflow-x: auto;
}

.mermaid-error {
  background: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-border);
  border-radius: 4px;
  padding: 1rem;
  color: var(--vp-c-text-2);
  font-family: var(--vp-font-family-mono);
}

:deep(.mermaid svg) {
  max-width: 100%;
  height: auto;
}
</style> 