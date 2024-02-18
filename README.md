# Fiber with Auth

[Postman collection](https://www.getpostman.com/collections/c862d012d5dcf50326f7)

## Endpoints


Berikut adalah daftar URL API yang terdapat dalam kode:

Hello World Endpoint:

Method: GET
URL: /api/
Handler: handler.Hello
Deskripsi: Endpoint ini memberikan respons "Hello, World!" saat diakses dengan metode GET.
Login Endpoint:

Method: POST
URL: /api/auth/login
Handler: handler.Login
Deskripsi: Endpoint ini digunakan untuk melakukan proses login.
User Endpoints:

Get User by ID:
Method: GET
URL: /api/user/:id
Handler: handler.GetUser
Deskripsi: Endpoint ini digunakan untuk mendapatkan informasi pengguna berdasarkan ID.
Create User:
Method: POST
URL: /api/user/
Handler: handler.CreateUser
Deskripsi: Endpoint ini digunakan untuk membuat pengguna baru.
Update User:
Method: PATCH
URL: /api/user/:id
Middleware: middleware.Protected()
Handler: handler.UpdateUser
Deskripsi: Endpoint ini digunakan untuk memperbarui informasi pengguna berdasarkan ID. Memerlukan otentikasi.
Delete User:
Method: DELETE
URL: /api/user/:id
Middleware: middleware.Protected()
Handler: handler.DeleteUser
Deskripsi: Endpoint ini digunakan untuk menghapus pengguna berdasarkan ID. Memerlukan otentikasi.
Product Endpoints:

Get All Products:
Method: GET
URL: /api/product/
Handler: handler.GetAllProducts
Deskripsi: Endpoint ini digunakan untuk mendapatkan semua produk yang tersedia.
Get Product by ID:
Method: GET
URL: /api/product/:id
Handler: handler.GetProduct
Deskripsi: Endpoint ini digunakan untuk mendapatkan informasi produk berdasarkan ID.
Create Product:
Method: POST
URL: /api/product/
Middleware: middleware.Protected()
Handler: handler.CreateProduct
Deskripsi: Endpoint ini digunakan untuk membuat produk baru. Memerlukan otentikasi.
Delete Product:
Method: DELETE
URL: /api/product/:id
Middleware: middleware.Protected()
Handler: handler.DeleteProduct
Deskripsi: Endpoint ini digunakan untuk menghapus produk berdasarkan ID. Memerlukan otentikasi.