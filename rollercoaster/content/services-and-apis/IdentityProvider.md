---
draft: false
date: "2017-10-05T15:23:02"
tags: ["Crypto",
"MessageBus-A",
]
title: IdentityProvider
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style IdentityProvider fill:#33d840,stroke:#27792f,stroke-width:2px",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
]
---
			