---
draft: false
date: "2017-10-05T15:23:02"
tags: ["ServiceB-DB",
"IdentityProvider",
"MessageBus-B",
"MessageBus-C",
"LoggingDB",
]
title: Service-B
categories: ["everything"]
depmap: [ "graph LR",
"style Service-B fill:#b36431,stroke:#277d78,stroke-width:2px",
"Service-B -->|SQL|ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-B -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-B -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
]
---
			