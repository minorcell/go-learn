import Layout from './Layout.vue'
import type { Theme } from 'vitepress'
import './style.css'

export default {
    Layout,
    enhanceApp({ app, router, siteData }) {
        // app is the Vue 3 app instance from `createApp()`.
        // router is VitePress' custom router. `siteData` is
        // a `ref` of current site-level metadata.
    }
} satisfies Theme 