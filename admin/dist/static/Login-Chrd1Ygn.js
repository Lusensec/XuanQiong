import{d as I,u as x,o as b,a as S,r as m,c as y,f as r,w as u,g as t,ab as g,h as V,b as E,Z as T,ac as z,ad as B,e as C,n as f,t as v,v as l,ae as L,x as N,af as D,y as M,_ as K}from"./index-TQ5iY1BV.js";import{a as U}from"./el-message-ZkJEpQUo.js";const j={class:"login-container"},A={class:"login-option"},P=I({__name:"Login",setup(R){const{t:o,locale:Z}=x();function _(){let e=sessionStorage.getItem("token");if(e){try{const a=g(e);if(a.role!=1){sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar");return}if(Math.floor(Date.now()/1e3)>a.exp){sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar");return}}catch{sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar"),location.reload();return}d.push("/")}}b(_);const d=S(),n=m(""),i=m("");async function p(){if(n.value==""||i.value==""){l.error("账号或密码不能为空");return}try{const e=await U.post("/api/v1/login",{username:n.value,password:i.value});if(console.log(e.data),e.data.token){if(g(e.data.token).role!=1){l.error(o("app.webui.nopermation"));return}L({title:o("app.webui.loginsucc"),message:e.data.username+", "+o("app.webui.welcome"),type:"success"}),await new Promise(s=>setTimeout(s,1e3)),sessionStorage.setItem("token",e.data.token),sessionStorage.setItem("username",e.data.username),e.data.avatar==""?sessionStorage.setItem("avatar","/avatar.svg"):sessionStorage.setItem("avatar","/download/file?id="+e.data.avatar),d.push("/")}else l.error(e.data.msg)}catch(e){console.error(e),l.error("登录请求失败")}}return(e,a)=>{const s=N,w=D,h=M,k=V;return E(),y("div",j,[r(k,{class:"login-card",header:t(o)("app.webui.login"),shadow:"always"},{default:u(()=>[r(s,{modelValue:n.value,"onUpdate:modelValue":a[0]||(a[0]=c=>n.value=c),maxlength:"50",placeholder:t(o)("app.webui.username"),size:"large","prefix-icon":t(T),clearable:""},null,8,["modelValue","placeholder","prefix-icon"]),r(s,{modelValue:i.value,"onUpdate:modelValue":a[1]||(a[1]=c=>i.value=c),maxlength:"50",type:"password","show-password":"",placeholder:t(o)("app.webui.password"),size:"large","prefix-icon":t(z),clearable:"",onKeyup:B(p,["enter","native"])},null,8,["modelValue","placeholder","prefix-icon"]),C("div",A,[r(w,{href:"#/register",style:{"margin-right":"10px"}},{default:u(()=>[f(v(t(o)("app.webui.forgot")),1)]),_:1})]),r(h,{type:"success",size:"large",style:{width:"60%",margin:"2% 20%","margin-bottom":"20px","font-weight":"bold","font-size":"16px"},onClick:p,"auto-insert-space":""},{default:u(()=>[f(v(t(o)("app.webui.login")),1)]),_:1})]),_:1},8,["header"])])}}}),G=K(P,[["__scopeId","data-v-e469acc7"]]);export{G as default};
