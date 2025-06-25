<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

const modules = ref([
  {
    title: '学习 Learn',
    details: '从语言核心到并发编程，系统化地掌握Go的精髓。我们的目标是构建坚实的知识基座，而非零散技巧的堆砌。',
    link: '/go-learn/learn/'
  },
  {
    title: '实践 Practice',
    details: '将理论化为利器。通过真实世界的项目、设计模式和部署策略，将您的Go技能锻造成解决复杂问题的工程能力。',
    link: '/go-learn/practice/'
  },
  {
    title: '生态 Ecosystem',
    details: '探索Go语言繁荣的生态系统。从主流Web框架到强大的标准库，再到前沿的技术趋势，拓展您的技术视野。',
    link: '/go-learn/ecosystem/'
  }
]);

const cardRefs = ref([]);
const container = ref(null);

const onMouseMove = (el, event) => {
  const { left, top, width, height } = el.getBoundingClientRect();
  const x = event.clientX - left;
  const y = event.clientY - top;

  const rotateX = (y / height - 0.5) * -20;
  const rotateY = (x / width - 0.5) * 20;

  el.style.transform = `perspective(500px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) scale3d(1.1, 1.1, 1.1)`;
};

const onMouseLeave = (el) => {
  el.style.transform = 'perspective(500px) rotateX(0) rotateY(0) scale3d(1, 1, 1)';
};

onMounted(() => {
  cardRefs.value.forEach(card => {
    if (card) {
      card.addEventListener('mousemove', (e) => onMouseMove(card, e));
      card.addEventListener('mouseleave', () => onMouseLeave(card));
    }
  });
});

onUnmounted(() => {
    cardRefs.value.forEach(card => {
    if (card) {
      card.removeEventListener('mousemove', onMouseMove);
      card.removeEventListener('mouseleave', onMouseLeave);
    }
  });
});

</script>

<template>
  <div class="module-showcase-container" ref="container">
    <div
      v-for="(mod, index) in modules"
      :key="mod.title"
      :ref="el => cardRefs[index] = el"
      class="module-card"
      @click="() => { if (mod.link) $router.push(mod.link) }"
    >
      <div class="card-content">
        <div class="title">{{ mod.title }}</div>
        <p class="details">{{ mod.details }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.module-showcase-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
  padding: 2rem 0;
  max-width: 1152px;
  margin: 0 auto;
}

.module-card {
  background: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-bg-soft);
  padding: 2rem;
  transition: transform 0.3s ease-out, box-shadow 0.3s ease-out;
  transform-style: preserve-3d;
  will-change: transform;
}

.module-card:hover {
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  border-color: var(--vp-c-brand-1);
}

.card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.title {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--vp-c-text-1);
  margin-bottom: 0.75rem;
  border: none;
}

.details {
  font-size: 1rem;
  color: var(--vp-c-text-2);
  line-height: 1.6;
}

@media (max-width: 768px) {
  .module-showcase-container {
    padding: 1rem;
    gap: 1.5rem;
  }
  .module-card {
    padding: 1.5rem;
  }
}
</style> 