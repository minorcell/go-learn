# GitHub Pages 自动部署指南

本项目已配置GitHub Actions自动部署到GitHub Pages。

## 部署设置步骤

### 1. 启用GitHub Pages

1. 进入你的GitHub仓库 (`minorcell/go-learn`)
2. 点击 **Settings** 标签页
3. 在左侧菜单中找到 **Pages**
4. 在 **Source** 部分选择 **GitHub Actions**

### 2. 权限配置

确保工作流具有必要的权限：

1. 在仓库的 **Settings** > **Actions** > **General** 中
2. 在 **Workflow permissions** 部分选择：
   - **Read and write permissions** 
   - 勾选 **Allow GitHub Actions to create and approve pull requests**

### 3. 触发部署

部署会在以下情况自动触发：

- 推送到 `main` 或 `goang` 分支
- 手动在 Actions 页面触发工作流

### 4. 查看部署状态

1. 进入仓库的 **Actions** 标签页
2. 查看 "Deploy VitePress site to Pages" 工作流状态
3. 部署成功后，网站将在以下地址可访问：
   
   **https://minorcell.github.io/go-learn/**

## 工作流说明

- **构建阶段**：使用Node.js 18安装依赖并构建VitePress项目
- **部署阶段**：将构建产物上传到GitHub Pages

## 本地开发

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建项目
npm run build

# 预览构建结果
npm run preview
```

## 故障排除

### 常见问题

1. **部署失败**：检查Actions权限设置
2. **页面404**：确认base路径配置正确 (`/go-learn/`)
3. **样式丢失**：检查资源路径是否正确

### 查看日志

在GitHub仓库的Actions页面可以查看详细的构建和部署日志。

## 注意事项

- 首次部署可能需要几分钟才能生效
- 推送到配置的分支会自动触发新的部署
- VitePress配置中的 `base: '/go-learn/'` 是必需的，因为GitHub Pages会在子路径下提供服务 