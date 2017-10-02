---
draft: false
date: "2017-10-02T11:28:22"
tags: ["MessageBus-B",
"MessageBus-C",
"IdentityProvider",
]
title: Service-B
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style Service-B fill:#d72d53,stroke:#7e6a1b,stroke-width:2px",
"Service-B -->|WCF|MessageBus-B{\"fa:fa-tasks MessageBus-B\"}",
"Service-B -->|WCF|MessageBus-C{\"fa:fa-tasks MessageBus-C\"}",
"Service-B -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			