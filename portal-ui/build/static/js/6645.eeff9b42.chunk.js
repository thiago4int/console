"use strict";(self.webpackChunkportal_ui=self.webpackChunkportal_ui||[]).push([[6645],{46645:function(e,n,t){t.r(n);var a=t(23430),s=t(18489),c=t(50390),r=t(34424),i=t(86509),o=t(4285),u=t(72462),l=t(25594),f=t(44149),h=t(28948),p=t(30324),g=t(13089),Z=t(62559),v=(0,r.$j)((function(e){return{loadingTenant:e.tenants.tenantDetails.loadingTenant}}),{setErrorSnackMessage:f.Ih});n.default=(0,o.Z)((function(e){return(0,i.Z)((0,s.Z)((0,s.Z)((0,s.Z)((0,s.Z)({tableWrapper:{height:"450px"}},u.OR),u.qg),u.VX),(0,u.Bz)(e.spacing(4))))}))(v((function(e){var n=e.classes,t=e.match,s=e.loadingTenant,r=e.setErrorSnackMessage,i=(0,c.useState)([]),o=(0,a.Z)(i,2),u=o[0],f=o[1],v=(0,c.useState)(!0),m=(0,a.Z)(v,2),d=m[0],k=m[1],x=t.params.tenantName,E=t.params.tenantNamespace;return(0,c.useEffect)((function(){s&&k(!0)}),[s]),(0,c.useEffect)((function(){d&&p.Z.invoke("GET","/api/v1/namespaces/".concat(E,"/tenants/").concat(x,"/events")).then((function(e){for(var n=0;n<e.length;n++){var t=Date.now()/1e3|0;e[n].seen=(0,h.v1)((t-e[n].last_seen).toString())}f(e),k(!1)})).catch((function(e){r(e),k(!1)}))}),[d,E,x,r]),(0,Z.jsxs)(c.Fragment,{children:[(0,Z.jsx)("h1",{className:n.sectionTitle,children:"Events"}),(0,Z.jsx)(l.ZP,{item:!0,xs:12,children:(0,Z.jsx)(g.Z,{events:u,loading:d})})]})})))}}]);
//# sourceMappingURL=6645.eeff9b42.chunk.js.map