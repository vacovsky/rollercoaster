---
draft: false
date: "2017-10-05T14:31:35"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#d9ae30,stroke:#509148,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-A -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
]
---
			