# Task Tracker 

My First solution for the  [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).


## How to run

Clone the repository, and use following command

```bash
git clone https://github.com/itocode21/cli-todo.git
cd main
```




### Run the following command to build:

```bash
go build main.go
```
-----------------------------

## Run the following command to  run the project:
```bash
# To add a task
./main add "u text here"

# To update aa task
./main update n "update u task here" #|n --> task id

# To delete a task
./main delete n #|n --> task id


# To set any status| default status "Pending"
./main done n #set status "done" for task| n --> task id
./main pending n #set status "pending" for task | n --> task id
./main cancel n #set status "status" for task | n --> task id


# To list task
./main list #list all task
./main list-done #list all task with status "done"
./main list-pending #list all task with status "pending"
./main list-cancel #list all task with status "cancel"

```
