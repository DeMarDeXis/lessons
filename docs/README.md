# All routes and their descriptions

- **POST** > sign_up = '***/auth/sign-up***'
    - JSON-Body:
  ```
  {
  "name": "Tim",
  "username": "TimAnderson7",
  "password": "qwerty"
  }
  ```


- **POST** > sign_in = '***/auth/sign-in***'
  - JSON-Body:
  ```
  {
  "username": "TimAnderson7",
  "password": "qwerty"
  }
  ```
  - Output:
  ```
  {
  "token": "<HERE_IS_TOKEN>"
  }
  ```
  
- **POST** > tasks = `***/app/tasks***`
  + Authorization > Bearer Token
  + JSON-Body:
  ```
    {
    "title": "Serbia Ultras",
    "description": "Partizan",
    "doe_date": "2026-01-02T15:04:05Z"
    }
  ```
  + Output:
  ```
    {
    "id": 4
    }
  ```
- **GET** > tasks = `***/app/tasks***`
  + Authorization > Bearer Token
  + Output:
  ```
  <Tareas de personas que estan en la base de datos>
  ```
- **GET** > tasks = `***/app/tasks/id***`
  + Authorization > Bearer Token
  + Output:
  ```
  <Tarea de usario que esta en la base de datos>
  ```
**PUT** > tasks = `***/app/tasks/id***`
  + Authorization > Bearer Token
  + JSON-Body:
  ```
    {
    "title": "Serbia Ultras",
    "description": "Crvena Zvezda",
    "doe_date": "2026-12-22T15:04:05Z"
    }
  ```
  + Output:
  ```
    {
     "status": "ok"
    }