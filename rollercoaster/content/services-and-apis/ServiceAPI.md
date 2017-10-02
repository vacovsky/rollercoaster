---
draft: false
date: "2017-10-02T11:28:22"
tags: ["IdentityProvider",
"MessageBus-A",
"Service-A",
"Service-B",
]
title: ServiceAPI
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceAPI fill:#3cde8c,stroke:#552523,stroke-width:2px",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-B{\"fa:fa-tasks Service-B\"}",
]
---
			