(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["a650eadc"],{4917:function(M,t,j){"use strict";var e=j("cb7c"),u=j("9def"),L=j("0390"),i=j("5f1b");j("214f")("match",1,(function(M,t,j,r){return[function(j){var e=M(this),u=void 0==j?void 0:j[t];return void 0!==u?u.call(j,e):new RegExp(j)[t](String(e))},function(M){var t=r(j,M,this);if(t.done)return t.value;var N=e(M),n=String(this);if(!N.global)return i(N,n);var a=N.unicode;N.lastIndex=0;var o,c=[],s=0;while(null!==(o=i(N,n))){var y=String(o[0]);c[s]=y,""===y&&(N.lastIndex=L(n,u(N.lastIndex),a)),s++}return 0===s?null:c}]}))},"51e0":function(M,t,j){},9224:function(M){M.exports=JSON.parse('{"name":"traefik-ui","version":"2.0.0","description":"Traefik UI","productName":"Traefik","cordovaId":"us.containo.traefik","private":true,"scripts":{"transfer":"node dev/scripts/transfer.js","lint":"eslint --ext .js,.vue src","test-unit":"mocha-webpack --mode=production \'./src/**/*.spec.js\'","dev":"export APP_ENV=\'development\' && quasar dev","build-quasar":"quasar build","build-staging":"export NODE_ENV=\'production\' && export APP_ENV=\'development\' && npm run build-quasar","build":"export NODE_ENV=\'production\' && export APP_ENV=\'production\' && npm run build-quasar && npm run transfer spa","build:nc":"npm run build"},"dependencies":{"@quasar/extras":"^1.0.0","axios":"^0.19.0","bowser":"^2.5.2","chart.js":"^2.8.0","dot-prop":"^5.2.0","iframe-resizer":"^4.2.11","lodash.isequal":"4.5.0","moment":"^2.24.0","quasar":"^1.4.4","query-string":"^6.13.1","vh-check":"^2.0.5","vue-chartjs":"^3.4.2","vuex-map-fields":"^1.3.4"},"devDependencies":{"@quasar/app":"^1.2.4","@vue/eslint-config-standard":"^4.0.0","@vue/test-utils":"^1.0.0-beta.29","babel-eslint":"^10.0.1","chai":"4.2.0","eslint":"^5.10.0","eslint-loader":"^2.1.1","eslint-plugin-prettier":"3.1.1","eslint-plugin-mocha":"6.2.1","eslint-plugin-vue":"^5.0.0","mocha":"^6.2.2","mocha-webpack":"^2.0.0-beta.0","node-sass":"^4.12.0","prettier":"1.19.1","sass-loader":"^7.1.0"},"engines":{"node":">= 8.9.0","npm":">= 5.6.0","yarn":">= 1.6.0"},"browserslist":["last 1 version, not dead, ie >= 11"]}')},"9b19":function(M,t){M.exports="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI4NCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDg0IDI0Ij4KICAgIDxnIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0ibm9uemVybyI+CiAgICAgICAgPHBhdGggZmlsbD0iI0ZGRiIgZD0iTTcuMTQyIDI0Yy0xLjI0MSAwLTIuMjMtLjQwNi0yLjk2My0xLjIxNy0uNzM1LS44MjItMS4xMDItMi4wODItMS4xMDItMy43ODF2LTkuNDdIMS4xMDFjLS4zNTYgMC0uNjMxLS4wODgtLjgyNi0uMjYzQy4wOTIgOS4wOTMgMCA4Ljg1OCAwIDguNTYyYzAtLjI4NS4wOTItLjUxNi4yNzUtLjY5MS4xODQtLjE4Ni40NTQtLjI4LjgxLS4yOGgxLjk5MmwuMzg5LTQuNjY5Yy4wMjItLjM1LjE3My0uNjUyLjQ1My0uOTA0LjI5Mi0uMjUyLjYxLS4zNzguOTU2LS4zNzguMzI0IDAgLjU5NC4xMDQuODEuMzEyLjIxNi4xOTcuMzI0LjQ4OC4zMjQuODcydjQuNzY4aDIuOThjLjM1NiAwIC42Mi4wOTMuNzkzLjI3OS4xODQuMTg2LjI3Ni40MjguMjc2LjcyMyAwIC42MjUtLjM1Ny45MzgtMS4wNy45MzhoLTIuOTh2OC45NmMwIDEuMTMuMTU3IDEuOTQ2LjQ3IDIuNDUuMzE0LjQ5My44MS43NCAxLjQ5Ljc0aC4xNzlsMS4xMDEtLjA1aC4wOTdjLjM1NiAwIC42MzcuMTE2Ljg0Mi4zNDYuMjA1LjIzLjMwOC40OTguMzA4LjgwNSAwIC4yNzQtLjA2LjUtLjE3OC42NzQtLjExOS4xNjUtLjMwMi4yODUtLjU1LjM2MmEzLjMxIDMuMzEgMCAwIDEtLjczLjE0OGMtLjIzNy4wMjItLjU1LjAzMy0uOTM5LjAzM2gtLjk1NnptNy43NzEgMGMtLjQwNSAwLS43NS0uMTItMS4wMzUtLjM1OC0uMjg0LS4yMzgtLjQyNy0uNTc0LS40MjctMS4wMDdWOS4yMjhjMC0uNDM0LjE0OC0uNzY0LjQ0NC0uOTkyYTEuNjIzIDEuNjIzIDAgMCAxIDEuMDUxLS4zNTdjLjQwNSAwIC43NS4xMTkgMS4wMzUuMzU3LjI5NS4yMjguNDQzLjU0Ny40NDMuOTZsLS4wMTYgMS45NWE0LjgxNSA0LjgxNSAwIDAgMSAxLjY0Mi0yLjUyYy44MzItLjY4MiAxLjgyOS0xLjAyNCAyLjk5LTEuMDI0LjM1IDAgLjYyNC4xMy44Mi4zOS4yMDkuMjYuMzEzLjU2NC4zMTMuOTEgMCAuMzI2LS4wOTMuNjA3LS4yOC44NDYtLjE4Ni4yMjctLjQ1NC4zNC0uODA0LjM0LTEuNDc4IDAtMi42MjguMzgtMy40NSAxLjEzOS0uODEuNzU4LTEuMjE1IDEuNzctMS4yMTUgMy4wMzh2OC4zN2MwIC40MzMtLjE1My43Ny0uNDYgMS4wMDdhMS42MjMgMS42MjMgMCAwIDEtMS4wNS4zNTh6bTQwLjYxLS4zODNhMS41MjIgMS41MjIgMCAwIDEtMS4wMzcuMzgzYy0uNCAwLS43NDUtLjEyOC0xLjAzNy0uMzgzLS4yOC0uMjY3LS40Mi0uNjMzLS40Mi0xLjFWOS41NzhoLTIuMTIzYy0uMzU3IDAtLjYyMS0uMDktLjc5NC0uMjY3YS45OTcuOTk3IDAgMCAxLS4yNi0uN2MwLS4yODguMDg3LS41MjcuMjYtLjcxNS4xNzMtLjE5LjQzNy0uMjg0Ljc5NC0uMjg0aDIuMTIyVjUuNDk2YzAtMS45ODcuNDEtMy4zOTcgMS4yMzItNC4yM0M1NS4wOC40MjIgNTYuMjE1IDAgNTcuNjYyIDBoMS45NzdjLjMxMyAwIC41NTYuMTE3LjcyOS4zNS4xODMuMjIyLjI3NS40ODguMjc1LjggMCAuMzEtLjA5Mi41NzctLjI3NS43OTktLjE3My4yMjItLjQyMS4zMzMtLjc0NS4zMzNoLTEuMzc4Yy0uNDY0IDAtLjgyNi4wNDQtMS4wODUuMTMzLS4yNi4wNzgtLjQ5Mi4yNi0uNjk3LjU1LS4xOTQuMjc3LS4zMy42ODItLjQwNSAxLjIxNS0uMDY1LjUyMi0uMDk3IDEuMjMzLS4wOTcgMi4xMzJ2MS4zaDIuNzM4Yy4zNTYgMCAuNjE2LjA5NC43NzguMjgzLjE3My4xODguMjU5LjQyNy4yNTkuNzE2IDAgLjY0NC0uMzQ2Ljk2Ni0xLjAzNy45NjZoLTIuNzM4djEyLjk0YzAgLjQ2Ny0uMTQ2LjgzMy0uNDM4IDEuMXptOS41NTctMjAuMDRjLS41OCAwLTEuMDQ1LS4xNjEtMS4zOTYtLjQ4NS0uMzUtLjMzNS0uNTI1LS43NjYtLjUyNS0xLjI5NSAwLS41MjkuMTc1LS45Ni41MjUtMS4yOTUuMzYyLS4zMzUuODMyLS41MDIgMS40MTMtLjUwMi41NyAwIDEuMDMuMTY3IDEuMzguNTAyLjM1LjMyNC41MjUuNzU1LjUyNSAxLjI5NSAwIC41MjktLjE3NS45Ni0uNTI2IDEuMjk1LS4zNS4zMjQtLjgxNS40ODYtMS4zOTYuNDg2ek02NS4wNTYgMjRhMS40OSAxLjQ5IDAgMCAxLTEuMDMtLjM4Yy0uMjgyLS4yNjMtLjQyNC0uNjM3LS40MjQtMS4xMjFWOS4wODdjMC0uNDYyLjE0Ny0uODI1LjQ0MS0xLjA4OWExLjU3MyAxLjU3MyAwIDAgMSAxLjA2Mi0uMzk2Yy4zOTIgMCAuNzMuMTMyIDEuMDEzLjM5Ni4yOTQuMjY0LjQ0LjYyNy40NCAxLjA5djEzLjQxYzAgLjQ2My0uMTUyLjgzMS0uNDU3IDEuMTA2YTEuNTEzIDEuNTEzIDAgMCAxLTEuMDQ1LjM5NnptOC40NC0uNDA3YTEuNTMgMS41MyAwIDAgMS0xLjA0OC4zOWMtLjQwMyAwLS43NDctLjEyNC0xLjAzLS4zNzQtLjI4NC0uMjYtLjQyNi0uNjI5LS40MjYtMS4xMDZWMS40NjRjMC0uNDU1LjE0Ny0uODEzLjQ0Mi0xLjA3My4zMDUtLjI2LjY2LS4zOTEgMS4wNjMtLjM5MS4zOTMgMCAuNzMxLjEzIDEuMDE1LjM5LjI5NC4yNi40NDIuNjE5LjQ0MiAxLjA3NFYxNC42Nmw3LjQyOC02Ljc2OGMuMjQtLjIwNy40OTYtLjMxLjc3LS4zMS4zNDggMCAuNjU0LjEzNi45MTUuNDA3LjI3My4yNzEuNDEuNTc1LjQxLjkxMSAwIC4zMDQtLjEzMi41OC0uMzkzLjgzbC01Ljk0IDUuMjg4IDYuNDk2IDYuODAyYy4yNC4yNS4zNi41MzEuMzYuODQ2IDAgLjM0Ny0uMTQyLjY1Ni0uNDI1LjkyNy0uMjczLjI3MS0uNTkuNDA3LS45NS40MDctLjM4MSAwLS43MDgtLjE1Mi0uOTgxLS40NTZsLTcuNjktOC4xMzV2Ny4wOTRjMCAuNDY2LS4xNTMuODMtLjQ1OSAxLjA5eiIvPgogICAgICAgIDxwYXRoIGZpbGw9IiM0Qjc2REQiIGQ9Ik0yOC45MyAyMy45OTNjLTEuOTQyLS4xNjMtMy40MjctLjczNC00LjUwNi0xLjczMi0xLjEtMS4wMTYtMS41Ny0yLjM0LTEuMzg0LTMuODg4LjIxNC0xLjc4MSAxLjA4My0zLjAzOSAyLjYwOC0zLjc3MyAxLjE3NS0uNTY1IDIuMzQ4LS43OCA0LjI1Ny0uNzgxIDEuMzU3IDAgMi40OTMuMTA4IDQuMjI3LjQwMS40MjkuMDczLjguMTMyLjgyNS4xMzIuMDM2IDAgLjA2NS0uMTMzLjE0NC0uNjY0LjE0Mi0uOTYuMTY3LTEuMjkyLjEyOS0xLjY3Ni0uMDQ2LS40NDgtLjEwNC0uNjYyLS4yNzktMS4wMDktLjQ1LS44OTUtMS41MTctMS40NDMtMy4yMzItMS42NTgtLjQ5LS4wNjItMS42MjItLjA3LTIuMTE0LS4wMTYtLjgyMy4wOS0xLjI4Mi4yMDYtMi4zMDYuNTgtLjYuMjE4LS42MDIuMjE4LS44NDEuMjA1LS40MTEtLjAyNC0uNjkyLS4yMDUtLjg3Mi0uNTYzLS4wNzQtLjE0Ny0uMDg3LS4yMDgtLjA4Ny0uNDA3IDAtLjI5OC4wOC0uNDU2LjM1OC0uNzA5LjYyMy0uNTY3IDEuNjcxLTEuMDUgMi43My0xLjI1OC42NzctLjEzMiAxLjAxLS4xNiAxLjkwMy0uMTU5LjkyMy4wMDIgMS41MTQuMDUyIDIuMzM1LjIgMS4zNjUuMjQ4IDIuNDAyLjY4MyAzLjE4MyAxLjMzOC4yMzIuMTk0LjU4Mi41NzcuNzMyLjhhLjQzMy40MzMgMCAwIDAgLjEwNS4xMjRjLjAxMiAwIC4wODctLjA4Mi4xNjgtLjE4MkMzOC4xOTcgNy44MyA0MC4xOSA2Ljk5OCA0Mi41MjYgN2MxLjY1MiAwIDMuMTg3LjQwOCA0LjMzMiAxLjE0OC45MDguNTg3IDEuNjQ1IDEuNDkgMS45NDMgMi4zODIuMTcxLjUxMi4xOTcuNzAxLjE5OSAxLjQ0NCAwIC40NjMtLjAxMi43NzUtLjAzOS45Mi0uMjY4IDEuNDYtLjkwNiAyLjQ3My0yLjAxMyAzLjE5My0uODkxLjU4LTIuMDEuOTItMy40NzQgMS4wNTMtLjczNi4wNjgtMi4wODQuMDYxLTIuOTU3LS4wMTQtLjc0LS4wNjMtMS42NS0uMTc1LTIuMDM4LS4yNWEzNC4xMiAzNC4xMiAwIDAgMC0xLjQ0Ni0uMjM2IDIyLjE1IDIyLjE1IDAgMCAwLS4xMTguNzJjLS4wOTcuNjI1LS4xMS43ODItLjExMyAxLjI4MS0uMDAyLjUwOC4wMDYuNTk3LjA2OS44MzIuMjAyLjc1NC42NDggMS4yNTggMS40NTggMS42NDguNzk3LjM4MyAxLjYzMS41NSAyLjg5My41NzUuNTU5LjAxMS44NDUuMDAzIDEuMTk0LS4wMzMuODQ3LS4wODcgMS4yMzMtLjE4MyAyLjM1OC0uNTg2LjI4LS4xLjU2Ny0uMTkxLjYzNi0uMjAzLjE0OS0uMDI0LjQyNi4wMjYuNTg3LjEwNi4yOTIuMTQ0LjUzNy41MjguNTM1Ljg0LS4wMDEuMzYyLS4xMjcuNTgyLS40OTcuODY4LS44ODMuNjg0LTIuMDEgMS4xMDQtMy4zNTcgMS4yNS0uNDguMDUtMS41NzYuMDU5LTIuMTA2LjAxNC0xLjQ5LS4xMjYtMi41MjYtLjM3NC0zLjQ3Mi0uODMzLS43MzgtLjM1Ny0xLjMwMS0uOC0xLjcyMi0xLjM1NmEyNC42OSAyNC42OSAwIDAgMC0uMTk1LS4yNTUgMS4zMDEgMS4zMDEgMCAwIDAtLjE1My4xNzUgNi4xMDQgNi4xMDQgMCAwIDEtMS40MjMgMS4yNTZjLS43NjMuNDgyLTEuODU5Ljg1LTIuOTY3Ljk5Ny0uMzMxLjA0NC0xLjQyOC4wOC0xLjcxLjA1N3ptMS45MDQtMi4wOGMuOTA2LS4xOTUgMS43MTItLjYzNSAyLjI5OC0xLjI1My41MzQtLjU2My45NzItMS40MDggMS4yMTMtMi4zMzguMTM2LS41MjUuMzYyLTEuOTg1LjMxMi0yLjAxNS0uMDItLjAxMi0uMTQtLjAzLS4yNjYtLjA0MmEzMS4zNjQgMzEuMzY0IDAgMCAxLS43OTktLjA5MmMtMS44Ni0uMjM0LTMuMDAyLS4zMjUtMy42NDItLjI5LTEuMjE3LjA2Ni0xLjcxMy4xNy0yLjQ0My41MS0uMzc0LjE3My0uNjMzLjM2Mi0uODM3LjYxLS4yOTMuMzU4LS40Ny43MTMtLjU5NCAxLjE5MS0uMDkuMzQ4LS4wOTkgMS4yNTktLjAxNSAxLjU1Mi4xNC40OTUuMzA3Ljc3NS42ODEgMS4xNDUuNTEzLjUwOCAxLjE5Ljg0NiAyLjAyOSAxLjAxNS40MjkuMDg3LjQxNi4wODYgMS4xMzQuMDc2LjUyMS0uMDA4LjcxNi0uMDIyLjkzLS4wNjh6bTExLjg1Ny02Ljg0NWMuNzU3LS4wNjUgMS4yMjItLjE4IDEuNzk0LS40NC40NzYtLjIxOC43NTItLjQ0IDEuMDE0LS44MTYuMzg1LS41NTIuNTM1LTEuMDc1LjUzMS0xLjg1Mi0uMDAzLS41NTQtLjA1LS43ODQtLjI0Ny0xLjE4Mi0uNDQtLjg4OC0xLjQ3LTEuNTI0LTIuODIyLTEuNzQxLS4zMzgtLjA1NS0xLjI0My0uMDQ1LTEuNTkuMDE2LS42NjQuMTE5LTEuMjI3LjMzNC0xLjcyMy42Ni0uOTE2LjYwNC0xLjU0OSAxLjUxMy0xLjkgMi43MzQtLjE0LjQ4LS4xNzYuNjYzLS4zIDEuNDlsLS4xMTMuNzYyLjE1Mi4wMTdjLjQ2LjA1IDEuNDk4LjE3MiAyLjEwMS4yNDcuMzguMDQ4Ljg0NS4wOTkgMS4wMzQuMTEzLjU0OS4wNDIgMS41MzEuMDM4IDIuMDY5LS4wMDh6Ii8+CiAgICA8L2c+Cjwvc3ZnPgo="},"9d96":function(M,t,j){},a17c:function(M,t,j){"use strict";var e=j("51e0"),u=j.n(e);u.a},bf59:function(M,t,j){"use strict";var e=j("9d96"),u=j.n(e);u.a},ea2a:function(M,t,j){"use strict";j.r(t);var e=function(){var M=this,t=M.$createElement,j=M._self._c||t;return j("q-layout",{attrs:{view:"hHh lpR fFf"}},[j("nav-bar",[j("router-view",{attrs:{name:"NavBar"}})],1),j("q-page-container",[j("router-view")],1)],1)},u=[],L=function(){var M=this,t=M.$createElement,e=M._self._c||t;return e("q-header",{staticClass:"shadow-1"},[e("section",{staticClass:"app-section bg-primary text-white"},[e("div",{staticClass:"app-section-wrap app-boxed app-boxed-xl"},[e("q-toolbar",{staticClass:"row no-wrap items-center"},[e("div",{staticClass:"q-pr-md logo"},[e("img",{attrs:{alt:"logo",src:j("9b19")}}),M.version?e("q-btn",{staticClass:"btn-menu version",attrs:{type:"a",href:"https://github.com/traefik/traefik/",target:"_blank",stretch:"",flat:"","no-caps":"",label:M.version}}):M._e()],1),e("q-tabs",{attrs:{align:"left","inline-label":"","indicator-color":"transparent","active-color":"white",stretch:""}},[e("q-route-tab",{attrs:{to:"/",icon:"eva-home-outline","no-caps":"",label:"Dashboard"}}),e("q-route-tab",{attrs:{to:"/http",icon:"eva-globe-outline","no-caps":"",label:"HTTP"}}),e("q-route-tab",{attrs:{to:"/tcp",icon:"eva-globe-2-outline","no-caps":"",label:"TCP"}}),e("q-route-tab",{attrs:{to:"/udp",icon:"eva-globe-2-outline","no-caps":"",label:"UDP"}})],1),e("div",{staticClass:"right-menu"},[e("q-tabs",[e("q-btn",{staticClass:"btn-menu",attrs:{stretch:"",flat:"","no-caps":"",icon:"invert_colors",label:(M.$q.dark.isActive?"Light":"Dark")+" theme"},on:{click:function(t){return M.$q.dark.toggle()}}}),e("q-btn",{attrs:{stretch:"",flat:"",icon:"eva-question-mark-circle-outline"}},[e("q-menu",{attrs:{anchor:"bottom left","auto-close":""}},[e("q-item",[e("q-btn",{staticClass:"btn-submenu full-width",attrs:{type:"a",href:"https://doc.traefik.io/traefik/"+M.parsedVersion,target:"_blank",flat:"",color:"accent",align:"left",icon:"eva-book-open-outline","no-caps":"",label:"Documentation"}})],1),e("q-separator"),e("q-item",[e("q-btn",{staticClass:"btn-submenu full-width",attrs:{type:"a",href:"https://github.com/traefik/traefik/",target:"_blank",flat:"",color:"accent",align:"left",icon:"eva-github-outline","no-caps":"",label:"Github repository"}})],1)],1)],1)],1),M.pilotEnabled?e("platform-auth-state"):M._e()],1)],1)],1)]),e("section",{staticClass:"app-section text-black sub-nav",class:{"bg-white":!M.$q.dark.isActive}},[e("div",{staticClass:"app-section-wrap app-boxed app-boxed-xl"},[M._t("default")],2)])])},i=[],r=(j("8e6e"),j("8a81"),j("ac6a"),j("cadf"),j("06db"),j("456d"),j("4917"),j("c47a")),N=j.n(r),n=j("9224"),a=function(){var M=this,t=M.$createElement,j=M._self._c||t;return M.isOnline?j("div",{staticClass:"iframe-wrapper"},[M.renderIrame?j("iframe",{directives:[{name:"resize",rawName:"v-resize",value:M.resizeOpts,expression:"resizeOpts"}],attrs:{id:"platform-auth-state",src:M.iFrameUrl,height:"64px",frameBorder:"0"}}):M._e()]):M._e()},o=[],c=(j("28a5"),j("6762"),j("2fdb"),j("2f62")),s=j("72bf"),y=j.n(s);j("1d6c");function S(M,t){var j=Object.keys(M);if(Object.getOwnPropertySymbols){var e=Object.getOwnPropertySymbols(M);t&&(e=e.filter((function(t){return Object.getOwnPropertyDescriptor(M,t).enumerable}))),j.push.apply(j,e)}return j}function D(M){for(var t=1;t<arguments.length;t++){var j=null!=arguments[t]?arguments[t]:{};t%2?S(j,!0).forEach((function(t){N()(M,t,j[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(M,Object.getOwnPropertyDescriptors(j)):S(j).forEach((function(t){Object.defineProperty(M,t,Object.getOwnPropertyDescriptor(j,t))}))}return M}var g={name:"PlatformPanel",data:function(){var M=this;return{renderIrame:!0,resizeOpts:{log:!1,onMessage:function(t){t.iframe;var j=t.message;"string"===typeof j&&("open:profile"===j?M.openPlatform("/"):j.includes("open:")?M.openPlatform(j.split("open:")[1]):"logout"===j&&M.closePlatform())}}}},created:function(){this.getInstanceInfos()},computed:D({},Object(c["c"])("platform",{isPlatformOpen:"isOpen",platformPath:"path"}),{},Object(c["c"])("core",{instanceInfos:"version"}),{isOnline:function(){return window.navigator.onLine},iFrameUrl:function(){var M=JSON.stringify(this.instanceInfos),t="".concat(window.location.href.split("?")[0],"?platform=").concat(this.platformPath);return y.a.stringifyUrl({url:"".concat(this.platformUrl,"/partials/auth-state"),query:{authRedirectUrl:t,instanceInfos:M}})}}),methods:D({},Object(c["b"])("platform",{openPlatform:"open"},{closePlatform:"close"}),{},Object(c["b"])("core",{getInstanceInfos:"getVersion"})),watch:{isPlatformOpen:function(M,t){var j=this;!M&&t&&(this.renderIrame=!1,this.$nextTick().then((function(){j.renderIrame=!0})))}}},z=g,I=(j("bf59"),j("2877")),O=Object(I["a"])(z,a,o,!1,null,"8c9ba272",null),l=O.exports;function T(M,t){var j=Object.keys(M);if(Object.getOwnPropertySymbols){var e=Object.getOwnPropertySymbols(M);t&&(e=e.filter((function(t){return Object.getOwnPropertyDescriptor(M,t).enumerable}))),j.push.apply(j,e)}return j}function A(M){for(var t=1;t<arguments.length;t++){var j=null!=arguments[t]?arguments[t]:{};t%2?T(j,!0).forEach((function(t){N()(M,t,j[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(M,Object.getOwnPropertyDescriptors(j)):T(j).forEach((function(t){Object.defineProperty(M,t,Object.getOwnPropertyDescriptor(j,t))}))}return M}var x={name:"NavBar",components:{PlatformAuthState:l},computed:A({},Object(c["c"])("core",{coreVersion:"version"}),{version:function(){return this.coreVersion.Version?/^(v?\d+\.\d+)/.test(this.coreVersion.Version)?this.coreVersion.Version:this.coreVersion.Version.substring(0,7):null},pilotEnabled:function(){return this.coreVersion.pilotEnabled},parsedVersion:function(){if(!this.version)return"master";if("dev"===this.version)return"master";var M=this.version.match(/^(v?\d+\.\d+)/);return M?"v"+M[1]:"master"},name:function(){return n.productName}}),methods:A({},Object(c["b"])("core",{getVersion:"getVersion"})),created:function(){this.getVersion()}},E=x,p=(j("a17c"),Object(I["a"])(E,L,i,!1,null,"19d12b1a",null)),w=p.exports,C={name:"Default",components:{NavBar:w},data:function(){return{}}},d=C,f=Object(I["a"])(d,e,u,!1,null,"69aed192",null);t["default"]=f.exports}}]);