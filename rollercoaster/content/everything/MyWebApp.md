---
draft: false
date: "2017-10-06T11:07:21"
tags: ["LoggingDB",
"IdentityProvider",
"Service-A",
"ServiceAPI",
"File_Share",
]
title: MyWebApp
categories: ["everything"]
depmap: [ "graph LR",
"style MyWebApp fill:#21babc,stroke:#6f1f35,stroke-width:2px",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->LoggingDB(\"fa:fa-database LoggingDB\")",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->IdentityProvider((\"fa:fa-globe IdentityProvider\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->Service-A{\"fa:fa-tasks Service-A\"}",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->ServiceAPI((\"fa:fa-globe ServiceAPI\"))",
"MyWebApp>\"fa:fa-arrow-right MyWebApp\"] -->File_Share[\"fa:fa-files-o File_Share\"]",
]
---
