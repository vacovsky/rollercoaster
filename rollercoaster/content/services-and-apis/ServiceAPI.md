---
draft: false
date: "2017-10-06T11:07:21"
tags: ["IdentityProvider",
"MessageBus-A",
"Service-A",
"Service-B",
"File_Share",
]
title: ServiceAPI
categories: ["services-and-apis"]
depmap: [ "graph LR",
"style ServiceAPI fill:#2e4caf,stroke:#2e7c52,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->File_Share[\"fa:fa-files-o File_Share\"]",
]
---
