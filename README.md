# GOCAS (Go Clean Architecture Service)

Dự án này là một hệ thống được thiết kế theo mô hình Clean Architecture trong Golang, bao gồm các service chính như Server (HTTP) và PubSub, được quản lý và điều hướng bằng công cụ giao diện dòng lệnh Cobra (Cobra CLI).

---

## 🚀 Quá trình Xây dựng và Refactor (Nhật ký phát triển)

Dưới đây là chi tiết các vấn đề đã được xử lý và tối ưu trong quá trình cấu trúc lại dự án:

### 1. Xử lý lỗi Import và Quy tắc thư mục `internal` trong Go
**Vấn đề ban đầu:** 
File `cmd/root.go` báo lỗi đỏ không thể import được package `gocas/cmd/pubsub` và `gocas/cmd/server`. Nguyên nhân là do các file code chính bị đặt sâu bên trong thư mục `internal/` (ví dụ: `cmd/pubsub/internal/pubsub.go`). Theo quy tắc bảo mật của trình biên dịch Go, các package bên ngoài không được phép truy cập trực tiếp vào thư mục `internal`.

**Giải pháp (Mô hình Bridge/Cầu nối):**
Để giữ nguyên cấu trúc `internal` (nhằm mục đích đóng gói và bảo vệ logic bên trong) đúng như mong muốn, chúng ta đã áp dụng mô hình "Cầu nối":
- **File Export bên ngoài:** Tạo các file `cmd/pubsub/pubsub.go` và `cmd/server/server.go`. Các file này chỉ làm nhiệm vụ khai báo `cobra.Command` và đóng vai trò như một API công khai để `root.go` có thể gọi tới.
- **Ẩn logic thực thi bên trong:** Toàn bộ code logic lõi (khởi tạo Fiber app, logger, context...) được đặt an toàn bên trong `cmd/pubsub/internal/pubsub.go` và `cmd/server/internal/server.go`. Khi có lệnh chạy, file bên ngoài sẽ uỷ quyền bằng cách gọi hàm `internal.Run()`.
- **Cấu trúc thư mục hiện tại:**
  ```text
  cmd/
   ├── root.go
   ├── pubsub/
   │    ├── pubsub.go           <-- (Cầu nối, export PubSubCmd)
   │    └── internal/
   │         └── pubsub.go      <-- (Core logic xử lý pubsub)
   └── server/
        ├── server.go           <-- (Cầu nối, export ServerCmd)
        └── internal/
             └── server.go      <-- (Core logic xử lý server)
  ```

### 2. Xử lý lỗi Makefile thiếu file cấu hình
**Vấn đề ban đầu:** Khi chạy lệnh khởi tạo qua Make (`make server pubsub`), hệ thống báo lỗi không tìm thấy file `.env`.
**Giải pháp:** Tiến hành clone file `.env.sample` thành `.env` thực tế để ứng dụng có thể đọc các cấu hình môi trường khởi tạo (như FIBER_PORT, LOG_PATH...).

### 3. Cấu hình Redis và Bảo mật (ACL)
Thiết lập file cấu hình cho Redis (`docker/redis/redis.conf`) áp dụng quy tắc ACL (Access Control List) mới của Redis để bảo mật:
- Tạo một user cụ thể `crm-gocas` với đầy đủ quyền (allcommands, allkeys, allchannels).
- Cấu hình bắt buộc phải có mật khẩu truy cập (`requirepass 123456Xyz`) để đảm bảo an toàn nếu triển khai trên môi trường Docker.

---

## 🛠 Hướng dẫn tạo Service mới

Dựa vào cấu trúc hiện tại của dự án (GOCAS - theo Clean Architecture và mô hình Bridge/Cầu nối sử dụng Cobra CLI), để tạo một service mới (ví dụ ta đặt tên service là `worker`), bạn cần thực hiện các bước sau:

### Bước 1: Tạo cấu trúc thư mục cho Service mới
Trong thư mục `cmd/`, bạn tạo thư mục chứa service và thư mục `internal` bên trong nó để chứa file quản lý routes (đúng với quy tắc bảo vệ code hiện tại của dự án).
```text
cmd/
 └── worker/
      ├── worker.go
      └── internal/
           └── routes_worker.go
```

### Bước 2: Tạo file điều hướng (Internal Routes)
Tạo file `cmd/worker/internal/routes_worker.go`. Đây là nơi bạn sẽ gọi các hàm điều hướng (routes) hoặc thiết lập event listener cho service của bạn.

```go
package internal

import (
	"github.com/teoit/gosctx"
)

func RoutesWorker(serviceCtx gosctx.ServiceContext) {
	// Setup worker routes, queues, or cron jobs here
}
```

### Bước 3: Tạo file Cầu nối và Khởi tạo Service (Export)
Tạo file `cmd/worker/worker.go`. File này sử dụng Cobra để tạo ra một Command, cấu hình Context (Gosctx) và đăng ký các Component cần thiết (Fiber, Gorm, Redis, Logger...) y hệt như server.

```go
package worker

import (
	"fmt"
	"gocas/cmd/worker/internal" // Thay 'gocas' bằng tên module trong go.mod của bạn

	"github.com/spf13/cobra"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/fiberapp"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/component/redisc"
	"github.com/teoit/gosctx/configs"
)

var (
	serviceName = "worker-service"
	version     = "1.0.0"
)

func newWorkerServiceCtx() gosctx.ServiceContext {
	return gosctx.NewServiceContext(
		gosctx.WithName(serviceName),
		gosctx.WithComponent(fiberapp.NewFiber(configs.KeyCompFIBER)),
		gosctx.WithComponent(gormc.NewGormDB(configs.KeyCompGorm, "")),
		gosctx.WithComponent(redisc.NewRedisc(configs.KeyCompRedis)),
		gosctx.WithComponent(gosctx.NewAppLoggerDaily(configs.KeyLoggerDaily)),
	)
}

var WorkerCmd = &cobra.Command{
	Use:     "worker",
	Version: version,
	Short:   "Worker service cho hệ thống GOCAS",
	Long:    `Worker service chạy ngầm xử lý các tác vụ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------------------- worker --------------")
		serviceCtx := newWorkerServiceCtx()
		
		logSrv := serviceCtx.MustGet(configs.KeyLoggerDaily).(gosctx.AppLoggerDaily).GetLogger("worker")
		logSrv.Info("------------>> start worker service <<------------------")

		if err := serviceCtx.Load(); err != nil {
			logSrv.Fatal(err)
		}

		// Gọi file điều hướng (Routing)
		internal.RoutesWorker(serviceCtx)

		// Chặn tiến trình không cho thoát
		var forever chan struct{}
		<-forever
	},
}
```

### Bước 4: Đăng ký Service vào Root Command
Mở file `cmd/root.go`, bạn import package service mới tạo và thêm lệnh của nó vào `rootCmd` trong hàm `init()`:

```go
package cmd

import (
	// ... (các import khác)
	"gocas/cmd/worker" // 1. Import package worker mới
)

func init() {
	// ... (các command khác)
	rootCmd.AddCommand(worker.WorkerCmd) // 2. Thêm command vào root
}
```

### Bước 5: Cập nhật Makefile (Tùy chọn)
Để dễ dàng khởi chạy, bạn có thể mở file `Makefile` và thêm một target để chạy service này giống như `server` hay `pubsub`:

```makefile
worker:
	go run main.go worker
```

Sau khi hoàn tất, bạn có thể chạy `make worker` để khởi động service mới.

---

## ⚙️ Hướng dẫn khởi chạy dự án chi tiết

Để chạy dự án mượt mà, bạn cần đảm bảo hệ thống đã cài đặt sẵn **Go**, **Docker**, và **Docker Compose**. Thực hiện lần lượt các bước sau:

### Bước 1: Chuẩn bị biến môi trường
Môi trường mặc định đã được cấu hình trong file `.env.sample`. Bạn cần nhân bản file này ra để tạo file `.env` thực tế:
```bash
cp .env.sample .env
```
*(Lưu ý: File `.env` chứa các thông số kết nối Database và Redis. Nếu bạn đổi mật khẩu trong `docker-compose.yml` hoặc `redis.conf`, hãy nhớ update cả vào file `.env` này nhé).*

### Bước 2: Khởi động Database (Postgres) và Cache (Redis)
Dự án phụ thuộc vào Postgres và Redis. Mình đã cấu hình sẵn mọi thứ trong `docker-compose.yml`. Mở terminal và chạy lệnh sau để khởi động 2 service này ngầm dưới background:
```bash
docker compose up -d
```
Bạn có thể kiểm tra xem DB và Redis đã chạy thành công chưa bằng lệnh `docker compose ps`.

### Bước 3: Cài đặt thư viện Go
Tải toàn bộ các thư viện cần thiết của dự án về máy:
```bash
go mod tidy
```
Đảm bảo lệnh này không báo lỗi gì trước khi đi tiếp.

### Bước 4: Khởi chạy các Service (Server & Pubsub)
Dự án bao gồm 2 service riêng biệt. Bạn cần **mở 2 tab terminal khác nhau** để chạy song song 2 service này, vì chúng đều là các tiến trình chạy liên tục (long-running).

**Terminal 1 - Khởi chạy HTTP Server (Fiber):**
```bash
make server
```
*Kết quả mong đợi: Terminal báo kết nối thành công Database, Redis và hiện logo của Fiber đang lắng nghe ở cổng `3007`.*

**Terminal 2 - Khởi chạy PubSub Worker:**
```bash
make pubsub
```
*Kết quả mong đợi: Terminal in ra dòng `start pubsub service` và lắng nghe các event từ hệ thống.*

---

## 🛠 Công nghệ sử dụng (Tech Stack)
- **Ngôn ngữ:** Golang
- **Framework Web:** [Go Fiber v2](https://gofiber.io/)
- **CLI Tool:** [Cobra](https://github.com/spf13/cobra)
- **Context/DI Management:** Gosctx (quản lý Component Lifecycle: Fiber, Gorm, Redis, Logger...)
- **Database/Cache:** Postgres & Redis (chạy qua Docker)
- **Live Reloading:** [Air](https://github.com/cosmtrek/air)
- **HTML Templating:** [Templ](https://templ.guide)

---

## 🚀 Hướng dẫn chạy môi trường Development (Live Reloading với Air và Templ)

Để giúp quá trình phát triển nhanh hơn, dự án đã được tích hợp **Air** (tự động build lại khi code thay đổi) và **Templ** (tạo giao diện HTML trong Go).

### 1. Cài đặt các công cụ cần thiết (Global)
Bạn cần mở terminal và chạy 2 lệnh sau để cài đặt `air` và `templ` vào máy:
```bash
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
```

### 2. Khởi chạy Server với Live-Reload
Thay vì dùng `make server`, giờ đây bạn chỉ cần gõ lệnh:
```bash
make dev
```
*(Lệnh này sẽ tự động gọi `air`, sau đó `air` sẽ tự gọi `templ generate` mỗi khi file `.templ` thay đổi, tiếp đó là compile code Go và chạy lại server).*

Giờ đây bạn có thể vào sửa giao diện HTML tại `views/users/index.templ` hoặc sửa code logic, server sẽ tự động cập nhật ngay lập tức mà không cần tắt bật lại bằng tay!


---

## Quy trình chạy code 

đầu tiên chạy make server thì đi vô file Makefile, sau đó vô file main.go, sau đó main đi vô root.go, sau đó root.go đi vô init trước, vô init thì có server.go, server.go nó gọi vô routes_server.go, sau đó nó sẽ gọi vô file routes_user_group.go, khai báo composer rồi sau đó vô transport, rồi nhảy vô handler hoặc là api, sau đó nó vô usecase của handler hoặc api, sau đó đi vô repository, sau đó đi qua entity, sau đó nó vô database để lấy dữ liệu 
