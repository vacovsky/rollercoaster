---
draft: false
date: "2017-10-06T12:02:14"
tags: ["MessageBus-C",
"IdentityProvider",
"MessageBus-B",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#55d32c,stroke:#6a821f,stroke-width:2px",
"Service-B -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-B -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
]
---
