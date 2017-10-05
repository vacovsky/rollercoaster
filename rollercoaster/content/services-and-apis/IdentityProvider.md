---
draft: false
date: "2017-10-05T14:31:35"
tags: ["Crypto",
"MessageBus-A",
]
title: IdentityProvider
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style IdentityProvider fill:#3cca75,stroke:#39906b,stroke-width:2px",
"IdentityProvider -->|HTTP|Crypto((\"fa:fa-globe Crypto\"))",
"IdentityProvider -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
]
---
			