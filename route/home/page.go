package home

import (
	"fmt"
	"net/http"
)

const AddForm = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>test</title><style>input{width:100%}.d_none{display:none}.d_visible{display:block}</style></head><body><h1>Данные с сервера в dev tools console</h1><div class="create"><p>user_id: <input id="id" value="5fb6d31a2fbd626a956190fe"></p><button id="generate_tokens">generate tokens by id</button></div><br><br><div class="refresh"><p>id: <input id="id" value="5fb6d31a2fbd626a956190fe"></p><p>idTokens: <input id="idTokens" value="5fb8169937e1c5ff37be3d81"></p><p>accessToken: <input id="accessToken" value="eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU5MDA1MjcsInNvbWV0aGluZyI6IjVmYjZkMzFhMmZiZDYyNmE5NTYxOTBmZSJ9.-p-vQ7wfqE8sbG38zHMk7OaDFMqGTLG7_hAjq7HXtgFMgy24EVnQXs8WOnqQeuYEgVygR4XgnNRmA8Zun6s6gA"></p><p>refreshToken: <input id="refreshToken" value="ZXlKaGJHY2lPaUpJVXpVeE1pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmhkWFJvYjNKcGVtVmtJanAwY25WbExDSmxlSEFpT2pFMk1EVTVNREExTWpjc0luTnZiV1YwYUdsdVp5STZJalZtWWpaa016RmhNbVppWkRZeU5tRTVOVFl4T1RCbVpTSjkuLXAtdlE3d2ZxRThzYkczOHpITWs3T2FERk1xR1RMRzdfaEFqcTdIWHRnRk1neTI0RVZuUVhzOFdPbnFRZXVZRWdWeWdSNFhnbk5SbUE4WnVuNnM2Z0E="></p><button id="refresh">Refreshing</button></div><br><br><div class="delete"><p>id: <input id="id" value="5fb6d31a2fbd626a956190fe"></p><p>tokenId: <input id="tokenId" value="5fb816e937e1c5ff37be3d82"></p><button id="delte_token">Delete token</button></div><br><br><div class="delete_all"><p>id: <input id="id" value="5fb6d31a2fbd626a956190fe"></p><button id="delte_all_tokens">Delete Refresh Tokens</button></div><script defer>const uri="http://eldiyar.herokuapp.com/",generate_tokens=document.getElementById("generate_tokens"),refresh=document.getElementById("refresh"),delete_token=document.getElementById("delte_token"),delte_all_tokens=document.getElementById("delte_all_tokens");generate_tokens.addEventListener("click",create_tokens),refresh.addEventListener("click",refreshing),delete_token.addEventListener("click",deleteToken),delte_all_tokens.addEventListener("click",deleteAllTokens);async function create_tokens(){let a=document.querySelector(".create #id").value;const b={method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:a})};deleteButtons();const c=await fetch("http://eldiyar.herokuapp.com/api/token",b),d=await c.json();deleteButtons(),console.log(d)}async function refreshing(){let a=document.querySelector(".refresh #id").value,b=document.querySelector(".refresh #idTokens").value,c=document.querySelector(".refresh #accessToken").value,d=document.querySelector(".refresh #refreshToken").value;const e={method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:a,idTokens:b,accessToken:c,refreshToken:d})};deleteButtons();const f=await fetch("http://eldiyar.herokuapp.com/api/refresh",e),g=await f.json();deleteButtons(),console.log(g)}async function deleteToken(){let a=document.querySelector(".delete #id").value,b=document.querySelector(".delete #tokenId").value;const c={method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:a,tokenId:b})};deleteButtons();const d=await fetch("http://eldiyar.herokuapp.com/api/delete/refresh",c),e=await d.json();deleteButtons(),console.log(e)}async function deleteAllTokens(){let a=document.querySelector(".delete_all #id").value;const b={method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({id:a})};deleteButtons();const c=await fetch("http://eldiyar.herokuapp.com/api/delete/all/refresh",b),d=await c.json();deleteButtons(),console.log(d)}function deleteButtons(){const a=document.querySelectorAll("button");for(let b=0;b<a.length;b++)a[b].classList.toggle("d_none")}</script></body></html>`

func RespondDoc (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, AddForm)
}