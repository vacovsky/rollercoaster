---
draft: false
date: "2017-10-05T15:23:02"
tags: ["MessageBus-C",
"IdentityProvider",
"MessageBus-B",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#76e40a,stroke:#809549,stroke-width:2px",
"Service-B -->|AMQP|MessageBus-C[\"fa:fa-sitemap MessageBus-C\"]",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|AMQP|MessageBus-B[\"fa:fa-sitemap MessageBus-B\"]",
]
---
			