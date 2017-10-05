---
draft: false
date: "2017-10-05T14:31:35"
tags: ["MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceA-DB",
"IdentityProvider",
]
title: Service-A
categories: ["everything"]
depmap: [ "graph LR",
"style Service-A fill:#2943b7,stroke:#4e2d62,stroke-width:2px",
"Service-A -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-A -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			