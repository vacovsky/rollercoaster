---
draft: false
date: "2017-10-05T15:23:02"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceA-DB",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#2943b7,stroke:#4e2d62,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-A -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
]
---
			