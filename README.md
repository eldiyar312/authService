# authServiceGolang


**Api routes in https://eldiyar.herokuapp.com/ address**
CORS alll

|API ROUTES             |description                                             |exemple data|
|-----------------------|--------------------------------------------------------|------------|
|/api/token             |generatin access and refresh tokens by GUID(user id)    |<code><pre>{id: "12345fwfsf"}</code></pre>|
|/api/refresh           |generatin new access and refresh tokens by refresh token|<code><pre>{id: '512asdf',idRefreshToken:'1234124',refreshToken:'ZXl=='}</code></pre>|
|/api/delete/refresh    |remove refresh token                                    |<code><pre>{id: '5f12e12',refreshId: '12d'}</code></pre>|
|/api/delete/all/refresh|remove all refresh token for a specific user            |<code><pre>{id: '5fb62112',}</code></pre>|