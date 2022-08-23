# golang-base-code

How to run it ?

just type "air" to run it

Current Features : 

1. JWT validation
2. Filter Query Parser

How to use FQP ?

FQP needs 2 params to use it, those are filter-query and filter-argument. filter-query for raw query and filter-argument for argument value of raw query.

supported value types :
1. integer

   example : @users.id[int]=1
   
2. string

   example : @users.name[string]=abdul rofiq
   
3. date

   example : @users.created_at[date]=2022-08-23
   
4. array

   array value type only support for integer, string, date and boolean.
   example : @users.ids[array]=[int](1,10,11,100)
   
5. boolean

   example : @users.is_active[boolean]=true
   
example :
{{HOST}}?filter-query=users.id = @users.id AND ( users.name != @users.name OR users.email != @users.email )&filter-argument=@users.name[string]=abdul rofiq&&@users.email[string]=abdulrofiq@mail.com&&@users.id[int]=1
