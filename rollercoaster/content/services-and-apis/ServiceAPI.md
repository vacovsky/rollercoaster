---
draft: false
date: "2017-10-06T12:02:14"
tags: ["IdentityProvider",
"MessageBus-A",
"Service-A",
"Service-B",
"File_Share",
]
title: ServiceAPI
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceAPI fill:#e41695,stroke:#723652,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->File_Share[\"fa:fa-files-o File_Share\"]",
]
---
