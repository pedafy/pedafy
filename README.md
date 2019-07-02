# Pedafy

[![travis](https://travis-ci.com/pedafy/pedafy.svg?branch=master)](https://travis-ci.com/pedafy/pedafy)
[![codefactor](https://www.codefactor.io/repository/github/pedafy/pedafy/badge?style=flat-square)](https://www.codefactor.io/repository/github/pedafy/pedafy)

Pedafy is a web application allowing Epitech staff to assign "TIG" to students.

### Version 1
- **Release date:** 01/07/2019
- **Contains**:
  - - [x] OAuth 2 Signin
  - - [ ] Main dashboard (display all the available tabs)
  - - [ ] User rights management (system of groups)
  - - [ ] A "non tigeable" list of user
  - - [ ] Admin interface
    - - [ ] Set user group
    - - [ ] Create/Modify/Delete TIG status
    - - [ ] Ban a user
    - - [ ] Manage the "non-tigeable" users
  - - [x] "My Assignement" tab
    - - [x] View of an assignement
    - - [x] View all our current TIG
    - - [x] View all our old TIG (history)
    - - [ ] Comment the current TIG
  - - [x] "Assign" tab
    - - [x] View of an assignement
    - - [x] Create a new TIG (assign it to someone)
    - - [ ] Comment a TIG
    - - [x] We can remove a TIG
    - - [x] We can see the history of all TIG
    - - [ ] We can seach TIG by status or by assigned's email
  - - [x] "ToDo" tab
    - - [x] View a ToDo
    - - [x] View all the current ToDos
    - - [x] View all the old ToDos
    - - [x] Create an assignement from a ToDo
    - - [x] We can add a new ToDo
    - - [x] We can modify or delete a ToDo
    - - [x] We can see if a ToDo is assigned
    - - [ ] We can comment a ToDo

### Deploy / Dev

To deploy you will need to contact a developer in order to be added to the Google App Engine team.

First, you will need to download Golang version 1.12 (minimum), and the Google App Engine SDK (you can find it on google app engine website). You will also need to the download the cloud SQL proxy that you can find on the same website.

Once it is ready, clone the runner repository and copy all the bin in your path. Then clone all the Go code source in your GOPATH.

You can use `pedafy_start` to run the server locally but beforehand run `pedafy_sql_proxy_start`.

To deploy, push on master and it will automatically push the new version on google app engine thanks to travis.