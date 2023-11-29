# Pemilu Backend

## Installation

1. go mod tidy
2. go run app/main.go

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

| DB Configuration |
| ---------------- |
| DB_PORT          |
| DB_USERNAME      |
| DB_PASSWORD      |
| DB_DATABASE_NAME |

| Cloudinary Configuration |
| ------------------------ |
| CLOUD_NAME               |
| API_KEY                  |
| API_SECRET               |

| Authorization |
| ------------- |
| SECRET_KEY    |

## API Reference

if you have import to the post man you'll see this

![alt text](https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700769955/ivtphz612yfl2yhf2zmc.png)

## APi Documentation

### Auth

- Register User
- Login User
- List User

#### Register User 👤

- URL : http://localhost:3000/api/v1/register
- Method : POST
- Required Token : No
- Request Body :

```sh
{
    "full_name" : "Jon Doe",
    "alamat" : "Jalan Semangka No 1",
    "jenis_kelamin" : "Laki - Laki",
    "username" : "jons",
    "password" : "jons123",
    "role" : "user"
}
```

- Response Body :

```sh
{
    "status": 201,
    "message": "Success",
    "data": {
        "full_name": "Jon Doe",
        "alamat": "Jalan Semangka No 1",
        "jenis_kelamin": "Laki - Laki",
        "username": "jons",
        "role": "user",
        "password": "$2b$10$ZiJTmgnmA5KWiHk7GQBr8.LR6NCgRJTU0BbRbT.xf/jFQKaplrP0i",
        "id": 4
    }
}
```

#### Login User 👤

- URL : http://localhost:3000/api/v1/login
- Method : POST
- Required Token : No
- Request Body :

```sh
{
    "username" : "jons",
    "password" : "jons123"
}
```

- Response Body :

```sh
{
    "status": 200,
    "message": "Success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjQsInVzZXJuYW1lIjoiam9ucyIsImlhdCI6MTcwMDgxODMwMywiZXhwIjoxNzAwODIxOTAzfQ.r0AV5LRxncmWUw678kjudiwIpySNytm6Jr40-BYPkwg"
}
```

#### List user 👤

- URL : http://localhost:3000/api/v1/list-user
- Method : GET
- Required Token : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "id": 1,
            "full_name": "Joko samudra",
            "alamat": "Jalan Kucing No 1",
            "jenis_kelamin": "Laki - Laki",
            "username": "joko123",
            "role": "admin",
            "password": "$2b$10$8WeZAbYOKkZnxPpf8QGodetWeQ8nue2XONv77CJdEUIVU6n9zK1qi"
        }
    ]
}
```

### Pemilu News (article)

- FindAll Article
- Detail Article
- Add Article
- Update Article
- Delete Article

#### FindAll Article 📰

- URL : http://localhost:3000/api/v1/pemilu-news
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "id": 3,
            "title": "Nomor Urut Paslon sudah di umumkan",
            "author": "Gunawan",
            "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
            "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
            "posted_at": "2023-11-23T18:46:57.990Z"
        }
    ]
}
```

#### Detail Article 📰

- URL : http://localhost:3000/api/v1/pemilu-news/details/3
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success",
    "data": {
        "id": 3,
        "title": "Nomor Urut Paslon sudah di umumkan",
        "author": "Gunawan",
        "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
        "posted_at": "2023-11-23T18:46:57.990Z"
    }
}
```

#### Add Article 📰

- URL : http://localhost:3000/api/v1/pemilu-news
- Method : POST
- Required Token : Yes
- Request Body :

```sh
{
    "title": "Pasangan calon dari Partai PDI yakin menang 1 putaran",
    "author": "Super Admin",
    "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
}
```

- Response Body :

```sh
{
    "status": 201,
    "message": "data created",
    "data": {
        "title": "Pasangan calon dari Partai PDI yakin menang 1 putaran",
        "author": "Super Admin",
        "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
        "posted_at": "2023-11-24T09:49:47.389Z",
        "id": 4,
        "created_at": "2023-11-24T09:49:47.417Z",
        "updated_at": "2023-11-24T09:49:47.417Z"
    }
}
```

#### Update Article 📰

- URL : http://localhost:3000/api/v1/pemilu-news/update/4
- Method : PUT
- Required Token : Yes
- Request Body :

```sh
{
    "title": "Pasangan calon dari Partai PDI yakin menang 1 putaran Kaga BOONG",
    "author": "Super Admin",
    "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
}
```

- Response Body :

```sh
{
    "status": 200,
    "message": "data updated",
    "data": {
        "title": "Pasangan calon dari Partai PDI yakin menang 1 putaran Kaga BOONG",
        "author": "Super Admin",
        "image": "https://images.unsplash.com/photo-1476158085676-e67f57ed9ed7?q=80&w=2344&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
    }
}
```

#### Delete Article 📰

- URL : http://localhost:3000/api/v1/pemilu-news/delete/3
- Method : DELETE
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "data deleted"
}
```

### Paslon

- FindAll Paslon
- Detail Paslon
- Add Paslon
- Update Paslon

#### FindAll Paslon 🤴🫅

- URL : http://localhost:3000/api/v1/paslon
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "3 data found",
    "data": [
         {
            "id": 2,
            "nama": "Prabowo Subianto & Gibran",
            "nomor_urut": "2",
            "visi_misi": "Membantu rakyat miskin",
            "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700715395/n8phga362fbwkrpxbxsi.jpg",
            "partai_pengusung": [
                {
                    "partai_data": {
                        "nama_partai": "Gerindra"
                    }
                },
                {
                    "partai_data": {
                        "nama_partai": "Demokrat"
                    }
                }
            ]
        },
    ]
}
```

#### Detail Paslon 🤴🫅

- URL : http://localhost:3000/api/v1/paslon
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "success get data",
    "data": {
        "id": 2,
        "nama": "Prabowo Subianto & Gibran",
        "nomor_urut": "2",
        "visi_misi": "Membantu rakyat miskin",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700715395/n8phga362fbwkrpxbxsi.jpg",
        "partai_pengusung": [
            {
                "nama_partai": "Gerindra"
            },
            {
                "nama_partai": "Demokrat"
            }
        ]
    }
}
```

#### Add Paslon 🤴🫅

- URL : http://localhost:3000/api/v1/paslon
- Method : POST
- Required Token : Yes
- Request Body :
  ![alt](https://i.imgur.com/7tZlE0d.png)
- Response Body :

```sh
{
    "message": "Success Add",
    "data": {
        "nama": "Firman & Agus",
        "nomor_urut": "4",
        "visi_misi": "Membantu rakyat fakir miskin",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700822817/zppnwaugmwvv4hwlijba.jpg",
        "id": 4,
        "created_at": "2023-11-24T10:46:57.542Z",
        "updated_at": "2023-11-24T10:46:57.542Z"
    }
}
```

#### Update Paslon 🤴🫅

- URL : http://localhost:3000/api/v1/paslon/update/4
- Method : PUT
- Required Token : Yes
- Request Body :
  ![alt](https://i.imgur.com/qvNZKom.png)
- Response Body :

```sh
{
    "message": "Success Update",
    "data": {
        "nama": "Firman & Agus",
        "nomor_urut": "299",
        "visi_misi": "Membantu rakyat miskin",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700823253/fhlwziqj2fpxjyi7bd5s.png"
    }
}
```

### Partai

- FindAll Partai
- Detail Partai
- Add Partai
- Update Partai
- Delete Partai

#### FindAll Partai 🐨

- URL : http://localhost:3000/api/v1/partai
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "4 data found",
    "data": [
        {
            "id": 4,
            "nama": "Demokrat",
            "ketua_umum": "Agus Hari Murti Yudoyono",
            "visi_misi": "Bensin akan murah",
            "alamat": "Salemba Jakarta Pusat",
            "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700822271/fz19tpnq0eu6s5yoa576.svg",
            "created_at": "2023-11-24T10:37:54.026Z",
            "updated_at": "2023-11-24T10:37:54.026Z"
        }
    ]
}
```

#### Detail Partai 🐨

- URL : http://localhost:3000/api/v1/partai/1
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "success get data",
    "data": {
        "id": 1,
        "nama": "PDI Perjuangan",
        "ketua_umum": "Megawati",
        "visi_misi": "Mewujudkan demokrasi indonesia",
        "alamat": "Menteng Jakarta Pusat",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700675276/nxewzarzjqppet9l4yqf.png",
        "created_at": "2023-11-23T18:46:58.063Z",
        "updated_at": "2023-11-23T18:46:58.063Z"
    }
}
```

#### Add Partai 🐨

- URL : http://localhost:3000/api/v1/partai/add
- Method : POST
- Required Token : Yes
- Request Body :
  ![ALt](https://i.imgur.com/PR3RMci.png)
- Response Body :

```sh
{
    "status": 201,
    "message": "Success",
    "data": {
        "nama": "Demokrat",
        "ketua_umum": "Agus Hari Murti Yudoyono",
        "visi_misi": "Bensin akan murah",
        "alamat": "Salemba Jakarta Pusat",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700822271/fz19tpnq0eu6s5yoa576.svg",
        "id": 4,
        "created_at": "2023-11-24T10:37:54.026Z",
        "updated_at": "2023-11-24T10:37:54.026Z"
    }
}
```

#### Update Partai 🐨

- URL : http://localhost:3000/api/v1/partai/update/4
- Method : PUT
- Required Token : Yes
- Request Body :
  ![ALt](https://i.imgur.com/PR3RMci.png)
- Response Body :

```sh
{
    "status": 200,
    "message": "success update",
    "data": {
        "nama": "Demokrat",
        "ketua_umum": "Agus Bambang Pacul",
        "visi_misi": "Menjadikan manusia yang demokratis",
        "alamat": "Kampung Bandan",
        "image": "https://res.cloudinary.com/dp3rsk2xa/image/upload/v1700826893/q6wipd45pqkbidg71ryc.png"
    }
}
```

#### Delete Partai 🐨

- URL : http://localhost:3000/api/v1/partai/delete/5
- Method : DELETE
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success delete",
}
```

### Paslon Partai (relation between paslon and partai)

- Add Paslon Partai

#### Add Paslon Partai 🧑‍💼

- URL : http://localhost:3000/api/v1/paslon-partai
- Method : POST
- Required Token : Yes
- Request Body :

```sh
{
    "paslonId": 2,
    "partaiId": 4
}
```

- Response Body :

```sh
{
    "message": "Success Add",
    "data": {
        "paslonId": 2,
        "partaiId": 4,
        "id": 4
    }
}
```

### Voter Paslon

- List Data Voters
- Vote Paslon
- Count Data Voter

#### List Data Voters 🗳️

- URL : http://localhost:3000/api/v1/voters
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success Get data",
    "data": [
        {
            "nama": "Arif Luthfi Romadhoni",
            "alamat": "Jalan Kucing No 1",
            "jenis_kelamin": "Laki - Laki",
            "paslon_name": "Anis Baswedan & Muhaimin"
        },
        {
            "nama": "Parman Agus",
            "alamat": "Jalan Belimbing No 1",
            "jenis_kelamin": "Laki - Laki",
            "paslon_name": "Prabowo Subianto & Gibran"
        },
        {
            "nama": "Juan",
            "alamat": "Jalan Semangka No 1",
            "jenis_kelamin": "Laki - Laki",
            "paslon_name": "Anis Baswedan & Muhaimin"
        }
    ]
}
```

#### Vote Paslon 🗳️

- URL : http://localhost:3000/api/v1/vote
- Method : POST
- Required Token : Yes
- Request Body :

```sh
{
    "paslonId": 1
}
```

- Response Body :

```sh
{
    {
    "message": "Success",
    "data": {
        "userId": 4,
        "paslonId": 1,
        "id": 17
    }
}
}
```

#### Count Data Voter 🗳️

- URL : http://localhost:3000/api/v1/vote
- Method : GET
- Required Token : Yes
- Request Body : No
- Response Body :

```sh
{
    "status": 200,
    "message": "Success Get data",
    "data": [
       {
            "nomer_urut": "1",
            "paslon_name": "Anis Baswedan & Muhaimin",
            "Pemilih": 3
        },
        {
            "nomer_urut": "2",
            "paslon_name": "Prabowo Subianto & Gibran",
            "Pemilih": 1
        },
        {
            "nomer_urut": "3",
            "paslon_name": "Ganjar Pranowo & Mahfud MD",
            "Pemilih": 0
        }
    ]
}
```