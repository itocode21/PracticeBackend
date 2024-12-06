# Task Tracker (REFACTORING NOW!)

My First solution for the  [task-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).


## How to run

Clone the repository, and use following command

```bash
git clone https://github.com/itocode21/ExpenseTracker.git
cd main
```




### Run the following command to build:

```bash
go build main.go
```
-----------------------------

## Run the following command to  run the project:
```bash
# To add a expense
./main add "Description" 10  #10 --> amount 

# To delete expense
./main delete n #|n --> expense id

# To list summaty expense
./main summary #|list summary all expense

# To list
./main list #|list all expense

# To list expense by month
./main month n #|n --> month 1-12

# export json file with data
./main export #|u get link http://localhost:8080/download ctrl+rmb for dowloand. ctrl+c for exit


```