---
draft: false
date: "2017-10-01T22:40:26"
tags: ["MessageBus-C",
"LoggingDB",
"ServiceB-DB",
"Identity",
"MessageBus-B",
]
title: ServiceB
categories: ["everything"]
depmap: [ "graph LR",
"style ServiceB fill:#b36431,stroke:#277d78,stroke-width:2px",
"ServiceB -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"ServiceB -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"ServiceB -->|SQL|ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
"ServiceB -->|HTTP|Identity((\"fa:fa-globe Identity\"))",
"ServiceB -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
]
---
			