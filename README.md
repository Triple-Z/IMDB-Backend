# IMDB Backend

数据库课设 —— IMDB 查找系统（后端仓库）

<!-- TOC -->

- [IMDB Backend](#imdb-backend)
    - [API References](#api-references)
        - [Search](#search)
            - [Fuzzy search titles](#fuzzy-search-titles)
                - [Parameters](#parameters)
                - [Response](#response)
            - [Fuzzy search names](#fuzzy-search-names)
                - [Parameters](#parameters-1)
                - [Response](#response-1)
            - [Advanced search titles](#advanced-search-titles)
                - [Parameters](#parameters-2)
                - [Response](#response-2)
            - [Advanced search names](#advanced-search-names)
                - [Parameters](#parameters-3)
                - [Response](#response-3)
        - [Titles](#titles)
            - [List titles](#list-titles)
                - [Parameters](#parameters-4)
                - [Response](#response-4)
            - [Get a title](#get-a-title)
                - [Response](#response-5)
            - [Create a title](#create-a-title)
                - [Parameters](#parameters-5)
                - [Response](#response-6)
            - [Update a title](#update-a-title)
                - [Parameters](#parameters-6)
                - [Response](#response-7)
            - [Delete a title](#delete-a-title)
                - [Response](#response-8)
            - [Get a title details](#get-a-title-details)
                - [Response](#response-9)
        - [Names](#names)
            - [List names](#list-names)
                - [Parameters](#parameters-7)
                - [Response](#response-10)
            - [Get a name](#get-a-name)
                - [Response](#response-11)
            - [Create a name](#create-a-name)
                - [Parameters](#parameters-8)
                - [Response](#response-12)
            - [Update a name](#update-a-name)
                - [Parameters](#parameters-9)
                - [Response](#response-13)
            - [Delete a name](#delete-a-name)
                - [Response](#response-14)
        - [Principals](#principals)
            - [Get a principal](#get-a-principal)
                - [Response](#response-15)
            - [Create a principal](#create-a-principal)
                - [Parameters](#parameters-10)
                - [Response](#response-16)
            - [Update a principal](#update-a-principal)
                - [Parameters](#parameters-11)
                - [Response](#response-17)
            - [Delete a principal](#delete-a-principal)
                - [Response](#response-18)

<!-- /TOC -->

## API References

API root: http://localhost:2333

### Search

#### Fuzzy search titles

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

#### Fuzzy search names

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

#### Advanced search titles

```api
GET /api/v1/advanced_search/titles
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`q` | `string`| 待搜索字段（支持布尔表达式） |
|`tconst` | `string` | 影片唯一标识号 |
|`isAdult` | `string` | 成人分类标识|
|`startYearStart` | `int` | 影片/电视剧开始年份的起始值|
|`startYearEnd` | `int` | 影片/电视剧开始年份的结束值|
|`endYearStart` | `int` | 电视剧结束年份的起始值|
|`endYearEnd` | `int` | 电视剧结束年份的结束值|
|`runtimeMinStart` | `int` | 影片持续时间的起始值|
|`runtimeMinEnd` | `int` | 影片持续时间的结束值|
|`page`| `int` | 当前请求的页号 |

> 以上值都为非必需，若值为空可以不传。

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 1,
    "cur_page": 1,
    "data": [
        {
            "Id": {
                "Int64": 458041,
                "Valid": true
            },
            "TConst": {
                "String": "tt0475784",
                "Valid": true
            },
            "TitleType": {
                "String": "tvSeries",
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
                "Int64": 2016,
                "Valid": true
            },
            "EndYear": {
                "Int64": 0,
                "Valid": false
            },
            "RuntimeMinutes": {
                "Int64": 62,
                "Valid": true
            },
            "Genres": {
                "String": "Drama,Mystery,Sci-Fi",
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
    "next_page": -1,
    "prev_page": -1,
    "server_time": "2018-06-16T15:46:04.38158079+08:00",
    "start_id": 0,
    "total_page": 1
}
```

#### Advanced search names

```api
GET /api/v1/advanced_search/names
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`q` | `string`| 待搜索字段（支持布尔表达式） |
|`tconst` | `string` | 影片唯一标识号 |
|`isAdult` | `string` | 成人分类标识|
|`startYearStart` | `int` | 影片/电视剧开始年份的起始值|
|`startYearEnd` | `int` | 影片/电视剧开始年份的结束值|
|`endYearStart` | `int` | 电视剧结束年份的起始值|
|`endYearEnd` | `int` | 电视剧结束年份的结束值|
|`runtimeMinStart` | `int` | 影片持续时间的起始值|
|`runtimeMinEnd` | `int` | 影片持续时间的结束值|
|`page`| `int` | 当前请求的页号 |

> 以上值都为非必需，若值为空可以不传。

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
                "Int64": 887972,
                "Valid": true
            },
            "NConst": {
                "String": "nm0939697",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Evan Rachel Wood",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1987,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 0,
                "Valid": false
            },
            "PrimaryProfession": {
                "String": "actress,soundtrack,music_department",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt0328538,tt1178663,tt1124035,tt0475784",
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
        ...  // These lines have been omitted.
        {
            "Id": {
                "Int64": 4764388,
                "Valid": true
            },
            "NConst": {
                "String": "nm5279037",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Evan Kopczyk",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1992,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 0,
                "Valid": false
            },
            "PrimaryProfession": {
                "String": "miscellaneous,editorial_department,actor",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt1793223,tt1094666,tt1831806,tt2882174",
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
        {
            "Id": {
                "Int64": 4780054,
                "Valid": true
            },
            "NConst": {
                "String": "nm5295960",
                "Valid": true
            },
            "PrimaryName": {
                "String": "Evan Chester",
                "Valid": true
            },
            "BirthYear": {
                "Int64": 1988,
                "Valid": true
            },
            "DeathYear": {
                "Int64": 0,
                "Valid": false
            },
            "PrimaryProfession": {
                "String": "miscellaneous,actor",
                "Valid": true
            },
            "KnownForTitles": {
                "String": "tt5867314,tt2908446,tt2975590,tt2109248",
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
    "server_time": "2018-06-16T15:55:51.299654741+08:00",
    "start_id": 0,
    "total_page": 3
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

#### Get a title details

```api
GET /api/v1/titles/:id/details
```

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 10,
    "data": [
        {
            "Id": {
                "Int64": 3295656,
                "Valid": true
            },
            "Ordering": {
                "Int64": 1,
                "Valid": true
            },
            "NameId": {
                "Int64": 1379242,
                "Valid": true
            },
            "PrimaryName": {
                "String": "Evan Rachel Wood",
                "Valid": true
            },
            "Category": {
                "String": "actress",
                "Valid": true
            },
            "Job": {
                "String": "",
                "Valid": false
            },
            "Characters": {
                "String": "Dolores Abernathy",
                "Valid": true
            }
        },
        ...  // These lines have been omitted
        {
            "Id": {
                "Int64": 3295655,
                "Valid": true
            },
            "NameId": {
                "Int64": 438228,
                "Valid": true
            },
            "Ordering": {
                "Int64": 10,
                "Valid": true
            },
            "PrimaryName": {
                "String": "Sidse Babett Knudsen",
                "Valid": true
            },
            "Category": {
                "String": "actress",
                "Valid": true
            },
            "Job": {
                "String": "",
                "Valid": false
            },
            "Characters": {
                "String": "Theresa Cullen",
                "Valid": true
            }
        }
    ],
    "server_time": "2018-06-16T14:58:01.641212628+08:00"
}
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

### Principals

#### Get a principal

```api
GET /api/v1/principals/:id
```

##### Response

```text
Status: 200 OK
```

```json
{
    "data": {
        "Id": {
            "Int64": 3295656,
            "Valid": true
        },
        "TConst": {
            "String": "tt0475784",
            "Valid": true
        },
        "Ordering": {
            "Int64": 1,
            "Valid": true
        },
        "NConst": {
            "String": "nm0939697",
            "Valid": true
        },
        "Category": {
            "String": "actress",
            "Valid": true
        },
        "Job": {
            "String": "",
            "Valid": false
        },
        "Characters": {
            "String": "Dolores Abernathy",
            "Valid": true
        },
        "CreateDate": {
            "Time": "2018-06-14T21:24:15Z",
            "Valid": true
        },
        "LastUpdated": {
            "Time": "2018-06-14T21:24:15Z",
            "Valid": true
        }
    },
    "server_time": "2018-06-16T15:06:06.861335428+08:00"
}
```

#### Create a principal

```api
POST /api/v1/principals
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`tconst`| `string`| **必填**. 影片唯一标识号|
|`ordering`| `int`| **必填**. 人物在影片中的排序|
|`nconst`| `string`| **必填**. 人物唯一标识号|
|`category`| `string`| 人物主要职务|
|`job`| `string`| 人物主要职责|
|`characters`| `string`| 人物扮演的角色名|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。

##### Response

```text
Status: 200 OK
```

```json
{
    "insert_id": 28507726,
    "server_time": "2018-06-16T15:17:36.334594154+08:00"
}
```

#### Update a principal

```api
PUT /api/v1/principals/:id
```

##### Parameters

| Name | Type | Description |
|:----:|:----:| :-----: |
|`tconst`| `string`| **必填**. 影片唯一标识号|
|`ordering`| `int`| **必填**. 人物在影片中的排序|
|`nconst`| `string`| **必填**. 人物唯一标识号|
|`category`| `string`| 人物主要职务|
|`job`| `string`| 人物主要职责|
|`characters`| `string`| 人物扮演的角色名|

> `form-data` 和 `application/json` 都能被接受，且字段名称类型相同。

##### Response

```text
Status: 200 OK
```

```json
{
    "server_time": "2018-06-16T15:20:03.23159969+08:00",
    "updated_id": "28507726"
}
```

#### Delete a principal

```api
DELETE /api/v1/principals/:id
```

##### Response

```text
Status: 204 No Content
```
