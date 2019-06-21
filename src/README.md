# Frontend routes / templates

| URL | User | Template |
| --- | --- | --- |
| `/tig` | `student` | `my_assignments` |
| `/task/{task_id}` | `student` | `one_task` |
| `/tig/{tig_id}` | `student`, `helper`, `admin` | `one_assignment` |
| `/tig` | `helper`, `admin` | `all_assignments` |
| `/tig/modify/{tig_id}` | `helper`, `admin` | `modify_assignment` |
| `/tig/new` | `helper`, `admin` | `new_assignment` |
| `/tig/review/{tig_id}` | `admin` | `review_assignment` |
| `/task/{task_id}` | `student`, `helper`, `admin` | `one_task` |
| `/task` | `helper`, `admin` | `all_tasks` |
| `/task/new` | `helper`, `admin` | `new_task` |
| `/task/modify/{task_id}` | `helper`, `admin` | `modify_task` |
| `/login` | `any` | `login` |