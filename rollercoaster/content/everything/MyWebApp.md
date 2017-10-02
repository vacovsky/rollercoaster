---
draft: false
date: "2017-10-02T14:57:28"
tags: ["LoggingDB",
"IdentityProvider",
"Service-A",
"ServiceAPI",
"File_Share",
]
title: MyWebApp
categories: ["everything"]
depmap: [ "graph LR",
"style MyWebApp fill:#43bc16,stroke:#4a5f25,stroke-width:2px",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp[\"fa:fa-arrow-right MyWebApp\"] -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
]
---
			