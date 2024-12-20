# go-SOLID
The switch from OOP Languages to golang made me look to find a way to create a basic WEB Template in a SOLID Clean Code Oop Fashion that i can use as starting point.

The MiddlewareRunner Concept is just a tryout. For Production where will be a authRouteMiddlewareRunner and a
normal Middleware Runner. So i can switch without breaking too much. I hope :-D

```bash
go run . -h localhost:8888
```

```bash
curl -X GET localhost:8888/api/v1

2024/12/21 00:18:35 Running Server: localhost:8888
2024/12/21 00:18:49 DebugMiddleware: GET Path: /api/v1
2024/12/21 00:18:49 LoginMiddleware: GET Path: /api/v1
2024/12/21 00:18:49 HeaderMiddleware: GET Path: /api/v1
2024/12/21 00:18:49 Method: GET Path: /api/v1
2024/12/21 00:18:49 AfterVisitRouteMiddleware: GET Path: /api/v1
```