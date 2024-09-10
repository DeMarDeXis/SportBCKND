# All routes and their descriptions

- **POST** > sign_up = '***/sign-up***'
  - JSON-Body:
    ```
    {
    "name": "Tim",
    "username": "TimAnderson7"
    }
    ```


- **POST** > sign_in = '***/sign-in***'
  - JSON-Body:
  ```
    {
    "username": "TimAnderson7"
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
    "title": "UK Drill",
    "description": "I'll go to the gym and do some exercises",
    "doe_date": {
        "year": 2025,
        "month": 4,
        "day": 17
    }
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
