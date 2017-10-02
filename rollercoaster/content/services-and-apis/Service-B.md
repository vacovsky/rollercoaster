---
draft: false
date: "2017-10-02T14:57:28"
tags: ["MessageBus-C",
"IdentityProvider",
"MessageBus-B",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#10b864,stroke:#4e1137,stroke-width:2px",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
]
---
			