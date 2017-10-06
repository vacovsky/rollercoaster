---
draft: false
date: "2017-10-06T11:03:36"
tags: ["IdentityProvider",
"Service-A",
"Service-B",
"LoggingDB",
"ServiceDB",
]
title: ServiceAPI
categories: ["databases-and-clients"]
depmap: [ "graph LR",
"style ServiceAPI fill:#22b031,stroke:#5c1c28,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->LoggingDB(\"fa:fa-database LoggingDB\")",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->ServiceDB(\"fa:fa-database ServiceDB\")",
]
---
