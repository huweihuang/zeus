[TOC]

#  实例管理

## 1. 创建实例 

> curl -X POST -H 'Content-Type: application/json' -d@create.json x.x.x.x:xxx/api/v1/instance

**Request**

`POST /api/v1/instance`

**Path Parameters**

| Parameter | Description                              |
| --------- | ---------------------------------------- |
| -     | -                   |


**Query Parameters**

- 


**Request Body**

> 创建的request body



```json
{
  "name": "test",
  "namespace": "test",
  "spec": {
    "image": ""
  },
  "shareStorage": true,
  "hostIDs": ["xxx","xxxx"]
}
```



- 200 OK

**Response Body**


```json
{
  "code": 200,
  "message": "create succeed",
  "data": {
    "name": "tnng6qrl9mdt8kjb6rmh"
  }
}
```

- 409 数据冲突

```json
{
  "code": 409,
  "message": "create failed",
  "data": {
    "error": "database conflict"
  }
}
```

- 500 内部错误

```json
{
  "code": 500,
  "message": "create failed",
  "data": {
    "error": "xxxx"
  }
}
```


## 2. 更新实例

> curl -X PUT -H 'Content-Type: application/json' -d@update.json x.x.x.x:xxx/api/v1/instance

**Request**

`POST /api/v1/instance`

**Path Parameters**

| Parameter | Description                              |
| --------- | ---------------------------------------- |
| -     | -                   |


**Query Parameters**

- 


**Request Body**

```json
{
  "name": "test",
  "spec": {
    "image": "xxx"
  },
  "hostIDs": ["xxxx","xxxx"]
}
```



**Response Body**

- 200 OK

```json
{
  "code": 200,
  "message": "create succeed",
  "data": {
    "name": "tnng6qrl9mdt8kjb6rmh"
  }
}
```

- 409 数据冲突

```json
{
  "code": 409,
  "message": "create failed",
  "data": {
    "error": "database conflict"
  }
}
```


- 500 内部错误

```json
{
  "code": 500,
  "message": "create failed",
  "data": {
    "error": "xxxx"
  }
}
```


## 3. 查询实例

> curl -X GET  x.x.x.x:xxx/api/v1/instance?name={name}

**Request**

`GET /api/v1/instance`

**Path Parameters**

| Parameter | Description                              |
| --------- | ---------------------------------------- |
| -      | -                   |


**Request Query**

- name

**Request Body**

```json
{}   
```

**Response Body**

- 200 OK

```json
{
  "code": 200,
  "message": "get succeed",
  "data":{}
}
```


## 4. 删除实例


> curl -X DELETE -H 'Content-Type: application/json' -d@delete.json x.x.x.x:xxx/api/v1/instance

**Request**

`DELETE /api/v1/instance`

**Path Parameters**

| Parameter | Description                              |
| --------- | ---------------------------------------- |
| -      | -                   |


**Request Query**



**Request Body**

```json
{

}   
```

**Response Body**

- 200 OK

```json
{
  "code": 200,
  "message": "delete succeed",
  "data": {
    "name": "xxxxxxxx"
  }
}
```

- 404  未找到

```json
{
  "code": 404,
  "message": "delete failed",
  "data": {
    "error": "host id not fount",
    "podIPs": [
      "x.x.x.x"
    ]
  }
}
```

- 500 内部错误

```json
{
  "code": 500,
  "message": "delete failed",
  "data": {
    "error": "xxxx"
  }
}
```
