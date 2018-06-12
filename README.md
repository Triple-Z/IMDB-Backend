# IMDB Backend

数据库课设 —— IMDB 查找系统（后端仓库）

<!-- TOC -->

- [IMDB Backend](#imdb-backend)
    - [API References](#api-references)
        - [Titles](#titles)
            - [List titles](#list-titles)
                - [Parameters](#parameters)
                - [Response](#response)
            - [Get a title](#get-a-title)
                - [Response](#response-1)
            - [Create a title](#create-a-title)
                - [Parameters](#parameters-1)
                - [Response](#response-2)
            - [Update a title](#update-a-title)
                - [Parameters](#parameters-2)
                - [Response](#response-3)
            - [Delete a title](#delete-a-title)
                - [Response](#response-4)

<!-- /TOC -->

## API References

API root: http://localhost:2333

### Titles

#### List titles

```api
GET /api/v1/titles
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`page`| `int` | 当前请求的页号 |

> 若缺少该字段，则默认返回第一页的信息。

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 20,
    "cur_page": 1,
    "data": [
        {
            "Id": {
                "Int64": 1,
                "Valid": true
            },
            "TConst": {
                "String": "tt0000001",
                "Valid": true
            },
            "TitleType": {
                "String": "short",
                "Valid": true
            },
            "PrimaryTitle": {
                "String": "Carmencita",
                "Valid": true
            },
            "OriginalTitle": {
                "String": "Carmencita",
                "Valid": true
            },
            "IsAdult": {
                "Bool": false,
                "Valid": true
            },
            "StartYear": {
                "Int64": 1894,
                "Valid": true
            },
            "EndYear": {
                "Int64": 0,
                "Valid": false
            },
            "RuntimeMinutes": {
                "Int64": 1,
                "Valid": true
            },
            "Genres": {
                "String": "Documentary,Short",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T00:34:03Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T00:34:03Z",
                "Valid": true
            }
        },
        ...  // Some data have been omitted
        {
            "Id": {
                "Int64": 20,
                "Valid": true
            },
            "TConst": {
                "String": "tt0000020",
                "Valid": true
            },
            "TitleType": {
                "String": "short",
                "Valid": true
            },
            "PrimaryTitle": {
                "String": "The Derby 1895",
                "Valid": true
            },
            "OriginalTitle": {
                "String": "The Derby 1895",
                "Valid": true
            },
            "IsAdult": {
                "Bool": false,
                "Valid": true
            },
            "StartYear": {
                "Int64": 1895,
                "Valid": true
            },
            "EndYear": {
                "Int64": 0,
                "Valid": false
            },
            "RuntimeMinutes": {
                "Int64": 1,
                "Valid": true
            },
            "Genres": {
                "String": "Documentary,Short,Sport",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T00:34:03Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T00:34:03Z",
                "Valid": true
            }
        }
    ],
    "next_page": 2,
    "prev_page": -1,
    "server_time": "2018-06-12T20:36:20.989811937+08:00",
    "start_id": 0,
    "total_page": 251881
}
```

#### Get a title

```api
GET /api/v1/titles/:id
```

##### Response

```text
Status: 200 OK
```

```json
{
    "data": {
        "Id": {
            "Int64": 1,
            "Valid": true
        },
        "TConst": {
            "String": "tt0000001",
            "Valid": true
        },
        "TitleType": {
            "String": "short",
            "Valid": true
        },
        "PrimaryTitle": {
            "String": "Carmencita",
            "Valid": true
        },
        "OriginalTitle": {
            "String": "Carmencita",
            "Valid": true
        },
        "IsAdult": {
            "Bool": false,
            "Valid": true
        },
        "StartYear": {
            "Int64": 1894,
            "Valid": true
        },
        "EndYear": {
            "Int64": 0,
            "Valid": false
        },
        "RuntimeMinutes": {
            "Int64": 1,
            "Valid": true
        },
        "Genres": {
            "String": "Documentary,Short",
            "Valid": true
        },
        "CreateDate": {
            "Time": "2018-06-12T00:34:03Z",
            "Valid": true
        },
        "LastUpdated": {
            "Time": "2018-06-12T00:34:03Z",
            "Valid": true
        }
    },
    "server_time": "2018-06-12T20:41:17.100390307+08:00"
}
```

#### Create a title

```api
POST /api/v1/titles
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`tconst`| `string`| **必填**. 影片唯一标识号|
|`title_type`| `string`| **必填**. 影片类型|
|`primary_title`| `string`| **必填**. 影片英文名称|
|`original_title`| `string`| 影片源名称|
|`is_adult`| `bool`| **必填**. 成人分类标识 |
|`start_year`| `int`| 影片/电视剧开始年份|
|`end_year`| `int`| 电视剧结束年份|
|`runtime_minutes`| `int`| 影片持续时间|
|`genres`| `string`| 影片分类（最多三个）|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。
 
##### Response

```text
Status: 200 OK
```

```json
{
    "insert_id": 5037605,
    "server_time": "2018-06-12T20:57:10.172114944+08:00"
}
```

#### Update a title

```api
PUT /api/v1/titles/:id
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`tconst`| `string`| **必填**. 影片唯一标识号|
|`title_type`| `string`| **必填**. 影片类型|
|`primary_title`| `string`| **必填**. 影片英文名称|
|`original_title`| `string`| 影片源名称|
|`is_adult`| `bool`| **必填**. 成人分类标识 |
|`start_year`| `int`| 影片/电视剧开始年份|
|`end_year`| `int`| 电视剧结束年份|
|`runtime_minutes`| `int`| 影片持续时间|
|`genres`| `string`| 影片分类（最多三个）|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。
 
##### Response

```text
Status: 200 OK
```

```json
{
    "server_time": "2018-06-12T20:58:21.282769281+08:00",
    "updated_id": "5037605"
}
```

#### Delete a title

```api
DELETE /api/v1/titles/:id
```

##### Response

```text
Status: 204 No Content
```
