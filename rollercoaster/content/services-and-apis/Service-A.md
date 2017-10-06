---
draft: false
date: "2017-10-06T12:02:14"
tags: ["MessageBus-B",
"MessageBus-C",
"IdentityProvider",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#2614c1,stroke:#242a7d,stroke-width:2px",
"Service-A -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-A -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
