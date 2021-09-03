# REST API CRUD Customer

​	Dalam repo ini menggunakan stack teknologi golang dalam pembuatan REST API dengan menggunkan bantuan *dependencies*: 

- Gofiber untuk web framework
- GORM untuk pemetaan database ke dalam struct di golang
- Go-Validator untuk menambahkan validasi dalam struktur field di struct
- MySQL untuk penyimpanan database

​	Request dengan HTTP Method POST dan struktur Json berupa berikut :

```json
POST /customer 

{
    "name": "Ahmad Irfaan Hibatullah",
    "gender" : "Male",
    "amount" : 5000000,
    "tenor": 6,
    "birth_date": "2001-08-22T00:00:00.000Z",
    "address": "Surabaya",
    "nationality": "Indonesian",
    "ktp": "3578222208960001"
}
```

akan mendapatkan balikan responsen yaitu : 

**Status code : 201**

```json

{
    "data": {
        "id": "231113b1-2954-4015-9023-c82d977aa458",
        "name": "Ahmad Irfaan Hibatullah",
        "gender": "Male",
        "amount": 5000000,
        "tenor": 6,
        "birth_date": "2001-08-22T00:00:00Z",
        "address": "Surabaya",
        "nationality": "Indonesian",
        "address_province": "Jawa Timur",
        "ktp": "3578222208960006",
        "status": "rejected"
    },
    "error": null
}
```

