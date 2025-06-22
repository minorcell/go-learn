import { defineConfig } from 'vitepress'

export default defineConfig({
    title: 'Go语言学习指南',
    description: '系统化的Go编程教程，从基础语法到项目实战',

    lang: 'zh-CN',
    base: '/',
    cleanUrls: true,

    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }],
        ['meta', { name: 'theme-color', content: '#00ADD8' }],
        ['meta', { name: 'og:type', content: 'website' }],
        ['meta', { name: 'og:locale', content: 'zh_CN' }],
        ['meta', { name: 'og:site_name', content: 'Go语言学习指南' }],
        ['meta', { name: 'og:title', content: 'Go语言学习指南 | 系统化Go编程教程' }],
        ['meta', { name: 'og:description', content: '从基础语法到项目实战，助您掌握现代Go开发' }],
    ],

    themeConfig: {
        logo: '/logo.png',
        siteTitle: 'Go学习指南',

        nav: [
            { text: '首页', link: '/' },
            { text: '学习指南', link: '/guide/' },
            {
                text: '教程',
                items: [
                    { text: '基础语法', link: '/basics/' },
                    { text: '进阶内容', link: '/advanced/' },
                    { text: '实战项目', link: '/projects/' }
                ]
            },
        ],

        sidebar: {
            '/guide/': [
                {
                    text: '学习指南',
                    items: [
                        { text: 'Go语言简介', link: '/guide/introduction' },
                        { text: '环境搭建', link: '/guide/setup' }
                    ]
                }
            ],

            '/basics/': [
                {
                    text: '基础语法',
                    items: [
                        { text: '概览', link: '/basics/' },
                        { text: '变量和类型', link: '/basics/variables-types' },
                        { text: '控制流程', link: '/basics/control-flow' },
                        { text: '函数', link: '/basics/functions' },
                        { text: '数组切片映射', link: '/basics/arrays-slices' },
                        { text: '结构体和方法', link: '/basics/structs-methods' },
                        { text: '接口', link: '/basics/interfaces' },
                        { text: '错误处理', link: '/basics/error-handling' }
                    ]
                }
            ],

            '/advanced/': [
                {
                    text: '进阶内容',
                    items: [
                        { text: '概览', link: '/advanced/' },
                        { text: '包管理', link: '/advanced/packages' },
                        { text: '并发编程', link: '/advanced/concurrency' },
                        { text: '文件操作', link: '/advanced/file-operations' },
                        { text: '网络编程', link: '/advanced/network-http' },
                        { text: '字符串和正则', link: '/advanced/strings-regexp' },
                        { text: '时间处理和加密', link: '/advanced/time-crypto' }
                    ]
                }
            ],

            '/projects/': [
                {
                    text: '实战项目',
                    items: [
                        { text: '项目概览', link: '/projects/' },
                        {
                            text: '计算器项目',
                            collapsed: false,
                            items: [
                                { text: '项目概述', link: '/projects/calculator/' },
                                { text: '架构设计', link: '/projects/calculator/architecture' },
                                { text: '代码实现', link: '/projects/calculator/implementation' }
                            ]
                        },
                        { text: 'TODO CLI', link: '/projects/todo-cli/' },
                        { text: 'Web服务器', link: '/projects/web-server/' }
                    ]
                }
            ]
        },

        socialLinks: [
            { icon: 'github', link: 'https://github.com/minorcell/go-learn' }
        ],

        footer: {
            message: '基于 <a href="https://vitepress.dev">VitePress</a> 构建的Go语言学习平台',
            copyright: `Copyright © 2024 <a href="https://github.com/minorcell/go-learn">Go学习指南</a>`
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
                        
                        .dark {
                            --vp-c-brand-1: #00D9FF;
                            --vp-c-brand-2: #00ADD8;
                            --vp-c-brand-3: #0099CC;
                            --vp-c-brand-soft: rgba(0, 217, 255, 0.16);
                            --vp-c-brand-softer: rgba(0, 217, 255, 0.08);
                            --vp-c-brand-softest: rgba(0, 217, 255, 0.04);
                        }
                    `
                }
            }
        }
    }
})
