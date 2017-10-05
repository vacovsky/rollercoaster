---
draft: false
date: "2017-10-05T15:23:02"
tags: ["ServiceAPI",
"File_Share",
"LoggingDB",
"IdentityProvider",
"Service-A",
]
title: MyWebApp
categories: ["all-except-unknown"]
depmap: [ "graph LR",
"style MyWebApp fill:#2fa797,stroke:#4b6c21,stroke-width:2px",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->|SMB/CIFS/NFS|File_Share[\"fa:fa-files-o File_Share\"]",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->|SQL|LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->|HTTP|IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->|WCF|Service-A{\"fa:fa-tasks Service-A\"}",
]
---
			