---
draft: false
date: "2017-10-02T11:37:29"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#33d840,stroke:#27792f,stroke-width:2px",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
]
---
			