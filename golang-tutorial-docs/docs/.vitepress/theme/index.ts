import DefaultTheme from 'vitepress/theme'
import MermaidDiagram from '../components/MermaidDiagram.vue'
import './style.css'

export default {
    extends: DefaultTheme,
    enhanceApp({ app, router, siteData }) {
        // 注册全局组件
        app.component('MermaidDiagram', MermaidDiagram)
    }
} 