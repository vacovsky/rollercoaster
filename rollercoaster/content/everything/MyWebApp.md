---
draft: false
date: "2017-10-02T15:03:19"
tags: ["IdentityProvider",
"Service-A",
"ServiceAPI",
"File_Share",
"LoggingDB",
]
title: MyWebApp
categories: ["everything"]
depmap: [ "graph LR",
"style MyWebApp fill:#1cb460,stroke:#602f1f,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
]
---
			