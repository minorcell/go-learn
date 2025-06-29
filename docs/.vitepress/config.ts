import { defineConfig } from 'vitepress'

export default defineConfig({
    title: 'Gopher成长指南',
    description: '全栈式Go语言学习路径，从语言基础到工程实践，从生态认知到职业成长',

    lang: 'zh-CN',
    base: '/go-learn/',
    cleanUrls: true,

    head: [
        ['link', { rel: 'icon', href: '/logo.png' }],
        ['meta', { name: 'theme-color', content: '#00ADD8' }],
        ['meta', { name: 'og:type', content: 'website' }],
        ['meta', { name: 'og:locale', content: 'zh_CN' }],
        ['meta', { name: 'og:site_name', content: 'Gopher成长指南' }],
        ['meta', { name: 'og:title', content: 'Gopher成长指南 | 全栈式Go语言学习路径' }],
        ['meta', { name: 'og:description', content: '从语言基础到工程实践，从生态认知到职业成长的完整指南' }],
    ],

    themeConfig: {
        logo: '/logo.png',
        siteTitle: 'Gopher成长指南',

        nav: [
            { text: '首页', link: '/' },
            {
                text: '学习 Learn',
                items: [
                    { text: '学习首页', link: '/learn/' },
                    { text: '基础语法', link: '/learn/fundamentals/' },
                    { text: '进阶特性', link: '/learn/advanced/' },
                    { text: '核心概念', link: '/learn/concepts/' }
                ]
            },
            {
                text: '工程实践 Practice',
                items: [
                    { text: '实践首页', link: '/practice/' },
                    { text: '实战项目', link: '/practice/projects/' },
                    { text: '设计模式', link: '/practice/patterns/' },
                    { text: '工具链', link: '/practice/tools/' },
                    { text: '部署运维', link: '/practice/deployment/' }
                ]
            },
            {
                text: '生态 Ecosystem',
                items: [
                    { text: '生态首页', link: '/ecosystem/' },
                    { text: '框架生态', link: '/ecosystem/frameworks/' },
                    { text: '常用库', link: '/ecosystem/libraries/' },
                    { text: '社区资源', link: '/ecosystem/community/' },
                    { text: '技术趋势', link: '/ecosystem/trends/' }
                ]
            }
        ],

        sidebar: {
            // 学习模块侧边栏
            '/learn/': [
                {
                    text: '学习模块',
                    items: [
                        { text: '学习首页', link: '/learn/' }
                    ]
                },
                {
                    text: '基础入门',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/learn/fundamentals/' },
                        { text: '变量与类型', link: '/learn/fundamentals/variables-types' },
                        { text: '数组、切片和映射', link: '/learn/fundamentals/arrays-slices-maps' },
                        { text: '控制流', link: '/learn/fundamentals/control-flow' },
                        { text: '函数', link: '/learn/fundamentals/functions' },
                        { text: '指针', link: '/learn/fundamentals/pointers' },
                        { text: '结构体', link: '/learn/fundamentals/structs' }
                    ]
                },
                {
                    text: '深入核心特性',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/learn/advanced/' },
                        { text: '方法', link: '/learn/advanced/methods' },
                        { text: '接口', link: '/learn/advanced/interfaces' },
                        { text: '并发', link: '/learn/advanced/concurrency' },
                        { text: '泛型', link: '/learn/advanced/generics' },
                        { text: '测试', link: '/learn/advanced/testing' }
                    ]
                },
                {
                    text: '核心概念',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/learn/concepts/' },
                        { text: '错误处理', link: '/learn/concepts/errors' },
                        { text: '设计哲学', link: '/learn/concepts/philosophy' },
                        { text: '类型系统', link: '/learn/concepts/type-system' },
                        { text: '内存模型', link: '/learn/concepts/memory-model' },
                        { text: '运行时', link: '/learn/concepts/runtime' },
                        { text: '编译器', link: '/learn/concepts/compiler' }
                    ]
                }
            ],

            // 工程实践模块侧边栏
            '/practice/': [
                {
                    text: '工程实践模块',
                    items: [
                        { text: '实践首页', link: '/practice/' }
                    ]
                },
                {
                    text: '实战项目',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/practice/projects/' },
                        { text: 'CLI命令行工具', link: '/practice/projects/cli-tools' },
                        { text: 'Web API服务', link: '/practice/projects/web-api' },
                        { text: '微服务架构实践', link: '/practice/projects/microservices' },
                        { text: '数据库应用开发', link: '/practice/projects/database-app' },
                        { text: '分布式系统设计', link: '/practice/projects/distributed-systems' }
                    ]
                },
                {
                    text: '设计模式',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/practice/patterns/' },
                        { text: '创建型模式', link: '/practice/patterns/creational' },
                        { text: '结构型模式', link: '/practice/patterns/structural' },
                        { text: '行为型模式', link: '/practice/patterns/behavioral' },
                        { text: '并发模式', link: '/practice/patterns/concurrency' },
                        { text: '错误处理模式', link: '/practice/patterns/error-handling' }
                    ]
                },
                {
                    text: '工具链',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/practice/tools/' },
                        { text: '工欲善其事：搭建开发环境', link: '/practice/tools/development-setup' },
                        { text: '代码的准星：静态分析', link: '/practice/tools/code-quality' },
                        { text: '精密瞄准镜：测试套件', link: '/practice/tools/testing' },
                        { text: '高精度放大镜：性能剖析', link: '/practice/tools/profiling' },
                        { text: '铸造成器：编译与构建', link: '/practice/tools/build-deploy' }
                    ]
                },
                {
                    text: '部署运维',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/practice/deployment/' },
                        { text: '容器化实践', link: '/practice/deployment/containerization' },
                        { text: 'CI/CD流水线', link: '/practice/deployment/cicd' },
                        { text: '监控和日志', link: '/practice/deployment/monitoring' },
                        { text: '性能优化', link: '/practice/deployment/performance' },
                        { text: '云原生部署', link: '/practice/deployment/cloud-native' }
                    ]
                }
            ],

            // 生态模块侧边栏
            '/ecosystem/': [
                {
                    text: '生态模块',
                    items: [
                        { text: '生态首页', link: '/ecosystem/' }
                    ]
                },
                {
                    text: '框架生态',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/ecosystem/frameworks/' },
                        { text: 'Web框架', link: '/ecosystem/frameworks/web' },
                        { text: '微服务框架', link: '/ecosystem/frameworks/microservices' },
                        { text: 'ORM框架', link: '/ecosystem/frameworks/orm' },
                        { text: 'RPC框架', link: '/ecosystem/frameworks/rpc' },
                        { text: '消息队列客户端', link: '/ecosystem/frameworks/messaging' }
                    ]
                },
                {
                    text: '常用库',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/ecosystem/libraries/' },
                        { text: 'HTTP客户端库', link: '/ecosystem/libraries/http-clients' },
                        { text: 'JSON和序列化', link: '/ecosystem/libraries/serialization' },
                        { text: '加密和安全', link: '/ecosystem/libraries/security' },
                        { text: '日志库', link: '/ecosystem/libraries/logging' },
                        { text: '配置管理', link: '/ecosystem/libraries/configuration' }
                    ]
                },
                {
                    text: '社区资源',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/ecosystem/community/' },
                        { text: '官方资源', link: '/ecosystem/community/official' },
                        { text: '技术博客和文章', link: '/ecosystem/community/blogs' },
                        { text: '开源项目推荐', link: '/ecosystem/community/open-source' },
                        { text: '学习资料和书籍', link: '/ecosystem/community/learning-resources' },
                        { text: '社区论坛和交流', link: '/ecosystem/community/forums' }
                    ]
                },
                {
                    text: '技术趋势',
                    collapsed: false,
                    items: [
                        { text: '概览', link: '/ecosystem/trends/' },
                        { text: 'Go与云原生', link: '/ecosystem/trends/cloud-native' },
                        { text: '编译器与性能', link: '/ecosystem/trends/compiler-improvements' },
                        { text: '版本历史与演进', link: '/ecosystem/trends/version-history' },
                        { text: '泛型演进史', link: '/ecosystem/trends/generics-evolution' },
                        { text: 'AI与Go', link: '/ecosystem/trends/ai-integration' }
                    ]
                }
            ]
        },

        socialLinks: [
            { icon: 'github', link: 'https://github.com/minorcell/go-learn' }
        ],

        footer: {
            copyright: `Gopher成长指南 Copyright © 2025 <a href="https://github.com/minorcell/go-learn">mCell</a>`
        },

        editLink: {
            pattern: 'https://github.com/minorcell/go-learn/edit/main/docs/:path',
            text: '在 GitHub 上编辑此页'
        },

        lastUpdated: {
            text: '最后更新于',
            formatOptions: {
                dateStyle: 'short',
                timeStyle: 'medium'
            }
        },

        docFooter: {
            prev: '上一页',
            next: '下一页'
        },

        outline: {
            level: [2, 3],
            label: '页面导航'
        },

        returnToTopLabel: '回到顶部',
        sidebarMenuLabel: '菜单',
        darkModeSwitchLabel: '主题',
        lightModeSwitchTitle: '切换到浅色模式',
        darkModeSwitchTitle: '切换到深色模式',

        search: {
            provider: 'local',
            options: {
                locales: {
                    root: {
                        translations: {
                            button: {
                                buttonText: '搜索文档',
                                buttonAriaLabel: '搜索文档'
                            },
                            modal: {
                                noResultsText: '无法找到相关结果',
                                resetButtonTitle: '清除查询条件',
                                footer: {
                                    selectText: '选择',
                                    navigateText: '切换',
                                    closeText: '关闭'
                                }
                            }
                        }
                    }
                }
            }
        }
    },

    markdown: {
        theme: {
            light: 'github-light',
            dark: 'github-dark'
        },
        lineNumbers: true,
        container: {
            tipLabel: '提示',
            warningLabel: '警告',
            dangerLabel: '危险',
            infoLabel: '信息',
            detailsLabel: '详细信息'
        }
    },

    vite: {
        define: {
            __VUE_OPTIONS_API__: false
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `
                        :root {
                            --vp-c-brand-1: #00ADD8;
                            --vp-c-brand-2: #00ADD8;
                            --vp-c-brand-3: #00ADD8;
                            --vp-c-brand-soft: rgba(0, 173, 216, 0.14);
                            --vp-c-brand-softer: rgba(0, 173, 216, 0.07);
                            --vp-c-brand-softest: rgba(0, 173, 216, 0.04);
                        }
                    `
                }
            }
        }
    }
})
