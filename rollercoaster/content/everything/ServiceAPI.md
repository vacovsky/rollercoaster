---
draft: false
date: "2017-10-05T15:23:02"
tags: ["MessageBus-A",
"Service-A",
"Service-B",
"File_Share",
"LoggingDB",
"ServiceDB",
"IdentityProvider",
]
title: ServiceAPI
categories: ["everything"]
depmap: [ "graph LR",
"style ServiceAPI fill:#9520bb,stroke:#74521f,stroke-width:2px",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|AMQP|MessageBus-A[\"fa:fa-sitemap MessageBus-A\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|WCF|Service-B{\"fa:fa-tasks Service-B\"}",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|SQL|ServiceDB(\"fa:fa-database ServiceDB\")",
"ServiceAPI>\"fa:fa-arrow-right ServiceAPI\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
]
---
			