# All routes and their descriptions

- **POST** > _create lesson_ = '***.../lessons/create***'
  - JSON-Body:
```
{
"lesson_type": "lecture",
"name": "TimAnderson",
"description": "Desc of checking"
}
```

- **GET** > _get lesson by **name**_ = '***.../lessons/id/{id}***'
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
  
- **PUT** > update lesson = `***/lessons/update/{id}***`
  + JSON-Body:
```
{
"lesson_type": "practice",
//"name": "UpdName",
"description": "Upd desc for debug"
}
```

- **POST** > attach file to the lesson = `***/lessons/upload/{id_lesson}/{filename}***`
  + Output:
  ```
  File uploaded successfully
  ```
  
- **POST** > tasks = `***/lessons/send/{id}***`
  + Output:
  ```
  <status:ok>
  ```