---
draft: false
date: "2017-10-05T14:31:35"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#d72d53,stroke:#7e6a1b,stroke-width:2px",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
"Service-B -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
]
---
			