---
draft: false
date: "2017-10-02T14:57:28"
tags: ["MessageBus-B",
"MessageBus-C",
"IdentityProvider",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#dc6a12,stroke:#8c2231,stroke-width:2px",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			