---
draft: false
date: "2017-10-02T11:37:29"
tags: ["IdentityProvider",
"MessageBus-B",
"MessageBus-C",
]
title: Service-A
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-A fill:#10b864,stroke:#4e1137,stroke-width:2px",
"Service-A -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-A -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-A -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
]
---
			