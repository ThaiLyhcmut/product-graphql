├── go.mod                  - File cấu hình Go, chứa các thông tin về module và các phụ thuộc.
├── go.sum                  - File chứa các checksum của các module, đảm bảo tính toàn vẹn khi tải xuống.
├── gqlgen.yml              - File cấu hình của gqlgen, chứa các tùy chọn để điều khiển mã nguồn được sinh ra.
├── graph
│   ├── generated           - Thư mục chứa mã runtime được sinh ra, không thay đổi khi bạn tái sinh mã.
│   │   └── generated.go    - File chứa mã runtime đã được sinh ra từ GraphQL.
│   ├── model               - Thư mục chứa tất cả các model GraphQL, bao gồm cả các model đã sinh ra hoặc tự tạo.
│   │   └── models_gen.go   - File chứa các model đã được sinh ra từ GraphQL schema.
│   ├── resolver.go         - File chứa type root resolver của GraphQL. Đây là file bạn tự viết và không được tái sinh lại.
│   ├── schema.graphqls     - Một file chứa GraphQL schema. Bạn có thể chia schema thành nhiều file nếu muốn.
│   └── schema.resolvers.go - File chứa phần cài đặt resolvers cho schema.graphql.
└── server.go               - Điểm vào (entry point) của ứng dụng. Bạn có thể tùy chỉnh file này theo yêu cầu.
