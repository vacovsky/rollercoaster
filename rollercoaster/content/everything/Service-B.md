---
draft: false
date: "2017-10-06T11:07:21"
tags: ["MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceB-DB",
"IdentityProvider",
]
title: Service-B
categories: ["everything"]
depmap: [ "graph LR",
"style Service-B fill:#4fcc06,stroke:#1b805e,stroke-width:2px",
"Service-B -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-B -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-B -->LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-B -->ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
"Service-B -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
