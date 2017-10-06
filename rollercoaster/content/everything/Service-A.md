---
draft: false
date: "2017-10-06T11:07:21"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceA-DB",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#c9ae15,stroke:#3c7852,stroke-width:2px",
"Service-A -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-A -->LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
]
---
