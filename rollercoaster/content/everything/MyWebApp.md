---
draft: false
date: "2017-10-06T12:02:14"
tags: ["ServiceAPI",
"File_Share",
"LoggingDB",
"IdentityProvider",
"Service-A",
]
title: MyWebApp
categories: ["everything"]
depmap: [ "graph LR",
"style MyWebApp fill:#b36431,stroke:#277d78,stroke-width:2px",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->File_Share[\"fa:fa-files-o File_Share\"]",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
]
---
