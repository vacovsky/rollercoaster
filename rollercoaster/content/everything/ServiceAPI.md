---
draft: false
date: "2017-10-06T11:07:22"
tags: ["Service-B",
"File_Share",
"LoggingDB",
"ServiceDB",
"IdentityProvider",
"MessageBus-A",
"Service-A",
]
title: ServiceAPI
categories: ["everything"]
depmap: [ "graph LR",
"style ServiceAPI fill:#ceb627,stroke:#482954,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->File_Share[\"fa:fa-files-o File_Share\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->LoggingDB(\"fa:fa-database LoggingDB\")",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->ServiceDB(\"fa:fa-database ServiceDB\")",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
]
---
