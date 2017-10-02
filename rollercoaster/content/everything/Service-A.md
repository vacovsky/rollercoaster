---
draft: false
date: "2017-10-02T15:03:19"
tags: ["MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceA-DB",
"IdentityProvider",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#0e9f79,stroke:#34777c,stroke-width:2px",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-A -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			