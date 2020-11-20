# authServiceGolang

**Api routes in https://eldiyar.herokuapp.com/ address**
CORS alll

<pre>
<p>/api/token</p>
<p>generatin access and refresh tokens by GUID(user id)</p>
<code>{id: "12345fwfsf"}</code>
</pre>

</br>
<hr>

<pre>
<p>/api/refresh</p>
<p>generatin new access and refresh tokens by refresh token</p>
<code>
{
    id: '512asdf',
    idRefreshToken:'1234124',
    refreshToken:'ZXl=='
}
</code>
</pre>

</br>
<hr>

<pre>
<p>/api/delete/refresh</p>
<p>remove refresh token</p>
<code>
{
    id: '5f12e12',
    refreshId: '12d'
}
</code>
</pre>

</br>
<hr>

<pre>
<p>/api/delete/all/refresh</p>
<p>remove all refresh token for a specific user</p>
<code>
{
    id: '5fb62112'
}
</code>
</pre>