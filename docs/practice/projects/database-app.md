---
title: "项目复盘：GORM 与仓储模式的最佳实践"
description: "如何通过仓储模式（Repository Pattern）和优雅的事务处理，构建一个高内聚、低耦合、易于测试的 Go 数据库应用。"
---

# 项目复盘：GORM 与仓储模式的最佳实践

## 1. 问题的起点：交织的业务与数据

在使用 GORM 或其他 ORM 框架时，我们常常会在业务逻辑层（通常称为 `Service` 层）直接调用 GORM 的方法。一个典型的例子可能是"用户创建订单"：

```go
// service/order_service.go
func (s *OrderService) CreateOrder(userID uint, productID uint, amount int) error {
    // 1. 检查库存
    var product models.Product
    if err := s.db.First(&product, productID).Error; err != nil {
        return errors.New("商品不存在")
    }
    if product.Stock < amount {
        return errors.New("库存不足")
    }

    // 2. 扣减库存
    product.Stock -= amount
    if err := s.db.Save(&product).Error; err != nil {
        return err
    }
    
    // 3. 创建订单
    order := models.Order{UserID: userID, ProductID: productID, Amount: amount}
    if err := s.db.Create(&order).Error; err != nil {
        // 严重问题：库存已经扣减，但订单创建失败！
        return err
    }

    return nil
}
```

这段代码存在两个致命问题：
1.  **缺乏原子性**：如果在扣减库存后、创建订单前发生任何错误（如数据库连接中断），系统将处于数据不一致的状态（商品库存减少了，但用户并没有对应的订单）。
2.  **职责混淆**：`OrderService` 不仅承担了业务流程编排的职责，还深度耦合了 GORM 的具体实现 (`s.db.First`, `s.db.Save`)。这使得对 `OrderService` 的单元测试变得异常困难，我们必须依赖一个真实的数据库连接。

为了解决这两个问题，我们引入了仓储模式（Repository Pattern）并设计了一套优雅的事务处理机制。

## 2. 解决方案：仓储模式与事务管理

### 2.1. 定义仓储接口 (Repository Interface)

仓储模式的核心思想是：**通过接口，将业务逻辑与数据持久化机制分离开**。我们在 `domain/repository` 目录下为每个聚合根（如 `Order`, `Product`）定义一个仓储接口。

`domain/repository/product_repository.go`:
```go
type ProductRepository interface {
    FindByID(id uint) (*model.Product, error)
    Update(product *model.Product) error
}
```
`domain/repository/order_repository.go`:
```go
type OrderRepository interface {
    Create(order *model.Order) error
}
```

最关键的是，我们定义一个统一的 `Repository` 接口，它聚合了所有其他的仓储接口，并增加了一个 `Transaction` 方法。

`domain/repository/repository.go`:
```go
type Repositories interface {
    Order() OrderRepository
    Product() ProductRepository
    // Transaction 方法接收一个函数作为参数
    // 在这个函数内执行的所有操作，都将处于同一个事务中
    Transaction(func(Repositories) error) error
}
```

### 2.2. 实现仓储 (Repository Implementation)

我们在 `infrastructure/persistence` 目录下提供基于 GORM 的具体实现。

`infrastructure/persistence/repositories.go`:
```go
type repositories struct {
    db      *gorm.DB
    order   OrderRepository
    product ProductRepository
}

func NewRepositories(db *gorm.DB) repository.Repositories {
    return &repositories{
        db:      db,
        order:   newOrderRepository(db),
        product: newProductRepository(db),
    }
}

// withTx 返回一个新的 repositories 实例，但它内部的 db 是一个事务对象
func (r *repositories) withTx(tx *gorm.DB) repository.Repositories {
    return &repositories{
        db:      tx, // 注意，这里传入的是 tx
        order:   newOrderRepository(tx),
        product: newProductRepository(tx),
    }
}

func (r *repositories) Transaction(fn func(repository.Repositories) error) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // 使用 withTx 创建一个在事务内的仓储实例
        txRepo := r.withTx(tx)
        // 执行传入的业务逻辑函数
        return fn(txRepo)
    })
}
// ... Order() 和 Product() 方法的实现 ...
```
`GORM` 的 `db.Transaction` 方法会自动处理事务的开始、提交和回滚，非常方便。我们的 `withTx` 方法是关键，它确保了在事务闭包内调用的所有仓储方法，都使用的是同一个事务连接 `tx`。

## 3. 重构业务逻辑

现在，我们可以重构 `OrderService`，让它依赖于我们定义的 `Repositories` 接口，而不是 `gorm.DB`。

`service/order_service.go`:
```go
type OrderService struct {
    repos repository.Repositories
}

func NewOrderService(repos repository.Repositories) *OrderService {
    return &OrderService{repos: repos}
}

func (s *OrderService) CreateOrder(userID uint, productID uint, amount int) error {
    // 调用 Transaction 方法，将所有数据库操作包裹在一个原子操作中
    return s.repos.Transaction(func(txRepos repository.Repositories) error {
        // 1. 检查库存
        product, err := txRepos.Product().FindByID(productID)
        if err != nil {
            return errors.New("商品不存在")
        }
        if product.Stock < amount {
            return errors.New("库存不足")
        }

        // 2. 扣减库存
        product.Stock -= amount
        if err := txRepos.Product().Update(product); err != nil {
            return err // 发生错误，整个事务将回滚
        }

        // 3. 创建订单
        order := model.Order{UserID: userID, ProductID: productID, Amount: amount}
        if err := txRepos.Order().Create(&order); err != nil {
            return err // 发生错误，整个事务将回滚
        }

        return nil // 一切顺利，事务将在方法结束时自动提交
    })
}
```

## 4. 复盘与反思

通过这种模式，我们完美地解决了最初的两个问题：
-   **保证了原子性**：`CreateOrder` 中的所有数据库操作现在都在一个事务中执行。任何一步的失败都会导致整个事务回滚，保证了数据的一致性。
-   **实现了关注点分离**：`OrderService` 现在只负责编排业务逻辑，它对 GORM 一无所知，只与稳定的 `Repositories` 接口交互。

最大的收益体现在**可测试性**上。我们可以轻松地为 `Repositories` 接口创建一个 Mock 实现，用于 `OrderService` 的单元测试。

`service/order_service_test.go`:
```go
// MockRepositories 实现了 repository.Repositories 接口
type MockRepositories struct {
    // ...
}
// ... 实现 MockRepositories 的所有方法 ...

func (m *MockRepositories) Transaction(fn func(repository.Repositories) error) error {
    // 在 Mock 中，我们可以直接执行该函数，因为不需要真正的事务
    return fn(m)
}

func TestCreateOrder_Success(t *testing.T) {
    mockRepos := &MockRepositories{}
    // ... 设置 mock 对象的预期行为 ...
    
    orderService := NewOrderService(mockRepos)
    err := orderService.CreateOrder(1, 101, 2)
    
    assert.NoError(t, err)
    // ... 验证 mock 对象的方法是否被正确调用 ...
}
```
这种架构模式将业务逻辑从持久化细节中解放出来，使得我们的代码更加清晰、健壮，并且极易测试。对于任何需要处理复杂数据交互的应用来说，这都是一个值得投入的最佳实践。
