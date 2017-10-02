---
draft: false
date: "2017-10-02T14:57:28"
tags: ["MessageBus-C",
"LoggingDB",
"ServiceB-DB",
"IdentityProvider",
"MessageBus-B",
]
title: Service-B
categories: ["everything"]
depmap: [ "graph LR",
"style Service-B fill:#1cb460,stroke:#602f1f,stroke-width:2px",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-B -->|SQL|ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
]
---
			