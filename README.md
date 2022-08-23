# golang-base-code

current features
1. JWT validation
2. Filter Query Parser

How to use FQP ?

FQP needs 2 params to use it, those are filter-query and filter-argument. filter-query for query raw and filter-argument for argument value of raw query.
e.g :
{{HOST}}?filter-query=users.id = @users.id AND ( users.name != @users.name OR users.email != @users.email )&filter-argument=@users.name[string]=abdul rofiq&&@users.email[string]=abdulrofiq@mail.com&&@users.id[int]=1
