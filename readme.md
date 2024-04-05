# Standard Lib Go Simple Server
For testing and templating future Go server.

**Sources**

https://www.youtube.com/watch?v=H7tbjKFSg58

## Test with
 - curl -X POST localhost:8080/item -d "{\"id\":78, \"name\": \"defName\"}"
 - curl -X POST localhost:8080/v2/item  -H "Authorization: Bearer Y2lhbwo=" -d "{\"id\":78, \"name\": \"defName\"}"
 - curl localhost:8080/item/34
 - curl -H "Authorization: Bearer Y2lhbwo=" -v localhost:8080/item/88
