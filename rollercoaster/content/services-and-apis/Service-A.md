---
draft: false
date: "2017-10-02T11:28:22"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#d9ae30,stroke:#509148,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
]
---
			