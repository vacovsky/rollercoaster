---
draft: false
date: "2017-10-05T15:23:02"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
"LoggingDB",
"ServiceB-DB",
]
title: Service-B
categories: ["all-except-unknown"]
depmap: [ "graph LR",
"style Service-B fill:#bf365e,stroke:#5c246b,stroke-width:2px",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-B -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-B -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"Service-B -->|SQL|ServiceB-DB(\"fa:fa-database ServiceB-DB\")",
]
---
			