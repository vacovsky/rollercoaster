---
draft: false
date: "2017-10-02T14:57:28"
tags: ["MessageBus-C",
"LoggingDB",
"ServiceA-DB",
"IdentityProvider",
"MessageBus-B",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#ceb627,stroke:#482954,stroke-width:2px",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-A -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
]
---
			