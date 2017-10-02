---
draft: false
date: "2017-10-02T15:03:19"
tags: ["MessageBus-A",
"Service-A",
"Service-B",
"IdentityProvider",
]
title: ServiceAPI
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceAPI fill:#2a419a,stroke:#497199,stroke-width:2px",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|MessageBus-A{\"fa:fa-tasks MessageBus-A\"}",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI[\"fa:fa-arrow-right ServiceAPI\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			