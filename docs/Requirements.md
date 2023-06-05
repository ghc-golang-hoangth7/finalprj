## Khách hàng là một hãng hàng không đang cần viết một ứng dụng giả lập với nội dung như sau:

### Xây dựng hệ thống phần mềm đáp ứng các yêu cầu nghiệp vụ sau
> #### Là người sử dụng dịch vụ của sân bay, người dùng mong muốn có các API endpoint để thực hiện việc đặt chỗ trên máy bay
> - Trước khi đặt chỗ cần kiểm tra xem máy bay còn chỗ hay không
> - Người dùng có thể search chuyến bay theo giờ, theo địa điểm (bao gồm điểm đi và điểm đến)
> - Người dùng muốn biết có các chuyến bay nào hợp lệ và trống bao nhiêu ghế
> - Khi người dùng đặt chỗ thì chuyến bay đó sẽ giảm ~~1 chỗ trống~~ => cho tùy chỉnh số ghế muốn đặt
> - Người dùng chỉ có thể đặt chỗ trước ít nhất 45 phút khi máy bay khởi hành

> #### Là người quản lý lên lịch các chuyến bay, người dùng mong muốn
> - Xem danh sách các máy bay và thông tin của máy bay
> - Đã có chuyến bay được lên lịch hay chưa, ngày giờ, điểm đi điểm đến của máy bay
> - Trạng thái máy bay: đang dọn dẹp, sửa chữa, sẵn sàng để bay
> - Lên lịch cho chuyến bay với đầu vào bao gồm
>   - Số hiệu máy bay
>   - Điểm đi
>   - Điểm đến
>   - Giờ xuất phát
>   - Giờ dự kiến sẽ đến điểm đến

> #### Là người quản lý công tác hậu cần ở sân bay, người dùng mong muốn
> - Cập nhật trạng thái máy bay
> - Thêm/Cập nhật/ máy bay với các thông tin cơ bản của máy bay
> - Số hiệu => unique & un-updatable
> - Tổng số chỗ
> - Trạng thái: đang dọn dẹp, sửa chữa, sẵn sàng để bay

> #### Note (service call service here, do not make call from API to server to handle logic below)
> - Không thể xóa máy bay khi đã có chuyến bay được lên lịch
> - Khi tạo chuyến bay cần lấy thông tin mới nhất về máy bay trước khi tạo (số ghế)

> #### Các yêu cầu project
> - Yêu cầu sử dụng GraphQL với API endpoint cho các client
> - Yêu cầu sử dụng gRPC cho giao tiếp giữa các services
#### Gợi ý về cấu trúc thư mục code

```
.
├── grpc
    ├── <your service name>
        ├──handlers                         # service logic handling
        ├──models
        ├──repository
        ├──requests
        ├──responses
        ├──main.go                          # main file of service, contain startup, load config, etc
├──config                                   # Store config
├──logger                                   # Logging
├──pb                                       # grpc
├──proto                                    # Protobuff file
├──clients                                  # API endpoint
    ├──rest (if you cannot write graph)     # Rest Endpoint
    ├──graphql                              # GraphQL endpoint
        ├──cmd                              # main files to start graphql server
        ├──generated                        # Generated files
        ├──models                           # Other model files
        ├──resolver                         # GraphQL resolver
        ├──schemas                          # GraphQL schemas
        ├──services                         # Comunicate with services

```

##### Lưu ý trong bài lab thì toàn bộ code store chung nhưng trên thực tế source code của các services/ Graph API là độc lập, hạn chế import code vượt đường bao
        - Ví dụ tầng clients chỉ có thể giao tiếp với tầng service thông qua call grpc
        - Các thành phần config, logger, proto thường sẽ được build thành module và có thể re-use
        - Thành phần pb là phần code gen tùy theo ngôn ngữ từ file proto
##### Một số thư viện cơ bản
- gRPC [gRPC-Go](https://pkg.go.dev/google.golang.org/grpc)
- Protobuff [Go support for Protocol Buffers](https://pkg.go.dev/google.golang.org/protobuf)


> #### Các phần nâng cao (không bắt buộc)
> - Sử dụng proto để định nghĩa các cấu trúc dữ liệu cần dùng
> - Xử lý vấn đề race condition khi có nhiều người dùng thực hiện book chuyến bay (nhiều hơn tổng số ghế còn lại, và tất cả yêu cầu đến cùng 1 thời điểm)
> - Tất cả các API endpoint cho client được expose ra internel với 1 địa chỉ duy nhất (Tham khảo keyword Apollo Federation)
> - Thay thế việc sử dụng gRPC giữa các services bằng event driven (nếu cần và giải thích tại sao)
