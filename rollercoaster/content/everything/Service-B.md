---
draft: false
date: "2017-10-06T12:02:14"
tags: ["LoggingDB",
"ServiceB-DB",
"IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-B
categories: ["everything"]
depmap: [ "graph LR",
"style Service-B fill:#ceb627,stroke:#482954,stroke-width:2px",
"Service-B -->LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-B -->ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
"Service-B -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-B -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
]
---
