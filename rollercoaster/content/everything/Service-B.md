---
draft: false
date: "2017-10-02T15:03:19"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceB-DB",
]
title: Service-B
categories: ["everything"]
depmap: [ "graph LR",
"style Service-B fill:#2943b7,stroke:#4e2d62,stroke-width:2px",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-B -->|SQL|ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
]
---
			