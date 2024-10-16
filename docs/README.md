# All routes and their descriptions
## Courses
- **POST** > _create courses_ = '***/courses/create***'
  - JSON-Body:
```
{
  "name": "ADMINISTRATION",
  "desc": "Its ADMINISTRATION course"
}
```
  - Output:
```
{
  "id": 3
}
```

- **GET** > _get course by id_ = '***/courses/id/3***'
  - Output:
```
{
  "course_id": 3,
  "name": "Check",
  "desc": "Its check course",
  "created_at": "2024-10-16T12:07:54.902191Z",
  "updated_at": "2024-10-16T12:07:54.902191Z",
  "owner_id": 1
}
```

- **PUT** > _update course by id_ = '***/courses/update/3***'
  - JSON-Body:
```
{
  "name": "UPD check",
  "desc": "Its UPD check course"
}
```
  - Output:
```
{status: "ok"}
```

- **GET** > _get all courses_ = '***/courses/all***'
- Output:
```
{
    "courses": [
        {
            "course_id": 1,
            "name": "IT",
            "desc": "Its IT course",
            "created_at": "2024-10-16T08:26:15.965443Z",
            "updated_at": "2024-10-16T08:26:15.965443Z",
            "owner_id": 1
        },
        {
            "course_id": 2,
            "name": "ADMINISTRATION",
            "desc": "Its ADMINISTRATION course",
            "created_at": "2024-10-16T08:29:25.181718Z",
            "updated_at": "2024-10-16T08:29:25.181718Z",
            "owner_id": 1
        },
        {
            "course_id": 3,
            "name": "UPD check",
            "desc": "Its UPD check course",
            "created_at": "2024-10-16T12:07:54.902191Z",
            "updated_at": "2024-10-16T12:13:04.017882Z",
            "owner_id": 1
        }
    ]
}
```

## Lessons
- **POST** > _create lesson_ = '***/courses/2/lessons/create***'
  - JSON-Body:
```
{
  "lesson_type": "lecture", //practice
  "name": "checkdouble",
  "description": "Desc of checking"
}
```

- **GET** > _get lesson by **id**_ = '***/courses/2/lessons/id/3***'
  - Output:
```
{
  "lesson_id": 3,
  "lesson_type": "practice",
  "name": "UpdName",
  "description": "Upd desc for debug",
  "lesson_file": "check.txt",
  "file_content": "TWFtYSBzaXRhIQ==",
  "lesson_status": "send"
}
```
  
- **GET** > _get lesson by **name**_ = '***/courses/2/lessons/name/check***'
  - Output:
```
{
  "lesson_id": 2,
  "lesson_type": "practice",
  "name": "TimAnderson",
  "description": "Upd desc for debug",
  "lesson_file": "triple.txt",
  "file_content": "dGltQW5kZXJzb243IQ==",
  "lesson_status": "send"
}
```

- **GET** > _get all lessons_ = '***/courses/2/lessons/all***'
- Output:
```
[
    {
        "lesson_id": 3,
        "lesson_type": "practice",
        "name": "UpdName",
        "description": "Upd desc for debug",
        "lesson_file": null,
        "file_content": null,
        "lesson_status": "send"
    },
    {
        "lesson_id": 4,
        "lesson_type": "lecture",
        "name": "check",
        "description": "Desc of checking",
        "lesson_file": null,
        "file_content": null,
        "lesson_status": "Not start"
    },
    {
        "lesson_id": 5,
        "lesson_type": "lecture",
        "name": "checkdouble",
        "description": "Desc of checking",
        "lesson_file": null,
        "file_content": null,
        "lesson_status": "Not start"
    }
]
```
  
- **PUT** > update lesson = '***/courses/{course_id}/lessons/update/{id}***'
  + JSON-Body:
```
{
  "lesson_type": "practice",
  "name": "UpdName",
  "description": "Upd desc for debug"
}
```

- **POST** > attach file to the lesson = '***/courses/2/lessons/upload/{id_lesson}/{filename}***' 
  + Output:
  ```
  File uploaded successfully
  ```
  
- **POST** > tasks = '***/courses/{course_id}/lessons/send/{id}***'
  + Output:
  ```
  <status:ok>
  ```