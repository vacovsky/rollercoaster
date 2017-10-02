---
draft: false
date: "2017-10-01T22:40:26"
tags: ["LoggingDB",
"ServiceA-DB",
"Identity",
"MessageBus-B",
"MessageBus-A",
]
title: ServiceA
categories: ["everything"]
depmap: [ "graph LR",
"style ServiceA fill:#2943b7,stroke:#4e2d62,stroke-width:2px",
"ServiceA -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"ServiceA -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"ServiceA -->|HTTP|Identity((\"fa:fa-globe Identity\"))",
"ServiceA -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"ServiceA -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
]
---
			