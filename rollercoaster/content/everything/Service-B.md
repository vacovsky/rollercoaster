---
draft: false
date: "2017-10-02T11:28:22"
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
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
]
---
			