---
draft: false
date: "2017-10-05T15:23:02"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#3cde8c,stroke:#552523,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
]
---
			