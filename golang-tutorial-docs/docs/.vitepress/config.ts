import { defineConfig } from 'vitepress'

export default defineConfig({
    title: 'Go语言教程',
    description: '从零开始学习Go语言编程',

    head: [
        ['link', { rel: 'icon', href: '/favicon.ico' }],
        ['meta', { name: 'keywords', content: 'Go,Golang,教程,编程,后端开发' }],
        ['meta', { name: 'author', content: 'Go语言教程' }]
    ],

    themeConfig: {
        logo: '/logo.png',

        nav: [
            { text: '首页', link: '/' },
            { text: '学习指南', link: '/guide/' },
            { text: '基础语法', link: '/basics/' },
            { text: '进阶内容', link: '/advanced/' },
            { text: '实战项目', link: '/projects/' }
        ],

        sidebar: {
            '/guide/': [
                {
                    text: '学习指南',
                    items: [
                        { text: '概览', link: '/guide/' },
                        { text: 'Go语言简介', link: '/guide/introduction' },
                        { text: '环境搭建', link: '/guide/setup' },
                        { text: '学习路线', link: '/guide/roadmap' }
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
                        { text: '时间和加密', link: '/advanced/time-crypto' }
                    ]
                }
            ],

            '/projects/': [
                {
                    text: '实战项目',
                    items: [
                        { text: '概览', link: '/projects/' },
                        { text: '计算器程序', link: '/projects/calculator/' },
                        { text: '命令行工具', link: '/projects/todo-cli/' },
                        { text: 'Web服务器', link: '/projects/web-server/' }
                    ]
                }
            ]
        },

        socialLinks: [
            { icon: 'github', link: 'https://github.com' }
        ],

        footer: {
            message: '基于 VitePress 构建',
            copyright: 'Copyright © 2024 Go语言教程'
        },

        search: {
            provider: 'local'
        },

        editLink: {
            pattern: 'https://github.com/your-username/golang-tutorial/edit/main/docs/:path',
            text: '在 GitHub 上编辑此页'
        },

        lastUpdated: {
            text: '最后更新于',
            formatOptions: {
                dateStyle: 'short',
                timeStyle: 'medium'
            }
        }
    },

    markdown: {
        lineNumbers: true,
        theme: {
            light: 'github-light',
            dark: 'github-dark'
        }
    }
})
