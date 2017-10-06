---
draft: false
date: "2017-10-06T12:02:14"
tags: ["MessageBus-C",
"LoggingDB",
"ServiceA-DB",
"IdentityProvider",
"MessageBus-B",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#43bc16,stroke:#4a5f25,stroke-width:2px",
"Service-A -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-A -->LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
]
---
