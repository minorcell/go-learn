# 安全库：构建可信的Go应用程序

> 安全不是功能，而是质量属性。本文从威胁建模角度深入分析Go安全生态，通过真实漏洞案例和合规实践，帮你构建安全防线完整的应用程序。

在数据泄露成本平均达到424万美元的今天，安全已经从"可选项"变成"必选项"。Go的安全库生态虽然相对年轻，但已经涵盖了从密码学到身份认证的各个关键领域。

---

## 🛡️ 威胁模型与防护策略

### 常见威胁场景分析

在Go应用中，我们面临的主要威胁包括：

| 威胁类型 | 影响程度 | 常见场景 | 防护优先级 |
|----------|----------|----------|------------|
| **密码泄露** | 🔴 极高 | 数据库泄露、日志记录 | P0 |
| **会话劫持** | 🟠 高 | 中间人攻击、XSS | P0 |
| **注入攻击** | 🟠 高 | SQL注入、命令注入 | P0 |
| **权限绕过** | 🟡 中 | JWT伪造、权限提升 | P1 |
| **敏感信息泄露** | 🟡 中 | 配置文件、错误消息 | P1 |

### 安全架构原则

**纵深防御（Defense in Depth）**：
- 输入验证 → 身份认证 → 授权控制 → 数据加密 → 审计日志

**最小权限原则（Principle of Least Privilege）**：
- 服务账户最小权限
- API访问最小范围
- 数据访问最小集合

---

## 🔐 密码学与加密

### 密码哈希：抵御彩虹表攻击

**错误示例**（易受攻击）：
::: details ❌ 危险：使用MD5或SHA1存储密码
```go
// ❌ 危险：使用MD5或SHA1存储密码
func badPasswordHash(password string) string {
    h := md5.Sum([]byte(password))
    return hex.EncodeToString(h[:])
}
```
:::
**正确实现**：

::: details bcrypt：工业级密码哈希
```go
package security

import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "golang.org/x/crypto/scrypt"
    "time"
)

// bcrypt密码哈希（推荐）
type PasswordHasher struct {
    cost int // bcrypt成本参数
}

func NewPasswordHasher() *PasswordHasher {
    return &PasswordHasher{
        cost: 12, // 平衡安全性和性能，大约250ms
    }
}

func (ph *PasswordHasher) HashPassword(password string) (string, error) {
    // bcrypt自动处理salt生成
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), ph.cost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    
    return string(hashedBytes), nil
}

func (ph *PasswordHasher) VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

// 定时检查成本参数是否需要调整
func (ph *PasswordHasher) BenchmarkCost() {
    start := time.Now()
    ph.HashPassword("test-password")
    duration := time.Since(start)
    
    if duration < 100*time.Millisecond {
        log.Printf("Warning: bcrypt cost too low (took %v), consider increasing", duration)
    }
    if duration > 500*time.Millisecond {
        log.Printf("Warning: bcrypt cost too high (took %v), consider decreasing", duration)
    }
}

// scrypt实现（更灵活的参数控制）
type ScryptHasher struct {
    n      int // CPU/内存成本参数
    r      int // 块大小参数
    p      int // 并行参数
    keyLen int // 输出密钥长度
}

func NewScryptHasher() *ScryptHasher {
    return &ScryptHasher{
        n:      32768,  // 2^15
        r:      8,
        p:      1,
        keyLen: 32,
    }
}

func (sh *ScryptHasher) HashPassword(password string) (string, error) {
    // 生成随机salt
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", fmt.Errorf("failed to generate salt: %w", err)
    }
    
    // 使用scrypt派生密钥
    dk, err := scrypt.Key([]byte(password), salt, sh.n, sh.r, sh.p, sh.keyLen)
    if err != nil {
        return "", fmt.Errorf("failed to derive key: %w", err)
    }
    
    // 组合salt和hash
    combined := append(salt, dk...)
    return base64.StdEncoding.EncodeToString(combined), nil
}

func (sh *ScryptHasher) VerifyPassword(hashedPassword, password string) bool {
    // 解码存储的哈希
    combined, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil || len(combined) < 16 {
        return false
    }
    
    salt := combined[:16]
    storedHash := combined[16:]
    
    // 重新计算哈希
    dk, err := scrypt.Key([]byte(password), salt, sh.n, sh.r, sh.p, sh.keyLen)
    if err != nil {
        return false
    }
    
    // 使用constant-time比较防止时序攻击
    return subtle.ConstantTimeCompare(storedHash, dk) == 1
}

// 密码强度检查
func CheckPasswordStrength(password string) []string {
    var issues []string
    
    if len(password) < 8 {
        issues = append(issues, "密码长度至少8位")
    }
    
    var hasUpper, hasLower, hasDigit, hasSpecial bool
    for _, char := range password {
        switch {
        case char >= 'A' && char <= 'Z':
            hasUpper = true
        case char >= 'a' && char <= 'z':
            hasLower = true
        case char >= '0' && char <= '9':
            hasDigit = true
        case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
            hasSpecial = true
        }
    }
    
    if !hasUpper { issues = append(issues, "缺少大写字母") }
    if !hasLower { issues = append(issues, "缺少小写字母") }
    if !hasDigit { issues = append(issues, "缺少数字") }
    if !hasSpecial { issues = append(issues, "缺少特殊字符") }
    
    return issues
}
```
:::

### 对称加密：数据保护

::: details ❌ 危险：使用AES-GCM加密
```go
package encryption

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

type AESGCMEncryptor struct {
    key []byte
}

func NewAESGCMEncryptor(key []byte) (*AESGCMEncryptor, error) {
    if len(key) != 16 && len(key) != 24 && len(key) != 32 {
        return nil, fmt.Errorf("invalid key size: must be 16, 24, or 32 bytes")
    }
    
    return &AESGCMEncryptor{key: key}, nil
}

func (e *AESGCMEncryptor) Encrypt(plaintext []byte) (string, error) {
    // 创建AES cipher
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", fmt.Errorf("failed to create cipher: %w", err)
    }
    
    // 创建GCM模式
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", fmt.Errorf("failed to create GCM: %w", err)
    }
    
    // 生成随机nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", fmt.Errorf("failed to generate nonce: %w", err)
    }
    
    // 加密数据
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    
    // Base64编码便于存储
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *AESGCMEncryptor) Decrypt(ciphertext string) ([]byte, error) {
    // 解码Base64
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return nil, fmt.Errorf("failed to decode base64: %w", err)
    }
    
    // 创建AES cipher
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }
    
    // 创建GCM模式
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }
    
    // 验证数据长度
    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    
    // 分离nonce和密文
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    
    // 解密数据
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to decrypt: %w", err)
    }
    
    return plaintext, nil
}

// 密钥派生函数（从密码生成密钥）
func DeriveKeyFromPassword(password, salt string) []byte {
    return pbkdf2.Key([]byte(password), []byte(salt), 100000, 32, sha256.New)
}

// 安全的随机密钥生成
func GenerateRandomKey(size int) ([]byte, error) {
    key := make([]byte, size)
    if _, err := rand.Read(key); err != nil {
        return nil, fmt.Errorf("failed to generate random key: %w", err)
    }
    return key, nil
}

// 密钥轮换支持
type RotatingEncryptor struct {
    currentKey []byte
    oldKeys    [][]byte
    encryptor  *AESGCMEncryptor
}

func (re *RotatingEncryptor) Decrypt(ciphertext string) ([]byte, error) {
    // 先尝试当前密钥
    if plaintext, err := re.encryptor.Decrypt(ciphertext); err == nil {
        return plaintext, nil
    }
    
    // 尝试旧密钥
    for _, oldKey := range re.oldKeys {
        if oldEncryptor, err := NewAESGCMEncryptor(oldKey); err == nil {
            if plaintext, err := oldEncryptor.Decrypt(ciphertext); err == nil {
                // 成功解密，考虑重新加密使用新密钥
                log.Printf("Decrypted with old key, consider re-encryption")
                return plaintext, nil
            }
        }
    }
    
    return nil, fmt.Errorf("failed to decrypt with any available key")
}
```
:::

---

## 🎫 身份认证与授权

### JWT：无状态认证的安全实现

::: details ❌ 危险：使用JWT进行身份认证
```go
package auth

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "time"
    
    "github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
    privateKey     *rsa.PrivateKey
    publicKey      *rsa.PublicKey
    issuer         string
    tokenDuration  time.Duration
    refreshDuration time.Duration
}

type Claims struct {
    UserID   string   `json:"user_id"`
    Username string   `json:"username"`
    Roles    []string `json:"roles"`
    jwt.RegisteredClaims
}

func NewJWTManager(issuer string) (*JWTManager, error) {
    // 生成RSA密钥对（生产环境应从安全存储加载）
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, fmt.Errorf("failed to generate RSA key: %w", err)
    }
    
    return &JWTManager{
        privateKey:      privateKey,
        publicKey:       &privateKey.PublicKey,
        issuer:          issuer,
        tokenDuration:   15 * time.Minute,  // 访问令牌短期有效
        refreshDuration: 7 * 24 * time.Hour, // 刷新令牌长期有效
    }, nil
}

func (jm *JWTManager) GenerateTokenPair(userID, username string, roles []string) (accessToken, refreshToken string, err error) {
    now := time.Now()
    
    // 访问令牌
    accessClaims := &Claims{
        UserID:   userID,
        Username: username,
        Roles:    roles,
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    jm.issuer,
            Subject:   userID,
            Audience:  []string{"api"},
            ExpiresAt: jwt.NewNumericDate(now.Add(jm.tokenDuration)),
            NotBefore: jwt.NewNumericDate(now),
            IssuedAt:  jwt.NewNumericDate(now),
            ID:        generateJTI(), // 唯一标识符
        },
    }
    
    accessToken, err = jm.signToken(accessClaims)
    if err != nil {
        return "", "", fmt.Errorf("failed to sign access token: %w", err)
    }
    
    // 刷新令牌（只包含基本信息）
    refreshClaims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            Issuer:    jm.issuer,
            Subject:   userID,
            Audience:  []string{"refresh"},
            ExpiresAt: jwt.NewNumericDate(now.Add(jm.refreshDuration)),
            NotBefore: jwt.NewNumericDate(now),
            IssuedAt:  jwt.NewNumericDate(now),
            ID:        generateJTI(),
        },
    }
    
    refreshToken, err = jm.signToken(refreshClaims)
    if err != nil {
        return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
    }
    
    return accessToken, refreshToken, nil
}

func (jm *JWTManager) signToken(claims *Claims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    
    // 添加额外的头部信息
    token.Header["kid"] = "key-1" // 密钥ID，支持密钥轮换
    
    return token.SignedString(jm.privateKey)
}

func (jm *JWTManager) VerifyToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // 验证签名算法
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        
        // 验证密钥ID（支持密钥轮换）
        kid, ok := token.Header["kid"].(string)
        if !ok || kid != "key-1" {
            return nil, fmt.Errorf("invalid key ID")
        }
        
        return jm.publicKey, nil
    })
    
    if err != nil {
        return nil, fmt.Errorf("failed to parse token: %w", err)
    }
    
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    
    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, fmt.Errorf("invalid claims")
    }
    
    // 额外验证
    if err := jm.validateClaims(claims); err != nil {
        return nil, fmt.Errorf("claims validation failed: %w", err)
    }
    
    return claims, nil
}

func (jm *JWTManager) validateClaims(claims *Claims) error {
    // 验证issuer
    if claims.Issuer != jm.issuer {
        return fmt.Errorf("invalid issuer")
    }
    
    // 验证用户ID格式
    if claims.UserID == "" {
        return fmt.Errorf("missing user ID")
    }
    
    // 验证角色（可选）
    validRoles := map[string]bool{
        "user": true, "admin": true, "moderator": true,
    }
    for _, role := range claims.Roles {
        if !validRoles[role] {
            return fmt.Errorf("invalid role: %s", role)
        }
    }
    
    return nil
}

// 中间件：验证JWT并提取用户信息
func (jm *JWTManager) AuthMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 从Authorization头或Cookie提取token
            token := jm.extractToken(r)
            if token == "" {
                http.Error(w, "Missing authorization token", http.StatusUnauthorized)
                return
            }
            
            // 验证token
            claims, err := jm.VerifyToken(token)
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }
            
            // 将用户信息添加到context
            ctx := context.WithValue(r.Context(), "user", claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func (jm *JWTManager) extractToken(r *http.Request) string {
    // 优先从Authorization头提取
    auth := r.Header.Get("Authorization")
    if strings.HasPrefix(auth, "Bearer ") {
        return strings.TrimPrefix(auth, "Bearer ")
    }
    
    // 从Cookie提取（适用于SPA）
    if cookie, err := r.Cookie("access_token"); err == nil {
        return cookie.Value
    }
    
    return ""
}

// JWT黑名单支持（登出、密钥泄露等场景）
type TokenBlacklist struct {
    redis *redis.Client
}

func (tb *TokenBlacklist) BlacklistToken(jti string, exp time.Time) error {
    ttl := time.Until(exp)
    if ttl <= 0 {
        return nil // 已过期的token无需加入黑名单
    }
    
    return tb.redis.Set(context.Background(), "blacklist:"+jti, "1", ttl).Err()
}

func (tb *TokenBlacklist) IsBlacklisted(jti string) (bool, error) {
    exists, err := tb.redis.Exists(context.Background(), "blacklist:"+jti).Result()
    return exists > 0, err
}
```
:::

### RBAC：基于角色的访问控制

::: details 权限控制系统实现
```go
package rbac

import (
    "fmt"
    "strings"
    "sync"
)

type Permission struct {
    Resource string // 资源类型，如 "user", "order"
    Action   string // 操作类型，如 "read", "write", "delete"
}

type Role struct {
    Name        string
    Permissions []Permission
}

type User struct {
    ID    string
    Roles []string
}

type RBAC struct {
    roles      map[string]*Role
    roleCache  map[string][]Permission // 角色权限缓存
    mu         sync.RWMutex
}

func NewRBAC() *RBAC {
    return &RBAC{
        roles:     make(map[string]*Role),
        roleCache: make(map[string][]Permission),
    }
}

func (rbac *RBAC) AddRole(role *Role) {
    rbac.mu.Lock()
    defer rbac.mu.Unlock()
    
    rbac.roles[role.Name] = role
    rbac.roleCache[role.Name] = role.Permissions
}

func (rbac *RBAC) CheckPermission(userRoles []string, resource, action string) bool {
    rbac.mu.RLock()
    defer rbac.mu.RUnlock()
    
    requiredPerm := Permission{Resource: resource, Action: action}
    
    for _, roleName := range userRoles {
        permissions, exists := rbac.roleCache[roleName]
        if !exists {
            continue
        }
        
        for _, perm := range permissions {
            if rbac.matchPermission(perm, requiredPerm) {
                return true
            }
        }
    }
    
    return false
}

func (rbac *RBAC) matchPermission(granted, required Permission) bool {
    // 支持通配符匹配
    resourceMatch := granted.Resource == "*" || granted.Resource == required.Resource
    actionMatch := granted.Action == "*" || granted.Action == required.Action
    
    return resourceMatch && actionMatch
}

// 权限中间件
func (rbac *RBAC) RequirePermission(resource, action string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 从context获取用户信息
            user, ok := r.Context().Value("user").(*Claims)
            if !ok {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            
            // 检查权限
            if !rbac.CheckPermission(user.Roles, resource, action) {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}

// 动态权限检查（基于资源所有权）
func (rbac *RBAC) CheckResourceOwnership(userID, resourceOwnerID string, userRoles []string, resource, action string) bool {
    // 如果是资源所有者，允许特定操作
    if userID == resourceOwnerID {
        ownerActions := []string{"read", "update"}
        for _, allowedAction := range ownerActions {
            if action == allowedAction {
                return true
            }
        }
    }
    
    // 否则检查角色权限
    return rbac.CheckPermission(userRoles, resource, action)
}

// 预定义角色系统
func (rbac *RBAC) SetupDefaultRoles() {
    // 超级管理员
    rbac.AddRole(&Role{
        Name: "super_admin",
        Permissions: []Permission{
            {Resource: "*", Action: "*"},
        },
    })
    
    // 普通管理员
    rbac.AddRole(&Role{
        Name: "admin",
        Permissions: []Permission{
            {Resource: "user", Action: "*"},
            {Resource: "order", Action: "*"},
            {Resource: "product", Action: "read"},
            {Resource: "product", Action: "update"},
        },
    })
    
    // 普通用户
    rbac.AddRole(&Role{
        Name: "user",
        Permissions: []Permission{
            {Resource: "profile", Action: "read"},
            {Resource: "profile", Action: "update"},
            {Resource: "order", Action: "read"},
            {Resource: "product", Action: "read"},
        },
    })
    
    // 只读用户
    rbac.AddRole(&Role{
        Name: "readonly",
        Permissions: []Permission{
            {Resource: "*", Action: "read"},
        },
    })
}
```
:::

---

## 🔒 输入验证与防护

### SQL注入防护

::: details 安全的数据库操作
```go
package database

import (
    "database/sql"
    "fmt"
    "regexp"
    "strings"
)

// 参数化查询（推荐）
type UserRepository struct {
    db *sql.DB
}

func (ur *UserRepository) GetUserByID(userID string) (*User, error) {
    // ✅ 使用参数化查询防止SQL注入
    query := "SELECT id, username, email FROM users WHERE id = ?"
    
    var user User
    err := ur.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, fmt.Errorf("failed to query user: %w", err)
    }
    
    return &user, nil
}

func (ur *UserRepository) SearchUsers(keyword string) ([]*User, error) {
    // 输入验证：防止恶意输入
    if err := validateSearchKeyword(keyword); err != nil {
        return nil, fmt.Errorf("invalid search keyword: %w", err)
    }
    
    // 使用LIKE查询时的安全处理
    query := "SELECT id, username, email FROM users WHERE username LIKE ? ESCAPE '\\' LIMIT 100"
    safeKeyword := escapeLikePattern(keyword)
    
    rows, err := ur.db.Query(query, "%"+safeKeyword+"%")
    if err != nil {
        return nil, fmt.Errorf("failed to search users: %w", err)
    }
    defer rows.Close()
    
    var users []*User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
            return nil, fmt.Errorf("failed to scan user: %w", err)
        }
        users = append(users, &user)
    }
    
    return users, nil
}

// 输入验证函数
func validateSearchKeyword(keyword string) error {
    if len(keyword) == 0 {
        return fmt.Errorf("keyword cannot be empty")
    }
    
    if len(keyword) > 50 {
        return fmt.Errorf("keyword too long")
    }
    
    // 检查是否包含可疑字符
    suspiciousPatterns := []string{
        "--", "/*", "*/", "xp_", "sp_", "UNION", "SELECT", "INSERT", "UPDATE", "DELETE",
    }
    
    upperKeyword := strings.ToUpper(keyword)
    for _, pattern := range suspiciousPatterns {
        if strings.Contains(upperKeyword, pattern) {
            return fmt.Errorf("keyword contains suspicious pattern: %s", pattern)
        }
    }
    
    return nil
}

// LIKE模式转义
func escapeLikePattern(pattern string) string {
    // 转义LIKE特殊字符
    pattern = strings.ReplaceAll(pattern, "\\", "\\\\")
    pattern = strings.ReplaceAll(pattern, "%", "\\%")
    pattern = strings.ReplaceAll(pattern, "_", "\\_")
    return pattern
}

// 白名单验证（用于表名、列名等不能参数化的场景）
func validateTableName(tableName string) error {
    // 只允许字母、数字和下划线
    validTableName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
    if !validTableName.MatchString(tableName) {
        return fmt.Errorf("invalid table name")
    }
    
    // 白名单检查
    allowedTables := map[string]bool{
        "users": true, "orders": true, "products": true,
    }
    
    if !allowedTables[tableName] {
        return fmt.Errorf("table not allowed: %s", tableName)
    }
    
    return nil
}
```
:::

### XSS防护与内容安全

::: details 输出编码和内容安全策略
```go
package security

import (
    "html"
    "html/template"
    "net/url"
    "regexp"
    "strings"
)

// HTML输出编码
type SafeRenderer struct {
    templates *template.Template
}

func NewSafeRenderer() *SafeRenderer {
    // 创建带有安全函数的模板
    funcMap := template.FuncMap{
        "safeHTML": func(s string) template.HTML {
            // 只有明确标记为安全的内容才使用此函数
            return template.HTML(sanitizeHTML(s))
        },
        "safeURL": func(s string) template.URL {
            return template.URL(sanitizeURL(s))
        },
        "safeJS": func(s string) template.JS {
            return template.JS(sanitizeJS(s))
        },
    }
    
    tmpl := template.New("").Funcs(funcMap)
    
    return &SafeRenderer{templates: tmpl}
}

// HTML内容净化
func sanitizeHTML(input string) string {
    // 简单的HTML标签白名单
    allowedTags := map[string]bool{
        "p": true, "br": true, "strong": true, "em": true,
    }
    
    // 移除所有不在白名单中的标签
    tagRegex := regexp.MustCompile(`</?([a-zA-Z]+)[^>]*>`)
    return tagRegex.ReplaceAllStringFunc(input, func(tag string) string {
        tagNameRegex := regexp.MustCompile(`</?([a-zA-Z]+)`)
        matches := tagNameRegex.FindStringSubmatch(tag)
        if len(matches) > 1 && allowedTags[strings.ToLower(matches[1])] {
            return tag
        }
        return "" // 移除不允许的标签
    })
}

// URL净化
func sanitizeURL(input string) string {
    // 只允许HTTP和HTTPS协议
    u, err := url.Parse(input)
    if err != nil {
        return ""
    }
    
    if u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "" {
        return ""
    }
    
    return u.String()
}

// JavaScript净化（基础版）
func sanitizeJS(input string) string {
    // 移除潜在危险的JavaScript代码
    dangerousPatterns := []string{
        "eval(", "Function(", "setTimeout(", "setInterval(",
        "document.cookie", "document.write", "innerHTML",
    }
    
    for _, pattern := range dangerousPatterns {
        input = strings.ReplaceAll(input, pattern, "")
    }
    
    return input
}

// 内容安全策略(CSP)中间件
func CSPMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 设置严格的CSP策略
            csp := strings.Join([]string{
                "default-src 'self'",
                "script-src 'self' 'unsafe-inline'", // 生产环境应避免unsafe-inline
                "style-src 'self' 'unsafe-inline'",
                "img-src 'self' data: https:",
                "font-src 'self'",
                "connect-src 'self'",
                "media-src 'none'",
                "object-src 'none'",
                "child-src 'none'",
                "frame-ancestors 'none'",
                "base-uri 'self'",
                "form-action 'self'",
            }, "; ")
            
            w.Header().Set("Content-Security-Policy", csp)
            
            // 其他安全头
            w.Header().Set("X-Content-Type-Options", "nosniff")
            w.Header().Set("X-Frame-Options", "DENY")
            w.Header().Set("X-XSS-Protection", "1; mode=block")
            w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
            
            next.ServeHTTP(w, r)
        })
    }
}

// API输入验证
type InputValidator struct {
    maxFieldLength int
    allowedFields  map[string]bool
}

func NewInputValidator() *InputValidator {
    return &InputValidator{
        maxFieldLength: 1000,
        allowedFields: map[string]bool{
            "username": true, "email": true, "password": true,
        },
    }
}

func (iv *InputValidator) ValidateUserInput(input map[string]string) error {
    for field, value := range input {
        // 检查字段是否允许
        if !iv.allowedFields[field] {
            return fmt.Errorf("field not allowed: %s", field)
        }
        
        // 检查长度
        if len(value) > iv.maxFieldLength {
            return fmt.Errorf("field %s too long", field)
        }
        
        // 字段特定验证
        if err := iv.validateField(field, value); err != nil {
            return fmt.Errorf("validation failed for %s: %w", field, err)
        }
    }
    
    return nil
}

func (iv *InputValidator) validateField(field, value string) error {
    switch field {
    case "email":
        emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
        if !emailRegex.MatchString(value) {
            return fmt.Errorf("invalid email format")
        }
    case "username":
        usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
        if !usernameRegex.MatchString(value) {
            return fmt.Errorf("invalid username format")
        }
    }
    
    return nil
}
```
:::

---

## 📋 安全合规检查清单

### OWASP Top 10 防护状况

✅ **A01 - 权限控制失效**：
- [ ] 实现基于角色的访问控制(RBAC)
- [ ] 验证用户权限在每个请求
- [ ] 最小权限原则
- [ ] 定期审计权限分配

✅ **A02 - 加密机制失效**：
- [ ] 使用强加密算法(AES-256, RSA-2048+)
- [ ] 实现正确的密钥管理
- [ ] 敏感数据传输加密(TLS 1.3)
- [ ] 密码使用bcrypt或scrypt哈希

✅ **A03 - 注入攻击**：
- [ ] 使用参数化查询
- [ ] 输入验证和净化
- [ ] 使用ORM框架
- [ ] 定期代码审计

✅ **A04 - 不安全设计**：
- [ ] 实施威胁建模
- [ ] 安全架构评审
- [ ] 纵深防御策略
- [ ] 安全测试集成到CI/CD

### 安全监控指标

```yaml
安全监控指标:
  认证失败率: < 5%
  异常登录地理位置: 监控并告警
  权限提升尝试: 零容忍
  SQL注入尝试: 自动阻止
  XSS攻击尝试: 自动阻止
  暴力破解: 自动限流
  
日志记录要求:
  认证事件: 必须记录
  授权决策: 必须记录
  敏感操作: 必须记录
  错误信息: 脱敏记录
```

### 应急响应流程

::: details 安全事件响应
```go
// 安全事件响应
type SecurityIncident struct {
    ID          string
    Type        string    // 事件类型
    Severity    string    // 严重程度
    UserID      string    // 涉及用户
    Description string    // 事件描述
    Timestamp   time.Time
    Status      string    // 处理状态
}

func (si *SecurityIncident) HandleDataBreach() {
    // 1. 立即隔离受影响系统
    si.isolateAffectedSystems()
    
    // 2. 通知安全团队
    si.notifySecurityTeam()
    
    // 3. 收集证据
    si.collectEvidence()
    
    // 4. 通知受影响用户（72小时内）
    si.notifyAffectedUsers()
    
    // 5. 监管报告（GDPR等要求）
    si.reportToRegulators()
}
```
:::

安全是一个持续的过程，而不是一次性的任务。定期进行安全评估、漏洞扫描和渗透测试，保持对新威胁的敏感度，是构建安全应用的关键。记住：**安全的成本总是小于数据泄露的代价**。
