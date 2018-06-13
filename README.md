# IMDB Backend

数据库课设 —— IMDB 查找系统（后端仓库）

<!-- TOC -->

- [IMDB Backend](#imdb-backend)
    - [API References](#api-references)
        - [Search](#search)
            - [Search titles](#search-titles)
                - [Parameters](#parameters)
                - [Response](#response)
            - [Search names](#search-names)
                - [Parameters](#parameters-1)
                - [Response](#response-1)
        - [Titles](#titles)
            - [List titles](#list-titles)
                - [Parameters](#parameters-2)
                - [Response](#response-2)
            - [Get a title](#get-a-title)
                - [Response](#response-3)
            - [Create a title](#create-a-title)
                - [Parameters](#parameters-3)
                - [Response](#response-4)
            - [Update a title](#update-a-title)
                - [Parameters](#parameters-4)
                - [Response](#response-5)
            - [Delete a title](#delete-a-title)
                - [Response](#response-6)
        - [Names](#names)
            - [List names](#list-names)
                - [Parameters](#parameters-5)
                - [Response](#response-7)
            - [Get a name](#get-a-name)
                - [Response](#response-8)
            - [Create a name](#create-a-name)
                - [Parameters](#parameters-6)
                - [Response](#response-9)
            - [Update a name](#update-a-name)
                - [Parameters](#parameters-7)
                - [Response](#response-10)
            - [Delete a name](#delete-a-name)
                - [Response](#response-11)

<!-- /TOC -->

## API References

API root: http://localhost:2333

### Search

#### Search titles

```api
GET /api/v1/search/titles
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`q` | `string`| 待搜索字段 |
|`page`| `int` | 当前请求的页号 |

> 若缺少 `page` 字段，则默认返回第一页的信息。

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
                "Int64": 69495,
                "Valid": true
            },
            "TConst": {
                "String": "tt0070909",
                "Valid": true
            },
            "TitleType": {
                "String": "movie",
                "Valid": true
            },
            "PrimaryTitle": {
                "String": "Westworld",
                "Valid": true
            },
            "OriginalTitle": {
                "String": "Westworld",
                "Valid": true
            },
            "IsAdult": {
                "Bool": false,
                "Valid": true
            },
            "StartYear": {
                "Int64": 1973,
                "Valid": true
            },
            "EndYear": {
                "Int64": 0,
                "Valid": false
            },
            "RuntimeMinutes": {
                "Int64": 88,
                "Valid": true
            },
            "Genres": {
                "String": "Action,Sci-Fi,Thriller",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T22:14:51Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T22:14:51Z",
                "Valid": true
            }
        },
        ...  // These lines have been omitted
        {
            "Id": {
                "Int64": 4115389,
                "Valid": true
            },
            "TConst": {
                "String": "tt6528044",
                "Valid": true
            },
            "TitleType": {
                "String": "tvEpisode",
                "Valid": true
            },
            "PrimaryTitle": {
                "String": "WestWorld",
                "Valid": true
            },
            "OriginalTitle": {
                "String": "WestWorld",
                "Valid": true
            },
            "IsAdult": {
                "Bool": false,
                "Valid": true
            },
            "StartYear": {
                "Int64": 2016,
                "Valid": true
            },
            "EndYear": {
                "Int64": 0,
                "Valid": false
            },
            "RuntimeMinutes": {
                "Int64": 0,
                "Valid": false
            },
            "Genres": {
                "String": "Comedy",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T22:14:51Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T22:14:51Z",
                "Valid": true
            }
        }
    ],
    "next_page": 2,
    "prev_page": -1,
    "server_time": "2018-06-13T23:34:27.568742891+08:00",
    "start_id": 0,
    "total_page": 3
}
```

#### Search names

```api
GET /api/v1/search/names
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`q` | `string`| 待搜索字段 |
|`page`| `int` | 当前请求的页号 |

> 若缺少 `page` 字段，则默认返回第一页的信息。

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 50,
    "cur_page": 1,
    "data": [
        {
            "Id": {
                "Int64": 600811,
                "Valid": true
            },
            "NConst": {
                "String": "nm0634240",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Christopher Nolan",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1970,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 0,
                "Valid": false
            },
            "PrimaryProfession": {
                "String": "writer,producer,director",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt0482571,tt1375666,tt0816692,tt0154506",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            }
        },
        ...  // These lines have been omitted
        {
            "Id": {
                "Int64": 600795,
                "Valid": true
            },
            "NConst": {
                "String": "nm0634224",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Bill Nolan",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 0,
                "Valid": false
            },
            "DeathYear": {
                "Int64": 0,
                "Valid": false
            },
            "PrimaryProfession": {
                "String": "actor",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt0116322",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            }
        }
    ],
    "next_page": 2,
    "prev_page": -1,
    "server_time": "2018-06-13T23:38:58.420526838+08:00",
    "start_id": 0,
    "total_page": 14
}
```

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

> 若页号为 `-1`，代表该页不可用。

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

### Names

#### List names

```api
GET /api/v1/names
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
    "count": 50,
    "cur_page": 1,
    "data": [
        {
            "Id": {
                "Int64": 1,
                "Valid": true
            },
            "NConst": {
                "String": "nm0000001",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Fred Astaire",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1899,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 1987,
                "Valid": true
            },
            "PrimaryProfession": {
                "String": "soundtrack,actor,miscellaneous",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt0050419,tt0072308,tt0043044,tt0045537",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            }
        },
        ...  // Some data have been omitted
        {
            "Id": {
                "Int64": 50,
                "Valid": true
            },
            "NConst": {
                "String": "nm0000050",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Groucho Marx",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1890,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 1977,
                "Valid": true
            },
            "PrimaryProfession": {
                "String": "soundtrack,actor,writer",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt0026778,tt0023969,tt0022158,tt0032536",
                "Valid": true
            },
            "CreateDate": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            },
            "LastUpdated": {
                "Time": "2018-06-12T23:01:04Z",
                "Valid": true
            }
        }
    ],
    "next_page": 2,
    "prev_page": -1,
    "server_time": "2018-06-12T23:48:53.499463315+08:00",
    "start_id": 0,
    "total_page": 172809
}
```

> 若页号为 `-1`，代表该页不可用。

#### Get a name

```api
GET /api/v1/name/:id
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
        "NConst": {
            "String": "nm0000001",
            "Valid": true
        },
        "PrimaryName": {
            "String": "Fred Astaire",
            "Valid": true
        },
        "BirthYear": {
            "Int64": 1899,
            "Valid": true
        },
        "DeathYear": {
            "Int64": 1987,
            "Valid": true
        },
        "PrimaryProfession": {
            "String": "soundtrack,actor,miscellaneous",
            "Valid": true
        },
        "KnownForTitles": {
            "String": "tt0050419,tt0072308,tt0043044,tt0045537",
            "Valid": true
        },
        "CreateDate": {
            "Time": "2018-06-12T23:01:04Z",
            "Valid": true
        },
        "LastUpdated": {
            "Time": "2018-06-12T23:01:04Z",
            "Valid": true
        }
    },
    "server_time": "2018-06-12T23:50:59.574007901+08:00"
}
```

#### Create a name

```api
POST /api/v1/names
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`nconst`| `string`| **必填**. 人物唯一标识号|
|`primary_name`| `string`| **必填**. 姓名|
|`birth_year`| `int`| 出生年|
|`death_year`| `int`| 去世年|
|`primary_profession`| `string`| 人物主要领域/技能|
|`known_for_titles`| `string`| 知名作品（按影片唯一标识号填写）|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。
 
##### Response

```text
Status: 200 OK
```

```json
{
    "insert_id": 8650621,
    "server_time": "2018-06-12T23:57:15.220742086+08:00"
}
```

#### Update a name

```api
PUT /api/v1/names/:id
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`nconst`| `string`| **必填**. 人物唯一标识号|
|`primary_name`| `string`| **必填**. 姓名|
|`birth_year`| `int`| 出生年|
|`death_year`| `int`| 去世年|
|`primary_profession`| `string`| 人物主要领域/技能|
|`known_for_titles`| `string`| 知名作品（按影片唯一标识号填写）|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。
 
##### Response

```text
Status: 200 OK
```

```json
{
    "server_time": "2018-06-12T23:57:53.108846482+08:00",
    "updated_id": "8650621"
}
```

#### Delete a name

```api
DELETE /api/v1/names/:id
```

##### Response

```text
Status: 204 No Content
```
