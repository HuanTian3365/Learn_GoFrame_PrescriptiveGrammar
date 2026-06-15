# 新 GoFrame 项目 CRUD 风格训练任务文档

## Summary
目标是在一个全新的 GoFrame v2 项目中，从零实现一套“类似当前项目风格”的增删改查开发模式。该训练不依赖原项目代码，不复制原项目目录或业务，只学习它的开发思想：清晰分层、统一响应、service 注册、logic 承载业务、dao/model 管理数据、多租户和权限预留。

练习项目：`gf-crud-style-practice`  
练习业务：商品管理  
练习表：`shop_product`  
接口前缀：`/api/v1/product`

最终你需要能独立完成一套标准 CRUD：

```text
api 定义接口
 -> router 绑定 controller
 -> controller 调 service
 -> service 暴露接口
 -> logic 实现业务
 -> dao/model 操作数据库
 -> middleware 统一响应
```

## Key Tasks

### 任务 1：创建新 GoFrame 项目并跑通最小接口
创建一个全新的 GoFrame v2 项目，不从原项目复制代码。

完成内容：

- 初始化 GoFrame 项目。
- 配置 `config.yaml`。
- 启动 HTTP 服务。
- 新增 `/api/v1/health` 接口。
- 接口返回统一 JSON 格式。

推荐响应格式：

```json
{
  "code": 0,
  "message": "",
  "data": {}
}
```

验收材料：

- 项目目录树。
- `main.go`
- HTTP 启动代码。
- health 接口代码。
- health 请求结果。

通过标准：

- 项目能启动。
- `/api/v1/health` 能访问。
- 返回格式统一。

### 任务 2：建立 CRUD 分层结构
在新项目中建立一套固定分层，不要求完全等同原项目，但要体现同类风格。

建议目录：

```text
api/product/v1
internal/controller/product
internal/logic/product
internal/service
internal/model
internal/dao
internal/router
internal/middleware
internal/response
```

完成内容：

- 建立 `product` 模块。
- `api` 只定义请求和响应结构。
- `controller` 只做参数转发和返回组装。
- `service` 只定义接口和注册方法。
- `logic` 写实际业务。
- `dao/model` 负责数据库相关结构。

验收材料：

- 目录树。
- 每层各一个最小文件。
- 说明一次请求会经过哪些层。

通过标准：

- 没有把 SQL 写进 controller。
- 没有让 api 层依赖 logic。
- 每层职责清楚。

### 任务 3：实现商品表和基础 DAO
创建商品表：

```sql
CREATE TABLE shop_product (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  product_name VARCHAR(100) NOT NULL,
  product_code VARCHAR(100) NOT NULL,
  price DECIMAL(10,2) NOT NULL DEFAULT 0,
  status TINYINT NOT NULL DEFAULT 1,
  remark VARCHAR(255),
  tenant_id VARCHAR(64),
  created_by BIGINT,
  created_at DATETIME,
  updated_by BIGINT,
  updated_at DATETIME
);
```

完成内容：

- 接入数据库。
- 使用 GoFrame ORM 查询 `shop_product`。
- 建立 `dao.ShopProduct` 或等价 DAO。
- 建立商品实体模型。
- logic 中通过 DAO 查询数据库，不直接散落表名。

验收材料：

- 建表 SQL。
- DAO 代码。
- Entity/Model 代码。
- 一次数据库查询结果。

通过标准：

- 数据访问入口统一。
- 字段命名清晰。
- 后续 CRUD 能复用 DAO。

### 任务 4：实现 List 查询
实现接口：

```text
GET /api/v1/product/list
```

支持参数：

- `page`
- `pageSize`
- `productName`
- `productCode`
- `status`

返回：

```json
{
  "items": [],
  "page": 1,
  "pageSize": 10,
  "total": 0
}
```

完成内容：

- `api` 定义 `ProductListReq` / `ProductListRes`。
- `controller` 调 `service.Product().List(...)`。
- `service` 定义 `List` 接口。
- `logic` 实现分页、筛选、排序。
- 默认按 `id desc` 排序。

验收材料：

- List 的 api/controller/service/logic/model 代码。
- 无条件查询结果。
- 带商品名称筛选结果。
- 带状态筛选结果。

通过标准：

- 分页正确。
- 查询条件在 logic 层。
- 返回结构稳定。

### 任务 5：实现 View 详情
实现接口：

```text
GET /api/v1/product/view?id=1
```

完成内容：

- 根据主键查询商品详情。
- 不存在时返回业务错误。
- 返回商品完整信息。

验收材料：

- View 相关代码。
- 查询存在数据的结果。
- 查询不存在数据的结果。

通过标准：

- 错误经过统一响应。
- controller 不处理数据库错误细节。
- logic 负责业务判断。

### 任务 6：实现 Edit 新增和修改
实现接口：

```text
POST /api/v1/product/edit
```

规则：

- `id = 0` 或不传：新增。
- `id > 0`：修改。
- `productCode` 必须唯一。
- 新增时写入 `createdBy`、`createdAt`。
- 修改时写入 `updatedBy`、`updatedAt`。
- 使用事务。
- 只允许写入白名单字段。

完成内容：

- `ProductEditReq`
- `ProductEditParam`
- 唯一性校验方法。
- 新增逻辑。
- 修改逻辑。
- 事务包裹。

验收材料：

- Edit 完整代码。
- 新增成功结果。
- 修改成功结果。
- 重复 `productCode` 的失败结果。
- 数据库记录截图或查询结果。

通过标准：

- 新增和修改逻辑分支清楚。
- 不会覆盖创建字段。
- 唯一性校验有效。
- 事务使用正确。

### 任务 7：实现 Delete 删除
实现接口：

```text
POST /api/v1/product/delete
```

请求参数：

```json
{
  "id": 1
}
```

完成内容：

- 根据主键删除商品。
- 删除不存在的数据要返回明确错误。
- 先实现物理删除。

验收材料：

- Delete 相关代码。
- 删除成功结果。
- 删除不存在数据的结果。
- 删除后 List 查询结果。

通过标准：

- 删除前有存在性判断。
- 错误响应统一。
- 删除逻辑只在 logic 层。

### 任务 8：加入统一响应中间件
实现统一响应能力，让业务方法只返回 `res, err`。

完成内容：

- 成功时统一输出：
  ```json
  {
    "code": 0,
    "message": "",
    "data": {}
  }
  ```
- 失败时统一输出：
  ```json
  {
    "code": 500,
    "message": "错误信息",
    "data": null
  }
  ```
- controller 不直接调用 `WriteJson`。

验收材料：

- middleware 代码。
- router 挂载代码。
- 成功响应示例。
- 失败响应示例。

通过标准：

- 所有 CRUD 接口响应格式一致。
- 错误不在 controller 中手动包装。
- 中间件职责清楚。

### 任务 9：加入简化权限码
给接口增加权限标记，但不要求接入真实菜单系统。

权限码：

```text
product:list
product:view
product:edit
product:delete
```

完成内容：

- 在 `api` 的 `g.Meta` 中声明权限码。
- 实现 Auth 中间件。
- 请求头必须携带 `Authorization`。
- 用内存数组模拟当前用户权限。
- 没有权限时拒绝访问。

验收材料：

- `g.Meta` 示例。
- Auth 中间件代码。
- 有权限请求结果。
- 无权限请求结果。
- 无 token 请求结果。

通过标准：

- 权限码跟接口绑定。
- 权限判断在中间件。
- 业务 logic 不关心鉴权。

### 任务 10：加入简化多租户过滤
给商品表增加租户意识，模拟真实后台项目的查询风格。

完成内容：

- 模拟当前租户：
  ```go
  GetTenantId(ctx) string
  ```
- List/View/Edit/Delete 默认只操作当前租户数据。
- 新增时自动写入 `tenant_id`。
- 封装一个统一 Model 方法：
  ```go
  func (s *sProduct) Model(ctx context.Context) *gdb.Model
  ```
- 不允许每个查询随手写散乱的租户条件。

验收材料：

- 租户上下文代码。
- Product 的 `Model(ctx)` 方法。
- 两个租户的数据准备 SQL。
- 当前租户只能查到自己数据的结果。

通过标准：

- 多租户过滤集中处理。
- 新增数据带租户。
- 不会串租户查询或删除。

### 任务 11：加入关联查询
新增商品分类表：

```sql
CREATE TABLE shop_category (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  category_name VARCHAR(100) NOT NULL,
  tenant_id VARCHAR(64)
);
```

商品表增加：

```sql
category_id BIGINT
```

List 返回 `categoryName`。

完成内容：

- 商品列表关联分类表。
- 支持 `categoryId` 过滤。
- 返回模型包含 `categoryName`。
- 编辑参数不允许传 `categoryName`。

验收材料：

- 分类表 SQL。
- ProductListParam。
- ProductListModel。
- List 关联查询代码。
- 带分类筛选的请求结果。

通过标准：

- 查询模型和编辑模型不混用。
- 关联字段只出现在返回模型。
- 连表逻辑在 logic 层。

### 任务 12：最终自检
完成内容：

- 跑通全部 CRUD。
- 运行：
  ```bash
  go test ./...
  go vet ./...
  ```
- 写一份简短说明，描述你的项目分层和请求链路。

验收材料：

- `go test ./...` 输出。
- `go vet ./...` 输出。
- 项目目录树。
- 请求链路说明。
- 你认为自己掌握的 3 个点。
- 你还不确定的 3 个点。

通过标准：

- 项目可运行。
- CRUD 完整。
- 统一响应、权限、多租户都可解释。
- 能独立说明每层职责。

## Test Plan
每个任务完成后至少验证：

- 服务能启动。
- 接口能访问。
- 成功和失败响应格式一致。
- controller 不写数据库逻辑。
- logic 不直接处理 HTTP 输出。
- 新增、查询、修改、删除均可重复验证。
- 权限失败和租户隔离场景可复现。

## Assumptions
- 这是一个全新 GoFrame 项目，不依赖原 IOTLCM 项目。
- 训练目标是学会“这种风格的 CRUD”，不是复制原项目代码。
- 业务选择商品管理只是练习载体，可替换成设备、客户、订单等普通业务。
- 权限、用户、租户先用模拟数据，重点是掌握代码结构和链路。
- 每完成一个任务，你把代码片段、目录树和请求结果发给我，我负责校验风格和指出问题。
