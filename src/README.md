# Frontend routes / templates

| URL | User | Template |
| --- | --- | --- |
| `/tig` | `student` | `my_assignments` | DONE
| `/task/{task_id}` | `student` | `one_task` | DONE
| `/tig/{tig_id}` | `student`, `helper`, `admin` | `one_assignment` | DONE
| `/tig` | `helper`, `admin` | `all_assignments` | DONE
| `/tig/modify/{tig_id}` | `helper`, `admin` | `modify_assignment` | PROGRESSING
| `/tig/new` | `helper`, `admin` | `new_assignment` | 
| `/task/{task_id}` | `student`, `helper`, `admin` | `one_task` | DONE
| `/task` | `helper`, `admin` | `all_tasks` | DONE
| `/task/new` | `helper`, `admin` | `new_task` | DONE
| `/task/modify/{task_id}` | `helper`, `admin` | `modify_task` | DONE
| `/login` | `any` | `login` | DONE