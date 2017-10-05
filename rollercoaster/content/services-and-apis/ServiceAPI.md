---
draft: false
date: "2017-10-05T15:23:02"
tags: ["IdentityProvider",
"MessageBus-A",
"Service-A",
"Service-B",
]
title: ServiceAPI
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceAPI fill:#2a419a,stroke:#497199,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-B{\"fa:fa-tasks Service-B\"}",
]
---
			