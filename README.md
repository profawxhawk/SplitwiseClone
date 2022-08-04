# SplitwiseClone

To run the application, install docker and run:
```
make dev
```
Supports 5 APIs

```
"/api/users/add", POST, input: name ,adds a new user and returns userID
"/api/groups/add", POST, input: name ,adds a new group and returns groupID
"/api/users/addToGroup", POST, input: { userId, groupId } , adds a user to a group
"/api/groups/addExpenseToGroup, POST, input : {
   "userId":1,
   "groupId":1,
   "expenseName":"out",
   "totalAmount":500,
   "expenseType":0, // enum 0 (equal split) or 1 (exact split)
   "userAmountMap":{
      "2":200,
      "3":300
   }
} , creates an expense record
"/api/users/getAllExpenses", GET, input: userId , gets all transactions for a user
```

HTTP server is exposed on port 8080
