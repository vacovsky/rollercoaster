---
draft: false
date: "2017-10-05T15:23:02"
tags: ["LoggingDB",
"ServiceA-DB",
"IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["all-except-unknown"]
depmap: [ "graph LR",
"style Service-A fill:#2ebb53,stroke:#53592a,stroke-width:2px",
"Service-A -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-A -->|SQL|ServiceA-DB(\"fa:fa-database ServiceA-DB\")",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
]
---
			