h1. REST API сервис ФСТР API для взаимодействия приложения с сервером БД ФСТР

*Version:* 1.0.0

----

{toc:printable=true|style=square|minLevel=2|maxLevel=3|type=list|outline=false|include=.*}

h2. Endpoints

    h3. submitData
    {status:colour=Yellow|title=post|subtle=false}
    {code}
    post /submitData
    {code}
    *Summary:* Внести данные
    *Description:* Вносит данные в карточку объекта


    h4. Parameters

        h5. Body Parameter
        ||Name||Description||Required||Default||Pattern||
        |submitData |Inventory item to add |(x) | |  |







    h4. Responses
        *Status Code:* 201
        *Message:*     item created
        {code:title=Response Type}

        {code}
        See [#models]



        {code:title=Response Schema |collapse=true}
{
  "description" : "item created"
}
        {code}
        *Status Code:* 400
        *Message:*     invalid input, object invalid
        {code:title=Response Type}

        {code}
        See [#models]



        {code:title=Response Schema |collapse=true}
{
  "description" : "invalid input, object invalid"
}
        {code}
        *Status Code:* 409
        *Message:*     an existing item already exists
        {code:title=Response Type}

        {code}
        See [#models]



        {code:title=Response Schema |collapse=true}
{
  "description" : "an existing item already exists"
}
        {code}
    ----

h2. Models

        h3. InventoryItem
        ||Field Name||Required||Type||Description||
         |releaseDate | |Date | |
