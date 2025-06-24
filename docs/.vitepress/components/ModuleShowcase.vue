<template>
    <div class="module-showcase">
        <div class="modules-container">
            <div v-for="module in modules" :key="module.id" class="module-card"
                :class="{ active: activeModule === module.id }" @mouseenter="handleMouseEnter(module.id)"
                @mouseleave="handleMouseLeave" @click="navigateToModule(module.link)">
                <div class="module-icon">
                    <span :class="module.iconClass">{{ module.icon }}</span>
                </div>
                <div class="module-content">
                    <h3 class="module-title">{{ module.title }}</h3>
                    <p class="module-description">{{ module.description }}</p>

                    <!-- 使用固定高度容器避免重排 -->
                    <div class="module-features-container">
                        <div class="module-features" :class="{ visible: activeModule === module.id }">
                            <ul>
                                <li v-for="feature in module.features" :key="feature">
                                    <span class="feature-dot">•</span>
                                    {{ feature }}
                                </li>
                            </ul>
                        </div>
                    </div>

                    <div class="module-action">
                        <span class="action-text">{{ getActionText(module.id) }}</span>
                        <span class="action-arrow">→</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vitepress'

const router = useRouter()
const activeModule = ref(null)
const showPath = ref(false)
const hoverTimeout = ref(null)

const modules = [
    {
        id: 'learn',
        title: '学习 Learn',
        icon: '📚',
        iconClass: 'icon-learn',
        description: '系统性学习Go语言核心概念，从基础语法到高级特性，构建扎实的理论基础',
        features: [
            '基础语法：变量类型、控制流、函数、接口等核心语法',
            '进阶特性：并发编程、反射、CGO、内存管理等高级特性',
            '核心概念：Go的设计理念、类型系统、接口机制、内存模型'
        ],
        link: 'go-learn/learn/'
    },
    {
        id: 'practice',
        title: '工程实践 Practice',
        icon: '⚡',
        iconClass: 'icon-practice',
        description: '通过实战项目、设计模式和工具链掌握，提升Go工程开发能力',
        features: [
            '实战项目：从简单CLI工具到复杂Web服务的完整项目',
            '设计模式：Go语言中的设计模式最佳实践',
            '工具链：开发、测试、调试、性能分析工具的使用',
            '部署运维：容器化、CI/CD、监控等生产环境实践'
        ],
        link: 'go-learn/practice/'
    },
    {
        id: 'ecosystem',
        title: '生态 Ecosystem',
        icon: '🌍',
        iconClass: 'icon-ecosystem',
        description: '深入了解Go生态体系，掌握主流框架、库和社区资源',
        features: [
            '框架生态：Web框架、微服务框架、数据库ORM等',
            '常用库：HTTP客户端、JSON处理、加密、日志等实用库',
            '社区资源：官方文档、技术博客、开源项目推荐',
            '技术趋势：Go语言发展趋势和新特性预览'
        ],
        link: 'go-learn/ecosystem/'
    }
]

// 防抖处理hover事件
const handleMouseEnter = (moduleId) => {
    if (hoverTimeout.value) {
        clearTimeout(hoverTimeout.value)
    }
    hoverTimeout.value = setTimeout(() => {
        activeModule.value = moduleId
    }, 100) // 100ms延迟避免频繁触发
}

const handleMouseLeave = () => {
    if (hoverTimeout.value) {
        clearTimeout(hoverTimeout.value)
    }
    hoverTimeout.value = setTimeout(() => {
        activeModule.value = null
    }, 150) // 稍长的延迟避免快速移出
}

const getActionText = computed(() => (moduleId) => {
    return activeModule.value === moduleId ? '点击进入' : '悬停查看详情'
})

const navigateToModule = (link) => {
    router.go(link)
}

const togglePath = () => {
    showPath.value = !showPath.value
}
</script>

<style scoped>
.module-showcase {
    margin: 2rem 0;
    padding: 2rem 0;
}

.modules-container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
    gap: 2rem;
    margin-bottom: 3rem;
}

.module-card {
    background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
    border: 2px solid #e2e8f0;
    border-radius: 16px;
    padding: 2rem;
    cursor: pointer;
    transition: all 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
    position: relative;
    overflow: hidden;
    will-change: transform, box-shadow;
}

.module-card::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: linear-gradient(90deg, #00ADD8, #00D9FF);
    transform: scaleX(0);
    transition: transform 0.6s cubic-bezier(0.25, 0.8, 0.25, 1);
    transform-origin: left;
}

.module-card:hover::before,
.module-card.active::before {
    transform: scaleX(1);
}

.module-card:hover,
.module-card.active {
    transform: translateY(-4px);
    box-shadow: 0 16px 32px rgba(0, 173, 216, 0.12);
    border-color: #00ADD8;
}

.dark .module-card {
    background: linear-gradient(135deg, #1e293b 0%, #334155 100%);
    border-color: #475569;
}

.dark .module-card:hover,
.dark .module-card.active {
    border-color: #00D9FF;
    box-shadow: 0 16px 32px rgba(0, 217, 255, 0.15);
}

.module-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 64px;
    height: 64px;
    background: linear-gradient(135deg, #00ADD8, #00D9FF);
    border-radius: 16px;
    margin-bottom: 1.5rem;
    font-size: 2rem;
    transition: transform 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
    will-change: transform;
}

.module-card:hover .module-icon,
.module-card.active .module-icon {
    transform: scale(1.05) rotate(3deg);
}

.module-title {
    font-size: 1.5rem;
    font-weight: 700;
    color: #1e293b;
    margin-bottom: 1rem;
}

.dark .module-title {
    color: #f1f5f9;
}

.module-description {
    color: #64748b;
    line-height: 1.6;
    margin-bottom: 1rem;
}

.dark .module-description {
    color: #cbd5e1;
}

/* 固定高度容器避免重排 */
.module-features-container {
    height: 0;
    overflow: hidden;
    margin-bottom: 1rem;
    transition: height 0.5s cubic-bezier(0.25, 0.8, 0.25, 1);
}

.module-card.active .module-features-container {
    height: 120px;
    /* 根据内容调整固定高度 */
}

.module-features {
    opacity: 0;
    transform: translateY(-10px);
    transition: all 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
    will-change: opacity, transform;
}

.module-features.visible {
    opacity: 1;
    transform: translateY(0);
}

.module-features ul {
    list-style: none;
    padding: 0;
    margin: 0;
}

.module-features li {
    display: flex;
    align-items: flex-start;
    margin-bottom: 0.5rem;
    color: #475569;
    font-size: 0.9rem;
    line-height: 1.5;
}

.dark .module-features li {
    color: #94a3b8;
}

.feature-dot {
    color: #00ADD8;
    font-weight: bold;
    margin-right: 0.5rem;
    flex-shrink: 0;
}

.module-action {
    display: flex;
    align-items: center;
    justify-content: space-between;
    color: #00ADD8;
    font-weight: 600;
    font-size: 0.9rem;
}

.action-arrow {
    transition: transform 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
    will-change: transform;
}

.module-card:hover .action-arrow,
.module-card.active .action-arrow {
    transform: translateX(4px);
}

/* 响应式设计 */
@media (max-width: 768px) {
    .modules-container {
        grid-template-columns: 1fr;
    }

    .path-flow {
        flex-direction: column;
    }

    .module-card {
        padding: 1.5rem;
    }

    .module-card.active .module-features-container {
        height: 140px;
        /* 移动端调整高度 */
    }
}

/* 性能优化 */
@media (prefers-reduced-motion: reduce) {
    * {
        animation-duration: 0.01ms !important;
        animation-iteration-count: 1 !important;
        transition-duration: 0.01ms !important;
    }
}
</style>